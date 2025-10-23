---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_cycle"
sidebar_current: "docs-oci-resource-fleet_software_update-fsu_cycle"
description: |-
  Provides the Fsu Cycle resource in Oracle Cloud Infrastructure Fleet Software Update service
---

# oci_fleet_software_update_fsu_cycle
This resource provides the Fsu Cycle resource in Oracle Cloud Infrastructure Fleet Software Update service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/edsfu/latest/FsuCycle

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleetsoftwareupdate

Creates a new Exadata Fleet Update Cycle.


## Example Usage

```hcl
resource "oci_fleet_software_update_fsu_cycle" "test_fsu_cycle" {
	#Required
	compartment_id = var.compartment_id
	fsu_collection_id = oci_fleet_software_update_fsu_collection.test_fsu_collection.id
	goal_version_details {
		#Required
		type = var.fsu_cycle_goal_version_details_type

		#Optional
		components {
			#Required
			component_type = var.fsu_cycle_goal_version_details_components_component_type
			goal_version_details {
				#Required
				goal_software_image_id = oci_core_image.test_image.id
				goal_type = var.fsu_cycle_goal_version_details_components_goal_version_details_goal_type

				#Optional
				goal_version = var.fsu_cycle_goal_version_details_components_goal_version_details_goal_version
			}

			#Optional
			home_policy = var.fsu_cycle_goal_version_details_components_home_policy
			new_home_prefix = var.fsu_cycle_goal_version_details_components_new_home_prefix
		}
		home_policy = var.fsu_cycle_goal_version_details_home_policy
		new_home_prefix = var.fsu_cycle_goal_version_details_new_home_prefix
		software_image_id = oci_core_image.test_image.id
		version = var.fsu_cycle_goal_version_details_version
	}
	type = var.fsu_cycle_type

	#Optional
	apply_action_schedule {
		#Required
		time_to_start = var.fsu_cycle_apply_action_schedule_time_to_start
		type = var.fsu_cycle_apply_action_schedule_type
	}
	batching_strategy {

		#Optional
		is_force_rolling = var.fsu_cycle_batching_strategy_is_force_rolling
		is_wait_for_batch_resume = var.fsu_cycle_batching_strategy_is_wait_for_batch_resume
		percentage = var.fsu_cycle_batching_strategy_percentage
		type = var.fsu_cycle_batching_strategy_type
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	diagnostics_collection {

		#Optional
		log_collection_mode = var.fsu_cycle_diagnostics_collection_log_collection_mode
	}
	display_name = var.fsu_cycle_display_name
	freeform_tags = {"bar-key"= "value"}
	is_ignore_missing_patches = var.fsu_cycle_is_ignore_missing_patches
	is_ignore_patches = var.fsu_cycle_is_ignore_patches
	is_keep_placement = var.fsu_cycle_is_keep_placement
	max_drain_timeout_in_seconds = var.fsu_cycle_max_drain_timeout_in_seconds
	stage_action_schedule {
		#Required
		time_to_start = var.fsu_cycle_stage_action_schedule_time_to_start
		type = var.fsu_cycle_stage_action_schedule_type
	}
	upgrade_details {
		#Required
		collection_type = var.fsu_cycle_upgrade_details_collection_type

		#Optional
		is_ignore_post_upgrade_errors = var.fsu_cycle_upgrade_details_is_ignore_post_upgrade_errors
		is_ignore_prerequisites = var.fsu_cycle_upgrade_details_is_ignore_prerequisites
		is_recompile_invalid_objects = var.fsu_cycle_upgrade_details_is_recompile_invalid_objects
		is_time_zone_upgrade = var.fsu_cycle_upgrade_details_is_time_zone_upgrade
		max_drain_timeout_in_seconds = var.fsu_cycle_upgrade_details_max_drain_timeout_in_seconds
	}
}
```

## Argument Reference

The following arguments are supported:

