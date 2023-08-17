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

package admission

import (
	"k8s.io/apiserver/pkg/admission"
	mutatingwebhook "k8s.io/apiserver/pkg/admission/plugin/webhook/mutating"

	"github.com/kcp-dev/contrib-tmc/pkg/admission/pathannotation"
)

// tmcOrderedPlugins is the list of TMC plugins in order.
var tmcOrderedPlugins = []string{
	pathannotation.PluginName,
}

func beforeWebhooks(currents []string, plugins []string) []string {
	ret := make([]string, 0, len(currents)+len(plugins))
	for _, plugin := range currents {
		if plugin == mutatingwebhook.PluginName {
			ret = append(ret, plugins...)
		}
		ret = append(ret, plugin)
	}
	return ret
}

func AddTMCOrderedPlugins(currents []string) []string {
	return beforeWebhooks(currents, tmcOrderedPlugins)
}

// RegisterAllTMCAdmissionPlugins registers all admission plugins.
// The order of registration is irrelevant, see AllOrderedPlugins for execution order.
func RegisterAllTMCAdmissionPlugins(plugins *admission.Plugins) {
	pathannotation.Register(plugins)
}
