// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationParameter The parameter of an application.
type ApplicationParameter struct {

	// The name of the parameter.  It must be a string of one or more word characters
	// (a-z, A-Z, 0-9, _).
	// Examples: "iterations", "input_file"
	Name *string `mandatory:"true" json:"name"`

	// The value of the parameter. It must be a string of 0 or more characters of any kind.
	// Examples: "" (empty string), "10", "mydata.xml", "${x}"
	Value *string `mandatory:"true" json:"value"`
}

func (m ApplicationParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
