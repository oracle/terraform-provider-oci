// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// MaintenanceRunLifecycleStateEnum Enum with underlying type: string
type MaintenanceRunLifecycleStateEnum string

// Set of constants representing the allowable values for MaintenanceRunLifecycleStateEnum
const (
	MaintenanceRunLifecycleStateCreating       MaintenanceRunLifecycleStateEnum = "CREATING"
	MaintenanceRunLifecycleStateScheduled      MaintenanceRunLifecycleStateEnum = "SCHEDULED"
	MaintenanceRunLifecycleStateInProgress     MaintenanceRunLifecycleStateEnum = "IN_PROGRESS"
	MaintenanceRunLifecycleStateSucceeded      MaintenanceRunLifecycleStateEnum = "SUCCEEDED"
	MaintenanceRunLifecycleStateSkipped        MaintenanceRunLifecycleStateEnum = "SKIPPED"
	MaintenanceRunLifecycleStateFailed         MaintenanceRunLifecycleStateEnum = "FAILED"
	MaintenanceRunLifecycleStateUpdating       MaintenanceRunLifecycleStateEnum = "UPDATING"
	MaintenanceRunLifecycleStateDeleting       MaintenanceRunLifecycleStateEnum = "DELETING"
	MaintenanceRunLifecycleStateDeleted        MaintenanceRunLifecycleStateEnum = "DELETED"
	MaintenanceRunLifecycleStateCanceled       MaintenanceRunLifecycleStateEnum = "CANCELED"
	MaintenanceRunLifecycleStatePartialSuccess MaintenanceRunLifecycleStateEnum = "PARTIAL_SUCCESS"
)

var mappingMaintenanceRunLifecycleStateEnum = map[string]MaintenanceRunLifecycleStateEnum{
	"CREATING":        MaintenanceRunLifecycleStateCreating,
	"SCHEDULED":       MaintenanceRunLifecycleStateScheduled,
	"IN_PROGRESS":     MaintenanceRunLifecycleStateInProgress,
	"SUCCEEDED":       MaintenanceRunLifecycleStateSucceeded,
	"SKIPPED":         MaintenanceRunLifecycleStateSkipped,
	"FAILED":          MaintenanceRunLifecycleStateFailed,
	"UPDATING":        MaintenanceRunLifecycleStateUpdating,
	"DELETING":        MaintenanceRunLifecycleStateDeleting,
	"DELETED":         MaintenanceRunLifecycleStateDeleted,
	"CANCELED":        MaintenanceRunLifecycleStateCanceled,
	"PARTIAL_SUCCESS": MaintenanceRunLifecycleStatePartialSuccess,
}

var mappingMaintenanceRunLifecycleStateEnumLowerCase = map[string]MaintenanceRunLifecycleStateEnum{
	"creating":        MaintenanceRunLifecycleStateCreating,
	"scheduled":       MaintenanceRunLifecycleStateScheduled,
	"in_progress":     MaintenanceRunLifecycleStateInProgress,
	"succeeded":       MaintenanceRunLifecycleStateSucceeded,
	"skipped":         MaintenanceRunLifecycleStateSkipped,
	"failed":          MaintenanceRunLifecycleStateFailed,
	"updating":        MaintenanceRunLifecycleStateUpdating,
	"deleting":        MaintenanceRunLifecycleStateDeleting,
	"deleted":         MaintenanceRunLifecycleStateDeleted,
	"canceled":        MaintenanceRunLifecycleStateCanceled,
	"partial_success": MaintenanceRunLifecycleStatePartialSuccess,
}

// GetMaintenanceRunLifecycleStateEnumValues Enumerates the set of values for MaintenanceRunLifecycleStateEnum
func GetMaintenanceRunLifecycleStateEnumValues() []MaintenanceRunLifecycleStateEnum {
	values := make([]MaintenanceRunLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceRunLifecycleStateEnumStringValues Enumerates the set of values in String for MaintenanceRunLifecycleStateEnum
func GetMaintenanceRunLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SCHEDULED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"SKIPPED",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"CANCELED",
		"PARTIAL_SUCCESS",
	}
}

// GetMappingMaintenanceRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceRunLifecycleStateEnum(val string) (MaintenanceRunLifecycleStateEnum, bool) {
	enum, ok := mappingMaintenanceRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
