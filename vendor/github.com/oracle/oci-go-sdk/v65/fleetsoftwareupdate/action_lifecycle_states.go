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

// ActionLifecycleStatesEnum Enum with underlying type: string
type ActionLifecycleStatesEnum string

// Set of constants representing the allowable values for ActionLifecycleStatesEnum
const (
	ActionLifecycleStatesAccepted       ActionLifecycleStatesEnum = "ACCEPTED"
	ActionLifecycleStatesInProgress     ActionLifecycleStatesEnum = "IN_PROGRESS"
	ActionLifecycleStatesWaiting        ActionLifecycleStatesEnum = "WAITING"
	ActionLifecycleStatesUpdating       ActionLifecycleStatesEnum = "UPDATING"
	ActionLifecycleStatesFailed         ActionLifecycleStatesEnum = "FAILED"
	ActionLifecycleStatesNeedsAttention ActionLifecycleStatesEnum = "NEEDS_ATTENTION"
	ActionLifecycleStatesSucceeded      ActionLifecycleStatesEnum = "SUCCEEDED"
	ActionLifecycleStatesCanceling      ActionLifecycleStatesEnum = "CANCELING"
	ActionLifecycleStatesCanceled       ActionLifecycleStatesEnum = "CANCELED"
	ActionLifecycleStatesUnknown        ActionLifecycleStatesEnum = "UNKNOWN"
	ActionLifecycleStatesDeleting       ActionLifecycleStatesEnum = "DELETING"
	ActionLifecycleStatesDeleted        ActionLifecycleStatesEnum = "DELETED"
)

var mappingActionLifecycleStatesEnum = map[string]ActionLifecycleStatesEnum{
	"ACCEPTED":        ActionLifecycleStatesAccepted,
	"IN_PROGRESS":     ActionLifecycleStatesInProgress,
	"WAITING":         ActionLifecycleStatesWaiting,
	"UPDATING":        ActionLifecycleStatesUpdating,
	"FAILED":          ActionLifecycleStatesFailed,
	"NEEDS_ATTENTION": ActionLifecycleStatesNeedsAttention,
	"SUCCEEDED":       ActionLifecycleStatesSucceeded,
	"CANCELING":       ActionLifecycleStatesCanceling,
	"CANCELED":        ActionLifecycleStatesCanceled,
	"UNKNOWN":         ActionLifecycleStatesUnknown,
	"DELETING":        ActionLifecycleStatesDeleting,
	"DELETED":         ActionLifecycleStatesDeleted,
}

var mappingActionLifecycleStatesEnumLowerCase = map[string]ActionLifecycleStatesEnum{
	"accepted":        ActionLifecycleStatesAccepted,
	"in_progress":     ActionLifecycleStatesInProgress,
	"waiting":         ActionLifecycleStatesWaiting,
	"updating":        ActionLifecycleStatesUpdating,
	"failed":          ActionLifecycleStatesFailed,
	"needs_attention": ActionLifecycleStatesNeedsAttention,
	"succeeded":       ActionLifecycleStatesSucceeded,
	"canceling":       ActionLifecycleStatesCanceling,
	"canceled":        ActionLifecycleStatesCanceled,
	"unknown":         ActionLifecycleStatesUnknown,
	"deleting":        ActionLifecycleStatesDeleting,
	"deleted":         ActionLifecycleStatesDeleted,
}

// GetActionLifecycleStatesEnumValues Enumerates the set of values for ActionLifecycleStatesEnum
func GetActionLifecycleStatesEnumValues() []ActionLifecycleStatesEnum {
	values := make([]ActionLifecycleStatesEnum, 0)
	for _, v := range mappingActionLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetActionLifecycleStatesEnumStringValues Enumerates the set of values in String for ActionLifecycleStatesEnum
func GetActionLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"UPDATING",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"UNKNOWN",
		"DELETING",
		"DELETED",
	}
}

// GetMappingActionLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionLifecycleStatesEnum(val string) (ActionLifecycleStatesEnum, bool) {
	enum, ok := mappingActionLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
