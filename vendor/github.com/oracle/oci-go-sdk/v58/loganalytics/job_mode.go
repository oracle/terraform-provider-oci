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

// JobModeEnum Enum with underlying type: string
type JobModeEnum string

// Set of constants representing the allowable values for JobModeEnum
const (
	JobModeForeground JobModeEnum = "FOREGROUND"
	JobModeBackground JobModeEnum = "BACKGROUND"
)

var mappingJobModeEnum = map[string]JobModeEnum{
	"FOREGROUND": JobModeForeground,
	"BACKGROUND": JobModeBackground,
}

// GetJobModeEnumValues Enumerates the set of values for JobModeEnum
func GetJobModeEnumValues() []JobModeEnum {
	values := make([]JobModeEnum, 0)
	for _, v := range mappingJobModeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobModeEnumStringValues Enumerates the set of values in String for JobModeEnum
func GetJobModeEnumStringValues() []string {
	return []string{
		"FOREGROUND",
		"BACKGROUND",
	}
}

// GetMappingJobModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobModeEnum(val string) (JobModeEnum, bool) {
	mappingJobModeEnumIgnoreCase := make(map[string]JobModeEnum)
	for k, v := range mappingJobModeEnum {
		mappingJobModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
