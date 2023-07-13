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
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	apiresourcev1alpha1 "github.com/kcp-dev/contrib-tmc/apis/apiresource/v1alpha1"
)

// NegotiatedAPIResourceSpecApplyConfiguration represents an declarative configuration of the NegotiatedAPIResourceSpec type for use
// with apply.
type NegotiatedAPIResourceSpecApplyConfiguration struct {
	CommonAPIResourceSpecApplyConfiguration `json:",inline"`
	Publish                                 *bool `json:"publish,omitempty"`
}

// NegotiatedAPIResourceSpecApplyConfiguration constructs an declarative configuration of the NegotiatedAPIResourceSpec type for use with
// apply.
func NegotiatedAPIResourceSpec() *NegotiatedAPIResourceSpecApplyConfiguration {
	return &NegotiatedAPIResourceSpecApplyConfiguration{}
}

// WithGroupVersion sets the GroupVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupVersion field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithGroupVersion(value *GroupVersionApplyConfiguration) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.GroupVersion = value
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithScope(value v1.ResourceScope) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.Scope = &value
	return b
}

// WithPlural sets the Plural field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Plural field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithPlural(value string) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.Plural = &value
	return b
}

// WithSingular sets the Singular field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Singular field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithSingular(value string) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.Singular = &value
	return b
}

// WithShortNames adds the given value to the ShortNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ShortNames field.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithShortNames(values ...string) *NegotiatedAPIResourceSpecApplyConfiguration {
	for i := range values {
		b.ShortNames = append(b.ShortNames, values[i])
	}
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithKind(value string) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.Kind = &value
	return b
}

// WithListKind sets the ListKind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ListKind field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithListKind(value string) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.ListKind = &value
	return b
}

// WithCategories adds the given value to the Categories field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Categories field.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithCategories(values ...string) *NegotiatedAPIResourceSpecApplyConfiguration {
	for i := range values {
		b.Categories = append(b.Categories, values[i])
	}
	return b
}

// WithOpenAPIV3Schema sets the OpenAPIV3Schema field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenAPIV3Schema field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithOpenAPIV3Schema(value runtime.RawExtension) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.OpenAPIV3Schema = &value
	return b
}

// WithSubResources sets the SubResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubResources field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithSubResources(value apiresourcev1alpha1.SubResources) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.SubResources = &value
	return b
}

// WithColumnDefinitions sets the ColumnDefinitions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ColumnDefinitions field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithColumnDefinitions(value apiresourcev1alpha1.ColumnDefinitions) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.ColumnDefinitions = &value
	return b
}

// WithPublish sets the Publish field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Publish field is set to the value of the last call.
func (b *NegotiatedAPIResourceSpecApplyConfiguration) WithPublish(value bool) *NegotiatedAPIResourceSpecApplyConfiguration {
	b.Publish = &value
	return b
}
