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

// LocalPeeringConnection Details regarding a local peering connection, which is an entity that allows two VCNs to communicate
// without traversing the Internet.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type LocalPeeringConnection struct {

	// The OCID of the compartment containing the local peering connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The local peering connection's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The local peering connection's current lifecycle state.
	LifecycleState LocalPeeringConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates whether the local peering connection is peered with another local peering connection.
	PeeringStatus LocalPeeringConnectionPeeringStatusEnum `mandatory:"true" json:"peeringStatus"`

	// The date and time the local peering connection was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the VCN the local peering connection belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Indicates whether the peer local peering connection is contained within another tenancy.
	IsCrossTenancyPeering *bool `mandatory:"false" json:"isCrossTenancyPeering"`

	// Indicates the range of IPs available on the peer. `null` if not peered.
	PeerAdvertisedCidr *string `mandatory:"false" json:"peerAdvertisedCidr"`

	// Additional information regarding the peering status if applicable.
	PeeringStatusDetails *string `mandatory:"false" json:"peeringStatusDetails"`
}

func (m LocalPeeringConnection) String() string {
	return common.PointerString(m)
}

// LocalPeeringConnectionLifecycleStateEnum Enum with underlying type: string
type LocalPeeringConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for LocalPeeringConnectionLifecycleStateEnum
const (
	LocalPeeringConnectionLifecycleStateProvisioning LocalPeeringConnectionLifecycleStateEnum = "PROVISIONING"
	LocalPeeringConnectionLifecycleStateAvailable    LocalPeeringConnectionLifecycleStateEnum = "AVAILABLE"
	LocalPeeringConnectionLifecycleStateTerminating  LocalPeeringConnectionLifecycleStateEnum = "TERMINATING"
	LocalPeeringConnectionLifecycleStateTerminated   LocalPeeringConnectionLifecycleStateEnum = "TERMINATED"
)

var mappingLocalPeeringConnectionLifecycleState = map[string]LocalPeeringConnectionLifecycleStateEnum{
	"PROVISIONING": LocalPeeringConnectionLifecycleStateProvisioning,
	"AVAILABLE":    LocalPeeringConnectionLifecycleStateAvailable,
	"TERMINATING":  LocalPeeringConnectionLifecycleStateTerminating,
	"TERMINATED":   LocalPeeringConnectionLifecycleStateTerminated,
}

// GetLocalPeeringConnectionLifecycleStateEnumValues Enumerates the set of values for LocalPeeringConnectionLifecycleStateEnum
func GetLocalPeeringConnectionLifecycleStateEnumValues() []LocalPeeringConnectionLifecycleStateEnum {
	values := make([]LocalPeeringConnectionLifecycleStateEnum, 0)
	for _, v := range mappingLocalPeeringConnectionLifecycleState {
		values = append(values, v)
	}
	return values
}

// LocalPeeringConnectionPeeringStatusEnum Enum with underlying type: string
type LocalPeeringConnectionPeeringStatusEnum string

// Set of constants representing the allowable values for LocalPeeringConnectionPeeringStatusEnum
const (
	LocalPeeringConnectionPeeringStatusInvalid LocalPeeringConnectionPeeringStatusEnum = "INVALID"
	LocalPeeringConnectionPeeringStatusNew     LocalPeeringConnectionPeeringStatusEnum = "NEW"
	LocalPeeringConnectionPeeringStatusPeered  LocalPeeringConnectionPeeringStatusEnum = "PEERED"
	LocalPeeringConnectionPeeringStatusPending LocalPeeringConnectionPeeringStatusEnum = "PENDING"
	LocalPeeringConnectionPeeringStatusRevoked LocalPeeringConnectionPeeringStatusEnum = "REVOKED"
)

var mappingLocalPeeringConnectionPeeringStatus = map[string]LocalPeeringConnectionPeeringStatusEnum{
	"INVALID": LocalPeeringConnectionPeeringStatusInvalid,
	"NEW":     LocalPeeringConnectionPeeringStatusNew,
	"PEERED":  LocalPeeringConnectionPeeringStatusPeered,
	"PENDING": LocalPeeringConnectionPeeringStatusPending,
	"REVOKED": LocalPeeringConnectionPeeringStatusRevoked,
}

// GetLocalPeeringConnectionPeeringStatusEnumValues Enumerates the set of values for LocalPeeringConnectionPeeringStatusEnum
func GetLocalPeeringConnectionPeeringStatusEnumValues() []LocalPeeringConnectionPeeringStatusEnum {
	values := make([]LocalPeeringConnectionPeeringStatusEnum, 0)
	for _, v := range mappingLocalPeeringConnectionPeeringStatus {
		values = append(values, v)
	}
	return values
}
