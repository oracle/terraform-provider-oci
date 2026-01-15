// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"strings"
)

// ScheduledQueryProcessingTypeEnum Enum with underlying type: string
type ScheduledQueryProcessingTypeEnum string

// Set of constants representing the allowable values for ScheduledQueryProcessingTypeEnum
const (
	ScheduledQueryProcessingTypeExport ScheduledQueryProcessingTypeEnum = "EXPORT"
	ScheduledQueryProcessingTypeQuery  ScheduledQueryProcessingTypeEnum = "QUERY"
	ScheduledQueryProcessingTypeAlert  ScheduledQueryProcessingTypeEnum = "ALERT"
)

var mappingScheduledQueryProcessingTypeEnum = map[string]ScheduledQueryProcessingTypeEnum{
	"EXPORT": ScheduledQueryProcessingTypeExport,
	"QUERY":  ScheduledQueryProcessingTypeQuery,
	"ALERT":  ScheduledQueryProcessingTypeAlert,
}

var mappingScheduledQueryProcessingTypeEnumLowerCase = map[string]ScheduledQueryProcessingTypeEnum{
	"export": ScheduledQueryProcessingTypeExport,
	"query":  ScheduledQueryProcessingTypeQuery,
	"alert":  ScheduledQueryProcessingTypeAlert,
}

// GetScheduledQueryProcessingTypeEnumValues Enumerates the set of values for ScheduledQueryProcessingTypeEnum
func GetScheduledQueryProcessingTypeEnumValues() []ScheduledQueryProcessingTypeEnum {
	values := make([]ScheduledQueryProcessingTypeEnum, 0)
	for _, v := range mappingScheduledQueryProcessingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledQueryProcessingTypeEnumStringValues Enumerates the set of values in String for ScheduledQueryProcessingTypeEnum
func GetScheduledQueryProcessingTypeEnumStringValues() []string {
	return []string{
		"EXPORT",
		"QUERY",
		"ALERT",
	}
}

// GetMappingScheduledQueryProcessingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledQueryProcessingTypeEnum(val string) (ScheduledQueryProcessingTypeEnum, bool) {
	enum, ok := mappingScheduledQueryProcessingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
