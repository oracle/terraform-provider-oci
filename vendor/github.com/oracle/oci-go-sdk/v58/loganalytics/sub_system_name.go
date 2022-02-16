// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// SubSystemNameEnum Enum with underlying type: string
type SubSystemNameEnum string

// Set of constants representing the allowable values for SubSystemNameEnum
const (
	SubSystemNameLog SubSystemNameEnum = "LOG"
)

var mappingSubSystemNameEnum = map[string]SubSystemNameEnum{
	"LOG": SubSystemNameLog,
}

// GetSubSystemNameEnumValues Enumerates the set of values for SubSystemNameEnum
func GetSubSystemNameEnumValues() []SubSystemNameEnum {
	values := make([]SubSystemNameEnum, 0)
	for _, v := range mappingSubSystemNameEnum {
		values = append(values, v)
	}
	return values
}

// GetSubSystemNameEnumStringValues Enumerates the set of values in String for SubSystemNameEnum
func GetSubSystemNameEnumStringValues() []string {
	return []string{
		"LOG",
	}
}

// GetMappingSubSystemNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubSystemNameEnum(val string) (SubSystemNameEnum, bool) {
	mappingSubSystemNameEnumIgnoreCase := make(map[string]SubSystemNameEnum)
	for k, v := range mappingSubSystemNameEnum {
		mappingSubSystemNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSubSystemNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
