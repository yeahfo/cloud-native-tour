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

package v2

// FooConfigApplyConfiguration represents a declarative configuration of the FooConfig type for use
// with apply.
type FooConfigApplyConfiguration struct {
	Msg         *string `json:"msg,omitempty"`
	Description *string `json:"description,omitempty"`
}

// FooConfigApplyConfiguration constructs a declarative configuration of the FooConfig type for use with
// apply.
func FooConfig() *FooConfigApplyConfiguration {
	return &FooConfigApplyConfiguration{}
}

// WithMsg sets the Msg field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Msg field is set to the value of the last call.
func (b *FooConfigApplyConfiguration) WithMsg(value string) *FooConfigApplyConfiguration {
	b.Msg = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *FooConfigApplyConfiguration) WithDescription(value string) *FooConfigApplyConfiguration {
	b.Description = &value
	return b
}