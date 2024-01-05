// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// HealthEnum Enum with underlying type: string
type HealthEnum string

// Set of constants representing the allowable values for HealthEnum
const (
	HealthProtected HealthEnum = "PROTECTED"
	HealthWarning   HealthEnum = "WARNING"
	HealthAlert     HealthEnum = "ALERT"
)

var mappingHealthEnum = map[string]HealthEnum{
	"PROTECTED": HealthProtected,
	"WARNING":   HealthWarning,
	"ALERT":     HealthAlert,
}

var mappingHealthEnumLowerCase = map[string]HealthEnum{
	"protected": HealthProtected,
	"warning":   HealthWarning,
	"alert":     HealthAlert,
}

// GetHealthEnumValues Enumerates the set of values for HealthEnum
func GetHealthEnumValues() []HealthEnum {
	values := make([]HealthEnum, 0)
	for _, v := range mappingHealthEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthEnumStringValues Enumerates the set of values in String for HealthEnum
func GetHealthEnumStringValues() []string {
	return []string{
		"PROTECTED",
		"WARNING",
		"ALERT",
	}
}

// GetMappingHealthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthEnum(val string) (HealthEnum, bool) {
	enum, ok := mappingHealthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
