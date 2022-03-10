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
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// ExecuteOperationJobDetails Contains details of executeOperationJob.
type ExecuteOperationJobDetails struct {

	// Job id to track job status.
	ExecuteOperationJobId *string `mandatory:"true" json:"executeOperationJobId"`

	// The status of job
	Status ExecuteOperationJobDetailsStatusEnum `mandatory:"true" json:"status"`

	// Error message if job creation is failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m ExecuteOperationJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteOperationJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecuteOperationJobDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetExecuteOperationJobDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteOperationJobDetailsStatusEnum Enum with underlying type: string
type ExecuteOperationJobDetailsStatusEnum string

// Set of constants representing the allowable values for ExecuteOperationJobDetailsStatusEnum
const (
	ExecuteOperationJobDetailsStatusFailed  ExecuteOperationJobDetailsStatusEnum = "FAILED"
	ExecuteOperationJobDetailsStatusSuccess ExecuteOperationJobDetailsStatusEnum = "SUCCESS"
)

var mappingExecuteOperationJobDetailsStatusEnum = map[string]ExecuteOperationJobDetailsStatusEnum{
	"FAILED":  ExecuteOperationJobDetailsStatusFailed,
	"SUCCESS": ExecuteOperationJobDetailsStatusSuccess,
}

var mappingExecuteOperationJobDetailsStatusEnumLowerCase = map[string]ExecuteOperationJobDetailsStatusEnum{
	"failed":  ExecuteOperationJobDetailsStatusFailed,
	"success": ExecuteOperationJobDetailsStatusSuccess,
}

// GetExecuteOperationJobDetailsStatusEnumValues Enumerates the set of values for ExecuteOperationJobDetailsStatusEnum
func GetExecuteOperationJobDetailsStatusEnumValues() []ExecuteOperationJobDetailsStatusEnum {
	values := make([]ExecuteOperationJobDetailsStatusEnum, 0)
	for _, v := range mappingExecuteOperationJobDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteOperationJobDetailsStatusEnumStringValues Enumerates the set of values in String for ExecuteOperationJobDetailsStatusEnum
func GetExecuteOperationJobDetailsStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCESS",
	}
}

// GetMappingExecuteOperationJobDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteOperationJobDetailsStatusEnum(val string) (ExecuteOperationJobDetailsStatusEnum, bool) {
	enum, ok := mappingExecuteOperationJobDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
