// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UiPassword A text password that enables a user to sign in to the Console, the user interface for interacting with Oracle
// Cloud Infrastructure.
// For more information about user credentials, see [User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Concepts/usercredentials.htm).
type UiPassword struct {

	// The user's password for the Console.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the user.
	UserId *string `mandatory:"false" json:"userId"`

	// Date and time the password was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState UiPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus"`
}

func (m UiPassword) String() string {
	return common.PointerString(m)
}

// UiPasswordLifecycleStateEnum Enum with underlying type: string
type UiPasswordLifecycleStateEnum string

// Set of constants representing the allowable values for UiPasswordLifecycleState
const (
	UiPasswordLifecycleStateCreating UiPasswordLifecycleStateEnum = "CREATING"
	UiPasswordLifecycleStateActive   UiPasswordLifecycleStateEnum = "ACTIVE"
	UiPasswordLifecycleStateInactive UiPasswordLifecycleStateEnum = "INACTIVE"
	UiPasswordLifecycleStateDeleting UiPasswordLifecycleStateEnum = "DELETING"
	UiPasswordLifecycleStateDeleted  UiPasswordLifecycleStateEnum = "DELETED"
	UiPasswordLifecycleStateUnknown  UiPasswordLifecycleStateEnum = "UNKNOWN"
)

var mappingUiPasswordLifecycleState = map[string]UiPasswordLifecycleStateEnum{
	"CREATING": UiPasswordLifecycleStateCreating,
	"ACTIVE":   UiPasswordLifecycleStateActive,
	"INACTIVE": UiPasswordLifecycleStateInactive,
	"DELETING": UiPasswordLifecycleStateDeleting,
	"DELETED":  UiPasswordLifecycleStateDeleted,
	"UNKNOWN":  UiPasswordLifecycleStateUnknown,
}

// GetUiPasswordLifecycleStateEnumValues Enumerates the set of values for UiPasswordLifecycleState
func GetUiPasswordLifecycleStateEnumValues() []UiPasswordLifecycleStateEnum {
	values := make([]UiPasswordLifecycleStateEnum, 0)
	for _, v := range mappingUiPasswordLifecycleState {
		if v != UiPasswordLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
