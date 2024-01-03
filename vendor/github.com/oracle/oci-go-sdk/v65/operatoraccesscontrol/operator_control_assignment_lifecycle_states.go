// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"strings"
)

// OperatorControlAssignmentLifecycleStatesEnum Enum with underlying type: string
type OperatorControlAssignmentLifecycleStatesEnum string

// Set of constants representing the allowable values for OperatorControlAssignmentLifecycleStatesEnum
const (
	OperatorControlAssignmentLifecycleStatesCreated        OperatorControlAssignmentLifecycleStatesEnum = "CREATED"
	OperatorControlAssignmentLifecycleStatesApplied        OperatorControlAssignmentLifecycleStatesEnum = "APPLIED"
	OperatorControlAssignmentLifecycleStatesApplyfailed    OperatorControlAssignmentLifecycleStatesEnum = "APPLYFAILED"
	OperatorControlAssignmentLifecycleStatesUpdating       OperatorControlAssignmentLifecycleStatesEnum = "UPDATING"
	OperatorControlAssignmentLifecycleStatesUpdatefailed   OperatorControlAssignmentLifecycleStatesEnum = "UPDATEFAILED"
	OperatorControlAssignmentLifecycleStatesDeleting       OperatorControlAssignmentLifecycleStatesEnum = "DELETING"
	OperatorControlAssignmentLifecycleStatesDeleted        OperatorControlAssignmentLifecycleStatesEnum = "DELETED"
	OperatorControlAssignmentLifecycleStatesDeletionfailed OperatorControlAssignmentLifecycleStatesEnum = "DELETIONFAILED"
)

var mappingOperatorControlAssignmentLifecycleStatesEnum = map[string]OperatorControlAssignmentLifecycleStatesEnum{
	"CREATED":        OperatorControlAssignmentLifecycleStatesCreated,
	"APPLIED":        OperatorControlAssignmentLifecycleStatesApplied,
	"APPLYFAILED":    OperatorControlAssignmentLifecycleStatesApplyfailed,
	"UPDATING":       OperatorControlAssignmentLifecycleStatesUpdating,
	"UPDATEFAILED":   OperatorControlAssignmentLifecycleStatesUpdatefailed,
	"DELETING":       OperatorControlAssignmentLifecycleStatesDeleting,
	"DELETED":        OperatorControlAssignmentLifecycleStatesDeleted,
	"DELETIONFAILED": OperatorControlAssignmentLifecycleStatesDeletionfailed,
}

var mappingOperatorControlAssignmentLifecycleStatesEnumLowerCase = map[string]OperatorControlAssignmentLifecycleStatesEnum{
	"created":        OperatorControlAssignmentLifecycleStatesCreated,
	"applied":        OperatorControlAssignmentLifecycleStatesApplied,
	"applyfailed":    OperatorControlAssignmentLifecycleStatesApplyfailed,
	"updating":       OperatorControlAssignmentLifecycleStatesUpdating,
	"updatefailed":   OperatorControlAssignmentLifecycleStatesUpdatefailed,
	"deleting":       OperatorControlAssignmentLifecycleStatesDeleting,
	"deleted":        OperatorControlAssignmentLifecycleStatesDeleted,
	"deletionfailed": OperatorControlAssignmentLifecycleStatesDeletionfailed,
}

// GetOperatorControlAssignmentLifecycleStatesEnumValues Enumerates the set of values for OperatorControlAssignmentLifecycleStatesEnum
func GetOperatorControlAssignmentLifecycleStatesEnumValues() []OperatorControlAssignmentLifecycleStatesEnum {
	values := make([]OperatorControlAssignmentLifecycleStatesEnum, 0)
	for _, v := range mappingOperatorControlAssignmentLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorControlAssignmentLifecycleStatesEnumStringValues Enumerates the set of values in String for OperatorControlAssignmentLifecycleStatesEnum
func GetOperatorControlAssignmentLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPLIED",
		"APPLYFAILED",
		"UPDATING",
		"UPDATEFAILED",
		"DELETING",
		"DELETED",
		"DELETIONFAILED",
	}
}

// GetMappingOperatorControlAssignmentLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorControlAssignmentLifecycleStatesEnum(val string) (OperatorControlAssignmentLifecycleStatesEnum, bool) {
	enum, ok := mappingOperatorControlAssignmentLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
