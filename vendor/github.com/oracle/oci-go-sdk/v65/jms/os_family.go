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

// OsFamilyEnum Enum with underlying type: string
type OsFamilyEnum string

// Set of constants representing the allowable values for OsFamilyEnum
const (
	OsFamilyLinux   OsFamilyEnum = "LINUX"
	OsFamilyWindows OsFamilyEnum = "WINDOWS"
	OsFamilyMacos   OsFamilyEnum = "MACOS"
	OsFamilyUnknown OsFamilyEnum = "UNKNOWN"
)

var mappingOsFamilyEnum = map[string]OsFamilyEnum{
	"LINUX":   OsFamilyLinux,
	"WINDOWS": OsFamilyWindows,
	"MACOS":   OsFamilyMacos,
	"UNKNOWN": OsFamilyUnknown,
}

var mappingOsFamilyEnumLowerCase = map[string]OsFamilyEnum{
	"linux":   OsFamilyLinux,
	"windows": OsFamilyWindows,
	"macos":   OsFamilyMacos,
	"unknown": OsFamilyUnknown,
}

// GetOsFamilyEnumValues Enumerates the set of values for OsFamilyEnum
func GetOsFamilyEnumValues() []OsFamilyEnum {
	values := make([]OsFamilyEnum, 0)
	for _, v := range mappingOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetOsFamilyEnumStringValues Enumerates the set of values in String for OsFamilyEnum
func GetOsFamilyEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"MACOS",
		"UNKNOWN",
	}
}

// GetMappingOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsFamilyEnum(val string) (OsFamilyEnum, bool) {
	enum, ok := mappingOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
