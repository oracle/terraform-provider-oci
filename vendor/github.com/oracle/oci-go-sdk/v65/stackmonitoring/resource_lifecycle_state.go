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

// ResourceLifecycleStateEnum Enum with underlying type: string
type ResourceLifecycleStateEnum string

// Set of constants representing the allowable values for ResourceLifecycleStateEnum
const (
	ResourceLifecycleStateCreating ResourceLifecycleStateEnum = "CREATING"
	ResourceLifecycleStateUpdating ResourceLifecycleStateEnum = "UPDATING"
	ResourceLifecycleStateActive   ResourceLifecycleStateEnum = "ACTIVE"
	ResourceLifecycleStateDeleting ResourceLifecycleStateEnum = "DELETING"
	ResourceLifecycleStateDeleted  ResourceLifecycleStateEnum = "DELETED"
	ResourceLifecycleStateFailed   ResourceLifecycleStateEnum = "FAILED"
)

var mappingResourceLifecycleStateEnum = map[string]ResourceLifecycleStateEnum{
	"CREATING": ResourceLifecycleStateCreating,
	"UPDATING": ResourceLifecycleStateUpdating,
	"ACTIVE":   ResourceLifecycleStateActive,
	"DELETING": ResourceLifecycleStateDeleting,
	"DELETED":  ResourceLifecycleStateDeleted,
	"FAILED":   ResourceLifecycleStateFailed,
}

var mappingResourceLifecycleStateEnumLowerCase = map[string]ResourceLifecycleStateEnum{
	"creating": ResourceLifecycleStateCreating,
	"updating": ResourceLifecycleStateUpdating,
	"active":   ResourceLifecycleStateActive,
	"deleting": ResourceLifecycleStateDeleting,
	"deleted":  ResourceLifecycleStateDeleted,
	"failed":   ResourceLifecycleStateFailed,
}

// GetResourceLifecycleStateEnumValues Enumerates the set of values for ResourceLifecycleStateEnum
func GetResourceLifecycleStateEnumValues() []ResourceLifecycleStateEnum {
	values := make([]ResourceLifecycleStateEnum, 0)
	for _, v := range mappingResourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceLifecycleStateEnumStringValues Enumerates the set of values in String for ResourceLifecycleStateEnum
func GetResourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingResourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceLifecycleStateEnum(val string) (ResourceLifecycleStateEnum, bool) {
	enum, ok := mappingResourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
