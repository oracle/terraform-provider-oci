// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// DiscoveryLifecycleStateEnum Enum with underlying type: string
type DiscoveryLifecycleStateEnum string

// Set of constants representing the allowable values for DiscoveryLifecycleStateEnum
const (
	DiscoveryLifecycleStateCreating DiscoveryLifecycleStateEnum = "CREATING"
	DiscoveryLifecycleStateActive   DiscoveryLifecycleStateEnum = "ACTIVE"
	DiscoveryLifecycleStateUpdating DiscoveryLifecycleStateEnum = "UPDATING"
	DiscoveryLifecycleStateDeleting DiscoveryLifecycleStateEnum = "DELETING"
	DiscoveryLifecycleStateDeleted  DiscoveryLifecycleStateEnum = "DELETED"
	DiscoveryLifecycleStateFailed   DiscoveryLifecycleStateEnum = "FAILED"
)

var mappingDiscoveryLifecycleStateEnum = map[string]DiscoveryLifecycleStateEnum{
	"CREATING": DiscoveryLifecycleStateCreating,
	"ACTIVE":   DiscoveryLifecycleStateActive,
	"UPDATING": DiscoveryLifecycleStateUpdating,
	"DELETING": DiscoveryLifecycleStateDeleting,
	"DELETED":  DiscoveryLifecycleStateDeleted,
	"FAILED":   DiscoveryLifecycleStateFailed,
}

var mappingDiscoveryLifecycleStateEnumLowerCase = map[string]DiscoveryLifecycleStateEnum{
	"creating": DiscoveryLifecycleStateCreating,
	"active":   DiscoveryLifecycleStateActive,
	"updating": DiscoveryLifecycleStateUpdating,
	"deleting": DiscoveryLifecycleStateDeleting,
	"deleted":  DiscoveryLifecycleStateDeleted,
	"failed":   DiscoveryLifecycleStateFailed,
}

// GetDiscoveryLifecycleStateEnumValues Enumerates the set of values for DiscoveryLifecycleStateEnum
func GetDiscoveryLifecycleStateEnumValues() []DiscoveryLifecycleStateEnum {
	values := make([]DiscoveryLifecycleStateEnum, 0)
	for _, v := range mappingDiscoveryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryLifecycleStateEnumStringValues Enumerates the set of values in String for DiscoveryLifecycleStateEnum
func GetDiscoveryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDiscoveryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryLifecycleStateEnum(val string) (DiscoveryLifecycleStateEnum, bool) {
	enum, ok := mappingDiscoveryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
