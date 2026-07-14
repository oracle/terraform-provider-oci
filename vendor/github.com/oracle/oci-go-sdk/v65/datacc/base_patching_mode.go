// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// BasePatchingModeEnum Enum with underlying type: string
type BasePatchingModeEnum string

// Set of constants representing the allowable values for BasePatchingModeEnum
const (
	BasePatchingModeRolling    BasePatchingModeEnum = "ROLLING"
	BasePatchingModeNonrolling BasePatchingModeEnum = "NONROLLING"
)

var mappingBasePatchingModeEnum = map[string]BasePatchingModeEnum{
	"ROLLING":    BasePatchingModeRolling,
	"NONROLLING": BasePatchingModeNonrolling,
}

var mappingBasePatchingModeEnumLowerCase = map[string]BasePatchingModeEnum{
	"rolling":    BasePatchingModeRolling,
	"nonrolling": BasePatchingModeNonrolling,
}

// GetBasePatchingModeEnumValues Enumerates the set of values for BasePatchingModeEnum
func GetBasePatchingModeEnumValues() []BasePatchingModeEnum {
	values := make([]BasePatchingModeEnum, 0)
	for _, v := range mappingBasePatchingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetBasePatchingModeEnumStringValues Enumerates the set of values in String for BasePatchingModeEnum
func GetBasePatchingModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingBasePatchingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBasePatchingModeEnum(val string) (BasePatchingModeEnum, bool) {
	enum, ok := mappingBasePatchingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
