// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// ResourceTypeLifecycleStateEnum Enum with underlying type: string
type ResourceTypeLifecycleStateEnum string

// Set of constants representing the allowable values for ResourceTypeLifecycleStateEnum
const (
	ResourceTypeLifecycleStateCreating ResourceTypeLifecycleStateEnum = "CREATING"
	ResourceTypeLifecycleStateUpdating ResourceTypeLifecycleStateEnum = "UPDATING"
	ResourceTypeLifecycleStateActive   ResourceTypeLifecycleStateEnum = "ACTIVE"
	ResourceTypeLifecycleStateDeleting ResourceTypeLifecycleStateEnum = "DELETING"
	ResourceTypeLifecycleStateDeleted  ResourceTypeLifecycleStateEnum = "DELETED"
	ResourceTypeLifecycleStateFailed   ResourceTypeLifecycleStateEnum = "FAILED"
)

var mappingResourceTypeLifecycleStateEnum = map[string]ResourceTypeLifecycleStateEnum{
	"CREATING": ResourceTypeLifecycleStateCreating,
	"UPDATING": ResourceTypeLifecycleStateUpdating,
	"ACTIVE":   ResourceTypeLifecycleStateActive,
	"DELETING": ResourceTypeLifecycleStateDeleting,
	"DELETED":  ResourceTypeLifecycleStateDeleted,
	"FAILED":   ResourceTypeLifecycleStateFailed,
}

var mappingResourceTypeLifecycleStateEnumLowerCase = map[string]ResourceTypeLifecycleStateEnum{
	"creating": ResourceTypeLifecycleStateCreating,
	"updating": ResourceTypeLifecycleStateUpdating,
	"active":   ResourceTypeLifecycleStateActive,
	"deleting": ResourceTypeLifecycleStateDeleting,
	"deleted":  ResourceTypeLifecycleStateDeleted,
	"failed":   ResourceTypeLifecycleStateFailed,
}

// GetResourceTypeLifecycleStateEnumValues Enumerates the set of values for ResourceTypeLifecycleStateEnum
func GetResourceTypeLifecycleStateEnumValues() []ResourceTypeLifecycleStateEnum {
	values := make([]ResourceTypeLifecycleStateEnum, 0)
	for _, v := range mappingResourceTypeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeLifecycleStateEnumStringValues Enumerates the set of values in String for ResourceTypeLifecycleStateEnum
func GetResourceTypeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingResourceTypeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeLifecycleStateEnum(val string) (ResourceTypeLifecycleStateEnum, bool) {
	enum, ok := mappingResourceTypeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
