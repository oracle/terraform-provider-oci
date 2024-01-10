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

// TaskRunLineageDetails The task lineage object provides information on the lineage information of a task after execution.
type TaskRunLineageDetails struct {

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

	// Task name
	TaskName *string `mandatory:"false" json:"taskName"`

	// Task name
	TaskType *string `mandatory:"false" json:"taskType"`

	// The object key.
	TaskKey *string `mandatory:"false" json:"taskKey"`

	// This value is used to track if lineage generation for a task is completed or not.
	IsLineageGenCompleted *bool `mandatory:"false" json:"isLineageGenCompleted"`

	// The status of the task run.
	TaskExecutionStatus TaskRunLineageDetailsTaskExecutionStatusEnum `mandatory:"false" json:"taskExecutionStatus,omitempty"`

	Flow *DataFlow `mandatory:"false" json:"flow"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskRunLineageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskRunLineageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTaskRunLineageDetailsTaskExecutionStatusEnum(string(m.TaskExecutionStatus)); !ok && m.TaskExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskExecutionStatus: %s. Supported values are: %s.", m.TaskExecutionStatus, strings.Join(GetTaskRunLineageDetailsTaskExecutionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskRunLineageDetailsTaskExecutionStatusEnum Enum with underlying type: string
type TaskRunLineageDetailsTaskExecutionStatusEnum string

// Set of constants representing the allowable values for TaskRunLineageDetailsTaskExecutionStatusEnum
const (
	TaskRunLineageDetailsTaskExecutionStatusSuccess    TaskRunLineageDetailsTaskExecutionStatusEnum = "SUCCESS"
	TaskRunLineageDetailsTaskExecutionStatusError      TaskRunLineageDetailsTaskExecutionStatusEnum = "ERROR"
	TaskRunLineageDetailsTaskExecutionStatusTerminated TaskRunLineageDetailsTaskExecutionStatusEnum = "TERMINATED"
)

var mappingTaskRunLineageDetailsTaskExecutionStatusEnum = map[string]TaskRunLineageDetailsTaskExecutionStatusEnum{
	"SUCCESS":    TaskRunLineageDetailsTaskExecutionStatusSuccess,
	"ERROR":      TaskRunLineageDetailsTaskExecutionStatusError,
	"TERMINATED": TaskRunLineageDetailsTaskExecutionStatusTerminated,
}

var mappingTaskRunLineageDetailsTaskExecutionStatusEnumLowerCase = map[string]TaskRunLineageDetailsTaskExecutionStatusEnum{
	"success":    TaskRunLineageDetailsTaskExecutionStatusSuccess,
	"error":      TaskRunLineageDetailsTaskExecutionStatusError,
	"terminated": TaskRunLineageDetailsTaskExecutionStatusTerminated,
}

// GetTaskRunLineageDetailsTaskExecutionStatusEnumValues Enumerates the set of values for TaskRunLineageDetailsTaskExecutionStatusEnum
func GetTaskRunLineageDetailsTaskExecutionStatusEnumValues() []TaskRunLineageDetailsTaskExecutionStatusEnum {
	values := make([]TaskRunLineageDetailsTaskExecutionStatusEnum, 0)
	for _, v := range mappingTaskRunLineageDetailsTaskExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunLineageDetailsTaskExecutionStatusEnumStringValues Enumerates the set of values in String for TaskRunLineageDetailsTaskExecutionStatusEnum
func GetTaskRunLineageDetailsTaskExecutionStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"ERROR",
		"TERMINATED",
	}
}

// GetMappingTaskRunLineageDetailsTaskExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunLineageDetailsTaskExecutionStatusEnum(val string) (TaskRunLineageDetailsTaskExecutionStatusEnum, bool) {
	enum, ok := mappingTaskRunLineageDetailsTaskExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
