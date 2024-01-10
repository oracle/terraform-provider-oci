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

// UserGroupMembership An object that represents the membership of a user in a group. When you add a user to a group, the result is a
// `UserGroupMembership` with its own OCID. To remove a user from a group, you delete the `UserGroupMembership` object.
type UserGroupMembership struct {

	// The OCID of the membership.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the user, group, and membership object.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the group.
	GroupId *string `mandatory:"true" json:"groupId"`

	// The OCID of the user.
	UserId *string `mandatory:"true" json:"userId"`

	// Date and time the membership was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes
	// from CREATING to ACTIVE before using it.
	LifecycleState UserGroupMembershipLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m UserGroupMembership) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserGroupMembership) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserGroupMembershipLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUserGroupMembershipLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserGroupMembershipLifecycleStateEnum Enum with underlying type: string
type UserGroupMembershipLifecycleStateEnum string

// Set of constants representing the allowable values for UserGroupMembershipLifecycleStateEnum
const (
	UserGroupMembershipLifecycleStateCreating UserGroupMembershipLifecycleStateEnum = "CREATING"
	UserGroupMembershipLifecycleStateActive   UserGroupMembershipLifecycleStateEnum = "ACTIVE"
	UserGroupMembershipLifecycleStateInactive UserGroupMembershipLifecycleStateEnum = "INACTIVE"
	UserGroupMembershipLifecycleStateDeleting UserGroupMembershipLifecycleStateEnum = "DELETING"
	UserGroupMembershipLifecycleStateDeleted  UserGroupMembershipLifecycleStateEnum = "DELETED"
)

var mappingUserGroupMembershipLifecycleStateEnum = map[string]UserGroupMembershipLifecycleStateEnum{
	"CREATING": UserGroupMembershipLifecycleStateCreating,
	"ACTIVE":   UserGroupMembershipLifecycleStateActive,
	"INACTIVE": UserGroupMembershipLifecycleStateInactive,
	"DELETING": UserGroupMembershipLifecycleStateDeleting,
	"DELETED":  UserGroupMembershipLifecycleStateDeleted,
}

var mappingUserGroupMembershipLifecycleStateEnumLowerCase = map[string]UserGroupMembershipLifecycleStateEnum{
	"creating": UserGroupMembershipLifecycleStateCreating,
	"active":   UserGroupMembershipLifecycleStateActive,
	"inactive": UserGroupMembershipLifecycleStateInactive,
	"deleting": UserGroupMembershipLifecycleStateDeleting,
	"deleted":  UserGroupMembershipLifecycleStateDeleted,
}

// GetUserGroupMembershipLifecycleStateEnumValues Enumerates the set of values for UserGroupMembershipLifecycleStateEnum
func GetUserGroupMembershipLifecycleStateEnumValues() []UserGroupMembershipLifecycleStateEnum {
	values := make([]UserGroupMembershipLifecycleStateEnum, 0)
	for _, v := range mappingUserGroupMembershipLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUserGroupMembershipLifecycleStateEnumStringValues Enumerates the set of values in String for UserGroupMembershipLifecycleStateEnum
func GetUserGroupMembershipLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingUserGroupMembershipLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserGroupMembershipLifecycleStateEnum(val string) (UserGroupMembershipLifecycleStateEnum, bool) {
	enum, ok := mappingUserGroupMembershipLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
