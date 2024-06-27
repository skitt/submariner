/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

import (
	v1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
)

// ConnectionApplyConfiguration represents an declarative configuration of the Connection type for use
// with apply.
type ConnectionApplyConfiguration struct {
	Status        *v1.ConnectionStatus              `json:"status,omitempty"`
	StatusMessage *string                           `json:"statusMessage,omitempty"`
	Endpoint      *EndpointSpecApplyConfiguration   `json:"endpoint,omitempty"`
	UsingIP       *string                           `json:"usingIP,omitempty"`
	UsingNAT      *bool                             `json:"usingNAT,omitempty"`
	LatencyRTT    *LatencyRTTSpecApplyConfiguration `json:"latencyRTT,omitempty"`
}

// ConnectionApplyConfiguration constructs an declarative configuration of the Connection type for use with
// apply.
func Connection() *ConnectionApplyConfiguration {
	return &ConnectionApplyConfiguration{}
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithStatus(value v1.ConnectionStatus) *ConnectionApplyConfiguration {
	b.Status = &value
	return b
}

// WithStatusMessage sets the StatusMessage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StatusMessage field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithStatusMessage(value string) *ConnectionApplyConfiguration {
	b.StatusMessage = &value
	return b
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithEndpoint(value *EndpointSpecApplyConfiguration) *ConnectionApplyConfiguration {
	b.Endpoint = value
	return b
}

// WithUsingIP sets the UsingIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UsingIP field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithUsingIP(value string) *ConnectionApplyConfiguration {
	b.UsingIP = &value
	return b
}

// WithUsingNAT sets the UsingNAT field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UsingNAT field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithUsingNAT(value bool) *ConnectionApplyConfiguration {
	b.UsingNAT = &value
	return b
}

// WithLatencyRTT sets the LatencyRTT field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatencyRTT field is set to the value of the last call.
func (b *ConnectionApplyConfiguration) WithLatencyRTT(value *LatencyRTTSpecApplyConfiguration) *ConnectionApplyConfiguration {
	b.LatencyRTT = value
	return b
}
