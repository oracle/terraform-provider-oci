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

// CreateTaskScheduleDetails The create task details.
type CreateTaskScheduleDetails struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	// Generated key that can be used in API calls to identify taskSchedule. On scenarios where reference to the taskSchedule is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	ScheduleRef *Schedule `mandatory:"false" json:"scheduleRef"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// Whether the task schedule is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The number of retries.
	NumberOfRetries *int `mandatory:"false" json:"numberOfRetries"`

	// The retry delay, the unit for measurement is in the property retry delay unit.
	RetryDelay *float64 `mandatory:"false" json:"retryDelay"`

	// The unit for the retry delay.
	RetryDelayUnit CreateTaskScheduleDetailsRetryDelayUnitEnum `mandatory:"false" json:"retryDelayUnit,omitempty"`

	// The start time in milliseconds.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The end time in milliseconds.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	// Whether the backfill is enabled.
	IsBackfillEnabled *bool `mandatory:"false" json:"isBackfillEnabled"`

	// The authorization mode for the task.
	AuthMode CreateTaskScheduleDetailsAuthModeEnum `mandatory:"false" json:"authMode,omitempty"`

	// The expected duration of the task execution.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	// The expected duration unit of the task execution.
	ExpectedDurationUnit CreateTaskScheduleDetailsExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateTaskScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTaskScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateTaskScheduleDetailsRetryDelayUnitEnum(string(m.RetryDelayUnit)); !ok && m.RetryDelayUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetryDelayUnit: %s. Supported values are: %s.", m.RetryDelayUnit, strings.Join(GetCreateTaskScheduleDetailsRetryDelayUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateTaskScheduleDetailsAuthModeEnum(string(m.AuthMode)); !ok && m.AuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthMode: %s. Supported values are: %s.", m.AuthMode, strings.Join(GetCreateTaskScheduleDetailsAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateTaskScheduleDetailsExpectedDurationUnitEnum(string(m.ExpectedDurationUnit)); !ok && m.ExpectedDurationUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedDurationUnit: %s. Supported values are: %s.", m.ExpectedDurationUnit, strings.Join(GetCreateTaskScheduleDetailsExpectedDurationUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTaskScheduleDetailsRetryDelayUnitEnum Enum with underlying type: string
type CreateTaskScheduleDetailsRetryDelayUnitEnum string

// Set of constants representing the allowable values for CreateTaskScheduleDetailsRetryDelayUnitEnum
const (
	CreateTaskScheduleDetailsRetryDelayUnitSeconds CreateTaskScheduleDetailsRetryDelayUnitEnum = "SECONDS"
	CreateTaskScheduleDetailsRetryDelayUnitMinutes CreateTaskScheduleDetailsRetryDelayUnitEnum = "MINUTES"
	CreateTaskScheduleDetailsRetryDelayUnitHours   CreateTaskScheduleDetailsRetryDelayUnitEnum = "HOURS"
	CreateTaskScheduleDetailsRetryDelayUnitDays    CreateTaskScheduleDetailsRetryDelayUnitEnum = "DAYS"
)

var mappingCreateTaskScheduleDetailsRetryDelayUnitEnum = map[string]CreateTaskScheduleDetailsRetryDelayUnitEnum{
	"SECONDS": CreateTaskScheduleDetailsRetryDelayUnitSeconds,
	"MINUTES": CreateTaskScheduleDetailsRetryDelayUnitMinutes,
	"HOURS":   CreateTaskScheduleDetailsRetryDelayUnitHours,
	"DAYS":    CreateTaskScheduleDetailsRetryDelayUnitDays,
}

// GetCreateTaskScheduleDetailsRetryDelayUnitEnumValues Enumerates the set of values for CreateTaskScheduleDetailsRetryDelayUnitEnum
func GetCreateTaskScheduleDetailsRetryDelayUnitEnumValues() []CreateTaskScheduleDetailsRetryDelayUnitEnum {
	values := make([]CreateTaskScheduleDetailsRetryDelayUnitEnum, 0)
	for _, v := range mappingCreateTaskScheduleDetailsRetryDelayUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskScheduleDetailsRetryDelayUnitEnumStringValues Enumerates the set of values in String for CreateTaskScheduleDetailsRetryDelayUnitEnum
func GetCreateTaskScheduleDetailsRetryDelayUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingCreateTaskScheduleDetailsRetryDelayUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskScheduleDetailsRetryDelayUnitEnum(val string) (CreateTaskScheduleDetailsRetryDelayUnitEnum, bool) {
	mappingCreateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase := make(map[string]CreateTaskScheduleDetailsRetryDelayUnitEnum)
	for k, v := range mappingCreateTaskScheduleDetailsRetryDelayUnitEnum {
		mappingCreateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// CreateTaskScheduleDetailsAuthModeEnum Enum with underlying type: string
type CreateTaskScheduleDetailsAuthModeEnum string

// Set of constants representing the allowable values for CreateTaskScheduleDetailsAuthModeEnum
const (
	CreateTaskScheduleDetailsAuthModeObo               CreateTaskScheduleDetailsAuthModeEnum = "OBO"
	CreateTaskScheduleDetailsAuthModeResourcePrincipal CreateTaskScheduleDetailsAuthModeEnum = "RESOURCE_PRINCIPAL"
	CreateTaskScheduleDetailsAuthModeUserCertificate   CreateTaskScheduleDetailsAuthModeEnum = "USER_CERTIFICATE"
)

var mappingCreateTaskScheduleDetailsAuthModeEnum = map[string]CreateTaskScheduleDetailsAuthModeEnum{
	"OBO":                CreateTaskScheduleDetailsAuthModeObo,
	"RESOURCE_PRINCIPAL": CreateTaskScheduleDetailsAuthModeResourcePrincipal,
	"USER_CERTIFICATE":   CreateTaskScheduleDetailsAuthModeUserCertificate,
}

// GetCreateTaskScheduleDetailsAuthModeEnumValues Enumerates the set of values for CreateTaskScheduleDetailsAuthModeEnum
func GetCreateTaskScheduleDetailsAuthModeEnumValues() []CreateTaskScheduleDetailsAuthModeEnum {
	values := make([]CreateTaskScheduleDetailsAuthModeEnum, 0)
	for _, v := range mappingCreateTaskScheduleDetailsAuthModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskScheduleDetailsAuthModeEnumStringValues Enumerates the set of values in String for CreateTaskScheduleDetailsAuthModeEnum
func GetCreateTaskScheduleDetailsAuthModeEnumStringValues() []string {
	return []string{
		"OBO",
		"RESOURCE_PRINCIPAL",
		"USER_CERTIFICATE",
	}
}

// GetMappingCreateTaskScheduleDetailsAuthModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskScheduleDetailsAuthModeEnum(val string) (CreateTaskScheduleDetailsAuthModeEnum, bool) {
	mappingCreateTaskScheduleDetailsAuthModeEnumIgnoreCase := make(map[string]CreateTaskScheduleDetailsAuthModeEnum)
	for k, v := range mappingCreateTaskScheduleDetailsAuthModeEnum {
		mappingCreateTaskScheduleDetailsAuthModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateTaskScheduleDetailsAuthModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// CreateTaskScheduleDetailsExpectedDurationUnitEnum Enum with underlying type: string
type CreateTaskScheduleDetailsExpectedDurationUnitEnum string

// Set of constants representing the allowable values for CreateTaskScheduleDetailsExpectedDurationUnitEnum
const (
	CreateTaskScheduleDetailsExpectedDurationUnitSeconds CreateTaskScheduleDetailsExpectedDurationUnitEnum = "SECONDS"
	CreateTaskScheduleDetailsExpectedDurationUnitMinutes CreateTaskScheduleDetailsExpectedDurationUnitEnum = "MINUTES"
	CreateTaskScheduleDetailsExpectedDurationUnitHours   CreateTaskScheduleDetailsExpectedDurationUnitEnum = "HOURS"
	CreateTaskScheduleDetailsExpectedDurationUnitDays    CreateTaskScheduleDetailsExpectedDurationUnitEnum = "DAYS"
)

var mappingCreateTaskScheduleDetailsExpectedDurationUnitEnum = map[string]CreateTaskScheduleDetailsExpectedDurationUnitEnum{
	"SECONDS": CreateTaskScheduleDetailsExpectedDurationUnitSeconds,
	"MINUTES": CreateTaskScheduleDetailsExpectedDurationUnitMinutes,
	"HOURS":   CreateTaskScheduleDetailsExpectedDurationUnitHours,
	"DAYS":    CreateTaskScheduleDetailsExpectedDurationUnitDays,
}

// GetCreateTaskScheduleDetailsExpectedDurationUnitEnumValues Enumerates the set of values for CreateTaskScheduleDetailsExpectedDurationUnitEnum
func GetCreateTaskScheduleDetailsExpectedDurationUnitEnumValues() []CreateTaskScheduleDetailsExpectedDurationUnitEnum {
	values := make([]CreateTaskScheduleDetailsExpectedDurationUnitEnum, 0)
	for _, v := range mappingCreateTaskScheduleDetailsExpectedDurationUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskScheduleDetailsExpectedDurationUnitEnumStringValues Enumerates the set of values in String for CreateTaskScheduleDetailsExpectedDurationUnitEnum
func GetCreateTaskScheduleDetailsExpectedDurationUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingCreateTaskScheduleDetailsExpectedDurationUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskScheduleDetailsExpectedDurationUnitEnum(val string) (CreateTaskScheduleDetailsExpectedDurationUnitEnum, bool) {
	mappingCreateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase := make(map[string]CreateTaskScheduleDetailsExpectedDurationUnitEnum)
	for k, v := range mappingCreateTaskScheduleDetailsExpectedDurationUnitEnum {
		mappingCreateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
