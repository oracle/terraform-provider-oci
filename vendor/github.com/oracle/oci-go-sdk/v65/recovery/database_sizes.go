// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// DatabaseSizesEnum Enum with underlying type: string
type DatabaseSizesEnum string

// Set of constants representing the allowable values for DatabaseSizesEnum
const (
	DatabaseSizesXs   DatabaseSizesEnum = "XS"
	DatabaseSizesS    DatabaseSizesEnum = "S"
	DatabaseSizesM    DatabaseSizesEnum = "M"
	DatabaseSizesL    DatabaseSizesEnum = "L"
	DatabaseSizesXl   DatabaseSizesEnum = "XL"
	DatabaseSizesXxl  DatabaseSizesEnum = "XXL"
	DatabaseSizesAuto DatabaseSizesEnum = "AUTO"
)

var mappingDatabaseSizesEnum = map[string]DatabaseSizesEnum{
	"XS":   DatabaseSizesXs,
	"S":    DatabaseSizesS,
	"M":    DatabaseSizesM,
	"L":    DatabaseSizesL,
	"XL":   DatabaseSizesXl,
	"XXL":  DatabaseSizesXxl,
	"AUTO": DatabaseSizesAuto,
}

var mappingDatabaseSizesEnumLowerCase = map[string]DatabaseSizesEnum{
	"xs":   DatabaseSizesXs,
	"s":    DatabaseSizesS,
	"m":    DatabaseSizesM,
	"l":    DatabaseSizesL,
	"xl":   DatabaseSizesXl,
	"xxl":  DatabaseSizesXxl,
	"auto": DatabaseSizesAuto,
}

// GetDatabaseSizesEnumValues Enumerates the set of values for DatabaseSizesEnum
func GetDatabaseSizesEnumValues() []DatabaseSizesEnum {
	values := make([]DatabaseSizesEnum, 0)
	for _, v := range mappingDatabaseSizesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSizesEnumStringValues Enumerates the set of values in String for DatabaseSizesEnum
func GetDatabaseSizesEnumStringValues() []string {
	return []string{
		"XS",
		"S",
		"M",
		"L",
		"XL",
		"XXL",
		"AUTO",
	}
}

// GetMappingDatabaseSizesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSizesEnum(val string) (DatabaseSizesEnum, bool) {
	enum, ok := mappingDatabaseSizesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
