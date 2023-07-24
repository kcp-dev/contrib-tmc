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
	"k8s.io/apiserver/pkg/admission"

	tmcclientset "github.com/kcp-dev/contrib-tmc/client/clientset/versioned/cluster"
	tmcinformers "github.com/kcp-dev/contrib-tmc/client/informers/externalversions"
)

// NewTmcInformersInitializer returns an admission plugin initializer that injects
// both local and global tmc shared informer factories into admission plugins.
func NewTmcInformersInitializer(
	local, global tmcinformers.SharedInformerFactory,
) *tmcInformersInitializer {
	return &tmcInformersInitializer{
		localTmcInformers:  local,
		globalTmcInformers: global,
	}
}

type tmcInformersInitializer struct {
	localTmcInformers, globalTmcInformers tmcinformers.SharedInformerFactory
}

func (i *tmcInformersInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsTmcInformers); ok {
		wants.SetTmcInformers(i.localTmcInformers, i.globalTmcInformers)
	}
}

// NewTmcClusterClientInitializer returns an admission plugin initializer that injects
// a tmc cluster client into admission plugins.
func NewTmcClusterClientInitializer(
	tmcClusterClient tmcclientset.ClusterInterface,
) *tmcClusterClientInitializer {
	return &tmcClusterClientInitializer{
		tmcClusterClient: tmcClusterClient,
	}
}

type tmcClusterClientInitializer struct {
	tmcClusterClient tmcclientset.ClusterInterface
}

func (i *tmcClusterClientInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsTmcClusterClient); ok {
		wants.SetTmcClusterClient(i.tmcClusterClient)
	}
}
