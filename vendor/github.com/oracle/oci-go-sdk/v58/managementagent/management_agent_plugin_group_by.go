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

// ManagementAgentPluginGroupByEnum Enum with underlying type: string
type ManagementAgentPluginGroupByEnum string

// Set of constants representing the allowable values for ManagementAgentPluginGroupByEnum
const (
	ManagementAgentPluginGroupByPluginName ManagementAgentPluginGroupByEnum = "pluginName"
)

var mappingManagementAgentPluginGroupByEnum = map[string]ManagementAgentPluginGroupByEnum{
	"pluginName": ManagementAgentPluginGroupByPluginName,
}

// GetManagementAgentPluginGroupByEnumValues Enumerates the set of values for ManagementAgentPluginGroupByEnum
func GetManagementAgentPluginGroupByEnumValues() []ManagementAgentPluginGroupByEnum {
	values := make([]ManagementAgentPluginGroupByEnum, 0)
	for _, v := range mappingManagementAgentPluginGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementAgentPluginGroupByEnumStringValues Enumerates the set of values in String for ManagementAgentPluginGroupByEnum
func GetManagementAgentPluginGroupByEnumStringValues() []string {
	return []string{
		"pluginName",
	}
}

// GetMappingManagementAgentPluginGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementAgentPluginGroupByEnum(val string) (ManagementAgentPluginGroupByEnum, bool) {
	mappingManagementAgentPluginGroupByEnumIgnoreCase := make(map[string]ManagementAgentPluginGroupByEnum)
	for k, v := range mappingManagementAgentPluginGroupByEnum {
		mappingManagementAgentPluginGroupByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingManagementAgentPluginGroupByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
