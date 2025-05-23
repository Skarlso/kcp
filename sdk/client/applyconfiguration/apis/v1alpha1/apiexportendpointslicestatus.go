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
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// APIExportEndpointSliceStatusApplyConfiguration represents a declarative configuration of the APIExportEndpointSliceStatus type for use
// with apply.
type APIExportEndpointSliceStatusApplyConfiguration struct {
	Conditions         *conditionsv1alpha1.Conditions        `json:"conditions,omitempty"`
	APIExportEndpoints []APIExportEndpointApplyConfiguration `json:"endpoints,omitempty"`
	ShardSelector      *string                               `json:"shardSelector,omitempty"`
}

// APIExportEndpointSliceStatusApplyConfiguration constructs a declarative configuration of the APIExportEndpointSliceStatus type for use with
// apply.
func APIExportEndpointSliceStatus() *APIExportEndpointSliceStatusApplyConfiguration {
	return &APIExportEndpointSliceStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *APIExportEndpointSliceStatusApplyConfiguration) WithConditions(value conditionsv1alpha1.Conditions) *APIExportEndpointSliceStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithAPIExportEndpoints adds the given value to the APIExportEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIExportEndpoints field.
func (b *APIExportEndpointSliceStatusApplyConfiguration) WithAPIExportEndpoints(values ...*APIExportEndpointApplyConfiguration) *APIExportEndpointSliceStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAPIExportEndpoints")
		}
		b.APIExportEndpoints = append(b.APIExportEndpoints, *values[i])
	}
	return b
}

// WithShardSelector sets the ShardSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShardSelector field is set to the value of the last call.
func (b *APIExportEndpointSliceStatusApplyConfiguration) WithShardSelector(value string) *APIExportEndpointSliceStatusApplyConfiguration {
	b.ShardSelector = &value
	return b
}
