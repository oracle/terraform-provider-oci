// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// ValueTypeEnum Enum with underlying type: string
type ValueTypeEnum string

// Set of constants representing the allowable values for ValueTypeEnum
const (
	ValueTypeString  ValueTypeEnum = "STRING"
	ValueTypeNumeric ValueTypeEnum = "NUMERIC"
)

var mappingValueTypeEnum = map[string]ValueTypeEnum{
	"STRING":  ValueTypeString,
	"NUMERIC": ValueTypeNumeric,
}

var mappingValueTypeEnumLowerCase = map[string]ValueTypeEnum{
	"string":  ValueTypeString,
	"numeric": ValueTypeNumeric,
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
		"STRING",
		"NUMERIC",
	}
}

// GetMappingValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValueTypeEnum(val string) (ValueTypeEnum, bool) {
	enum, ok := mappingValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
