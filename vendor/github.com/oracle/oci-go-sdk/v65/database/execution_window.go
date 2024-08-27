// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecutionWindow Details of an execution window.
type ExecutionWindow struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the execution window.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the execution resource the execution window belongs to.
	ExecutionResourceId *string `mandatory:"true" json:"executionResourceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the execution window. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Schedule Policy. Valid states are CREATED, SCHEDULED, IN_PROGRESS, FAILED, CANCELED,
	// UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS.
	LifecycleState ExecutionWindowLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The scheduled start date and time of the execution window.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes.
	WindowDurationInMins *int `mandatory:"true" json:"windowDurationInMins"`

	// Description of the execution window.
	Description *string `mandatory:"false" json:"description"`

	// The current sub-state of the execution window. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING.
	LifecycleSubstate ExecutionWindowLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the execution window was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the execution window was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time that the execution window was started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the execution window ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
	IsEnforcedDuration *bool `mandatory:"false" json:"isEnforcedDuration"`

	// The estimated time of the execution window in minutes.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// The total time taken by corresponding resource activity in minutes.
	TotalTimeTakenInMins *int `mandatory:"false" json:"totalTimeTakenInMins"`

	// The execution window is of PLANNED or UNPLANNED type.
	WindowType ExecutionWindowWindowTypeEnum `mandatory:"false" json:"windowType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExecutionWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecutionWindowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExecutionWindowLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExecutionWindowLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetExecutionWindowLifecycleSubstateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionWindowWindowTypeEnum(string(m.WindowType)); !ok && m.WindowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WindowType: %s. Supported values are: %s.", m.WindowType, strings.Join(GetExecutionWindowWindowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecutionWindowLifecycleStateEnum Enum with underlying type: string
type ExecutionWindowLifecycleStateEnum string

// Set of constants representing the allowable values for ExecutionWindowLifecycleStateEnum
const (
	ExecutionWindowLifecycleStateCreated        ExecutionWindowLifecycleStateEnum = "CREATED"
	ExecutionWindowLifecycleStateScheduled      ExecutionWindowLifecycleStateEnum = "SCHEDULED"
	ExecutionWindowLifecycleStateInProgress     ExecutionWindowLifecycleStateEnum = "IN_PROGRESS"
	ExecutionWindowLifecycleStateFailed         ExecutionWindowLifecycleStateEnum = "FAILED"
	ExecutionWindowLifecycleStateCanceled       ExecutionWindowLifecycleStateEnum = "CANCELED"
	ExecutionWindowLifecycleStateUpdating       ExecutionWindowLifecycleStateEnum = "UPDATING"
	ExecutionWindowLifecycleStateDeleted        ExecutionWindowLifecycleStateEnum = "DELETED"
	ExecutionWindowLifecycleStateSucceeded      ExecutionWindowLifecycleStateEnum = "SUCCEEDED"
	ExecutionWindowLifecycleStatePartialSuccess ExecutionWindowLifecycleStateEnum = "PARTIAL_SUCCESS"
	ExecutionWindowLifecycleStateCreating       ExecutionWindowLifecycleStateEnum = "CREATING"
	ExecutionWindowLifecycleStateDeleting       ExecutionWindowLifecycleStateEnum = "DELETING"
)

var mappingExecutionWindowLifecycleStateEnum = map[string]ExecutionWindowLifecycleStateEnum{
	"CREATED":         ExecutionWindowLifecycleStateCreated,
	"SCHEDULED":       ExecutionWindowLifecycleStateScheduled,
	"IN_PROGRESS":     ExecutionWindowLifecycleStateInProgress,
	"FAILED":          ExecutionWindowLifecycleStateFailed,
	"CANCELED":        ExecutionWindowLifecycleStateCanceled,
	"UPDATING":        ExecutionWindowLifecycleStateUpdating,
	"DELETED":         ExecutionWindowLifecycleStateDeleted,
	"SUCCEEDED":       ExecutionWindowLifecycleStateSucceeded,
	"PARTIAL_SUCCESS": ExecutionWindowLifecycleStatePartialSuccess,
	"CREATING":        ExecutionWindowLifecycleStateCreating,
	"DELETING":        ExecutionWindowLifecycleStateDeleting,
}

var mappingExecutionWindowLifecycleStateEnumLowerCase = map[string]ExecutionWindowLifecycleStateEnum{
	"created":         ExecutionWindowLifecycleStateCreated,
	"scheduled":       ExecutionWindowLifecycleStateScheduled,
	"in_progress":     ExecutionWindowLifecycleStateInProgress,
	"failed":          ExecutionWindowLifecycleStateFailed,
	"canceled":        ExecutionWindowLifecycleStateCanceled,
	"updating":        ExecutionWindowLifecycleStateUpdating,
	"deleted":         ExecutionWindowLifecycleStateDeleted,
	"succeeded":       ExecutionWindowLifecycleStateSucceeded,
	"partial_success": ExecutionWindowLifecycleStatePartialSuccess,
	"creating":        ExecutionWindowLifecycleStateCreating,
	"deleting":        ExecutionWindowLifecycleStateDeleting,
}

// GetExecutionWindowLifecycleStateEnumValues Enumerates the set of values for ExecutionWindowLifecycleStateEnum
func GetExecutionWindowLifecycleStateEnumValues() []ExecutionWindowLifecycleStateEnum {
	values := make([]ExecutionWindowLifecycleStateEnum, 0)
	for _, v := range mappingExecutionWindowLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowLifecycleStateEnumStringValues Enumerates the set of values in String for ExecutionWindowLifecycleStateEnum
func GetExecutionWindowLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"SCHEDULED",
		"IN_PROGRESS",
		"FAILED",
		"CANCELED",
		"UPDATING",
		"DELETED",
		"SUCCEEDED",
		"PARTIAL_SUCCESS",
		"CREATING",
		"DELETING",
	}
}

// GetMappingExecutionWindowLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowLifecycleStateEnum(val string) (ExecutionWindowLifecycleStateEnum, bool) {
	enum, ok := mappingExecutionWindowLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionWindowLifecycleSubstateEnum Enum with underlying type: string
type ExecutionWindowLifecycleSubstateEnum string

// Set of constants representing the allowable values for ExecutionWindowLifecycleSubstateEnum
const (
	ExecutionWindowLifecycleSubstateDurationExceeded      ExecutionWindowLifecycleSubstateEnum = "DURATION_EXCEEDED"
	ExecutionWindowLifecycleSubstateMaintenanceInProgress ExecutionWindowLifecycleSubstateEnum = "MAINTENANCE_IN_PROGRESS"
	ExecutionWindowLifecycleSubstateWaiting               ExecutionWindowLifecycleSubstateEnum = "WAITING"
	ExecutionWindowLifecycleSubstateRescheduled           ExecutionWindowLifecycleSubstateEnum = "RESCHEDULED"
)

var mappingExecutionWindowLifecycleSubstateEnum = map[string]ExecutionWindowLifecycleSubstateEnum{
	"DURATION_EXCEEDED":       ExecutionWindowLifecycleSubstateDurationExceeded,
	"MAINTENANCE_IN_PROGRESS": ExecutionWindowLifecycleSubstateMaintenanceInProgress,
	"WAITING":                 ExecutionWindowLifecycleSubstateWaiting,
	"RESCHEDULED":             ExecutionWindowLifecycleSubstateRescheduled,
}

var mappingExecutionWindowLifecycleSubstateEnumLowerCase = map[string]ExecutionWindowLifecycleSubstateEnum{
	"duration_exceeded":       ExecutionWindowLifecycleSubstateDurationExceeded,
	"maintenance_in_progress": ExecutionWindowLifecycleSubstateMaintenanceInProgress,
	"waiting":                 ExecutionWindowLifecycleSubstateWaiting,
	"rescheduled":             ExecutionWindowLifecycleSubstateRescheduled,
}

// GetExecutionWindowLifecycleSubstateEnumValues Enumerates the set of values for ExecutionWindowLifecycleSubstateEnum
func GetExecutionWindowLifecycleSubstateEnumValues() []ExecutionWindowLifecycleSubstateEnum {
	values := make([]ExecutionWindowLifecycleSubstateEnum, 0)
	for _, v := range mappingExecutionWindowLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowLifecycleSubstateEnumStringValues Enumerates the set of values in String for ExecutionWindowLifecycleSubstateEnum
func GetExecutionWindowLifecycleSubstateEnumStringValues() []string {
	return []string{
		"DURATION_EXCEEDED",
		"MAINTENANCE_IN_PROGRESS",
		"WAITING",
		"RESCHEDULED",
	}
}

// GetMappingExecutionWindowLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowLifecycleSubstateEnum(val string) (ExecutionWindowLifecycleSubstateEnum, bool) {
	enum, ok := mappingExecutionWindowLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionWindowWindowTypeEnum Enum with underlying type: string
type ExecutionWindowWindowTypeEnum string

// Set of constants representing the allowable values for ExecutionWindowWindowTypeEnum
const (
	ExecutionWindowWindowTypePlanned   ExecutionWindowWindowTypeEnum = "PLANNED"
	ExecutionWindowWindowTypeUnplanned ExecutionWindowWindowTypeEnum = "UNPLANNED"
)

var mappingExecutionWindowWindowTypeEnum = map[string]ExecutionWindowWindowTypeEnum{
	"PLANNED":   ExecutionWindowWindowTypePlanned,
	"UNPLANNED": ExecutionWindowWindowTypeUnplanned,
}

var mappingExecutionWindowWindowTypeEnumLowerCase = map[string]ExecutionWindowWindowTypeEnum{
	"planned":   ExecutionWindowWindowTypePlanned,
	"unplanned": ExecutionWindowWindowTypeUnplanned,
}

// GetExecutionWindowWindowTypeEnumValues Enumerates the set of values for ExecutionWindowWindowTypeEnum
func GetExecutionWindowWindowTypeEnumValues() []ExecutionWindowWindowTypeEnum {
	values := make([]ExecutionWindowWindowTypeEnum, 0)
	for _, v := range mappingExecutionWindowWindowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowWindowTypeEnumStringValues Enumerates the set of values in String for ExecutionWindowWindowTypeEnum
func GetExecutionWindowWindowTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingExecutionWindowWindowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowWindowTypeEnum(val string) (ExecutionWindowWindowTypeEnum, bool) {
	enum, ok := mappingExecutionWindowWindowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
