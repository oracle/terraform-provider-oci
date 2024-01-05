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

// OperatorActionLifecycleStatesEnum Enum with underlying type: string
type OperatorActionLifecycleStatesEnum string

// Set of constants representing the allowable values for OperatorActionLifecycleStatesEnum
const (
	OperatorActionLifecycleStatesActive   OperatorActionLifecycleStatesEnum = "ACTIVE"
	OperatorActionLifecycleStatesInactive OperatorActionLifecycleStatesEnum = "INACTIVE"
)

var mappingOperatorActionLifecycleStatesEnum = map[string]OperatorActionLifecycleStatesEnum{
	"ACTIVE":   OperatorActionLifecycleStatesActive,
	"INACTIVE": OperatorActionLifecycleStatesInactive,
}

var mappingOperatorActionLifecycleStatesEnumLowerCase = map[string]OperatorActionLifecycleStatesEnum{
	"active":   OperatorActionLifecycleStatesActive,
	"inactive": OperatorActionLifecycleStatesInactive,
}

// GetOperatorActionLifecycleStatesEnumValues Enumerates the set of values for OperatorActionLifecycleStatesEnum
func GetOperatorActionLifecycleStatesEnumValues() []OperatorActionLifecycleStatesEnum {
	values := make([]OperatorActionLifecycleStatesEnum, 0)
	for _, v := range mappingOperatorActionLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperatorActionLifecycleStatesEnumStringValues Enumerates the set of values in String for OperatorActionLifecycleStatesEnum
func GetOperatorActionLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingOperatorActionLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperatorActionLifecycleStatesEnum(val string) (OperatorActionLifecycleStatesEnum, bool) {
	enum, ok := mappingOperatorActionLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
