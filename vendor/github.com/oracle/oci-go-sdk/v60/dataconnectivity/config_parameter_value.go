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

// ConfigParameterValue Contains the parameter configuration values.
type ConfigParameterValue struct {

	// A string value of the parameter.
	StringValue *string `mandatory:"false" json:"stringValue"`

	// An integer value of the parameter.
	IntValue *int `mandatory:"false" json:"intValue"`

	// An object value of the parameter.
	ObjectValue *interface{} `mandatory:"false" json:"objectValue"`

	// The root object reference value.
	RefValue *interface{} `mandatory:"false" json:"refValue"`

	// Reference to the parameter by its key.
	ParameterValue *string `mandatory:"false" json:"parameterValue"`
}

func (m ConfigParameterValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigParameterValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
