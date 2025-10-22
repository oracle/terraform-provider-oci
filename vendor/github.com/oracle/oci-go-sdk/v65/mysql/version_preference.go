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

// VersionPreferenceEnum Enum with underlying type: string
type VersionPreferenceEnum string

// Set of constants representing the allowable values for VersionPreferenceEnum
const (
	VersionPreferenceOldest       VersionPreferenceEnum = "OLDEST"
	VersionPreferenceSecondNewest VersionPreferenceEnum = "SECOND_NEWEST"
	VersionPreferenceNewest       VersionPreferenceEnum = "NEWEST"
)

var mappingVersionPreferenceEnum = map[string]VersionPreferenceEnum{
	"OLDEST":        VersionPreferenceOldest,
	"SECOND_NEWEST": VersionPreferenceSecondNewest,
	"NEWEST":        VersionPreferenceNewest,
}

var mappingVersionPreferenceEnumLowerCase = map[string]VersionPreferenceEnum{
	"oldest":        VersionPreferenceOldest,
	"second_newest": VersionPreferenceSecondNewest,
	"newest":        VersionPreferenceNewest,
}

// GetVersionPreferenceEnumValues Enumerates the set of values for VersionPreferenceEnum
func GetVersionPreferenceEnumValues() []VersionPreferenceEnum {
	values := make([]VersionPreferenceEnum, 0)
	for _, v := range mappingVersionPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetVersionPreferenceEnumStringValues Enumerates the set of values in String for VersionPreferenceEnum
func GetVersionPreferenceEnumStringValues() []string {
	return []string{
		"OLDEST",
		"SECOND_NEWEST",
		"NEWEST",
	}
}

// GetMappingVersionPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVersionPreferenceEnum(val string) (VersionPreferenceEnum, bool) {
	enum, ok := mappingVersionPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
