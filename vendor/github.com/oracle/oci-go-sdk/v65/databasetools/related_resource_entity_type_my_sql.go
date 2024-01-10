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

// RelatedResourceEntityTypeMySqlEnum Enum with underlying type: string
type RelatedResourceEntityTypeMySqlEnum string

// Set of constants representing the allowable values for RelatedResourceEntityTypeMySqlEnum
const (
	RelatedResourceEntityTypeMySqlMysqldbsystem RelatedResourceEntityTypeMySqlEnum = "MYSQLDBSYSTEM"
)

var mappingRelatedResourceEntityTypeMySqlEnum = map[string]RelatedResourceEntityTypeMySqlEnum{
	"MYSQLDBSYSTEM": RelatedResourceEntityTypeMySqlMysqldbsystem,
}

var mappingRelatedResourceEntityTypeMySqlEnumLowerCase = map[string]RelatedResourceEntityTypeMySqlEnum{
	"mysqldbsystem": RelatedResourceEntityTypeMySqlMysqldbsystem,
}

// GetRelatedResourceEntityTypeMySqlEnumValues Enumerates the set of values for RelatedResourceEntityTypeMySqlEnum
func GetRelatedResourceEntityTypeMySqlEnumValues() []RelatedResourceEntityTypeMySqlEnum {
	values := make([]RelatedResourceEntityTypeMySqlEnum, 0)
	for _, v := range mappingRelatedResourceEntityTypeMySqlEnum {
		values = append(values, v)
	}
	return values
}

// GetRelatedResourceEntityTypeMySqlEnumStringValues Enumerates the set of values in String for RelatedResourceEntityTypeMySqlEnum
func GetRelatedResourceEntityTypeMySqlEnumStringValues() []string {
	return []string{
		"MYSQLDBSYSTEM",
	}
}

// GetMappingRelatedResourceEntityTypeMySqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelatedResourceEntityTypeMySqlEnum(val string) (RelatedResourceEntityTypeMySqlEnum, bool) {
	enum, ok := mappingRelatedResourceEntityTypeMySqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
