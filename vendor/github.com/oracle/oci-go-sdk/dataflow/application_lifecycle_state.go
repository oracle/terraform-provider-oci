// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

// ApplicationLifecycleStateEnum Enum with underlying type: string
type ApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationLifecycleStateEnum
const (
	ApplicationLifecycleStateActive   ApplicationLifecycleStateEnum = "ACTIVE"
	ApplicationLifecycleStateDeleted  ApplicationLifecycleStateEnum = "DELETED"
	ApplicationLifecycleStateInactive ApplicationLifecycleStateEnum = "INACTIVE"
)

var mappingApplicationLifecycleState = map[string]ApplicationLifecycleStateEnum{
	"ACTIVE":   ApplicationLifecycleStateActive,
	"DELETED":  ApplicationLifecycleStateDeleted,
	"INACTIVE": ApplicationLifecycleStateInactive,
}

// GetApplicationLifecycleStateEnumValues Enumerates the set of values for ApplicationLifecycleStateEnum
func GetApplicationLifecycleStateEnumValues() []ApplicationLifecycleStateEnum {
	values := make([]ApplicationLifecycleStateEnum, 0)
	for _, v := range mappingApplicationLifecycleState {
		values = append(values, v)
	}
	return values
}
