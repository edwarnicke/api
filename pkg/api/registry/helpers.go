// Copyright (c) 2020 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"github.com/golang/protobuf/proto"
)

// ServiceNames - returns grpc ServiceNames implemented by impl
func ServiceNames(impl interface{}, existingServiceNames ...string) []string {
	if _, ok := impl.(NetworkServiceRegistryServer); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceRegistryClient); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceEndpointRegistryServer); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceEndpointRegistry_serviceDesc.ServiceName)
	}
	if _, ok := impl.(NetworkServiceEndpointRegistryClient); ok {
		existingServiceNames = append(existingServiceNames, _NetworkServiceEndpointRegistry_serviceDesc.ServiceName)
	}
	return existingServiceNames
}

// ReadNetworkServiceList read list of NetworkServices from passed stream
func ReadNetworkServiceList(stream NetworkServiceRegistry_FindClient) []*NetworkService {
	var result []*NetworkService
	for msg, err := stream.Recv(); err == nil; msg, err = stream.Recv() {
		result = append(result, msg)
	}
	return result
}

// ReadNetworkServiceChannel reads NetworkServices from the stream and puts it into result channel
func ReadNetworkServiceChannel(stream NetworkServiceRegistry_FindClient) <-chan *NetworkService {
	result := make(chan *NetworkService)
	go func() {
		for msg, err := stream.Recv(); err == nil; msg, err = stream.Recv() {
			result <- msg
		}
		close(result)
	}()
	return result
}

// ReadNetworkServiceEndpointChannel reads NetworkServiceEndpoint from the stream and puts it into result channel
func ReadNetworkServiceEndpointChannel(stream NetworkServiceEndpointRegistry_FindClient) <-chan *NetworkServiceEndpoint {
	result := make(chan *NetworkServiceEndpoint)
	go func() {
		for msg, err := stream.Recv(); err == nil; msg, err = stream.Recv() {
			result <- msg
		}
		close(result)
	}()
	return result
}

// ReadNetworkServiceEndpointList read list of NetworkServiceEndpoints from passed stream
func ReadNetworkServiceEndpointList(stream NetworkServiceEndpointRegistry_FindClient) []*NetworkServiceEndpoint {
	var result []*NetworkServiceEndpoint
	for msg, err := stream.Recv(); err == nil; msg, err = stream.Recv() {
		result = append(result, msg)
	}
	return result
}

// Clone clones request
func (m *NetworkService) Clone() *NetworkService {
	return proto.Clone(m).(*NetworkService)
}

// Clone clones request
func (m *NetworkServiceEndpoint) Clone() *NetworkServiceEndpoint {
	return proto.Clone(m).(*NetworkServiceEndpoint)
}
