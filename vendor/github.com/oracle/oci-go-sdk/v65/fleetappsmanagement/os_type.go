// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// OsTypeEnum Enum with underlying type: string
type OsTypeEnum string

// Set of constants representing the allowable values for OsTypeEnum
const (
	OsTypeWindows OsTypeEnum = "WINDOWS"
	OsTypeLinux   OsTypeEnum = "LINUX"
	OsTypeGeneric OsTypeEnum = "GENERIC"
)

var mappingOsTypeEnum = map[string]OsTypeEnum{
	"WINDOWS": OsTypeWindows,
	"LINUX":   OsTypeLinux,
	"GENERIC": OsTypeGeneric,
}

var mappingOsTypeEnumLowerCase = map[string]OsTypeEnum{
	"windows": OsTypeWindows,
	"linux":   OsTypeLinux,
	"generic": OsTypeGeneric,
}

// GetOsTypeEnumValues Enumerates the set of values for OsTypeEnum
func GetOsTypeEnumValues() []OsTypeEnum {
	values := make([]OsTypeEnum, 0)
	for _, v := range mappingOsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOsTypeEnumStringValues Enumerates the set of values in String for OsTypeEnum
func GetOsTypeEnumStringValues() []string {
	return []string{
		"WINDOWS",
		"LINUX",
		"GENERIC",
	}
}

// GetMappingOsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsTypeEnum(val string) (OsTypeEnum, bool) {
	enum, ok := mappingOsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
