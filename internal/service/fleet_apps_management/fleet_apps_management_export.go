package fleet_apps_management

import (
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("fleet_apps_management", fleetAppsManagementResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportFleetAppsManagementTaskRecordHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_task_record",
	DatasourceClass:        "oci_fleet_apps_management_task_records",
	DatasourceItemsAttr:    "task_record_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "task_record",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.TaskRecordLifecycleStateActive),
	},
}

var exportFleetAppsManagementMaintenanceWindowHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_maintenance_window",
	DatasourceClass:        "oci_fleet_apps_management_maintenance_windows",
	DatasourceItemsAttr:    "maintenance_window_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "maintenance_window",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateActive),
		string(oci_fleet_apps_management.MaintenanceWindowLifecycleStateNeedsAttention),
	},
}

var exportFleetAppsManagementFleetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_fleet",
	DatasourceClass:        "oci_fleet_apps_management_fleets",
	DatasourceItemsAttr:    "fleet_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "fleet",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.FleetLifecycleStateActive),
		string(oci_fleet_apps_management.FleetLifecycleStateNeedsAttention),
	},
}

var exportFleetAppsManagementSchedulerDefinitionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_scheduler_definition",
	DatasourceClass:        "oci_fleet_apps_management_scheduler_definitions",
	DatasourceItemsAttr:    "scheduler_definition_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "scheduler_definition",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateActive),
	},
}

var exportFleetAppsManagementPropertyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_property",
	DatasourceClass:        "oci_fleet_apps_management_properties",
	DatasourceItemsAttr:    "property_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "property",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.PropertyLifecycleStateActive),
	},
}

var exportFleetAppsManagementRunbookHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_runbook",
	DatasourceClass:        "oci_fleet_apps_management_runbooks",
	DatasourceItemsAttr:    "runbook_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "runbook",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.RunbookLifecycleStateActive),
		string(oci_fleet_apps_management.RunbookLifecycleStateInactive),
	},
}

var exportFleetAppsManagementPlatformConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_platform_configuration",
	DatasourceClass:        "oci_fleet_apps_management_platform_configurations",
	DatasourceItemsAttr:    "platform_configuration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "platform_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.PlatformConfigurationLifecycleStateActive),
	},
}

var exportFleetAppsManagementCompliancePolicyRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_compliance_policy_rule",
	DatasourceClass:        "oci_fleet_apps_management_compliance_policy_rules",
	DatasourceItemsAttr:    "compliance_policy_rule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "compliance_policy_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.CompliancePolicyRuleLifecycleStateActive),
	},
}

var exportFleetAppsManagementPatchHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_fleet_apps_management_patch",
	DatasourceClass:        "oci_fleet_apps_management_patches",
	DatasourceItemsAttr:    "patch_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "patch",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_fleet_apps_management.PatchLifecycleStateActive),
	},
}

var fleetAppsManagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFleetAppsManagementTaskRecordHints},
		{TerraformResourceHints: exportFleetAppsManagementMaintenanceWindowHints},
		{TerraformResourceHints: exportFleetAppsManagementFleetHints},
		{TerraformResourceHints: exportFleetAppsManagementSchedulerDefinitionHints},
		{TerraformResourceHints: exportFleetAppsManagementPropertyHints},
		{TerraformResourceHints: exportFleetAppsManagementRunbookHints},
		{TerraformResourceHints: exportFleetAppsManagementPlatformConfigurationHints},
		{TerraformResourceHints: exportFleetAppsManagementCompliancePolicyRuleHints},
		{TerraformResourceHints: exportFleetAppsManagementPatchHints},
	},
}
