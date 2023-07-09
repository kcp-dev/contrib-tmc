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

package indexers

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	workloadv1alpha1 "github.com/kcp-dev/contrib-tmc/apis/workload/v1alpha1"
	syncershared "github.com/kcp-dev/contrib-tmc/pkg/syncer/shared"
)

const (
	// BySyncerFinalizerKey is the name for the index that indexes by syncer finalizer label keys.
	BySyncerFinalizerKey = "bySyncerFinalizerKey"
	// ByClusterResourceStateLabelKey indexes resources based on the cluster state label key.
	ByClusterResourceStateLabelKey = "ByClusterResourceStateLabelKey"
)

// IndexBySyncerFinalizerKey indexes by syncer finalizer label keys.
func IndexBySyncerFinalizerKey(obj interface{}) ([]string, error) {
	metaObj, ok := obj.(metav1.Object)
	if !ok {
		return []string{}, fmt.Errorf("obj is supposed to be a metav1.Object, but is %T", obj)
	}

	syncerFinalizers := []string{}
	for _, f := range metaObj.GetFinalizers() {
		if strings.HasPrefix(f, syncershared.SyncerFinalizerNamePrefix) {
			syncerFinalizers = append(syncerFinalizers, f)
		}
	}

	return syncerFinalizers, nil
}

// IndexByClusterResourceStateLabelKey indexes resources based on the cluster state key label.
func IndexByClusterResourceStateLabelKey(obj interface{}) ([]string, error) {
	metaObj, ok := obj.(metav1.Object)
	if !ok {
		return []string{}, fmt.Errorf("obj is supposed to be a metav1.Object, but is %T", obj)
	}

	ClusterResourceStateLabelKeys := []string{}
	for k := range metaObj.GetLabels() {
		if strings.HasPrefix(k, workloadv1alpha1.ClusterResourceStateLabelPrefix) {
			ClusterResourceStateLabelKeys = append(ClusterResourceStateLabelKeys, k)
		}
	}
	return ClusterResourceStateLabelKeys, nil
}
