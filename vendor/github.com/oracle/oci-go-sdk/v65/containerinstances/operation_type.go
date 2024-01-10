// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateContainerInstance  OperationTypeEnum = "CREATE_CONTAINER_INSTANCE"
	OperationTypeUpdateContainerInstance  OperationTypeEnum = "UPDATE_CONTAINER_INSTANCE"
	OperationTypeDeleteContainerInstance  OperationTypeEnum = "DELETE_CONTAINER_INSTANCE"
	OperationTypeMoveContainerInstance    OperationTypeEnum = "MOVE_CONTAINER_INSTANCE"
	OperationTypeStartContainerInstance   OperationTypeEnum = "START_CONTAINER_INSTANCE"
	OperationTypeStopContainerInstance    OperationTypeEnum = "STOP_CONTAINER_INSTANCE"
	OperationTypeRestartContainerInstance OperationTypeEnum = "RESTART_CONTAINER_INSTANCE"
	OperationTypeUpdateContainer          OperationTypeEnum = "UPDATE_CONTAINER"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_CONTAINER_INSTANCE":  OperationTypeCreateContainerInstance,
	"UPDATE_CONTAINER_INSTANCE":  OperationTypeUpdateContainerInstance,
	"DELETE_CONTAINER_INSTANCE":  OperationTypeDeleteContainerInstance,
	"MOVE_CONTAINER_INSTANCE":    OperationTypeMoveContainerInstance,
	"START_CONTAINER_INSTANCE":   OperationTypeStartContainerInstance,
	"STOP_CONTAINER_INSTANCE":    OperationTypeStopContainerInstance,
	"RESTART_CONTAINER_INSTANCE": OperationTypeRestartContainerInstance,
	"UPDATE_CONTAINER":           OperationTypeUpdateContainer,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_container_instance":  OperationTypeCreateContainerInstance,
	"update_container_instance":  OperationTypeUpdateContainerInstance,
	"delete_container_instance":  OperationTypeDeleteContainerInstance,
	"move_container_instance":    OperationTypeMoveContainerInstance,
	"start_container_instance":   OperationTypeStartContainerInstance,
	"stop_container_instance":    OperationTypeStopContainerInstance,
	"restart_container_instance": OperationTypeRestartContainerInstance,
	"update_container":           OperationTypeUpdateContainer,
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
		"CREATE_CONTAINER_INSTANCE",
		"UPDATE_CONTAINER_INSTANCE",
		"DELETE_CONTAINER_INSTANCE",
		"MOVE_CONTAINER_INSTANCE",
		"START_CONTAINER_INSTANCE",
		"STOP_CONTAINER_INSTANCE",
		"RESTART_CONTAINER_INSTANCE",
		"UPDATE_CONTAINER",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
