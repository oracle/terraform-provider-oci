// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// RelatedResourceEntityTypeEnum Enum with underlying type: string
type RelatedResourceEntityTypeEnum string

// Set of constants representing the allowable values for RelatedResourceEntityTypeEnum
const (
	RelatedResourceEntityTypeAutonomousdatabase RelatedResourceEntityTypeEnum = "AUTONOMOUSDATABASE"
	RelatedResourceEntityTypeDatabase           RelatedResourceEntityTypeEnum = "DATABASE"
	RelatedResourceEntityTypePluggabledatabase  RelatedResourceEntityTypeEnum = "PLUGGABLEDATABASE"
)

var mappingRelatedResourceEntityTypeEnum = map[string]RelatedResourceEntityTypeEnum{
	"AUTONOMOUSDATABASE": RelatedResourceEntityTypeAutonomousdatabase,
	"DATABASE":           RelatedResourceEntityTypeDatabase,
	"PLUGGABLEDATABASE":  RelatedResourceEntityTypePluggabledatabase,
}

var mappingRelatedResourceEntityTypeEnumLowerCase = map[string]RelatedResourceEntityTypeEnum{
	"autonomousdatabase": RelatedResourceEntityTypeAutonomousdatabase,
	"database":           RelatedResourceEntityTypeDatabase,
	"pluggabledatabase":  RelatedResourceEntityTypePluggabledatabase,
}

// GetRelatedResourceEntityTypeEnumValues Enumerates the set of values for RelatedResourceEntityTypeEnum
func GetRelatedResourceEntityTypeEnumValues() []RelatedResourceEntityTypeEnum {
	values := make([]RelatedResourceEntityTypeEnum, 0)
	for _, v := range mappingRelatedResourceEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedResourceEntityTypeEnumStringValues Enumerates the set of values in String for RelatedResourceEntityTypeEnum
func GetRelatedResourceEntityTypeEnumStringValues() []string {
	return []string{
		"AUTONOMOUSDATABASE",
		"DATABASE",
		"PLUGGABLEDATABASE",
	}
}

// GetMappingRelatedResourceEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedResourceEntityTypeEnum(val string) (RelatedResourceEntityTypeEnum, bool) {
	enum, ok := mappingRelatedResourceEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
