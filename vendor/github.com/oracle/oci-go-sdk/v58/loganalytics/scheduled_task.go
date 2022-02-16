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

// ScheduledTask Log analytics scheduled task resource.
type ScheduledTask interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data plane resource.
	GetId() *string

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	GetDisplayName() *string

	// Task type.
	GetTaskType() TaskTypeEnum

	// Schedules.
	GetSchedules() []Schedule

	GetAction() Action

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	GetCompartmentId() *string

	// The date and time the scheduled task was created, in the format defined by RFC3339.
	GetTimeCreated() *common.SDKTime

	// The date and time the scheduled task was last updated, in the format defined by RFC3339.
	GetTimeUpdated() *common.SDKTime

	// The current state of the scheduled task.
	GetLifecycleState() ScheduledTaskLifecycleStateEnum

	// Status of the scheduled task.
	GetTaskStatus() ScheduledTaskTaskStatusEnum

	// reason for taskStatus PAUSED.
	GetPauseReason() ScheduledTaskPauseReasonEnum

	// most recent Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	GetWorkRequestId() *string

	// Number of execution occurrences.
	GetNumOccurrences() *int64

	// The date and time the scheduled task will execute next,
	// in the format defined by RFC3339.
	GetTimeOfNextExecution() *common.SDKTime

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type scheduledtask struct {
	JsonData            []byte
	Id                  *string                           `mandatory:"true" json:"id"`
	DisplayName         *string                           `mandatory:"true" json:"displayName"`
	TaskType            TaskTypeEnum                      `mandatory:"true" json:"taskType"`
	Schedules           json.RawMessage                   `mandatory:"true" json:"schedules"`
	Action              Action                            `mandatory:"true" json:"action"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated         *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState      ScheduledTaskLifecycleStateEnum   `mandatory:"true" json:"lifecycleState"`
	TaskStatus          ScheduledTaskTaskStatusEnum       `mandatory:"false" json:"taskStatus,omitempty"`
	PauseReason         ScheduledTaskPauseReasonEnum      `mandatory:"false" json:"pauseReason,omitempty"`
	WorkRequestId       *string                           `mandatory:"false" json:"workRequestId"`
	NumOccurrences      *int64                            `mandatory:"false" json:"numOccurrences"`
	TimeOfNextExecution *common.SDKTime                   `mandatory:"false" json:"timeOfNextExecution"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Kind                string                            `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *scheduledtask) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduledtask scheduledtask
	s := struct {
		Model Unmarshalerscheduledtask
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.TaskType = s.Model.TaskType
	m.Schedules = s.Model.Schedules
	m.Action = s.Model.Action
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.TaskStatus = s.Model.TaskStatus
	m.PauseReason = s.Model.PauseReason
	m.WorkRequestId = s.Model.WorkRequestId
	m.NumOccurrences = s.Model.NumOccurrences
	m.TimeOfNextExecution = s.Model.TimeOfNextExecution
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduledtask) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STANDARD":
		mm := StandardTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m scheduledtask) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m scheduledtask) GetDisplayName() *string {
	return m.DisplayName
}

//GetTaskType returns TaskType
func (m scheduledtask) GetTaskType() TaskTypeEnum {
	return m.TaskType
}

//GetSchedules returns Schedules
func (m scheduledtask) GetSchedules() json.RawMessage {
	return m.Schedules
}

//GetAction returns Action
func (m scheduledtask) GetAction() Action {
	return m.Action
}

//GetCompartmentId returns CompartmentId
func (m scheduledtask) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m scheduledtask) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m scheduledtask) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m scheduledtask) GetLifecycleState() ScheduledTaskLifecycleStateEnum {
	return m.LifecycleState
}

//GetTaskStatus returns TaskStatus
func (m scheduledtask) GetTaskStatus() ScheduledTaskTaskStatusEnum {
	return m.TaskStatus
}

//GetPauseReason returns PauseReason
func (m scheduledtask) GetPauseReason() ScheduledTaskPauseReasonEnum {
	return m.PauseReason
}

//GetWorkRequestId returns WorkRequestId
func (m scheduledtask) GetWorkRequestId() *string {
	return m.WorkRequestId
}

//GetNumOccurrences returns NumOccurrences
func (m scheduledtask) GetNumOccurrences() *int64 {
	return m.NumOccurrences
}

//GetTimeOfNextExecution returns TimeOfNextExecution
func (m scheduledtask) GetTimeOfNextExecution() *common.SDKTime {
	return m.TimeOfNextExecution
}

//GetFreeformTags returns FreeformTags
func (m scheduledtask) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m scheduledtask) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m scheduledtask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduledtask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledTaskLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScheduledTaskTaskStatusEnum(string(m.TaskStatus)); !ok && m.TaskStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskStatus: %s. Supported values are: %s.", m.TaskStatus, strings.Join(GetScheduledTaskTaskStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledTaskPauseReasonEnum(string(m.PauseReason)); !ok && m.PauseReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PauseReason: %s. Supported values are: %s.", m.PauseReason, strings.Join(GetScheduledTaskPauseReasonEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledTaskTaskStatusEnum Enum with underlying type: string
type ScheduledTaskTaskStatusEnum string

// Set of constants representing the allowable values for ScheduledTaskTaskStatusEnum
const (
	ScheduledTaskTaskStatusReady     ScheduledTaskTaskStatusEnum = "READY"
	ScheduledTaskTaskStatusPaused    ScheduledTaskTaskStatusEnum = "PAUSED"
	ScheduledTaskTaskStatusCompleted ScheduledTaskTaskStatusEnum = "COMPLETED"
	ScheduledTaskTaskStatusBlocked   ScheduledTaskTaskStatusEnum = "BLOCKED"
)

var mappingScheduledTaskTaskStatusEnum = map[string]ScheduledTaskTaskStatusEnum{
	"READY":     ScheduledTaskTaskStatusReady,
	"PAUSED":    ScheduledTaskTaskStatusPaused,
	"COMPLETED": ScheduledTaskTaskStatusCompleted,
	"BLOCKED":   ScheduledTaskTaskStatusBlocked,
}

// GetScheduledTaskTaskStatusEnumValues Enumerates the set of values for ScheduledTaskTaskStatusEnum
func GetScheduledTaskTaskStatusEnumValues() []ScheduledTaskTaskStatusEnum {
	values := make([]ScheduledTaskTaskStatusEnum, 0)
	for _, v := range mappingScheduledTaskTaskStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskTaskStatusEnumStringValues Enumerates the set of values in String for ScheduledTaskTaskStatusEnum
func GetScheduledTaskTaskStatusEnumStringValues() []string {
	return []string{
		"READY",
		"PAUSED",
		"COMPLETED",
		"BLOCKED",
	}
}

// GetMappingScheduledTaskTaskStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskTaskStatusEnum(val string) (ScheduledTaskTaskStatusEnum, bool) {
	mappingScheduledTaskTaskStatusEnumIgnoreCase := make(map[string]ScheduledTaskTaskStatusEnum)
	for k, v := range mappingScheduledTaskTaskStatusEnum {
		mappingScheduledTaskTaskStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskTaskStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledTaskPauseReasonEnum Enum with underlying type: string
type ScheduledTaskPauseReasonEnum string

// Set of constants representing the allowable values for ScheduledTaskPauseReasonEnum
const (
	ScheduledTaskPauseReasonMetricExtractionNotValid ScheduledTaskPauseReasonEnum = "METRIC_EXTRACTION_NOT_VALID"
	ScheduledTaskPauseReasonSavedSearchNotValid      ScheduledTaskPauseReasonEnum = "SAVED_SEARCH_NOT_VALID"
	ScheduledTaskPauseReasonSavedSearchNotFound      ScheduledTaskPauseReasonEnum = "SAVED_SEARCH_NOT_FOUND"
	ScheduledTaskPauseReasonQueryStringNotValid      ScheduledTaskPauseReasonEnum = "QUERY_STRING_NOT_VALID"
	ScheduledTaskPauseReasonUserAction               ScheduledTaskPauseReasonEnum = "USER_ACTION"
	ScheduledTaskPauseReasonTenancyLifecycle         ScheduledTaskPauseReasonEnum = "TENANCY_LIFECYCLE"
	ScheduledTaskPauseReasonPurgeResourceNotFound    ScheduledTaskPauseReasonEnum = "PURGE_RESOURCE_NOT_FOUND"
)

var mappingScheduledTaskPauseReasonEnum = map[string]ScheduledTaskPauseReasonEnum{
	"METRIC_EXTRACTION_NOT_VALID": ScheduledTaskPauseReasonMetricExtractionNotValid,
	"SAVED_SEARCH_NOT_VALID":      ScheduledTaskPauseReasonSavedSearchNotValid,
	"SAVED_SEARCH_NOT_FOUND":      ScheduledTaskPauseReasonSavedSearchNotFound,
	"QUERY_STRING_NOT_VALID":      ScheduledTaskPauseReasonQueryStringNotValid,
	"USER_ACTION":                 ScheduledTaskPauseReasonUserAction,
	"TENANCY_LIFECYCLE":           ScheduledTaskPauseReasonTenancyLifecycle,
	"PURGE_RESOURCE_NOT_FOUND":    ScheduledTaskPauseReasonPurgeResourceNotFound,
}

// GetScheduledTaskPauseReasonEnumValues Enumerates the set of values for ScheduledTaskPauseReasonEnum
func GetScheduledTaskPauseReasonEnumValues() []ScheduledTaskPauseReasonEnum {
	values := make([]ScheduledTaskPauseReasonEnum, 0)
	for _, v := range mappingScheduledTaskPauseReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskPauseReasonEnumStringValues Enumerates the set of values in String for ScheduledTaskPauseReasonEnum
func GetScheduledTaskPauseReasonEnumStringValues() []string {
	return []string{
		"METRIC_EXTRACTION_NOT_VALID",
		"SAVED_SEARCH_NOT_VALID",
		"SAVED_SEARCH_NOT_FOUND",
		"QUERY_STRING_NOT_VALID",
		"USER_ACTION",
		"TENANCY_LIFECYCLE",
		"PURGE_RESOURCE_NOT_FOUND",
	}
}

// GetMappingScheduledTaskPauseReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskPauseReasonEnum(val string) (ScheduledTaskPauseReasonEnum, bool) {
	mappingScheduledTaskPauseReasonEnumIgnoreCase := make(map[string]ScheduledTaskPauseReasonEnum)
	for k, v := range mappingScheduledTaskPauseReasonEnum {
		mappingScheduledTaskPauseReasonEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskPauseReasonEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledTaskLifecycleStateEnum Enum with underlying type: string
type ScheduledTaskLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledTaskLifecycleStateEnum
const (
	ScheduledTaskLifecycleStateActive  ScheduledTaskLifecycleStateEnum = "ACTIVE"
	ScheduledTaskLifecycleStateDeleted ScheduledTaskLifecycleStateEnum = "DELETED"
)

var mappingScheduledTaskLifecycleStateEnum = map[string]ScheduledTaskLifecycleStateEnum{
	"ACTIVE":  ScheduledTaskLifecycleStateActive,
	"DELETED": ScheduledTaskLifecycleStateDeleted,
}

// GetScheduledTaskLifecycleStateEnumValues Enumerates the set of values for ScheduledTaskLifecycleStateEnum
func GetScheduledTaskLifecycleStateEnumValues() []ScheduledTaskLifecycleStateEnum {
	values := make([]ScheduledTaskLifecycleStateEnum, 0)
	for _, v := range mappingScheduledTaskLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledTaskLifecycleStateEnum
func GetScheduledTaskLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingScheduledTaskLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskLifecycleStateEnum(val string) (ScheduledTaskLifecycleStateEnum, bool) {
	mappingScheduledTaskLifecycleStateEnumIgnoreCase := make(map[string]ScheduledTaskLifecycleStateEnum)
	for k, v := range mappingScheduledTaskLifecycleStateEnum {
		mappingScheduledTaskLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledTaskKindEnum Enum with underlying type: string
type ScheduledTaskKindEnum string

// Set of constants representing the allowable values for ScheduledTaskKindEnum
const (
	ScheduledTaskKindAcceleration ScheduledTaskKindEnum = "ACCELERATION"
	ScheduledTaskKindStandard     ScheduledTaskKindEnum = "STANDARD"
)

var mappingScheduledTaskKindEnum = map[string]ScheduledTaskKindEnum{
	"ACCELERATION": ScheduledTaskKindAcceleration,
	"STANDARD":     ScheduledTaskKindStandard,
}

// GetScheduledTaskKindEnumValues Enumerates the set of values for ScheduledTaskKindEnum
func GetScheduledTaskKindEnumValues() []ScheduledTaskKindEnum {
	values := make([]ScheduledTaskKindEnum, 0)
	for _, v := range mappingScheduledTaskKindEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledTaskKindEnumStringValues Enumerates the set of values in String for ScheduledTaskKindEnum
func GetScheduledTaskKindEnumStringValues() []string {
	return []string{
		"ACCELERATION",
		"STANDARD",
	}
}

// GetMappingScheduledTaskKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledTaskKindEnum(val string) (ScheduledTaskKindEnum, bool) {
	mappingScheduledTaskKindEnumIgnoreCase := make(map[string]ScheduledTaskKindEnum)
	for k, v := range mappingScheduledTaskKindEnum {
		mappingScheduledTaskKindEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduledTaskKindEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
