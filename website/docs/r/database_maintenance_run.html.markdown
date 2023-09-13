---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run"
sidebar_current: "docs-oci-resource-database-maintenance_run"
description: |-
  Provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service
---

# oci_database_maintenance_run
This resource provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service.

Updates the properties of a maintenance run, such as the state of a maintenance run.

## Example Usage

```hcl
resource "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	maintenance_run_id = oci_database_maintenance_run.test_maintenance_run.id

	#Optional
	current_custom_action_timeout_in_mins = var.maintenance_run_current_custom_action_timeout_in_mins
	custom_action_timeout_in_mins = var.maintenance_run_custom_action_timeout_in_mins
	is_custom_action_timeout_enabled = var.maintenance_run_is_custom_action_timeout_enabled
	is_enabled = var.maintenance_run_is_enabled
	is_patch_now_enabled = var.maintenance_run_is_patch_now_enabled
	is_resume_patching = var.maintenance_run_is_resume_patching
	patch_id = oci_database_patch.test_patch.id
	patching_mode = var.maintenance_run_patching_mode
	target_db_server_version = var.maintenance_run_target_db_server_version
	target_storage_server_version = var.maintenance_run_target_storage_server_version
	time_scheduled = var.maintenance_run_time_scheduled
}
```

## Argument Reference

The following arguments are supported:

* `current_custom_action_timeout_in_mins` - (Optional) (Updatable) The current custom action timeout between the current database servers during waiting state in addition to custom action timeout, from 0 (zero) to 30 minutes.
* `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Specify a number of minutes from 15 to 120. 
* `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database servers patching operations.
* `is_enabled` - (Optional) (Updatable) If `FALSE`, skips the maintenance run.
* `is_patch_now_enabled` - (Optional) (Updatable) If set to `TRUE`, starts patching immediately.
* `is_resume_patching` - (Optional) (Updatable) If true, then the patching is resumed and the next component will be patched immediately.
* `maintenance_run_id` - (Required) The maintenance run OCID.
* `patch_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

	*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
* `target_db_server_version` - (Optional) (Updatable) The target database server system software version for the patching operation.
* `target_storage_server_version` - (Optional) (Updatable) The target storage cell system software version for the patching operation.
* `time_scheduled` - (Optional) (Updatable) The scheduled date and time of the maintenance run to update.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `current_custom_action_timeout_in_mins` - Extend current custom action timeout between the current database servers during waiting state, from 0 (zero) to 30 minutes.
* `current_patching_component` - The name of the current infrastruture component that is getting patched.
* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Specify a number of minutes, from 15 to 120. 
* `description` - Description of the maintenance run.
* `display_name` - The user-friendly name for the maintenance run.
* `estimated_component_patching_start_time` - The estimated start time of the next infrastruture component patching operation.
* `estimated_patching_time` - The estimated total time required in minutes for all patching operations (database server, storage server, and network switch patching). 
	* `estimated_db_server_patching_time` - The estimated time required in minutes for database server patching.
	* `estimated_network_switches_patching_time` - The estimated time required in minutes for network switch patching.
	* `estimated_storage_server_patching_time` - The estimated time required in minutes for storage server patching.
	* `total_estimated_patching_time` - The estimated total time required in minutes for all patching operations.
* `id` - The OCID of the maintenance run.
* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database servers patching operations.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `patch_failure_count` - Contain the patch failure count.
* `patch_id` - The unique identifier of the patch. The identifier string includes the patch type, the Oracle Database version, and the patch creation date (using the format YYMMDD). For example, the identifier `ru_patch_19.9.0.0_201030` is used for an RU patch for Oracle Database 19.9.0.0 that was released October 30, 2020.
* `patching_end_time` - The time when the patching operation ended.
* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

	*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
* `patching_start_time` - The time when the patching operation started.
* `patching_status` - The status of the patching operation.
* `peer_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
* `state` - The current state of the maintenance run. For Autonomous Database Serverless instances, valid states are IN_PROGRESS, SUCCEEDED, and FAILED. 
* `target_db_server_version` - The target software version for the database server patching operation.
* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
* `target_storage_server_version` - The target Cell version that is to be patched to.
* `time_ended` - The date and time the maintenance run was completed.
* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
* `time_started` - The date and time the maintenance run starts.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Run
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Run
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Run


## Import

MaintenanceRuns can be imported using the `id`, e.g.

```
$ terraform import oci_database_maintenance_run.test_maintenance_run "id"
```

