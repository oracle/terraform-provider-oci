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

// UpdateBastionDetails The configuration to update on an existing bastion.
type UpdateBastionDetails struct {

	// The maximum amount of time that any session on the bastion can remain active.
	MaxSessionTtlInSeconds *int `mandatory:"false" json:"maxSessionTtlInSeconds"`

	// A list of IP addresses of the hosts that the bastion has access to. Not applicable to `standard` bastions.
	StaticJumpHostIpAddresses []string `mandatory:"false" json:"staticJumpHostIpAddresses"`

	// A list of address ranges in CIDR notation that you want to allow to connect to sessions hosted by this bastion.
	ClientCidrBlockAllowList []string `mandatory:"false" json:"clientCidrBlockAllowList"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBastionDetails) String() string {
	return common.PointerString(m)
}
