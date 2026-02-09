// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

	// The version of the database was scheduled for update.
	Version *string `mandatory:"true" json:"version"`

	// The managed software update readiness status
	UpdateReadinessStatus ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum `mandatory:"false" json:"updateReadinessStatus,omitempty"`

	// This field will contain actual cause of update readiness state.
	UpdateReadinessStatusDetails *string `mandatory:"false" json:"updateReadinessStatusDetails"`

	// The date and time of when the status was updated.
	TimeOfStatusUpdate *common.SDKTime `mandatory:"false" json:"timeOfStatusUpdate"`

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

	if _, ok := GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum(string(m.UpdateReadinessStatus)); !ok && m.UpdateReadinessStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateReadinessStatus: %s. Supported values are: %s.", m.UpdateReadinessStatus, strings.Join(GetManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetManagedSoftwareUpdateMaintenanceDetailsUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum Enum with underlying type: string
type ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum string

// Set of constants representing the allowable values for ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum
const (
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusScheduled                ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "SCHEDULED"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusSucceeded                ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "SUCCEEDED"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusReadyForManagedUpdate    ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "READY_FOR_MANAGED_UPDATE"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNotReadyForManagedUpdate ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "NOT_READY_FOR_MANAGED_UPDATE"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusInProgress               ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "IN_PROGRESS"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNeedsAttention           ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "NEEDS_ATTENTION"
	ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusAwaitingResolution       ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = "AWAITING_RESOLUTION"
)

var mappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum = map[string]ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum{
	"SCHEDULED":                    ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusScheduled,
	"SUCCEEDED":                    ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusSucceeded,
	"READY_FOR_MANAGED_UPDATE":     ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusReadyForManagedUpdate,
	"NOT_READY_FOR_MANAGED_UPDATE": ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNotReadyForManagedUpdate,
	"IN_PROGRESS":                  ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusInProgress,
	"NEEDS_ATTENTION":              ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNeedsAttention,
	"AWAITING_RESOLUTION":          ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusAwaitingResolution,
}

var mappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumLowerCase = map[string]ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum{
	"scheduled":                    ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusScheduled,
	"succeeded":                    ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusSucceeded,
	"ready_for_managed_update":     ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusReadyForManagedUpdate,
	"not_ready_for_managed_update": ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNotReadyForManagedUpdate,
	"in_progress":                  ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusInProgress,
	"needs_attention":              ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusNeedsAttention,
	"awaiting_resolution":          ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusAwaitingResolution,
}

// GetManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumValues Enumerates the set of values for ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum
func GetManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumValues() []ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum {
	values := make([]ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum, 0)
	for _, v := range mappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumStringValues Enumerates the set of values in String for ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum
func GetManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"SUCCEEDED",
		"READY_FOR_MANAGED_UPDATE",
		"NOT_READY_FOR_MANAGED_UPDATE",
		"IN_PROGRESS",
		"NEEDS_ATTENTION",
		"AWAITING_RESOLUTION",
	}
}

// GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum(val string) (ManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnum, bool) {
	enum, ok := mappingManagedSoftwareUpdateMaintenanceDetailsUpdateReadinessStatusEnumLowerCase[strings.ToLower(val)]
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
