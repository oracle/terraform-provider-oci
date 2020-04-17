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

// TaskRunSummary The information about TaskRun.
type TaskRunSummary struct {

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
	Status TaskRunSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

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
	TaskType TaskRunSummaryTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskRunSummary) String() string {
	return common.PointerString(m)
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

var mappingTaskRunSummaryStatus = map[string]TaskRunSummaryStatusEnum{
	"NOT_STARTED": TaskRunSummaryStatusNotStarted,
	"QUEUED":      TaskRunSummaryStatusQueued,
	"RUNNING":     TaskRunSummaryStatusRunning,
	"TERMINATING": TaskRunSummaryStatusTerminating,
	"TERMINATED":  TaskRunSummaryStatusTerminated,
	"SUCCESS":     TaskRunSummaryStatusSuccess,
	"ERROR":       TaskRunSummaryStatusError,
}

// GetTaskRunSummaryStatusEnumValues Enumerates the set of values for TaskRunSummaryStatusEnum
func GetTaskRunSummaryStatusEnumValues() []TaskRunSummaryStatusEnum {
	values := make([]TaskRunSummaryStatusEnum, 0)
	for _, v := range mappingTaskRunSummaryStatus {
		values = append(values, v)
	}
	return values
}

// TaskRunSummaryTaskTypeEnum Enum with underlying type: string
type TaskRunSummaryTaskTypeEnum string

// Set of constants representing the allowable values for TaskRunSummaryTaskTypeEnum
const (
	TaskRunSummaryTaskTypeIntegrationTask TaskRunSummaryTaskTypeEnum = "INTEGRATION_TASK"
	TaskRunSummaryTaskTypeDataLoaderTask  TaskRunSummaryTaskTypeEnum = "DATA_LOADER_TASK"
)

var mappingTaskRunSummaryTaskType = map[string]TaskRunSummaryTaskTypeEnum{
	"INTEGRATION_TASK": TaskRunSummaryTaskTypeIntegrationTask,
	"DATA_LOADER_TASK": TaskRunSummaryTaskTypeDataLoaderTask,
}

// GetTaskRunSummaryTaskTypeEnumValues Enumerates the set of values for TaskRunSummaryTaskTypeEnum
func GetTaskRunSummaryTaskTypeEnumValues() []TaskRunSummaryTaskTypeEnum {
	values := make([]TaskRunSummaryTaskTypeEnum, 0)
	for _, v := range mappingTaskRunSummaryTaskType {
		values = append(values, v)
	}
	return values
}
