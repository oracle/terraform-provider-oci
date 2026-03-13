// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbStorageManagementEnum Enum with underlying type: string
type BaseDbStorageManagementEnum string

// Set of constants representing the allowable values for BaseDbStorageManagementEnum
const (
	BaseDbStorageManagementAsm BaseDbStorageManagementEnum = "ASM"
	BaseDbStorageManagementLvm BaseDbStorageManagementEnum = "LVM"
)

var mappingBaseDbStorageManagementEnum = map[string]BaseDbStorageManagementEnum{
	"ASM": BaseDbStorageManagementAsm,
	"LVM": BaseDbStorageManagementLvm,
}

var mappingBaseDbStorageManagementEnumLowerCase = map[string]BaseDbStorageManagementEnum{
	"asm": BaseDbStorageManagementAsm,
	"lvm": BaseDbStorageManagementLvm,
}

// GetBaseDbStorageManagementEnumValues Enumerates the set of values for BaseDbStorageManagementEnum
func GetBaseDbStorageManagementEnumValues() []BaseDbStorageManagementEnum {
	values := make([]BaseDbStorageManagementEnum, 0)
	for _, v := range mappingBaseDbStorageManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbStorageManagementEnumStringValues Enumerates the set of values in String for BaseDbStorageManagementEnum
func GetBaseDbStorageManagementEnumStringValues() []string {
	return []string{
		"ASM",
		"LVM",
	}
}

// GetMappingBaseDbStorageManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbStorageManagementEnum(val string) (BaseDbStorageManagementEnum, bool) {
	enum, ok := mappingBaseDbStorageManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
