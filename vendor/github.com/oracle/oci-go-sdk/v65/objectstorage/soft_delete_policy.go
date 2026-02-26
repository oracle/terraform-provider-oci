// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftDeletePolicy The details of a soft delete policy.
type SoftDeletePolicy struct {

	// The date and time that the soft delete policy becomes active.
	TimeActivation *common.SDKTime `mandatory:"true" json:"timeActivation"`

	// Set the soft delete status on the bucket. By default, a bucket is created with soft delete `Disabled`. Use this option to enable soft delete during bucket creation. Objects in a soft delete enabled bucket are protected from permanent deletions. Soft deleted versions of the object will be available in the bucket for restoration.
	SoftDeleteMode SoftDeletePolicySoftDeleteModeEnum `mandatory:"false" json:"softDeleteMode,omitempty"`

	// The duration (specified as an ISO 8601 extended format string) that a soft-deleted object is available for recovery.
	// Valid range: P1D to P365D (e.g., "P30D" for 30 days).
	// After this soft delete restore period elapses, the object is permanently deleted and can no longer be restored.
	RestoreDuration *string `mandatory:"false" json:"restoreDuration"`

	// The duration (specified as an ISO 8601 extended format string) to wait before the updated soft delete configuration becomes effective.
	// Valid range: PT1H to PT240H (e.g., "PT24H" for 24 hours).
	// During the cool-off period, the current soft delete policy continues to apply. After the cool-off period elapses, the new configuration is activated.
	CoolOffDuration *string `mandatory:"false" json:"coolOffDuration"`
}

func (m SoftDeletePolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoftDeletePolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSoftDeletePolicySoftDeleteModeEnum(string(m.SoftDeleteMode)); !ok && m.SoftDeleteMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftDeleteMode: %s. Supported values are: %s.", m.SoftDeleteMode, strings.Join(GetSoftDeletePolicySoftDeleteModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SoftDeletePolicySoftDeleteModeEnum Enum with underlying type: string
type SoftDeletePolicySoftDeleteModeEnum string

// Set of constants representing the allowable values for SoftDeletePolicySoftDeleteModeEnum
const (
	SoftDeletePolicySoftDeleteModeEnabled   SoftDeletePolicySoftDeleteModeEnum = "Enabled"
	SoftDeletePolicySoftDeleteModeSuspended SoftDeletePolicySoftDeleteModeEnum = "Suspended"
	SoftDeletePolicySoftDeleteModeDisabled  SoftDeletePolicySoftDeleteModeEnum = "Disabled"
)

var mappingSoftDeletePolicySoftDeleteModeEnum = map[string]SoftDeletePolicySoftDeleteModeEnum{
	"Enabled":   SoftDeletePolicySoftDeleteModeEnabled,
	"Suspended": SoftDeletePolicySoftDeleteModeSuspended,
	"Disabled":  SoftDeletePolicySoftDeleteModeDisabled,
}

var mappingSoftDeletePolicySoftDeleteModeEnumLowerCase = map[string]SoftDeletePolicySoftDeleteModeEnum{
	"enabled":   SoftDeletePolicySoftDeleteModeEnabled,
	"suspended": SoftDeletePolicySoftDeleteModeSuspended,
	"disabled":  SoftDeletePolicySoftDeleteModeDisabled,
}

// GetSoftDeletePolicySoftDeleteModeEnumValues Enumerates the set of values for SoftDeletePolicySoftDeleteModeEnum
func GetSoftDeletePolicySoftDeleteModeEnumValues() []SoftDeletePolicySoftDeleteModeEnum {
	values := make([]SoftDeletePolicySoftDeleteModeEnum, 0)
	for _, v := range mappingSoftDeletePolicySoftDeleteModeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftDeletePolicySoftDeleteModeEnumStringValues Enumerates the set of values in String for SoftDeletePolicySoftDeleteModeEnum
func GetSoftDeletePolicySoftDeleteModeEnumStringValues() []string {
	return []string{
		"Enabled",
		"Suspended",
		"Disabled",
	}
}

// GetMappingSoftDeletePolicySoftDeleteModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftDeletePolicySoftDeleteModeEnum(val string) (SoftDeletePolicySoftDeleteModeEnum, bool) {
	enum, ok := mappingSoftDeletePolicySoftDeleteModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
