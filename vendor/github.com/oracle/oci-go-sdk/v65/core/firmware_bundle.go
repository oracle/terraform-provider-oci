// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FirmwareBundle A collection of pinned component firmware versions organized by platform.
type FirmwareBundle struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this firmware bundle.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name of this firmware bundle.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A brief description or metadata about this firmware bundle.
	Description *string `mandatory:"true" json:"description"`

	// A map of platforms to pinned firmware versions.
	Platforms []PlatformVersions `mandatory:"true" json:"platforms"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of this firmware bundle.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the firmware bundle was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the firmware bundle was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the firmware bundle.
	LifecycleState FirmwareBundleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	AllowableTransitions *FirmwareBundleTransitions `mandatory:"false" json:"allowableTransitions"`
}

func (m FirmwareBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FirmwareBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFirmwareBundleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFirmwareBundleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FirmwareBundleLifecycleStateEnum Enum with underlying type: string
type FirmwareBundleLifecycleStateEnum string

// Set of constants representing the allowable values for FirmwareBundleLifecycleStateEnum
const (
	FirmwareBundleLifecycleStateActive          FirmwareBundleLifecycleStateEnum = "ACTIVE"
	FirmwareBundleLifecycleStateInactive        FirmwareBundleLifecycleStateEnum = "INACTIVE"
	FirmwareBundleLifecycleStateDeleteScheduled FirmwareBundleLifecycleStateEnum = "DELETE_SCHEDULED"
)

var mappingFirmwareBundleLifecycleStateEnum = map[string]FirmwareBundleLifecycleStateEnum{
	"ACTIVE":           FirmwareBundleLifecycleStateActive,
	"INACTIVE":         FirmwareBundleLifecycleStateInactive,
	"DELETE_SCHEDULED": FirmwareBundleLifecycleStateDeleteScheduled,
}

var mappingFirmwareBundleLifecycleStateEnumLowerCase = map[string]FirmwareBundleLifecycleStateEnum{
	"active":           FirmwareBundleLifecycleStateActive,
	"inactive":         FirmwareBundleLifecycleStateInactive,
	"delete_scheduled": FirmwareBundleLifecycleStateDeleteScheduled,
}

// GetFirmwareBundleLifecycleStateEnumValues Enumerates the set of values for FirmwareBundleLifecycleStateEnum
func GetFirmwareBundleLifecycleStateEnumValues() []FirmwareBundleLifecycleStateEnum {
	values := make([]FirmwareBundleLifecycleStateEnum, 0)
	for _, v := range mappingFirmwareBundleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFirmwareBundleLifecycleStateEnumStringValues Enumerates the set of values in String for FirmwareBundleLifecycleStateEnum
func GetFirmwareBundleLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETE_SCHEDULED",
	}
}

// GetMappingFirmwareBundleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirmwareBundleLifecycleStateEnum(val string) (FirmwareBundleLifecycleStateEnum, bool) {
	enum, ok := mappingFirmwareBundleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
