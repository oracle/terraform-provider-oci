// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

// PlatformTypesEnum Enum with underlying type: string
type PlatformTypesEnum string

// Set of constants representing the allowable values for PlatformTypesEnum
const (
	PlatformTypesLinux   PlatformTypesEnum = "LINUX"
	PlatformTypesWindows PlatformTypesEnum = "WINDOWS"
)

var mappingPlatformTypes = map[string]PlatformTypesEnum{
	"LINUX":   PlatformTypesLinux,
	"WINDOWS": PlatformTypesWindows,
}

// GetPlatformTypesEnumValues Enumerates the set of values for PlatformTypesEnum
func GetPlatformTypesEnumValues() []PlatformTypesEnum {
	values := make([]PlatformTypesEnum, 0)
	for _, v := range mappingPlatformTypes {
		values = append(values, v)
	}
	return values
}
