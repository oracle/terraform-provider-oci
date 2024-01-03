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

// PoolLifecycleStateEnum Enum with underlying type: string
type PoolLifecycleStateEnum string

// Set of constants representing the allowable values for PoolLifecycleStateEnum
const (
	PoolLifecycleStateAccepted  PoolLifecycleStateEnum = "ACCEPTED"
	PoolLifecycleStateScheduled PoolLifecycleStateEnum = "SCHEDULED"
	PoolLifecycleStateCreating  PoolLifecycleStateEnum = "CREATING"
	PoolLifecycleStateActive    PoolLifecycleStateEnum = "ACTIVE"
	PoolLifecycleStateStopping  PoolLifecycleStateEnum = "STOPPING"
	PoolLifecycleStateStopped   PoolLifecycleStateEnum = "STOPPED"
	PoolLifecycleStateUpdating  PoolLifecycleStateEnum = "UPDATING"
	PoolLifecycleStateDeleting  PoolLifecycleStateEnum = "DELETING"
	PoolLifecycleStateDeleted   PoolLifecycleStateEnum = "DELETED"
	PoolLifecycleStateFailed    PoolLifecycleStateEnum = "FAILED"
)

var mappingPoolLifecycleStateEnum = map[string]PoolLifecycleStateEnum{
	"ACCEPTED":  PoolLifecycleStateAccepted,
	"SCHEDULED": PoolLifecycleStateScheduled,
	"CREATING":  PoolLifecycleStateCreating,
	"ACTIVE":    PoolLifecycleStateActive,
	"STOPPING":  PoolLifecycleStateStopping,
	"STOPPED":   PoolLifecycleStateStopped,
	"UPDATING":  PoolLifecycleStateUpdating,
	"DELETING":  PoolLifecycleStateDeleting,
	"DELETED":   PoolLifecycleStateDeleted,
	"FAILED":    PoolLifecycleStateFailed,
}

var mappingPoolLifecycleStateEnumLowerCase = map[string]PoolLifecycleStateEnum{
	"accepted":  PoolLifecycleStateAccepted,
	"scheduled": PoolLifecycleStateScheduled,
	"creating":  PoolLifecycleStateCreating,
	"active":    PoolLifecycleStateActive,
	"stopping":  PoolLifecycleStateStopping,
	"stopped":   PoolLifecycleStateStopped,
	"updating":  PoolLifecycleStateUpdating,
	"deleting":  PoolLifecycleStateDeleting,
	"deleted":   PoolLifecycleStateDeleted,
	"failed":    PoolLifecycleStateFailed,
}

// GetPoolLifecycleStateEnumValues Enumerates the set of values for PoolLifecycleStateEnum
func GetPoolLifecycleStateEnumValues() []PoolLifecycleStateEnum {
	values := make([]PoolLifecycleStateEnum, 0)
	for _, v := range mappingPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPoolLifecycleStateEnumStringValues Enumerates the set of values in String for PoolLifecycleStateEnum
func GetPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"SCHEDULED",
		"CREATING",
		"ACTIVE",
		"STOPPING",
		"STOPPED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPoolLifecycleStateEnum(val string) (PoolLifecycleStateEnum, bool) {
	enum, ok := mappingPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
