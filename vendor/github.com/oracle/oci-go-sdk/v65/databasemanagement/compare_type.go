// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// CompareTypeEnum Enum with underlying type: string
type CompareTypeEnum string

// Set of constants representing the allowable values for CompareTypeEnum
const (
	CompareTypeHour CompareTypeEnum = "HOUR"
	CompareTypeDay  CompareTypeEnum = "DAY"
	CompareTypeWeek CompareTypeEnum = "WEEK"
)

var mappingCompareTypeEnum = map[string]CompareTypeEnum{
	"HOUR": CompareTypeHour,
	"DAY":  CompareTypeDay,
	"WEEK": CompareTypeWeek,
}

var mappingCompareTypeEnumLowerCase = map[string]CompareTypeEnum{
	"hour": CompareTypeHour,
	"day":  CompareTypeDay,
	"week": CompareTypeWeek,
}

// GetCompareTypeEnumValues Enumerates the set of values for CompareTypeEnum
func GetCompareTypeEnumValues() []CompareTypeEnum {
	values := make([]CompareTypeEnum, 0)
	for _, v := range mappingCompareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCompareTypeEnumStringValues Enumerates the set of values in String for CompareTypeEnum
func GetCompareTypeEnumStringValues() []string {
	return []string{
		"HOUR",
		"DAY",
		"WEEK",
	}
}

// GetMappingCompareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompareTypeEnum(val string) (CompareTypeEnum, bool) {
	enum, ok := mappingCompareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
