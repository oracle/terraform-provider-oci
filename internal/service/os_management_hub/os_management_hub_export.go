package os_management_hub

import (
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	//tf_export.RegisterCompartmentGraphs("os_management_hub", osManagementHubResourceGraph)
}

// Hints for discovering and exporting this resource to configuration and state files
var exportOsManagementHubManagedInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_managed_instance",
	DatasourceClass:        "oci_os_management_hub_managed_instances",
	DatasourceItemsAttr:    "managed_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "managed_instance",
	RequireResourceRefresh: true,
}

var exportOsManagementHubManagedInstanceUpdatePackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_update_packages_management",
	ResourceAbbreviation: "managed_instance_update_packages_management",
}

var exportOsManagementHubManagedInstanceInstallWindowsUpdatesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_install_windows_updates_management",
	ResourceAbbreviation: "managed_instance_install_windows_updates_management",
}

var exportOsManagementHubManagedInstanceAttachProfileManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_attach_profile_management",
	ResourceAbbreviation: "managed_instance_attach_profile_management",
}

var exportOsManagementHubManagedInstanceDetachProfileManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_detach_profile_management",
	ResourceAbbreviation: "managed_instance_detach_profile_management",
}

var exportOsManagementHubManagedInstanceRebootManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_reboot_management",
	ResourceAbbreviation: "managed_instance_reboot_management",
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

var exportOsManagementHubProfileAttachLifecycleStageManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_profile_attach_lifecycle_stage_management",
	ResourceAbbreviation: "profile_attach_lifecycle_stage_management",
}

var exportOsManagementHubProfileAttachManagedInstanceGroupManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_profile_attach_managed_instance_group_management",
	ResourceAbbreviation: "profile_attach_managed_instance_group_management",
}

var exportOsManagementHubProfileAttachManagementStationManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_profile_attach_management_station_management",
	ResourceAbbreviation: "profile_attach_management_station_management",
}

var exportOsManagementHubProfileAttachSoftwareSourcesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_profile_attach_software_sources_management",
	ResourceAbbreviation: "profile_attach_software_sources_management",
}

var exportOsManagementHubProfileDetachSoftwareSourcesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_profile_detach_software_sources_management",
	ResourceAbbreviation: "profile_detach_software_sources_management",
}

var exportOsManagementHubScheduledJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_scheduled_job",
	DatasourceClass:        "oci_os_management_hub_scheduled_jobs",
	DatasourceItemsAttr:    "scheduled_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "scheduled_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.ScheduledJobLifecycleStateActive),
	},
}

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

var exportOsManagementHubLifecycleStageAttachManagedInstancesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_lifecycle_stage_attach_managed_instances_management",
	ResourceAbbreviation: "lifecycle_stage_attach_managed_instances_management",
}

var exportOsManagementHubLifecycleStageDetachManagedInstancesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_lifecycle_stage_detach_managed_instances_management",
	ResourceAbbreviation: "lifecycle_stage_detach_managed_instances_management",
}

var exportOsManagementHubLifecycleStagePromoteSoftwareSourceManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_lifecycle_stage_promote_software_source_management",
	ResourceAbbreviation: "lifecycle_stage_promote_software_source_management",
}

var exportOsManagementHubLifecycleStageRebootManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_lifecycle_stage_reboot_management",
	ResourceAbbreviation: "lifecycle_stage_reboot_management",
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

var exportOsManagementHubManagedInstanceGroupAttachManagedInstancesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_attach_managed_instances_management",
	ResourceAbbreviation: "managed_instance_group_attach_managed_instances_management",
}

var exportOsManagementHubManagedInstanceGroupDetachManagedInstancesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_detach_managed_instances_management",
	ResourceAbbreviation: "managed_instance_group_detach_managed_instances_management",
}

var exportOsManagementHubManagedInstanceGroupAttachSoftwareSourcesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_attach_software_sources_management",
	ResourceAbbreviation: "managed_instance_group_attach_software_sources_management",
}

var exportOsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_detach_software_sources_management",
	ResourceAbbreviation: "managed_instance_group_detach_software_sources_management",
}

