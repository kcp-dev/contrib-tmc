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

package options

import (
	kcpcoreoptions "github.com/kcp-dev/kcp/pkg/server/options"

	cliflag "k8s.io/component-base/cli/flag"

	tmcadmission "github.com/kcp-dev/contrib-tmc/pkg/admission"
	tmcvirtualoptions "github.com/kcp-dev/contrib-tmc/tmc/virtual/options"
)

type Options struct {
	Core                 kcpcoreoptions.Options
	TmcControllers       Controllers
	TmcVirtualWorkspaces tmcvirtualoptions.Options

	Extra ExtraOptions
}

type ExtraOptions struct {
}

type completedOptions struct {
	Core                 kcpcoreoptions.CompletedOptions
	Controllers          Controllers
	TmcVirtualWorkspaces tmcvirtualoptions.Options

	Extra ExtraOptions
}

type CompletedOptions struct {
	*completedOptions
}

// NewOptions creates a new Options with default parameters.
func NewOptions(rootDir string) *Options {
	o := &Options{
		Core:                 *kcpcoreoptions.NewOptions(rootDir),
		TmcControllers:       *NewTmcControllers(),
		TmcVirtualWorkspaces: *tmcvirtualoptions.NewOptions(),

		Extra: ExtraOptions{},
	}
	// add TMC admission plugins
	tmcadmission.RegisterAllTMCAdmissionPlugins(o.Core.GenericControlPlane.Admission.Plugins)
	orderedPlugins := tmcadmission.AddTMCOrderedPlugins(o.Core.GenericControlPlane.Admission.RecommendedPluginOrder)
	o.Core.GenericControlPlane.Admission.RecommendedPluginOrder = orderedPlugins

	return o
}

func (o *Options) AddFlags(fss *cliflag.NamedFlagSets) {
	o.Core.AddFlags(fss)
	o.TmcControllers.AddFlags(fss.FlagSet("TMC Controllers"))
	o.TmcVirtualWorkspaces.AddFlags(fss.FlagSet("TMC Virtual Workspaces"))
}

func (o *CompletedOptions) Validate() []error {
	var errs []error

	errs = append(errs, o.Core.Validate()...)
	errs = append(errs, o.Controllers.Validate()...)

	return errs
}

func (o *Options) Complete(rootDir string) (*CompletedOptions, error) {
	core, err := o.Core.Complete(rootDir)
	if err != nil {
		return nil, err
	}
	if err := o.TmcControllers.Complete(rootDir); err != nil {
		return nil, err
	}

	return &CompletedOptions{
		completedOptions: &completedOptions{
			Core:                 *core,
			Controllers:          o.TmcControllers,
			TmcVirtualWorkspaces: o.TmcVirtualWorkspaces,
			Extra:                o.Extra,
		},
	}, nil
}
