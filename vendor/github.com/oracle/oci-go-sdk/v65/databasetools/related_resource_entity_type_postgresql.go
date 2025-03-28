// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RelatedResourceEntityTypePostgresqlEnum Enum with underlying type: string
type RelatedResourceEntityTypePostgresqlEnum string

// Set of constants representing the allowable values for RelatedResourceEntityTypePostgresqlEnum
const (
	RelatedResourceEntityTypePostgresqlPostgresqldbsystem RelatedResourceEntityTypePostgresqlEnum = "POSTGRESQLDBSYSTEM"
)

var mappingRelatedResourceEntityTypePostgresqlEnum = map[string]RelatedResourceEntityTypePostgresqlEnum{
	"POSTGRESQLDBSYSTEM": RelatedResourceEntityTypePostgresqlPostgresqldbsystem,
}

var mappingRelatedResourceEntityTypePostgresqlEnumLowerCase = map[string]RelatedResourceEntityTypePostgresqlEnum{
	"postgresqldbsystem": RelatedResourceEntityTypePostgresqlPostgresqldbsystem,
}

// GetRelatedResourceEntityTypePostgresqlEnumValues Enumerates the set of values for RelatedResourceEntityTypePostgresqlEnum
func GetRelatedResourceEntityTypePostgresqlEnumValues() []RelatedResourceEntityTypePostgresqlEnum {
	values := make([]RelatedResourceEntityTypePostgresqlEnum, 0)
	for _, v := range mappingRelatedResourceEntityTypePostgresqlEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedResourceEntityTypePostgresqlEnumStringValues Enumerates the set of values in String for RelatedResourceEntityTypePostgresqlEnum
func GetRelatedResourceEntityTypePostgresqlEnumStringValues() []string {
	return []string{
		"POSTGRESQLDBSYSTEM",
	}
}

// GetMappingRelatedResourceEntityTypePostgresqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedResourceEntityTypePostgresqlEnum(val string) (RelatedResourceEntityTypePostgresqlEnum, bool) {
	enum, ok := mappingRelatedResourceEntityTypePostgresqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
