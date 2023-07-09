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

package clientset

import (
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

	client "github.com/kcp-dev/contrib-tmc/client/clientset/versioned"
	schedulingv1alpha1 "github.com/kcp-dev/contrib-tmc/client/clientset/versioned/cluster/typed/scheduling/v1alpha1"
	workloadv1alpha1 "github.com/kcp-dev/contrib-tmc/client/clientset/versioned/cluster/typed/workload/v1alpha1"
)

type ClusterInterface interface {
	Cluster(logicalcluster.Path) client.Interface
	Discovery() discovery.DiscoveryInterface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1ClusterInterface
	WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1ClusterInterface
}

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*discovery.DiscoveryClient
	clientCache        kcpclient.Cache[*client.Clientset]
	schedulingV1alpha1 *schedulingv1alpha1.SchedulingV1alpha1ClusterClient
	workloadV1alpha1   *workloadv1alpha1.WorkloadV1alpha1ClusterClient
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1ClusterClient.
func (c *ClusterClientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1ClusterInterface {
	return c.schedulingV1alpha1
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1ClusterClient.
func (c *ClusterClientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1ClusterInterface {
	return c.workloadV1alpha1
}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

// NewForConfig creates a new ClusterClientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new ClusterClientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[*client.Clientset]{
		NewForConfigAndClient: client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	var cs ClusterClientset
	cs.clientCache = cache
	var err error
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.workloadV1alpha1, err = workloadv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new ClusterClientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterClientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}
