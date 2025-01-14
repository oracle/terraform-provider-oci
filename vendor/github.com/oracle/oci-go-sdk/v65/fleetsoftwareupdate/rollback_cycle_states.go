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

// RollbackCycleStatesEnum Enum with underlying type: string
type RollbackCycleStatesEnum string

// Set of constants representing the allowable values for RollbackCycleStatesEnum
const (
	RollbackCycleStatesAbleToExecute  RollbackCycleStatesEnum = "ABLE_TO_EXECUTE"
	RollbackCycleStatesInProgress     RollbackCycleStatesEnum = "IN_PROGRESS"
	RollbackCycleStatesFailed         RollbackCycleStatesEnum = "FAILED"
	RollbackCycleStatesNeedsAttention RollbackCycleStatesEnum = "NEEDS_ATTENTION"
	RollbackCycleStatesSucceeded      RollbackCycleStatesEnum = "SUCCEEDED"
)

var mappingRollbackCycleStatesEnum = map[string]RollbackCycleStatesEnum{
	"ABLE_TO_EXECUTE": RollbackCycleStatesAbleToExecute,
	"IN_PROGRESS":     RollbackCycleStatesInProgress,
	"FAILED":          RollbackCycleStatesFailed,
	"NEEDS_ATTENTION": RollbackCycleStatesNeedsAttention,
	"SUCCEEDED":       RollbackCycleStatesSucceeded,
}

var mappingRollbackCycleStatesEnumLowerCase = map[string]RollbackCycleStatesEnum{
	"able_to_execute": RollbackCycleStatesAbleToExecute,
	"in_progress":     RollbackCycleStatesInProgress,
	"failed":          RollbackCycleStatesFailed,
	"needs_attention": RollbackCycleStatesNeedsAttention,
	"succeeded":       RollbackCycleStatesSucceeded,
}

// GetRollbackCycleStatesEnumValues Enumerates the set of values for RollbackCycleStatesEnum
func GetRollbackCycleStatesEnumValues() []RollbackCycleStatesEnum {
	values := make([]RollbackCycleStatesEnum, 0)
	for _, v := range mappingRollbackCycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetRollbackCycleStatesEnumStringValues Enumerates the set of values in String for RollbackCycleStatesEnum
func GetRollbackCycleStatesEnumStringValues() []string {
	return []string{
		"ABLE_TO_EXECUTE",
		"IN_PROGRESS",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
	}
}

// GetMappingRollbackCycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRollbackCycleStatesEnum(val string) (RollbackCycleStatesEnum, bool) {
	enum, ok := mappingRollbackCycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
