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

// PropertyDefinition Defines an property of a macro, variable or query.
type PropertyDefinition struct {

	// Name of the property.
	Name *string `mandatory:"false" json:"name"`

	// Value of the property.
	Value *string `mandatory:"false" json:"value"`

	// Type of the property.
	Type PropertyDefinitionTypeEnum `mandatory:"false" json:"type,omitempty"`

	// True if property is for all macros.  Not applicable for macro variables or query.
	IsGlobal *bool `mandatory:"false" json:"isGlobal"`
}

func (m PropertyDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PropertyDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPropertyDefinitionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPropertyDefinitionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PropertyDefinitionTypeEnum Enum with underlying type: string
type PropertyDefinitionTypeEnum string

// Set of constants representing the allowable values for PropertyDefinitionTypeEnum
const (
	PropertyDefinitionTypeString    PropertyDefinitionTypeEnum = "STRING"
	PropertyDefinitionTypeDouble    PropertyDefinitionTypeEnum = "DOUBLE"
	PropertyDefinitionTypeFloat     PropertyDefinitionTypeEnum = "FLOAT"
	PropertyDefinitionTypeLong      PropertyDefinitionTypeEnum = "LONG"
	PropertyDefinitionTypeInteger   PropertyDefinitionTypeEnum = "INTEGER"
	PropertyDefinitionTypeTimestamp PropertyDefinitionTypeEnum = "TIMESTAMP"
)

var mappingPropertyDefinitionTypeEnum = map[string]PropertyDefinitionTypeEnum{
	"STRING":    PropertyDefinitionTypeString,
	"DOUBLE":    PropertyDefinitionTypeDouble,
	"FLOAT":     PropertyDefinitionTypeFloat,
	"LONG":      PropertyDefinitionTypeLong,
	"INTEGER":   PropertyDefinitionTypeInteger,
	"TIMESTAMP": PropertyDefinitionTypeTimestamp,
}

var mappingPropertyDefinitionTypeEnumLowerCase = map[string]PropertyDefinitionTypeEnum{
	"string":    PropertyDefinitionTypeString,
	"double":    PropertyDefinitionTypeDouble,
	"float":     PropertyDefinitionTypeFloat,
	"long":      PropertyDefinitionTypeLong,
	"integer":   PropertyDefinitionTypeInteger,
	"timestamp": PropertyDefinitionTypeTimestamp,
}

// GetPropertyDefinitionTypeEnumValues Enumerates the set of values for PropertyDefinitionTypeEnum
func GetPropertyDefinitionTypeEnumValues() []PropertyDefinitionTypeEnum {
	values := make([]PropertyDefinitionTypeEnum, 0)
	for _, v := range mappingPropertyDefinitionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertyDefinitionTypeEnumStringValues Enumerates the set of values in String for PropertyDefinitionTypeEnum
func GetPropertyDefinitionTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"DOUBLE",
		"FLOAT",
		"LONG",
		"INTEGER",
		"TIMESTAMP",
	}
}

// GetMappingPropertyDefinitionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertyDefinitionTypeEnum(val string) (PropertyDefinitionTypeEnum, bool) {
	enum, ok := mappingPropertyDefinitionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
