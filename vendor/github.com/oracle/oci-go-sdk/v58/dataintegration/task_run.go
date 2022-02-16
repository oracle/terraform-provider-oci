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

// TaskRun The information about a task run.
type TaskRun struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ConfigProvider *ConfigProvider `mandatory:"false" json:"configProvider"`

	// The status of the task run.
	Status TaskRunStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The start time.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The end time.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// The date and time the object was last updated.
	LastUpdated *int64 `mandatory:"false" json:"lastUpdated"`

	// The number of records processed in the task run.
	RecordsWritten *int64 `mandatory:"false" json:"recordsWritten"`

	// The number of bytes processed in the task run.
	BytesProcessed *int64 `mandatory:"false" json:"bytesProcessed"`

	// Contains an error message if status is `ERROR`.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The expected duration for the task run.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	// The expected duration unit of measure.
	ExpectedDurationUnit TaskRunExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	// Task Key of the task for which TaskRun is being created. If not specified, the AggregatorKey in RegistryMetadata will be assumed to be the TaskKey
	TaskKey *string `mandatory:"false" json:"taskKey"`

	// The external identifier for the task run.
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Holds the particular attempt number.
	RetryAttempt *int `mandatory:"false" json:"retryAttempt"`

	TaskSchedule *TaskSchedule `mandatory:"false" json:"taskSchedule"`

	// A map of metrics for the run.
	Metrics map[string]float32 `mandatory:"false" json:"metrics"`

	// A map of the outputs of the run.
	Outputs map[string]ParameterValue `mandatory:"false" json:"outputs"`

	// An array of execution errors from the run.
	ExecutionErrors []string `mandatory:"false" json:"executionErrors"`

	// An array of termination errors from the run.
	TerminationErrors []string `mandatory:"false" json:"terminationErrors"`

	// The autorization mode for when the task was executed.
	AuthMode TaskRunAuthModeEnum `mandatory:"false" json:"authMode,omitempty"`

	// The OPC request ID of execution of the task run.
	OpcRequestId *string `mandatory:"false" json:"opcRequestId"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The type of task run.
	TaskType TaskRunTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m TaskRun) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskRun) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTaskRunStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTaskRunStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRunExpectedDurationUnitEnum(string(m.ExpectedDurationUnit)); !ok && m.ExpectedDurationUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExpectedDurationUnit: %s. Supported values are: %s.", m.ExpectedDurationUnit, strings.Join(GetTaskRunExpectedDurationUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRunAuthModeEnum(string(m.AuthMode)); !ok && m.AuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthMode: %s. Supported values are: %s.", m.AuthMode, strings.Join(GetTaskRunAuthModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRunTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetTaskRunTaskTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskRunStatusEnum Enum with underlying type: string
type TaskRunStatusEnum string

// Set of constants representing the allowable values for TaskRunStatusEnum
const (
	TaskRunStatusNotStarted  TaskRunStatusEnum = "NOT_STARTED"
	TaskRunStatusQueued      TaskRunStatusEnum = "QUEUED"
	TaskRunStatusRunning     TaskRunStatusEnum = "RUNNING"
	TaskRunStatusTerminating TaskRunStatusEnum = "TERMINATING"
	TaskRunStatusTerminated  TaskRunStatusEnum = "TERMINATED"
	TaskRunStatusSuccess     TaskRunStatusEnum = "SUCCESS"
	TaskRunStatusError       TaskRunStatusEnum = "ERROR"
)

var mappingTaskRunStatusEnum = map[string]TaskRunStatusEnum{
	"NOT_STARTED": TaskRunStatusNotStarted,
	"QUEUED":      TaskRunStatusQueued,
	"RUNNING":     TaskRunStatusRunning,
	"TERMINATING": TaskRunStatusTerminating,
	"TERMINATED":  TaskRunStatusTerminated,
	"SUCCESS":     TaskRunStatusSuccess,
	"ERROR":       TaskRunStatusError,
}

// GetTaskRunStatusEnumValues Enumerates the set of values for TaskRunStatusEnum
func GetTaskRunStatusEnumValues() []TaskRunStatusEnum {
	values := make([]TaskRunStatusEnum, 0)
	for _, v := range mappingTaskRunStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunStatusEnumStringValues Enumerates the set of values in String for TaskRunStatusEnum
func GetTaskRunStatusEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"QUEUED",
		"RUNNING",
		"TERMINATING",
		"TERMINATED",
		"SUCCESS",
		"ERROR",
	}
}

// GetMappingTaskRunStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunStatusEnum(val string) (TaskRunStatusEnum, bool) {
	mappingTaskRunStatusEnumIgnoreCase := make(map[string]TaskRunStatusEnum)
	for k, v := range mappingTaskRunStatusEnum {
		mappingTaskRunStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskRunStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRunExpectedDurationUnitEnum Enum with underlying type: string
type TaskRunExpectedDurationUnitEnum string

// Set of constants representing the allowable values for TaskRunExpectedDurationUnitEnum
const (
	TaskRunExpectedDurationUnitSeconds TaskRunExpectedDurationUnitEnum = "SECONDS"
	TaskRunExpectedDurationUnitMinutes TaskRunExpectedDurationUnitEnum = "MINUTES"
	TaskRunExpectedDurationUnitHours   TaskRunExpectedDurationUnitEnum = "HOURS"
	TaskRunExpectedDurationUnitDays    TaskRunExpectedDurationUnitEnum = "DAYS"
)

var mappingTaskRunExpectedDurationUnitEnum = map[string]TaskRunExpectedDurationUnitEnum{
	"SECONDS": TaskRunExpectedDurationUnitSeconds,
	"MINUTES": TaskRunExpectedDurationUnitMinutes,
	"HOURS":   TaskRunExpectedDurationUnitHours,
	"DAYS":    TaskRunExpectedDurationUnitDays,
}

// GetTaskRunExpectedDurationUnitEnumValues Enumerates the set of values for TaskRunExpectedDurationUnitEnum
func GetTaskRunExpectedDurationUnitEnumValues() []TaskRunExpectedDurationUnitEnum {
	values := make([]TaskRunExpectedDurationUnitEnum, 0)
	for _, v := range mappingTaskRunExpectedDurationUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunExpectedDurationUnitEnumStringValues Enumerates the set of values in String for TaskRunExpectedDurationUnitEnum
func GetTaskRunExpectedDurationUnitEnumStringValues() []string {
	return []string{
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
}

// GetMappingTaskRunExpectedDurationUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunExpectedDurationUnitEnum(val string) (TaskRunExpectedDurationUnitEnum, bool) {
	mappingTaskRunExpectedDurationUnitEnumIgnoreCase := make(map[string]TaskRunExpectedDurationUnitEnum)
	for k, v := range mappingTaskRunExpectedDurationUnitEnum {
		mappingTaskRunExpectedDurationUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskRunExpectedDurationUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRunAuthModeEnum Enum with underlying type: string
type TaskRunAuthModeEnum string

// Set of constants representing the allowable values for TaskRunAuthModeEnum
const (
	TaskRunAuthModeObo               TaskRunAuthModeEnum = "OBO"
	TaskRunAuthModeResourcePrincipal TaskRunAuthModeEnum = "RESOURCE_PRINCIPAL"
	TaskRunAuthModeUserCertificate   TaskRunAuthModeEnum = "USER_CERTIFICATE"
)

var mappingTaskRunAuthModeEnum = map[string]TaskRunAuthModeEnum{
	"OBO":                TaskRunAuthModeObo,
	"RESOURCE_PRINCIPAL": TaskRunAuthModeResourcePrincipal,
	"USER_CERTIFICATE":   TaskRunAuthModeUserCertificate,
}

// GetTaskRunAuthModeEnumValues Enumerates the set of values for TaskRunAuthModeEnum
func GetTaskRunAuthModeEnumValues() []TaskRunAuthModeEnum {
	values := make([]TaskRunAuthModeEnum, 0)
	for _, v := range mappingTaskRunAuthModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunAuthModeEnumStringValues Enumerates the set of values in String for TaskRunAuthModeEnum
func GetTaskRunAuthModeEnumStringValues() []string {
	return []string{
		"OBO",
		"RESOURCE_PRINCIPAL",
		"USER_CERTIFICATE",
	}
}

// GetMappingTaskRunAuthModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunAuthModeEnum(val string) (TaskRunAuthModeEnum, bool) {
	mappingTaskRunAuthModeEnumIgnoreCase := make(map[string]TaskRunAuthModeEnum)
	for k, v := range mappingTaskRunAuthModeEnum {
		mappingTaskRunAuthModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskRunAuthModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRunTaskTypeEnum Enum with underlying type: string
type TaskRunTaskTypeEnum string

// Set of constants representing the allowable values for TaskRunTaskTypeEnum
const (
	TaskRunTaskTypeIntegrationTask TaskRunTaskTypeEnum = "INTEGRATION_TASK"
	TaskRunTaskTypeDataLoaderTask  TaskRunTaskTypeEnum = "DATA_LOADER_TASK"
	TaskRunTaskTypePipelineTask    TaskRunTaskTypeEnum = "PIPELINE_TASK"
	TaskRunTaskTypeSqlTask         TaskRunTaskTypeEnum = "SQL_TASK"
	TaskRunTaskTypeOciDataflowTask TaskRunTaskTypeEnum = "OCI_DATAFLOW_TASK"
	TaskRunTaskTypeRestTask        TaskRunTaskTypeEnum = "REST_TASK"
)

var mappingTaskRunTaskTypeEnum = map[string]TaskRunTaskTypeEnum{
	"INTEGRATION_TASK":  TaskRunTaskTypeIntegrationTask,
	"DATA_LOADER_TASK":  TaskRunTaskTypeDataLoaderTask,
	"PIPELINE_TASK":     TaskRunTaskTypePipelineTask,
	"SQL_TASK":          TaskRunTaskTypeSqlTask,
	"OCI_DATAFLOW_TASK": TaskRunTaskTypeOciDataflowTask,
	"REST_TASK":         TaskRunTaskTypeRestTask,
}

// GetTaskRunTaskTypeEnumValues Enumerates the set of values for TaskRunTaskTypeEnum
func GetTaskRunTaskTypeEnumValues() []TaskRunTaskTypeEnum {
	values := make([]TaskRunTaskTypeEnum, 0)
	for _, v := range mappingTaskRunTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunTaskTypeEnumStringValues Enumerates the set of values in String for TaskRunTaskTypeEnum
func GetTaskRunTaskTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingTaskRunTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunTaskTypeEnum(val string) (TaskRunTaskTypeEnum, bool) {
	mappingTaskRunTaskTypeEnumIgnoreCase := make(map[string]TaskRunTaskTypeEnum)
	for k, v := range mappingTaskRunTaskTypeEnum {
		mappingTaskRunTaskTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskRunTaskTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
