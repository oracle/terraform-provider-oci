// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

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
)

var mappingRunLifecycleState = map[string]RunLifecycleStateEnum{
	"ACCEPTED":    RunLifecycleStateAccepted,
	"IN_PROGRESS": RunLifecycleStateInProgress,
	"CANCELING":   RunLifecycleStateCanceling,
	"CANCELED":    RunLifecycleStateCanceled,
	"FAILED":      RunLifecycleStateFailed,
	"SUCCEEDED":   RunLifecycleStateSucceeded,
}

// GetRunLifecycleStateEnumValues Enumerates the set of values for RunLifecycleStateEnum
func GetRunLifecycleStateEnumValues() []RunLifecycleStateEnum {
	values := make([]RunLifecycleStateEnum, 0)
	for _, v := range mappingRunLifecycleState {
		values = append(values, v)
	}
	return values
}
