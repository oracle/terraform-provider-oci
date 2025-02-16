// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

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
	LifecycleStateNeedsAttention LifecycleStateEnum = "NEEDS_ATTENTION"
	LifecycleStateDeleting       LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted        LifecycleStateEnum = "DELETED"
	LifecycleStateFailed         LifecycleStateEnum = "FAILED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING":        LifecycleStateCreating,
	"UPDATING":        LifecycleStateUpdating,
	"ACTIVE":          LifecycleStateActive,
	"INACTIVE":        LifecycleStateInactive,
	"NEEDS_ATTENTION": LifecycleStateNeedsAttention,
	"DELETING":        LifecycleStateDeleting,
	"DELETED":         LifecycleStateDeleted,
	"FAILED":          LifecycleStateFailed,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"creating":        LifecycleStateCreating,
	"updating":        LifecycleStateUpdating,
	"active":          LifecycleStateActive,
	"inactive":        LifecycleStateInactive,
	"needs_attention": LifecycleStateNeedsAttention,
	"deleting":        LifecycleStateDeleting,
	"deleted":         LifecycleStateDeleted,
	"failed":          LifecycleStateFailed,
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
		"NEEDS_ATTENTION",
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
