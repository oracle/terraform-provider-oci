// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ScheduledTaskSummary Summary information about a scheduled task.
type ScheduledTaskSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data plane resource.
	Id *string `mandatory:"true" json:"id"`

	// Task type.
	TaskType TaskTypeEnum `mandatory:"true" json:"taskType"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the schedule task was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the scheduled task was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the scheduled task.
	LifecycleState ScheduledTaskLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Status of the scheduled task.
	TaskStatus ScheduledTaskSummaryTaskStatusEnum `mandatory:"false" json:"taskStatus,omitempty"`

	// reason for taskStatus PAUSED.
	PauseReason ScheduledTaskPauseReasonEnum `mandatory:"false" json:"pauseReason,omitempty"`

	// most recent Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The most recent task execution status.
	LastExecutionStatus ScheduledTaskSummaryLastExecutionStatusEnum `mandatory:"false" json:"lastExecutionStatus,omitempty"`

	// The date and time the scheduled task last executed, in the format defined by RFC3339.
	TimeLastExecuted *common.SDKTime `mandatory:"false" json:"timeLastExecuted"`
}

func (m ScheduledTaskSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledTaskSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledTaskLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScheduledTaskSummaryTaskStatusEnum(string(m.TaskStatus)); !ok && m.TaskStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskStatus: %s. Supported values are: %s.", m.TaskStatus, strings.Join(GetScheduledTaskSummaryTaskStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskPauseReasonEnum(string(m.PauseReason)); !ok && m.PauseReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PauseReason: %s. Supported values are: %s.", m.PauseReason, strings.Join(GetScheduledTaskPauseReasonEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskSummaryLastExecutionStatusEnum(string(m.LastExecutionStatus)); !ok && m.LastExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastExecutionStatus: %s. Supported values are: %s.", m.LastExecutionStatus, strings.Join(GetScheduledTaskSummaryLastExecutionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledTaskSummaryTaskStatusEnum Enum with underlying type: string
type ScheduledTaskSummaryTaskStatusEnum string

// Set of constants representing the allowable values for ScheduledTaskSummaryTaskStatusEnum
const (
	ScheduledTaskSummaryTaskStatusReady     ScheduledTaskSummaryTaskStatusEnum = "READY"
	ScheduledTaskSummaryTaskStatusPaused    ScheduledTaskSummaryTaskStatusEnum = "PAUSED"
	ScheduledTaskSummaryTaskStatusCompleted ScheduledTaskSummaryTaskStatusEnum = "COMPLETED"
	ScheduledTaskSummaryTaskStatusBlocked   ScheduledTaskSummaryTaskStatusEnum = "BLOCKED"
)

var mappingScheduledTaskSummaryTaskStatusEnum = map[string]ScheduledTaskSummaryTaskStatusEnum{
	"READY":     ScheduledTaskSummaryTaskStatusReady,
	"PAUSED":    ScheduledTaskSummaryTaskStatusPaused,
	"COMPLETED": ScheduledTaskSummaryTaskStatusCompleted,
	"BLOCKED":   ScheduledTaskSummaryTaskStatusBlocked,
}

// GetScheduledTaskSummaryTaskStatusEnumValues Enumerates the set of values for ScheduledTaskSummaryTaskStatusEnum
func GetScheduledTaskSummaryTaskStatusEnumValues() []ScheduledTaskSummaryTaskStatusEnum {
	values := make([]ScheduledTaskSummaryTaskStatusEnum, 0)
	for _, v := range mappingScheduledTaskSummaryTaskStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskSummaryTaskStatusEnumStringValues Enumerates the set of values in String for ScheduledTaskSummaryTaskStatusEnum
func GetScheduledTaskSummaryTaskStatusEnumStringValues() []string {
	return []string{
		"READY",
		"PAUSED",
		"COMPLETED",
		"BLOCKED",
	}
}

// GetMappingScheduledTaskSummaryTaskStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskSummaryTaskStatusEnum(val string) (ScheduledTaskSummaryTaskStatusEnum, bool) {
	mappingScheduledTaskSummaryTaskStatusEnumIgnoreCase := make(map[string]ScheduledTaskSummaryTaskStatusEnum)
	for k, v := range mappingScheduledTaskSummaryTaskStatusEnum {
		mappingScheduledTaskSummaryTaskStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskSummaryTaskStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledTaskSummaryLastExecutionStatusEnum Enum with underlying type: string
type ScheduledTaskSummaryLastExecutionStatusEnum string

// Set of constants representing the allowable values for ScheduledTaskSummaryLastExecutionStatusEnum
const (
	ScheduledTaskSummaryLastExecutionStatusFailed    ScheduledTaskSummaryLastExecutionStatusEnum = "FAILED"
	ScheduledTaskSummaryLastExecutionStatusSucceeded ScheduledTaskSummaryLastExecutionStatusEnum = "SUCCEEDED"
)

var mappingScheduledTaskSummaryLastExecutionStatusEnum = map[string]ScheduledTaskSummaryLastExecutionStatusEnum{
	"FAILED":    ScheduledTaskSummaryLastExecutionStatusFailed,
	"SUCCEEDED": ScheduledTaskSummaryLastExecutionStatusSucceeded,
}

// GetScheduledTaskSummaryLastExecutionStatusEnumValues Enumerates the set of values for ScheduledTaskSummaryLastExecutionStatusEnum
func GetScheduledTaskSummaryLastExecutionStatusEnumValues() []ScheduledTaskSummaryLastExecutionStatusEnum {
	values := make([]ScheduledTaskSummaryLastExecutionStatusEnum, 0)
	for _, v := range mappingScheduledTaskSummaryLastExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskSummaryLastExecutionStatusEnumStringValues Enumerates the set of values in String for ScheduledTaskSummaryLastExecutionStatusEnum
func GetScheduledTaskSummaryLastExecutionStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingScheduledTaskSummaryLastExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskSummaryLastExecutionStatusEnum(val string) (ScheduledTaskSummaryLastExecutionStatusEnum, bool) {
	mappingScheduledTaskSummaryLastExecutionStatusEnumIgnoreCase := make(map[string]ScheduledTaskSummaryLastExecutionStatusEnum)
	for k, v := range mappingScheduledTaskSummaryLastExecutionStatusEnum {
		mappingScheduledTaskSummaryLastExecutionStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskSummaryLastExecutionStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
