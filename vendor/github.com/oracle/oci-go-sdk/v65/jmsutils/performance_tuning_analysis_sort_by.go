// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"strings"
)

// PerformanceTuningAnalysisSortByEnum Enum with underlying type: string
type PerformanceTuningAnalysisSortByEnum string

// Set of constants representing the allowable values for PerformanceTuningAnalysisSortByEnum
const (
	PerformanceTuningAnalysisSortByTimeCreated  PerformanceTuningAnalysisSortByEnum = "TIME_CREATED"
	PerformanceTuningAnalysisSortByTimeStarted  PerformanceTuningAnalysisSortByEnum = "TIME_STARTED"
	PerformanceTuningAnalysisSortByTimeFinished PerformanceTuningAnalysisSortByEnum = "TIME_FINISHED"
)

var mappingPerformanceTuningAnalysisSortByEnum = map[string]PerformanceTuningAnalysisSortByEnum{
	"TIME_CREATED":  PerformanceTuningAnalysisSortByTimeCreated,
	"TIME_STARTED":  PerformanceTuningAnalysisSortByTimeStarted,
	"TIME_FINISHED": PerformanceTuningAnalysisSortByTimeFinished,
}

var mappingPerformanceTuningAnalysisSortByEnumLowerCase = map[string]PerformanceTuningAnalysisSortByEnum{
	"time_created":  PerformanceTuningAnalysisSortByTimeCreated,
	"time_started":  PerformanceTuningAnalysisSortByTimeStarted,
	"time_finished": PerformanceTuningAnalysisSortByTimeFinished,
}

// GetPerformanceTuningAnalysisSortByEnumValues Enumerates the set of values for PerformanceTuningAnalysisSortByEnum
func GetPerformanceTuningAnalysisSortByEnumValues() []PerformanceTuningAnalysisSortByEnum {
	values := make([]PerformanceTuningAnalysisSortByEnum, 0)
	for _, v := range mappingPerformanceTuningAnalysisSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPerformanceTuningAnalysisSortByEnumStringValues Enumerates the set of values in String for PerformanceTuningAnalysisSortByEnum
func GetPerformanceTuningAnalysisSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"TIME_STARTED",
		"TIME_FINISHED",
	}
}

// GetMappingPerformanceTuningAnalysisSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPerformanceTuningAnalysisSortByEnum(val string) (PerformanceTuningAnalysisSortByEnum, bool) {
	enum, ok := mappingPerformanceTuningAnalysisSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
