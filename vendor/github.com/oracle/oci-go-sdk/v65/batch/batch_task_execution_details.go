// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTaskExecutionDetails Execution details for a batch task.
type BatchTaskExecutionDetails interface {
}

type batchtaskexecutiondetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *batchtaskexecutiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbatchtaskexecutiondetails batchtaskexecutiondetails
	s := struct {
		Model Unmarshalerbatchtaskexecutiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *batchtaskexecutiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE_TASK_EXECUTION_DETAILS":
		mm := ComputeTaskExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BatchTaskExecutionDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m batchtaskexecutiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m batchtaskexecutiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchTaskExecutionDetailsTypeEnum Enum with underlying type: string
type BatchTaskExecutionDetailsTypeEnum string

// Set of constants representing the allowable values for BatchTaskExecutionDetailsTypeEnum
const (
	BatchTaskExecutionDetailsTypeComputeTaskExecutionDetails BatchTaskExecutionDetailsTypeEnum = "COMPUTE_TASK_EXECUTION_DETAILS"
)

var mappingBatchTaskExecutionDetailsTypeEnum = map[string]BatchTaskExecutionDetailsTypeEnum{
	"COMPUTE_TASK_EXECUTION_DETAILS": BatchTaskExecutionDetailsTypeComputeTaskExecutionDetails,
}

var mappingBatchTaskExecutionDetailsTypeEnumLowerCase = map[string]BatchTaskExecutionDetailsTypeEnum{
	"compute_task_execution_details": BatchTaskExecutionDetailsTypeComputeTaskExecutionDetails,
}

// GetBatchTaskExecutionDetailsTypeEnumValues Enumerates the set of values for BatchTaskExecutionDetailsTypeEnum
func GetBatchTaskExecutionDetailsTypeEnumValues() []BatchTaskExecutionDetailsTypeEnum {
	values := make([]BatchTaskExecutionDetailsTypeEnum, 0)
	for _, v := range mappingBatchTaskExecutionDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskExecutionDetailsTypeEnumStringValues Enumerates the set of values in String for BatchTaskExecutionDetailsTypeEnum
func GetBatchTaskExecutionDetailsTypeEnumStringValues() []string {
	return []string{
		"COMPUTE_TASK_EXECUTION_DETAILS",
	}
}

// GetMappingBatchTaskExecutionDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskExecutionDetailsTypeEnum(val string) (BatchTaskExecutionDetailsTypeEnum, bool) {
	enum, ok := mappingBatchTaskExecutionDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
