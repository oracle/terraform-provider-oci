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

// ConnectivityValidation The information about connectivity validation results
type ConnectivityValidation struct {

	// Total number of validation messages
	TotalMessageCount *int `mandatory:"true" json:"totalMessageCount"`

	// Total number of validation error messages
	ErrorMessageCount *int `mandatory:"false" json:"errorMessageCount"`

	// Total number of validation warning messages
	WarnMessageCount *int `mandatory:"false" json:"warnMessageCount"`

	// Total number of validation information messages
	InfoMessageCount *int `mandatory:"false" json:"infoMessageCount"`

	// Detailed information of the connectivity validation.
	ValidationMessages map[string][]ValidationMessage `mandatory:"false" json:"validationMessages"`
}

func (m ConnectivityValidation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectivityValidation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
