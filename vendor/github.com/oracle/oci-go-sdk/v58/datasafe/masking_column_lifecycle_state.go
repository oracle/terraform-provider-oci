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

// MaskingColumnLifecycleStateEnum Enum with underlying type: string
type MaskingColumnLifecycleStateEnum string

// Set of constants representing the allowable values for MaskingColumnLifecycleStateEnum
const (
	MaskingColumnLifecycleStateCreating       MaskingColumnLifecycleStateEnum = "CREATING"
	MaskingColumnLifecycleStateActive         MaskingColumnLifecycleStateEnum = "ACTIVE"
	MaskingColumnLifecycleStateUpdating       MaskingColumnLifecycleStateEnum = "UPDATING"
	MaskingColumnLifecycleStateDeleting       MaskingColumnLifecycleStateEnum = "DELETING"
	MaskingColumnLifecycleStateNeedsAttention MaskingColumnLifecycleStateEnum = "NEEDS_ATTENTION"
	MaskingColumnLifecycleStateFailed         MaskingColumnLifecycleStateEnum = "FAILED"
)

var mappingMaskingColumnLifecycleStateEnum = map[string]MaskingColumnLifecycleStateEnum{
	"CREATING":        MaskingColumnLifecycleStateCreating,
	"ACTIVE":          MaskingColumnLifecycleStateActive,
	"UPDATING":        MaskingColumnLifecycleStateUpdating,
	"DELETING":        MaskingColumnLifecycleStateDeleting,
	"NEEDS_ATTENTION": MaskingColumnLifecycleStateNeedsAttention,
	"FAILED":          MaskingColumnLifecycleStateFailed,
}

// GetMaskingColumnLifecycleStateEnumValues Enumerates the set of values for MaskingColumnLifecycleStateEnum
func GetMaskingColumnLifecycleStateEnumValues() []MaskingColumnLifecycleStateEnum {
	values := make([]MaskingColumnLifecycleStateEnum, 0)
	for _, v := range mappingMaskingColumnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingColumnLifecycleStateEnumStringValues Enumerates the set of values in String for MaskingColumnLifecycleStateEnum
func GetMaskingColumnLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingMaskingColumnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingColumnLifecycleStateEnum(val string) (MaskingColumnLifecycleStateEnum, bool) {
	mappingMaskingColumnLifecycleStateEnumIgnoreCase := make(map[string]MaskingColumnLifecycleStateEnum)
	for k, v := range mappingMaskingColumnLifecycleStateEnum {
		mappingMaskingColumnLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaskingColumnLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
