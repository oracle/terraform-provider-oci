// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// PerformanceTuningAnalysisResultSortByEnum Enum with underlying type: string
type PerformanceTuningAnalysisResultSortByEnum string

// Set of constants representing the allowable values for PerformanceTuningAnalysisResultSortByEnum
const (
	PerformanceTuningAnalysisResultSortByTimeCreated       PerformanceTuningAnalysisResultSortByEnum = "timeCreated"
	PerformanceTuningAnalysisResultSortByManagedInstanceId PerformanceTuningAnalysisResultSortByEnum = "managedInstanceId"
	PerformanceTuningAnalysisResultSortByWorkRequestId     PerformanceTuningAnalysisResultSortByEnum = "workRequestId"
)

var mappingPerformanceTuningAnalysisResultSortByEnum = map[string]PerformanceTuningAnalysisResultSortByEnum{
	"timeCreated":       PerformanceTuningAnalysisResultSortByTimeCreated,
	"managedInstanceId": PerformanceTuningAnalysisResultSortByManagedInstanceId,
	"workRequestId":     PerformanceTuningAnalysisResultSortByWorkRequestId,
}

var mappingPerformanceTuningAnalysisResultSortByEnumLowerCase = map[string]PerformanceTuningAnalysisResultSortByEnum{
	"timecreated":       PerformanceTuningAnalysisResultSortByTimeCreated,
	"managedinstanceid": PerformanceTuningAnalysisResultSortByManagedInstanceId,
	"workrequestid":     PerformanceTuningAnalysisResultSortByWorkRequestId,
}

// GetPerformanceTuningAnalysisResultSortByEnumValues Enumerates the set of values for PerformanceTuningAnalysisResultSortByEnum
func GetPerformanceTuningAnalysisResultSortByEnumValues() []PerformanceTuningAnalysisResultSortByEnum {
	values := make([]PerformanceTuningAnalysisResultSortByEnum, 0)
	for _, v := range mappingPerformanceTuningAnalysisResultSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPerformanceTuningAnalysisResultSortByEnumStringValues Enumerates the set of values in String for PerformanceTuningAnalysisResultSortByEnum
func GetPerformanceTuningAnalysisResultSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
	}
}

// GetMappingPerformanceTuningAnalysisResultSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPerformanceTuningAnalysisResultSortByEnum(val string) (PerformanceTuningAnalysisResultSortByEnum, bool) {
	enum, ok := mappingPerformanceTuningAnalysisResultSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
