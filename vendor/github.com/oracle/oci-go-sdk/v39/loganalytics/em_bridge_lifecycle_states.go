// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// EmBridgeLifecycleStatesEnum Enum with underlying type: string
type EmBridgeLifecycleStatesEnum string

// Set of constants representing the allowable values for EmBridgeLifecycleStatesEnum
const (
	EmBridgeLifecycleStatesCreating       EmBridgeLifecycleStatesEnum = "CREATING"
	EmBridgeLifecycleStatesActive         EmBridgeLifecycleStatesEnum = "ACTIVE"
	EmBridgeLifecycleStatesDeleted        EmBridgeLifecycleStatesEnum = "DELETED"
	EmBridgeLifecycleStatesNeedsAttention EmBridgeLifecycleStatesEnum = "NEEDS_ATTENTION"
)

var mappingEmBridgeLifecycleStates = map[string]EmBridgeLifecycleStatesEnum{
	"CREATING":        EmBridgeLifecycleStatesCreating,
	"ACTIVE":          EmBridgeLifecycleStatesActive,
	"DELETED":         EmBridgeLifecycleStatesDeleted,
	"NEEDS_ATTENTION": EmBridgeLifecycleStatesNeedsAttention,
}

// GetEmBridgeLifecycleStatesEnumValues Enumerates the set of values for EmBridgeLifecycleStatesEnum
func GetEmBridgeLifecycleStatesEnumValues() []EmBridgeLifecycleStatesEnum {
	values := make([]EmBridgeLifecycleStatesEnum, 0)
	for _, v := range mappingEmBridgeLifecycleStates {
		values = append(values, v)
	}
	return values
}
