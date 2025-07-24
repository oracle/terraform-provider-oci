// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesDeployPlugin           OperationTypesEnum = "DEPLOY_PLUGIN"
	OperationTypesUpgradePlugin          OperationTypesEnum = "UPGRADE_PLUGIN"
	OperationTypesCreatePlugin           OperationTypesEnum = "CREATE_PLUGIN"
	OperationTypesCreateUpgradePlugins   OperationTypesEnum = "CREATE_UPGRADE_PLUGINS"
	OperationTypesAgentimageUpgrade      OperationTypesEnum = "AGENTIMAGE_UPGRADE"
	OperationTypesCreateDataSource       OperationTypesEnum = "CREATE_DATA_SOURCE"
	OperationTypesUpdateDataSource       OperationTypesEnum = "UPDATE_DATA_SOURCE"
	OperationTypesDeleteDataSource       OperationTypesEnum = "DELETE_DATA_SOURCE"
	OperationTypesCreateNamedcredentials OperationTypesEnum = "CREATE_NAMEDCREDENTIALS"
	OperationTypesUpdateNamedcredentials OperationTypesEnum = "UPDATE_NAMEDCREDENTIALS"
	OperationTypesDeleteNamedcredentials OperationTypesEnum = "DELETE_NAMEDCREDENTIALS"
	OperationTypesAgentUpgrade           OperationTypesEnum = "AGENT_UPGRADE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"DEPLOY_PLUGIN":           OperationTypesDeployPlugin,
	"UPGRADE_PLUGIN":          OperationTypesUpgradePlugin,
	"CREATE_PLUGIN":           OperationTypesCreatePlugin,
	"CREATE_UPGRADE_PLUGINS":  OperationTypesCreateUpgradePlugins,
	"AGENTIMAGE_UPGRADE":      OperationTypesAgentimageUpgrade,
	"CREATE_DATA_SOURCE":      OperationTypesCreateDataSource,
	"UPDATE_DATA_SOURCE":      OperationTypesUpdateDataSource,
	"DELETE_DATA_SOURCE":      OperationTypesDeleteDataSource,
	"CREATE_NAMEDCREDENTIALS": OperationTypesCreateNamedcredentials,
	"UPDATE_NAMEDCREDENTIALS": OperationTypesUpdateNamedcredentials,
	"DELETE_NAMEDCREDENTIALS": OperationTypesDeleteNamedcredentials,
	"AGENT_UPGRADE":           OperationTypesAgentUpgrade,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"deploy_plugin":           OperationTypesDeployPlugin,
	"upgrade_plugin":          OperationTypesUpgradePlugin,
	"create_plugin":           OperationTypesCreatePlugin,
	"create_upgrade_plugins":  OperationTypesCreateUpgradePlugins,
	"agentimage_upgrade":      OperationTypesAgentimageUpgrade,
	"create_data_source":      OperationTypesCreateDataSource,
	"update_data_source":      OperationTypesUpdateDataSource,
	"delete_data_source":      OperationTypesDeleteDataSource,
	"create_namedcredentials": OperationTypesCreateNamedcredentials,
	"update_namedcredentials": OperationTypesUpdateNamedcredentials,
	"delete_namedcredentials": OperationTypesDeleteNamedcredentials,
	"agent_upgrade":           OperationTypesAgentUpgrade,
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
		"CREATE_PLUGIN",
		"CREATE_UPGRADE_PLUGINS",
		"AGENTIMAGE_UPGRADE",
		"CREATE_DATA_SOURCE",
		"UPDATE_DATA_SOURCE",
		"DELETE_DATA_SOURCE",
		"CREATE_NAMEDCREDENTIALS",
		"UPDATE_NAMEDCREDENTIALS",
		"DELETE_NAMEDCREDENTIALS",
		"AGENT_UPGRADE",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
