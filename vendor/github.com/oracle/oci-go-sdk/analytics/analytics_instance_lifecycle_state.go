// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

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

var mappingAnalyticsInstanceLifecycleState = map[string]AnalyticsInstanceLifecycleStateEnum{
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
	for _, v := range mappingAnalyticsInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
