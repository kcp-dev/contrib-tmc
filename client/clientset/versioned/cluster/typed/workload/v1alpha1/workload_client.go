//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/rest"

	workloadv1alpha1 "github.com/faroshq/tmc/client/clientset/versioned/typed/workload/v1alpha1"
)

type WorkloadV1alpha1ClusterInterface interface {
	WorkloadV1alpha1ClusterScoper
	SyncTargetsClusterGetter
}

type WorkloadV1alpha1ClusterScoper interface {
	Cluster(logicalcluster.Path) workloadv1alpha1.WorkloadV1alpha1Interface
}

type WorkloadV1alpha1ClusterClient struct {
	clientCache kcpclient.Cache[*workloadv1alpha1.WorkloadV1alpha1Client]
}

func (c *WorkloadV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) workloadv1alpha1.WorkloadV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *WorkloadV1alpha1ClusterClient) SyncTargets() SyncTargetClusterInterface {
	return &syncTargetsClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new WorkloadV1alpha1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*WorkloadV1alpha1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new WorkloadV1alpha1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*WorkloadV1alpha1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*workloadv1alpha1.WorkloadV1alpha1Client]{
		NewForConfigAndClient: workloadv1alpha1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &WorkloadV1alpha1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new WorkloadV1alpha1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WorkloadV1alpha1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
