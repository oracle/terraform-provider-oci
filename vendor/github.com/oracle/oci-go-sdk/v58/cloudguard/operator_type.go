// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// OperatorTypeEnum Enum with underlying type: string
type OperatorTypeEnum string

// Set of constants representing the allowable values for OperatorTypeEnum
const (
	OperatorTypeIn        OperatorTypeEnum = "IN"
	OperatorTypeNotIn     OperatorTypeEnum = "NOT_IN"
	OperatorTypeEquals    OperatorTypeEnum = "EQUALS"
	OperatorTypeNotEquals OperatorTypeEnum = "NOT_EQUALS"
)

var mappingOperatorTypeEnum = map[string]OperatorTypeEnum{
	"IN":         OperatorTypeIn,
	"NOT_IN":     OperatorTypeNotIn,
	"EQUALS":     OperatorTypeEquals,
	"NOT_EQUALS": OperatorTypeNotEquals,
}

// GetOperatorTypeEnumValues Enumerates the set of values for OperatorTypeEnum
func GetOperatorTypeEnumValues() []OperatorTypeEnum {
	values := make([]OperatorTypeEnum, 0)
	for _, v := range mappingOperatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorTypeEnumStringValues Enumerates the set of values in String for OperatorTypeEnum
func GetOperatorTypeEnumStringValues() []string {
	return []string{
		"IN",
		"NOT_IN",
		"EQUALS",
		"NOT_EQUALS",
	}
}

// GetMappingOperatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorTypeEnum(val string) (OperatorTypeEnum, bool) {
	mappingOperatorTypeEnumIgnoreCase := make(map[string]OperatorTypeEnum)
	for k, v := range mappingOperatorTypeEnum {
		mappingOperatorTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperatorTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
