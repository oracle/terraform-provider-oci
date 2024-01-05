// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// ArchitectureTypeEnum Enum with underlying type: string
type ArchitectureTypeEnum string

// Set of constants representing the allowable values for ArchitectureTypeEnum
const (
	ArchitectureTypeX86 ArchitectureTypeEnum = "X86"
	ArchitectureTypeArm ArchitectureTypeEnum = "ARM"
)

var mappingArchitectureTypeEnum = map[string]ArchitectureTypeEnum{
	"X86": ArchitectureTypeX86,
	"ARM": ArchitectureTypeArm,
}

var mappingArchitectureTypeEnumLowerCase = map[string]ArchitectureTypeEnum{
	"x86": ArchitectureTypeX86,
	"arm": ArchitectureTypeArm,
}

// GetArchitectureTypeEnumValues Enumerates the set of values for ArchitectureTypeEnum
func GetArchitectureTypeEnumValues() []ArchitectureTypeEnum {
	values := make([]ArchitectureTypeEnum, 0)
	for _, v := range mappingArchitectureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetArchitectureTypeEnumStringValues Enumerates the set of values in String for ArchitectureTypeEnum
func GetArchitectureTypeEnumStringValues() []string {
	return []string{
		"X86",
		"ARM",
	}
}

// GetMappingArchitectureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchitectureTypeEnum(val string) (ArchitectureTypeEnum, bool) {
	enum, ok := mappingArchitectureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
