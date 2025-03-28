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

// CollectionLifecycleStatesEnum Enum with underlying type: string
type CollectionLifecycleStatesEnum string

// Set of constants representing the allowable values for CollectionLifecycleStatesEnum
const (
	CollectionLifecycleStatesCreating       CollectionLifecycleStatesEnum = "CREATING"
	CollectionLifecycleStatesUpdating       CollectionLifecycleStatesEnum = "UPDATING"
	CollectionLifecycleStatesActive         CollectionLifecycleStatesEnum = "ACTIVE"
	CollectionLifecycleStatesNeedsAttention CollectionLifecycleStatesEnum = "NEEDS_ATTENTION"
	CollectionLifecycleStatesDeleting       CollectionLifecycleStatesEnum = "DELETING"
	CollectionLifecycleStatesDeleted        CollectionLifecycleStatesEnum = "DELETED"
	CollectionLifecycleStatesFailed         CollectionLifecycleStatesEnum = "FAILED"
)

var mappingCollectionLifecycleStatesEnum = map[string]CollectionLifecycleStatesEnum{
	"CREATING":        CollectionLifecycleStatesCreating,
	"UPDATING":        CollectionLifecycleStatesUpdating,
	"ACTIVE":          CollectionLifecycleStatesActive,
	"NEEDS_ATTENTION": CollectionLifecycleStatesNeedsAttention,
	"DELETING":        CollectionLifecycleStatesDeleting,
	"DELETED":         CollectionLifecycleStatesDeleted,
	"FAILED":          CollectionLifecycleStatesFailed,
}

var mappingCollectionLifecycleStatesEnumLowerCase = map[string]CollectionLifecycleStatesEnum{
	"creating":        CollectionLifecycleStatesCreating,
	"updating":        CollectionLifecycleStatesUpdating,
	"active":          CollectionLifecycleStatesActive,
	"needs_attention": CollectionLifecycleStatesNeedsAttention,
	"deleting":        CollectionLifecycleStatesDeleting,
	"deleted":         CollectionLifecycleStatesDeleted,
	"failed":          CollectionLifecycleStatesFailed,
}

// GetCollectionLifecycleStatesEnumValues Enumerates the set of values for CollectionLifecycleStatesEnum
func GetCollectionLifecycleStatesEnumValues() []CollectionLifecycleStatesEnum {
	values := make([]CollectionLifecycleStatesEnum, 0)
	for _, v := range mappingCollectionLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectionLifecycleStatesEnumStringValues Enumerates the set of values in String for CollectionLifecycleStatesEnum
func GetCollectionLifecycleStatesEnumStringValues() []string {
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

// GetMappingCollectionLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectionLifecycleStatesEnum(val string) (CollectionLifecycleStatesEnum, bool) {
	enum, ok := mappingCollectionLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
