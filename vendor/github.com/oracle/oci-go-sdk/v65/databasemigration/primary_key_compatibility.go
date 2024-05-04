// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// PrimaryKeyCompatibilityEnum Enum with underlying type: string
type PrimaryKeyCompatibilityEnum string

// Set of constants representing the allowable values for PrimaryKeyCompatibilityEnum
const (
	PrimaryKeyCompatibilityNone               PrimaryKeyCompatibilityEnum = "NONE"
	PrimaryKeyCompatibilityIgnoreMissingPks   PrimaryKeyCompatibilityEnum = "IGNORE_MISSING_PKS"
	PrimaryKeyCompatibilityCreateInvisiblePks PrimaryKeyCompatibilityEnum = "CREATE_INVISIBLE_PKS"
)

var mappingPrimaryKeyCompatibilityEnum = map[string]PrimaryKeyCompatibilityEnum{
	"NONE":                 PrimaryKeyCompatibilityNone,
	"IGNORE_MISSING_PKS":   PrimaryKeyCompatibilityIgnoreMissingPks,
	"CREATE_INVISIBLE_PKS": PrimaryKeyCompatibilityCreateInvisiblePks,
}

var mappingPrimaryKeyCompatibilityEnumLowerCase = map[string]PrimaryKeyCompatibilityEnum{
	"none":                 PrimaryKeyCompatibilityNone,
	"ignore_missing_pks":   PrimaryKeyCompatibilityIgnoreMissingPks,
	"create_invisible_pks": PrimaryKeyCompatibilityCreateInvisiblePks,
}

// GetPrimaryKeyCompatibilityEnumValues Enumerates the set of values for PrimaryKeyCompatibilityEnum
func GetPrimaryKeyCompatibilityEnumValues() []PrimaryKeyCompatibilityEnum {
	values := make([]PrimaryKeyCompatibilityEnum, 0)
	for _, v := range mappingPrimaryKeyCompatibilityEnum {
		values = append(values, v)
	}
	return values
}

// GetPrimaryKeyCompatibilityEnumStringValues Enumerates the set of values in String for PrimaryKeyCompatibilityEnum
func GetPrimaryKeyCompatibilityEnumStringValues() []string {
	return []string{
		"NONE",
		"IGNORE_MISSING_PKS",
		"CREATE_INVISIBLE_PKS",
	}
}

// GetMappingPrimaryKeyCompatibilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrimaryKeyCompatibilityEnum(val string) (PrimaryKeyCompatibilityEnum, bool) {
	enum, ok := mappingPrimaryKeyCompatibilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
