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

// ConnectionTypeEnum Enum with underlying type: string
type ConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionTypeEnum
const (
	ConnectionTypeOracleDatabase ConnectionTypeEnum = "ORACLE_DATABASE"
	ConnectionTypeMysql          ConnectionTypeEnum = "MYSQL"
	ConnectionTypePostgresql     ConnectionTypeEnum = "POSTGRESQL"
	ConnectionTypeGenericJdbc    ConnectionTypeEnum = "GENERIC_JDBC"
)

var mappingConnectionTypeEnum = map[string]ConnectionTypeEnum{
	"ORACLE_DATABASE": ConnectionTypeOracleDatabase,
	"MYSQL":           ConnectionTypeMysql,
	"POSTGRESQL":      ConnectionTypePostgresql,
	"GENERIC_JDBC":    ConnectionTypeGenericJdbc,
}

var mappingConnectionTypeEnumLowerCase = map[string]ConnectionTypeEnum{
	"oracle_database": ConnectionTypeOracleDatabase,
	"mysql":           ConnectionTypeMysql,
	"postgresql":      ConnectionTypePostgresql,
	"generic_jdbc":    ConnectionTypeGenericJdbc,
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
		"MYSQL",
		"POSTGRESQL",
		"GENERIC_JDBC",
	}
}

// GetMappingConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionTypeEnum(val string) (ConnectionTypeEnum, bool) {
	enum, ok := mappingConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
