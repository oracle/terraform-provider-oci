// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// ValueTypeEnum Enum with underlying type: string
type ValueTypeEnum string

// Set of constants representing the allowable values for ValueTypeEnum
const (
	ValueTypeBoolean   ValueTypeEnum = "BOOLEAN"
	ValueTypeString    ValueTypeEnum = "STRING"
	ValueTypeDouble    ValueTypeEnum = "DOUBLE"
	ValueTypeFloat     ValueTypeEnum = "FLOAT"
	ValueTypeLong      ValueTypeEnum = "LONG"
	ValueTypeInteger   ValueTypeEnum = "INTEGER"
	ValueTypeTimestamp ValueTypeEnum = "TIMESTAMP"
	ValueTypeFacet     ValueTypeEnum = "FACET"
	ValueTypeTable     ValueTypeEnum = "TABLE"
)

var mappingValueTypeEnum = map[string]ValueTypeEnum{
	"BOOLEAN":   ValueTypeBoolean,
	"STRING":    ValueTypeString,
	"DOUBLE":    ValueTypeDouble,
	"FLOAT":     ValueTypeFloat,
	"LONG":      ValueTypeLong,
	"INTEGER":   ValueTypeInteger,
	"TIMESTAMP": ValueTypeTimestamp,
	"FACET":     ValueTypeFacet,
	"TABLE":     ValueTypeTable,
}

var mappingValueTypeEnumLowerCase = map[string]ValueTypeEnum{
	"boolean":   ValueTypeBoolean,
	"string":    ValueTypeString,
	"double":    ValueTypeDouble,
	"float":     ValueTypeFloat,
	"long":      ValueTypeLong,
	"integer":   ValueTypeInteger,
	"timestamp": ValueTypeTimestamp,
	"facet":     ValueTypeFacet,
	"table":     ValueTypeTable,
}

// GetValueTypeEnumValues Enumerates the set of values for ValueTypeEnum
func GetValueTypeEnumValues() []ValueTypeEnum {
	values := make([]ValueTypeEnum, 0)
	for _, v := range mappingValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetValueTypeEnumStringValues Enumerates the set of values in String for ValueTypeEnum
func GetValueTypeEnumStringValues() []string {
	return []string{
		"BOOLEAN",
		"STRING",
		"DOUBLE",
		"FLOAT",
		"LONG",
		"INTEGER",
		"TIMESTAMP",
		"FACET",
		"TABLE",
	}
}

// GetMappingValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValueTypeEnum(val string) (ValueTypeEnum, bool) {
	enum, ok := mappingValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
