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

// ConditionValueTypeEnum Enum with underlying type: string
type ConditionValueTypeEnum string

// Set of constants representing the allowable values for ConditionValueTypeEnum
const (
	ConditionValueTypeManaged ConditionValueTypeEnum = "MANAGED"
	ConditionValueTypeCustom  ConditionValueTypeEnum = "CUSTOM"
)

var mappingConditionValueTypeEnum = map[string]ConditionValueTypeEnum{
	"MANAGED": ConditionValueTypeManaged,
	"CUSTOM":  ConditionValueTypeCustom,
}

// GetConditionValueTypeEnumValues Enumerates the set of values for ConditionValueTypeEnum
func GetConditionValueTypeEnumValues() []ConditionValueTypeEnum {
	values := make([]ConditionValueTypeEnum, 0)
	for _, v := range mappingConditionValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionValueTypeEnumStringValues Enumerates the set of values in String for ConditionValueTypeEnum
func GetConditionValueTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingConditionValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionValueTypeEnum(val string) (ConditionValueTypeEnum, bool) {
	mappingConditionValueTypeEnumIgnoreCase := make(map[string]ConditionValueTypeEnum)
	for k, v := range mappingConditionValueTypeEnum {
		mappingConditionValueTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConditionValueTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
