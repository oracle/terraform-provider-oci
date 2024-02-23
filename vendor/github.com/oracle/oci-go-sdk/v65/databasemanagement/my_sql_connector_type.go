// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// MySqlConnectorTypeEnum Enum with underlying type: string
type MySqlConnectorTypeEnum string

// Set of constants representing the allowable values for MySqlConnectorTypeEnum
const (
	MySqlConnectorTypeMacs MySqlConnectorTypeEnum = "MACS"
)

var mappingMySqlConnectorTypeEnum = map[string]MySqlConnectorTypeEnum{
	"MACS": MySqlConnectorTypeMacs,
}

var mappingMySqlConnectorTypeEnumLowerCase = map[string]MySqlConnectorTypeEnum{
	"macs": MySqlConnectorTypeMacs,
}

// GetMySqlConnectorTypeEnumValues Enumerates the set of values for MySqlConnectorTypeEnum
func GetMySqlConnectorTypeEnumValues() []MySqlConnectorTypeEnum {
	values := make([]MySqlConnectorTypeEnum, 0)
	for _, v := range mappingMySqlConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlConnectorTypeEnumStringValues Enumerates the set of values in String for MySqlConnectorTypeEnum
func GetMySqlConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingMySqlConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlConnectorTypeEnum(val string) (MySqlConnectorTypeEnum, bool) {
	enum, ok := mappingMySqlConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
