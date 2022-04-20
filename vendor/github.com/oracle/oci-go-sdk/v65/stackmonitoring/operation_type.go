// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateResources         OperationTypeEnum = "CREATE_RESOURCES"
	OperationTypeUpdateResources         OperationTypeEnum = "UPDATE_RESOURCES"
	OperationTypeDeleteResources         OperationTypeEnum = "DELETE_RESOURCES"
	OperationTypeMoveResources           OperationTypeEnum = "MOVE_RESOURCES"
	OperationTypeEnableExternalDatabase  OperationTypeEnum = "ENABLE_EXTERNAL_DATABASE"
	OperationTypeDisableExternalDatabase OperationTypeEnum = "DISABLE_EXTERNAL_DATABASE"
	OperationTypeAddSourcesToAgent       OperationTypeEnum = "ADD_SOURCES_TO_AGENT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_RESOURCES":          OperationTypeCreateResources,
	"UPDATE_RESOURCES":          OperationTypeUpdateResources,
	"DELETE_RESOURCES":          OperationTypeDeleteResources,
	"MOVE_RESOURCES":            OperationTypeMoveResources,
	"ENABLE_EXTERNAL_DATABASE":  OperationTypeEnableExternalDatabase,
	"DISABLE_EXTERNAL_DATABASE": OperationTypeDisableExternalDatabase,
	"ADD_SOURCES_TO_AGENT":      OperationTypeAddSourcesToAgent,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_resources":          OperationTypeCreateResources,
	"update_resources":          OperationTypeUpdateResources,
	"delete_resources":          OperationTypeDeleteResources,
	"move_resources":            OperationTypeMoveResources,
	"enable_external_database":  OperationTypeEnableExternalDatabase,
	"disable_external_database": OperationTypeDisableExternalDatabase,
	"add_sources_to_agent":      OperationTypeAddSourcesToAgent,
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
		"CREATE_RESOURCES",
		"UPDATE_RESOURCES",
		"DELETE_RESOURCES",
		"MOVE_RESOURCES",
		"ENABLE_EXTERNAL_DATABASE",
		"DISABLE_EXTERNAL_DATABASE",
		"ADD_SOURCES_TO_AGENT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
