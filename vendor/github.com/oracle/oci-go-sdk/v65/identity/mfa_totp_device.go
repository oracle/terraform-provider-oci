// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MfaTotpDevice Users can enable multi-factor authentication (MFA) for their own user accounts. After MFA is enabled, the
// user is prompted for a time-based one-time password (TOTP) to authenticate before they can sign in to the
// Console. To enable multi-factor authentication, the user must register a mobile device with a TOTP authenticator app
// installed. The registration process creates the `MfaTotpDevice` object. The registration process requires
// interaction with the Console and cannot be completed programmatically. For more information, see
// Managing Multi-Factor Authentication (https://docs.oracle.com/iaas/Content/Identity/mfa/understand-multi-factor-authentication.htm).
type MfaTotpDevice struct {

	// The OCID of the MFA TOTP device.
	Id *string `mandatory:"true" json:"id"`

	// The seed for the MFA TOTP device (Base32 encoded).
	Seed *string `mandatory:"true" json:"seed"`

	// The OCID of the user the MFA TOTP device belongs to.
	UserId *string `mandatory:"true" json:"userId"`

	// Date and time the `MfaTotpDevice` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The MFA TOTP device's current state. After creating the MFA TOTP device, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState MfaTotpDeviceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Flag to indicate if the MFA TOTP device has been activated.
	IsActivated *bool `mandatory:"true" json:"isActivated"`

	// Date and time when this MFA TOTP device will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The detailed status of INACTIVE lifecycleState.
	// Allowed values are:
	//  - 1 - SUSPENDED
	//  - 2 - DISABLED
	//  - 4 - BLOCKED
	//  - 8 - LOCKED
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m MfaTotpDevice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MfaTotpDevice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMfaTotpDeviceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMfaTotpDeviceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingMfaTotpDeviceLifecycleStateEnum = map[string]MfaTotpDeviceLifecycleStateEnum{
	"CREATING": MfaTotpDeviceLifecycleStateCreating,
	"ACTIVE":   MfaTotpDeviceLifecycleStateActive,
	"INACTIVE": MfaTotpDeviceLifecycleStateInactive,
	"DELETING": MfaTotpDeviceLifecycleStateDeleting,
	"DELETED":  MfaTotpDeviceLifecycleStateDeleted,
}

var mappingMfaTotpDeviceLifecycleStateEnumLowerCase = map[string]MfaTotpDeviceLifecycleStateEnum{
	"creating": MfaTotpDeviceLifecycleStateCreating,
	"active":   MfaTotpDeviceLifecycleStateActive,
	"inactive": MfaTotpDeviceLifecycleStateInactive,
	"deleting": MfaTotpDeviceLifecycleStateDeleting,
	"deleted":  MfaTotpDeviceLifecycleStateDeleted,
}

// GetMfaTotpDeviceLifecycleStateEnumValues Enumerates the set of values for MfaTotpDeviceLifecycleStateEnum
func GetMfaTotpDeviceLifecycleStateEnumValues() []MfaTotpDeviceLifecycleStateEnum {
	values := make([]MfaTotpDeviceLifecycleStateEnum, 0)
	for _, v := range mappingMfaTotpDeviceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMfaTotpDeviceLifecycleStateEnumStringValues Enumerates the set of values in String for MfaTotpDeviceLifecycleStateEnum
func GetMfaTotpDeviceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingMfaTotpDeviceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMfaTotpDeviceLifecycleStateEnum(val string) (MfaTotpDeviceLifecycleStateEnum, bool) {
	enum, ok := mappingMfaTotpDeviceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
