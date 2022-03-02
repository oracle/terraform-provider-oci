// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// OperationExecResult Operation execution result for a single input set.
type OperationExecResult struct {

	// Status of the operation job for particular set of input.
	ExecutionStatus OperationExecResultExecutionStatusEnum `mandatory:"false" json:"executionStatus,omitempty"`

	// Error message if execution of operation is failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Metrics of operation execution job.
	Metrics *interface{} `mandatory:"false" json:"metrics"`

	// List of emitted rows for each OUT/INOUT param.
	OutputValues [][]interface{} `mandatory:"false" json:"outputValues"`

	// True, if error message should be displayed on UI.
	IsWhitelistedErrorMessage *bool `mandatory:"false" json:"isWhitelistedErrorMessage"`
}

func (m OperationExecResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperationExecResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOperationExecResultExecutionStatusEnum(string(m.ExecutionStatus)); !ok && m.ExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionStatus: %s. Supported values are: %s.", m.ExecutionStatus, strings.Join(GetOperationExecResultExecutionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OperationExecResultExecutionStatusEnum Enum with underlying type: string
type OperationExecResultExecutionStatusEnum string

// Set of constants representing the allowable values for OperationExecResultExecutionStatusEnum
const (
	OperationExecResultExecutionStatusFailed  OperationExecResultExecutionStatusEnum = "FAILED"
	OperationExecResultExecutionStatusSuccess OperationExecResultExecutionStatusEnum = "SUCCESS"
	OperationExecResultExecutionStatusQueued  OperationExecResultExecutionStatusEnum = "QUEUED"
	OperationExecResultExecutionStatusRunning OperationExecResultExecutionStatusEnum = "RUNNING"
)

var mappingOperationExecResultExecutionStatusEnum = map[string]OperationExecResultExecutionStatusEnum{
	"FAILED":  OperationExecResultExecutionStatusFailed,
	"SUCCESS": OperationExecResultExecutionStatusSuccess,
	"QUEUED":  OperationExecResultExecutionStatusQueued,
	"RUNNING": OperationExecResultExecutionStatusRunning,
}

var mappingOperationExecResultExecutionStatusEnumLowerCase = map[string]OperationExecResultExecutionStatusEnum{
	"failed":  OperationExecResultExecutionStatusFailed,
	"success": OperationExecResultExecutionStatusSuccess,
	"queued":  OperationExecResultExecutionStatusQueued,
	"running": OperationExecResultExecutionStatusRunning,
}

// GetOperationExecResultExecutionStatusEnumValues Enumerates the set of values for OperationExecResultExecutionStatusEnum
func GetOperationExecResultExecutionStatusEnumValues() []OperationExecResultExecutionStatusEnum {
	values := make([]OperationExecResultExecutionStatusEnum, 0)
	for _, v := range mappingOperationExecResultExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationExecResultExecutionStatusEnumStringValues Enumerates the set of values in String for OperationExecResultExecutionStatusEnum
func GetOperationExecResultExecutionStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCESS",
		"QUEUED",
		"RUNNING",
	}
}

// GetMappingOperationExecResultExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationExecResultExecutionStatusEnum(val string) (OperationExecResultExecutionStatusEnum, bool) {
	enum, ok := mappingOperationExecResultExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
