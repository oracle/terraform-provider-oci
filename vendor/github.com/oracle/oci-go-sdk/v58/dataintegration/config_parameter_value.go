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
	"github.com/oracle/oci-go-sdk/v58/common"
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

	// The root object value, used in custom parameters.
	RootObjectValue *interface{} `mandatory:"false" json:"rootObjectValue"`
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
