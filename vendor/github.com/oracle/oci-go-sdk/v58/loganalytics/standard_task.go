// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StandardTask Log analytics scheduled task resource.
type StandardTask struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data plane resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Schedules.
	Schedules []Schedule `mandatory:"true" json:"schedules"`

	Action Action `mandatory:"true" json:"action"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the scheduled task was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the scheduled task was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// most recent Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// Number of execution occurrences.
	NumOccurrences *int64 `mandatory:"false" json:"numOccurrences"`

	// The date and time the scheduled task will execute next,
	// in the format defined by RFC3339.
	TimeOfNextExecution *common.SDKTime `mandatory:"false" json:"timeOfNextExecution"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The date and time the scheduled task last executed, in the format defined by RFC3339.
	TimeLastExecuted *common.SDKTime `mandatory:"false" json:"timeLastExecuted"`

	// The most recent task execution status.
	LastExecutionStatus StandardTaskLastExecutionStatusEnum `mandatory:"false" json:"lastExecutionStatus,omitempty"`

	// Task type.
	TaskType TaskTypeEnum `mandatory:"true" json:"taskType"`

	// Status of the scheduled task.
	TaskStatus ScheduledTaskTaskStatusEnum `mandatory:"false" json:"taskStatus,omitempty"`

	// reason for taskStatus PAUSED.
	PauseReason ScheduledTaskPauseReasonEnum `mandatory:"false" json:"pauseReason,omitempty"`

	// The current state of the scheduled task.
	LifecycleState ScheduledTaskLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m StandardTask) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m StandardTask) GetDisplayName() *string {
	return m.DisplayName
}

//GetTaskType returns TaskType
func (m StandardTask) GetTaskType() TaskTypeEnum {
	return m.TaskType
}

//GetSchedules returns Schedules
func (m StandardTask) GetSchedules() []Schedule {
	return m.Schedules
}

//GetAction returns Action
func (m StandardTask) GetAction() Action {
	return m.Action
}

//GetTaskStatus returns TaskStatus
func (m StandardTask) GetTaskStatus() ScheduledTaskTaskStatusEnum {
	return m.TaskStatus
}

//GetPauseReason returns PauseReason
func (m StandardTask) GetPauseReason() ScheduledTaskPauseReasonEnum {
	return m.PauseReason
}

//GetWorkRequestId returns WorkRequestId
func (m StandardTask) GetWorkRequestId() *string {
	return m.WorkRequestId
}

//GetNumOccurrences returns NumOccurrences
func (m StandardTask) GetNumOccurrences() *int64 {
	return m.NumOccurrences
}

//GetCompartmentId returns CompartmentId
func (m StandardTask) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m StandardTask) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m StandardTask) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetTimeOfNextExecution returns TimeOfNextExecution
func (m StandardTask) GetTimeOfNextExecution() *common.SDKTime {
	return m.TimeOfNextExecution
}

//GetLifecycleState returns LifecycleState
func (m StandardTask) GetLifecycleState() ScheduledTaskLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m StandardTask) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m StandardTask) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m StandardTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StandardTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStandardTaskLastExecutionStatusEnum(string(m.LastExecutionStatus)); !ok && m.LastExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastExecutionStatus: %s. Supported values are: %s.", m.LastExecutionStatus, strings.Join(GetStandardTaskLastExecutionStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskTaskStatusEnum(string(m.TaskStatus)); !ok && m.TaskStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskStatus: %s. Supported values are: %s.", m.TaskStatus, strings.Join(GetScheduledTaskTaskStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskPauseReasonEnum(string(m.PauseReason)); !ok && m.PauseReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PauseReason: %s. Supported values are: %s.", m.PauseReason, strings.Join(GetScheduledTaskPauseReasonEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StandardTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStandardTask StandardTask
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStandardTask
	}{
		"STANDARD",
		(MarshalTypeStandardTask)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *StandardTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskStatus          ScheduledTaskTaskStatusEnum         `json:"taskStatus"`
		PauseReason         ScheduledTaskPauseReasonEnum        `json:"pauseReason"`
		WorkRequestId       *string                             `json:"workRequestId"`
		NumOccurrences      *int64                              `json:"numOccurrences"`
		TimeOfNextExecution *common.SDKTime                     `json:"timeOfNextExecution"`
		FreeformTags        map[string]string                   `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}   `json:"definedTags"`
		LastExecutionStatus StandardTaskLastExecutionStatusEnum `json:"lastExecutionStatus"`
		TimeLastExecuted    *common.SDKTime                     `json:"timeLastExecuted"`
		Id                  *string                             `json:"id"`
		DisplayName         *string                             `json:"displayName"`
		TaskType            TaskTypeEnum                        `json:"taskType"`
		Schedules           []schedule                          `json:"schedules"`
		Action              action                              `json:"action"`
		CompartmentId       *string                             `json:"compartmentId"`
		TimeCreated         *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                     `json:"timeUpdated"`
		LifecycleState      ScheduledTaskLifecycleStateEnum     `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskStatus = model.TaskStatus

	m.PauseReason = model.PauseReason

	m.WorkRequestId = model.WorkRequestId

	m.NumOccurrences = model.NumOccurrences

	m.TimeOfNextExecution = model.TimeOfNextExecution

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.LastExecutionStatus = model.LastExecutionStatus

	m.TimeLastExecuted = model.TimeLastExecuted

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.TaskType = model.TaskType

	m.Schedules = make([]Schedule, len(model.Schedules))
	for i, n := range model.Schedules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Schedules[i] = nn.(Schedule)
		} else {
			m.Schedules[i] = nil
		}
	}

	nn, e = model.Action.UnmarshalPolymorphicJSON(model.Action.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Action = nn.(Action)
	} else {
		m.Action = nil
	}

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}

// StandardTaskLastExecutionStatusEnum Enum with underlying type: string
type StandardTaskLastExecutionStatusEnum string

// Set of constants representing the allowable values for StandardTaskLastExecutionStatusEnum
const (
	StandardTaskLastExecutionStatusFailed    StandardTaskLastExecutionStatusEnum = "FAILED"
	StandardTaskLastExecutionStatusSucceeded StandardTaskLastExecutionStatusEnum = "SUCCEEDED"
)

var mappingStandardTaskLastExecutionStatusEnum = map[string]StandardTaskLastExecutionStatusEnum{
	"FAILED":    StandardTaskLastExecutionStatusFailed,
	"SUCCEEDED": StandardTaskLastExecutionStatusSucceeded,
}

// GetStandardTaskLastExecutionStatusEnumValues Enumerates the set of values for StandardTaskLastExecutionStatusEnum
func GetStandardTaskLastExecutionStatusEnumValues() []StandardTaskLastExecutionStatusEnum {
	values := make([]StandardTaskLastExecutionStatusEnum, 0)
	for _, v := range mappingStandardTaskLastExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStandardTaskLastExecutionStatusEnumStringValues Enumerates the set of values in String for StandardTaskLastExecutionStatusEnum
func GetStandardTaskLastExecutionStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingStandardTaskLastExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStandardTaskLastExecutionStatusEnum(val string) (StandardTaskLastExecutionStatusEnum, bool) {
	mappingStandardTaskLastExecutionStatusEnumIgnoreCase := make(map[string]StandardTaskLastExecutionStatusEnum)
	for k, v := range mappingStandardTaskLastExecutionStatusEnum {
		mappingStandardTaskLastExecutionStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStandardTaskLastExecutionStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
