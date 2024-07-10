// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ComputeClusterCommandLifecycleStateEnum Enum with underlying type: string
type ComputeClusterCommandLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeClusterCommandLifecycleStateEnum
const (
	ComputeClusterCommandLifecycleStateCancelled  ComputeClusterCommandLifecycleStateEnum = "CANCELLED"
	ComputeClusterCommandLifecycleStateCancelling ComputeClusterCommandLifecycleStateEnum = "CANCELLING"
	ComputeClusterCommandLifecycleStateFailed     ComputeClusterCommandLifecycleStateEnum = "FAILED"
	ComputeClusterCommandLifecycleStateSucceeded  ComputeClusterCommandLifecycleStateEnum = "SUCCEEDED"
	ComputeClusterCommandLifecycleStateQueued     ComputeClusterCommandLifecycleStateEnum = "QUEUED"
	ComputeClusterCommandLifecycleStateRunning    ComputeClusterCommandLifecycleStateEnum = "RUNNING"
)

var mappingComputeClusterCommandLifecycleStateEnum = map[string]ComputeClusterCommandLifecycleStateEnum{
	"CANCELLED":  ComputeClusterCommandLifecycleStateCancelled,
	"CANCELLING": ComputeClusterCommandLifecycleStateCancelling,
	"FAILED":     ComputeClusterCommandLifecycleStateFailed,
	"SUCCEEDED":  ComputeClusterCommandLifecycleStateSucceeded,
	"QUEUED":     ComputeClusterCommandLifecycleStateQueued,
	"RUNNING":    ComputeClusterCommandLifecycleStateRunning,
}

var mappingComputeClusterCommandLifecycleStateEnumLowerCase = map[string]ComputeClusterCommandLifecycleStateEnum{
	"cancelled":  ComputeClusterCommandLifecycleStateCancelled,
	"cancelling": ComputeClusterCommandLifecycleStateCancelling,
	"failed":     ComputeClusterCommandLifecycleStateFailed,
	"succeeded":  ComputeClusterCommandLifecycleStateSucceeded,
	"queued":     ComputeClusterCommandLifecycleStateQueued,
	"running":    ComputeClusterCommandLifecycleStateRunning,
}

// GetComputeClusterCommandLifecycleStateEnumValues Enumerates the set of values for ComputeClusterCommandLifecycleStateEnum
func GetComputeClusterCommandLifecycleStateEnumValues() []ComputeClusterCommandLifecycleStateEnum {
	values := make([]ComputeClusterCommandLifecycleStateEnum, 0)
	for _, v := range mappingComputeClusterCommandLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterCommandLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeClusterCommandLifecycleStateEnum
func GetComputeClusterCommandLifecycleStateEnumStringValues() []string {
	return []string{
		"CANCELLED",
		"CANCELLING",
		"FAILED",
		"SUCCEEDED",
		"QUEUED",
		"RUNNING",
	}
}

// GetMappingComputeClusterCommandLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterCommandLifecycleStateEnum(val string) (ComputeClusterCommandLifecycleStateEnum, bool) {
	enum, ok := mappingComputeClusterCommandLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
