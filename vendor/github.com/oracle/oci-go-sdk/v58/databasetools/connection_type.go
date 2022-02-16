// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"strings"
)

// ConnectionTypeEnum Enum with underlying type: string
type ConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionTypeEnum
const (
	ConnectionTypeOracleDatabase ConnectionTypeEnum = "ORACLE_DATABASE"
)

var mappingConnectionTypeEnum = map[string]ConnectionTypeEnum{
	"ORACLE_DATABASE": ConnectionTypeOracleDatabase,
}

// GetConnectionTypeEnumValues Enumerates the set of values for ConnectionTypeEnum
func GetConnectionTypeEnumValues() []ConnectionTypeEnum {
	values := make([]ConnectionTypeEnum, 0)
	for _, v := range mappingConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionTypeEnum
func GetConnectionTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATABASE",
	}
}

// GetMappingConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionTypeEnum(val string) (ConnectionTypeEnum, bool) {
	mappingConnectionTypeEnumIgnoreCase := make(map[string]ConnectionTypeEnum)
	for k, v := range mappingConnectionTypeEnum {
		mappingConnectionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConnectionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
