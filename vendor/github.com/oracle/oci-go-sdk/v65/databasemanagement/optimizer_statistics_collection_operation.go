// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsCollectionOperation The summary of the Optimizer Statistics Collection tasks, which includes details of the Managed Database and the execution.
type OptimizerStatisticsCollectionOperation struct {

	// The ID of the operation.
	Id *int `mandatory:"true" json:"id"`

	// The name of the operation.
	OperationName *string `mandatory:"true" json:"operationName"`

	// The target object type such as Table, Index, and Partition.
	Target *string `mandatory:"true" json:"target"`

	// The name of the job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The status of the operation such as Completed, and Failed.
	Status OptimizerStatisticsCollectionOperationStatusEnum `mandatory:"true" json:"status"`

	// The start time of the operation.
	StartTime *string `mandatory:"true" json:"startTime"`

	// The end time of the operation.
	EndTime *string `mandatory:"true" json:"endTime"`

	// The time it takes to complete the operation (in seconds).
	DurationInSeconds *float32 `mandatory:"true" json:"durationInSeconds"`

	// The number of objects for which statistics collection is completed.
	CompletedCount *int `mandatory:"false" json:"completedCount"`

	// The number of objects for which statistics collection is in progress.
	InProgressCount *int `mandatory:"false" json:"inProgressCount"`

	// The number of objects for which statistics collection failed.
	FailedCount *int `mandatory:"false" json:"failedCount"`

	// The number of objects for which statistics collection timed out.
	TimedOutCount *int `mandatory:"false" json:"timedOutCount"`

	// The total number of objects for which statistics is collected. This number is the sum of all the objects
	// with various statuses: completed, inProgress, failed, and timedOut.
	TotalObjectsCount *int `mandatory:"false" json:"totalObjectsCount"`

	Database *OptimizerDatabase `mandatory:"false" json:"database"`

	// An array of Optimizer Statistics Collection task details.
	Tasks []OptimizerStatisticsOperationTask `mandatory:"false" json:"tasks"`
}

func (m OptimizerStatisticsCollectionOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsCollectionOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOptimizerStatisticsCollectionOperationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOptimizerStatisticsCollectionOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OptimizerStatisticsCollectionOperationStatusEnum Enum with underlying type: string
type OptimizerStatisticsCollectionOperationStatusEnum string

// Set of constants representing the allowable values for OptimizerStatisticsCollectionOperationStatusEnum
const (
	OptimizerStatisticsCollectionOperationStatusInProgress OptimizerStatisticsCollectionOperationStatusEnum = "IN_PROGRESS"
	OptimizerStatisticsCollectionOperationStatusCompleted  OptimizerStatisticsCollectionOperationStatusEnum = "COMPLETED"
	OptimizerStatisticsCollectionOperationStatusFailed     OptimizerStatisticsCollectionOperationStatusEnum = "FAILED"
	OptimizerStatisticsCollectionOperationStatusTimedOut   OptimizerStatisticsCollectionOperationStatusEnum = "TIMED_OUT"
)

var mappingOptimizerStatisticsCollectionOperationStatusEnum = map[string]OptimizerStatisticsCollectionOperationStatusEnum{
	"IN_PROGRESS": OptimizerStatisticsCollectionOperationStatusInProgress,
	"COMPLETED":   OptimizerStatisticsCollectionOperationStatusCompleted,
	"FAILED":      OptimizerStatisticsCollectionOperationStatusFailed,
	"TIMED_OUT":   OptimizerStatisticsCollectionOperationStatusTimedOut,
}

var mappingOptimizerStatisticsCollectionOperationStatusEnumLowerCase = map[string]OptimizerStatisticsCollectionOperationStatusEnum{
	"in_progress": OptimizerStatisticsCollectionOperationStatusInProgress,
	"completed":   OptimizerStatisticsCollectionOperationStatusCompleted,
	"failed":      OptimizerStatisticsCollectionOperationStatusFailed,
	"timed_out":   OptimizerStatisticsCollectionOperationStatusTimedOut,
}

// GetOptimizerStatisticsCollectionOperationStatusEnumValues Enumerates the set of values for OptimizerStatisticsCollectionOperationStatusEnum
func GetOptimizerStatisticsCollectionOperationStatusEnumValues() []OptimizerStatisticsCollectionOperationStatusEnum {
	values := make([]OptimizerStatisticsCollectionOperationStatusEnum, 0)
	for _, v := range mappingOptimizerStatisticsCollectionOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsCollectionOperationStatusEnumStringValues Enumerates the set of values in String for OptimizerStatisticsCollectionOperationStatusEnum
func GetOptimizerStatisticsCollectionOperationStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"TIMED_OUT",
	}
}

// GetMappingOptimizerStatisticsCollectionOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsCollectionOperationStatusEnum(val string) (OptimizerStatisticsCollectionOperationStatusEnum, bool) {
	enum, ok := mappingOptimizerStatisticsCollectionOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
