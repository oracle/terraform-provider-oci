// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDesktopPool OperationTypeEnum = "CREATE_DESKTOP_POOL"
	OperationTypeUpdateDesktopPool OperationTypeEnum = "UPDATE_DESKTOP_POOL"
	OperationTypeDeleteDesktopPool OperationTypeEnum = "DELETE_DESKTOP_POOL"
	OperationTypeMoveDesktopPool   OperationTypeEnum = "MOVE_DESKTOP_POOL"
	OperationTypeStartDesktopPool  OperationTypeEnum = "START_DESKTOP_POOL"
	OperationTypeStopDesktopPool   OperationTypeEnum = "STOP_DESKTOP_POOL"
	OperationTypeDeleteDesktop     OperationTypeEnum = "DELETE_DESKTOP"
	OperationTypeUpdateDesktop     OperationTypeEnum = "UPDATE_DESKTOP"
	OperationTypeStartDesktop      OperationTypeEnum = "START_DESKTOP"
	OperationTypeStopDesktop       OperationTypeEnum = "STOP_DESKTOP"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DESKTOP_POOL": OperationTypeCreateDesktopPool,
	"UPDATE_DESKTOP_POOL": OperationTypeUpdateDesktopPool,
	"DELETE_DESKTOP_POOL": OperationTypeDeleteDesktopPool,
	"MOVE_DESKTOP_POOL":   OperationTypeMoveDesktopPool,
	"START_DESKTOP_POOL":  OperationTypeStartDesktopPool,
	"STOP_DESKTOP_POOL":   OperationTypeStopDesktopPool,
	"DELETE_DESKTOP":      OperationTypeDeleteDesktop,
	"UPDATE_DESKTOP":      OperationTypeUpdateDesktop,
	"START_DESKTOP":       OperationTypeStartDesktop,
	"STOP_DESKTOP":        OperationTypeStopDesktop,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_desktop_pool": OperationTypeCreateDesktopPool,
	"update_desktop_pool": OperationTypeUpdateDesktopPool,
	"delete_desktop_pool": OperationTypeDeleteDesktopPool,
	"move_desktop_pool":   OperationTypeMoveDesktopPool,
	"start_desktop_pool":  OperationTypeStartDesktopPool,
	"stop_desktop_pool":   OperationTypeStopDesktopPool,
	"delete_desktop":      OperationTypeDeleteDesktop,
	"update_desktop":      OperationTypeUpdateDesktop,
	"start_desktop":       OperationTypeStartDesktop,
	"stop_desktop":        OperationTypeStopDesktop,
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
		"CREATE_DESKTOP_POOL",
		"UPDATE_DESKTOP_POOL",
		"DELETE_DESKTOP_POOL",
		"MOVE_DESKTOP_POOL",
		"START_DESKTOP_POOL",
		"STOP_DESKTOP_POOL",
		"DELETE_DESKTOP",
		"UPDATE_DESKTOP",
		"START_DESKTOP",
		"STOP_DESKTOP",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
