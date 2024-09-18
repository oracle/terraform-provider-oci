// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// DbSourceMajorVersionsEnum Enum with underlying type: string
type DbSourceMajorVersionsEnum string

// Set of constants representing the allowable values for DbSourceMajorVersionsEnum
const (
	DbSourceMajorVersionsDb11204 DbSourceMajorVersionsEnum = "DB_11204"
	DbSourceMajorVersionsDb121   DbSourceMajorVersionsEnum = "DB_121"
	DbSourceMajorVersionsDb122   DbSourceMajorVersionsEnum = "DB_122"
	DbSourceMajorVersionsDb18    DbSourceMajorVersionsEnum = "DB_18"
	DbSourceMajorVersionsDb19    DbSourceMajorVersionsEnum = "DB_19"
	DbSourceMajorVersionsDb23    DbSourceMajorVersionsEnum = "DB_23"
)

var mappingDbSourceMajorVersionsEnum = map[string]DbSourceMajorVersionsEnum{
	"DB_11204": DbSourceMajorVersionsDb11204,
	"DB_121":   DbSourceMajorVersionsDb121,
	"DB_122":   DbSourceMajorVersionsDb122,
	"DB_18":    DbSourceMajorVersionsDb18,
	"DB_19":    DbSourceMajorVersionsDb19,
	"DB_23":    DbSourceMajorVersionsDb23,
}

var mappingDbSourceMajorVersionsEnumLowerCase = map[string]DbSourceMajorVersionsEnum{
	"db_11204": DbSourceMajorVersionsDb11204,
	"db_121":   DbSourceMajorVersionsDb121,
	"db_122":   DbSourceMajorVersionsDb122,
	"db_18":    DbSourceMajorVersionsDb18,
	"db_19":    DbSourceMajorVersionsDb19,
	"db_23":    DbSourceMajorVersionsDb23,
}

// GetDbSourceMajorVersionsEnumValues Enumerates the set of values for DbSourceMajorVersionsEnum
func GetDbSourceMajorVersionsEnumValues() []DbSourceMajorVersionsEnum {
	values := make([]DbSourceMajorVersionsEnum, 0)
	for _, v := range mappingDbSourceMajorVersionsEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSourceMajorVersionsEnumStringValues Enumerates the set of values in String for DbSourceMajorVersionsEnum
func GetDbSourceMajorVersionsEnumStringValues() []string {
	return []string{
		"DB_11204",
		"DB_121",
		"DB_122",
		"DB_18",
		"DB_19",
		"DB_23",
	}
}

// GetMappingDbSourceMajorVersionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSourceMajorVersionsEnum(val string) (DbSourceMajorVersionsEnum, bool) {
	enum, ok := mappingDbSourceMajorVersionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
