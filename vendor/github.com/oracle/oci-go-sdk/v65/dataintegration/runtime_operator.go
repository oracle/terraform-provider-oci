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

// RuntimeOperator Runtime operator model which holds the runtime metadata of the task operator executed.
type RuntimeOperator struct {

	// The TaskRun key.
	TaskRunKey *string `mandatory:"false" json:"taskRunKey"`

	// The runtime operator start time.
	StartTimeInMillis *int64 `mandatory:"false" json:"startTimeInMillis"`

	// The runtime operator end time.
	EndTimeInMillis *int64 `mandatory:"false" json:"endTimeInMillis"`

	// status
	Status RuntimeOperatorStatusEnum `mandatory:"false" json:"status,omitempty"`

	// status
	ExecutionState RuntimeOperatorExecutionStateEnum `mandatory:"false" json:"executionState,omitempty"`

	// A list of parameters for the pipeline, this allows certain aspects of the pipeline to be configured when the pipeline is executed.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	// The configuration provider bindings.
	Inputs map[string]ParameterValue `mandatory:"false" json:"inputs"`

	// The configuration provider bindings.
	Outputs map[string]ParameterValue `mandatory:"false" json:"outputs"`

	// A map metrics for the task run.
	Metrics map[string]float32 `mandatory:"false" json:"metrics"`
}

func (m RuntimeOperator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuntimeOperator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeOperatorStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRuntimeOperatorStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorExecutionStateEnum(string(m.ExecutionState)); !ok && m.ExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionState: %s. Supported values are: %s.", m.ExecutionState, strings.Join(GetRuntimeOperatorExecutionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuntimeOperatorStatusEnum Enum with underlying type: string
type RuntimeOperatorStatusEnum string

// Set of constants representing the allowable values for RuntimeOperatorStatusEnum
const (
	RuntimeOperatorStatusNotStarted  RuntimeOperatorStatusEnum = "NOT_STARTED"
	RuntimeOperatorStatusQueued      RuntimeOperatorStatusEnum = "QUEUED"
	RuntimeOperatorStatusRunning     RuntimeOperatorStatusEnum = "RUNNING"
	RuntimeOperatorStatusTerminating RuntimeOperatorStatusEnum = "TERMINATING"
	RuntimeOperatorStatusTerminated  RuntimeOperatorStatusEnum = "TERMINATED"
	RuntimeOperatorStatusSuccess     RuntimeOperatorStatusEnum = "SUCCESS"
	RuntimeOperatorStatusError       RuntimeOperatorStatusEnum = "ERROR"
)

var mappingRuntimeOperatorStatusEnum = map[string]RuntimeOperatorStatusEnum{
	"NOT_STARTED": RuntimeOperatorStatusNotStarted,
	"QUEUED":      RuntimeOperatorStatusQueued,
	"RUNNING":     RuntimeOperatorStatusRunning,
	"TERMINATING": RuntimeOperatorStatusTerminating,
	"TERMINATED":  RuntimeOperatorStatusTerminated,
	"SUCCESS":     RuntimeOperatorStatusSuccess,
	"ERROR":       RuntimeOperatorStatusError,
}

var mappingRuntimeOperatorStatusEnumLowerCase = map[string]RuntimeOperatorStatusEnum{
	"not_started": RuntimeOperatorStatusNotStarted,
	"queued":      RuntimeOperatorStatusQueued,
	"running":     RuntimeOperatorStatusRunning,
	"terminating": RuntimeOperatorStatusTerminating,
	"terminated":  RuntimeOperatorStatusTerminated,
	"success":     RuntimeOperatorStatusSuccess,
	"error":       RuntimeOperatorStatusError,
}

// GetRuntimeOperatorStatusEnumValues Enumerates the set of values for RuntimeOperatorStatusEnum
func GetRuntimeOperatorStatusEnumValues() []RuntimeOperatorStatusEnum {
	values := make([]RuntimeOperatorStatusEnum, 0)
	for _, v := range mappingRuntimeOperatorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorStatusEnumStringValues Enumerates the set of values in String for RuntimeOperatorStatusEnum
func GetRuntimeOperatorStatusEnumStringValues() []string {
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

// GetMappingRuntimeOperatorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorStatusEnum(val string) (RuntimeOperatorStatusEnum, bool) {
	enum, ok := mappingRuntimeOperatorStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorExecutionStateEnum Enum with underlying type: string
type RuntimeOperatorExecutionStateEnum string

// Set of constants representing the allowable values for RuntimeOperatorExecutionStateEnum
const (
	RuntimeOperatorExecutionStateNotStarted RuntimeOperatorExecutionStateEnum = "NOT_STARTED"
	RuntimeOperatorExecutionStateRunning    RuntimeOperatorExecutionStateEnum = "RUNNING"
	RuntimeOperatorExecutionStateTerminated RuntimeOperatorExecutionStateEnum = "TERMINATED"
	RuntimeOperatorExecutionStateSuccess    RuntimeOperatorExecutionStateEnum = "SUCCESS"
	RuntimeOperatorExecutionStateError      RuntimeOperatorExecutionStateEnum = "ERROR"
	RuntimeOperatorExecutionStateSkipped    RuntimeOperatorExecutionStateEnum = "SKIPPED"
	RuntimeOperatorExecutionStateUnknown    RuntimeOperatorExecutionStateEnum = "UNKNOWN"
)

var mappingRuntimeOperatorExecutionStateEnum = map[string]RuntimeOperatorExecutionStateEnum{
	"NOT_STARTED": RuntimeOperatorExecutionStateNotStarted,
	"RUNNING":     RuntimeOperatorExecutionStateRunning,
	"TERMINATED":  RuntimeOperatorExecutionStateTerminated,
	"SUCCESS":     RuntimeOperatorExecutionStateSuccess,
	"ERROR":       RuntimeOperatorExecutionStateError,
	"SKIPPED":     RuntimeOperatorExecutionStateSkipped,
	"UNKNOWN":     RuntimeOperatorExecutionStateUnknown,
}

var mappingRuntimeOperatorExecutionStateEnumLowerCase = map[string]RuntimeOperatorExecutionStateEnum{
	"not_started": RuntimeOperatorExecutionStateNotStarted,
	"running":     RuntimeOperatorExecutionStateRunning,
	"terminated":  RuntimeOperatorExecutionStateTerminated,
	"success":     RuntimeOperatorExecutionStateSuccess,
	"error":       RuntimeOperatorExecutionStateError,
	"skipped":     RuntimeOperatorExecutionStateSkipped,
	"unknown":     RuntimeOperatorExecutionStateUnknown,
}

// GetRuntimeOperatorExecutionStateEnumValues Enumerates the set of values for RuntimeOperatorExecutionStateEnum
func GetRuntimeOperatorExecutionStateEnumValues() []RuntimeOperatorExecutionStateEnum {
	values := make([]RuntimeOperatorExecutionStateEnum, 0)
	for _, v := range mappingRuntimeOperatorExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorExecutionStateEnumStringValues Enumerates the set of values in String for RuntimeOperatorExecutionStateEnum
func GetRuntimeOperatorExecutionStateEnumStringValues() []string {
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

// GetMappingRuntimeOperatorExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorExecutionStateEnum(val string) (RuntimeOperatorExecutionStateEnum, bool) {
	enum, ok := mappingRuntimeOperatorExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
