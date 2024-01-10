// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// DayEnum Enum with underlying type: string
type DayEnum string

// Set of constants representing the allowable values for DayEnum
const (
	DayMonday    DayEnum = "MONDAY"
	DayTuesday   DayEnum = "TUESDAY"
	DayWednesday DayEnum = "WEDNESDAY"
	DayThursday  DayEnum = "THURSDAY"
	DayFriday    DayEnum = "FRIDAY"
	DaySaturday  DayEnum = "SATURDAY"
	DaySunday    DayEnum = "SUNDAY"
)

var mappingDayEnum = map[string]DayEnum{
	"MONDAY":    DayMonday,
	"TUESDAY":   DayTuesday,
	"WEDNESDAY": DayWednesday,
	"THURSDAY":  DayThursday,
	"FRIDAY":    DayFriday,
	"SATURDAY":  DaySaturday,
	"SUNDAY":    DaySunday,
}

var mappingDayEnumLowerCase = map[string]DayEnum{
	"monday":    DayMonday,
	"tuesday":   DayTuesday,
	"wednesday": DayWednesday,
	"thursday":  DayThursday,
	"friday":    DayFriday,
	"saturday":  DaySaturday,
	"sunday":    DaySunday,
}

// GetDayEnumValues Enumerates the set of values for DayEnum
func GetDayEnumValues() []DayEnum {
	values := make([]DayEnum, 0)
	for _, v := range mappingDayEnum {
		values = append(values, v)
	}
	return values
}

// GetDayEnumStringValues Enumerates the set of values in String for DayEnum
func GetDayEnumStringValues() []string {
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

// GetMappingDayEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDayEnum(val string) (DayEnum, bool) {
	enum, ok := mappingDayEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
