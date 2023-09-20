// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesDeployPlugin         OperationTypesEnum = "DEPLOY_PLUGIN"
	OperationTypesUpgradePlugin        OperationTypesEnum = "UPGRADE_PLUGIN"
	OperationTypesCreateUpgradePlugins OperationTypesEnum = "CREATE_UPGRADE_PLUGINS"
	OperationTypesAgentimageUpgrade    OperationTypesEnum = "AGENTIMAGE_UPGRADE"
	OperationTypesAddDataSource        OperationTypesEnum = "ADD_DATA_SOURCE"
	OperationTypesUpdateDataSource     OperationTypesEnum = "UPDATE_DATA_SOURCE"
	OperationTypesRemoveDataSource     OperationTypesEnum = "REMOVE_DATA_SOURCE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"DEPLOY_PLUGIN":          OperationTypesDeployPlugin,
	"UPGRADE_PLUGIN":         OperationTypesUpgradePlugin,
	"CREATE_UPGRADE_PLUGINS": OperationTypesCreateUpgradePlugins,
	"AGENTIMAGE_UPGRADE":     OperationTypesAgentimageUpgrade,
	"ADD_DATA_SOURCE":        OperationTypesAddDataSource,
	"UPDATE_DATA_SOURCE":     OperationTypesUpdateDataSource,
	"REMOVE_DATA_SOURCE":     OperationTypesRemoveDataSource,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"deploy_plugin":          OperationTypesDeployPlugin,
	"upgrade_plugin":         OperationTypesUpgradePlugin,
	"create_upgrade_plugins": OperationTypesCreateUpgradePlugins,
	"agentimage_upgrade":     OperationTypesAgentimageUpgrade,
	"add_data_source":        OperationTypesAddDataSource,
	"update_data_source":     OperationTypesUpdateDataSource,
	"remove_data_source":     OperationTypesRemoveDataSource,
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
		"DEPLOY_PLUGIN",
		"UPGRADE_PLUGIN",
		"CREATE_UPGRADE_PLUGINS",
		"AGENTIMAGE_UPGRADE",
		"ADD_DATA_SOURCE",
		"UPDATE_DATA_SOURCE",
		"REMOVE_DATA_SOURCE",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
