// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating LifecycleStateEnum = "CREATING"
	LifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive LifecycleStateEnum = "INACTIVE"
	LifecycleStateUpdating LifecycleStateEnum = "UPDATING"
	LifecycleStateDeleting LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted  LifecycleStateEnum = "DELETED"
	LifecycleStateFailed   LifecycleStateEnum = "FAILED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING": LifecycleStateCreating,
	"ACTIVE":   LifecycleStateActive,
	"INACTIVE": LifecycleStateInactive,
	"UPDATING": LifecycleStateUpdating,
	"DELETING": LifecycleStateDeleting,
	"DELETED":  LifecycleStateDeleted,
	"FAILED":   LifecycleStateFailed,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"creating": LifecycleStateCreating,
	"active":   LifecycleStateActive,
	"inactive": LifecycleStateInactive,
	"updating": LifecycleStateUpdating,
	"deleting": LifecycleStateDeleting,
	"deleted":  LifecycleStateDeleted,
	"failed":   LifecycleStateFailed,
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
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
