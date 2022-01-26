// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

// EventTypeEnum Enum with underlying type: string
type EventTypeEnum string

// Set of constants representing the allowable values for EventTypeEnum
const (
	EventTypeKernelOops       EventTypeEnum = "KERNEL_OOPS"
	EventTypeKernelCrash      EventTypeEnum = "KERNEL_CRASH"
	EventTypeCrash            EventTypeEnum = "CRASH"
	EventTypeExploitAttempt   EventTypeEnum = "EXPLOIT_ATTEMPT"
	EventTypeCompliance       EventTypeEnum = "COMPLIANCE"
	EventTypeTuningSuggestion EventTypeEnum = "TUNING_SUGGESTION"
	EventTypeTuningApplied    EventTypeEnum = "TUNING_APPLIED"
	EventTypeSecurity         EventTypeEnum = "SECURITY"
	EventTypeError            EventTypeEnum = "ERROR"
	EventTypeWarning          EventTypeEnum = "WARNING"
)

var mappingEventType = map[string]EventTypeEnum{
	"KERNEL_OOPS":       EventTypeKernelOops,
	"KERNEL_CRASH":      EventTypeKernelCrash,
	"CRASH":             EventTypeCrash,
	"EXPLOIT_ATTEMPT":   EventTypeExploitAttempt,
	"COMPLIANCE":        EventTypeCompliance,
	"TUNING_SUGGESTION": EventTypeTuningSuggestion,
	"TUNING_APPLIED":    EventTypeTuningApplied,
	"SECURITY":          EventTypeSecurity,
	"ERROR":             EventTypeError,
	"WARNING":           EventTypeWarning,
}

// GetEventTypeEnumValues Enumerates the set of values for EventTypeEnum
func GetEventTypeEnumValues() []EventTypeEnum {
	values := make([]EventTypeEnum, 0)
	for _, v := range mappingEventType {
		values = append(values, v)
	}
	return values
}
