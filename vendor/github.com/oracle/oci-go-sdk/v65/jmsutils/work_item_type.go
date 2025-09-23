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

// WorkItemTypeEnum Enum with underlying type: string
type WorkItemTypeEnum string

// Set of constants representing the allowable values for WorkItemTypeEnum
const (
	WorkItemTypePerformanceTuning WorkItemTypeEnum = "PERFORMANCE_TUNING"
	WorkItemTypeJavaMigration     WorkItemTypeEnum = "JAVA_MIGRATION"
)

var mappingWorkItemTypeEnum = map[string]WorkItemTypeEnum{
	"PERFORMANCE_TUNING": WorkItemTypePerformanceTuning,
	"JAVA_MIGRATION":     WorkItemTypeJavaMigration,
}

var mappingWorkItemTypeEnumLowerCase = map[string]WorkItemTypeEnum{
	"performance_tuning": WorkItemTypePerformanceTuning,
	"java_migration":     WorkItemTypeJavaMigration,
}

// GetWorkItemTypeEnumValues Enumerates the set of values for WorkItemTypeEnum
func GetWorkItemTypeEnumValues() []WorkItemTypeEnum {
	values := make([]WorkItemTypeEnum, 0)
	for _, v := range mappingWorkItemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkItemTypeEnumStringValues Enumerates the set of values in String for WorkItemTypeEnum
func GetWorkItemTypeEnumStringValues() []string {
	return []string{
		"PERFORMANCE_TUNING",
		"JAVA_MIGRATION",
	}
}

// GetMappingWorkItemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkItemTypeEnum(val string) (WorkItemTypeEnum, bool) {
	enum, ok := mappingWorkItemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
