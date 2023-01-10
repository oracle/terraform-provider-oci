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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuntimeOperatorSummary The information about RuntimeOperator.
type RuntimeOperatorSummary struct {

	// The TaskRun key.
	TaskRunKey *string `mandatory:"false" json:"taskRunKey"`

	// The runtime operator start time.
	StartTimeInMillis *int64 `mandatory:"false" json:"startTimeInMillis"`

	// The runtime operator end time.
	EndTimeInMillis *int64 `mandatory:"false" json:"endTimeInMillis"`

	// status
	Status RuntimeOperatorSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// status
	ExecutionState RuntimeOperatorSummaryExecutionStateEnum `mandatory:"false" json:"executionState,omitempty"`

	// A list of parameters for the pipeline, this allows certain aspects of the pipeline to be configured when the pipeline is executed.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	// The configuration provider bindings.
	Inputs map[string]ParameterValue `mandatory:"false" json:"inputs"`

	// The configuration provider bindings.
	Outputs map[string]ParameterValue `mandatory:"false" json:"outputs"`

	// A map metrics for the task run.
	Metrics map[string]float32 `mandatory:"false" json:"metrics"`
}

func (m RuntimeOperatorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuntimeOperatorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeOperatorSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRuntimeOperatorSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorSummaryExecutionStateEnum(string(m.ExecutionState)); !ok && m.ExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionState: %s. Supported values are: %s.", m.ExecutionState, strings.Join(GetRuntimeOperatorSummaryExecutionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuntimeOperatorSummaryStatusEnum Enum with underlying type: string
type RuntimeOperatorSummaryStatusEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryStatusEnum
const (
	RuntimeOperatorSummaryStatusNotStarted  RuntimeOperatorSummaryStatusEnum = "NOT_STARTED"
	RuntimeOperatorSummaryStatusQueued      RuntimeOperatorSummaryStatusEnum = "QUEUED"
	RuntimeOperatorSummaryStatusRunning     RuntimeOperatorSummaryStatusEnum = "RUNNING"
	RuntimeOperatorSummaryStatusTerminating RuntimeOperatorSummaryStatusEnum = "TERMINATING"
	RuntimeOperatorSummaryStatusTerminated  RuntimeOperatorSummaryStatusEnum = "TERMINATED"
	RuntimeOperatorSummaryStatusSuccess     RuntimeOperatorSummaryStatusEnum = "SUCCESS"
	RuntimeOperatorSummaryStatusError       RuntimeOperatorSummaryStatusEnum = "ERROR"
)

var mappingRuntimeOperatorSummaryStatusEnum = map[string]RuntimeOperatorSummaryStatusEnum{
	"NOT_STARTED": RuntimeOperatorSummaryStatusNotStarted,
	"QUEUED":      RuntimeOperatorSummaryStatusQueued,
	"RUNNING":     RuntimeOperatorSummaryStatusRunning,
	"TERMINATING": RuntimeOperatorSummaryStatusTerminating,
	"TERMINATED":  RuntimeOperatorSummaryStatusTerminated,
	"SUCCESS":     RuntimeOperatorSummaryStatusSuccess,
	"ERROR":       RuntimeOperatorSummaryStatusError,
}

var mappingRuntimeOperatorSummaryStatusEnumLowerCase = map[string]RuntimeOperatorSummaryStatusEnum{
	"not_started": RuntimeOperatorSummaryStatusNotStarted,
	"queued":      RuntimeOperatorSummaryStatusQueued,
	"running":     RuntimeOperatorSummaryStatusRunning,
	"terminating": RuntimeOperatorSummaryStatusTerminating,
	"terminated":  RuntimeOperatorSummaryStatusTerminated,
	"success":     RuntimeOperatorSummaryStatusSuccess,
	"error":       RuntimeOperatorSummaryStatusError,
}

// GetRuntimeOperatorSummaryStatusEnumValues Enumerates the set of values for RuntimeOperatorSummaryStatusEnum
func GetRuntimeOperatorSummaryStatusEnumValues() []RuntimeOperatorSummaryStatusEnum {
	values := make([]RuntimeOperatorSummaryStatusEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryStatusEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryStatusEnum
func GetRuntimeOperatorSummaryStatusEnumStringValues() []string {
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

// GetMappingRuntimeOperatorSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryStatusEnum(val string) (RuntimeOperatorSummaryStatusEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorSummaryExecutionStateEnum Enum with underlying type: string
type RuntimeOperatorSummaryExecutionStateEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryExecutionStateEnum
const (
	RuntimeOperatorSummaryExecutionStateNotStarted RuntimeOperatorSummaryExecutionStateEnum = "NOT_STARTED"
	RuntimeOperatorSummaryExecutionStateRunning    RuntimeOperatorSummaryExecutionStateEnum = "RUNNING"
	RuntimeOperatorSummaryExecutionStateTerminated RuntimeOperatorSummaryExecutionStateEnum = "TERMINATED"
	RuntimeOperatorSummaryExecutionStateSuccess    RuntimeOperatorSummaryExecutionStateEnum = "SUCCESS"
	RuntimeOperatorSummaryExecutionStateError      RuntimeOperatorSummaryExecutionStateEnum = "ERROR"
	RuntimeOperatorSummaryExecutionStateSkipped    RuntimeOperatorSummaryExecutionStateEnum = "SKIPPED"
	RuntimeOperatorSummaryExecutionStateUnknown    RuntimeOperatorSummaryExecutionStateEnum = "UNKNOWN"
)

var mappingRuntimeOperatorSummaryExecutionStateEnum = map[string]RuntimeOperatorSummaryExecutionStateEnum{
	"NOT_STARTED": RuntimeOperatorSummaryExecutionStateNotStarted,
	"RUNNING":     RuntimeOperatorSummaryExecutionStateRunning,
	"TERMINATED":  RuntimeOperatorSummaryExecutionStateTerminated,
	"SUCCESS":     RuntimeOperatorSummaryExecutionStateSuccess,
	"ERROR":       RuntimeOperatorSummaryExecutionStateError,
	"SKIPPED":     RuntimeOperatorSummaryExecutionStateSkipped,
	"UNKNOWN":     RuntimeOperatorSummaryExecutionStateUnknown,
}

var mappingRuntimeOperatorSummaryExecutionStateEnumLowerCase = map[string]RuntimeOperatorSummaryExecutionStateEnum{
	"not_started": RuntimeOperatorSummaryExecutionStateNotStarted,
	"running":     RuntimeOperatorSummaryExecutionStateRunning,
	"terminated":  RuntimeOperatorSummaryExecutionStateTerminated,
	"success":     RuntimeOperatorSummaryExecutionStateSuccess,
	"error":       RuntimeOperatorSummaryExecutionStateError,
	"skipped":     RuntimeOperatorSummaryExecutionStateSkipped,
	"unknown":     RuntimeOperatorSummaryExecutionStateUnknown,
}

// GetRuntimeOperatorSummaryExecutionStateEnumValues Enumerates the set of values for RuntimeOperatorSummaryExecutionStateEnum
func GetRuntimeOperatorSummaryExecutionStateEnumValues() []RuntimeOperatorSummaryExecutionStateEnum {
	values := make([]RuntimeOperatorSummaryExecutionStateEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryExecutionStateEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryExecutionStateEnum
func GetRuntimeOperatorSummaryExecutionStateEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"RUNNING",
		"TERMINATED",
		"SUCCESS",
		"ERROR",
		"SKIPPED",
		"UNKNOWN",
	}
}

// GetMappingRuntimeOperatorSummaryExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryExecutionStateEnum(val string) (RuntimeOperatorSummaryExecutionStateEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
