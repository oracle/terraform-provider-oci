// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// SoftDeleteEnum Enum with underlying type: string
type SoftDeleteEnum string

// Set of constants representing the allowable values for SoftDeleteEnum
const (
	SoftDeleteEnabled  SoftDeleteEnum = "ENABLED"
	SoftDeleteDisabled SoftDeleteEnum = "DISABLED"
)

var mappingSoftDeleteEnum = map[string]SoftDeleteEnum{
	"ENABLED":  SoftDeleteEnabled,
	"DISABLED": SoftDeleteDisabled,
}

var mappingSoftDeleteEnumLowerCase = map[string]SoftDeleteEnum{
	"enabled":  SoftDeleteEnabled,
	"disabled": SoftDeleteDisabled,
}

// GetSoftDeleteEnumValues Enumerates the set of values for SoftDeleteEnum
func GetSoftDeleteEnumValues() []SoftDeleteEnum {
	values := make([]SoftDeleteEnum, 0)
	for _, v := range mappingSoftDeleteEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftDeleteEnumStringValues Enumerates the set of values in String for SoftDeleteEnum
func GetSoftDeleteEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSoftDeleteEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftDeleteEnum(val string) (SoftDeleteEnum, bool) {
	enum, ok := mappingSoftDeleteEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
