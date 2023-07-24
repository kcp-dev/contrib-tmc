/*
Copyright 2022 The KCP Authors.

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

package tmc

import (
	"context"
	"embed"

	kcpdynamic "github.com/kcp-dev/client-go/dynamic"
	confighelpers "github.com/kcp-dev/kcp/config/helpers"
	"github.com/kcp-dev/kcp/sdk/apis/core"
	"github.com/kcp-dev/logicalcluster/v3"

	kcpapiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/kcp/clientset/versioned"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/kcp-dev/contrib-tmc/config/tmc/resources"
)

//go:embed *.yaml
var fs embed.FS

// RootTMCClusterName is the workspace to host common tmc APIs.
var RootTMCClusterName = logicalcluster.NewPath("root:tmc")

// Bootstrap creates resources in this package by continuously retrying the list.
// This is blocking, i.e. it only returns (with error) when the context is closed or with nil when
// the bootstrapping is successfully completed.
func Bootstrap(ctx context.Context, apiExtensionClusterClient kcpapiextensionsclientset.ClusterInterface, dynamicClusterClient kcpdynamic.ClusterInterface, batteriesIncluded sets.Set[string]) error {
	rootDiscoveryClient := apiExtensionClusterClient.Cluster(core.RootCluster.Path()).Discovery()
	rootDynamicClient := dynamicClusterClient.Cluster(core.RootCluster.Path())
	if err := confighelpers.Bootstrap(ctx, rootDiscoveryClient, rootDynamicClient, batteriesIncluded, fs); err != nil {
		return err
	}

	computeDiscoveryClient := apiExtensionClusterClient.Cluster(RootTMCClusterName).Discovery()
	computeDynamicClient := dynamicClusterClient.Cluster(RootTMCClusterName)

	return resources.Bootstrap(ctx, computeDiscoveryClient, computeDynamicClient, batteriesIncluded)
}
