// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// MySqlCredTypeEnum Enum with underlying type: string
type MySqlCredTypeEnum string

// Set of constants representing the allowable values for MySqlCredTypeEnum
const (
	MySqlCredTypeMysqlExternalNonSslCredentials MySqlCredTypeEnum = "MYSQL_EXTERNAL_NON_SSL_CREDENTIALS"
	MySqlCredTypeMysqlExternalSslCredentials    MySqlCredTypeEnum = "MYSQL_EXTERNAL_SSL_CREDENTIALS"
	MySqlCredTypeMysqlExternalSocketCredentials MySqlCredTypeEnum = "MYSQL_EXTERNAL_SOCKET_CREDENTIALS"
)

var mappingMySqlCredTypeEnum = map[string]MySqlCredTypeEnum{
	"MYSQL_EXTERNAL_NON_SSL_CREDENTIALS": MySqlCredTypeMysqlExternalNonSslCredentials,
	"MYSQL_EXTERNAL_SSL_CREDENTIALS":     MySqlCredTypeMysqlExternalSslCredentials,
	"MYSQL_EXTERNAL_SOCKET_CREDENTIALS":  MySqlCredTypeMysqlExternalSocketCredentials,
}

var mappingMySqlCredTypeEnumLowerCase = map[string]MySqlCredTypeEnum{
	"mysql_external_non_ssl_credentials": MySqlCredTypeMysqlExternalNonSslCredentials,
	"mysql_external_ssl_credentials":     MySqlCredTypeMysqlExternalSslCredentials,
	"mysql_external_socket_credentials":  MySqlCredTypeMysqlExternalSocketCredentials,
}

// GetMySqlCredTypeEnumValues Enumerates the set of values for MySqlCredTypeEnum
func GetMySqlCredTypeEnumValues() []MySqlCredTypeEnum {
	values := make([]MySqlCredTypeEnum, 0)
	for _, v := range mappingMySqlCredTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlCredTypeEnumStringValues Enumerates the set of values in String for MySqlCredTypeEnum
func GetMySqlCredTypeEnumStringValues() []string {
	return []string{
		"MYSQL_EXTERNAL_NON_SSL_CREDENTIALS",
		"MYSQL_EXTERNAL_SSL_CREDENTIALS",
		"MYSQL_EXTERNAL_SOCKET_CREDENTIALS",
	}
}

// GetMappingMySqlCredTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlCredTypeEnum(val string) (MySqlCredTypeEnum, bool) {
	enum, ok := mappingMySqlCredTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
