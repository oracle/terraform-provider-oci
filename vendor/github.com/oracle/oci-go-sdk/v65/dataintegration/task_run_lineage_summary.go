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

// TaskRunLineageSummary The information about TaskRunLineage.
type TaskRunLineageSummary struct {

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
	TaskExecutionStatus TaskRunLineageSummaryTaskExecutionStatusEnum `mandatory:"false" json:"taskExecutionStatus,omitempty"`

	Flow *DataFlow `mandatory:"false" json:"flow"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`
}

func (m TaskRunLineageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskRunLineageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTaskRunLineageSummaryTaskExecutionStatusEnum(string(m.TaskExecutionStatus)); !ok && m.TaskExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskExecutionStatus: %s. Supported values are: %s.", m.TaskExecutionStatus, strings.Join(GetTaskRunLineageSummaryTaskExecutionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskRunLineageSummaryTaskExecutionStatusEnum Enum with underlying type: string
type TaskRunLineageSummaryTaskExecutionStatusEnum string

// Set of constants representing the allowable values for TaskRunLineageSummaryTaskExecutionStatusEnum
const (
	TaskRunLineageSummaryTaskExecutionStatusSuccess    TaskRunLineageSummaryTaskExecutionStatusEnum = "SUCCESS"
	TaskRunLineageSummaryTaskExecutionStatusError      TaskRunLineageSummaryTaskExecutionStatusEnum = "ERROR"
	TaskRunLineageSummaryTaskExecutionStatusTerminated TaskRunLineageSummaryTaskExecutionStatusEnum = "TERMINATED"
)

var mappingTaskRunLineageSummaryTaskExecutionStatusEnum = map[string]TaskRunLineageSummaryTaskExecutionStatusEnum{
	"SUCCESS":    TaskRunLineageSummaryTaskExecutionStatusSuccess,
	"ERROR":      TaskRunLineageSummaryTaskExecutionStatusError,
	"TERMINATED": TaskRunLineageSummaryTaskExecutionStatusTerminated,
}

var mappingTaskRunLineageSummaryTaskExecutionStatusEnumLowerCase = map[string]TaskRunLineageSummaryTaskExecutionStatusEnum{
	"success":    TaskRunLineageSummaryTaskExecutionStatusSuccess,
	"error":      TaskRunLineageSummaryTaskExecutionStatusError,
	"terminated": TaskRunLineageSummaryTaskExecutionStatusTerminated,
}

// GetTaskRunLineageSummaryTaskExecutionStatusEnumValues Enumerates the set of values for TaskRunLineageSummaryTaskExecutionStatusEnum
func GetTaskRunLineageSummaryTaskExecutionStatusEnumValues() []TaskRunLineageSummaryTaskExecutionStatusEnum {
	values := make([]TaskRunLineageSummaryTaskExecutionStatusEnum, 0)
	for _, v := range mappingTaskRunLineageSummaryTaskExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRunLineageSummaryTaskExecutionStatusEnumStringValues Enumerates the set of values in String for TaskRunLineageSummaryTaskExecutionStatusEnum
func GetTaskRunLineageSummaryTaskExecutionStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"ERROR",
		"TERMINATED",
	}
}

// GetMappingTaskRunLineageSummaryTaskExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRunLineageSummaryTaskExecutionStatusEnum(val string) (TaskRunLineageSummaryTaskExecutionStatusEnum, bool) {
	enum, ok := mappingTaskRunLineageSummaryTaskExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
