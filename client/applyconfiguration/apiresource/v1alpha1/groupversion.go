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

// GroupVersionApplyConfiguration represents an declarative configuration of the GroupVersion type for use
// with apply.
type GroupVersionApplyConfiguration struct {
	Group   *string `json:"group,omitempty"`
	Version *string `json:"version,omitempty"`
}

// GroupVersionApplyConfiguration constructs an declarative configuration of the GroupVersion type for use with
// apply.
func GroupVersion() *GroupVersionApplyConfiguration {
	return &GroupVersionApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *GroupVersionApplyConfiguration) WithGroup(value string) *GroupVersionApplyConfiguration {
	b.Group = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *GroupVersionApplyConfiguration) WithVersion(value string) *GroupVersionApplyConfiguration {
	b.Version = &value
	return b
}
