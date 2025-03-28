// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// RunLifecycleStateEnum Enum with underlying type: string
type RunLifecycleStateEnum string

// Set of constants representing the allowable values for RunLifecycleStateEnum
const (
	RunLifecycleStateAccepted   RunLifecycleStateEnum = "ACCEPTED"
	RunLifecycleStateInProgress RunLifecycleStateEnum = "IN_PROGRESS"
	RunLifecycleStateCanceling  RunLifecycleStateEnum = "CANCELING"
	RunLifecycleStateCanceled   RunLifecycleStateEnum = "CANCELED"
	RunLifecycleStateFailed     RunLifecycleStateEnum = "FAILED"
	RunLifecycleStateSucceeded  RunLifecycleStateEnum = "SUCCEEDED"
	RunLifecycleStateStopping   RunLifecycleStateEnum = "STOPPING"
	RunLifecycleStateStopped    RunLifecycleStateEnum = "STOPPED"
)

var mappingRunLifecycleStateEnum = map[string]RunLifecycleStateEnum{
	"ACCEPTED":    RunLifecycleStateAccepted,
	"IN_PROGRESS": RunLifecycleStateInProgress,
	"CANCELING":   RunLifecycleStateCanceling,
	"CANCELED":    RunLifecycleStateCanceled,
	"FAILED":      RunLifecycleStateFailed,
	"SUCCEEDED":   RunLifecycleStateSucceeded,
	"STOPPING":    RunLifecycleStateStopping,
	"STOPPED":     RunLifecycleStateStopped,
}

var mappingRunLifecycleStateEnumLowerCase = map[string]RunLifecycleStateEnum{
	"accepted":    RunLifecycleStateAccepted,
	"in_progress": RunLifecycleStateInProgress,
	"canceling":   RunLifecycleStateCanceling,
	"canceled":    RunLifecycleStateCanceled,
	"failed":      RunLifecycleStateFailed,
	"succeeded":   RunLifecycleStateSucceeded,
	"stopping":    RunLifecycleStateStopping,
	"stopped":     RunLifecycleStateStopped,
}

// GetRunLifecycleStateEnumValues Enumerates the set of values for RunLifecycleStateEnum
func GetRunLifecycleStateEnumValues() []RunLifecycleStateEnum {
	values := make([]RunLifecycleStateEnum, 0)
	for _, v := range mappingRunLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRunLifecycleStateEnumStringValues Enumerates the set of values in String for RunLifecycleStateEnum
func GetRunLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"FAILED",
		"SUCCEEDED",
		"STOPPING",
		"STOPPED",
	}
}

// GetMappingRunLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunLifecycleStateEnum(val string) (RunLifecycleStateEnum, bool) {
	enum, ok := mappingRunLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
