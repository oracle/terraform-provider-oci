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

// DiscoveryLifecycleStatesEnum Enum with underlying type: string
type DiscoveryLifecycleStatesEnum string

// Set of constants representing the allowable values for DiscoveryLifecycleStatesEnum
const (
	DiscoveryLifecycleStatesAccepted   DiscoveryLifecycleStatesEnum = "ACCEPTED"
	DiscoveryLifecycleStatesInProgress DiscoveryLifecycleStatesEnum = "IN_PROGRESS"
	DiscoveryLifecycleStatesFailed     DiscoveryLifecycleStatesEnum = "FAILED"
	DiscoveryLifecycleStatesSucceeded  DiscoveryLifecycleStatesEnum = "SUCCEEDED"
	DiscoveryLifecycleStatesCanceling  DiscoveryLifecycleStatesEnum = "CANCELING"
	DiscoveryLifecycleStatesCanceled   DiscoveryLifecycleStatesEnum = "CANCELED"
	DiscoveryLifecycleStatesDeleting   DiscoveryLifecycleStatesEnum = "DELETING"
	DiscoveryLifecycleStatesDeleted    DiscoveryLifecycleStatesEnum = "DELETED"
)

var mappingDiscoveryLifecycleStatesEnum = map[string]DiscoveryLifecycleStatesEnum{
	"ACCEPTED":    DiscoveryLifecycleStatesAccepted,
	"IN_PROGRESS": DiscoveryLifecycleStatesInProgress,
	"FAILED":      DiscoveryLifecycleStatesFailed,
	"SUCCEEDED":   DiscoveryLifecycleStatesSucceeded,
	"CANCELING":   DiscoveryLifecycleStatesCanceling,
	"CANCELED":    DiscoveryLifecycleStatesCanceled,
	"DELETING":    DiscoveryLifecycleStatesDeleting,
	"DELETED":     DiscoveryLifecycleStatesDeleted,
}

var mappingDiscoveryLifecycleStatesEnumLowerCase = map[string]DiscoveryLifecycleStatesEnum{
	"accepted":    DiscoveryLifecycleStatesAccepted,
	"in_progress": DiscoveryLifecycleStatesInProgress,
	"failed":      DiscoveryLifecycleStatesFailed,
	"succeeded":   DiscoveryLifecycleStatesSucceeded,
	"canceling":   DiscoveryLifecycleStatesCanceling,
	"canceled":    DiscoveryLifecycleStatesCanceled,
	"deleting":    DiscoveryLifecycleStatesDeleting,
	"deleted":     DiscoveryLifecycleStatesDeleted,
}

// GetDiscoveryLifecycleStatesEnumValues Enumerates the set of values for DiscoveryLifecycleStatesEnum
func GetDiscoveryLifecycleStatesEnumValues() []DiscoveryLifecycleStatesEnum {
	values := make([]DiscoveryLifecycleStatesEnum, 0)
	for _, v := range mappingDiscoveryLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryLifecycleStatesEnumStringValues Enumerates the set of values in String for DiscoveryLifecycleStatesEnum
func GetDiscoveryLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingDiscoveryLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryLifecycleStatesEnum(val string) (DiscoveryLifecycleStatesEnum, bool) {
	enum, ok := mappingDiscoveryLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
