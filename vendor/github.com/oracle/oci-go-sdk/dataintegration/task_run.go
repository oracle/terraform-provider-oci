// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TaskRun The information about TaskRun.
type TaskRun struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	ConfigProvider *ConfigProvider `mandatory:"false" json:"configProvider"`

	// status
	Status TaskRunStatusEnum `mandatory:"false" json:"status,omitempty"`

	// startTimeMillis
	StartTimeMillis *int64 `mandatory:"false" json:"startTimeMillis"`

	// endTimeMillis
	EndTimeMillis *int64 `mandatory:"false" json:"endTimeMillis"`

	// lastUpdated
	LastUpdated *int64 `mandatory:"false" json:"lastUpdated"`

	// Number of records processed in task run.
	RecordsWritten *int64 `mandatory:"false" json:"recordsWritten"`

	// Number of bytes processed in task run.
	BytesProcessed *int64 `mandatory:"false" json:"bytesProcessed"`

	// Error message if status is ERROR
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Opc request id of execution of task run
	OpcRequestId *string `mandatory:"false" json:"opcRequestId"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The type of the task for the run.
	TaskType TaskRunTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`
}

func (m TaskRun) String() string {
	return common.PointerString(m)
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

var mappingTaskRunStatus = map[string]TaskRunStatusEnum{
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
	for _, v := range mappingTaskRunStatus {
		values = append(values, v)
	}
	return values
}

// TaskRunTaskTypeEnum Enum with underlying type: string
type TaskRunTaskTypeEnum string

// Set of constants representing the allowable values for TaskRunTaskTypeEnum
const (
	TaskRunTaskTypeIntegrationTask TaskRunTaskTypeEnum = "INTEGRATION_TASK"
	TaskRunTaskTypeDataLoaderTask  TaskRunTaskTypeEnum = "DATA_LOADER_TASK"
)

var mappingTaskRunTaskType = map[string]TaskRunTaskTypeEnum{
	"INTEGRATION_TASK": TaskRunTaskTypeIntegrationTask,
	"DATA_LOADER_TASK": TaskRunTaskTypeDataLoaderTask,
}

// GetTaskRunTaskTypeEnumValues Enumerates the set of values for TaskRunTaskTypeEnum
func GetTaskRunTaskTypeEnumValues() []TaskRunTaskTypeEnum {
	values := make([]TaskRunTaskTypeEnum, 0)
	for _, v := range mappingTaskRunTaskType {
		values = append(values, v)
	}
	return values
}
