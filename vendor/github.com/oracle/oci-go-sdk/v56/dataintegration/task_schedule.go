// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TaskSchedule A model that holds Schedule and other information required for scheduling a task.
type TaskSchedule struct {

	// Generated key that can be used in API calls to identify taskSchedule. On scenarios where reference to the taskSchedule is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ScheduleRef *Schedule `mandatory:"false" json:"scheduleRef"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// Whether the schedule is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The number of retry attempts.
	RetryAttempts *int `mandatory:"false" json:"retryAttempts"`

	// The unit for the retry delay.
	RetryDelayUnit TaskScheduleRetryDelayUnitEnum `mandatory:"false" json:"retryDelayUnit,omitempty"`

	// The retry delay, the unit for measurement is in the property retry delay unit.
	RetryDelay *float64 `mandatory:"false" json:"retryDelay"`

	// The start time in milliseconds.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The end time in milliseconds.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	// Whether the backfill is enabled
	IsBackfillEnabled *bool `mandatory:"false" json:"isBackfillEnabled"`

	// The authorization mode for the task.
	AuthMode TaskScheduleAuthModeEnum `mandatory:"false" json:"authMode,omitempty"`

	// The expected duration of the task execution.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	// The expected duration unit of the task execution.
	ExpectedDurationUnit TaskScheduleExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	LastRunDetails *LastRunDetails `mandatory:"false" json:"lastRunDetails"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskSchedule) String() string {
	return common.PointerString(m)
}

// TaskScheduleRetryDelayUnitEnum Enum with underlying type: string
type TaskScheduleRetryDelayUnitEnum string

// Set of constants representing the allowable values for TaskScheduleRetryDelayUnitEnum
const (
	TaskScheduleRetryDelayUnitSeconds TaskScheduleRetryDelayUnitEnum = "SECONDS"
	TaskScheduleRetryDelayUnitMinutes TaskScheduleRetryDelayUnitEnum = "MINUTES"
	TaskScheduleRetryDelayUnitHours   TaskScheduleRetryDelayUnitEnum = "HOURS"
	TaskScheduleRetryDelayUnitDays    TaskScheduleRetryDelayUnitEnum = "DAYS"
)

var mappingTaskScheduleRetryDelayUnit = map[string]TaskScheduleRetryDelayUnitEnum{
	"SECONDS": TaskScheduleRetryDelayUnitSeconds,
	"MINUTES": TaskScheduleRetryDelayUnitMinutes,
	"HOURS":   TaskScheduleRetryDelayUnitHours,
	"DAYS":    TaskScheduleRetryDelayUnitDays,
}

// GetTaskScheduleRetryDelayUnitEnumValues Enumerates the set of values for TaskScheduleRetryDelayUnitEnum
func GetTaskScheduleRetryDelayUnitEnumValues() []TaskScheduleRetryDelayUnitEnum {
	values := make([]TaskScheduleRetryDelayUnitEnum, 0)
	for _, v := range mappingTaskScheduleRetryDelayUnit {
		values = append(values, v)
	}
	return values
}

// TaskScheduleAuthModeEnum Enum with underlying type: string
type TaskScheduleAuthModeEnum string

// Set of constants representing the allowable values for TaskScheduleAuthModeEnum
const (
	TaskScheduleAuthModeObo               TaskScheduleAuthModeEnum = "OBO"
	TaskScheduleAuthModeResourcePrincipal TaskScheduleAuthModeEnum = "RESOURCE_PRINCIPAL"
	TaskScheduleAuthModeUserCertificate   TaskScheduleAuthModeEnum = "USER_CERTIFICATE"
)

var mappingTaskScheduleAuthMode = map[string]TaskScheduleAuthModeEnum{
	"OBO":                TaskScheduleAuthModeObo,
	"RESOURCE_PRINCIPAL": TaskScheduleAuthModeResourcePrincipal,
	"USER_CERTIFICATE":   TaskScheduleAuthModeUserCertificate,
}

// GetTaskScheduleAuthModeEnumValues Enumerates the set of values for TaskScheduleAuthModeEnum
func GetTaskScheduleAuthModeEnumValues() []TaskScheduleAuthModeEnum {
	values := make([]TaskScheduleAuthModeEnum, 0)
	for _, v := range mappingTaskScheduleAuthMode {
		values = append(values, v)
	}
	return values
}

// TaskScheduleExpectedDurationUnitEnum Enum with underlying type: string
type TaskScheduleExpectedDurationUnitEnum string

// Set of constants representing the allowable values for TaskScheduleExpectedDurationUnitEnum
const (
	TaskScheduleExpectedDurationUnitSeconds TaskScheduleExpectedDurationUnitEnum = "SECONDS"
	TaskScheduleExpectedDurationUnitMinutes TaskScheduleExpectedDurationUnitEnum = "MINUTES"
	TaskScheduleExpectedDurationUnitHours   TaskScheduleExpectedDurationUnitEnum = "HOURS"
	TaskScheduleExpectedDurationUnitDays    TaskScheduleExpectedDurationUnitEnum = "DAYS"
)

var mappingTaskScheduleExpectedDurationUnit = map[string]TaskScheduleExpectedDurationUnitEnum{
	"SECONDS": TaskScheduleExpectedDurationUnitSeconds,
	"MINUTES": TaskScheduleExpectedDurationUnitMinutes,
	"HOURS":   TaskScheduleExpectedDurationUnitHours,
	"DAYS":    TaskScheduleExpectedDurationUnitDays,
}

// GetTaskScheduleExpectedDurationUnitEnumValues Enumerates the set of values for TaskScheduleExpectedDurationUnitEnum
func GetTaskScheduleExpectedDurationUnitEnumValues() []TaskScheduleExpectedDurationUnitEnum {
	values := make([]TaskScheduleExpectedDurationUnitEnum, 0)
	for _, v := range mappingTaskScheduleExpectedDurationUnit {
		values = append(values, v)
	}
	return values
}
