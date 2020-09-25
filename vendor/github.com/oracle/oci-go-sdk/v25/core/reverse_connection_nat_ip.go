// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// ReverseConnectionNatIp The current NAT IP address that corresponds to a specific customer IP address.
// To establish a reverse connection to a customer IP address, use the NAT IP
// address as the destination.
type ReverseConnectionNatIp struct {

	// The reverse connection NAT IP's current state.
	LifecycleState ReverseConnectionNatIpLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the reverse connection NAT IP was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The customer's IP address that corresponds to the reverse connection NAT IP address.
	ReverseConnectionCustomerIp *string `mandatory:"true" json:"reverseConnectionCustomerIp"`

	// The reverse connection NAT IP address corresonding to the customer IP and private endpoint.
	// Use this address as the destination when establishing a reverse connection to a customer's
	// IP address.
	ReverseConnectionNatIp *string `mandatory:"true" json:"reverseConnectionNatIp"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private endpoint
	// associated with the reverse connection.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`
}

func (m ReverseConnectionNatIp) String() string {
	return common.PointerString(m)
}

// ReverseConnectionNatIpLifecycleStateEnum Enum with underlying type: string
type ReverseConnectionNatIpLifecycleStateEnum string

// Set of constants representing the allowable values for ReverseConnectionNatIpLifecycleStateEnum
const (
	ReverseConnectionNatIpLifecycleStateProvisioning ReverseConnectionNatIpLifecycleStateEnum = "PROVISIONING"
	ReverseConnectionNatIpLifecycleStateAvailable    ReverseConnectionNatIpLifecycleStateEnum = "AVAILABLE"
	ReverseConnectionNatIpLifecycleStateTerminating  ReverseConnectionNatIpLifecycleStateEnum = "TERMINATING"
	ReverseConnectionNatIpLifecycleStateTerminated   ReverseConnectionNatIpLifecycleStateEnum = "TERMINATED"
)

var mappingReverseConnectionNatIpLifecycleState = map[string]ReverseConnectionNatIpLifecycleStateEnum{
	"PROVISIONING": ReverseConnectionNatIpLifecycleStateProvisioning,
	"AVAILABLE":    ReverseConnectionNatIpLifecycleStateAvailable,
	"TERMINATING":  ReverseConnectionNatIpLifecycleStateTerminating,
	"TERMINATED":   ReverseConnectionNatIpLifecycleStateTerminated,
}

// GetReverseConnectionNatIpLifecycleStateEnumValues Enumerates the set of values for ReverseConnectionNatIpLifecycleStateEnum
func GetReverseConnectionNatIpLifecycleStateEnumValues() []ReverseConnectionNatIpLifecycleStateEnum {
	values := make([]ReverseConnectionNatIpLifecycleStateEnum, 0)
	for _, v := range mappingReverseConnectionNatIpLifecycleState {
		values = append(values, v)
	}
	return values
}
