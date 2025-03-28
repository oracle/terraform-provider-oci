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

// MfaTotpDeviceSummary As the name suggests, a `MfaTotpDeviceSummary` object contains information about a `MfaTotpDevice`.
type MfaTotpDeviceSummary struct {

	// The OCID of the MFA TOTP Device.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the user the MFA TOTP device belongs to.
	UserId *string `mandatory:"true" json:"userId"`

	// Date and time the `MfaTotpDevice` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The MFA TOTP device's current state.
	LifecycleState MfaTotpDeviceSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Flag to indicate if the MFA TOTP device has been activated
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

func (m MfaTotpDeviceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MfaTotpDeviceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMfaTotpDeviceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMfaTotpDeviceSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MfaTotpDeviceSummaryLifecycleStateEnum Enum with underlying type: string
type MfaTotpDeviceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for MfaTotpDeviceSummaryLifecycleStateEnum
const (
	MfaTotpDeviceSummaryLifecycleStateCreating MfaTotpDeviceSummaryLifecycleStateEnum = "CREATING"
	MfaTotpDeviceSummaryLifecycleStateActive   MfaTotpDeviceSummaryLifecycleStateEnum = "ACTIVE"
	MfaTotpDeviceSummaryLifecycleStateInactive MfaTotpDeviceSummaryLifecycleStateEnum = "INACTIVE"
	MfaTotpDeviceSummaryLifecycleStateDeleting MfaTotpDeviceSummaryLifecycleStateEnum = "DELETING"
	MfaTotpDeviceSummaryLifecycleStateDeleted  MfaTotpDeviceSummaryLifecycleStateEnum = "DELETED"
)

var mappingMfaTotpDeviceSummaryLifecycleStateEnum = map[string]MfaTotpDeviceSummaryLifecycleStateEnum{
	"CREATING": MfaTotpDeviceSummaryLifecycleStateCreating,
	"ACTIVE":   MfaTotpDeviceSummaryLifecycleStateActive,
	"INACTIVE": MfaTotpDeviceSummaryLifecycleStateInactive,
	"DELETING": MfaTotpDeviceSummaryLifecycleStateDeleting,
	"DELETED":  MfaTotpDeviceSummaryLifecycleStateDeleted,
}

var mappingMfaTotpDeviceSummaryLifecycleStateEnumLowerCase = map[string]MfaTotpDeviceSummaryLifecycleStateEnum{
	"creating": MfaTotpDeviceSummaryLifecycleStateCreating,
	"active":   MfaTotpDeviceSummaryLifecycleStateActive,
	"inactive": MfaTotpDeviceSummaryLifecycleStateInactive,
	"deleting": MfaTotpDeviceSummaryLifecycleStateDeleting,
	"deleted":  MfaTotpDeviceSummaryLifecycleStateDeleted,
}

// GetMfaTotpDeviceSummaryLifecycleStateEnumValues Enumerates the set of values for MfaTotpDeviceSummaryLifecycleStateEnum
func GetMfaTotpDeviceSummaryLifecycleStateEnumValues() []MfaTotpDeviceSummaryLifecycleStateEnum {
	values := make([]MfaTotpDeviceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingMfaTotpDeviceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMfaTotpDeviceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for MfaTotpDeviceSummaryLifecycleStateEnum
func GetMfaTotpDeviceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingMfaTotpDeviceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMfaTotpDeviceSummaryLifecycleStateEnum(val string) (MfaTotpDeviceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingMfaTotpDeviceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
