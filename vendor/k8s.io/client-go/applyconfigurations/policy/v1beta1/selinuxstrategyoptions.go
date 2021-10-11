/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// SELinuxStrategyOptionsApplyConfiguration represents an declarative configuration of the SELinuxStrategyOptions type for use
// with apply.
type SELinuxStrategyOptionsApplyConfiguration struct {
	Rule           *v1beta1.SELinuxStrategy             `json:"rule,omitempty"`
	SELinuxOptions *v1.SELinuxOptionsApplyConfiguration `json:"seLinuxOptions,omitempty"`
}

// SELinuxStrategyOptionsApplyConfiguration constructs an declarative configuration of the SELinuxStrategyOptions type for use with
// apply.
func SELinuxStrategyOptions() *SELinuxStrategyOptionsApplyConfiguration {
	return &SELinuxStrategyOptionsApplyConfiguration{}
}

// WithRule sets the Rule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rule field is set to the value of the last call.
func (b *SELinuxStrategyOptionsApplyConfiguration) WithRule(value v1beta1.SELinuxStrategy) *SELinuxStrategyOptionsApplyConfiguration {
	b.Rule = &value
	return b
}

// WithSELinuxOptions sets the SELinuxOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SELinuxOptions field is set to the value of the last call.
func (b *SELinuxStrategyOptionsApplyConfiguration) WithSELinuxOptions(value *v1.SELinuxOptionsApplyConfiguration) *SELinuxStrategyOptionsApplyConfiguration {
	b.SELinuxOptions = value
	return b
}
