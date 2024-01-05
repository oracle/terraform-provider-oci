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

var mappingPrivateEndpointLifecycleStateEnum = map[string]PrivateEndpointLifecycleStateEnum{
	"CREATING": PrivateEndpointLifecycleStateCreating,
	"ACTIVE":   PrivateEndpointLifecycleStateActive,
	"INACTIVE": PrivateEndpointLifecycleStateInactive,
	"UPDATING": PrivateEndpointLifecycleStateUpdating,
	"DELETING": PrivateEndpointLifecycleStateDeleting,
	"DELETED":  PrivateEndpointLifecycleStateDeleted,
	"FAILED":   PrivateEndpointLifecycleStateFailed,
}

var mappingPrivateEndpointLifecycleStateEnumLowerCase = map[string]PrivateEndpointLifecycleStateEnum{
	"creating": PrivateEndpointLifecycleStateCreating,
	"active":   PrivateEndpointLifecycleStateActive,
	"inactive": PrivateEndpointLifecycleStateInactive,
	"updating": PrivateEndpointLifecycleStateUpdating,
	"deleting": PrivateEndpointLifecycleStateDeleting,
	"deleted":  PrivateEndpointLifecycleStateDeleted,
	"failed":   PrivateEndpointLifecycleStateFailed,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateEndpointLifecycleStateEnum(val string) (PrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
