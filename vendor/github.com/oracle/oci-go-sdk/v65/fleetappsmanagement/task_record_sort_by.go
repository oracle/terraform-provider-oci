// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// TaskRecordSortByEnum Enum with underlying type: string
type TaskRecordSortByEnum string

// Set of constants representing the allowable values for TaskRecordSortByEnum
const (
	TaskRecordSortByTimeCreated TaskRecordSortByEnum = "timeCreated"
	TaskRecordSortByDisplayName TaskRecordSortByEnum = "displayName"
)

var mappingTaskRecordSortByEnum = map[string]TaskRecordSortByEnum{
	"timeCreated": TaskRecordSortByTimeCreated,
	"displayName": TaskRecordSortByDisplayName,
}

var mappingTaskRecordSortByEnumLowerCase = map[string]TaskRecordSortByEnum{
	"timecreated": TaskRecordSortByTimeCreated,
	"displayname": TaskRecordSortByDisplayName,
}

// GetTaskRecordSortByEnumValues Enumerates the set of values for TaskRecordSortByEnum
func GetTaskRecordSortByEnumValues() []TaskRecordSortByEnum {
	values := make([]TaskRecordSortByEnum, 0)
	for _, v := range mappingTaskRecordSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRecordSortByEnumStringValues Enumerates the set of values in String for TaskRecordSortByEnum
func GetTaskRecordSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingTaskRecordSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRecordSortByEnum(val string) (TaskRecordSortByEnum, bool) {
	enum, ok := mappingTaskRecordSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
