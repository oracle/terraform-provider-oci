// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ImageLifecycleStatesEnum Enum with underlying type: string
type ImageLifecycleStatesEnum string

// Set of constants representing the allowable values for ImageLifecycleStatesEnum
const (
	ImageLifecycleStatesCreating       ImageLifecycleStatesEnum = "CREATING"
	ImageLifecycleStatesUpdating       ImageLifecycleStatesEnum = "UPDATING"
	ImageLifecycleStatesActive         ImageLifecycleStatesEnum = "ACTIVE"
	ImageLifecycleStatesNeedsAttention ImageLifecycleStatesEnum = "NEEDS_ATTENTION"
	ImageLifecycleStatesDeleting       ImageLifecycleStatesEnum = "DELETING"
	ImageLifecycleStatesDeleted        ImageLifecycleStatesEnum = "DELETED"
	ImageLifecycleStatesFailed         ImageLifecycleStatesEnum = "FAILED"
)

var mappingImageLifecycleStatesEnum = map[string]ImageLifecycleStatesEnum{
	"CREATING":        ImageLifecycleStatesCreating,
	"UPDATING":        ImageLifecycleStatesUpdating,
	"ACTIVE":          ImageLifecycleStatesActive,
	"NEEDS_ATTENTION": ImageLifecycleStatesNeedsAttention,
	"DELETING":        ImageLifecycleStatesDeleting,
	"DELETED":         ImageLifecycleStatesDeleted,
	"FAILED":          ImageLifecycleStatesFailed,
}

var mappingImageLifecycleStatesEnumLowerCase = map[string]ImageLifecycleStatesEnum{
	"creating":        ImageLifecycleStatesCreating,
	"updating":        ImageLifecycleStatesUpdating,
	"active":          ImageLifecycleStatesActive,
	"needs_attention": ImageLifecycleStatesNeedsAttention,
	"deleting":        ImageLifecycleStatesDeleting,
	"deleted":         ImageLifecycleStatesDeleted,
	"failed":          ImageLifecycleStatesFailed,
}

// GetImageLifecycleStatesEnumValues Enumerates the set of values for ImageLifecycleStatesEnum
func GetImageLifecycleStatesEnumValues() []ImageLifecycleStatesEnum {
	values := make([]ImageLifecycleStatesEnum, 0)
	for _, v := range mappingImageLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetImageLifecycleStatesEnumStringValues Enumerates the set of values in String for ImageLifecycleStatesEnum
func GetImageLifecycleStatesEnumStringValues() []string {
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

// GetMappingImageLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageLifecycleStatesEnum(val string) (ImageLifecycleStatesEnum, bool) {
	enum, ok := mappingImageLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
