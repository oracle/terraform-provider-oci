// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating       LifecycleStateEnum = "CREATING"
	LifecycleStateUpdating       LifecycleStateEnum = "UPDATING"
	LifecycleStateActive         LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive       LifecycleStateEnum = "INACTIVE"
	LifecycleStateDeleting       LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted        LifecycleStateEnum = "DELETED"
	LifecycleStateFailed         LifecycleStateEnum = "FAILED"
	LifecycleStateNeedsAttention LifecycleStateEnum = "NEEDS_ATTENTION"
	LifecycleStateInProgress     LifecycleStateEnum = "IN_PROGRESS"
	LifecycleStateCanceling      LifecycleStateEnum = "CANCELING"
	LifecycleStateCanceled       LifecycleStateEnum = "CANCELED"
	LifecycleStateSucceeded      LifecycleStateEnum = "SUCCEEDED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING":        LifecycleStateCreating,
	"UPDATING":        LifecycleStateUpdating,
	"ACTIVE":          LifecycleStateActive,
	"INACTIVE":        LifecycleStateInactive,
	"DELETING":        LifecycleStateDeleting,
	"DELETED":         LifecycleStateDeleted,
	"FAILED":          LifecycleStateFailed,
	"NEEDS_ATTENTION": LifecycleStateNeedsAttention,
	"IN_PROGRESS":     LifecycleStateInProgress,
	"CANCELING":       LifecycleStateCanceling,
	"CANCELED":        LifecycleStateCanceled,
	"SUCCEEDED":       LifecycleStateSucceeded,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleStateEnum
func GetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"IN_PROGRESS",
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	mappingLifecycleStateEnumIgnoreCase := make(map[string]LifecycleStateEnum)
	for k, v := range mappingLifecycleStateEnum {
		mappingLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
