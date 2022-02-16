// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// MaskingLifecycleStateEnum Enum with underlying type: string
type MaskingLifecycleStateEnum string

// Set of constants representing the allowable values for MaskingLifecycleStateEnum
const (
	MaskingLifecycleStateCreating       MaskingLifecycleStateEnum = "CREATING"
	MaskingLifecycleStateActive         MaskingLifecycleStateEnum = "ACTIVE"
	MaskingLifecycleStateUpdating       MaskingLifecycleStateEnum = "UPDATING"
	MaskingLifecycleStateDeleting       MaskingLifecycleStateEnum = "DELETING"
	MaskingLifecycleStateDeleted        MaskingLifecycleStateEnum = "DELETED"
	MaskingLifecycleStateNeedsAttention MaskingLifecycleStateEnum = "NEEDS_ATTENTION"
	MaskingLifecycleStateFailed         MaskingLifecycleStateEnum = "FAILED"
)

var mappingMaskingLifecycleStateEnum = map[string]MaskingLifecycleStateEnum{
	"CREATING":        MaskingLifecycleStateCreating,
	"ACTIVE":          MaskingLifecycleStateActive,
	"UPDATING":        MaskingLifecycleStateUpdating,
	"DELETING":        MaskingLifecycleStateDeleting,
	"DELETED":         MaskingLifecycleStateDeleted,
	"NEEDS_ATTENTION": MaskingLifecycleStateNeedsAttention,
	"FAILED":          MaskingLifecycleStateFailed,
}

// GetMaskingLifecycleStateEnumValues Enumerates the set of values for MaskingLifecycleStateEnum
func GetMaskingLifecycleStateEnumValues() []MaskingLifecycleStateEnum {
	values := make([]MaskingLifecycleStateEnum, 0)
	for _, v := range mappingMaskingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingLifecycleStateEnumStringValues Enumerates the set of values in String for MaskingLifecycleStateEnum
func GetMaskingLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingMaskingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingLifecycleStateEnum(val string) (MaskingLifecycleStateEnum, bool) {
	mappingMaskingLifecycleStateEnumIgnoreCase := make(map[string]MaskingLifecycleStateEnum)
	for k, v := range mappingMaskingLifecycleStateEnum {
		mappingMaskingLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaskingLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
