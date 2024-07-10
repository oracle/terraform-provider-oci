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

// ComputeClusterContextLifecycleStateEnum Enum with underlying type: string
type ComputeClusterContextLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeClusterContextLifecycleStateEnum
const (
	ComputeClusterContextLifecycleStateCreating   ComputeClusterContextLifecycleStateEnum = "CREATING"
	ComputeClusterContextLifecycleStateActive     ComputeClusterContextLifecycleStateEnum = "ACTIVE"
	ComputeClusterContextLifecycleStateCancelling ComputeClusterContextLifecycleStateEnum = "CANCELLING"
	ComputeClusterContextLifecycleStateCancelled  ComputeClusterContextLifecycleStateEnum = "CANCELLED"
	ComputeClusterContextLifecycleStateFailed     ComputeClusterContextLifecycleStateEnum = "FAILED"
)

var mappingComputeClusterContextLifecycleStateEnum = map[string]ComputeClusterContextLifecycleStateEnum{
	"CREATING":   ComputeClusterContextLifecycleStateCreating,
	"ACTIVE":     ComputeClusterContextLifecycleStateActive,
	"CANCELLING": ComputeClusterContextLifecycleStateCancelling,
	"CANCELLED":  ComputeClusterContextLifecycleStateCancelled,
	"FAILED":     ComputeClusterContextLifecycleStateFailed,
}

var mappingComputeClusterContextLifecycleStateEnumLowerCase = map[string]ComputeClusterContextLifecycleStateEnum{
	"creating":   ComputeClusterContextLifecycleStateCreating,
	"active":     ComputeClusterContextLifecycleStateActive,
	"cancelling": ComputeClusterContextLifecycleStateCancelling,
	"cancelled":  ComputeClusterContextLifecycleStateCancelled,
	"failed":     ComputeClusterContextLifecycleStateFailed,
}

// GetComputeClusterContextLifecycleStateEnumValues Enumerates the set of values for ComputeClusterContextLifecycleStateEnum
func GetComputeClusterContextLifecycleStateEnumValues() []ComputeClusterContextLifecycleStateEnum {
	values := make([]ComputeClusterContextLifecycleStateEnum, 0)
	for _, v := range mappingComputeClusterContextLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterContextLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeClusterContextLifecycleStateEnum
func GetComputeClusterContextLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
	}
}

// GetMappingComputeClusterContextLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterContextLifecycleStateEnum(val string) (ComputeClusterContextLifecycleStateEnum, bool) {
	enum, ok := mappingComputeClusterContextLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
