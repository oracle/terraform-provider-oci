// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ScheduleLifecycleStateEnum Enum with underlying type: string
type ScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduleLifecycleStateEnum
const (
	ScheduleLifecycleStateCreating ScheduleLifecycleStateEnum = "CREATING"
	ScheduleLifecycleStateActive   ScheduleLifecycleStateEnum = "ACTIVE"
	ScheduleLifecycleStateInactive ScheduleLifecycleStateEnum = "INACTIVE"
	ScheduleLifecycleStateUpdating ScheduleLifecycleStateEnum = "UPDATING"
	ScheduleLifecycleStateDeleting ScheduleLifecycleStateEnum = "DELETING"
	ScheduleLifecycleStateDeleted  ScheduleLifecycleStateEnum = "DELETED"
	ScheduleLifecycleStateFailed   ScheduleLifecycleStateEnum = "FAILED"
)

var mappingScheduleLifecycleStateEnum = map[string]ScheduleLifecycleStateEnum{
	"CREATING": ScheduleLifecycleStateCreating,
	"ACTIVE":   ScheduleLifecycleStateActive,
	"INACTIVE": ScheduleLifecycleStateInactive,
	"UPDATING": ScheduleLifecycleStateUpdating,
	"DELETING": ScheduleLifecycleStateDeleting,
	"DELETED":  ScheduleLifecycleStateDeleted,
	"FAILED":   ScheduleLifecycleStateFailed,
}

var mappingScheduleLifecycleStateEnumLowerCase = map[string]ScheduleLifecycleStateEnum{
	"creating": ScheduleLifecycleStateCreating,
	"active":   ScheduleLifecycleStateActive,
	"inactive": ScheduleLifecycleStateInactive,
	"updating": ScheduleLifecycleStateUpdating,
	"deleting": ScheduleLifecycleStateDeleting,
	"deleted":  ScheduleLifecycleStateDeleted,
	"failed":   ScheduleLifecycleStateFailed,
}

// GetScheduleLifecycleStateEnumValues Enumerates the set of values for ScheduleLifecycleStateEnum
func GetScheduleLifecycleStateEnumValues() []ScheduleLifecycleStateEnum {
	values := make([]ScheduleLifecycleStateEnum, 0)
	for _, v := range mappingScheduleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduleLifecycleStateEnum
func GetScheduleLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleLifecycleStateEnum(val string) (ScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
