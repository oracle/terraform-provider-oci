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

// TargetResourceTypeEnumEnum Enum with underlying type: string
type TargetResourceTypeEnumEnum string

// Set of constants representing the allowable values for TargetResourceTypeEnumEnum
const (
	TargetResourceTypeEnumDbCcInfrastructure TargetResourceTypeEnumEnum = "DB_CC_INFRASTRUCTURE"
)

var mappingTargetResourceTypeEnumEnum = map[string]TargetResourceTypeEnumEnum{
	"DB_CC_INFRASTRUCTURE": TargetResourceTypeEnumDbCcInfrastructure,
}

var mappingTargetResourceTypeEnumEnumLowerCase = map[string]TargetResourceTypeEnumEnum{
	"db_cc_infrastructure": TargetResourceTypeEnumDbCcInfrastructure,
}

// GetTargetResourceTypeEnumEnumValues Enumerates the set of values for TargetResourceTypeEnumEnum
func GetTargetResourceTypeEnumEnumValues() []TargetResourceTypeEnumEnum {
	values := make([]TargetResourceTypeEnumEnum, 0)
	for _, v := range mappingTargetResourceTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetResourceTypeEnumEnumStringValues Enumerates the set of values in String for TargetResourceTypeEnumEnum
func GetTargetResourceTypeEnumEnumStringValues() []string {
	return []string{
		"DB_CC_INFRASTRUCTURE",
	}
}

// GetMappingTargetResourceTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetResourceTypeEnumEnum(val string) (TargetResourceTypeEnumEnum, bool) {
	enum, ok := mappingTargetResourceTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
