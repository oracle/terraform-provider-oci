// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FullPushDownTaskResponse The full pushdown task.
type FullPushDownTaskResponse struct {

	// The type of of FullPushDownTask response.
	ModelType *string `mandatory:"true" json:"modelType"`

	// The error message in the response object.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The status of FullPushDownTask.
	Status *string `mandatory:"false" json:"status"`
}

func (m FullPushDownTaskResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FullPushDownTaskResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
