// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SnapModesEnum Enum with underlying type: string
type SnapModesEnum string

// Set of constants representing the allowable values for SnapModesEnum
const (
	SnapModesDefault  SnapModesEnum = "DEFAULT"
	SnapModesClassic  SnapModesEnum = "CLASSIC"
	SnapModesJailMode SnapModesEnum = "JAIL_MODE"
	SnapModesDevMode  SnapModesEnum = "DEV_MODE"
)

var mappingSnapModesEnum = map[string]SnapModesEnum{
	"DEFAULT":   SnapModesDefault,
	"CLASSIC":   SnapModesClassic,
	"JAIL_MODE": SnapModesJailMode,
	"DEV_MODE":  SnapModesDevMode,
}

var mappingSnapModesEnumLowerCase = map[string]SnapModesEnum{
	"default":   SnapModesDefault,
	"classic":   SnapModesClassic,
	"jail_mode": SnapModesJailMode,
	"dev_mode":  SnapModesDevMode,
}

// GetSnapModesEnumValues Enumerates the set of values for SnapModesEnum
func GetSnapModesEnumValues() []SnapModesEnum {
	values := make([]SnapModesEnum, 0)
	for _, v := range mappingSnapModesEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapModesEnumStringValues Enumerates the set of values in String for SnapModesEnum
func GetSnapModesEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CLASSIC",
		"JAIL_MODE",
		"DEV_MODE",
	}
}

// GetMappingSnapModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapModesEnum(val string) (SnapModesEnum, bool) {
	enum, ok := mappingSnapModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
