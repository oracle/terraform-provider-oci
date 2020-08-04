// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

// PrivateEndpointLifecycleStateEnum Enum with underlying type: string
type PrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateEndpointLifecycleStateEnum
const (
	PrivateEndpointLifecycleStateCreating PrivateEndpointLifecycleStateEnum = "CREATING"
	PrivateEndpointLifecycleStateActive   PrivateEndpointLifecycleStateEnum = "ACTIVE"
	PrivateEndpointLifecycleStateInactive PrivateEndpointLifecycleStateEnum = "INACTIVE"
	PrivateEndpointLifecycleStateUpdating PrivateEndpointLifecycleStateEnum = "UPDATING"
	PrivateEndpointLifecycleStateDeleting PrivateEndpointLifecycleStateEnum = "DELETING"
	PrivateEndpointLifecycleStateDeleted  PrivateEndpointLifecycleStateEnum = "DELETED"
	PrivateEndpointLifecycleStateFailed   PrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingPrivateEndpointLifecycleState = map[string]PrivateEndpointLifecycleStateEnum{
	"CREATING": PrivateEndpointLifecycleStateCreating,
	"ACTIVE":   PrivateEndpointLifecycleStateActive,
	"INACTIVE": PrivateEndpointLifecycleStateInactive,
	"UPDATING": PrivateEndpointLifecycleStateUpdating,
	"DELETING": PrivateEndpointLifecycleStateDeleting,
	"DELETED":  PrivateEndpointLifecycleStateDeleted,
	"FAILED":   PrivateEndpointLifecycleStateFailed,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleState {
		values = append(values, v)
	}
	return values
}
