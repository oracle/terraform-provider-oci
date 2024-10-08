// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MaintenanceWindowOperationTypeEnum Enum with underlying type: string
type MaintenanceWindowOperationTypeEnum string

// Set of constants representing the allowable values for MaintenanceWindowOperationTypeEnum
const (
	MaintenanceWindowOperationTypeUpdate MaintenanceWindowOperationTypeEnum = "UPDATE"
	MaintenanceWindowOperationTypeCreate MaintenanceWindowOperationTypeEnum = "CREATE"
	MaintenanceWindowOperationTypeDelete MaintenanceWindowOperationTypeEnum = "DELETE"
	MaintenanceWindowOperationTypeStop   MaintenanceWindowOperationTypeEnum = "STOP"
)

var mappingMaintenanceWindowOperationTypeEnum = map[string]MaintenanceWindowOperationTypeEnum{
	"UPDATE": MaintenanceWindowOperationTypeUpdate,
	"CREATE": MaintenanceWindowOperationTypeCreate,
	"DELETE": MaintenanceWindowOperationTypeDelete,
	"STOP":   MaintenanceWindowOperationTypeStop,
}

var mappingMaintenanceWindowOperationTypeEnumLowerCase = map[string]MaintenanceWindowOperationTypeEnum{
	"update": MaintenanceWindowOperationTypeUpdate,
	"create": MaintenanceWindowOperationTypeCreate,
	"delete": MaintenanceWindowOperationTypeDelete,
	"stop":   MaintenanceWindowOperationTypeStop,
}

// GetMaintenanceWindowOperationTypeEnumValues Enumerates the set of values for MaintenanceWindowOperationTypeEnum
func GetMaintenanceWindowOperationTypeEnumValues() []MaintenanceWindowOperationTypeEnum {
	values := make([]MaintenanceWindowOperationTypeEnum, 0)
	for _, v := range mappingMaintenanceWindowOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowOperationTypeEnumStringValues Enumerates the set of values in String for MaintenanceWindowOperationTypeEnum
func GetMaintenanceWindowOperationTypeEnumStringValues() []string {
	return []string{
		"UPDATE",
		"CREATE",
		"DELETE",
		"STOP",
	}
}

// GetMappingMaintenanceWindowOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowOperationTypeEnum(val string) (MaintenanceWindowOperationTypeEnum, bool) {
	enum, ok := mappingMaintenanceWindowOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
