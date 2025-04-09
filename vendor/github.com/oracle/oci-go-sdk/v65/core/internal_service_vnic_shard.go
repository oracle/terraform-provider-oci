// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalServiceVnicShard A ServiceVnicShard contains service vnic shard config details
type InternalServiceVnicShard struct {

	// Name of the service vnic shard
	Name *string `mandatory:"true" json:"name"`

	// It is used as key for shard level allocations like slotId, to ensure that they are unique for all service
	// vnics attached to a shard.
	IpAddressForShardLevelAllocations *string `mandatory:"true" json:"ipAddressForShardLevelAllocations"`

	// Regional anycast IP for NLB service VNICs
	IpAddressForNlbRegional *string `mandatory:"true" json:"ipAddressForNlbRegional"`

	// List of service vnic ad configs
	AdConfigs []InternalServiceVnicAdConfig `mandatory:"true" json:"adConfigs"`

	// Indicates whether service vnic shard is enabled for placing vnic attachments
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Indicates whether service vnic shard can be used for latency sensitive vnic attachments
	IsLatencySensitive *bool `mandatory:"false" json:"isLatencySensitive"`

	// Indicates whether this service vnic shard is only dedicated for placing latency sensitive vnic attachments
	IsDedicatedLatencySensitive *bool `mandatory:"false" json:"isDedicatedLatencySensitive"`

	// Indicates whether this service vnic shard is only dedicated for cpg
	IsDedicatedCpg *bool `mandatory:"false" json:"isDedicatedCpg"`

	// Max vnics supported on this service vnic shard
	MaxVnics *int `mandatory:"false" json:"maxVnics"`

	// List of resource types that are disabled on this shard
	DisabledResourceTypes []string `mandatory:"false" json:"disabledResourceTypes"`

	// List of tags on this shard
	Tags []string `mandatory:"false" json:"tags"`

	// The date and time ServiceVnicShard was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time ServiceVnicShard was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m InternalServiceVnicShard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalServiceVnicShard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
