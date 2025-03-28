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

// MySqlNetworkProtocolTypeEnum Enum with underlying type: string
type MySqlNetworkProtocolTypeEnum string

// Set of constants representing the allowable values for MySqlNetworkProtocolTypeEnum
const (
	MySqlNetworkProtocolTypeTcp     MySqlNetworkProtocolTypeEnum = "TCP"
	MySqlNetworkProtocolTypeTcps    MySqlNetworkProtocolTypeEnum = "TCPS"
	MySqlNetworkProtocolTypeSockets MySqlNetworkProtocolTypeEnum = "SOCKETS"
)

var mappingMySqlNetworkProtocolTypeEnum = map[string]MySqlNetworkProtocolTypeEnum{
	"TCP":     MySqlNetworkProtocolTypeTcp,
	"TCPS":    MySqlNetworkProtocolTypeTcps,
	"SOCKETS": MySqlNetworkProtocolTypeSockets,
}

var mappingMySqlNetworkProtocolTypeEnumLowerCase = map[string]MySqlNetworkProtocolTypeEnum{
	"tcp":     MySqlNetworkProtocolTypeTcp,
	"tcps":    MySqlNetworkProtocolTypeTcps,
	"sockets": MySqlNetworkProtocolTypeSockets,
}

// GetMySqlNetworkProtocolTypeEnumValues Enumerates the set of values for MySqlNetworkProtocolTypeEnum
func GetMySqlNetworkProtocolTypeEnumValues() []MySqlNetworkProtocolTypeEnum {
	values := make([]MySqlNetworkProtocolTypeEnum, 0)
	for _, v := range mappingMySqlNetworkProtocolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlNetworkProtocolTypeEnumStringValues Enumerates the set of values in String for MySqlNetworkProtocolTypeEnum
func GetMySqlNetworkProtocolTypeEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
		"SOCKETS",
	}
}

// GetMappingMySqlNetworkProtocolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlNetworkProtocolTypeEnum(val string) (MySqlNetworkProtocolTypeEnum, bool) {
	enum, ok := mappingMySqlNetworkProtocolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
