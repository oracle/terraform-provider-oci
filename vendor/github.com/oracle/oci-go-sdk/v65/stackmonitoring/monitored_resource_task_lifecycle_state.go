// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MonitoredResourceTaskLifecycleStateEnum Enum with underlying type: string
type MonitoredResourceTaskLifecycleStateEnum string

// Set of constants representing the allowable values for MonitoredResourceTaskLifecycleStateEnum
const (
	MonitoredResourceTaskLifecycleStateAccepted       MonitoredResourceTaskLifecycleStateEnum = "ACCEPTED"
	MonitoredResourceTaskLifecycleStateInProgress     MonitoredResourceTaskLifecycleStateEnum = "IN_PROGRESS"
	MonitoredResourceTaskLifecycleStateWaiting        MonitoredResourceTaskLifecycleStateEnum = "WAITING"
	MonitoredResourceTaskLifecycleStateFailed         MonitoredResourceTaskLifecycleStateEnum = "FAILED"
	MonitoredResourceTaskLifecycleStateSucceeded      MonitoredResourceTaskLifecycleStateEnum = "SUCCEEDED"
	MonitoredResourceTaskLifecycleStateCanceling      MonitoredResourceTaskLifecycleStateEnum = "CANCELING"
	MonitoredResourceTaskLifecycleStateCanceled       MonitoredResourceTaskLifecycleStateEnum = "CANCELED"
	MonitoredResourceTaskLifecycleStateNeedsAttention MonitoredResourceTaskLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingMonitoredResourceTaskLifecycleStateEnum = map[string]MonitoredResourceTaskLifecycleStateEnum{
	"ACCEPTED":        MonitoredResourceTaskLifecycleStateAccepted,
	"IN_PROGRESS":     MonitoredResourceTaskLifecycleStateInProgress,
	"WAITING":         MonitoredResourceTaskLifecycleStateWaiting,
	"FAILED":          MonitoredResourceTaskLifecycleStateFailed,
	"SUCCEEDED":       MonitoredResourceTaskLifecycleStateSucceeded,
	"CANCELING":       MonitoredResourceTaskLifecycleStateCanceling,
	"CANCELED":        MonitoredResourceTaskLifecycleStateCanceled,
	"NEEDS_ATTENTION": MonitoredResourceTaskLifecycleStateNeedsAttention,
}

var mappingMonitoredResourceTaskLifecycleStateEnumLowerCase = map[string]MonitoredResourceTaskLifecycleStateEnum{
	"accepted":        MonitoredResourceTaskLifecycleStateAccepted,
	"in_progress":     MonitoredResourceTaskLifecycleStateInProgress,
	"waiting":         MonitoredResourceTaskLifecycleStateWaiting,
	"failed":          MonitoredResourceTaskLifecycleStateFailed,
	"succeeded":       MonitoredResourceTaskLifecycleStateSucceeded,
	"canceling":       MonitoredResourceTaskLifecycleStateCanceling,
	"canceled":        MonitoredResourceTaskLifecycleStateCanceled,
	"needs_attention": MonitoredResourceTaskLifecycleStateNeedsAttention,
}

// GetMonitoredResourceTaskLifecycleStateEnumValues Enumerates the set of values for MonitoredResourceTaskLifecycleStateEnum
func GetMonitoredResourceTaskLifecycleStateEnumValues() []MonitoredResourceTaskLifecycleStateEnum {
	values := make([]MonitoredResourceTaskLifecycleStateEnum, 0)
	for _, v := range mappingMonitoredResourceTaskLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredResourceTaskLifecycleStateEnumStringValues Enumerates the set of values in String for MonitoredResourceTaskLifecycleStateEnum
func GetMonitoredResourceTaskLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingMonitoredResourceTaskLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredResourceTaskLifecycleStateEnum(val string) (MonitoredResourceTaskLifecycleStateEnum, bool) {
	enum, ok := mappingMonitoredResourceTaskLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
