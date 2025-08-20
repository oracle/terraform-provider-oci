// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VariableDefinition Defines a variable used in a macro or the initization section of a query.
type VariableDefinition struct {

	// Name of the variable preceded by a $.
	Name *string `mandatory:"true" json:"name"`

	// Descripion of the variable to show the user.
	Description *string `mandatory:"false" json:"description"`

	// Type of the variable to show the user.
	Type VariableDefinitionTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Default value of the variable is not already set.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Optional list of properties for the variable.
	Properties []PropertyDefinition `mandatory:"false" json:"properties"`
}

func (m VariableDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VariableDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVariableDefinitionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetVariableDefinitionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VariableDefinitionTypeEnum Enum with underlying type: string
type VariableDefinitionTypeEnum string

// Set of constants representing the allowable values for VariableDefinitionTypeEnum
const (
	VariableDefinitionTypeNumber    VariableDefinitionTypeEnum = "NUMBER"
	VariableDefinitionTypeString    VariableDefinitionTypeEnum = "STRING"
	VariableDefinitionTypeTimestamp VariableDefinitionTypeEnum = "TIMESTAMP"
)

var mappingVariableDefinitionTypeEnum = map[string]VariableDefinitionTypeEnum{
	"NUMBER":    VariableDefinitionTypeNumber,
	"STRING":    VariableDefinitionTypeString,
	"TIMESTAMP": VariableDefinitionTypeTimestamp,
}

var mappingVariableDefinitionTypeEnumLowerCase = map[string]VariableDefinitionTypeEnum{
	"number":    VariableDefinitionTypeNumber,
	"string":    VariableDefinitionTypeString,
	"timestamp": VariableDefinitionTypeTimestamp,
}

// GetVariableDefinitionTypeEnumValues Enumerates the set of values for VariableDefinitionTypeEnum
func GetVariableDefinitionTypeEnumValues() []VariableDefinitionTypeEnum {
	values := make([]VariableDefinitionTypeEnum, 0)
	for _, v := range mappingVariableDefinitionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVariableDefinitionTypeEnumStringValues Enumerates the set of values in String for VariableDefinitionTypeEnum
func GetVariableDefinitionTypeEnumStringValues() []string {
	return []string{
		"NUMBER",
		"STRING",
		"TIMESTAMP",
	}
}

// GetMappingVariableDefinitionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVariableDefinitionTypeEnum(val string) (VariableDefinitionTypeEnum, bool) {
	enum, ok := mappingVariableDefinitionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
