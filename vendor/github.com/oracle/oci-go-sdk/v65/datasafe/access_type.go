// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AccessTypeEnum Enum with underlying type: string
type AccessTypeEnum string

// Set of constants representing the allowable values for AccessTypeEnum
const (
	AccessTypeSelect AccessTypeEnum = "SELECT"
	AccessTypeUpdate AccessTypeEnum = "UPDATE"
	AccessTypeInsert AccessTypeEnum = "INSERT"
	AccessTypeDelete AccessTypeEnum = "DELETE"
	AccessTypeOwner  AccessTypeEnum = "OWNER"
)

var mappingAccessTypeEnum = map[string]AccessTypeEnum{
	"SELECT": AccessTypeSelect,
	"UPDATE": AccessTypeUpdate,
	"INSERT": AccessTypeInsert,
	"DELETE": AccessTypeDelete,
	"OWNER":  AccessTypeOwner,
}

var mappingAccessTypeEnumLowerCase = map[string]AccessTypeEnum{
	"select": AccessTypeSelect,
	"update": AccessTypeUpdate,
	"insert": AccessTypeInsert,
	"delete": AccessTypeDelete,
	"owner":  AccessTypeOwner,
}

// GetAccessTypeEnumValues Enumerates the set of values for AccessTypeEnum
func GetAccessTypeEnumValues() []AccessTypeEnum {
	values := make([]AccessTypeEnum, 0)
	for _, v := range mappingAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessTypeEnumStringValues Enumerates the set of values in String for AccessTypeEnum
func GetAccessTypeEnumStringValues() []string {
	return []string{
		"SELECT",
		"UPDATE",
		"INSERT",
		"DELETE",
		"OWNER",
	}
}

// GetMappingAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessTypeEnum(val string) (AccessTypeEnum, bool) {
	enum, ok := mappingAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
