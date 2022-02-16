// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// AnalyticsInstanceLifecycleStateEnum Enum with underlying type: string
type AnalyticsInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for AnalyticsInstanceLifecycleStateEnum
const (
	AnalyticsInstanceLifecycleStateActive   AnalyticsInstanceLifecycleStateEnum = "ACTIVE"
	AnalyticsInstanceLifecycleStateCreating AnalyticsInstanceLifecycleStateEnum = "CREATING"
	AnalyticsInstanceLifecycleStateDeleted  AnalyticsInstanceLifecycleStateEnum = "DELETED"
	AnalyticsInstanceLifecycleStateDeleting AnalyticsInstanceLifecycleStateEnum = "DELETING"
	AnalyticsInstanceLifecycleStateFailed   AnalyticsInstanceLifecycleStateEnum = "FAILED"
	AnalyticsInstanceLifecycleStateInactive AnalyticsInstanceLifecycleStateEnum = "INACTIVE"
	AnalyticsInstanceLifecycleStateUpdating AnalyticsInstanceLifecycleStateEnum = "UPDATING"
)

var mappingAnalyticsInstanceLifecycleStateEnum = map[string]AnalyticsInstanceLifecycleStateEnum{
	"ACTIVE":   AnalyticsInstanceLifecycleStateActive,
	"CREATING": AnalyticsInstanceLifecycleStateCreating,
	"DELETED":  AnalyticsInstanceLifecycleStateDeleted,
	"DELETING": AnalyticsInstanceLifecycleStateDeleting,
	"FAILED":   AnalyticsInstanceLifecycleStateFailed,
	"INACTIVE": AnalyticsInstanceLifecycleStateInactive,
	"UPDATING": AnalyticsInstanceLifecycleStateUpdating,
}

// GetAnalyticsInstanceLifecycleStateEnumValues Enumerates the set of values for AnalyticsInstanceLifecycleStateEnum
func GetAnalyticsInstanceLifecycleStateEnumValues() []AnalyticsInstanceLifecycleStateEnum {
	values := make([]AnalyticsInstanceLifecycleStateEnum, 0)
	for _, v := range mappingAnalyticsInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAnalyticsInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for AnalyticsInstanceLifecycleStateEnum
func GetAnalyticsInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"INACTIVE",
		"UPDATING",
	}
}

// GetMappingAnalyticsInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnalyticsInstanceLifecycleStateEnum(val string) (AnalyticsInstanceLifecycleStateEnum, bool) {
	mappingAnalyticsInstanceLifecycleStateEnumIgnoreCase := make(map[string]AnalyticsInstanceLifecycleStateEnum)
	for k, v := range mappingAnalyticsInstanceLifecycleStateEnum {
		mappingAnalyticsInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAnalyticsInstanceLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
