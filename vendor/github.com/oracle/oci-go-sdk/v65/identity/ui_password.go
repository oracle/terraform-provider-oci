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

// UiPassword A text password that enables a user to sign in to the Console, the user interface for interacting with Oracle
// Cloud Infrastructure.
// For more information about user credentials, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/usercred/usercredentials.htm).
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
	LifecycleState UiPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m UiPassword) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UiPassword) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUiPasswordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUiPasswordLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UiPasswordLifecycleStateEnum Enum with underlying type: string
type UiPasswordLifecycleStateEnum string

// Set of constants representing the allowable values for UiPasswordLifecycleStateEnum
const (
	UiPasswordLifecycleStateCreating UiPasswordLifecycleStateEnum = "CREATING"
	UiPasswordLifecycleStateActive   UiPasswordLifecycleStateEnum = "ACTIVE"
	UiPasswordLifecycleStateInactive UiPasswordLifecycleStateEnum = "INACTIVE"
	UiPasswordLifecycleStateDeleting UiPasswordLifecycleStateEnum = "DELETING"
	UiPasswordLifecycleStateDeleted  UiPasswordLifecycleStateEnum = "DELETED"
)

var mappingUiPasswordLifecycleStateEnum = map[string]UiPasswordLifecycleStateEnum{
	"CREATING": UiPasswordLifecycleStateCreating,
	"ACTIVE":   UiPasswordLifecycleStateActive,
	"INACTIVE": UiPasswordLifecycleStateInactive,
	"DELETING": UiPasswordLifecycleStateDeleting,
	"DELETED":  UiPasswordLifecycleStateDeleted,
}

var mappingUiPasswordLifecycleStateEnumLowerCase = map[string]UiPasswordLifecycleStateEnum{
	"creating": UiPasswordLifecycleStateCreating,
	"active":   UiPasswordLifecycleStateActive,
	"inactive": UiPasswordLifecycleStateInactive,
	"deleting": UiPasswordLifecycleStateDeleting,
	"deleted":  UiPasswordLifecycleStateDeleted,
}

// GetUiPasswordLifecycleStateEnumValues Enumerates the set of values for UiPasswordLifecycleStateEnum
func GetUiPasswordLifecycleStateEnumValues() []UiPasswordLifecycleStateEnum {
	values := make([]UiPasswordLifecycleStateEnum, 0)
	for _, v := range mappingUiPasswordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUiPasswordLifecycleStateEnumStringValues Enumerates the set of values in String for UiPasswordLifecycleStateEnum
func GetUiPasswordLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingUiPasswordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUiPasswordLifecycleStateEnum(val string) (UiPasswordLifecycleStateEnum, bool) {
	enum, ok := mappingUiPasswordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
