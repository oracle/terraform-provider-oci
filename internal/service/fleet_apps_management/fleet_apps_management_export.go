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

var fleetAppsManagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFleetAppsManagementMaintenanceWindowHints},
		{TerraformResourceHints: exportFleetAppsManagementFleetHints},
		{TerraformResourceHints: exportFleetAppsManagementSchedulerDefinitionHints},
		{TerraformResourceHints: exportFleetAppsManagementPropertyHints},
	},
}
