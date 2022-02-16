// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreatePrivateApplication OperationTypeEnum = "CREATE_PRIVATE_APPLICATION"
	OperationTypeUpdatePrivateApplication OperationTypeEnum = "UPDATE_PRIVATE_APPLICATION"
	OperationTypeDeletePrivateApplication OperationTypeEnum = "DELETE_PRIVATE_APPLICATION"
	OperationTypeMovePrivateApplication   OperationTypeEnum = "MOVE_PRIVATE_APPLICATION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PRIVATE_APPLICATION": OperationTypeCreatePrivateApplication,
	"UPDATE_PRIVATE_APPLICATION": OperationTypeUpdatePrivateApplication,
	"DELETE_PRIVATE_APPLICATION": OperationTypeDeletePrivateApplication,
	"MOVE_PRIVATE_APPLICATION":   OperationTypeMovePrivateApplication,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_PRIVATE_APPLICATION",
		"UPDATE_PRIVATE_APPLICATION",
		"DELETE_PRIVATE_APPLICATION",
		"MOVE_PRIVATE_APPLICATION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	mappingOperationTypeEnumIgnoreCase := make(map[string]OperationTypeEnum)
	for k, v := range mappingOperationTypeEnum {
		mappingOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
