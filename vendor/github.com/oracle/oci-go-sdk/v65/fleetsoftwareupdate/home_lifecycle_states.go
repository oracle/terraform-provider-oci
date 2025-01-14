// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// HomeLifecycleStatesEnum Enum with underlying type: string
type HomeLifecycleStatesEnum string

// Set of constants representing the allowable values for HomeLifecycleStatesEnum
const (
	HomeLifecycleStatesCreating       HomeLifecycleStatesEnum = "CREATING"
	HomeLifecycleStatesUpdating       HomeLifecycleStatesEnum = "UPDATING"
	HomeLifecycleStatesActive         HomeLifecycleStatesEnum = "ACTIVE"
	HomeLifecycleStatesNeedsAttention HomeLifecycleStatesEnum = "NEEDS_ATTENTION"
	HomeLifecycleStatesDeleting       HomeLifecycleStatesEnum = "DELETING"
	HomeLifecycleStatesDeleted        HomeLifecycleStatesEnum = "DELETED"
	HomeLifecycleStatesFailed         HomeLifecycleStatesEnum = "FAILED"
)

var mappingHomeLifecycleStatesEnum = map[string]HomeLifecycleStatesEnum{
	"CREATING":        HomeLifecycleStatesCreating,
	"UPDATING":        HomeLifecycleStatesUpdating,
	"ACTIVE":          HomeLifecycleStatesActive,
	"NEEDS_ATTENTION": HomeLifecycleStatesNeedsAttention,
	"DELETING":        HomeLifecycleStatesDeleting,
	"DELETED":         HomeLifecycleStatesDeleted,
	"FAILED":          HomeLifecycleStatesFailed,
}

var mappingHomeLifecycleStatesEnumLowerCase = map[string]HomeLifecycleStatesEnum{
	"creating":        HomeLifecycleStatesCreating,
	"updating":        HomeLifecycleStatesUpdating,
	"active":          HomeLifecycleStatesActive,
	"needs_attention": HomeLifecycleStatesNeedsAttention,
	"deleting":        HomeLifecycleStatesDeleting,
	"deleted":         HomeLifecycleStatesDeleted,
	"failed":          HomeLifecycleStatesFailed,
}

// GetHomeLifecycleStatesEnumValues Enumerates the set of values for HomeLifecycleStatesEnum
func GetHomeLifecycleStatesEnumValues() []HomeLifecycleStatesEnum {
	values := make([]HomeLifecycleStatesEnum, 0)
	for _, v := range mappingHomeLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetHomeLifecycleStatesEnumStringValues Enumerates the set of values in String for HomeLifecycleStatesEnum
func GetHomeLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingHomeLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHomeLifecycleStatesEnum(val string) (HomeLifecycleStatesEnum, bool) {
	enum, ok := mappingHomeLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
