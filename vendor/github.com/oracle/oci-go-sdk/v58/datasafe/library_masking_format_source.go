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

// LibraryMaskingFormatSourceEnum Enum with underlying type: string
type LibraryMaskingFormatSourceEnum string

// Set of constants representing the allowable values for LibraryMaskingFormatSourceEnum
const (
	LibraryMaskingFormatSourceOracle LibraryMaskingFormatSourceEnum = "ORACLE"
	LibraryMaskingFormatSourceUser   LibraryMaskingFormatSourceEnum = "USER"
)

var mappingLibraryMaskingFormatSourceEnum = map[string]LibraryMaskingFormatSourceEnum{
	"ORACLE": LibraryMaskingFormatSourceOracle,
	"USER":   LibraryMaskingFormatSourceUser,
}

// GetLibraryMaskingFormatSourceEnumValues Enumerates the set of values for LibraryMaskingFormatSourceEnum
func GetLibraryMaskingFormatSourceEnumValues() []LibraryMaskingFormatSourceEnum {
	values := make([]LibraryMaskingFormatSourceEnum, 0)
	for _, v := range mappingLibraryMaskingFormatSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryMaskingFormatSourceEnumStringValues Enumerates the set of values in String for LibraryMaskingFormatSourceEnum
func GetLibraryMaskingFormatSourceEnumStringValues() []string {
	return []string{
		"ORACLE",
		"USER",
	}
}

// GetMappingLibraryMaskingFormatSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryMaskingFormatSourceEnum(val string) (LibraryMaskingFormatSourceEnum, bool) {
	mappingLibraryMaskingFormatSourceEnumIgnoreCase := make(map[string]LibraryMaskingFormatSourceEnum)
	for k, v := range mappingLibraryMaskingFormatSourceEnum {
		mappingLibraryMaskingFormatSourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLibraryMaskingFormatSourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
