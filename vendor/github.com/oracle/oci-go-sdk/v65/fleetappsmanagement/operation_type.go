// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet                 OperationTypeEnum = "CREATE_FLEET"
	OperationTypeUpdateFleet                 OperationTypeEnum = "UPDATE_FLEET"
	OperationTypeDeleteFleet                 OperationTypeEnum = "DELETE_FLEET"
	OperationTypeMoveFleet                   OperationTypeEnum = "MOVE_FLEET"
	OperationTypeConfirmTarget               OperationTypeEnum = "CONFIRM_TARGET"
	OperationTypeGenerateCompliance          OperationTypeEnum = "GENERATE_COMPLIANCE"
	OperationTypeRequestTargetDiscovery      OperationTypeEnum = "REQUEST_TARGET_DISCOVERY"
	OperationTypeValidateResource            OperationTypeEnum = "VALIDATE_RESOURCE"
	OperationTypeCreateCredential            OperationTypeEnum = "CREATE_CREDENTIAL"
	OperationTypeUpdateCredential            OperationTypeEnum = "UPDATE_CREDENTIAL"
	OperationTypeDeleteCredential            OperationTypeEnum = "DELETE_CREDENTIAL"
	OperationTypeCreateSchedule              OperationTypeEnum = "CREATE_SCHEDULE"
	OperationTypeUpdateSchedule              OperationTypeEnum = "UPDATE_SCHEDULE"
	OperationTypeUpdateMaintenanceWindow     OperationTypeEnum = "UPDATE_MAINTENANCE_WINDOW"
	OperationTypeDeleteMaintenanceWindow     OperationTypeEnum = "DELETE_MAINTENANCE_WINDOW"
	OperationTypeCreateFleetResource         OperationTypeEnum = "CREATE_FLEET_RESOURCE"
	OperationTypeUpdateFleetResource         OperationTypeEnum = "UPDATE_FLEET_RESOURCE"
	OperationTypeDeleteFleetResource         OperationTypeEnum = "DELETE_FLEET_RESOURCE"
	OperationTypeCreateFamsOnboarding        OperationTypeEnum = "CREATE_FAMS_ONBOARDING"
	OperationTypeCreateRunbook               OperationTypeEnum = "CREATE_RUNBOOK"
	OperationTypeUpdateRunbook               OperationTypeEnum = "UPDATE_RUNBOOK"
	OperationTypeDeleteRunbook               OperationTypeEnum = "DELETE_RUNBOOK"
	OperationTypePublishRunbook              OperationTypeEnum = "PUBLISH_RUNBOOK"
	OperationTypeMoveRunbook                 OperationTypeEnum = "MOVE_RUNBOOK"
	OperationTypeCreateRunbookVersion        OperationTypeEnum = "CREATE_RUNBOOK_VERSION"
	OperationTypeUpdateRunbookVersion        OperationTypeEnum = "UPDATE_RUNBOOK_VERSION"
	OperationTypeDeleteRunbookVersion        OperationTypeEnum = "DELETE_RUNBOOK_VERSION"
	OperationTypePublishRunbookVersion       OperationTypeEnum = "PUBLISH_RUNBOOK_VERSION"
	OperationTypeMoveTask                    OperationTypeEnum = "MOVE_TASK"
	OperationTypeExportRunbook               OperationTypeEnum = "EXPORT_RUNBOOK"
	OperationTypeImportRunbook               OperationTypeEnum = "IMPORT_RUNBOOK"
	OperationTypeExportRunbookVersion        OperationTypeEnum = "EXPORT_RUNBOOK_VERSION"
	OperationTypeImportRunbookVersion        OperationTypeEnum = "IMPORT_RUNBOOK_VERSION"
	OperationTypeUpdateTask                  OperationTypeEnum = "UPDATE_TASK"
	OperationTypeDeleteTask                  OperationTypeEnum = "DELETE_TASK"
	OperationTypeUpdateFamsOnboarding        OperationTypeEnum = "UPDATE_FAMS_ONBOARDING"
	OperationTypeDeleteFamsOnboarding        OperationTypeEnum = "DELETE_FAMS_ONBOARDING"
	OperationTypeCreateCompliancePolicyRule  OperationTypeEnum = "CREATE_COMPLIANCE_POLICY_RULE"
	OperationTypeUpdateCompliancePolicyRule  OperationTypeEnum = "UPDATE_COMPLIANCE_POLICY_RULE"
	OperationTypeDeleteCompliancePolicyRule  OperationTypeEnum = "DELETE_COMPLIANCE_POLICY_RULE"
	OperationTypeUpdatePatch                 OperationTypeEnum = "UPDATE_PATCH"
	OperationTypeDeletePatch                 OperationTypeEnum = "DELETE_PATCH"
	OperationTypeMovePatch                   OperationTypeEnum = "MOVE_PATCH"
	OperationTypeManageJobExecution          OperationTypeEnum = "MANAGE_JOB_EXECUTION"
	OperationTypeDeletePlatformConfiguration OperationTypeEnum = "DELETE_PLATFORM_CONFIGURATION"
	OperationTypeUpdatePlatformConfiguration OperationTypeEnum = "UPDATE_PLATFORM_CONFIGURATION"
	OperationTypeMovePlatformConfiguration   OperationTypeEnum = "MOVE_PLATFORM_CONFIGURATION"
	OperationTypeCreatePlatformConfiguration OperationTypeEnum = "CREATE_PLATFORM_CONFIGURATION"
	OperationTypeMoveProperty                OperationTypeEnum = "MOVE_PROPERTY"
	OperationTypeCreateCatalogItem           OperationTypeEnum = "CREATE_CATALOG_ITEM"
	OperationTypeUpdateCatalogItem           OperationTypeEnum = "UPDATE_CATALOG_ITEM"
	OperationTypeDeleteCatalogItem           OperationTypeEnum = "DELETE_CATALOG_ITEM"
	OperationTypeMoveCatalogItem             OperationTypeEnum = "MOVE_CATALOG_ITEM"
	OperationTypeCloneCatalogItem            OperationTypeEnum = "CLONE_CATALOG_ITEM"
	OperationTypeCreateProvision             OperationTypeEnum = "CREATE_PROVISION"
	OperationTypeUpdateProvision             OperationTypeEnum = "UPDATE_PROVISION"
	OperationTypeDeleteProvision             OperationTypeEnum = "DELETE_PROVISION"
	OperationTypeMoveProvision               OperationTypeEnum = "MOVE_PROVISION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_FLEET":                  OperationTypeCreateFleet,
	"UPDATE_FLEET":                  OperationTypeUpdateFleet,
	"DELETE_FLEET":                  OperationTypeDeleteFleet,
	"MOVE_FLEET":                    OperationTypeMoveFleet,
	"CONFIRM_TARGET":                OperationTypeConfirmTarget,
	"GENERATE_COMPLIANCE":           OperationTypeGenerateCompliance,
	"REQUEST_TARGET_DISCOVERY":      OperationTypeRequestTargetDiscovery,
	"VALIDATE_RESOURCE":             OperationTypeValidateResource,
	"CREATE_CREDENTIAL":             OperationTypeCreateCredential,
	"UPDATE_CREDENTIAL":             OperationTypeUpdateCredential,
	"DELETE_CREDENTIAL":             OperationTypeDeleteCredential,
	"CREATE_SCHEDULE":               OperationTypeCreateSchedule,
	"UPDATE_SCHEDULE":               OperationTypeUpdateSchedule,
	"UPDATE_MAINTENANCE_WINDOW":     OperationTypeUpdateMaintenanceWindow,
	"DELETE_MAINTENANCE_WINDOW":     OperationTypeDeleteMaintenanceWindow,
	"CREATE_FLEET_RESOURCE":         OperationTypeCreateFleetResource,
	"UPDATE_FLEET_RESOURCE":         OperationTypeUpdateFleetResource,
	"DELETE_FLEET_RESOURCE":         OperationTypeDeleteFleetResource,
	"CREATE_FAMS_ONBOARDING":        OperationTypeCreateFamsOnboarding,
	"CREATE_RUNBOOK":                OperationTypeCreateRunbook,
	"UPDATE_RUNBOOK":                OperationTypeUpdateRunbook,
	"DELETE_RUNBOOK":                OperationTypeDeleteRunbook,
	"PUBLISH_RUNBOOK":               OperationTypePublishRunbook,
	"MOVE_RUNBOOK":                  OperationTypeMoveRunbook,
	"CREATE_RUNBOOK_VERSION":        OperationTypeCreateRunbookVersion,
	"UPDATE_RUNBOOK_VERSION":        OperationTypeUpdateRunbookVersion,
	"DELETE_RUNBOOK_VERSION":        OperationTypeDeleteRunbookVersion,
	"PUBLISH_RUNBOOK_VERSION":       OperationTypePublishRunbookVersion,
	"MOVE_TASK":                     OperationTypeMoveTask,
	"EXPORT_RUNBOOK":                OperationTypeExportRunbook,
	"IMPORT_RUNBOOK":                OperationTypeImportRunbook,
	"EXPORT_RUNBOOK_VERSION":        OperationTypeExportRunbookVersion,
	"IMPORT_RUNBOOK_VERSION":        OperationTypeImportRunbookVersion,
	"UPDATE_TASK":                   OperationTypeUpdateTask,
	"DELETE_TASK":                   OperationTypeDeleteTask,
	"UPDATE_FAMS_ONBOARDING":        OperationTypeUpdateFamsOnboarding,
	"DELETE_FAMS_ONBOARDING":        OperationTypeDeleteFamsOnboarding,
	"CREATE_COMPLIANCE_POLICY_RULE": OperationTypeCreateCompliancePolicyRule,
	"UPDATE_COMPLIANCE_POLICY_RULE": OperationTypeUpdateCompliancePolicyRule,
	"DELETE_COMPLIANCE_POLICY_RULE": OperationTypeDeleteCompliancePolicyRule,
	"UPDATE_PATCH":                  OperationTypeUpdatePatch,
	"DELETE_PATCH":                  OperationTypeDeletePatch,
	"MOVE_PATCH":                    OperationTypeMovePatch,
	"MANAGE_JOB_EXECUTION":          OperationTypeManageJobExecution,
	"DELETE_PLATFORM_CONFIGURATION": OperationTypeDeletePlatformConfiguration,
	"UPDATE_PLATFORM_CONFIGURATION": OperationTypeUpdatePlatformConfiguration,
	"MOVE_PLATFORM_CONFIGURATION":   OperationTypeMovePlatformConfiguration,
	"CREATE_PLATFORM_CONFIGURATION": OperationTypeCreatePlatformConfiguration,
	"MOVE_PROPERTY":                 OperationTypeMoveProperty,
	"CREATE_CATALOG_ITEM":           OperationTypeCreateCatalogItem,
	"UPDATE_CATALOG_ITEM":           OperationTypeUpdateCatalogItem,
	"DELETE_CATALOG_ITEM":           OperationTypeDeleteCatalogItem,
	"MOVE_CATALOG_ITEM":             OperationTypeMoveCatalogItem,
	"CLONE_CATALOG_ITEM":            OperationTypeCloneCatalogItem,
	"CREATE_PROVISION":              OperationTypeCreateProvision,
	"UPDATE_PROVISION":              OperationTypeUpdateProvision,
	"DELETE_PROVISION":              OperationTypeDeleteProvision,
	"MOVE_PROVISION":                OperationTypeMoveProvision,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_fleet":                  OperationTypeCreateFleet,
	"update_fleet":                  OperationTypeUpdateFleet,
	"delete_fleet":                  OperationTypeDeleteFleet,
	"move_fleet":                    OperationTypeMoveFleet,
	"confirm_target":                OperationTypeConfirmTarget,
	"generate_compliance":           OperationTypeGenerateCompliance,
	"request_target_discovery":      OperationTypeRequestTargetDiscovery,
	"validate_resource":             OperationTypeValidateResource,
	"create_credential":             OperationTypeCreateCredential,
	"update_credential":             OperationTypeUpdateCredential,
	"delete_credential":             OperationTypeDeleteCredential,
	"create_schedule":               OperationTypeCreateSchedule,
	"update_schedule":               OperationTypeUpdateSchedule,
	"update_maintenance_window":     OperationTypeUpdateMaintenanceWindow,
	"delete_maintenance_window":     OperationTypeDeleteMaintenanceWindow,
	"create_fleet_resource":         OperationTypeCreateFleetResource,
	"update_fleet_resource":         OperationTypeUpdateFleetResource,
	"delete_fleet_resource":         OperationTypeDeleteFleetResource,
	"create_fams_onboarding":        OperationTypeCreateFamsOnboarding,
	"create_runbook":                OperationTypeCreateRunbook,
	"update_runbook":                OperationTypeUpdateRunbook,
	"delete_runbook":                OperationTypeDeleteRunbook,
	"publish_runbook":               OperationTypePublishRunbook,
	"move_runbook":                  OperationTypeMoveRunbook,
	"create_runbook_version":        OperationTypeCreateRunbookVersion,
	"update_runbook_version":        OperationTypeUpdateRunbookVersion,
	"delete_runbook_version":        OperationTypeDeleteRunbookVersion,
	"publish_runbook_version":       OperationTypePublishRunbookVersion,
	"move_task":                     OperationTypeMoveTask,
	"export_runbook":                OperationTypeExportRunbook,
	"import_runbook":                OperationTypeImportRunbook,
	"export_runbook_version":        OperationTypeExportRunbookVersion,
	"import_runbook_version":        OperationTypeImportRunbookVersion,
	"update_task":                   OperationTypeUpdateTask,
	"delete_task":                   OperationTypeDeleteTask,
	"update_fams_onboarding":        OperationTypeUpdateFamsOnboarding,
	"delete_fams_onboarding":        OperationTypeDeleteFamsOnboarding,
	"create_compliance_policy_rule": OperationTypeCreateCompliancePolicyRule,
	"update_compliance_policy_rule": OperationTypeUpdateCompliancePolicyRule,
	"delete_compliance_policy_rule": OperationTypeDeleteCompliancePolicyRule,
	"update_patch":                  OperationTypeUpdatePatch,
	"delete_patch":                  OperationTypeDeletePatch,
	"move_patch":                    OperationTypeMovePatch,
	"manage_job_execution":          OperationTypeManageJobExecution,
	"delete_platform_configuration": OperationTypeDeletePlatformConfiguration,
	"update_platform_configuration": OperationTypeUpdatePlatformConfiguration,
	"move_platform_configuration":   OperationTypeMovePlatformConfiguration,
	"create_platform_configuration": OperationTypeCreatePlatformConfiguration,
	"move_property":                 OperationTypeMoveProperty,
	"create_catalog_item":           OperationTypeCreateCatalogItem,
	"update_catalog_item":           OperationTypeUpdateCatalogItem,
	"delete_catalog_item":           OperationTypeDeleteCatalogItem,
	"move_catalog_item":             OperationTypeMoveCatalogItem,
	"clone_catalog_item":            OperationTypeCloneCatalogItem,
	"create_provision":              OperationTypeCreateProvision,
	"update_provision":              OperationTypeUpdateProvision,
	"delete_provision":              OperationTypeDeleteProvision,
	"move_provision":                OperationTypeMoveProvision,
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
		"CREATE_FLEET",
		"UPDATE_FLEET",
		"DELETE_FLEET",
		"MOVE_FLEET",
		"CONFIRM_TARGET",
		"GENERATE_COMPLIANCE",
		"REQUEST_TARGET_DISCOVERY",
		"VALIDATE_RESOURCE",
		"CREATE_CREDENTIAL",
		"UPDATE_CREDENTIAL",
		"DELETE_CREDENTIAL",
		"CREATE_SCHEDULE",
		"UPDATE_SCHEDULE",
		"UPDATE_MAINTENANCE_WINDOW",
		"DELETE_MAINTENANCE_WINDOW",
		"CREATE_FLEET_RESOURCE",
		"UPDATE_FLEET_RESOURCE",
		"DELETE_FLEET_RESOURCE",
		"CREATE_FAMS_ONBOARDING",
		"CREATE_RUNBOOK",
		"UPDATE_RUNBOOK",
		"DELETE_RUNBOOK",
		"PUBLISH_RUNBOOK",
		"MOVE_RUNBOOK",
		"CREATE_RUNBOOK_VERSION",
		"UPDATE_RUNBOOK_VERSION",
		"DELETE_RUNBOOK_VERSION",
		"PUBLISH_RUNBOOK_VERSION",
		"MOVE_TASK",
		"EXPORT_RUNBOOK",
		"IMPORT_RUNBOOK",
		"EXPORT_RUNBOOK_VERSION",
		"IMPORT_RUNBOOK_VERSION",
		"UPDATE_TASK",
		"DELETE_TASK",
		"UPDATE_FAMS_ONBOARDING",
		"DELETE_FAMS_ONBOARDING",
		"CREATE_COMPLIANCE_POLICY_RULE",
		"UPDATE_COMPLIANCE_POLICY_RULE",
		"DELETE_COMPLIANCE_POLICY_RULE",
		"UPDATE_PATCH",
		"DELETE_PATCH",
		"MOVE_PATCH",
		"MANAGE_JOB_EXECUTION",
		"DELETE_PLATFORM_CONFIGURATION",
		"UPDATE_PLATFORM_CONFIGURATION",
		"MOVE_PLATFORM_CONFIGURATION",
		"CREATE_PLATFORM_CONFIGURATION",
		"MOVE_PROPERTY",
		"CREATE_CATALOG_ITEM",
		"UPDATE_CATALOG_ITEM",
		"DELETE_CATALOG_ITEM",
		"MOVE_CATALOG_ITEM",
		"CLONE_CATALOG_ITEM",
		"CREATE_PROVISION",
		"UPDATE_PROVISION",
		"DELETE_PROVISION",
		"MOVE_PROVISION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
