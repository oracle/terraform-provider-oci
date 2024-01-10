// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ParameterDefinition A parameter to a resource.
type ParameterDefinition struct {

	// The name of the parameter
	Name *string `mandatory:"true" json:"name"`

	// Enumerated parameter type.
	Type ParameterDefinitionTypeEnum `mandatory:"true" json:"type"`

	// Description of the parameter.
	Description *string `mandatory:"false" json:"description"`

	// Is this parameter required. Ignored for parameters with direction = OUTPUT.
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Is the data for this parameter sensitive (e.g. should the data be hidden in UI, encrypted if stored, etc.)
	IsSensitive *bool `mandatory:"false" json:"isSensitive"`

	// Default value for the parameter.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Used for character string types such as STRING to constrain the length of the value
	MinLength *int `mandatory:"false" json:"minLength"`

	// Used for character string types such as STRING to constrain the length of the value
	MaxLength *int `mandatory:"false" json:"maxLength"`

	// Regular expression used to validate the value of a string type such as STRING
	Pattern *string `mandatory:"false" json:"pattern"`

	// Is this parameter an input parameter, output parameter, or both?
	Direction ParameterDefinitionDirectionEnum `mandatory:"false" json:"direction,omitempty"`

	// A forward-slash-delimited 'path' in an imaginary hierarchy, at which this parameter's UI widgets should be placed
	UiPlacementHint *string `mandatory:"false" json:"uiPlacementHint"`

	// Any configuration needed to help the resource type process this parameter (e.g. link to manifest, etc.).
	ResourceTypeMetadata *interface{} `mandatory:"false" json:"resourceTypeMetadata"`
}

func (m ParameterDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ParameterDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingParameterDefinitionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetParameterDefinitionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingParameterDefinitionDirectionEnum(string(m.Direction)); !ok && m.Direction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Direction: %s. Supported values are: %s.", m.Direction, strings.Join(GetParameterDefinitionDirectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ParameterDefinitionTypeEnum Enum with underlying type: string
type ParameterDefinitionTypeEnum string

// Set of constants representing the allowable values for ParameterDefinitionTypeEnum
const (
	ParameterDefinitionTypeString  ParameterDefinitionTypeEnum = "STRING"
	ParameterDefinitionTypeUri     ParameterDefinitionTypeEnum = "URI"
	ParameterDefinitionTypeUrl     ParameterDefinitionTypeEnum = "URL"
	ParameterDefinitionTypeNumber  ParameterDefinitionTypeEnum = "NUMBER"
	ParameterDefinitionTypeBoolean ParameterDefinitionTypeEnum = "BOOLEAN"
)

var mappingParameterDefinitionTypeEnum = map[string]ParameterDefinitionTypeEnum{
	"STRING":  ParameterDefinitionTypeString,
	"URI":     ParameterDefinitionTypeUri,
	"URL":     ParameterDefinitionTypeUrl,
	"NUMBER":  ParameterDefinitionTypeNumber,
	"BOOLEAN": ParameterDefinitionTypeBoolean,
}

var mappingParameterDefinitionTypeEnumLowerCase = map[string]ParameterDefinitionTypeEnum{
	"string":  ParameterDefinitionTypeString,
	"uri":     ParameterDefinitionTypeUri,
	"url":     ParameterDefinitionTypeUrl,
	"number":  ParameterDefinitionTypeNumber,
	"boolean": ParameterDefinitionTypeBoolean,
}

// GetParameterDefinitionTypeEnumValues Enumerates the set of values for ParameterDefinitionTypeEnum
func GetParameterDefinitionTypeEnumValues() []ParameterDefinitionTypeEnum {
	values := make([]ParameterDefinitionTypeEnum, 0)
	for _, v := range mappingParameterDefinitionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterDefinitionTypeEnumStringValues Enumerates the set of values in String for ParameterDefinitionTypeEnum
func GetParameterDefinitionTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"URI",
		"URL",
		"NUMBER",
		"BOOLEAN",
	}
}

// GetMappingParameterDefinitionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterDefinitionTypeEnum(val string) (ParameterDefinitionTypeEnum, bool) {
	enum, ok := mappingParameterDefinitionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ParameterDefinitionDirectionEnum Enum with underlying type: string
type ParameterDefinitionDirectionEnum string

// Set of constants representing the allowable values for ParameterDefinitionDirectionEnum
const (
	ParameterDefinitionDirectionInput  ParameterDefinitionDirectionEnum = "INPUT"
	ParameterDefinitionDirectionOutput ParameterDefinitionDirectionEnum = "OUTPUT"
)

var mappingParameterDefinitionDirectionEnum = map[string]ParameterDefinitionDirectionEnum{
	"INPUT":  ParameterDefinitionDirectionInput,
	"OUTPUT": ParameterDefinitionDirectionOutput,
}

var mappingParameterDefinitionDirectionEnumLowerCase = map[string]ParameterDefinitionDirectionEnum{
	"input":  ParameterDefinitionDirectionInput,
	"output": ParameterDefinitionDirectionOutput,
}

// GetParameterDefinitionDirectionEnumValues Enumerates the set of values for ParameterDefinitionDirectionEnum
func GetParameterDefinitionDirectionEnumValues() []ParameterDefinitionDirectionEnum {
	values := make([]ParameterDefinitionDirectionEnum, 0)
	for _, v := range mappingParameterDefinitionDirectionEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterDefinitionDirectionEnumStringValues Enumerates the set of values in String for ParameterDefinitionDirectionEnum
func GetParameterDefinitionDirectionEnumStringValues() []string {
	return []string{
		"INPUT",
		"OUTPUT",
	}
}

// GetMappingParameterDefinitionDirectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterDefinitionDirectionEnum(val string) (ParameterDefinitionDirectionEnum, bool) {
	enum, ok := mappingParameterDefinitionDirectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
