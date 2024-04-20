// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// WindowsUpdateTypesEnum Enum with underlying type: string
type WindowsUpdateTypesEnum string

// Set of constants representing the allowable values for WindowsUpdateTypesEnum
const (
	WindowsUpdateTypesSecurity    WindowsUpdateTypesEnum = "SECURITY"
	WindowsUpdateTypesBugfix      WindowsUpdateTypesEnum = "BUGFIX"
	WindowsUpdateTypesEnhancement WindowsUpdateTypesEnum = "ENHANCEMENT"
	WindowsUpdateTypesOther       WindowsUpdateTypesEnum = "OTHER"
	WindowsUpdateTypesAll         WindowsUpdateTypesEnum = "ALL"
)

var mappingWindowsUpdateTypesEnum = map[string]WindowsUpdateTypesEnum{
	"SECURITY":    WindowsUpdateTypesSecurity,
	"BUGFIX":      WindowsUpdateTypesBugfix,
	"ENHANCEMENT": WindowsUpdateTypesEnhancement,
	"OTHER":       WindowsUpdateTypesOther,
	"ALL":         WindowsUpdateTypesAll,
}

var mappingWindowsUpdateTypesEnumLowerCase = map[string]WindowsUpdateTypesEnum{
	"security":    WindowsUpdateTypesSecurity,
	"bugfix":      WindowsUpdateTypesBugfix,
	"enhancement": WindowsUpdateTypesEnhancement,
	"other":       WindowsUpdateTypesOther,
	"all":         WindowsUpdateTypesAll,
}

// GetWindowsUpdateTypesEnumValues Enumerates the set of values for WindowsUpdateTypesEnum
func GetWindowsUpdateTypesEnumValues() []WindowsUpdateTypesEnum {
	values := make([]WindowsUpdateTypesEnum, 0)
	for _, v := range mappingWindowsUpdateTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetWindowsUpdateTypesEnumStringValues Enumerates the set of values in String for WindowsUpdateTypesEnum
func GetWindowsUpdateTypesEnumStringValues() []string {
	return []string{
		"SECURITY",
		"BUGFIX",
		"ENHANCEMENT",
		"OTHER",
		"ALL",
	}
}

// GetMappingWindowsUpdateTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWindowsUpdateTypesEnum(val string) (WindowsUpdateTypesEnum, bool) {
	enum, ok := mappingWindowsUpdateTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
