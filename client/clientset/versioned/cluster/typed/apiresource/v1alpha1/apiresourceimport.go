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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	apiresourcev1alpha1 "github.com/kcp-dev/contrib-tmc/apis/apiresource/v1alpha1"
	apiresourcev1alpha1client "github.com/kcp-dev/contrib-tmc/client/clientset/versioned/typed/apiresource/v1alpha1"
)

// APIResourceImportsClusterGetter has a method to return a APIResourceImportClusterInterface.
// A group's cluster client should implement this interface.
type APIResourceImportsClusterGetter interface {
	APIResourceImports() APIResourceImportClusterInterface
}

// APIResourceImportClusterInterface can operate on APIResourceImports across all clusters,
// or scope down to one cluster and return a apiresourcev1alpha1client.APIResourceImportInterface.
type APIResourceImportClusterInterface interface {
	Cluster(logicalcluster.Path) apiresourcev1alpha1client.APIResourceImportInterface
	List(ctx context.Context, opts metav1.ListOptions) (*apiresourcev1alpha1.APIResourceImportList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type aPIResourceImportsClusterInterface struct {
	clientCache kcpclient.Cache[*apiresourcev1alpha1client.ApiresourceV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *aPIResourceImportsClusterInterface) Cluster(clusterPath logicalcluster.Path) apiresourcev1alpha1client.APIResourceImportInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).APIResourceImports()
}

// List returns the entire collection of all APIResourceImports across all clusters.
func (c *aPIResourceImportsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*apiresourcev1alpha1.APIResourceImportList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResourceImports().List(ctx, opts)
}

// Watch begins to watch all APIResourceImports across all clusters.
func (c *aPIResourceImportsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResourceImports().Watch(ctx, opts)
}
