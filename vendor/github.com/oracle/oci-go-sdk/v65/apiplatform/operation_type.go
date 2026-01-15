// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APIP Control Plane API
//
// Control Plane designed to manage lifecycle of APIP Instances
//

package apiplatform

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateApipInstance OperationTypeEnum = "CREATE_APIP_INSTANCE"
	OperationTypeUpdateApipInstance OperationTypeEnum = "UPDATE_APIP_INSTANCE"
	OperationTypeDeleteApipInstance OperationTypeEnum = "DELETE_APIP_INSTANCE"
	OperationTypeMoveApipInstance   OperationTypeEnum = "MOVE_APIP_INSTANCE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_APIP_INSTANCE": OperationTypeCreateApipInstance,
	"UPDATE_APIP_INSTANCE": OperationTypeUpdateApipInstance,
	"DELETE_APIP_INSTANCE": OperationTypeDeleteApipInstance,
	"MOVE_APIP_INSTANCE":   OperationTypeMoveApipInstance,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_apip_instance": OperationTypeCreateApipInstance,
	"update_apip_instance": OperationTypeUpdateApipInstance,
	"delete_apip_instance": OperationTypeDeleteApipInstance,
	"move_apip_instance":   OperationTypeMoveApipInstance,
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
		"CREATE_APIP_INSTANCE",
		"UPDATE_APIP_INSTANCE",
		"DELETE_APIP_INSTANCE",
		"MOVE_APIP_INSTANCE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
