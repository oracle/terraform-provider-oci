// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// SensitiveTypeSourceEnum Enum with underlying type: string
type SensitiveTypeSourceEnum string

// Set of constants representing the allowable values for SensitiveTypeSourceEnum
const (
	SensitiveTypeSourceOracle SensitiveTypeSourceEnum = "ORACLE"
	SensitiveTypeSourceUser   SensitiveTypeSourceEnum = "USER"
)

var mappingSensitiveTypeSourceEnum = map[string]SensitiveTypeSourceEnum{
	"ORACLE": SensitiveTypeSourceOracle,
	"USER":   SensitiveTypeSourceUser,
}

// GetSensitiveTypeSourceEnumValues Enumerates the set of values for SensitiveTypeSourceEnum
func GetSensitiveTypeSourceEnumValues() []SensitiveTypeSourceEnum {
	values := make([]SensitiveTypeSourceEnum, 0)
	for _, v := range mappingSensitiveTypeSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveTypeSourceEnumStringValues Enumerates the set of values in String for SensitiveTypeSourceEnum
func GetSensitiveTypeSourceEnumStringValues() []string {
	return []string{
		"ORACLE",
		"USER",
	}
}

// GetMappingSensitiveTypeSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveTypeSourceEnum(val string) (SensitiveTypeSourceEnum, bool) {
	mappingSensitiveTypeSourceEnumIgnoreCase := make(map[string]SensitiveTypeSourceEnum)
	for k, v := range mappingSensitiveTypeSourceEnum {
		mappingSensitiveTypeSourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSensitiveTypeSourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
