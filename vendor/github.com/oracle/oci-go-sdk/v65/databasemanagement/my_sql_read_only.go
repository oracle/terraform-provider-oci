// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlReadOnlyEnum Enum with underlying type: string
type MySqlReadOnlyEnum string

// Set of constants representing the allowable values for MySqlReadOnlyEnum
const (
	MySqlReadOnlyOn    MySqlReadOnlyEnum = "ON"
	MySqlReadOnlySuper MySqlReadOnlyEnum = "SUPER"
	MySqlReadOnlyOff   MySqlReadOnlyEnum = "OFF"
)

var mappingMySqlReadOnlyEnum = map[string]MySqlReadOnlyEnum{
	"ON":    MySqlReadOnlyOn,
	"SUPER": MySqlReadOnlySuper,
	"OFF":   MySqlReadOnlyOff,
}

var mappingMySqlReadOnlyEnumLowerCase = map[string]MySqlReadOnlyEnum{
	"on":    MySqlReadOnlyOn,
	"super": MySqlReadOnlySuper,
	"off":   MySqlReadOnlyOff,
}

// GetMySqlReadOnlyEnumValues Enumerates the set of values for MySqlReadOnlyEnum
func GetMySqlReadOnlyEnumValues() []MySqlReadOnlyEnum {
	values := make([]MySqlReadOnlyEnum, 0)
	for _, v := range mappingMySqlReadOnlyEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlReadOnlyEnumStringValues Enumerates the set of values in String for MySqlReadOnlyEnum
func GetMySqlReadOnlyEnumStringValues() []string {
	return []string{
		"ON",
		"SUPER",
		"OFF",
	}
}

// GetMappingMySqlReadOnlyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlReadOnlyEnum(val string) (MySqlReadOnlyEnum, bool) {
	enum, ok := mappingMySqlReadOnlyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
