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

// ModuleStreamStatusEnum Enum with underlying type: string
type ModuleStreamStatusEnum string

// Set of constants representing the allowable values for ModuleStreamStatusEnum
const (
	ModuleStreamStatusEnabled  ModuleStreamStatusEnum = "ENABLED"
	ModuleStreamStatusDisabled ModuleStreamStatusEnum = "DISABLED"
	ModuleStreamStatusActive   ModuleStreamStatusEnum = "ACTIVE"
)

var mappingModuleStreamStatusEnum = map[string]ModuleStreamStatusEnum{
	"ENABLED":  ModuleStreamStatusEnabled,
	"DISABLED": ModuleStreamStatusDisabled,
	"ACTIVE":   ModuleStreamStatusActive,
}

var mappingModuleStreamStatusEnumLowerCase = map[string]ModuleStreamStatusEnum{
	"enabled":  ModuleStreamStatusEnabled,
	"disabled": ModuleStreamStatusDisabled,
	"active":   ModuleStreamStatusActive,
}

// GetModuleStreamStatusEnumValues Enumerates the set of values for ModuleStreamStatusEnum
func GetModuleStreamStatusEnumValues() []ModuleStreamStatusEnum {
	values := make([]ModuleStreamStatusEnum, 0)
	for _, v := range mappingModuleStreamStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetModuleStreamStatusEnumStringValues Enumerates the set of values in String for ModuleStreamStatusEnum
func GetModuleStreamStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"ACTIVE",
	}
}

// GetMappingModuleStreamStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModuleStreamStatusEnum(val string) (ModuleStreamStatusEnum, bool) {
	enum, ok := mappingModuleStreamStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
