// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ProfileTypeEnum Enum with underlying type: string
type ProfileTypeEnum string

// Set of constants representing the allowable values for ProfileTypeEnum
const (
	ProfileTypeSoftwaresource    ProfileTypeEnum = "SOFTWARESOURCE"
	ProfileTypeGroup             ProfileTypeEnum = "GROUP"
	ProfileTypeLifecycle         ProfileTypeEnum = "LIFECYCLE"
	ProfileTypeStation           ProfileTypeEnum = "STATION"
	ProfileTypeWindowsStandalone ProfileTypeEnum = "WINDOWS_STANDALONE"
)

var mappingProfileTypeEnum = map[string]ProfileTypeEnum{
	"SOFTWARESOURCE":     ProfileTypeSoftwaresource,
	"GROUP":              ProfileTypeGroup,
	"LIFECYCLE":          ProfileTypeLifecycle,
	"STATION":            ProfileTypeStation,
	"WINDOWS_STANDALONE": ProfileTypeWindowsStandalone,
}

var mappingProfileTypeEnumLowerCase = map[string]ProfileTypeEnum{
	"softwaresource":     ProfileTypeSoftwaresource,
	"group":              ProfileTypeGroup,
	"lifecycle":          ProfileTypeLifecycle,
	"station":            ProfileTypeStation,
	"windows_standalone": ProfileTypeWindowsStandalone,
}

// GetProfileTypeEnumValues Enumerates the set of values for ProfileTypeEnum
func GetProfileTypeEnumValues() []ProfileTypeEnum {
	values := make([]ProfileTypeEnum, 0)
	for _, v := range mappingProfileTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProfileTypeEnumStringValues Enumerates the set of values in String for ProfileTypeEnum
func GetProfileTypeEnumStringValues() []string {
	return []string{
		"SOFTWARESOURCE",
		"GROUP",
		"LIFECYCLE",
		"STATION",
		"WINDOWS_STANDALONE",
	}
}

// GetMappingProfileTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProfileTypeEnum(val string) (ProfileTypeEnum, bool) {
	enum, ok := mappingProfileTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
