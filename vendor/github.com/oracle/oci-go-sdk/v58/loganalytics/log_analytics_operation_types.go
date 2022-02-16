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

// LogAnalyticsOperationTypesEnum Enum with underlying type: string
type LogAnalyticsOperationTypesEnum string

// Set of constants representing the allowable values for LogAnalyticsOperationTypesEnum
const (
	LogAnalyticsOperationTypesCreateLogAnalytics LogAnalyticsOperationTypesEnum = "CREATE_LOG_ANALYTICS"
	LogAnalyticsOperationTypesDeleteLogAnalytics LogAnalyticsOperationTypesEnum = "DELETE_LOG_ANALYTICS"
)

var mappingLogAnalyticsOperationTypesEnum = map[string]LogAnalyticsOperationTypesEnum{
	"CREATE_LOG_ANALYTICS": LogAnalyticsOperationTypesCreateLogAnalytics,
	"DELETE_LOG_ANALYTICS": LogAnalyticsOperationTypesDeleteLogAnalytics,
}

// GetLogAnalyticsOperationTypesEnumValues Enumerates the set of values for LogAnalyticsOperationTypesEnum
func GetLogAnalyticsOperationTypesEnumValues() []LogAnalyticsOperationTypesEnum {
	values := make([]LogAnalyticsOperationTypesEnum, 0)
	for _, v := range mappingLogAnalyticsOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsOperationTypesEnumStringValues Enumerates the set of values in String for LogAnalyticsOperationTypesEnum
func GetLogAnalyticsOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_LOG_ANALYTICS",
		"DELETE_LOG_ANALYTICS",
	}
}

// GetMappingLogAnalyticsOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsOperationTypesEnum(val string) (LogAnalyticsOperationTypesEnum, bool) {
	mappingLogAnalyticsOperationTypesEnumIgnoreCase := make(map[string]LogAnalyticsOperationTypesEnum)
	for k, v := range mappingLogAnalyticsOperationTypesEnum {
		mappingLogAnalyticsOperationTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLogAnalyticsOperationTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
