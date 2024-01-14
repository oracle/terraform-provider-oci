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

// FindingLifecycleStateEnum Enum with underlying type: string
type FindingLifecycleStateEnum string

// Set of constants representing the allowable values for FindingLifecycleStateEnum
const (
	FindingLifecycleStateActive         FindingLifecycleStateEnum = "ACTIVE"
	FindingLifecycleStateUpdating       FindingLifecycleStateEnum = "UPDATING"
	FindingLifecycleStateNeedsAttention FindingLifecycleStateEnum = "NEEDS_ATTENTION"
	FindingLifecycleStateFailed         FindingLifecycleStateEnum = "FAILED"
)

var mappingFindingLifecycleStateEnum = map[string]FindingLifecycleStateEnum{
	"ACTIVE":          FindingLifecycleStateActive,
	"UPDATING":        FindingLifecycleStateUpdating,
	"NEEDS_ATTENTION": FindingLifecycleStateNeedsAttention,
	"FAILED":          FindingLifecycleStateFailed,
}

var mappingFindingLifecycleStateEnumLowerCase = map[string]FindingLifecycleStateEnum{
	"active":          FindingLifecycleStateActive,
	"updating":        FindingLifecycleStateUpdating,
	"needs_attention": FindingLifecycleStateNeedsAttention,
	"failed":          FindingLifecycleStateFailed,
}

// GetFindingLifecycleStateEnumValues Enumerates the set of values for FindingLifecycleStateEnum
func GetFindingLifecycleStateEnumValues() []FindingLifecycleStateEnum {
	values := make([]FindingLifecycleStateEnum, 0)
	for _, v := range mappingFindingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFindingLifecycleStateEnumStringValues Enumerates the set of values in String for FindingLifecycleStateEnum
func GetFindingLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingFindingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFindingLifecycleStateEnum(val string) (FindingLifecycleStateEnum, bool) {
	enum, ok := mappingFindingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
