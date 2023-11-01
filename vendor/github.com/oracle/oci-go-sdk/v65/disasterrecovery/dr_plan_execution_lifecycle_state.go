// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanExecutionLifecycleStateEnum Enum with underlying type: string
type DrPlanExecutionLifecycleStateEnum string

// Set of constants representing the allowable values for DrPlanExecutionLifecycleStateEnum
const (
	DrPlanExecutionLifecycleStateAccepted   DrPlanExecutionLifecycleStateEnum = "ACCEPTED"
	DrPlanExecutionLifecycleStateInProgress DrPlanExecutionLifecycleStateEnum = "IN_PROGRESS"
	DrPlanExecutionLifecycleStateWaiting    DrPlanExecutionLifecycleStateEnum = "WAITING"
	DrPlanExecutionLifecycleStateCanceling  DrPlanExecutionLifecycleStateEnum = "CANCELING"
	DrPlanExecutionLifecycleStateCanceled   DrPlanExecutionLifecycleStateEnum = "CANCELED"
	DrPlanExecutionLifecycleStateSucceeded  DrPlanExecutionLifecycleStateEnum = "SUCCEEDED"
	DrPlanExecutionLifecycleStateFailed     DrPlanExecutionLifecycleStateEnum = "FAILED"
	DrPlanExecutionLifecycleStateDeleting   DrPlanExecutionLifecycleStateEnum = "DELETING"
	DrPlanExecutionLifecycleStateDeleted    DrPlanExecutionLifecycleStateEnum = "DELETED"
	DrPlanExecutionLifecycleStatePausing    DrPlanExecutionLifecycleStateEnum = "PAUSING"
	DrPlanExecutionLifecycleStatePaused     DrPlanExecutionLifecycleStateEnum = "PAUSED"
	DrPlanExecutionLifecycleStateResuming   DrPlanExecutionLifecycleStateEnum = "RESUMING"
)

var mappingDrPlanExecutionLifecycleStateEnum = map[string]DrPlanExecutionLifecycleStateEnum{
	"ACCEPTED":    DrPlanExecutionLifecycleStateAccepted,
	"IN_PROGRESS": DrPlanExecutionLifecycleStateInProgress,
	"WAITING":     DrPlanExecutionLifecycleStateWaiting,
	"CANCELING":   DrPlanExecutionLifecycleStateCanceling,
	"CANCELED":    DrPlanExecutionLifecycleStateCanceled,
	"SUCCEEDED":   DrPlanExecutionLifecycleStateSucceeded,
	"FAILED":      DrPlanExecutionLifecycleStateFailed,
	"DELETING":    DrPlanExecutionLifecycleStateDeleting,
	"DELETED":     DrPlanExecutionLifecycleStateDeleted,
	"PAUSING":     DrPlanExecutionLifecycleStatePausing,
	"PAUSED":      DrPlanExecutionLifecycleStatePaused,
	"RESUMING":    DrPlanExecutionLifecycleStateResuming,
}

var mappingDrPlanExecutionLifecycleStateEnumLowerCase = map[string]DrPlanExecutionLifecycleStateEnum{
	"accepted":    DrPlanExecutionLifecycleStateAccepted,
	"in_progress": DrPlanExecutionLifecycleStateInProgress,
	"waiting":     DrPlanExecutionLifecycleStateWaiting,
	"canceling":   DrPlanExecutionLifecycleStateCanceling,
	"canceled":    DrPlanExecutionLifecycleStateCanceled,
	"succeeded":   DrPlanExecutionLifecycleStateSucceeded,
	"failed":      DrPlanExecutionLifecycleStateFailed,
	"deleting":    DrPlanExecutionLifecycleStateDeleting,
	"deleted":     DrPlanExecutionLifecycleStateDeleted,
	"pausing":     DrPlanExecutionLifecycleStatePausing,
	"paused":      DrPlanExecutionLifecycleStatePaused,
	"resuming":    DrPlanExecutionLifecycleStateResuming,
}

// GetDrPlanExecutionLifecycleStateEnumValues Enumerates the set of values for DrPlanExecutionLifecycleStateEnum
func GetDrPlanExecutionLifecycleStateEnumValues() []DrPlanExecutionLifecycleStateEnum {
	values := make([]DrPlanExecutionLifecycleStateEnum, 0)
	for _, v := range mappingDrPlanExecutionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanExecutionLifecycleStateEnumStringValues Enumerates the set of values in String for DrPlanExecutionLifecycleStateEnum
func GetDrPlanExecutionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
		"FAILED",
		"DELETING",
		"DELETED",
		"PAUSING",
		"PAUSED",
		"RESUMING",
	}
}

// GetMappingDrPlanExecutionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanExecutionLifecycleStateEnum(val string) (DrPlanExecutionLifecycleStateEnum, bool) {
	enum, ok := mappingDrPlanExecutionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
