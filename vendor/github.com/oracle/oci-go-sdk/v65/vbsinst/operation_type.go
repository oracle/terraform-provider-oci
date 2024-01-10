// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VbsControlplaneInstance API
//
// A description of the VbsControlplaneInstance API
//

package vbsinst

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateVbsInstance OperationTypeEnum = "CREATE_VBS_INSTANCE"
	OperationTypeUpdateVbsInstance OperationTypeEnum = "UPDATE_VBS_INSTANCE"
	OperationTypeDeleteVbsInstance OperationTypeEnum = "DELETE_VBS_INSTANCE"
	OperationTypeMoveVbsInstance   OperationTypeEnum = "MOVE_VBS_INSTANCE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_VBS_INSTANCE": OperationTypeCreateVbsInstance,
	"UPDATE_VBS_INSTANCE": OperationTypeUpdateVbsInstance,
	"DELETE_VBS_INSTANCE": OperationTypeDeleteVbsInstance,
	"MOVE_VBS_INSTANCE":   OperationTypeMoveVbsInstance,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_vbs_instance": OperationTypeCreateVbsInstance,
	"update_vbs_instance": OperationTypeUpdateVbsInstance,
	"delete_vbs_instance": OperationTypeDeleteVbsInstance,
	"move_vbs_instance":   OperationTypeMoveVbsInstance,
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
		"CREATE_VBS_INSTANCE",
		"UPDATE_VBS_INSTANCE",
		"DELETE_VBS_INSTANCE",
		"MOVE_VBS_INSTANCE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
