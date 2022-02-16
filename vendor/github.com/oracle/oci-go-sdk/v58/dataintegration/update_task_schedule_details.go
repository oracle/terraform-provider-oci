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

// UpdateTaskScheduleDetails The update task details.
type UpdateTaskScheduleDetails struct {

	// Generated key that can be used in API calls to identify taskSchedule. On scenarios where reference to the taskSchedule is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ScheduleRef *Schedule `mandatory:"false" json:"scheduleRef"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The number of retries.
	NumberOfRetries *int `mandatory:"false" json:"numberOfRetries"`

	// The retry delay, the unit for measurement is in the property retry delay unit.
	RetryDelay *float64 `mandatory:"false" json:"retryDelay"`

	// The unit for the retry delay.
	RetryDelayUnit UpdateTaskScheduleDetailsRetryDelayUnitEnum `mandatory:"false" json:"retryDelayUnit,omitempty"`

	// The start time in milliseconds.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The end time in milliseconds.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	// Whether the backfill is enabled.
	IsBackfillEnabled *bool `mandatory:"false" json:"isBackfillEnabled"`

	// The authorization mode for the task.
	AuthMode UpdateTaskScheduleDetailsAuthModeEnum `mandatory:"false" json:"authMode,omitempty"`

	// The expected duration of the task.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	// The expected duration of the task.
	ExpectedDurationUnit UpdateTaskScheduleDetailsExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m UpdateTaskScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTaskScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTaskScheduleDetailsRetryDelayUnitEnum(string(m.RetryDelayUnit)); !ok && m.RetryDelayUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetryDelayUnit: %s. Supported values are: %s.", m.RetryDelayUnit, strings.Join(GetUpdateTaskScheduleDetailsRetryDelayUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTaskScheduleDetailsAuthModeEnum(string(m.AuthMode)); !ok && m.AuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthMode: %s. Supported values are: %s.", m.AuthMode, strings.Join(GetUpdateTaskScheduleDetailsAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum(string(m.ExpectedDurationUnit)); !ok && m.ExpectedDurationUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedDurationUnit: %s. Supported values are: %s.", m.ExpectedDurationUnit, strings.Join(GetUpdateTaskScheduleDetailsExpectedDurationUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTaskScheduleDetailsRetryDelayUnitEnum Enum with underlying type: string
type UpdateTaskScheduleDetailsRetryDelayUnitEnum string

// Set of constants representing the allowable values for UpdateTaskScheduleDetailsRetryDelayUnitEnum
const (
	UpdateTaskScheduleDetailsRetryDelayUnitSeconds UpdateTaskScheduleDetailsRetryDelayUnitEnum = "SECONDS"
	UpdateTaskScheduleDetailsRetryDelayUnitMinutes UpdateTaskScheduleDetailsRetryDelayUnitEnum = "MINUTES"
	UpdateTaskScheduleDetailsRetryDelayUnitHours   UpdateTaskScheduleDetailsRetryDelayUnitEnum = "HOURS"
	UpdateTaskScheduleDetailsRetryDelayUnitDays    UpdateTaskScheduleDetailsRetryDelayUnitEnum = "DAYS"
)

var mappingUpdateTaskScheduleDetailsRetryDelayUnitEnum = map[string]UpdateTaskScheduleDetailsRetryDelayUnitEnum{
	"SECONDS": UpdateTaskScheduleDetailsRetryDelayUnitSeconds,
	"MINUTES": UpdateTaskScheduleDetailsRetryDelayUnitMinutes,
	"HOURS":   UpdateTaskScheduleDetailsRetryDelayUnitHours,
	"DAYS":    UpdateTaskScheduleDetailsRetryDelayUnitDays,
}

// GetUpdateTaskScheduleDetailsRetryDelayUnitEnumValues Enumerates the set of values for UpdateTaskScheduleDetailsRetryDelayUnitEnum
func GetUpdateTaskScheduleDetailsRetryDelayUnitEnumValues() []UpdateTaskScheduleDetailsRetryDelayUnitEnum {
	values := make([]UpdateTaskScheduleDetailsRetryDelayUnitEnum, 0)
	for _, v := range mappingUpdateTaskScheduleDetailsRetryDelayUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskScheduleDetailsRetryDelayUnitEnumStringValues Enumerates the set of values in String for UpdateTaskScheduleDetailsRetryDelayUnitEnum
func GetUpdateTaskScheduleDetailsRetryDelayUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingUpdateTaskScheduleDetailsRetryDelayUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskScheduleDetailsRetryDelayUnitEnum(val string) (UpdateTaskScheduleDetailsRetryDelayUnitEnum, bool) {
	mappingUpdateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase := make(map[string]UpdateTaskScheduleDetailsRetryDelayUnitEnum)
	for k, v := range mappingUpdateTaskScheduleDetailsRetryDelayUnitEnum {
		mappingUpdateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTaskScheduleDetailsRetryDelayUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateTaskScheduleDetailsAuthModeEnum Enum with underlying type: string
type UpdateTaskScheduleDetailsAuthModeEnum string

// Set of constants representing the allowable values for UpdateTaskScheduleDetailsAuthModeEnum
const (
	UpdateTaskScheduleDetailsAuthModeObo               UpdateTaskScheduleDetailsAuthModeEnum = "OBO"
	UpdateTaskScheduleDetailsAuthModeResourcePrincipal UpdateTaskScheduleDetailsAuthModeEnum = "RESOURCE_PRINCIPAL"
	UpdateTaskScheduleDetailsAuthModeUserCertificate   UpdateTaskScheduleDetailsAuthModeEnum = "USER_CERTIFICATE"
)

var mappingUpdateTaskScheduleDetailsAuthModeEnum = map[string]UpdateTaskScheduleDetailsAuthModeEnum{
	"OBO":                UpdateTaskScheduleDetailsAuthModeObo,
	"RESOURCE_PRINCIPAL": UpdateTaskScheduleDetailsAuthModeResourcePrincipal,
	"USER_CERTIFICATE":   UpdateTaskScheduleDetailsAuthModeUserCertificate,
}

// GetUpdateTaskScheduleDetailsAuthModeEnumValues Enumerates the set of values for UpdateTaskScheduleDetailsAuthModeEnum
func GetUpdateTaskScheduleDetailsAuthModeEnumValues() []UpdateTaskScheduleDetailsAuthModeEnum {
	values := make([]UpdateTaskScheduleDetailsAuthModeEnum, 0)
	for _, v := range mappingUpdateTaskScheduleDetailsAuthModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskScheduleDetailsAuthModeEnumStringValues Enumerates the set of values in String for UpdateTaskScheduleDetailsAuthModeEnum
func GetUpdateTaskScheduleDetailsAuthModeEnumStringValues() []string {
	return []string{
		"OBO",
		"RESOURCE_PRINCIPAL",
		"USER_CERTIFICATE",
	}
}

// GetMappingUpdateTaskScheduleDetailsAuthModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskScheduleDetailsAuthModeEnum(val string) (UpdateTaskScheduleDetailsAuthModeEnum, bool) {
	mappingUpdateTaskScheduleDetailsAuthModeEnumIgnoreCase := make(map[string]UpdateTaskScheduleDetailsAuthModeEnum)
	for k, v := range mappingUpdateTaskScheduleDetailsAuthModeEnum {
		mappingUpdateTaskScheduleDetailsAuthModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTaskScheduleDetailsAuthModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateTaskScheduleDetailsExpectedDurationUnitEnum Enum with underlying type: string
type UpdateTaskScheduleDetailsExpectedDurationUnitEnum string

// Set of constants representing the allowable values for UpdateTaskScheduleDetailsExpectedDurationUnitEnum
const (
	UpdateTaskScheduleDetailsExpectedDurationUnitSeconds UpdateTaskScheduleDetailsExpectedDurationUnitEnum = "SECONDS"
	UpdateTaskScheduleDetailsExpectedDurationUnitMinutes UpdateTaskScheduleDetailsExpectedDurationUnitEnum = "MINUTES"
	UpdateTaskScheduleDetailsExpectedDurationUnitHours   UpdateTaskScheduleDetailsExpectedDurationUnitEnum = "HOURS"
	UpdateTaskScheduleDetailsExpectedDurationUnitDays    UpdateTaskScheduleDetailsExpectedDurationUnitEnum = "DAYS"
)

var mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum = map[string]UpdateTaskScheduleDetailsExpectedDurationUnitEnum{
	"SECONDS": UpdateTaskScheduleDetailsExpectedDurationUnitSeconds,
	"MINUTES": UpdateTaskScheduleDetailsExpectedDurationUnitMinutes,
	"HOURS":   UpdateTaskScheduleDetailsExpectedDurationUnitHours,
	"DAYS":    UpdateTaskScheduleDetailsExpectedDurationUnitDays,
}

// GetUpdateTaskScheduleDetailsExpectedDurationUnitEnumValues Enumerates the set of values for UpdateTaskScheduleDetailsExpectedDurationUnitEnum
func GetUpdateTaskScheduleDetailsExpectedDurationUnitEnumValues() []UpdateTaskScheduleDetailsExpectedDurationUnitEnum {
	values := make([]UpdateTaskScheduleDetailsExpectedDurationUnitEnum, 0)
	for _, v := range mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskScheduleDetailsExpectedDurationUnitEnumStringValues Enumerates the set of values in String for UpdateTaskScheduleDetailsExpectedDurationUnitEnum
func GetUpdateTaskScheduleDetailsExpectedDurationUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum(val string) (UpdateTaskScheduleDetailsExpectedDurationUnitEnum, bool) {
	mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase := make(map[string]UpdateTaskScheduleDetailsExpectedDurationUnitEnum)
	for k, v := range mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnum {
		mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTaskScheduleDetailsExpectedDurationUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
