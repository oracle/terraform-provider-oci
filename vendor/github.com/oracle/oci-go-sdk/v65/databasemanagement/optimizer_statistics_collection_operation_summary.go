// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OptimizerStatisticsCollectionOperationSummary The summary of the Optimizer Statistics Collection operation.
type OptimizerStatisticsCollectionOperationSummary struct {

	// The ID of the operation.
	Id *int `mandatory:"true" json:"id"`

	// The name of the operation.
	OperationName *string `mandatory:"true" json:"operationName"`

	// The target object type such as Table, Index, and Partition.
	Target *string `mandatory:"true" json:"target"`

	// The name of the job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The status of the operation such as Completed, and Failed.
	Status OptimizerStatisticsCollectionOperationSummaryStatusEnum `mandatory:"true" json:"status"`

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
}

func (m OptimizerStatisticsCollectionOperationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerStatisticsCollectionOperationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOptimizerStatisticsCollectionOperationSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOptimizerStatisticsCollectionOperationSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OptimizerStatisticsCollectionOperationSummaryStatusEnum Enum with underlying type: string
type OptimizerStatisticsCollectionOperationSummaryStatusEnum string

// Set of constants representing the allowable values for OptimizerStatisticsCollectionOperationSummaryStatusEnum
const (
	OptimizerStatisticsCollectionOperationSummaryStatusInProgress OptimizerStatisticsCollectionOperationSummaryStatusEnum = "IN_PROGRESS"
	OptimizerStatisticsCollectionOperationSummaryStatusCompleted  OptimizerStatisticsCollectionOperationSummaryStatusEnum = "COMPLETED"
	OptimizerStatisticsCollectionOperationSummaryStatusFailed     OptimizerStatisticsCollectionOperationSummaryStatusEnum = "FAILED"
	OptimizerStatisticsCollectionOperationSummaryStatusTimedOut   OptimizerStatisticsCollectionOperationSummaryStatusEnum = "TIMED_OUT"
)

var mappingOptimizerStatisticsCollectionOperationSummaryStatusEnum = map[string]OptimizerStatisticsCollectionOperationSummaryStatusEnum{
	"IN_PROGRESS": OptimizerStatisticsCollectionOperationSummaryStatusInProgress,
	"COMPLETED":   OptimizerStatisticsCollectionOperationSummaryStatusCompleted,
	"FAILED":      OptimizerStatisticsCollectionOperationSummaryStatusFailed,
	"TIMED_OUT":   OptimizerStatisticsCollectionOperationSummaryStatusTimedOut,
}

var mappingOptimizerStatisticsCollectionOperationSummaryStatusEnumLowerCase = map[string]OptimizerStatisticsCollectionOperationSummaryStatusEnum{
	"in_progress": OptimizerStatisticsCollectionOperationSummaryStatusInProgress,
	"completed":   OptimizerStatisticsCollectionOperationSummaryStatusCompleted,
	"failed":      OptimizerStatisticsCollectionOperationSummaryStatusFailed,
	"timed_out":   OptimizerStatisticsCollectionOperationSummaryStatusTimedOut,
}

// GetOptimizerStatisticsCollectionOperationSummaryStatusEnumValues Enumerates the set of values for OptimizerStatisticsCollectionOperationSummaryStatusEnum
func GetOptimizerStatisticsCollectionOperationSummaryStatusEnumValues() []OptimizerStatisticsCollectionOperationSummaryStatusEnum {
	values := make([]OptimizerStatisticsCollectionOperationSummaryStatusEnum, 0)
	for _, v := range mappingOptimizerStatisticsCollectionOperationSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsCollectionOperationSummaryStatusEnumStringValues Enumerates the set of values in String for OptimizerStatisticsCollectionOperationSummaryStatusEnum
func GetOptimizerStatisticsCollectionOperationSummaryStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"TIMED_OUT",
	}
}

// GetMappingOptimizerStatisticsCollectionOperationSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsCollectionOperationSummaryStatusEnum(val string) (OptimizerStatisticsCollectionOperationSummaryStatusEnum, bool) {
	enum, ok := mappingOptimizerStatisticsCollectionOperationSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
