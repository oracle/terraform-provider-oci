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

// ConditionTypeEnumEnum Enum with underlying type: string
type ConditionTypeEnumEnum string

// Set of constants representing the allowable values for ConditionTypeEnumEnum
const (
	ConditionTypeEnumActivityCondition   ConditionTypeEnumEnum = "ActivityCondition"
	ConditionTypeEnumSecurityCondition   ConditionTypeEnumEnum = "SecurityCondition"
	ConditionTypeEnumCloudGuardCondition ConditionTypeEnumEnum = "CloudGuardCondition"
)

var mappingConditionTypeEnumEnum = map[string]ConditionTypeEnumEnum{
	"ActivityCondition":   ConditionTypeEnumActivityCondition,
	"SecurityCondition":   ConditionTypeEnumSecurityCondition,
	"CloudGuardCondition": ConditionTypeEnumCloudGuardCondition,
}

// GetConditionTypeEnumEnumValues Enumerates the set of values for ConditionTypeEnumEnum
func GetConditionTypeEnumEnumValues() []ConditionTypeEnumEnum {
	values := make([]ConditionTypeEnumEnum, 0)
	for _, v := range mappingConditionTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionTypeEnumEnumStringValues Enumerates the set of values in String for ConditionTypeEnumEnum
func GetConditionTypeEnumEnumStringValues() []string {
	return []string{
		"ActivityCondition",
		"SecurityCondition",
		"CloudGuardCondition",
	}
}

// GetMappingConditionTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionTypeEnumEnum(val string) (ConditionTypeEnumEnum, bool) {
	mappingConditionTypeEnumEnumIgnoreCase := make(map[string]ConditionTypeEnumEnum)
	for k, v := range mappingConditionTypeEnumEnum {
		mappingConditionTypeEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConditionTypeEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
