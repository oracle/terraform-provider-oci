// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// AcceptableDowntimeEnum Enum with underlying type: string
type AcceptableDowntimeEnum string

// Set of constants representing the allowable values for AcceptableDowntimeEnum
const (
	AcceptableDowntimeLessThan10Minutes AcceptableDowntimeEnum = "LESS_THAN_10_MINUTES"
	AcceptableDowntimeLessThan1Hour     AcceptableDowntimeEnum = "LESS_THAN_1_HOUR"
	AcceptableDowntimeLessThan4Hours    AcceptableDowntimeEnum = "LESS_THAN_4_HOURS"
	AcceptableDowntimeLessThan8Hours    AcceptableDowntimeEnum = "LESS_THAN_8_HOURS"
	AcceptableDowntimeLessThan12Hours   AcceptableDowntimeEnum = "LESS_THAN_12_HOURS"
	AcceptableDowntimeLessThan1Day      AcceptableDowntimeEnum = "LESS_THAN_1_DAY"
	AcceptableDowntimeLessThan2Days     AcceptableDowntimeEnum = "LESS_THAN_2_DAYS"
	AcceptableDowntimeMoreThan2Days     AcceptableDowntimeEnum = "MORE_THAN_2_DAYS"
)

var mappingAcceptableDowntimeEnum = map[string]AcceptableDowntimeEnum{
	"LESS_THAN_10_MINUTES": AcceptableDowntimeLessThan10Minutes,
	"LESS_THAN_1_HOUR":     AcceptableDowntimeLessThan1Hour,
	"LESS_THAN_4_HOURS":    AcceptableDowntimeLessThan4Hours,
	"LESS_THAN_8_HOURS":    AcceptableDowntimeLessThan8Hours,
	"LESS_THAN_12_HOURS":   AcceptableDowntimeLessThan12Hours,
	"LESS_THAN_1_DAY":      AcceptableDowntimeLessThan1Day,
	"LESS_THAN_2_DAYS":     AcceptableDowntimeLessThan2Days,
	"MORE_THAN_2_DAYS":     AcceptableDowntimeMoreThan2Days,
}

var mappingAcceptableDowntimeEnumLowerCase = map[string]AcceptableDowntimeEnum{
	"less_than_10_minutes": AcceptableDowntimeLessThan10Minutes,
	"less_than_1_hour":     AcceptableDowntimeLessThan1Hour,
	"less_than_4_hours":    AcceptableDowntimeLessThan4Hours,
	"less_than_8_hours":    AcceptableDowntimeLessThan8Hours,
	"less_than_12_hours":   AcceptableDowntimeLessThan12Hours,
	"less_than_1_day":      AcceptableDowntimeLessThan1Day,
	"less_than_2_days":     AcceptableDowntimeLessThan2Days,
	"more_than_2_days":     AcceptableDowntimeMoreThan2Days,
}

// GetAcceptableDowntimeEnumValues Enumerates the set of values for AcceptableDowntimeEnum
func GetAcceptableDowntimeEnumValues() []AcceptableDowntimeEnum {
	values := make([]AcceptableDowntimeEnum, 0)
	for _, v := range mappingAcceptableDowntimeEnum {
		values = append(values, v)
	}
	return values
}

// GetAcceptableDowntimeEnumStringValues Enumerates the set of values in String for AcceptableDowntimeEnum
func GetAcceptableDowntimeEnumStringValues() []string {
	return []string{
		"LESS_THAN_10_MINUTES",
		"LESS_THAN_1_HOUR",
		"LESS_THAN_4_HOURS",
		"LESS_THAN_8_HOURS",
		"LESS_THAN_12_HOURS",
		"LESS_THAN_1_DAY",
		"LESS_THAN_2_DAYS",
		"MORE_THAN_2_DAYS",
	}
}

// GetMappingAcceptableDowntimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAcceptableDowntimeEnum(val string) (AcceptableDowntimeEnum, bool) {
	enum, ok := mappingAcceptableDowntimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
