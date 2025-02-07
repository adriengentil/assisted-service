/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterStatusBuilder contains the data and logic needed to build 'cluster_status' objects.
//
// Detailed status of a cluster.
type ClusterStatusBuilder struct {
	bitmap_               uint32
	id                    string
	href                  string
	configurationMode     ClusterConfigurationMode
	description           string
	provisionErrorCode    string
	provisionErrorMessage string
	state                 ClusterState
	dnsReady              bool
}

// NewClusterStatus creates a new builder of 'cluster_status' objects.
func NewClusterStatus() *ClusterStatusBuilder {
	return &ClusterStatusBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *ClusterStatusBuilder) Link(value bool) *ClusterStatusBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *ClusterStatusBuilder) ID(value string) *ClusterStatusBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *ClusterStatusBuilder) HREF(value string) *ClusterStatusBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// DNSReady sets the value of the 'DNS_ready' attribute to the given value.
func (b *ClusterStatusBuilder) DNSReady(value bool) *ClusterStatusBuilder {
	b.dnsReady = value
	b.bitmap_ |= 8
	return b
}

// ConfigurationMode sets the value of the 'configuration_mode' attribute to the given value.
//
// Configuration mode of a cluster.
func (b *ClusterStatusBuilder) ConfigurationMode(value ClusterConfigurationMode) *ClusterStatusBuilder {
	b.configurationMode = value
	b.bitmap_ |= 16
	return b
}

// Description sets the value of the 'description' attribute to the given value.
func (b *ClusterStatusBuilder) Description(value string) *ClusterStatusBuilder {
	b.description = value
	b.bitmap_ |= 32
	return b
}

// ProvisionErrorCode sets the value of the 'provision_error_code' attribute to the given value.
func (b *ClusterStatusBuilder) ProvisionErrorCode(value string) *ClusterStatusBuilder {
	b.provisionErrorCode = value
	b.bitmap_ |= 64
	return b
}

// ProvisionErrorMessage sets the value of the 'provision_error_message' attribute to the given value.
func (b *ClusterStatusBuilder) ProvisionErrorMessage(value string) *ClusterStatusBuilder {
	b.provisionErrorMessage = value
	b.bitmap_ |= 128
	return b
}

// State sets the value of the 'state' attribute to the given value.
//
// Overall state of a cluster.
func (b *ClusterStatusBuilder) State(value ClusterState) *ClusterStatusBuilder {
	b.state = value
	b.bitmap_ |= 256
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterStatusBuilder) Copy(object *ClusterStatus) *ClusterStatusBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	b.dnsReady = object.dnsReady
	b.configurationMode = object.configurationMode
	b.description = object.description
	b.provisionErrorCode = object.provisionErrorCode
	b.provisionErrorMessage = object.provisionErrorMessage
	b.state = object.state
	return b
}

// Build creates a 'cluster_status' object using the configuration stored in the builder.
func (b *ClusterStatusBuilder) Build() (object *ClusterStatus, err error) {
	object = new(ClusterStatus)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	object.dnsReady = b.dnsReady
	object.configurationMode = b.configurationMode
	object.description = b.description
	object.provisionErrorCode = b.provisionErrorCode
	object.provisionErrorMessage = b.provisionErrorMessage
	object.state = b.state
	return
}
