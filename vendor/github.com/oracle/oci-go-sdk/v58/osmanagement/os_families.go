// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"strings"
)

// OsFamiliesEnum Enum with underlying type: string
type OsFamiliesEnum string

// Set of constants representing the allowable values for OsFamiliesEnum
const (
	OsFamiliesLinux   OsFamiliesEnum = "LINUX"
	OsFamiliesWindows OsFamiliesEnum = "WINDOWS"
	OsFamiliesAll     OsFamiliesEnum = "ALL"
)

var mappingOsFamiliesEnum = map[string]OsFamiliesEnum{
	"LINUX":   OsFamiliesLinux,
	"WINDOWS": OsFamiliesWindows,
	"ALL":     OsFamiliesAll,
}

// GetOsFamiliesEnumValues Enumerates the set of values for OsFamiliesEnum
func GetOsFamiliesEnumValues() []OsFamiliesEnum {
	values := make([]OsFamiliesEnum, 0)
	for _, v := range mappingOsFamiliesEnum {
		values = append(values, v)
	}
	return values
}

// GetOsFamiliesEnumStringValues Enumerates the set of values in String for OsFamiliesEnum
func GetOsFamiliesEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"ALL",
	}
}

// GetMappingOsFamiliesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOsFamiliesEnum(val string) (OsFamiliesEnum, bool) {
	mappingOsFamiliesEnumIgnoreCase := make(map[string]OsFamiliesEnum)
	for k, v := range mappingOsFamiliesEnum {
		mappingOsFamiliesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOsFamiliesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
