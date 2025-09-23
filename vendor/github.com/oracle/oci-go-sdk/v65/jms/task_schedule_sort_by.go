// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// TaskScheduleSortByEnum Enum with underlying type: string
type TaskScheduleSortByEnum string

// Set of constants representing the allowable values for TaskScheduleSortByEnum
const (
	TaskScheduleSortByName           TaskScheduleSortByEnum = "name"
	TaskScheduleSortByLifecycleState TaskScheduleSortByEnum = "lifecycleState"
	TaskScheduleSortByTimeCreated    TaskScheduleSortByEnum = "timeCreated"
	TaskScheduleSortByTimeNextRun    TaskScheduleSortByEnum = "timeNextRun"
	TaskScheduleSortByTimeLastRun    TaskScheduleSortByEnum = "timeLastRun"
)

var mappingTaskScheduleSortByEnum = map[string]TaskScheduleSortByEnum{
	"name":           TaskScheduleSortByName,
	"lifecycleState": TaskScheduleSortByLifecycleState,
	"timeCreated":    TaskScheduleSortByTimeCreated,
	"timeNextRun":    TaskScheduleSortByTimeNextRun,
	"timeLastRun":    TaskScheduleSortByTimeLastRun,
}

var mappingTaskScheduleSortByEnumLowerCase = map[string]TaskScheduleSortByEnum{
	"name":           TaskScheduleSortByName,
	"lifecyclestate": TaskScheduleSortByLifecycleState,
	"timecreated":    TaskScheduleSortByTimeCreated,
	"timenextrun":    TaskScheduleSortByTimeNextRun,
	"timelastrun":    TaskScheduleSortByTimeLastRun,
}

// GetTaskScheduleSortByEnumValues Enumerates the set of values for TaskScheduleSortByEnum
func GetTaskScheduleSortByEnumValues() []TaskScheduleSortByEnum {
	values := make([]TaskScheduleSortByEnum, 0)
	for _, v := range mappingTaskScheduleSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskScheduleSortByEnumStringValues Enumerates the set of values in String for TaskScheduleSortByEnum
func GetTaskScheduleSortByEnumStringValues() []string {
	return []string{
		"name",
		"lifecycleState",
		"timeCreated",
		"timeNextRun",
		"timeLastRun",
	}
}

// GetMappingTaskScheduleSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskScheduleSortByEnum(val string) (TaskScheduleSortByEnum, bool) {
	enum, ok := mappingTaskScheduleSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