var exportOsManagementHubManagedInstanceGroupInstallPackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_install_packages_management",
	ResourceAbbreviation: "managed_instance_group_install_packages_management",
}

var exportOsManagementHubManagedInstanceGroupRemovePackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_remove_packages_management",
	ResourceAbbreviation: "managed_instance_group_remove_packages_management",
}

var exportOsManagementHubManagedInstanceGroupUpdateAllPackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_update_all_packages_management",
	ResourceAbbreviation: "managed_instance_group_update_all_packages_management",
}

var exportOsManagementHubManagedInstanceGroupManageModuleStreamsManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_manage_module_streams_management",
	ResourceAbbreviation: "managed_instance_group_manage_module_streams_management",
}

var exportOsManagementHubManagedInstanceGroupInstallWindowsUpdatesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_install_windows_updates_management",
	ResourceAbbreviation: "managed_instance_group_install_windows_updates_management",
}

var exportOsManagementHubManagedInstanceGroupRebootManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_managed_instance_group_reboot_management",
	ResourceAbbreviation: "managed_instance_group_reboot_management",
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

var exportOsManagementHubManagementStationRefreshManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_management_station_refresh_management",
	ResourceAbbreviation: "management_station_refresh_management",
}

var exportOsManagementHubManagementStationAssociateManagedInstancesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_management_station_associate_managed_instances_management",
	ResourceAbbreviation: "management_station_associate_managed_instances_management",
}

var exportOsManagementHubManagementStationSynchronizeMirrorsManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_management_station_synchronize_mirrors_management",
	ResourceAbbreviation: "management_station_synchronize_mirrors_management",
}

var exportOsManagementHubManagementStationMirrorSynchronizeManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_management_station_mirror_synchronize_management",
	ResourceAbbreviation: "management_station_mirror_synchronize_management",
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
		string(oci_os_management_hub.SoftwareSourceLifecycleStateNeedsAttention),
	},
}

var exportOsManagementHubSoftwareSourceChangeAvailabilityManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_software_source_change_availability_management",
	ResourceAbbreviation: "software_source_change_availability_management",
}

var exportOsManagementHubSoftwareSourceAddPackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_software_source_add_packages_management",
	ResourceAbbreviation: "software_source_add_packages_management",
}

var exportOsManagementHubSoftwareSourceRemovePackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_software_source_remove_packages_management",
	ResourceAbbreviation: "software_source_remove_packages_management",
}

var exportOsManagementHubSoftwareSourceReplacePackagesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_software_source_replace_packages_management",
	ResourceAbbreviation: "software_source_replace_packages_management",
}

var exportOsManagementHubSoftwareSourceGenerateMetadataManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_software_source_generate_metadata_management",
	ResourceAbbreviation: "software_source_generate_metadata_management",
}

var exportOsManagementHubEventHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_os_management_hub_event",
	DatasourceClass:        "oci_os_management_hub_events",
	DatasourceItemsAttr:    "event_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "event",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_os_management_hub.EventLifecycleStateActive),
	},
}

var exportOsManagementHubWorkRequestRerunManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_os_management_hub_work_request_rerun_management",
	ResourceAbbreviation: "work_request_rerun_management",
}

var osManagementHubResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOsManagementHubManagedInstanceHints},
		{TerraformResourceHints: exportOsManagementHubProfileHints},
		{TerraformResourceHints: exportOsManagementHubScheduledJobHints},
		{TerraformResourceHints: exportOsManagementHubLifecycleEnvironmentHints},
		{TerraformResourceHints: exportOsManagementHubManagedInstanceGroupHints},
		{TerraformResourceHints: exportOsManagementHubManagementStationHints},
		{TerraformResourceHints: exportOsManagementHubSoftwareSourceHints},
		{TerraformResourceHints: exportOsManagementHubEventHints},
	},
	/*
		Below code is causing infinite loop in resource discovery
		"oci_os_management_hub_software_source": {
			{
				TerraformResourceHints: exportOsManagementHubSoftwareSourceHints,
				DatasourceQueryParams: map[string]string{
					"software_source_id": "id",
				},
			},
		},
	*/
}
