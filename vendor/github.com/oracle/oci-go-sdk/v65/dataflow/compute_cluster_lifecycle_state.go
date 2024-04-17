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

// ComputeClusterLifecycleStateEnum Enum with underlying type: string
type ComputeClusterLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeClusterLifecycleStateEnum
const (
	ComputeClusterLifecycleStateAccepted ComputeClusterLifecycleStateEnum = "ACCEPTED"
	ComputeClusterLifecycleStateCreating ComputeClusterLifecycleStateEnum = "CREATING"
	ComputeClusterLifecycleStateActive   ComputeClusterLifecycleStateEnum = "ACTIVE"
	ComputeClusterLifecycleStateDeleting ComputeClusterLifecycleStateEnum = "DELETING"
	ComputeClusterLifecycleStateDeleted  ComputeClusterLifecycleStateEnum = "DELETED"
	ComputeClusterLifecycleStateFailed   ComputeClusterLifecycleStateEnum = "FAILED"
	ComputeClusterLifecycleStateStopping ComputeClusterLifecycleStateEnum = "STOPPING"
	ComputeClusterLifecycleStateStopped  ComputeClusterLifecycleStateEnum = "STOPPED"
	ComputeClusterLifecycleStateUpdating ComputeClusterLifecycleStateEnum = "UPDATING"
)

var mappingComputeClusterLifecycleStateEnum = map[string]ComputeClusterLifecycleStateEnum{
	"ACCEPTED": ComputeClusterLifecycleStateAccepted,
	"CREATING": ComputeClusterLifecycleStateCreating,
	"ACTIVE":   ComputeClusterLifecycleStateActive,
	"DELETING": ComputeClusterLifecycleStateDeleting,
	"DELETED":  ComputeClusterLifecycleStateDeleted,
	"FAILED":   ComputeClusterLifecycleStateFailed,
	"STOPPING": ComputeClusterLifecycleStateStopping,
	"STOPPED":  ComputeClusterLifecycleStateStopped,
	"UPDATING": ComputeClusterLifecycleStateUpdating,
}

var mappingComputeClusterLifecycleStateEnumLowerCase = map[string]ComputeClusterLifecycleStateEnum{
	"accepted": ComputeClusterLifecycleStateAccepted,
	"creating": ComputeClusterLifecycleStateCreating,
	"active":   ComputeClusterLifecycleStateActive,
	"deleting": ComputeClusterLifecycleStateDeleting,
	"deleted":  ComputeClusterLifecycleStateDeleted,
	"failed":   ComputeClusterLifecycleStateFailed,
	"stopping": ComputeClusterLifecycleStateStopping,
	"stopped":  ComputeClusterLifecycleStateStopped,
	"updating": ComputeClusterLifecycleStateUpdating,
}

// GetComputeClusterLifecycleStateEnumValues Enumerates the set of values for ComputeClusterLifecycleStateEnum
func GetComputeClusterLifecycleStateEnumValues() []ComputeClusterLifecycleStateEnum {
	values := make([]ComputeClusterLifecycleStateEnum, 0)
	for _, v := range mappingComputeClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeClusterLifecycleStateEnum
func GetComputeClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"STOPPING",
		"STOPPED",
		"UPDATING",
	}
}

// GetMappingComputeClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterLifecycleStateEnum(val string) (ComputeClusterLifecycleStateEnum, bool) {
	enum, ok := mappingComputeClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
