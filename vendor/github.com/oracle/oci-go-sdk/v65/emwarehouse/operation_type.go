// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// EM Warehouse API
//
// Use the EM Warehouse API to manage EM Warehouse data collection.
//

package emwarehouse

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateEmWarehouse OperationTypeEnum = "CREATE_EM_WAREHOUSE"
	OperationTypeUpdateEmWarehouse OperationTypeEnum = "UPDATE_EM_WAREHOUSE"
	OperationTypeDeleteEmWarehouse OperationTypeEnum = "DELETE_EM_WAREHOUSE"
	OperationTypeMoveEmWarehouse   OperationTypeEnum = "MOVE_EM_WAREHOUSE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_EM_WAREHOUSE": OperationTypeCreateEmWarehouse,
	"UPDATE_EM_WAREHOUSE": OperationTypeUpdateEmWarehouse,
	"DELETE_EM_WAREHOUSE": OperationTypeDeleteEmWarehouse,
	"MOVE_EM_WAREHOUSE":   OperationTypeMoveEmWarehouse,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_em_warehouse": OperationTypeCreateEmWarehouse,
	"update_em_warehouse": OperationTypeUpdateEmWarehouse,
	"delete_em_warehouse": OperationTypeDeleteEmWarehouse,
	"move_em_warehouse":   OperationTypeMoveEmWarehouse,
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
		"CREATE_EM_WAREHOUSE",
		"UPDATE_EM_WAREHOUSE",
		"DELETE_EM_WAREHOUSE",
		"MOVE_EM_WAREHOUSE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
