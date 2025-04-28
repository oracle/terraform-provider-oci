// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedSoftwareUpdateMaintenanceDetails Provides details about actual Oracle Managed Database Software Updates scheduled time and version.
type ManagedSoftwareUpdateMaintenanceDetails struct {

	// The date and time of the database was scheduled for update.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// The state of the maintenance.
	State ManagedSoftwareUpdateMaintenanceDetailsStateEnum `mandatory:"true" json:"state"`

	// The version of the database was scheduled for update.
	Version *string `mandatory:"true" json:"version"`

	// Oracle Managed Database Software update method, either "ROLLING" or "NONROLLING"
	UpdateMode ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum `mandatory:"false" json:"updateMode,omitempty"`

	// The date and time of the last readiness check.
	TimeOfLastReadinessCheck *common.SDKTime `mandatory:"false" json:"timeOfLastReadinessCheck"`
}

func (m ManagedSoftwareUpdateMaintenanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedSoftwareUpdateMaintenanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedSoftwareUpdateMaintenanceDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetManagedSoftwareUpdateMaintenanceDetailsStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedSoftwareUpdateMaintenanceDetailsStateEnum Enum with underlying type: string
type ManagedSoftwareUpdateMaintenanceDetailsStateEnum string

// Set of constants representing the allowable values for ManagedSoftwareUpdateMaintenanceDetailsStateEnum
const (
	ManagedSoftwareUpdateMaintenanceDetailsStateScheduled           ManagedSoftwareUpdateMaintenanceDetailsStateEnum = "SCHEDULED"
	ManagedSoftwareUpdateMaintenanceDetailsStateReadyForPatching    ManagedSoftwareUpdateMaintenanceDetailsStateEnum = "READY_FOR_PATCHING"
	ManagedSoftwareUpdateMaintenanceDetailsStateSucceeded           ManagedSoftwareUpdateMaintenanceDetailsStateEnum = "SUCCEEDED"
	ManagedSoftwareUpdateMaintenanceDetailsStateNotReadyForPatching ManagedSoftwareUpdateMaintenanceDetailsStateEnum = "NOT_READY_FOR_PATCHING"
	ManagedSoftwareUpdateMaintenanceDetailsStateInProgress          ManagedSoftwareUpdateMaintenanceDetailsStateEnum = "IN_PROGRESS"
)

var mappingManagedSoftwareUpdateMaintenanceDetailsStateEnum = map[string]ManagedSoftwareUpdateMaintenanceDetailsStateEnum{
	"SCHEDULED":              ManagedSoftwareUpdateMaintenanceDetailsStateScheduled,
	"READY_FOR_PATCHING":     ManagedSoftwareUpdateMaintenanceDetailsStateReadyForPatching,
	"SUCCEEDED":              ManagedSoftwareUpdateMaintenanceDetailsStateSucceeded,
	"NOT_READY_FOR_PATCHING": ManagedSoftwareUpdateMaintenanceDetailsStateNotReadyForPatching,
	"IN_PROGRESS":            ManagedSoftwareUpdateMaintenanceDetailsStateInProgress,
}

var mappingManagedSoftwareUpdateMaintenanceDetailsStateEnumLowerCase = map[string]ManagedSoftwareUpdateMaintenanceDetailsStateEnum{
	"scheduled":              ManagedSoftwareUpdateMaintenanceDetailsStateScheduled,
	"ready_for_patching":     ManagedSoftwareUpdateMaintenanceDetailsStateReadyForPatching,
	"succeeded":              ManagedSoftwareUpdateMaintenanceDetailsStateSucceeded,
	"not_ready_for_patching": ManagedSoftwareUpdateMaintenanceDetailsStateNotReadyForPatching,
	"in_progress":            ManagedSoftwareUpdateMaintenanceDetailsStateInProgress,
}

// GetManagedSoftwareUpdateMaintenanceDetailsStateEnumValues Enumerates the set of values for ManagedSoftwareUpdateMaintenanceDetailsStateEnum
func GetManagedSoftwareUpdateMaintenanceDetailsStateEnumValues() []ManagedSoftwareUpdateMaintenanceDetailsStateEnum {
	values := make([]ManagedSoftwareUpdateMaintenanceDetailsStateEnum, 0)
	for _, v := range mappingManagedSoftwareUpdateMaintenanceDetailsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedSoftwareUpdateMaintenanceDetailsStateEnumStringValues Enumerates the set of values in String for ManagedSoftwareUpdateMaintenanceDetailsStateEnum
func GetManagedSoftwareUpdateMaintenanceDetailsStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"READY_FOR_PATCHING",
		"SUCCEEDED",
		"NOT_READY_FOR_PATCHING",
		"IN_PROGRESS",
	}
}

// GetMappingManagedSoftwareUpdateMaintenanceDetailsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedSoftwareUpdateMaintenanceDetailsStateEnum(val string) (ManagedSoftwareUpdateMaintenanceDetailsStateEnum, bool) {
	enum, ok := mappingManagedSoftwareUpdateMaintenanceDetailsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum Enum with underlying type: string
type ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum string

// Set of constants representing the allowable values for ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum
const (
	ManagedSoftwareUpdateMaintenanceDetailsUpdateModeRolling    ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum = "ROLLING"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateModeNonrolling ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum = "NONROLLING"
)

var mappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum = map[string]ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum{
	"ROLLING":    ManagedSoftwareUpdateMaintenanceDetailsUpdateModeRolling,
	"NONROLLING": ManagedSoftwareUpdateMaintenanceDetailsUpdateModeNonrolling,
}

var mappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumLowerCase = map[string]ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum{
	"rolling":    ManagedSoftwareUpdateMaintenanceDetailsUpdateModeRolling,
	"nonrolling": ManagedSoftwareUpdateMaintenanceDetailsUpdateModeNonrolling,
}

// GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumValues Enumerates the set of values for ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum
func GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumValues() []ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum {
	values := make([]ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum, 0)
	for _, v := range mappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumStringValues Enumerates the set of values in String for ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum
func GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum(val string) (ManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum, bool) {
	enum, ok := mappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
