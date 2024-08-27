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

// ExecutionWindowSummary Details of an execution window.
type ExecutionWindowSummary struct {

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
	LifecycleState ExecutionWindowSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The scheduled start date and time of the execution window.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes.
	WindowDurationInMins *int `mandatory:"true" json:"windowDurationInMins"`

	// Description of the execution window.
	Description *string `mandatory:"false" json:"description"`

	// The current sub-state of the execution window. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING.
	LifecycleSubstate ExecutionWindowSummaryLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

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
	WindowType ExecutionWindowSummaryWindowTypeEnum `mandatory:"false" json:"windowType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExecutionWindowSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionWindowSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecutionWindowSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExecutionWindowSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExecutionWindowSummaryLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetExecutionWindowSummaryLifecycleSubstateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionWindowSummaryWindowTypeEnum(string(m.WindowType)); !ok && m.WindowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WindowType: %s. Supported values are: %s.", m.WindowType, strings.Join(GetExecutionWindowSummaryWindowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecutionWindowSummaryLifecycleStateEnum Enum with underlying type: string
type ExecutionWindowSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExecutionWindowSummaryLifecycleStateEnum
const (
	ExecutionWindowSummaryLifecycleStateCreated        ExecutionWindowSummaryLifecycleStateEnum = "CREATED"
	ExecutionWindowSummaryLifecycleStateScheduled      ExecutionWindowSummaryLifecycleStateEnum = "SCHEDULED"
	ExecutionWindowSummaryLifecycleStateInProgress     ExecutionWindowSummaryLifecycleStateEnum = "IN_PROGRESS"
	ExecutionWindowSummaryLifecycleStateFailed         ExecutionWindowSummaryLifecycleStateEnum = "FAILED"
	ExecutionWindowSummaryLifecycleStateCanceled       ExecutionWindowSummaryLifecycleStateEnum = "CANCELED"
	ExecutionWindowSummaryLifecycleStateUpdating       ExecutionWindowSummaryLifecycleStateEnum = "UPDATING"
	ExecutionWindowSummaryLifecycleStateDeleted        ExecutionWindowSummaryLifecycleStateEnum = "DELETED"
	ExecutionWindowSummaryLifecycleStateSucceeded      ExecutionWindowSummaryLifecycleStateEnum = "SUCCEEDED"
	ExecutionWindowSummaryLifecycleStatePartialSuccess ExecutionWindowSummaryLifecycleStateEnum = "PARTIAL_SUCCESS"
	ExecutionWindowSummaryLifecycleStateCreating       ExecutionWindowSummaryLifecycleStateEnum = "CREATING"
	ExecutionWindowSummaryLifecycleStateDeleting       ExecutionWindowSummaryLifecycleStateEnum = "DELETING"
)

var mappingExecutionWindowSummaryLifecycleStateEnum = map[string]ExecutionWindowSummaryLifecycleStateEnum{
	"CREATED":         ExecutionWindowSummaryLifecycleStateCreated,
	"SCHEDULED":       ExecutionWindowSummaryLifecycleStateScheduled,
	"IN_PROGRESS":     ExecutionWindowSummaryLifecycleStateInProgress,
	"FAILED":          ExecutionWindowSummaryLifecycleStateFailed,
	"CANCELED":        ExecutionWindowSummaryLifecycleStateCanceled,
	"UPDATING":        ExecutionWindowSummaryLifecycleStateUpdating,
	"DELETED":         ExecutionWindowSummaryLifecycleStateDeleted,
	"SUCCEEDED":       ExecutionWindowSummaryLifecycleStateSucceeded,
	"PARTIAL_SUCCESS": ExecutionWindowSummaryLifecycleStatePartialSuccess,
	"CREATING":        ExecutionWindowSummaryLifecycleStateCreating,
	"DELETING":        ExecutionWindowSummaryLifecycleStateDeleting,
}

var mappingExecutionWindowSummaryLifecycleStateEnumLowerCase = map[string]ExecutionWindowSummaryLifecycleStateEnum{
	"created":         ExecutionWindowSummaryLifecycleStateCreated,
	"scheduled":       ExecutionWindowSummaryLifecycleStateScheduled,
	"in_progress":     ExecutionWindowSummaryLifecycleStateInProgress,
	"failed":          ExecutionWindowSummaryLifecycleStateFailed,
	"canceled":        ExecutionWindowSummaryLifecycleStateCanceled,
	"updating":        ExecutionWindowSummaryLifecycleStateUpdating,
	"deleted":         ExecutionWindowSummaryLifecycleStateDeleted,
	"succeeded":       ExecutionWindowSummaryLifecycleStateSucceeded,
	"partial_success": ExecutionWindowSummaryLifecycleStatePartialSuccess,
	"creating":        ExecutionWindowSummaryLifecycleStateCreating,
	"deleting":        ExecutionWindowSummaryLifecycleStateDeleting,
}

// GetExecutionWindowSummaryLifecycleStateEnumValues Enumerates the set of values for ExecutionWindowSummaryLifecycleStateEnum
func GetExecutionWindowSummaryLifecycleStateEnumValues() []ExecutionWindowSummaryLifecycleStateEnum {
	values := make([]ExecutionWindowSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExecutionWindowSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExecutionWindowSummaryLifecycleStateEnum
func GetExecutionWindowSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingExecutionWindowSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowSummaryLifecycleStateEnum(val string) (ExecutionWindowSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExecutionWindowSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionWindowSummaryLifecycleSubstateEnum Enum with underlying type: string
type ExecutionWindowSummaryLifecycleSubstateEnum string

// Set of constants representing the allowable values for ExecutionWindowSummaryLifecycleSubstateEnum
const (
	ExecutionWindowSummaryLifecycleSubstateDurationExceeded      ExecutionWindowSummaryLifecycleSubstateEnum = "DURATION_EXCEEDED"
	ExecutionWindowSummaryLifecycleSubstateMaintenanceInProgress ExecutionWindowSummaryLifecycleSubstateEnum = "MAINTENANCE_IN_PROGRESS"
	ExecutionWindowSummaryLifecycleSubstateWaiting               ExecutionWindowSummaryLifecycleSubstateEnum = "WAITING"
	ExecutionWindowSummaryLifecycleSubstateRescheduled           ExecutionWindowSummaryLifecycleSubstateEnum = "RESCHEDULED"
)

var mappingExecutionWindowSummaryLifecycleSubstateEnum = map[string]ExecutionWindowSummaryLifecycleSubstateEnum{
	"DURATION_EXCEEDED":       ExecutionWindowSummaryLifecycleSubstateDurationExceeded,
	"MAINTENANCE_IN_PROGRESS": ExecutionWindowSummaryLifecycleSubstateMaintenanceInProgress,
	"WAITING":                 ExecutionWindowSummaryLifecycleSubstateWaiting,
	"RESCHEDULED":             ExecutionWindowSummaryLifecycleSubstateRescheduled,
}

var mappingExecutionWindowSummaryLifecycleSubstateEnumLowerCase = map[string]ExecutionWindowSummaryLifecycleSubstateEnum{
	"duration_exceeded":       ExecutionWindowSummaryLifecycleSubstateDurationExceeded,
	"maintenance_in_progress": ExecutionWindowSummaryLifecycleSubstateMaintenanceInProgress,
	"waiting":                 ExecutionWindowSummaryLifecycleSubstateWaiting,
	"rescheduled":             ExecutionWindowSummaryLifecycleSubstateRescheduled,
}

// GetExecutionWindowSummaryLifecycleSubstateEnumValues Enumerates the set of values for ExecutionWindowSummaryLifecycleSubstateEnum
func GetExecutionWindowSummaryLifecycleSubstateEnumValues() []ExecutionWindowSummaryLifecycleSubstateEnum {
	values := make([]ExecutionWindowSummaryLifecycleSubstateEnum, 0)
	for _, v := range mappingExecutionWindowSummaryLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowSummaryLifecycleSubstateEnumStringValues Enumerates the set of values in String for ExecutionWindowSummaryLifecycleSubstateEnum
func GetExecutionWindowSummaryLifecycleSubstateEnumStringValues() []string {
	return []string{
		"DURATION_EXCEEDED",
		"MAINTENANCE_IN_PROGRESS",
		"WAITING",
		"RESCHEDULED",
	}
}

// GetMappingExecutionWindowSummaryLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowSummaryLifecycleSubstateEnum(val string) (ExecutionWindowSummaryLifecycleSubstateEnum, bool) {
	enum, ok := mappingExecutionWindowSummaryLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionWindowSummaryWindowTypeEnum Enum with underlying type: string
type ExecutionWindowSummaryWindowTypeEnum string

// Set of constants representing the allowable values for ExecutionWindowSummaryWindowTypeEnum
const (
	ExecutionWindowSummaryWindowTypePlanned   ExecutionWindowSummaryWindowTypeEnum = "PLANNED"
	ExecutionWindowSummaryWindowTypeUnplanned ExecutionWindowSummaryWindowTypeEnum = "UNPLANNED"
)

var mappingExecutionWindowSummaryWindowTypeEnum = map[string]ExecutionWindowSummaryWindowTypeEnum{
	"PLANNED":   ExecutionWindowSummaryWindowTypePlanned,
	"UNPLANNED": ExecutionWindowSummaryWindowTypeUnplanned,
}

var mappingExecutionWindowSummaryWindowTypeEnumLowerCase = map[string]ExecutionWindowSummaryWindowTypeEnum{
	"planned":   ExecutionWindowSummaryWindowTypePlanned,
	"unplanned": ExecutionWindowSummaryWindowTypeUnplanned,
}

// GetExecutionWindowSummaryWindowTypeEnumValues Enumerates the set of values for ExecutionWindowSummaryWindowTypeEnum
func GetExecutionWindowSummaryWindowTypeEnumValues() []ExecutionWindowSummaryWindowTypeEnum {
	values := make([]ExecutionWindowSummaryWindowTypeEnum, 0)
	for _, v := range mappingExecutionWindowSummaryWindowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionWindowSummaryWindowTypeEnumStringValues Enumerates the set of values in String for ExecutionWindowSummaryWindowTypeEnum
func GetExecutionWindowSummaryWindowTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingExecutionWindowSummaryWindowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionWindowSummaryWindowTypeEnum(val string) (ExecutionWindowSummaryWindowTypeEnum, bool) {
	enum, ok := mappingExecutionWindowSummaryWindowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
