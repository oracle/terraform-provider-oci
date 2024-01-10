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

// JobModeFilterEnum Enum with underlying type: string
type JobModeFilterEnum string

// Set of constants representing the allowable values for JobModeFilterEnum
const (
	JobModeFilterAll        JobModeFilterEnum = "ALL"
	JobModeFilterForeground JobModeFilterEnum = "FOREGROUND"
	JobModeFilterBackground JobModeFilterEnum = "BACKGROUND"
)

var mappingJobModeFilterEnum = map[string]JobModeFilterEnum{
	"ALL":        JobModeFilterAll,
	"FOREGROUND": JobModeFilterForeground,
	"BACKGROUND": JobModeFilterBackground,
}

var mappingJobModeFilterEnumLowerCase = map[string]JobModeFilterEnum{
	"all":        JobModeFilterAll,
	"foreground": JobModeFilterForeground,
	"background": JobModeFilterBackground,
}

// GetJobModeFilterEnumValues Enumerates the set of values for JobModeFilterEnum
func GetJobModeFilterEnumValues() []JobModeFilterEnum {
	values := make([]JobModeFilterEnum, 0)
	for _, v := range mappingJobModeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetJobModeFilterEnumStringValues Enumerates the set of values in String for JobModeFilterEnum
func GetJobModeFilterEnumStringValues() []string {
	return []string{
		"ALL",
		"FOREGROUND",
		"BACKGROUND",
	}
}

// GetMappingJobModeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobModeFilterEnum(val string) (JobModeFilterEnum, bool) {
	enum, ok := mappingJobModeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