* `apply_action_schedule` - (Optional) Scheduling related details for the Exadata Fleet Update Action during create operations. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails for Stage and Apply Actions in Exadata Fleet Update Cycle creation would not create Actions. Null scheduleDetails for CreateAction would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - (Required) The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - (Required) Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `batching_strategy` - (Optional) (Updatable) Batching strategy details to use during PRECHECK and APPLY Cycle Actions. 
	* `is_force_rolling` - (Applicable when type=FIFTY_FIFTY | SEQUENTIAL | SERVICE_AVAILABILITY_FACTOR) (Updatable) True to force rolling patching. 
	* `is_wait_for_batch_resume` - (Applicable when type=FIFTY_FIFTY) (Updatable) True to wait for customer to resume the Apply Action once the first half is done. False to automatically patch the second half. 
	* `percentage` - (Applicable when type=SERVICE_AVAILABILITY_FACTOR) (Updatable) Percentage of availability in the service during the Patch operation. 
	* `type` - (Optional) (Updatable) Supported batching strategies. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `diagnostics_collection` - (Optional) (Updatable) Details to configure diagnostics collection for targets affected by this Exadata Fleet Update Maintenance Cycle. 
	* `log_collection_mode` - (Optional) (Updatable) Enable incident logs and trace collection.  Allow Oracle to collect incident logs and traces to enable fault diagnosis and issue resolution according to the selected mode. 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Exadata Fleet Update Cycle. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fsu_collection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Collection which will be updated by the Exadata Fleet Update Cycle being created. 
* `goal_version_details` - (Required) (Updatable) Goal version or image details for the Exadata Fleet Update Cycle. 
	* `components` - (Required when type=EXADB_STACK) (Updatable) Details of goal versions for components in an Exadata software stack. 
		* `component_type` - (Required) (Updatable) Type of component in an Exadata software stack. 
		* `goal_version_details` - (Required) (Updatable) Details of goal 'GUEST_OS' software version. 
			* `goal_software_image_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom 'GI' software image. 
			* `goal_type` - (Required) (Updatable) Preference to use an Oracle released 'GI' software image or a custom 'GI' software image. 
			* `goal_version` - (Required when goal_type=GI_ORACLE_IMAGE | GUEST_OS_ORACLE_IMAGE) (Updatable) Goal version string matching an Oracle released 'GUEST_OS' software image. 
		* `home_policy` - (Applicable when component_type=GI) (Updatable) Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
		* `new_home_prefix` - (Applicable when component_type=GI) (Updatable) Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `home_policy` - (Optional) (Updatable) Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
	* `new_home_prefix` - (Optional) (Updatable) Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `software_image_id` - (Required when type=IMAGE_ID) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the goal database software image. 
	* `type` - (Required) (Updatable) Type of goal version specified 
	* `version` - (Required when type=VERSION) (Updatable) Goal version string for the Exadata Fleet Update Cycle. Applicable to Database, Grid Infrastructure, or Exadata Image software updates. 
* `is_ignore_missing_patches` - (Applicable when type=PATCH) (Updatable) List of identifiers of patches to ignore. This attribute will be ignored for Exadata Image (Guest OS) maintenance update. 
* `is_ignore_patches` - (Applicable when type=PATCH) (Updatable) Ignore patch conflicts or missing patches between the source and goal homes. This attribute will be ignored for Exadata Image (Guest OS) maintenance update. 
* `is_keep_placement` - (Applicable when type=PATCH) (Updatable) Ensure that database services are online on the same VMs before and after the maintenance update. 
* `max_drain_timeout_in_seconds` - (Applicable when type=PATCH) (Updatable) Timeout for session draining for database services specified in seconds. 
* `stage_action_schedule` - (Optional) Scheduling related details for the Exadata Fleet Update Action during create operations. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails for Stage and Apply Actions in Exadata Fleet Update Cycle creation would not create Actions. Null scheduleDetails for CreateAction would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - (Required) The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - (Required) Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `type` - (Required) (Updatable) Type of Exadata Fleet Update Cycle. 
* `upgrade_details` - (Applicable when type=UPGRADE) (Updatable) Details of supported upgrade options for DB or GI collection. 
	* `collection_type` - (Required) (Updatable) Type of Exadata Fleet Update collection being upgraded. 
	* `is_ignore_post_upgrade_errors` - (Applicable when collection_type=GI) (Updatable) Ignore errors during post Oracle Grid Infrastructure upgrade Cluster Verification Utility (CVU) check. 
	* `is_ignore_prerequisites` - (Applicable when collection_type=GI) (Updatable) Ignore the Cluster Verification Utility (CVU) prerequisite checks. 
	* `is_recompile_invalid_objects` - (Applicable when collection_type=DB) (Updatable) Enables or disables the recompilation of invalid objects. 
	* `is_time_zone_upgrade` - (Applicable when collection_type=DB) (Updatable) Enables or disables time zone upgrade. 
	* `max_drain_timeout_in_seconds` - (Applicable when collection_type=DB) (Updatable) Service drain timeout specified in seconds. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apply_action_schedule` - Scheduling related details for the Exadata Fleet Update Action. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `batching_strategy` - Batching strategy details to use during PRECHECK and APPLY Cycle Actions. 
	* `is_force_rolling` - True to force rolling patching. 
	* `is_wait_for_batch_resume` - True to wait for customer to resume the Apply Action once the first half is done. False to automatically patch the second half. 
	* `percentage` - Percentage of availability in the service during the Patch operation. 
	* `type` - Supported batching strategies. 
