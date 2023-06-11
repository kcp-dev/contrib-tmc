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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SyncTargetSpecApplyConfiguration represents an declarative configuration of the SyncTargetSpec type for use
// with apply.
type SyncTargetSpecApplyConfiguration struct {
	Unschedulable       *bool                         `json:"unschedulable,omitempty"`
	EvictAfter          *v1.Time                      `json:"evictAfter,omitempty"`
	SupportedAPIExports []v1alpha1.APIExportReference `json:"supportedAPIExports,omitempty"`
	Cells               map[string]string             `json:"cells,omitempty"`
}

// SyncTargetSpecApplyConfiguration constructs an declarative configuration of the SyncTargetSpec type for use with
// apply.
func SyncTargetSpec() *SyncTargetSpecApplyConfiguration {
	return &SyncTargetSpecApplyConfiguration{}
}

// WithUnschedulable sets the Unschedulable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Unschedulable field is set to the value of the last call.
func (b *SyncTargetSpecApplyConfiguration) WithUnschedulable(value bool) *SyncTargetSpecApplyConfiguration {
	b.Unschedulable = &value
	return b
}

// WithEvictAfter sets the EvictAfter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EvictAfter field is set to the value of the last call.
func (b *SyncTargetSpecApplyConfiguration) WithEvictAfter(value v1.Time) *SyncTargetSpecApplyConfiguration {
	b.EvictAfter = &value
	return b
}

// WithSupportedAPIExports adds the given value to the SupportedAPIExports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SupportedAPIExports field.
func (b *SyncTargetSpecApplyConfiguration) WithSupportedAPIExports(values ...v1alpha1.APIExportReference) *SyncTargetSpecApplyConfiguration {
	for i := range values {
		b.SupportedAPIExports = append(b.SupportedAPIExports, values[i])
	}
	return b
}

// WithCells puts the entries into the Cells field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Cells field,
// overwriting an existing map entries in Cells field with the same key.
func (b *SyncTargetSpecApplyConfiguration) WithCells(entries map[string]string) *SyncTargetSpecApplyConfiguration {
	if b.Cells == nil && len(entries) > 0 {
		b.Cells = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Cells[k] = v
	}
	return b
}
