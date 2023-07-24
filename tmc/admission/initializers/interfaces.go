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

package initializers

import (
	tmcclientset "github.com/kcp-dev/contrib-tmc/client/clientset/versioned/cluster"
	tmcinformers "github.com/kcp-dev/contrib-tmc/client/informers/externalversions"
)

// WantsTmcInformers interface should be implemented by admission plugins
// that want to have both local and global tmc informer factories injected.
type WantsTmcInformers interface {
	SetTmcInformers(local, global tmcinformers.SharedInformerFactory)
}

// WantsTmcClusterClient interface should be implemented by admission plugins
// that want to have a tmc cluster client injected.
type WantsTmcClusterClient interface {
	SetTmcClusterClient(tmcclientset.ClusterInterface)
}
