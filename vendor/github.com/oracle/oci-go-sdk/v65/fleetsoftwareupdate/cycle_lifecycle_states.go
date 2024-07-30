// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CycleLifecycleStatesEnum Enum with underlying type: string
type CycleLifecycleStatesEnum string

// Set of constants representing the allowable values for CycleLifecycleStatesEnum
const (
	CycleLifecycleStatesCreating       CycleLifecycleStatesEnum = "CREATING"
	CycleLifecycleStatesActive         CycleLifecycleStatesEnum = "ACTIVE"
	CycleLifecycleStatesUpdating       CycleLifecycleStatesEnum = "UPDATING"
	CycleLifecycleStatesInProgress     CycleLifecycleStatesEnum = "IN_PROGRESS"
	CycleLifecycleStatesFailed         CycleLifecycleStatesEnum = "FAILED"
	CycleLifecycleStatesNeedsAttention CycleLifecycleStatesEnum = "NEEDS_ATTENTION"
	CycleLifecycleStatesSucceeded      CycleLifecycleStatesEnum = "SUCCEEDED"
	CycleLifecycleStatesDeleting       CycleLifecycleStatesEnum = "DELETING"
	CycleLifecycleStatesDeleted        CycleLifecycleStatesEnum = "DELETED"
)

var mappingCycleLifecycleStatesEnum = map[string]CycleLifecycleStatesEnum{
	"CREATING":        CycleLifecycleStatesCreating,
	"ACTIVE":          CycleLifecycleStatesActive,
	"UPDATING":        CycleLifecycleStatesUpdating,
	"IN_PROGRESS":     CycleLifecycleStatesInProgress,
	"FAILED":          CycleLifecycleStatesFailed,
	"NEEDS_ATTENTION": CycleLifecycleStatesNeedsAttention,
	"SUCCEEDED":       CycleLifecycleStatesSucceeded,
	"DELETING":        CycleLifecycleStatesDeleting,
	"DELETED":         CycleLifecycleStatesDeleted,
}

var mappingCycleLifecycleStatesEnumLowerCase = map[string]CycleLifecycleStatesEnum{
	"creating":        CycleLifecycleStatesCreating,
	"active":          CycleLifecycleStatesActive,
	"updating":        CycleLifecycleStatesUpdating,
	"in_progress":     CycleLifecycleStatesInProgress,
	"failed":          CycleLifecycleStatesFailed,
	"needs_attention": CycleLifecycleStatesNeedsAttention,
	"succeeded":       CycleLifecycleStatesSucceeded,
	"deleting":        CycleLifecycleStatesDeleting,
	"deleted":         CycleLifecycleStatesDeleted,
}

// GetCycleLifecycleStatesEnumValues Enumerates the set of values for CycleLifecycleStatesEnum
func GetCycleLifecycleStatesEnumValues() []CycleLifecycleStatesEnum {
	values := make([]CycleLifecycleStatesEnum, 0)
	for _, v := range mappingCycleLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetCycleLifecycleStatesEnumStringValues Enumerates the set of values in String for CycleLifecycleStatesEnum
func GetCycleLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"IN_PROGRESS",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingCycleLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCycleLifecycleStatesEnum(val string) (CycleLifecycleStatesEnum, bool) {
	enum, ok := mappingCycleLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
