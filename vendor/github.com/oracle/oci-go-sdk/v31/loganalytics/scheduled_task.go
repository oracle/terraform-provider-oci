// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ScheduledTask Log analytics scheduled task resource.
type ScheduledTask struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the data plane resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name that is changeable and that does not have to be unique.
	// Format: a leading alphanumeric, followed by zero or more
	// alphanumerics, underscores, spaces, backslashes, or hyphens in any order).
	// No trailing spaces allowed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Task type.
	TaskType TaskTypeEnum `mandatory:"true" json:"taskType"`

	// Schedules.
	Schedules []Schedule `mandatory:"true" json:"schedules"`

	Action Action `mandatory:"true" json:"action"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the scheduled task was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the scheduled task was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the scheduled task.
	LifecycleState ScheduledTaskLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Status of the scheduled task.
	TaskStatus ScheduledTaskTaskStatusEnum `mandatory:"false" json:"taskStatus,omitempty"`

	// most recent Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`

	// Number of execution occurrences.
	NumOccurrences *int64 `mandatory:"false" json:"numOccurrences"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ScheduledTask) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ScheduledTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskStatus     ScheduledTaskTaskStatusEnum       `json:"taskStatus"`
		WorkRequestId  *string                           `json:"workRequestId"`
		NumOccurrences *int64                            `json:"numOccurrences"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		Id             *string                           `json:"id"`
		DisplayName    *string                           `json:"displayName"`
		TaskType       TaskTypeEnum                      `json:"taskType"`
		Schedules      []schedule                        `json:"schedules"`
		Action         action                            `json:"action"`
		CompartmentId  *string                           `json:"compartmentId"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState ScheduledTaskLifecycleStateEnum   `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskStatus = model.TaskStatus

	m.WorkRequestId = model.WorkRequestId

	m.NumOccurrences = model.NumOccurrences

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

// ScheduledTaskTaskStatusEnum Enum with underlying type: string
type ScheduledTaskTaskStatusEnum string

// Set of constants representing the allowable values for ScheduledTaskTaskStatusEnum
const (
	ScheduledTaskTaskStatusReady     ScheduledTaskTaskStatusEnum = "READY"
	ScheduledTaskTaskStatusPaused    ScheduledTaskTaskStatusEnum = "PAUSED"
	ScheduledTaskTaskStatusCompleted ScheduledTaskTaskStatusEnum = "COMPLETED"
	ScheduledTaskTaskStatusBlocked   ScheduledTaskTaskStatusEnum = "BLOCKED"
)

var mappingScheduledTaskTaskStatus = map[string]ScheduledTaskTaskStatusEnum{
	"READY":     ScheduledTaskTaskStatusReady,
	"PAUSED":    ScheduledTaskTaskStatusPaused,
	"COMPLETED": ScheduledTaskTaskStatusCompleted,
	"BLOCKED":   ScheduledTaskTaskStatusBlocked,
}

// GetScheduledTaskTaskStatusEnumValues Enumerates the set of values for ScheduledTaskTaskStatusEnum
func GetScheduledTaskTaskStatusEnumValues() []ScheduledTaskTaskStatusEnum {
	values := make([]ScheduledTaskTaskStatusEnum, 0)
	for _, v := range mappingScheduledTaskTaskStatus {
		values = append(values, v)
	}
	return values
}

// ScheduledTaskLifecycleStateEnum Enum with underlying type: string
type ScheduledTaskLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledTaskLifecycleStateEnum
const (
	ScheduledTaskLifecycleStateActive  ScheduledTaskLifecycleStateEnum = "ACTIVE"
	ScheduledTaskLifecycleStateDeleted ScheduledTaskLifecycleStateEnum = "DELETED"
)

var mappingScheduledTaskLifecycleState = map[string]ScheduledTaskLifecycleStateEnum{
	"ACTIVE":  ScheduledTaskLifecycleStateActive,
	"DELETED": ScheduledTaskLifecycleStateDeleted,
}

// GetScheduledTaskLifecycleStateEnumValues Enumerates the set of values for ScheduledTaskLifecycleStateEnum
func GetScheduledTaskLifecycleStateEnumValues() []ScheduledTaskLifecycleStateEnum {
	values := make([]ScheduledTaskLifecycleStateEnum, 0)
	for _, v := range mappingScheduledTaskLifecycleState {
		values = append(values, v)
	}
	return values
}
