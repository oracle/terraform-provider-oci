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

// DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum Enum with underlying type: string
type DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum
const (
	DatabaseApiGatewayConfigAdvancedPropertyConfigTypeGlobal DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum = "GLOBAL"
	DatabaseApiGatewayConfigAdvancedPropertyConfigTypePool   DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum = "POOL"
)

var mappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum = map[string]DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum{
	"GLOBAL": DatabaseApiGatewayConfigAdvancedPropertyConfigTypeGlobal,
	"POOL":   DatabaseApiGatewayConfigAdvancedPropertyConfigTypePool,
}

var mappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumLowerCase = map[string]DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum{
	"global": DatabaseApiGatewayConfigAdvancedPropertyConfigTypeGlobal,
	"pool":   DatabaseApiGatewayConfigAdvancedPropertyConfigTypePool,
}

// GetDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumValues Enumerates the set of values for DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum
func GetDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumValues() []DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum {
	values := make([]DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum
func GetDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"POOL",
	}
}

// GetMappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum(val string) (DatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigAdvancedPropertyConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
