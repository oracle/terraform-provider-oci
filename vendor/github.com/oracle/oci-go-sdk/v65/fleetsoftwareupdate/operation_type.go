// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDiscovery           OperationTypeEnum = "CREATE_DISCOVERY"
	OperationTypeDeleteDiscovery           OperationTypeEnum = "DELETE_DISCOVERY"
	OperationTypeCreateCollection          OperationTypeEnum = "CREATE_COLLECTION"
	OperationTypeUpdateCollection          OperationTypeEnum = "UPDATE_COLLECTION"
	OperationTypeDeleteCollection          OperationTypeEnum = "DELETE_COLLECTION"
	OperationTypeMoveCollection            OperationTypeEnum = "MOVE_COLLECTION"
	OperationTypeAddTargetsToCollection    OperationTypeEnum = "ADD_TARGETS_TO_COLLECTION"
	OperationTypeRemoveTargetsInCollection OperationTypeEnum = "REMOVE_TARGETS_IN_COLLECTION"
	OperationTypeCreateMaintenanceCycle    OperationTypeEnum = "CREATE_MAINTENANCE_CYCLE"
	OperationTypeUpdateMaintenanceCycle    OperationTypeEnum = "UPDATE_MAINTENANCE_CYCLE"
	OperationTypeDeleteMaintenanceCycle    OperationTypeEnum = "DELETE_MAINTENANCE_CYCLE"
	OperationTypeMoveMaintenanceCycle      OperationTypeEnum = "MOVE_MAINTENANCE_CYCLE"
	OperationTypeCloneMaintenanceCycle     OperationTypeEnum = "CLONE_MAINTENANCE_CYCLE"
	OperationTypeCreateAction              OperationTypeEnum = "CREATE_ACTION"
	OperationTypeUpdateAction              OperationTypeEnum = "UPDATE_ACTION"
	OperationTypeDeleteAction              OperationTypeEnum = "DELETE_ACTION"
	OperationTypeMoveAction                OperationTypeEnum = "MOVE_ACTION"
	OperationTypePatchAction               OperationTypeEnum = "PATCH_ACTION"
	OperationTypeCleanupAction             OperationTypeEnum = "CLEANUP_ACTION"
	OperationTypeRollbackAndRemoveAction   OperationTypeEnum = "ROLLBACK_AND_REMOVE_ACTION"
	OperationTypeApplyAction               OperationTypeEnum = "APPLY_ACTION"
	OperationTypePrecheckAction            OperationTypeEnum = "PRECHECK_ACTION"
	OperationTypeStageAction               OperationTypeEnum = "STAGE_ACTION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DISCOVERY":             OperationTypeCreateDiscovery,
	"DELETE_DISCOVERY":             OperationTypeDeleteDiscovery,
	"CREATE_COLLECTION":            OperationTypeCreateCollection,
	"UPDATE_COLLECTION":            OperationTypeUpdateCollection,
	"DELETE_COLLECTION":            OperationTypeDeleteCollection,
	"MOVE_COLLECTION":              OperationTypeMoveCollection,
	"ADD_TARGETS_TO_COLLECTION":    OperationTypeAddTargetsToCollection,
	"REMOVE_TARGETS_IN_COLLECTION": OperationTypeRemoveTargetsInCollection,
	"CREATE_MAINTENANCE_CYCLE":     OperationTypeCreateMaintenanceCycle,
	"UPDATE_MAINTENANCE_CYCLE":     OperationTypeUpdateMaintenanceCycle,
	"DELETE_MAINTENANCE_CYCLE":     OperationTypeDeleteMaintenanceCycle,
	"MOVE_MAINTENANCE_CYCLE":       OperationTypeMoveMaintenanceCycle,
	"CLONE_MAINTENANCE_CYCLE":      OperationTypeCloneMaintenanceCycle,
	"CREATE_ACTION":                OperationTypeCreateAction,
	"UPDATE_ACTION":                OperationTypeUpdateAction,
	"DELETE_ACTION":                OperationTypeDeleteAction,
	"MOVE_ACTION":                  OperationTypeMoveAction,
	"PATCH_ACTION":                 OperationTypePatchAction,
	"CLEANUP_ACTION":               OperationTypeCleanupAction,
	"ROLLBACK_AND_REMOVE_ACTION":   OperationTypeRollbackAndRemoveAction,
	"APPLY_ACTION":                 OperationTypeApplyAction,
	"PRECHECK_ACTION":              OperationTypePrecheckAction,
	"STAGE_ACTION":                 OperationTypeStageAction,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_discovery":             OperationTypeCreateDiscovery,
	"delete_discovery":             OperationTypeDeleteDiscovery,
	"create_collection":            OperationTypeCreateCollection,
	"update_collection":            OperationTypeUpdateCollection,
	"delete_collection":            OperationTypeDeleteCollection,
	"move_collection":              OperationTypeMoveCollection,
	"add_targets_to_collection":    OperationTypeAddTargetsToCollection,
	"remove_targets_in_collection": OperationTypeRemoveTargetsInCollection,
	"create_maintenance_cycle":     OperationTypeCreateMaintenanceCycle,
	"update_maintenance_cycle":     OperationTypeUpdateMaintenanceCycle,
	"delete_maintenance_cycle":     OperationTypeDeleteMaintenanceCycle,
	"move_maintenance_cycle":       OperationTypeMoveMaintenanceCycle,
	"clone_maintenance_cycle":      OperationTypeCloneMaintenanceCycle,
	"create_action":                OperationTypeCreateAction,
	"update_action":                OperationTypeUpdateAction,
	"delete_action":                OperationTypeDeleteAction,
	"move_action":                  OperationTypeMoveAction,
	"patch_action":                 OperationTypePatchAction,
	"cleanup_action":               OperationTypeCleanupAction,
	"rollback_and_remove_action":   OperationTypeRollbackAndRemoveAction,
	"apply_action":                 OperationTypeApplyAction,
	"precheck_action":              OperationTypePrecheckAction,
	"stage_action":                 OperationTypeStageAction,
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
		"CREATE_DISCOVERY",
		"DELETE_DISCOVERY",
		"CREATE_COLLECTION",
		"UPDATE_COLLECTION",
		"DELETE_COLLECTION",
		"MOVE_COLLECTION",
		"ADD_TARGETS_TO_COLLECTION",
		"REMOVE_TARGETS_IN_COLLECTION",
		"CREATE_MAINTENANCE_CYCLE",
		"UPDATE_MAINTENANCE_CYCLE",
		"DELETE_MAINTENANCE_CYCLE",
		"MOVE_MAINTENANCE_CYCLE",
		"CLONE_MAINTENANCE_CYCLE",
		"CREATE_ACTION",
		"UPDATE_ACTION",
		"DELETE_ACTION",
		"MOVE_ACTION",
		"PATCH_ACTION",
		"CLEANUP_ACTION",
		"ROLLBACK_AND_REMOVE_ACTION",
		"APPLY_ACTION",
		"PRECHECK_ACTION",
		"STAGE_ACTION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
