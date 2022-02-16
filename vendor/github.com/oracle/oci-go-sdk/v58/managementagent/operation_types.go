// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesUpgradePlugin        OperationTypesEnum = "UPGRADE_PLUGIN"
	OperationTypesCreateUpgradePlugins OperationTypesEnum = "CREATE_UPGRADE_PLUGINS"
	OperationTypesAgentimageUpgrade    OperationTypesEnum = "AGENTIMAGE_UPGRADE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"UPGRADE_PLUGIN":         OperationTypesUpgradePlugin,
	"CREATE_UPGRADE_PLUGINS": OperationTypesCreateUpgradePlugins,
	"AGENTIMAGE_UPGRADE":     OperationTypesAgentimageUpgrade,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"UPGRADE_PLUGIN",
		"CREATE_UPGRADE_PLUGINS",
		"AGENTIMAGE_UPGRADE",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	mappingOperationTypesEnumIgnoreCase := make(map[string]OperationTypesEnum)
	for k, v := range mappingOperationTypesEnum {
		mappingOperationTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
