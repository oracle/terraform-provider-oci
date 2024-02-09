// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LogTypesEnum Enum with underlying type: string
type LogTypesEnum string

// Set of constants representing the allowable values for LogTypesEnum
const (
	LogTypesLog       LogTypesEnum = "LOG"
	LogTypesLogEvents LogTypesEnum = "LOG_EVENTS"
)

var mappingLogTypesEnum = map[string]LogTypesEnum{
	"LOG":        LogTypesLog,
	"LOG_EVENTS": LogTypesLogEvents,
}

var mappingLogTypesEnumLowerCase = map[string]LogTypesEnum{
	"log":        LogTypesLog,
	"log_events": LogTypesLogEvents,
}

// GetLogTypesEnumValues Enumerates the set of values for LogTypesEnum
func GetLogTypesEnumValues() []LogTypesEnum {
	values := make([]LogTypesEnum, 0)
	for _, v := range mappingLogTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogTypesEnumStringValues Enumerates the set of values in String for LogTypesEnum
func GetLogTypesEnumStringValues() []string {
	return []string{
		"LOG",
		"LOG_EVENTS",
	}
}

// GetMappingLogTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogTypesEnum(val string) (LogTypesEnum, bool) {
	enum, ok := mappingLogTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
