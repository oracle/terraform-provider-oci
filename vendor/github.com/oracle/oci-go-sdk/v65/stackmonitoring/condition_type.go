// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// ConditionTypeEnum Enum with underlying type: string
type ConditionTypeEnum string

// Set of constants representing the allowable values for ConditionTypeEnum
const (
	ConditionTypeFixed        ConditionTypeEnum = "FIXED"
	ConditionTypeAvailability ConditionTypeEnum = "AVAILABILITY"
)

var mappingConditionTypeEnum = map[string]ConditionTypeEnum{
	"FIXED":        ConditionTypeFixed,
	"AVAILABILITY": ConditionTypeAvailability,
}

var mappingConditionTypeEnumLowerCase = map[string]ConditionTypeEnum{
	"fixed":        ConditionTypeFixed,
	"availability": ConditionTypeAvailability,
}

// GetConditionTypeEnumValues Enumerates the set of values for ConditionTypeEnum
func GetConditionTypeEnumValues() []ConditionTypeEnum {
	values := make([]ConditionTypeEnum, 0)
	for _, v := range mappingConditionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionTypeEnumStringValues Enumerates the set of values in String for ConditionTypeEnum
func GetConditionTypeEnumStringValues() []string {
	return []string{
		"FIXED",
		"AVAILABILITY",
	}
}

// GetMappingConditionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionTypeEnum(val string) (ConditionTypeEnum, bool) {
	enum, ok := mappingConditionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
