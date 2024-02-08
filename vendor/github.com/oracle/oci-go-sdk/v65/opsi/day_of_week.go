// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// DayOfWeekEnum Enum with underlying type: string
type DayOfWeekEnum string

// Set of constants representing the allowable values for DayOfWeekEnum
const (
	DayOfWeekMonday    DayOfWeekEnum = "MONDAY"
	DayOfWeekTuesday   DayOfWeekEnum = "TUESDAY"
	DayOfWeekWednesday DayOfWeekEnum = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeekEnum = "THURSDAY"
	DayOfWeekFriday    DayOfWeekEnum = "FRIDAY"
	DayOfWeekSaturday  DayOfWeekEnum = "SATURDAY"
	DayOfWeekSunday    DayOfWeekEnum = "SUNDAY"
)

var mappingDayOfWeekEnum = map[string]DayOfWeekEnum{
	"MONDAY":    DayOfWeekMonday,
	"TUESDAY":   DayOfWeekTuesday,
	"WEDNESDAY": DayOfWeekWednesday,
	"THURSDAY":  DayOfWeekThursday,
	"FRIDAY":    DayOfWeekFriday,
	"SATURDAY":  DayOfWeekSaturday,
	"SUNDAY":    DayOfWeekSunday,
}

var mappingDayOfWeekEnumLowerCase = map[string]DayOfWeekEnum{
	"monday":    DayOfWeekMonday,
	"tuesday":   DayOfWeekTuesday,
	"wednesday": DayOfWeekWednesday,
	"thursday":  DayOfWeekThursday,
	"friday":    DayOfWeekFriday,
	"saturday":  DayOfWeekSaturday,
	"sunday":    DayOfWeekSunday,
}

// GetDayOfWeekEnumValues Enumerates the set of values for DayOfWeekEnum
func GetDayOfWeekEnumValues() []DayOfWeekEnum {
	values := make([]DayOfWeekEnum, 0)
	for _, v := range mappingDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetDayOfWeekEnumStringValues Enumerates the set of values in String for DayOfWeekEnum
func GetDayOfWeekEnumStringValues() []string {
	return []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
}

// GetMappingDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDayOfWeekEnum(val string) (DayOfWeekEnum, bool) {
	enum, ok := mappingDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
