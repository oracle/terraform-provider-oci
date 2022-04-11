// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// OptimizerStatisticsCollectionOperation The response object includes the managed database, execution details and optimizer statistics collection tasks summary.
type OptimizerStatisticsCollectionOperation struct {

	// Operation id.
	Id *int `mandatory:"true" json:"id"`

	// Name of the operation.
	OperationName *string `mandatory:"true" json:"operationName"`

	// Target object type like Table/Index/Partition etc.
	Target *string `mandatory:"true" json:"target"`

	// Name of the Job.
	JobName *string `mandatory:"true" json:"jobName"`

	// Status of the operation like Completed/Failed etc.
	Status OptimizerStatisticsCollectionOperationStatusEnum `mandatory:"true" json:"status"`

	// Start time of the execution.
	StartTime *string `mandatory:"true" json:"startTime"`

	// End time of the execution.
	EndTime *string `mandatory:"true" json:"endTime"`

	// This is the time it takes to complete the task in seconds.
	DurationInSeconds *float32 `mandatory:"true" json:"durationInSeconds"`

	// Count of objects for which statistics collection is successfully.
	CompletedCount *int `mandatory:"false" json:"completedCount"`

	// Count of objects for which statistics gathering is still in progress.
	InProgressCount *int `mandatory:"false" json:"inProgressCount"`

	// Count of objects for which statistics collection failed.
	FailedCount *int `mandatory:"false" json:"failedCount"`

	// Count of objects statistics for which statistics collection timed out.
	TimedOutCount *int `mandatory:"false" json:"timedOutCount"`

	// Total number of objects statistics collected. This count includes completed, inProgress, failed and timedOut objects counts.
	TotalObjectsCount *int `mandatory:"false" json:"totalObjectsCount"`

	Database *OptimizerDatabase `mandatory:"false" json:"database"`

	// List of gather statistics task summary.
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
