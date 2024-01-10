// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// CapacityTypeEnum Enum with underlying type: string
type CapacityTypeEnum string

// Set of constants representing the allowable values for CapacityTypeEnum
const (
	CapacityTypeOlpuCount CapacityTypeEnum = "OLPU_COUNT"
	CapacityTypeUserCount CapacityTypeEnum = "USER_COUNT"
)

var mappingCapacityTypeEnum = map[string]CapacityTypeEnum{
	"OLPU_COUNT": CapacityTypeOlpuCount,
	"USER_COUNT": CapacityTypeUserCount,
}

var mappingCapacityTypeEnumLowerCase = map[string]CapacityTypeEnum{
	"olpu_count": CapacityTypeOlpuCount,
	"user_count": CapacityTypeUserCount,
}

// GetCapacityTypeEnumValues Enumerates the set of values for CapacityTypeEnum
func GetCapacityTypeEnumValues() []CapacityTypeEnum {
	values := make([]CapacityTypeEnum, 0)
	for _, v := range mappingCapacityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCapacityTypeEnumStringValues Enumerates the set of values in String for CapacityTypeEnum
func GetCapacityTypeEnumStringValues() []string {
	return []string{
		"OLPU_COUNT",
		"USER_COUNT",
	}
}

// GetMappingCapacityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCapacityTypeEnum(val string) (CapacityTypeEnum, bool) {
	enum, ok := mappingCapacityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
