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
	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apiresource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NegotiatedAPIResourceConditionApplyConfiguration represents an declarative configuration of the NegotiatedAPIResourceCondition type for use
// with apply.
type NegotiatedAPIResourceConditionApplyConfiguration struct {
	Type               *v1alpha1.NegotiatedAPIResourceConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                          `json:"status,omitempty"`
	LastTransitionTime *v1.Time                                     `json:"lastTransitionTime,omitempty"`
	Reason             *string                                      `json:"reason,omitempty"`
	Message            *string                                      `json:"message,omitempty"`
}

// NegotiatedAPIResourceConditionApplyConfiguration constructs an declarative configuration of the NegotiatedAPIResourceCondition type for use with
// apply.
func NegotiatedAPIResourceCondition() *NegotiatedAPIResourceConditionApplyConfiguration {
	return &NegotiatedAPIResourceConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *NegotiatedAPIResourceConditionApplyConfiguration) WithType(value v1alpha1.NegotiatedAPIResourceConditionType) *NegotiatedAPIResourceConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NegotiatedAPIResourceConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *NegotiatedAPIResourceConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *NegotiatedAPIResourceConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *NegotiatedAPIResourceConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *NegotiatedAPIResourceConditionApplyConfiguration) WithReason(value string) *NegotiatedAPIResourceConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *NegotiatedAPIResourceConditionApplyConfiguration) WithMessage(value string) *NegotiatedAPIResourceConditionApplyConfiguration {
	b.Message = &value
	return b
}
