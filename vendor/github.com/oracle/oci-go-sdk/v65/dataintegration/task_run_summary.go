// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaskRunSummary The information about a task run.
type TaskRunSummary struct {

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object type.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// status
	Status TaskRunSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The task run start time.
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// The task run end time.
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// The date and time the task run was last updated.
	LastUpdated *int64 `mandatory:"false" json:"lastUpdated"`

	// Number of records processed in task run.
	RecordsWritten *int64 `mandatory:"false" json:"recordsWritten"`

	// Number of bytes processed in task run.
	BytesProcessed *int64 `mandatory:"false" json:"bytesProcessed"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The type of the task for the run.
	TaskType TaskRunSummaryTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// Reference Task Run Id to be used for re-run
	RefTaskRunId *string `mandatory:"false" json:"refTaskRunId"`

	// Supported re-run types
	ReRunType TaskRunSummaryReRunTypeEnum `mandatory:"false" json:"reRunType,omitempty"`

	// Step Id for running from a certain step.
	StepId *string `mandatory:"false" json:"stepId"`

	// A map of the configuration provider input bindings of the run.
	Inputs map[string]ParameterValue `mandatory:"false" json:"inputs"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskRunSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskRunSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTaskRunSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTaskRunSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRunSummaryTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetTaskRunSummaryTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRunSummaryReRunTypeEnum(string(m.ReRunType)); !ok && m.ReRunType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReRunType: %s. Supported values are: %s.", m.ReRunType, strings.Join(GetTaskRunSummaryReRunTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskRunSummaryStatusEnum Enum with underlying type: string
type TaskRunSummaryStatusEnum string

// Set of constants representing the allowable values for TaskRunSummaryStatusEnum
const (
	TaskRunSummaryStatusNotStarted  TaskRunSummaryStatusEnum = "NOT_STARTED"
	TaskRunSummaryStatusQueued      TaskRunSummaryStatusEnum = "QUEUED"
	TaskRunSummaryStatusRunning     TaskRunSummaryStatusEnum = "RUNNING"
	TaskRunSummaryStatusTerminating TaskRunSummaryStatusEnum = "TERMINATING"
	TaskRunSummaryStatusTerminated  TaskRunSummaryStatusEnum = "TERMINATED"
	TaskRunSummaryStatusSuccess     TaskRunSummaryStatusEnum = "SUCCESS"
	TaskRunSummaryStatusError       TaskRunSummaryStatusEnum = "ERROR"
)

var mappingTaskRunSummaryStatusEnum = map[string]TaskRunSummaryStatusEnum{
	"NOT_STARTED": TaskRunSummaryStatusNotStarted,
	"QUEUED":      TaskRunSummaryStatusQueued,
	"RUNNING":     TaskRunSummaryStatusRunning,
	"TERMINATING": TaskRunSummaryStatusTerminating,
	"TERMINATED":  TaskRunSummaryStatusTerminated,
	"SUCCESS":     TaskRunSummaryStatusSuccess,
	"ERROR":       TaskRunSummaryStatusError,
}

var mappingTaskRunSummaryStatusEnumLowerCase = map[string]TaskRunSummaryStatusEnum{
	"not_started": TaskRunSummaryStatusNotStarted,
	"queued":      TaskRunSummaryStatusQueued,
	"running":     TaskRunSummaryStatusRunning,
	"terminating": TaskRunSummaryStatusTerminating,
	"terminated":  TaskRunSummaryStatusTerminated,
	"success":     TaskRunSummaryStatusSuccess,
	"error":       TaskRunSummaryStatusError,
}

// GetTaskRunSummaryStatusEnumValues Enumerates the set of values for TaskRunSummaryStatusEnum
func GetTaskRunSummaryStatusEnumValues() []TaskRunSummaryStatusEnum {
	values := make([]TaskRunSummaryStatusEnum, 0)
	for _, v := range mappingTaskRunSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunSummaryStatusEnumStringValues Enumerates the set of values in String for TaskRunSummaryStatusEnum
func GetTaskRunSummaryStatusEnumStringValues() []string {
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

// GetMappingTaskRunSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunSummaryStatusEnum(val string) (TaskRunSummaryStatusEnum, bool) {
	enum, ok := mappingTaskRunSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRunSummaryTaskTypeEnum Enum with underlying type: string
type TaskRunSummaryTaskTypeEnum string

// Set of constants representing the allowable values for TaskRunSummaryTaskTypeEnum
const (
	TaskRunSummaryTaskTypeIntegrationTask TaskRunSummaryTaskTypeEnum = "INTEGRATION_TASK"
	TaskRunSummaryTaskTypeDataLoaderTask  TaskRunSummaryTaskTypeEnum = "DATA_LOADER_TASK"
	TaskRunSummaryTaskTypePipelineTask    TaskRunSummaryTaskTypeEnum = "PIPELINE_TASK"
	TaskRunSummaryTaskTypeSqlTask         TaskRunSummaryTaskTypeEnum = "SQL_TASK"
	TaskRunSummaryTaskTypeOciDataflowTask TaskRunSummaryTaskTypeEnum = "OCI_DATAFLOW_TASK"
	TaskRunSummaryTaskTypeRestTask        TaskRunSummaryTaskTypeEnum = "REST_TASK"
)

var mappingTaskRunSummaryTaskTypeEnum = map[string]TaskRunSummaryTaskTypeEnum{
	"INTEGRATION_TASK":  TaskRunSummaryTaskTypeIntegrationTask,
	"DATA_LOADER_TASK":  TaskRunSummaryTaskTypeDataLoaderTask,
	"PIPELINE_TASK":     TaskRunSummaryTaskTypePipelineTask,
	"SQL_TASK":          TaskRunSummaryTaskTypeSqlTask,
	"OCI_DATAFLOW_TASK": TaskRunSummaryTaskTypeOciDataflowTask,
	"REST_TASK":         TaskRunSummaryTaskTypeRestTask,
}

var mappingTaskRunSummaryTaskTypeEnumLowerCase = map[string]TaskRunSummaryTaskTypeEnum{
	"integration_task":  TaskRunSummaryTaskTypeIntegrationTask,
	"data_loader_task":  TaskRunSummaryTaskTypeDataLoaderTask,
	"pipeline_task":     TaskRunSummaryTaskTypePipelineTask,
	"sql_task":          TaskRunSummaryTaskTypeSqlTask,
	"oci_dataflow_task": TaskRunSummaryTaskTypeOciDataflowTask,
	"rest_task":         TaskRunSummaryTaskTypeRestTask,
}

// GetTaskRunSummaryTaskTypeEnumValues Enumerates the set of values for TaskRunSummaryTaskTypeEnum
func GetTaskRunSummaryTaskTypeEnumValues() []TaskRunSummaryTaskTypeEnum {
	values := make([]TaskRunSummaryTaskTypeEnum, 0)
	for _, v := range mappingTaskRunSummaryTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunSummaryTaskTypeEnumStringValues Enumerates the set of values in String for TaskRunSummaryTaskTypeEnum
func GetTaskRunSummaryTaskTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingTaskRunSummaryTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunSummaryTaskTypeEnum(val string) (TaskRunSummaryTaskTypeEnum, bool) {
	enum, ok := mappingTaskRunSummaryTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRunSummaryReRunTypeEnum Enum with underlying type: string
type TaskRunSummaryReRunTypeEnum string

// Set of constants representing the allowable values for TaskRunSummaryReRunTypeEnum
const (
	TaskRunSummaryReRunTypeBeginning TaskRunSummaryReRunTypeEnum = "BEGINNING"
	TaskRunSummaryReRunTypeFailed    TaskRunSummaryReRunTypeEnum = "FAILED"
	TaskRunSummaryReRunTypeStep      TaskRunSummaryReRunTypeEnum = "STEP"
)

var mappingTaskRunSummaryReRunTypeEnum = map[string]TaskRunSummaryReRunTypeEnum{
	"BEGINNING": TaskRunSummaryReRunTypeBeginning,
	"FAILED":    TaskRunSummaryReRunTypeFailed,
	"STEP":      TaskRunSummaryReRunTypeStep,
}

var mappingTaskRunSummaryReRunTypeEnumLowerCase = map[string]TaskRunSummaryReRunTypeEnum{
	"beginning": TaskRunSummaryReRunTypeBeginning,
	"failed":    TaskRunSummaryReRunTypeFailed,
	"step":      TaskRunSummaryReRunTypeStep,
}

// GetTaskRunSummaryReRunTypeEnumValues Enumerates the set of values for TaskRunSummaryReRunTypeEnum
func GetTaskRunSummaryReRunTypeEnumValues() []TaskRunSummaryReRunTypeEnum {
	values := make([]TaskRunSummaryReRunTypeEnum, 0)
	for _, v := range mappingTaskRunSummaryReRunTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunSummaryReRunTypeEnumStringValues Enumerates the set of values in String for TaskRunSummaryReRunTypeEnum
func GetTaskRunSummaryReRunTypeEnumStringValues() []string {
	return []string{
		"BEGINNING",
		"FAILED",
		"STEP",
	}
}

// GetMappingTaskRunSummaryReRunTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunSummaryReRunTypeEnum(val string) (TaskRunSummaryReRunTypeEnum, bool) {
	enum, ok := mappingTaskRunSummaryReRunTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
