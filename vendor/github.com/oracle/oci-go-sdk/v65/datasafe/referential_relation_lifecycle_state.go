// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ReferentialRelationLifecycleStateEnum Enum with underlying type: string
type ReferentialRelationLifecycleStateEnum string

// Set of constants representing the allowable values for ReferentialRelationLifecycleStateEnum
const (
	ReferentialRelationLifecycleStateCreating ReferentialRelationLifecycleStateEnum = "CREATING"
	ReferentialRelationLifecycleStateActive   ReferentialRelationLifecycleStateEnum = "ACTIVE"
	ReferentialRelationLifecycleStateUpdating ReferentialRelationLifecycleStateEnum = "UPDATING"
	ReferentialRelationLifecycleStateDeleting ReferentialRelationLifecycleStateEnum = "DELETING"
	ReferentialRelationLifecycleStateFailed   ReferentialRelationLifecycleStateEnum = "FAILED"
)

var mappingReferentialRelationLifecycleStateEnum = map[string]ReferentialRelationLifecycleStateEnum{
	"CREATING": ReferentialRelationLifecycleStateCreating,
	"ACTIVE":   ReferentialRelationLifecycleStateActive,
	"UPDATING": ReferentialRelationLifecycleStateUpdating,
	"DELETING": ReferentialRelationLifecycleStateDeleting,
	"FAILED":   ReferentialRelationLifecycleStateFailed,
}

var mappingReferentialRelationLifecycleStateEnumLowerCase = map[string]ReferentialRelationLifecycleStateEnum{
	"creating": ReferentialRelationLifecycleStateCreating,
	"active":   ReferentialRelationLifecycleStateActive,
	"updating": ReferentialRelationLifecycleStateUpdating,
	"deleting": ReferentialRelationLifecycleStateDeleting,
	"failed":   ReferentialRelationLifecycleStateFailed,
}

// GetReferentialRelationLifecycleStateEnumValues Enumerates the set of values for ReferentialRelationLifecycleStateEnum
func GetReferentialRelationLifecycleStateEnumValues() []ReferentialRelationLifecycleStateEnum {
	values := make([]ReferentialRelationLifecycleStateEnum, 0)
	for _, v := range mappingReferentialRelationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReferentialRelationLifecycleStateEnumStringValues Enumerates the set of values in String for ReferentialRelationLifecycleStateEnum
func GetReferentialRelationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

// GetMappingReferentialRelationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReferentialRelationLifecycleStateEnum(val string) (ReferentialRelationLifecycleStateEnum, bool) {
	enum, ok := mappingReferentialRelationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
