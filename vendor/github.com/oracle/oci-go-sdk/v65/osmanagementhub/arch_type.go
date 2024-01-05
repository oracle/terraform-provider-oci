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

// ArchTypeEnum Enum with underlying type: string
type ArchTypeEnum string

// Set of constants representing the allowable values for ArchTypeEnum
const (
	ArchTypeX8664   ArchTypeEnum = "X86_64"
	ArchTypeAarch64 ArchTypeEnum = "AARCH64"
	ArchTypeI686    ArchTypeEnum = "I686"
	ArchTypeNoarch  ArchTypeEnum = "NOARCH"
	ArchTypeSrc     ArchTypeEnum = "SRC"
)

var mappingArchTypeEnum = map[string]ArchTypeEnum{
	"X86_64":  ArchTypeX8664,
	"AARCH64": ArchTypeAarch64,
	"I686":    ArchTypeI686,
	"NOARCH":  ArchTypeNoarch,
	"SRC":     ArchTypeSrc,
}

var mappingArchTypeEnumLowerCase = map[string]ArchTypeEnum{
	"x86_64":  ArchTypeX8664,
	"aarch64": ArchTypeAarch64,
	"i686":    ArchTypeI686,
	"noarch":  ArchTypeNoarch,
	"src":     ArchTypeSrc,
}

// GetArchTypeEnumValues Enumerates the set of values for ArchTypeEnum
func GetArchTypeEnumValues() []ArchTypeEnum {
	values := make([]ArchTypeEnum, 0)
	for _, v := range mappingArchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArchTypeEnumStringValues Enumerates the set of values in String for ArchTypeEnum
func GetArchTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"AARCH64",
		"I686",
		"NOARCH",
		"SRC",
	}
}

// GetMappingArchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchTypeEnum(val string) (ArchTypeEnum, bool) {
	enum, ok := mappingArchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
