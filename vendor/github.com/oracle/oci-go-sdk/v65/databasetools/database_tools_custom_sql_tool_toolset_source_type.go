// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsCustomSqlToolToolsetSourceTypeEnum Enum with underlying type: string
type DatabaseToolsCustomSqlToolToolsetSourceTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsCustomSqlToolToolsetSourceTypeEnum
const (
	DatabaseToolsCustomSqlToolToolsetSourceTypeInline DatabaseToolsCustomSqlToolToolsetSourceTypeEnum = "INLINE"
)

var mappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnum = map[string]DatabaseToolsCustomSqlToolToolsetSourceTypeEnum{
	"INLINE": DatabaseToolsCustomSqlToolToolsetSourceTypeInline,
}

var mappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnumLowerCase = map[string]DatabaseToolsCustomSqlToolToolsetSourceTypeEnum{
	"inline": DatabaseToolsCustomSqlToolToolsetSourceTypeInline,
}

// GetDatabaseToolsCustomSqlToolToolsetSourceTypeEnumValues Enumerates the set of values for DatabaseToolsCustomSqlToolToolsetSourceTypeEnum
func GetDatabaseToolsCustomSqlToolToolsetSourceTypeEnumValues() []DatabaseToolsCustomSqlToolToolsetSourceTypeEnum {
	values := make([]DatabaseToolsCustomSqlToolToolsetSourceTypeEnum, 0)
	for _, v := range mappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsCustomSqlToolToolsetSourceTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsCustomSqlToolToolsetSourceTypeEnum
func GetDatabaseToolsCustomSqlToolToolsetSourceTypeEnumStringValues() []string {
	return []string{
		"INLINE",
	}
}

// GetMappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnum(val string) (DatabaseToolsCustomSqlToolToolsetSourceTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
