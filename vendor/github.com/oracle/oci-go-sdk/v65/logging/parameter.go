// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Parameter Parameters that a resource category supports.
type Parameter struct {

	// Parameter name.
	Name *string `mandatory:"true" json:"name"`

	Type ParameterTypeEnum `mandatory:"true" json:"type"`

	// Java regex pattern to validate a parameter value.
	Pattern *string `mandatory:"false" json:"pattern"`
}

func (m Parameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Parameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingParameterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetParameterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ParameterTypeEnum Enum with underlying type: string
type ParameterTypeEnum string

// Set of constants representing the allowable values for ParameterTypeEnum
const (
	ParameterTypeInteger ParameterTypeEnum = "integer"
	ParameterTypeString  ParameterTypeEnum = "string"
	ParameterTypeBoolean ParameterTypeEnum = "boolean"
)

var mappingParameterTypeEnum = map[string]ParameterTypeEnum{
	"integer": ParameterTypeInteger,
	"string":  ParameterTypeString,
	"boolean": ParameterTypeBoolean,
}

var mappingParameterTypeEnumLowerCase = map[string]ParameterTypeEnum{
	"integer": ParameterTypeInteger,
	"string":  ParameterTypeString,
	"boolean": ParameterTypeBoolean,
}

// GetParameterTypeEnumValues Enumerates the set of values for ParameterTypeEnum
func GetParameterTypeEnumValues() []ParameterTypeEnum {
	values := make([]ParameterTypeEnum, 0)
	for _, v := range mappingParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterTypeEnumStringValues Enumerates the set of values in String for ParameterTypeEnum
func GetParameterTypeEnumStringValues() []string {
	return []string{
		"integer",
		"string",
		"boolean",
	}
}

// GetMappingParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterTypeEnum(val string) (ParameterTypeEnum, bool) {
	enum, ok := mappingParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
