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

// DatabaseApiGatewayConfigTypeEnum Enum with underlying type: string
type DatabaseApiGatewayConfigTypeEnum string

// Set of constants representing the allowable values for DatabaseApiGatewayConfigTypeEnum
const (
	DatabaseApiGatewayConfigTypeDefault DatabaseApiGatewayConfigTypeEnum = "DEFAULT"
)

var mappingDatabaseApiGatewayConfigTypeEnum = map[string]DatabaseApiGatewayConfigTypeEnum{
	"DEFAULT": DatabaseApiGatewayConfigTypeDefault,
}

var mappingDatabaseApiGatewayConfigTypeEnumLowerCase = map[string]DatabaseApiGatewayConfigTypeEnum{
	"default": DatabaseApiGatewayConfigTypeDefault,
}

// GetDatabaseApiGatewayConfigTypeEnumValues Enumerates the set of values for DatabaseApiGatewayConfigTypeEnum
func GetDatabaseApiGatewayConfigTypeEnumValues() []DatabaseApiGatewayConfigTypeEnum {
	values := make([]DatabaseApiGatewayConfigTypeEnum, 0)
	for _, v := range mappingDatabaseApiGatewayConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseApiGatewayConfigTypeEnumStringValues Enumerates the set of values in String for DatabaseApiGatewayConfigTypeEnum
func GetDatabaseApiGatewayConfigTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingDatabaseApiGatewayConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseApiGatewayConfigTypeEnum(val string) (DatabaseApiGatewayConfigTypeEnum, bool) {
	enum, ok := mappingDatabaseApiGatewayConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
