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

// TypeEnumEnum Enum with underlying type: string
type TypeEnumEnum string

// Set of constants representing the allowable values for TypeEnumEnum
const (
	TypeEnumNotify  TypeEnumEnum = "NOTIFY"
	TypeEnumExecute TypeEnumEnum = "EXECUTE"
)

var mappingTypeEnumEnum = map[string]TypeEnumEnum{
	"NOTIFY":  TypeEnumNotify,
	"EXECUTE": TypeEnumExecute,
}

var mappingTypeEnumEnumLowerCase = map[string]TypeEnumEnum{
	"notify":  TypeEnumNotify,
	"execute": TypeEnumExecute,
}

// GetTypeEnumEnumValues Enumerates the set of values for TypeEnumEnum
func GetTypeEnumEnumValues() []TypeEnumEnum {
	values := make([]TypeEnumEnum, 0)
	for _, v := range mappingTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetTypeEnumEnumStringValues Enumerates the set of values in String for TypeEnumEnum
func GetTypeEnumEnumStringValues() []string {
	return []string{
		"NOTIFY",
		"EXECUTE",
	}
}

// GetMappingTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTypeEnumEnum(val string) (TypeEnumEnum, bool) {
	enum, ok := mappingTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
