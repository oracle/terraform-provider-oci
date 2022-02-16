// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TaskScheduleSummary The tsk schedule summary information.
type TaskScheduleSummary struct {

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

	// Whether the task schedule is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The number of retries.
	NumberOfRetries *int `mandatory:"false" json:"numberOfRetries"`

	// The retry delay, the unit for measurement is in the property retry delay unit.
	RetryDelay *float64 `mandatory:"false" json:"retryDelay"`

	// The unit for the retry delay.
	RetryDelayUnit TaskScheduleSummaryRetryDelayUnitEnum `mandatory:"false" json:"retryDelayUnit,omitempty"`

	// The start time in milliseconds.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The end time in milliseconds.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	// Whether the backfill is enabled.
	IsBackfillEnabled *bool `mandatory:"false" json:"isBackfillEnabled"`

	// The authorization mode for the task.
	AuthMode TaskScheduleSummaryAuthModeEnum `mandatory:"false" json:"authMode,omitempty"`

	// The expected duration of the task execution.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	// The expected duration unit of the task execution.
	ExpectedDurationUnit TaskScheduleSummaryExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	// The time for next run in milliseconds.
	NextRunTimeMillis *int64 `mandatory:"false" json:"nextRunTimeMillis"`

	LastRunDetails *LastRunDetails `mandatory:"false" json:"lastRunDetails"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskScheduleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskScheduleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTaskScheduleSummaryRetryDelayUnitEnum(string(m.RetryDelayUnit)); !ok && m.RetryDelayUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetryDelayUnit: %s. Supported values are: %s.", m.RetryDelayUnit, strings.Join(GetTaskScheduleSummaryRetryDelayUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskScheduleSummaryAuthModeEnum(string(m.AuthMode)); !ok && m.AuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthMode: %s. Supported values are: %s.", m.AuthMode, strings.Join(GetTaskScheduleSummaryAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskScheduleSummaryExpectedDurationUnitEnum(string(m.ExpectedDurationUnit)); !ok && m.ExpectedDurationUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedDurationUnit: %s. Supported values are: %s.", m.ExpectedDurationUnit, strings.Join(GetTaskScheduleSummaryExpectedDurationUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskScheduleSummaryRetryDelayUnitEnum Enum with underlying type: string
type TaskScheduleSummaryRetryDelayUnitEnum string

// Set of constants representing the allowable values for TaskScheduleSummaryRetryDelayUnitEnum
const (
	TaskScheduleSummaryRetryDelayUnitSeconds TaskScheduleSummaryRetryDelayUnitEnum = "SECONDS"
	TaskScheduleSummaryRetryDelayUnitMinutes TaskScheduleSummaryRetryDelayUnitEnum = "MINUTES"
	TaskScheduleSummaryRetryDelayUnitHours   TaskScheduleSummaryRetryDelayUnitEnum = "HOURS"
	TaskScheduleSummaryRetryDelayUnitDays    TaskScheduleSummaryRetryDelayUnitEnum = "DAYS"
)

var mappingTaskScheduleSummaryRetryDelayUnitEnum = map[string]TaskScheduleSummaryRetryDelayUnitEnum{
	"SECONDS": TaskScheduleSummaryRetryDelayUnitSeconds,
	"MINUTES": TaskScheduleSummaryRetryDelayUnitMinutes,
	"HOURS":   TaskScheduleSummaryRetryDelayUnitHours,
	"DAYS":    TaskScheduleSummaryRetryDelayUnitDays,
}

// GetTaskScheduleSummaryRetryDelayUnitEnumValues Enumerates the set of values for TaskScheduleSummaryRetryDelayUnitEnum
func GetTaskScheduleSummaryRetryDelayUnitEnumValues() []TaskScheduleSummaryRetryDelayUnitEnum {
	values := make([]TaskScheduleSummaryRetryDelayUnitEnum, 0)
	for _, v := range mappingTaskScheduleSummaryRetryDelayUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskScheduleSummaryRetryDelayUnitEnumStringValues Enumerates the set of values in String for TaskScheduleSummaryRetryDelayUnitEnum
func GetTaskScheduleSummaryRetryDelayUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingTaskScheduleSummaryRetryDelayUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskScheduleSummaryRetryDelayUnitEnum(val string) (TaskScheduleSummaryRetryDelayUnitEnum, bool) {
	mappingTaskScheduleSummaryRetryDelayUnitEnumIgnoreCase := make(map[string]TaskScheduleSummaryRetryDelayUnitEnum)
	for k, v := range mappingTaskScheduleSummaryRetryDelayUnitEnum {
		mappingTaskScheduleSummaryRetryDelayUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskScheduleSummaryRetryDelayUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskScheduleSummaryAuthModeEnum Enum with underlying type: string
type TaskScheduleSummaryAuthModeEnum string

// Set of constants representing the allowable values for TaskScheduleSummaryAuthModeEnum
const (
	TaskScheduleSummaryAuthModeObo               TaskScheduleSummaryAuthModeEnum = "OBO"
	TaskScheduleSummaryAuthModeResourcePrincipal TaskScheduleSummaryAuthModeEnum = "RESOURCE_PRINCIPAL"
	TaskScheduleSummaryAuthModeUserCertificate   TaskScheduleSummaryAuthModeEnum = "USER_CERTIFICATE"
)

var mappingTaskScheduleSummaryAuthModeEnum = map[string]TaskScheduleSummaryAuthModeEnum{
	"OBO":                TaskScheduleSummaryAuthModeObo,
	"RESOURCE_PRINCIPAL": TaskScheduleSummaryAuthModeResourcePrincipal,
	"USER_CERTIFICATE":   TaskScheduleSummaryAuthModeUserCertificate,
}

// GetTaskScheduleSummaryAuthModeEnumValues Enumerates the set of values for TaskScheduleSummaryAuthModeEnum
func GetTaskScheduleSummaryAuthModeEnumValues() []TaskScheduleSummaryAuthModeEnum {
	values := make([]TaskScheduleSummaryAuthModeEnum, 0)
	for _, v := range mappingTaskScheduleSummaryAuthModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskScheduleSummaryAuthModeEnumStringValues Enumerates the set of values in String for TaskScheduleSummaryAuthModeEnum
func GetTaskScheduleSummaryAuthModeEnumStringValues() []string {
	return []string{
		"OBO",
		"RESOURCE_PRINCIPAL",
		"USER_CERTIFICATE",
	}
}

// GetMappingTaskScheduleSummaryAuthModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskScheduleSummaryAuthModeEnum(val string) (TaskScheduleSummaryAuthModeEnum, bool) {
	mappingTaskScheduleSummaryAuthModeEnumIgnoreCase := make(map[string]TaskScheduleSummaryAuthModeEnum)
	for k, v := range mappingTaskScheduleSummaryAuthModeEnum {
		mappingTaskScheduleSummaryAuthModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskScheduleSummaryAuthModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskScheduleSummaryExpectedDurationUnitEnum Enum with underlying type: string
type TaskScheduleSummaryExpectedDurationUnitEnum string

// Set of constants representing the allowable values for TaskScheduleSummaryExpectedDurationUnitEnum
const (
	TaskScheduleSummaryExpectedDurationUnitSeconds TaskScheduleSummaryExpectedDurationUnitEnum = "SECONDS"
	TaskScheduleSummaryExpectedDurationUnitMinutes TaskScheduleSummaryExpectedDurationUnitEnum = "MINUTES"
	TaskScheduleSummaryExpectedDurationUnitHours   TaskScheduleSummaryExpectedDurationUnitEnum = "HOURS"
	TaskScheduleSummaryExpectedDurationUnitDays    TaskScheduleSummaryExpectedDurationUnitEnum = "DAYS"
)

var mappingTaskScheduleSummaryExpectedDurationUnitEnum = map[string]TaskScheduleSummaryExpectedDurationUnitEnum{
	"SECONDS": TaskScheduleSummaryExpectedDurationUnitSeconds,
	"MINUTES": TaskScheduleSummaryExpectedDurationUnitMinutes,
	"HOURS":   TaskScheduleSummaryExpectedDurationUnitHours,
	"DAYS":    TaskScheduleSummaryExpectedDurationUnitDays,
}

// GetTaskScheduleSummaryExpectedDurationUnitEnumValues Enumerates the set of values for TaskScheduleSummaryExpectedDurationUnitEnum
func GetTaskScheduleSummaryExpectedDurationUnitEnumValues() []TaskScheduleSummaryExpectedDurationUnitEnum {
	values := make([]TaskScheduleSummaryExpectedDurationUnitEnum, 0)
	for _, v := range mappingTaskScheduleSummaryExpectedDurationUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskScheduleSummaryExpectedDurationUnitEnumStringValues Enumerates the set of values in String for TaskScheduleSummaryExpectedDurationUnitEnum
func GetTaskScheduleSummaryExpectedDurationUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingTaskScheduleSummaryExpectedDurationUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskScheduleSummaryExpectedDurationUnitEnum(val string) (TaskScheduleSummaryExpectedDurationUnitEnum, bool) {
	mappingTaskScheduleSummaryExpectedDurationUnitEnumIgnoreCase := make(map[string]TaskScheduleSummaryExpectedDurationUnitEnum)
	for k, v := range mappingTaskScheduleSummaryExpectedDurationUnitEnum {
		mappingTaskScheduleSummaryExpectedDurationUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskScheduleSummaryExpectedDurationUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
