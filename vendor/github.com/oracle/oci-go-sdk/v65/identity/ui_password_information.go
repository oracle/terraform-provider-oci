// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UiPasswordInformation Information about the UIPassword, which is a text password that enables a user to sign in to the Console,
// the user interface for interacting with Oracle Cloud Infrastructure.
// For more information about user credentials, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/Concepts/usercredentials.htm).
type UiPasswordInformation struct {

	// The OCID of the user.
	UserId *string `mandatory:"false" json:"userId"`

	// Date and time the password was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState UiPasswordInformationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UiPasswordInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UiPasswordInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUiPasswordInformationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUiPasswordInformationLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UiPasswordInformationLifecycleStateEnum Enum with underlying type: string
type UiPasswordInformationLifecycleStateEnum string

// Set of constants representing the allowable values for UiPasswordInformationLifecycleStateEnum
const (
	UiPasswordInformationLifecycleStateCreating UiPasswordInformationLifecycleStateEnum = "CREATING"
	UiPasswordInformationLifecycleStateActive   UiPasswordInformationLifecycleStateEnum = "ACTIVE"
	UiPasswordInformationLifecycleStateInactive UiPasswordInformationLifecycleStateEnum = "INACTIVE"
	UiPasswordInformationLifecycleStateDeleting UiPasswordInformationLifecycleStateEnum = "DELETING"
	UiPasswordInformationLifecycleStateDeleted  UiPasswordInformationLifecycleStateEnum = "DELETED"
)

var mappingUiPasswordInformationLifecycleStateEnum = map[string]UiPasswordInformationLifecycleStateEnum{
	"CREATING": UiPasswordInformationLifecycleStateCreating,
	"ACTIVE":   UiPasswordInformationLifecycleStateActive,
	"INACTIVE": UiPasswordInformationLifecycleStateInactive,
	"DELETING": UiPasswordInformationLifecycleStateDeleting,
	"DELETED":  UiPasswordInformationLifecycleStateDeleted,
}

var mappingUiPasswordInformationLifecycleStateEnumLowerCase = map[string]UiPasswordInformationLifecycleStateEnum{
	"creating": UiPasswordInformationLifecycleStateCreating,
	"active":   UiPasswordInformationLifecycleStateActive,
	"inactive": UiPasswordInformationLifecycleStateInactive,
	"deleting": UiPasswordInformationLifecycleStateDeleting,
	"deleted":  UiPasswordInformationLifecycleStateDeleted,
}

// GetUiPasswordInformationLifecycleStateEnumValues Enumerates the set of values for UiPasswordInformationLifecycleStateEnum
func GetUiPasswordInformationLifecycleStateEnumValues() []UiPasswordInformationLifecycleStateEnum {
	values := make([]UiPasswordInformationLifecycleStateEnum, 0)
	for _, v := range mappingUiPasswordInformationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUiPasswordInformationLifecycleStateEnumStringValues Enumerates the set of values in String for UiPasswordInformationLifecycleStateEnum
func GetUiPasswordInformationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingUiPasswordInformationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUiPasswordInformationLifecycleStateEnum(val string) (UiPasswordInformationLifecycleStateEnum, bool) {
	enum, ok := mappingUiPasswordInformationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
