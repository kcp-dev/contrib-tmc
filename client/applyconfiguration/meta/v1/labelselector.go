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

package v1

// LabelSelectorApplyConfiguration represents an declarative configuration of the LabelSelector type for use
// with apply.
type LabelSelectorApplyConfiguration struct {
	MatchLabels      map[string]string                            `json:"matchLabels,omitempty"`
	MatchExpressions []LabelSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
}

// LabelSelectorApplyConfiguration constructs an declarative configuration of the LabelSelector type for use with
// apply.
func LabelSelector() *LabelSelectorApplyConfiguration {
	return &LabelSelectorApplyConfiguration{}
}

// WithMatchLabels puts the entries into the MatchLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the MatchLabels field,
// overwriting an existing map entries in MatchLabels field with the same key.
func (b *LabelSelectorApplyConfiguration) WithMatchLabels(entries map[string]string) *LabelSelectorApplyConfiguration {
	if b.MatchLabels == nil && len(entries) > 0 {
		b.MatchLabels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.MatchLabels[k] = v
	}
	return b
}

// WithMatchExpressions adds the given value to the MatchExpressions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchExpressions field.
func (b *LabelSelectorApplyConfiguration) WithMatchExpressions(values ...*LabelSelectorRequirementApplyConfiguration) *LabelSelectorApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchExpressions")
		}
		b.MatchExpressions = append(b.MatchExpressions, *values[i])
	}
	return b
}
