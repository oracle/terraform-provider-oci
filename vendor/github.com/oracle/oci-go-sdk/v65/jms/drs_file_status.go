// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// DrsFileStatusEnum Enum with underlying type: string
type DrsFileStatusEnum string

// Set of constants representing the allowable values for DrsFileStatusEnum
const (
	DrsFileStatusPresent       DrsFileStatusEnum = "PRESENT"
	DrsFileStatusAbsent        DrsFileStatusEnum = "ABSENT"
	DrsFileStatusMismatch      DrsFileStatusEnum = "MISMATCH"
	DrsFileStatusNotConfigured DrsFileStatusEnum = "NOT_CONFIGURED"
)

var mappingDrsFileStatusEnum = map[string]DrsFileStatusEnum{
	"PRESENT":        DrsFileStatusPresent,
	"ABSENT":         DrsFileStatusAbsent,
	"MISMATCH":       DrsFileStatusMismatch,
	"NOT_CONFIGURED": DrsFileStatusNotConfigured,
}

var mappingDrsFileStatusEnumLowerCase = map[string]DrsFileStatusEnum{
	"present":        DrsFileStatusPresent,
	"absent":         DrsFileStatusAbsent,
	"mismatch":       DrsFileStatusMismatch,
	"not_configured": DrsFileStatusNotConfigured,
}

// GetDrsFileStatusEnumValues Enumerates the set of values for DrsFileStatusEnum
func GetDrsFileStatusEnumValues() []DrsFileStatusEnum {
	values := make([]DrsFileStatusEnum, 0)
	for _, v := range mappingDrsFileStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrsFileStatusEnumStringValues Enumerates the set of values in String for DrsFileStatusEnum
func GetDrsFileStatusEnumStringValues() []string {
	return []string{
		"PRESENT",
		"ABSENT",
		"MISMATCH",
		"NOT_CONFIGURED",
	}
}

// GetMappingDrsFileStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrsFileStatusEnum(val string) (DrsFileStatusEnum, bool) {
	enum, ok := mappingDrsFileStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
