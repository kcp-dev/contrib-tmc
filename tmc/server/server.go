/*
Copyright 2021 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"errors"
	"fmt"
	_ "net/http/pprof"
	"time"

	coreserver "github.com/kcp-dev/kcp/pkg/server"
	"github.com/kcp-dev/kcp/sdk/apis/core"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"

	"k8s.io/apimachinery/pkg/util/sets"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/klog/v2"

	"github.com/kcp-dev/contrib-tmc/config/root"
	configrootcompute "github.com/kcp-dev/contrib-tmc/config/rootcompute"
	configtmc "github.com/kcp-dev/contrib-tmc/config/tmc"
	"github.com/kcp-dev/contrib-tmc/tmc/server/bootstrap"
)

const resyncPeriod = 10 * time.Hour

type Server struct {
	CompletedConfig

	Core *coreserver.Server

	syncedPhase1Ch chan struct{}
	syncedPhase2Ch chan struct{}
}

func NewServer(c CompletedConfig) (*Server, error) {
	core, err := coreserver.NewServer(c.Core)
	if err != nil {
		return nil, err
	}

	s := &Server{
		CompletedConfig: c,
		Core:            core,
		// phase1 - crds, apiresourceschemas, workspaces
		syncedPhase1Ch: make(chan struct{}),
		// phase2 - informers started and running
		syncedPhase2Ch: make(chan struct{}),
	}

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	logger := klog.FromContext(ctx).WithValues("component", "tmc")
	ctx = klog.NewContext(ctx, logger)

	cacheHookName := "tmc-populate-cache-server"
	if err := s.Core.AddPostStartHook(cacheHookName, func(hookContext genericapiserver.PostStartHookContext) error {
		logger := logger.WithValues("postStartHook", cacheHookName)

		err := s.Core.WaitForSync(hookContext.StopCh)
		if err != nil {
			logger.Error(err, "failed to wait for sync")
			return nil
		}

		if err := bootstrap.Bootstrap(klog.NewContext(goContext(hookContext), logger), s.Core.ApiExtensionsClusterClient); err != nil {
			logger.Error(err, "failed creating the static CustomResourcesDefinitions")
			return nil // don't klog.Fatal. This only happens when context is cancelled.
		}
		close(s.syncedPhase1Ch)
		return nil
	}); err != nil {
		return err
	}

	hookName := "tmc-start-informers"
	if err := s.Core.AddPostStartHook(hookName, func(hookContext genericapiserver.PostStartHookContext) error {
		logger := logger.WithValues("postStartHook", hookName)
		ctx = klog.NewContext(ctx, logger)

		err := s.WaitForSyncPhase1(hookContext.StopCh)
		if err != nil {
			logger.Error(err, "failed to wait for phase1 sync")
			return err
		}

		// Poke the informers to start syncing
		s.TmcSharedInformerFactory.Workload().V1alpha1().SyncTargets()

		logger.Info("starting tmc informers")
		s.TmcSharedInformerFactory.Start(hookContext.StopCh)

		for v, synced := range s.TmcSharedInformerFactory.WaitForCacheSync(hookContext.StopCh) {
			if !synced {
				logger.Error(nil, "Error syncing informer", "informer", v)
				return fmt.Errorf("failed to sync informer %s", v)
			}
			logger.Info("synced informer", "informer", v)
		}

		logger.Info("synced all TMC informers")
		close(s.syncedPhase2Ch)

		select {
		case <-hookContext.StopCh:
			return nil // context closed, avoid reporting success below
		default:
		}

		logger.Info("finished starting tmc informers")
		return nil
	}); err != nil {
		return err
	}

	tmcBootstrapHook := "tmcBootstrap"
	if err := s.Core.AddPostStartHook(tmcBootstrapHook, func(hookContext genericapiserver.PostStartHookContext) error {
		logger := logger.WithValues("postStartHook", tmcBootstrapHook)
		err := s.WaitForSyncPhase1(hookContext.StopCh)
		if err != nil {
			logger.Error(err, "failed to wait for sync")
			return nil
		}
		if s.Core.Options.Extra.ShardName == corev1alpha1.RootShard {
			// the root ws is only present on the root shard
			logger.Info("starting bootstrapping root tmc assets")
			if err := configtmc.Bootstrap(goContext(hookContext),
				s.Core.BootstrapApiExtensionsClusterClient,
				s.Core.BootstrapDynamicClusterClient,
				sets.New[string](s.Core.Options.Extra.BatteriesIncluded...),
			); err != nil {
				logger.Error(err, "failed to bootstrap root tmc assets")
				return nil // don't klog.Fatal. This only tmc when context is cancelled.
			}
			logger.Info("finished bootstrapping root tmc assets")

			logger.Info("starting bootstrapping root workspace rbac")
			if err := root.Bootstrap(
				goContext(hookContext),
				s.Core.BootstrapApiExtensionsClusterClient.Cluster(core.RootCluster.Path()).Discovery(),
				s.Core.BootstrapDynamicClusterClient.Cluster(core.RootCluster.Path()),
				s.Core.Options.HomeWorkspaces.HomeCreatorGroups,
				sets.New[string](s.Core.Options.Extra.BatteriesIncluded...),
			); err != nil {
				logger.Error(err, "failed to bootstrap root workspace rbac")
				return nil // don't klog.Fatal. This only happens when context is cancelled.
			}
			logger.Info("finished bootstrapping root workspace rbac")
		}
		return nil
	}); err != nil {
		return err
	}

	// TODO(marun) Consider enabling each controller via a separate flag
	if err := s.installSyncTargetHeartbeatController(ctx); err != nil {
		return err
	}
	if err := s.installSyncTargetController(ctx); err != nil {
		return err
	}
	if err := s.installSchedulingLocationStatusController(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadResourceScheduler(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadSyncTargetExportController(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadReplicateClusterRoleControllers(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadReplicateClusterRoleBindingControllers(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadReplicateLogicalClusterControllers(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadNamespaceScheduler(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadPlacementScheduler(ctx); err != nil {
		return err
	}
	if err := s.installSchedulingPlacementController(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadAPIExportController(ctx); err != nil {
		return err
	}
	if err := s.installWorkloadDefaultLocationController(ctx); err != nil {
		return err
	}

	// bootstrap root compute workspace
	// not part of phases as it will be needed only when somebody starts consuming compute workspaces
	computeBootstrapHookName := "rootComputeBootstrap"
	if err := s.Core.AddPostStartHook(computeBootstrapHookName, func(hookContext genericapiserver.PostStartHookContext) error {
		logger := logger.WithValues("postStartHook", computeBootstrapHookName)
		err := s.Core.WaitForSync(hookContext.StopCh)
		if err != nil {
			logger.Error(err, "failed to wait for sync")
			return nil
		}

		if s.Core.Options.Extra.ShardName == corev1alpha1.RootShard {
			// the root ws is only present on the root shard

			logger.Info("starting bootstrapping root compute workspace")
			if err := configrootcompute.Bootstrap(goContext(hookContext),
				s.Core.BootstrapApiExtensionsClusterClient,
				s.Core.BootstrapDynamicClusterClient,
				sets.New[string](s.Core.Options.Extra.BatteriesIncluded...),
			); err != nil {
				logger.Error(err, "failed to bootstrap root compute workspace")
				return nil // don't klog.Fatal. This only happens when context is cancelled.
			}
			logger.Info("finished bootstrapping root compute workspace")
		}
		return nil
	}); err != nil {
		return err
	}

	return s.Core.Run(ctx)
}

// goContext turns the PostStartHookContext into a context.Context for use in routines that may or may not
// run inside of a post-start-hook. The k8s APIServer wrote the post-start-hook context code before contexts
// were part of the Go stdlib.
func goContext(parent genericapiserver.PostStartHookContext) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func(done <-chan struct{}) {
		<-done
		cancel()
	}(parent.StopCh)
	return ctx
}

func (s *Server) WaitForSyncPhase1(stop <-chan struct{}) error {
	// Wait for shared informer factories to by synced.
	// factory. Otherwise, informer list calls may go into backoff (before the CRDs are ready) and
	// take ~10 seconds to succeed.
	select {
	case <-stop:
		return errors.New("timed out waiting for core resources to sync")
	case <-s.syncedPhase1Ch:
		return nil
	}
}

func (s *Server) WaitForSyncPhase2(stop <-chan struct{}) error {
	// Wait for shared informer factories to by synced.
	// factory. Otherwise, informer list calls may go into backoff (before the CRDs are ready) and
	// take ~10 seconds to succeed.
	select {
	case <-stop:
		return errors.New("timed out waiting for informers to sync")
	case <-s.syncedPhase2Ch:
		return nil
	}
}
