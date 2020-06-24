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

// TaskRunDetails The task run object provides information on the execution of a task.
type TaskRunDetails struct {

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

	// status
	Status TaskRunDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

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

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The type of the task for the run.
	TaskType TaskRunDetailsTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskRunDetails) String() string {
	return common.PointerString(m)
}

// TaskRunDetailsStatusEnum Enum with underlying type: string
type TaskRunDetailsStatusEnum string

// Set of constants representing the allowable values for TaskRunDetailsStatusEnum
const (
	TaskRunDetailsStatusNotStarted  TaskRunDetailsStatusEnum = "NOT_STARTED"
	TaskRunDetailsStatusQueued      TaskRunDetailsStatusEnum = "QUEUED"
	TaskRunDetailsStatusRunning     TaskRunDetailsStatusEnum = "RUNNING"
	TaskRunDetailsStatusTerminating TaskRunDetailsStatusEnum = "TERMINATING"
	TaskRunDetailsStatusTerminated  TaskRunDetailsStatusEnum = "TERMINATED"
	TaskRunDetailsStatusSuccess     TaskRunDetailsStatusEnum = "SUCCESS"
	TaskRunDetailsStatusError       TaskRunDetailsStatusEnum = "ERROR"
)

var mappingTaskRunDetailsStatus = map[string]TaskRunDetailsStatusEnum{
	"NOT_STARTED": TaskRunDetailsStatusNotStarted,
	"QUEUED":      TaskRunDetailsStatusQueued,
	"RUNNING":     TaskRunDetailsStatusRunning,
	"TERMINATING": TaskRunDetailsStatusTerminating,
	"TERMINATED":  TaskRunDetailsStatusTerminated,
	"SUCCESS":     TaskRunDetailsStatusSuccess,
	"ERROR":       TaskRunDetailsStatusError,
}

// GetTaskRunDetailsStatusEnumValues Enumerates the set of values for TaskRunDetailsStatusEnum
func GetTaskRunDetailsStatusEnumValues() []TaskRunDetailsStatusEnum {
	values := make([]TaskRunDetailsStatusEnum, 0)
	for _, v := range mappingTaskRunDetailsStatus {
		values = append(values, v)
	}
	return values
}

// TaskRunDetailsTaskTypeEnum Enum with underlying type: string
type TaskRunDetailsTaskTypeEnum string

// Set of constants representing the allowable values for TaskRunDetailsTaskTypeEnum
const (
	TaskRunDetailsTaskTypeIntegrationTask TaskRunDetailsTaskTypeEnum = "INTEGRATION_TASK"
	TaskRunDetailsTaskTypeDataLoaderTask  TaskRunDetailsTaskTypeEnum = "DATA_LOADER_TASK"
)

var mappingTaskRunDetailsTaskType = map[string]TaskRunDetailsTaskTypeEnum{
	"INTEGRATION_TASK": TaskRunDetailsTaskTypeIntegrationTask,
	"DATA_LOADER_TASK": TaskRunDetailsTaskTypeDataLoaderTask,
}

// GetTaskRunDetailsTaskTypeEnumValues Enumerates the set of values for TaskRunDetailsTaskTypeEnum
func GetTaskRunDetailsTaskTypeEnumValues() []TaskRunDetailsTaskTypeEnum {
	values := make([]TaskRunDetailsTaskTypeEnum, 0)
	for _, v := range mappingTaskRunDetailsTaskType {
		values = append(values, v)
	}
	return values
}
