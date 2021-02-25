// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v35/common"
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
