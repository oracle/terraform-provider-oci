// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateBastionDetails The configuration details for a new bastion. A bastion provides secured, public access to target resources in the cloud that you cannot otherwise reach from the internet. A bastion resides in a public subnet and establishes the network infrastructure needed to connect a user to a target resource in a private subnet.
type CreateBastionDetails struct {

	// The type of bastion. Use `standard`.
	BastionType *string `mandatory:"true" json:"bastionType"`

	// The unique identifier (OCID) of the compartment where the bastion is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier (OCID) of the subnet that the bastion connects to.
	TargetSubnetId *string `mandatory:"true" json:"targetSubnetId"`

	// The name of the bastion, which can't be changed after creation.
	Name *string `mandatory:"false" json:"name"`

	// The phonebook entry of the customer's team, which can't be changed after creation. Not applicable to `standard` bastions.
	PhoneBookEntry *string `mandatory:"false" json:"phoneBookEntry"`

	// A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions.
	StaticJumpHostIpAddresses []string `mandatory:"false" json:"staticJumpHostIpAddresses"`

	// A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
	ClientCidrBlockAllowList []string `mandatory:"false" json:"clientCidrBlockAllowList"`

	// The maximum amount of time that any session on the bastion can remain active.
	MaxSessionTtlInSeconds *int `mandatory:"false" json:"maxSessionTtlInSeconds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBastionDetails) String() string {
	return common.PointerString(m)
}
