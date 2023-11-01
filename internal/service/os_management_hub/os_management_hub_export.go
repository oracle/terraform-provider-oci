package os_management_hub

import (
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("os_management_hub", osManagementHubResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOsManagementHubLifecycleEnvironmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_lifecycle_environment",
	DatasourceClass:        "oci_os_management_hub_lifecycle_environments",
	DatasourceItemsAttr:    "lifecycle_environment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "lifecycle_environment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.LifecycleEnvironmentLifecycleStateActive),
	},
}

var exportOsManagementHubSoftwareSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_software_source",
	DatasourceClass:        "oci_os_management_hub_software_sources",
	DatasourceItemsAttr:    "software_source_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "software_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.SoftwareSourceLifecycleStateActive),
	},
}

var exportOsManagementHubManagedInstanceGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_managed_instance_group",
	DatasourceClass:        "oci_os_management_hub_managed_instance_groups",
	DatasourceItemsAttr:    "managed_instance_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "managed_instance_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.ManagedInstanceGroupLifecycleStateActive),
	},
}

var exportOsManagementHubManagementStationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_management_station",
	DatasourceClass:        "oci_os_management_hub_management_stations",
	DatasourceItemsAttr:    "management_station_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "management_station",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.ManagementStationLifecycleStateActive),
	},
}

var exportOsManagementHubProfileHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_profile",
	DatasourceClass:        "oci_os_management_hub_profiles",
	DatasourceItemsAttr:    "profile_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "profile",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.ProfileLifecycleStateActive),
	},
}

var osManagementHubResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOsManagementHubLifecycleEnvironmentHints},
		{TerraformResourceHints: exportOsManagementHubSoftwareSourceHints},
		{TerraformResourceHints: exportOsManagementHubManagedInstanceGroupHints},
		{TerraformResourceHints: exportOsManagementHubManagementStationHints},
		{TerraformResourceHints: exportOsManagementHubProfileHints},
	},
}
