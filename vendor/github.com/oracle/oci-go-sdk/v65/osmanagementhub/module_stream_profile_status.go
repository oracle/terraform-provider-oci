// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// ModuleStreamProfileStatusEnum Enum with underlying type: string
type ModuleStreamProfileStatusEnum string

// Set of constants representing the allowable values for ModuleStreamProfileStatusEnum
const (
	ModuleStreamProfileStatusInstalled ModuleStreamProfileStatusEnum = "INSTALLED"
	ModuleStreamProfileStatusAvailable ModuleStreamProfileStatusEnum = "AVAILABLE"
)

var mappingModuleStreamProfileStatusEnum = map[string]ModuleStreamProfileStatusEnum{
	"INSTALLED": ModuleStreamProfileStatusInstalled,
	"AVAILABLE": ModuleStreamProfileStatusAvailable,
}

var mappingModuleStreamProfileStatusEnumLowerCase = map[string]ModuleStreamProfileStatusEnum{
	"installed": ModuleStreamProfileStatusInstalled,
	"available": ModuleStreamProfileStatusAvailable,
}

// GetModuleStreamProfileStatusEnumValues Enumerates the set of values for ModuleStreamProfileStatusEnum
func GetModuleStreamProfileStatusEnumValues() []ModuleStreamProfileStatusEnum {
	values := make([]ModuleStreamProfileStatusEnum, 0)
	for _, v := range mappingModuleStreamProfileStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetModuleStreamProfileStatusEnumStringValues Enumerates the set of values in String for ModuleStreamProfileStatusEnum
func GetModuleStreamProfileStatusEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"AVAILABLE",
	}
}

// GetMappingModuleStreamProfileStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModuleStreamProfileStatusEnum(val string) (ModuleStreamProfileStatusEnum, bool) {
	enum, ok := mappingModuleStreamProfileStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
