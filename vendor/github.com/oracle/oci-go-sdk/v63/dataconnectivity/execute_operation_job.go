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
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// ExecuteOperationJob Response of executeOperationJob.
type ExecuteOperationJob struct {

	// Status of the operation job for all sets of input.
	OperationStatus *string `mandatory:"true" json:"operationStatus"`

	// List of operation execution result for each input set.
	OperationResult []OperationExecResult `mandatory:"true" json:"operationResult"`

	// Error message, if whole operation is failed.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// Name of the operation.
	OperationName *string `mandatory:"false" json:"operationName"`

	// List of names of OUT/INOUT params.
	OutParams []string `mandatory:"false" json:"outParams"`
}

func (m ExecuteOperationJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteOperationJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
