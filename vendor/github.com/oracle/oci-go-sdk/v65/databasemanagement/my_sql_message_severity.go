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

// MySqlMessageSeverityEnum Enum with underlying type: string
type MySqlMessageSeverityEnum string

// Set of constants representing the allowable values for MySqlMessageSeverityEnum
const (
	MySqlMessageSeverityError   MySqlMessageSeverityEnum = "ERROR"
	MySqlMessageSeverityWarning MySqlMessageSeverityEnum = "WARNING"
	MySqlMessageSeverityNote    MySqlMessageSeverityEnum = "NOTE"
)

var mappingMySqlMessageSeverityEnum = map[string]MySqlMessageSeverityEnum{
	"ERROR":   MySqlMessageSeverityError,
	"WARNING": MySqlMessageSeverityWarning,
	"NOTE":    MySqlMessageSeverityNote,
}

var mappingMySqlMessageSeverityEnumLowerCase = map[string]MySqlMessageSeverityEnum{
	"error":   MySqlMessageSeverityError,
	"warning": MySqlMessageSeverityWarning,
	"note":    MySqlMessageSeverityNote,
}

// GetMySqlMessageSeverityEnumValues Enumerates the set of values for MySqlMessageSeverityEnum
func GetMySqlMessageSeverityEnumValues() []MySqlMessageSeverityEnum {
	values := make([]MySqlMessageSeverityEnum, 0)
	for _, v := range mappingMySqlMessageSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlMessageSeverityEnumStringValues Enumerates the set of values in String for MySqlMessageSeverityEnum
func GetMySqlMessageSeverityEnumStringValues() []string {
	return []string{
		"ERROR",
		"WARNING",
		"NOTE",
	}
}

// GetMappingMySqlMessageSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlMessageSeverityEnum(val string) (MySqlMessageSeverityEnum, bool) {
	enum, ok := mappingMySqlMessageSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
