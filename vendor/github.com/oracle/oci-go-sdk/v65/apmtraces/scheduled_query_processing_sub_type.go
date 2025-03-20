// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ScheduledQueryProcessingSubTypeEnum Enum with underlying type: string
type ScheduledQueryProcessingSubTypeEnum string

// Set of constants representing the allowable values for ScheduledQueryProcessingSubTypeEnum
const (
	ScheduledQueryProcessingSubTypeObjectStorage     ScheduledQueryProcessingSubTypeEnum = "OBJECT_STORAGE"
	ScheduledQueryProcessingSubTypeStreaming         ScheduledQueryProcessingSubTypeEnum = "STREAMING"
	ScheduledQueryProcessingSubTypeMonitoringService ScheduledQueryProcessingSubTypeEnum = "MONITORING_SERVICE"
	ScheduledQueryProcessingSubTypeNone              ScheduledQueryProcessingSubTypeEnum = "NONE"
)

var mappingScheduledQueryProcessingSubTypeEnum = map[string]ScheduledQueryProcessingSubTypeEnum{
	"OBJECT_STORAGE":     ScheduledQueryProcessingSubTypeObjectStorage,
	"STREAMING":          ScheduledQueryProcessingSubTypeStreaming,
	"MONITORING_SERVICE": ScheduledQueryProcessingSubTypeMonitoringService,
	"NONE":               ScheduledQueryProcessingSubTypeNone,
}

var mappingScheduledQueryProcessingSubTypeEnumLowerCase = map[string]ScheduledQueryProcessingSubTypeEnum{
	"object_storage":     ScheduledQueryProcessingSubTypeObjectStorage,
	"streaming":          ScheduledQueryProcessingSubTypeStreaming,
	"monitoring_service": ScheduledQueryProcessingSubTypeMonitoringService,
	"none":               ScheduledQueryProcessingSubTypeNone,
}

// GetScheduledQueryProcessingSubTypeEnumValues Enumerates the set of values for ScheduledQueryProcessingSubTypeEnum
func GetScheduledQueryProcessingSubTypeEnumValues() []ScheduledQueryProcessingSubTypeEnum {
	values := make([]ScheduledQueryProcessingSubTypeEnum, 0)
	for _, v := range mappingScheduledQueryProcessingSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledQueryProcessingSubTypeEnumStringValues Enumerates the set of values in String for ScheduledQueryProcessingSubTypeEnum
func GetScheduledQueryProcessingSubTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
		"STREAMING",
		"MONITORING_SERVICE",
		"NONE",
	}
}

// GetMappingScheduledQueryProcessingSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledQueryProcessingSubTypeEnum(val string) (ScheduledQueryProcessingSubTypeEnum, bool) {
	enum, ok := mappingScheduledQueryProcessingSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
