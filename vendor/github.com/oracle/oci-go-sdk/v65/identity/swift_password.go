// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SwiftPassword **Deprecated. Use AuthToken instead.**
// Swift is the OpenStack object storage service. A `SwiftPassword` is an Oracle-provided password for using a
// Swift client with the Object Storage Service. This password is associated with
// the user's Console login. Swift passwords never expire. A user can have up to two Swift passwords at a time.
// **Note:** The password is always an Oracle-generated string; you can't change it to a string of your choice.
// For more information, see Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm).
type SwiftPassword struct {

	// The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not
	// for `ListSwiftPasswords` or `UpdateSwiftPassword`.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the Swift password.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the user the password belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// The description you assign to the Swift password. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description"`

	// Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when this password will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	ExpiresOn *common.SDKTime `mandatory:"false" json:"expiresOn"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState SwiftPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m SwiftPassword) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SwiftPassword) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSwiftPasswordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSwiftPasswordLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SwiftPasswordLifecycleStateEnum Enum with underlying type: string
type SwiftPasswordLifecycleStateEnum string

// Set of constants representing the allowable values for SwiftPasswordLifecycleStateEnum
const (
	SwiftPasswordLifecycleStateCreating SwiftPasswordLifecycleStateEnum = "CREATING"
	SwiftPasswordLifecycleStateActive   SwiftPasswordLifecycleStateEnum = "ACTIVE"
	SwiftPasswordLifecycleStateInactive SwiftPasswordLifecycleStateEnum = "INACTIVE"
	SwiftPasswordLifecycleStateDeleting SwiftPasswordLifecycleStateEnum = "DELETING"
	SwiftPasswordLifecycleStateDeleted  SwiftPasswordLifecycleStateEnum = "DELETED"
)

var mappingSwiftPasswordLifecycleStateEnum = map[string]SwiftPasswordLifecycleStateEnum{
	"CREATING": SwiftPasswordLifecycleStateCreating,
	"ACTIVE":   SwiftPasswordLifecycleStateActive,
	"INACTIVE": SwiftPasswordLifecycleStateInactive,
	"DELETING": SwiftPasswordLifecycleStateDeleting,
	"DELETED":  SwiftPasswordLifecycleStateDeleted,
}

var mappingSwiftPasswordLifecycleStateEnumLowerCase = map[string]SwiftPasswordLifecycleStateEnum{
	"creating": SwiftPasswordLifecycleStateCreating,
	"active":   SwiftPasswordLifecycleStateActive,
	"inactive": SwiftPasswordLifecycleStateInactive,
	"deleting": SwiftPasswordLifecycleStateDeleting,
	"deleted":  SwiftPasswordLifecycleStateDeleted,
}

// GetSwiftPasswordLifecycleStateEnumValues Enumerates the set of values for SwiftPasswordLifecycleStateEnum
func GetSwiftPasswordLifecycleStateEnumValues() []SwiftPasswordLifecycleStateEnum {
	values := make([]SwiftPasswordLifecycleStateEnum, 0)
	for _, v := range mappingSwiftPasswordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSwiftPasswordLifecycleStateEnumStringValues Enumerates the set of values in String for SwiftPasswordLifecycleStateEnum
func GetSwiftPasswordLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSwiftPasswordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSwiftPasswordLifecycleStateEnum(val string) (SwiftPasswordLifecycleStateEnum, bool) {
	enum, ok := mappingSwiftPasswordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
