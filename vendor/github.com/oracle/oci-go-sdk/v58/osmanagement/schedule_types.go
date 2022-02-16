// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// ScheduleTypesEnum Enum with underlying type: string
type ScheduleTypesEnum string

// Set of constants representing the allowable values for ScheduleTypesEnum
const (
	ScheduleTypesOnetime   ScheduleTypesEnum = "ONETIME"
	ScheduleTypesRecurring ScheduleTypesEnum = "RECURRING"
)

var mappingScheduleTypesEnum = map[string]ScheduleTypesEnum{
	"ONETIME":   ScheduleTypesOnetime,
	"RECURRING": ScheduleTypesRecurring,
}

// GetScheduleTypesEnumValues Enumerates the set of values for ScheduleTypesEnum
func GetScheduleTypesEnumValues() []ScheduleTypesEnum {
	values := make([]ScheduleTypesEnum, 0)
	for _, v := range mappingScheduleTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleTypesEnumStringValues Enumerates the set of values in String for ScheduleTypesEnum
func GetScheduleTypesEnumStringValues() []string {
	return []string{
		"ONETIME",
		"RECURRING",
	}
}

// GetMappingScheduleTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleTypesEnum(val string) (ScheduleTypesEnum, bool) {
	mappingScheduleTypesEnumIgnoreCase := make(map[string]ScheduleTypesEnum)
	for k, v := range mappingScheduleTypesEnum {
		mappingScheduleTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingScheduleTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
