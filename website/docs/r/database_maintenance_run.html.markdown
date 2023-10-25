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

Creates a maintenance run with one of the following:
The latest available release update patch (RUP) for the Autonomous Container Database.
The latest available RUP and DST time zone (TZ) file updates for the Autonomous Container Database.
Creates a maintenance run to update the DST TZ file for the Autonomous Container Database.


## Example Usage

```hcl
resource "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	patch_type = var.maintenance_run_patch_type
	target_resource_id = oci_usage_proxy_resource.test_resource.id
	time_scheduled = var.maintenance_run_time_scheduled

	#Optional
	compartment_id = var.compartment_id
	is_dst_file_update_enabled = var.maintenance_run_is_dst_file_update_enabled
	patching_mode = var.maintenance_run_patching_mode
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Maintenance Run.
* `is_dst_file_update_enabled` - (Optional) Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
* `patch_type` - (Required) Patch type, either "QUARTERLY" or "TIMEZONE". 
* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

	*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
* `target_resource_id` - (Required) The ID of the target resource for which the maintenance run should be created.
* `time_scheduled` - (Required) (Updatable) The date and time that update should be scheduled.


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
* `is_dst_file_update_enabled` - Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
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

