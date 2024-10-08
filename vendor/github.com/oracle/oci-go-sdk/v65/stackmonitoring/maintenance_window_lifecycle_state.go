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

// MaintenanceWindowLifecycleStateEnum Enum with underlying type: string
type MaintenanceWindowLifecycleStateEnum string

// Set of constants representing the allowable values for MaintenanceWindowLifecycleStateEnum
const (
	MaintenanceWindowLifecycleStateCreating       MaintenanceWindowLifecycleStateEnum = "CREATING"
	MaintenanceWindowLifecycleStateUpdating       MaintenanceWindowLifecycleStateEnum = "UPDATING"
	MaintenanceWindowLifecycleStateInactive       MaintenanceWindowLifecycleStateEnum = "INACTIVE"
	MaintenanceWindowLifecycleStateActive         MaintenanceWindowLifecycleStateEnum = "ACTIVE"
	MaintenanceWindowLifecycleStateDeleting       MaintenanceWindowLifecycleStateEnum = "DELETING"
	MaintenanceWindowLifecycleStateDeleted        MaintenanceWindowLifecycleStateEnum = "DELETED"
	MaintenanceWindowLifecycleStateFailed         MaintenanceWindowLifecycleStateEnum = "FAILED"
	MaintenanceWindowLifecycleStateNeedsAttention MaintenanceWindowLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingMaintenanceWindowLifecycleStateEnum = map[string]MaintenanceWindowLifecycleStateEnum{
	"CREATING":        MaintenanceWindowLifecycleStateCreating,
	"UPDATING":        MaintenanceWindowLifecycleStateUpdating,
	"INACTIVE":        MaintenanceWindowLifecycleStateInactive,
	"ACTIVE":          MaintenanceWindowLifecycleStateActive,
	"DELETING":        MaintenanceWindowLifecycleStateDeleting,
	"DELETED":         MaintenanceWindowLifecycleStateDeleted,
	"FAILED":          MaintenanceWindowLifecycleStateFailed,
	"NEEDS_ATTENTION": MaintenanceWindowLifecycleStateNeedsAttention,
}

var mappingMaintenanceWindowLifecycleStateEnumLowerCase = map[string]MaintenanceWindowLifecycleStateEnum{
	"creating":        MaintenanceWindowLifecycleStateCreating,
	"updating":        MaintenanceWindowLifecycleStateUpdating,
	"inactive":        MaintenanceWindowLifecycleStateInactive,
	"active":          MaintenanceWindowLifecycleStateActive,
	"deleting":        MaintenanceWindowLifecycleStateDeleting,
	"deleted":         MaintenanceWindowLifecycleStateDeleted,
	"failed":          MaintenanceWindowLifecycleStateFailed,
	"needs_attention": MaintenanceWindowLifecycleStateNeedsAttention,
}

// GetMaintenanceWindowLifecycleStateEnumValues Enumerates the set of values for MaintenanceWindowLifecycleStateEnum
func GetMaintenanceWindowLifecycleStateEnumValues() []MaintenanceWindowLifecycleStateEnum {
	values := make([]MaintenanceWindowLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceWindowLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowLifecycleStateEnumStringValues Enumerates the set of values in String for MaintenanceWindowLifecycleStateEnum
func GetMaintenanceWindowLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"INACTIVE",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingMaintenanceWindowLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowLifecycleStateEnum(val string) (MaintenanceWindowLifecycleStateEnum, bool) {
	enum, ok := mappingMaintenanceWindowLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