* `collection_type` - Type of the Exadata Fleet Update Collection being updated by this Exadata Fleet Update Cycle. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `diagnostics_collection` - Details to configure diagnostics collection for targets affected by this Exadata Fleet Update Maintenance Cycle. 
	* `log_collection_mode` - Enable incident logs and trace collection.  Allow Oracle to collect incident logs and traces to enable fault diagnosis and issue resolution according to the selected mode. 
* `display_name` - The user-friendly name for the Exadata Fleet Update Cycle. 
* `executing_fsu_action_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Action that is currently in progress, if applicable. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fsu_collection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Collection being updated by this Exadata Fleet Update Cycle. 
* `goal_version_details` - Goal version or image details for the Exadata Fleet Update Cycle. 
	* `components` - Details of goal versions for components in an Exadata software stack. 
		* `component_type` - Type of component in an Exadata software stack. 
		* `goal_version_details` - Details of goal 'GUEST_OS' software version. 
			* `goal_software_image_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom 'GI' software image. 
			* `goal_type` - Preference to use an Oracle released 'GI' software image or a custom 'GI' software image. 
			* `goal_version` - Goal version string matching an Oracle released 'GUEST_OS' software image. 
		* `home_policy` - Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
		* `new_home_prefix` - Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `home_policy` - Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
	* `new_home_prefix` - Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `software_image_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the goal database software image. 
	* `type` - Type of goal version specified 
	* `version` - Goal version string for the Exadata Fleet Update Cycle. Applicable to Database, Grid Infrastructure, or Exadata Image software updates. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Cycle. 
* `is_ignore_missing_patches` - List of identifiers of patches to ignore. This attribute will be ignored for Exadata Image (Guest OS) maintenance update. 
* `is_ignore_patches` - Ignore patch conflicts or missing patches between the source and goal homes. This attribute will be ignored for Exadata Image (Guest OS) maintenance update. 
* `is_keep_placement` - Ensure that database services are online on the same VMs before and after the maintenance update. 
* `last_completed_action` - The latest Action type that was completed in the Exadata Fleet Update Cycle. No value would indicate that the Cycle has not completed any Action yet. 
* `last_completed_action_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the latest Action  in the Exadata Fleet Update Cycle. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `max_drain_timeout_in_seconds` - Timeout for session draining for database services specified in seconds. 
* `next_action_to_execute` - All possible Exadata Fleet Update Actions will be listed. The first element is the suggested Exadata Fleet Update Action. 
	* `time_to_start` - The date and time the Exadata Fleet Update Action is expected to start. Null if no Action has been scheduled. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of Exadata Fleet Update Action 
* `rollback_cycle_state` - Current rollback cycle state if rollback maintenance cycle action has been attempted. No value would indicate that the Cycle has not run a rollback maintenance cycle action before. 
* `stage_action_schedule` - Scheduling related details for the Exadata Fleet Update Action. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `state` - The current state of the Exadata Fleet Update Cycle.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Exadata Fleet Update Cycle was created, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `time_finished` - The date and time the Exadata Fleet Update Cycle was finished, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The date and time the Exadata Fleet Update Cycle was updated, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `type` - Type of Exadata Fleet Update Cycle. 
* `upgrade_details` - Details of supported upgrade options for DB or GI collection. 
	* `collection_type` - Type of Exadata Fleet Update collection being upgraded. 
	* `is_ignore_post_upgrade_errors` - Ignore errors during post Oracle Grid Infrastructure upgrade Cluster Verification Utility (CVU) check. 
	* `is_ignore_prerequisites` - Ignore the Cluster Verification Utility (CVU) prerequisite checks. 
	* `is_recompile_invalid_objects` - Enables or disables the recompilation of invalid objects. 
	* `is_time_zone_upgrade` - Enables or disables time zone upgrade. 
	* `max_drain_timeout_in_seconds` - Service drain timeout specified in seconds. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fsu Cycle
	* `update` - (Defaults to 20 minutes), when updating the Fsu Cycle
	* `delete` - (Defaults to 20 minutes), when destroying the Fsu Cycle


## Import

FsuCycles can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_software_update_fsu_cycle.test_fsu_cycle "id"
```

