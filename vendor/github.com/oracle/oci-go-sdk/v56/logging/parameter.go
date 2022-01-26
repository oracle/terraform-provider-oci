// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Parameter Parameters that a resource category supports.
type Parameter struct {

	// Parameter name.
	Name *string `mandatory:"true" json:"name"`

	// Parameter type. One of integer, string, boolean.
	Type ParameterTypeEnum `mandatory:"true" json:"type"`

	// Java regex pattern to validate a parameter value.
	Pattern *string `mandatory:"false" json:"pattern"`
}

func (m Parameter) String() string {
	return common.PointerString(m)
}

// ParameterTypeEnum Enum with underlying type: string
type ParameterTypeEnum string

// Set of constants representing the allowable values for ParameterTypeEnum
const (
	ParameterTypeInteger ParameterTypeEnum = "integer"
	ParameterTypeString  ParameterTypeEnum = "string"
	ParameterTypeBoolean ParameterTypeEnum = "boolean"
)

var mappingParameterType = map[string]ParameterTypeEnum{
	"integer": ParameterTypeInteger,
	"string":  ParameterTypeString,
	"boolean": ParameterTypeBoolean,
}

// GetParameterTypeEnumValues Enumerates the set of values for ParameterTypeEnum
func GetParameterTypeEnumValues() []ParameterTypeEnum {
	values := make([]ParameterTypeEnum, 0)
	for _, v := range mappingParameterType {
		values = append(values, v)
	}
	return values
}
