// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// DatabaseApiGatewayConfigGlobalKeyEnum Enum with underlying type: string
type DatabaseApiGatewayConfigGlobalKeyEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigGlobalKeyEnum
const (
	DatabaseApiGatewayConfigGlobalKeySettings DatabaseApiGatewayConfigGlobalKeyEnum = "SETTINGS"
)

var mappingDatabaseApiGatewayConfigGlobalKeyEnum = map[string]DatabaseApiGatewayConfigGlobalKeyEnum{
	"SETTINGS": DatabaseApiGatewayConfigGlobalKeySettings,
}

var mappingDatabaseApiGatewayConfigGlobalKeyEnumLowerCase = map[string]DatabaseApiGatewayConfigGlobalKeyEnum{
	"settings": DatabaseApiGatewayConfigGlobalKeySettings,
}

// GetDatabaseApiGatewayConfigGlobalKeyEnumValues Enumerates the set of values for DatabaseApiGatewayConfigGlobalKeyEnum
func GetDatabaseApiGatewayConfigGlobalKeyEnumValues() []DatabaseApiGatewayConfigGlobalKeyEnum {
	values := make([]DatabaseApiGatewayConfigGlobalKeyEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigGlobalKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigGlobalKeyEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigGlobalKeyEnum
func GetDatabaseApiGatewayConfigGlobalKeyEnumStringValues() []string {
	return []string{
		"SETTINGS",
	}
}

// GetMappingDatabaseApiGatewayConfigGlobalKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigGlobalKeyEnum(val string) (DatabaseApiGatewayConfigGlobalKeyEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigGlobalKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
