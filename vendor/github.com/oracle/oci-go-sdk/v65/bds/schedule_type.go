// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"strings"
)

// ScheduleTypeEnum Enum with underlying type: string
type ScheduleTypeEnum string

// Set of constants representing the allowable values for ScheduleTypeEnum
const (
	ScheduleTypeDayBased ScheduleTypeEnum = "DAY_BASED"
)

var mappingScheduleTypeEnum = map[string]ScheduleTypeEnum{
	"DAY_BASED": ScheduleTypeDayBased,
}

var mappingScheduleTypeEnumLowerCase = map[string]ScheduleTypeEnum{
	"day_based": ScheduleTypeDayBased,
}

// GetScheduleTypeEnumValues Enumerates the set of values for ScheduleTypeEnum
func GetScheduleTypeEnumValues() []ScheduleTypeEnum {
	values := make([]ScheduleTypeEnum, 0)
	for _, v := range mappingScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleTypeEnumStringValues Enumerates the set of values in String for ScheduleTypeEnum
func GetScheduleTypeEnumStringValues() []string {
	return []string{
		"DAY_BASED",
	}
}

// GetMappingScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleTypeEnum(val string) (ScheduleTypeEnum, bool) {
	enum, ok := mappingScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
