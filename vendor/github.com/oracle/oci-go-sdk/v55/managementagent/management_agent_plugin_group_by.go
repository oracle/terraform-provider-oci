// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

// ManagementAgentPluginGroupByEnum Enum with underlying type: string
type ManagementAgentPluginGroupByEnum string

// Set of constants representing the allowable values for ManagementAgentPluginGroupByEnum
const (
	ManagementAgentPluginGroupByPluginName ManagementAgentPluginGroupByEnum = "pluginName"
)

var mappingManagementAgentPluginGroupBy = map[string]ManagementAgentPluginGroupByEnum{
	"pluginName": ManagementAgentPluginGroupByPluginName,
}

// GetManagementAgentPluginGroupByEnumValues Enumerates the set of values for ManagementAgentPluginGroupByEnum
func GetManagementAgentPluginGroupByEnumValues() []ManagementAgentPluginGroupByEnum {
	values := make([]ManagementAgentPluginGroupByEnum, 0)
	for _, v := range mappingManagementAgentPluginGroupBy {
		values = append(values, v)
	}
	return values
}
