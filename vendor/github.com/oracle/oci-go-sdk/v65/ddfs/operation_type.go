// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Device Data FHIR Service API
//
// Use the Device Data FHIR Service API to manage DDFS instances.
//

package ddfs

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateInstance OperationTypeEnum = "CREATE_INSTANCE"
	OperationTypeUpdateInstance OperationTypeEnum = "UPDATE_INSTANCE"
	OperationTypeDeleteInstance OperationTypeEnum = "DELETE_INSTANCE"
	OperationTypeMoveInstance   OperationTypeEnum = "MOVE_INSTANCE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_INSTANCE": OperationTypeCreateInstance,
	"UPDATE_INSTANCE": OperationTypeUpdateInstance,
	"DELETE_INSTANCE": OperationTypeDeleteInstance,
	"MOVE_INSTANCE":   OperationTypeMoveInstance,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_instance": OperationTypeCreateInstance,
	"update_instance": OperationTypeUpdateInstance,
	"delete_instance": OperationTypeDeleteInstance,
	"move_instance":   OperationTypeMoveInstance,
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
		"CREATE_INSTANCE",
		"UPDATE_INSTANCE",
		"DELETE_INSTANCE",
		"MOVE_INSTANCE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
