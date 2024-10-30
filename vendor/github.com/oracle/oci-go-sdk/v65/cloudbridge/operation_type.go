// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateEnvironment     OperationTypeEnum = "CREATE_ENVIRONMENT"
	OperationTypeUpdateEnvironment     OperationTypeEnum = "UPDATE_ENVIRONMENT"
	OperationTypeDeleteEnvironment     OperationTypeEnum = "DELETE_ENVIRONMENT"
	OperationTypeMoveEnvironment       OperationTypeEnum = "MOVE_ENVIRONMENT"
	OperationTypeCreateOcbAgent        OperationTypeEnum = "CREATE_OCB_AGENT"
	OperationTypeUpdateOcbAgent        OperationTypeEnum = "UPDATE_OCB_AGENT"
	OperationTypeDeleteOcbAgent        OperationTypeEnum = "DELETE_OCB_AGENT"
	OperationTypeMoveOcbAgent          OperationTypeEnum = "MOVE_OCB_AGENT"
	OperationTypeCreateAgentDependency OperationTypeEnum = "CREATE_AGENT_DEPENDENCY"
	OperationTypeUpdateAgentDependency OperationTypeEnum = "UPDATE_AGENT_DEPENDENCY"
	OperationTypeDeleteAgentDependency OperationTypeEnum = "DELETE_AGENT_DEPENDENCY"
	OperationTypeMoveAgentDependency   OperationTypeEnum = "MOVE_AGENT_DEPENDENCY"
	OperationTypeCreateInventory       OperationTypeEnum = "CREATE_INVENTORY"
	OperationTypeDeleteInventory       OperationTypeEnum = "DELETE_INVENTORY"
	OperationTypeImportInventory       OperationTypeEnum = "IMPORT_INVENTORY"
	OperationTypeDeleteAssetSource     OperationTypeEnum = "DELETE_ASSET_SOURCE"
	OperationTypeRefreshAssetSource    OperationTypeEnum = "REFRESH_ASSET_SOURCE"
	OperationTypeCreateAssetSource     OperationTypeEnum = "CREATE_ASSET_SOURCE"
	OperationTypeUpdateAssetSource     OperationTypeEnum = "UPDATE_ASSET_SOURCE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_ENVIRONMENT":      OperationTypeCreateEnvironment,
	"UPDATE_ENVIRONMENT":      OperationTypeUpdateEnvironment,
	"DELETE_ENVIRONMENT":      OperationTypeDeleteEnvironment,
	"MOVE_ENVIRONMENT":        OperationTypeMoveEnvironment,
	"CREATE_OCB_AGENT":        OperationTypeCreateOcbAgent,
	"UPDATE_OCB_AGENT":        OperationTypeUpdateOcbAgent,
	"DELETE_OCB_AGENT":        OperationTypeDeleteOcbAgent,
	"MOVE_OCB_AGENT":          OperationTypeMoveOcbAgent,
	"CREATE_AGENT_DEPENDENCY": OperationTypeCreateAgentDependency,
	"UPDATE_AGENT_DEPENDENCY": OperationTypeUpdateAgentDependency,
	"DELETE_AGENT_DEPENDENCY": OperationTypeDeleteAgentDependency,
	"MOVE_AGENT_DEPENDENCY":   OperationTypeMoveAgentDependency,
	"CREATE_INVENTORY":        OperationTypeCreateInventory,
	"DELETE_INVENTORY":        OperationTypeDeleteInventory,
	"IMPORT_INVENTORY":        OperationTypeImportInventory,
	"DELETE_ASSET_SOURCE":     OperationTypeDeleteAssetSource,
	"REFRESH_ASSET_SOURCE":    OperationTypeRefreshAssetSource,
	"CREATE_ASSET_SOURCE":     OperationTypeCreateAssetSource,
	"UPDATE_ASSET_SOURCE":     OperationTypeUpdateAssetSource,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_environment":      OperationTypeCreateEnvironment,
	"update_environment":      OperationTypeUpdateEnvironment,
	"delete_environment":      OperationTypeDeleteEnvironment,
	"move_environment":        OperationTypeMoveEnvironment,
	"create_ocb_agent":        OperationTypeCreateOcbAgent,
	"update_ocb_agent":        OperationTypeUpdateOcbAgent,
	"delete_ocb_agent":        OperationTypeDeleteOcbAgent,
	"move_ocb_agent":          OperationTypeMoveOcbAgent,
	"create_agent_dependency": OperationTypeCreateAgentDependency,
	"update_agent_dependency": OperationTypeUpdateAgentDependency,
	"delete_agent_dependency": OperationTypeDeleteAgentDependency,
	"move_agent_dependency":   OperationTypeMoveAgentDependency,
	"create_inventory":        OperationTypeCreateInventory,
	"delete_inventory":        OperationTypeDeleteInventory,
	"import_inventory":        OperationTypeImportInventory,
	"delete_asset_source":     OperationTypeDeleteAssetSource,
	"refresh_asset_source":    OperationTypeRefreshAssetSource,
	"create_asset_source":     OperationTypeCreateAssetSource,
	"update_asset_source":     OperationTypeUpdateAssetSource,
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
		"CREATE_ENVIRONMENT",
		"UPDATE_ENVIRONMENT",
		"DELETE_ENVIRONMENT",
		"MOVE_ENVIRONMENT",
		"CREATE_OCB_AGENT",
		"UPDATE_OCB_AGENT",
		"DELETE_OCB_AGENT",
		"MOVE_OCB_AGENT",
		"CREATE_AGENT_DEPENDENCY",
		"UPDATE_AGENT_DEPENDENCY",
		"DELETE_AGENT_DEPENDENCY",
		"MOVE_AGENT_DEPENDENCY",
		"CREATE_INVENTORY",
		"DELETE_INVENTORY",
		"IMPORT_INVENTORY",
		"DELETE_ASSET_SOURCE",
		"REFRESH_ASSET_SOURCE",
		"CREATE_ASSET_SOURCE",
		"UPDATE_ASSET_SOURCE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
