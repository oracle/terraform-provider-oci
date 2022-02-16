// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// PlatformTypesEnum Enum with underlying type: string
type PlatformTypesEnum string

// Set of constants representing the allowable values for PlatformTypesEnum
const (
	PlatformTypesLinux   PlatformTypesEnum = "LINUX"
	PlatformTypesWindows PlatformTypesEnum = "WINDOWS"
	PlatformTypesSolaris PlatformTypesEnum = "SOLARIS"
)

var mappingPlatformTypesEnum = map[string]PlatformTypesEnum{
	"LINUX":   PlatformTypesLinux,
	"WINDOWS": PlatformTypesWindows,
	"SOLARIS": PlatformTypesSolaris,
}

// GetPlatformTypesEnumValues Enumerates the set of values for PlatformTypesEnum
func GetPlatformTypesEnumValues() []PlatformTypesEnum {
	values := make([]PlatformTypesEnum, 0)
	for _, v := range mappingPlatformTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformTypesEnumStringValues Enumerates the set of values in String for PlatformTypesEnum
func GetPlatformTypesEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"SOLARIS",
	}
}

// GetMappingPlatformTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformTypesEnum(val string) (PlatformTypesEnum, bool) {
	mappingPlatformTypesEnumIgnoreCase := make(map[string]PlatformTypesEnum)
	for k, v := range mappingPlatformTypesEnum {
		mappingPlatformTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPlatformTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
