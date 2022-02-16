// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// OperatorControlLifecycleStatesEnum Enum with underlying type: string
type OperatorControlLifecycleStatesEnum string

// Set of constants representing the allowable values for OperatorControlLifecycleStatesEnum
const (
	OperatorControlLifecycleStatesCreated    OperatorControlLifecycleStatesEnum = "CREATED"
	OperatorControlLifecycleStatesAssigned   OperatorControlLifecycleStatesEnum = "ASSIGNED"
	OperatorControlLifecycleStatesUnassigned OperatorControlLifecycleStatesEnum = "UNASSIGNED"
	OperatorControlLifecycleStatesDeleted    OperatorControlLifecycleStatesEnum = "DELETED"
)

var mappingOperatorControlLifecycleStatesEnum = map[string]OperatorControlLifecycleStatesEnum{
	"CREATED":    OperatorControlLifecycleStatesCreated,
	"ASSIGNED":   OperatorControlLifecycleStatesAssigned,
	"UNASSIGNED": OperatorControlLifecycleStatesUnassigned,
	"DELETED":    OperatorControlLifecycleStatesDeleted,
}

// GetOperatorControlLifecycleStatesEnumValues Enumerates the set of values for OperatorControlLifecycleStatesEnum
func GetOperatorControlLifecycleStatesEnumValues() []OperatorControlLifecycleStatesEnum {
	values := make([]OperatorControlLifecycleStatesEnum, 0)
	for _, v := range mappingOperatorControlLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorControlLifecycleStatesEnumStringValues Enumerates the set of values in String for OperatorControlLifecycleStatesEnum
func GetOperatorControlLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATED",
		"ASSIGNED",
		"UNASSIGNED",
		"DELETED",
	}
}

// GetMappingOperatorControlLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorControlLifecycleStatesEnum(val string) (OperatorControlLifecycleStatesEnum, bool) {
	mappingOperatorControlLifecycleStatesEnumIgnoreCase := make(map[string]OperatorControlLifecycleStatesEnum)
	for k, v := range mappingOperatorControlLifecycleStatesEnum {
		mappingOperatorControlLifecycleStatesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperatorControlLifecycleStatesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
