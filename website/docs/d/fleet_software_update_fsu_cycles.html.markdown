---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_cycles"
sidebar_current: "docs-oci-datasource-fleet_software_update-fsu_cycles"
description: |-
  Provides the list of Fsu Cycles in Oracle Cloud Infrastructure Fleet Software Update service
---

# Data Source: oci_fleet_software_update_fsu_cycles
This data source provides the list of Fsu Cycles in Oracle Cloud Infrastructure Fleet Software Update service.

Gets a list of all Exadata Fleet Update Cycles in a compartment.


## Example Usage

```hcl
data "oci_fleet_software_update_fsu_cycles" "test_fsu_cycles" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	collection_type = var.fsu_cycle_collection_type
	display_name = var.fsu_cycle_display_name
	fsu_collection_id = oci_fleet_software_update_fsu_collection.test_fsu_collection.id
	state = var.fsu_cycle_state
	target_version = var.fsu_cycle_target_version
}
```

## Argument Reference

The following arguments are supported:

* `collection_type` - (Optional) A filter to return only resources whose Collection type matches the given type. 
* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `fsu_collection_id` - (Optional) A filter to return only resources whose fsuCollectionId matches the given fsuCollectionId. 
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState. 
* `target_version` - (Optional) A filter to return only entries whose targetVersion matches the given targetVersion. 


## Attributes Reference

The following attributes are exported:

* `fsu_cycle_summary_collection` - The list of fsu_cycle_summary_collection.

### FsuCycle Reference

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

