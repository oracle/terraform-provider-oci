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
* `compartment_id` - (Required) (Updatable) Compartment Identifier. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `diagnostics_collection` - (Optional) (Updatable) Details to configure diagnostics collection for targets affected by this Exadata Fleet Update Maintenance Cycle. 
	* `log_collection_mode` - (Optional) (Updatable) Enable incident logs and trace collection.  Allow Oracle to collect incident logs and traces to enable fault diagnosis and issue resolution according to the selected mode. 
* `display_name` - (Optional) (Updatable) Exadata Fleet Update Cycle display name. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fsu_collection_id` - (Required) OCID identifier for the Collection ID the Exadata Fleet Update Cycle will be assigned to. 
* `goal_version_details` - (Required) (Updatable) Goal version or image details for the Exadata Fleet Update Cycle. 
	* `home_policy` - (Optional) (Updatable) Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
	* `new_home_prefix` - (Optional) (Updatable) Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `software_image_id` - (Required when type=IMAGE_ID) (Updatable) Target database software image OCID. 
	* `type` - (Required) (Updatable) Type of goal target version specified 
	* `version` - (Required when type=VERSION) (Updatable) Target DB or GI version string for the Exadata Fleet Update Cycle. 
* `is_ignore_missing_patches` - (Optional) (Updatable) List of patch IDs to ignore. 
* `is_ignore_patches` - (Optional) (Updatable) Ignore all patches between the source and target homes during patching. 
* `is_keep_placement` - (Optional) (Updatable) Ensure that services of administrator-managed Oracle RAC or Oracle RAC One databases are running on the same instances before and after the move operation. 
* `max_drain_timeout_in_seconds` - (Optional) (Updatable) Service drain timeout specified in seconds. 
* `stage_action_schedule` - (Optional) Scheduling related details for the Exadata Fleet Update Action during create operations. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails for Stage and Apply Actions in Exadata Fleet Update Cycle creation would not create Actions. Null scheduleDetails for CreateAction would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - (Required) The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - (Required) Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `type` - (Required) (Updatable) Type of Exadata Fleet Update Cycle. 


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
* `collection_type` - Type of Collection this Exadata Fleet Update Cycle belongs to. 
* `compartment_id` - Compartment Identifier. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `diagnostics_collection` - Details to configure diagnostics collection for targets affected by this Exadata Fleet Update Maintenance Cycle. 
	* `log_collection_mode` - Enable incident logs and trace collection.  Allow Oracle to collect incident logs and traces to enable fault diagnosis and issue resolution according to the selected mode. 
* `display_name` - Exadata Fleet Update Cycle display name. 
* `executing_fsu_action_id` - OCID identifier for the Action that is currently in execution, if applicable. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `fsu_collection_id` - OCID identifier for the Collection ID the Exadata Fleet Update Cycle is assigned to. 
* `goal_version_details` - Goal version or image details for the Exadata Fleet Update Cycle. 
	* `home_policy` - Goal home policy to use when Staging the Goal Version during patching. CREATE_NEW: Create a new DBHome (for Database Collections) for the specified image or version. USE_EXISTING: All database targets in the same VMCluster or CloudVmCluster will be moved to a shared database home.  If an existing home for the selected image or version is not found in the VM Cluster for a target database, then a new home will be created.  If more than one existing home for the selected image is found, then the home with the least number of databases will be used.  If multiple homes have the least number of databases, then a home will be selected at random. 
	* `new_home_prefix` - Prefix name used for new DB home resources created as part of the Stage Action. Format: <specified_prefix>_<timestamp> If not specified, a default Oracle Cloud Infrastructure DB home resource will be generated for the new DB home resources created. 
	* `software_image_id` - Target database software image OCID. 
	* `type` - Type of goal target version specified 
	* `version` - Target DB or GI version string for the Exadata Fleet Update Cycle. 
* `id` - OCID identifier for the Exadata Fleet Update Cycle. 
* `is_ignore_missing_patches` - List of bug numbers to ignore. 
* `is_ignore_patches` - Ignore all patches between the source and target homes during patching. 
* `is_keep_placement` - Ensure that services of administrator-managed Oracle RAC or Oracle RAC One databases are running on the same instances before and after the move operation. 
* `last_completed_action` - The latest Action type that was completed in the Exadata Fleet Update Cycle. No value would indicate that the Cycle has not completed any Action yet. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `max_drain_timeout_in_seconds` - Service drain timeout specified in seconds. 
* `next_action_to_execute` - In this array all the possible actions will be listed. The first element is the suggested Action. 
	* `time_to_start` - The date and time the Exadata Fleet Update Action is expected to start. Null if no Action has been scheduled. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of Exadata Fleet Update Action 
* `stage_action_schedule` - Scheduling related details for the Exadata Fleet Update Action. The specified time should not conflict with existing Exadata Infrastructure maintenance windows. Null scheduleDetails would execute the Exadata Fleet Update Action as soon as possible. 
	* `time_to_start` - The date and time the Exadata Fleet Update Action is expected to start. [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
	* `type` - Type of scheduling strategy to use for Fleet Patching Update Action execution. 
* `state` - The current state of the Exadata Fleet Update Cycle.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Exadata Fleet Update Cycle was created, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `time_finished` - The date and time the Exadata Fleet Update Cycle was finished, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The date and time the Exadata Fleet Update Cycle was updated, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `type` - Type of Exadata Fleet Update Cycle. 

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

