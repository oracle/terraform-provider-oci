// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MfaTotpDevice A `MfaTotpDevice` is an Mfa Totp device that the user can use to authenticate with OCI (Leslie will add more details here)
type MfaTotpDevice struct {

	// The OCID of the Mfa Totp Device.
	Id *string `mandatory:"true" json:"id"`

	// The seed for the Mfa Totp device (Base32 encoded)
	Seed *string `mandatory:"true" json:"seed"`

	// The OCID of the user the Mfa Totp Device belongs to.
	UserId *string `mandatory:"true" json:"userId"`

	// Date and time the `Mfa Totp Device` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The Mfa Totp Device's current state. After creating a Mfa Totp Device, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState MfaTotpDeviceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Flag to indicate if the Mfa Totp Device has been isActivated
	IsActivated *bool `mandatory:"true" json:"isActivated"`

	// Date and time when this Mfa Totp device will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The detailed status of INACTIVE lifecycleState. Possible values are 1(SUSPENDED), 2(DISABLED), 4(BLOCKED) and 8(LOCKED).
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m MfaTotpDevice) String() string {
	return common.PointerString(m)
}

// MfaTotpDeviceLifecycleStateEnum Enum with underlying type: string
type MfaTotpDeviceLifecycleStateEnum string

// Set of constants representing the allowable values for MfaTotpDeviceLifecycleStateEnum
const (
	MfaTotpDeviceLifecycleStateCreating MfaTotpDeviceLifecycleStateEnum = "CREATING"
	MfaTotpDeviceLifecycleStateActive   MfaTotpDeviceLifecycleStateEnum = "ACTIVE"
	MfaTotpDeviceLifecycleStateInactive MfaTotpDeviceLifecycleStateEnum = "INACTIVE"
	MfaTotpDeviceLifecycleStateDeleting MfaTotpDeviceLifecycleStateEnum = "DELETING"
	MfaTotpDeviceLifecycleStateDeleted  MfaTotpDeviceLifecycleStateEnum = "DELETED"
)

var mappingMfaTotpDeviceLifecycleState = map[string]MfaTotpDeviceLifecycleStateEnum{
	"CREATING": MfaTotpDeviceLifecycleStateCreating,
	"ACTIVE":   MfaTotpDeviceLifecycleStateActive,
	"INACTIVE": MfaTotpDeviceLifecycleStateInactive,
	"DELETING": MfaTotpDeviceLifecycleStateDeleting,
	"DELETED":  MfaTotpDeviceLifecycleStateDeleted,
}

// GetMfaTotpDeviceLifecycleStateEnumValues Enumerates the set of values for MfaTotpDeviceLifecycleStateEnum
func GetMfaTotpDeviceLifecycleStateEnumValues() []MfaTotpDeviceLifecycleStateEnum {
	values := make([]MfaTotpDeviceLifecycleStateEnum, 0)
	for _, v := range mappingMfaTotpDeviceLifecycleState {
		values = append(values, v)
	}
	return values
}
