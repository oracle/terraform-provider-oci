// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ArchTypesEnum Enum with underlying type: string
type ArchTypesEnum string

// Set of constants representing the allowable values for ArchTypesEnum
const (
	ArchTypesIa32        ArchTypesEnum = "IA_32"
	ArchTypesX8664       ArchTypesEnum = "X86_64"
	ArchTypesAarch64     ArchTypesEnum = "AARCH64"
	ArchTypesSparc       ArchTypesEnum = "SPARC"
	ArchTypesAmd64Debian ArchTypesEnum = "AMD64_DEBIAN"
)

var mappingArchTypesEnum = map[string]ArchTypesEnum{
	"IA_32":        ArchTypesIa32,
	"X86_64":       ArchTypesX8664,
	"AARCH64":      ArchTypesAarch64,
	"SPARC":        ArchTypesSparc,
	"AMD64_DEBIAN": ArchTypesAmd64Debian,
}

var mappingArchTypesEnumLowerCase = map[string]ArchTypesEnum{
	"ia_32":        ArchTypesIa32,
	"x86_64":       ArchTypesX8664,
	"aarch64":      ArchTypesAarch64,
	"sparc":        ArchTypesSparc,
	"amd64_debian": ArchTypesAmd64Debian,
}

// GetArchTypesEnumValues Enumerates the set of values for ArchTypesEnum
func GetArchTypesEnumValues() []ArchTypesEnum {
	values := make([]ArchTypesEnum, 0)
	for _, v := range mappingArchTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetArchTypesEnumStringValues Enumerates the set of values in String for ArchTypesEnum
func GetArchTypesEnumStringValues() []string {
	return []string{
		"IA_32",
		"X86_64",
		"AARCH64",
		"SPARC",
		"AMD64_DEBIAN",
	}
}

// GetMappingArchTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchTypesEnum(val string) (ArchTypesEnum, bool) {
	enum, ok := mappingArchTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
