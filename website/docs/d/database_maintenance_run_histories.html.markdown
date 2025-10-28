---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run_histories"
sidebar_current: "docs-oci-datasource-database-maintenance_run_histories"
description: |-
  Provides the list of Maintenance Run Histories in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_maintenance_run_histories
This data source provides the list of Maintenance Run Histories in Oracle Cloud Infrastructure Database service.

Gets a list of the maintenance run histories in the specified compartment.


## Example Usage

```hcl
data "oci_database_maintenance_run_histories" "test_maintenance_run_histories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.maintenance_run_history_availability_domain
	maintenance_type = var.maintenance_run_history_maintenance_type
	state = var.maintenance_run_history_state
	target_resource_id = oci_database_target_resource.test_target_resource.id
	target_resource_type = var.maintenance_run_history_target_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `maintenance_type` - (Optional) The maintenance type.
* `state` - (Optional) The state of the maintenance run history.
* `target_resource_id` - (Optional) The target resource ID.
* `target_resource_type` - (Optional) The type of the target resource.


## Attributes Reference

The following attributes are exported:

* `maintenance_run_histories` - The list of maintenance_run_histories.

### MaintenanceRunHistory Reference

The following attributes are exported:

* `current_execution_window` - The OCID of the current execution window.
* `db_servers_history_details` - List of database server history details.
	* `db_server_patching_details` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
		* `estimated_patch_duration` - Estimated time, in minutes, to patch one database server.
		* `patching_status` - The status of the patching operation.
		* `time_patching_ended` - The time when the patching operation ended.
		* `time_patching_started` - The time when the patching operation started.
	* `display_name` - The user-friendly name for the database server. The name does not need to be unique.
	* `id` - The OCID of the database server.
* `granular_maintenance_history` - The list of granular maintenance history details.
	* `execution_actions` - The list of execution actions for this granular maintenance history.
		* `action_members` - List of action members of this execution action.
			* `estimated_time_in_mins` - The estimated time of the execution action member in minutes.
			* `member_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent resource the execution action belongs to.
			* `member_order` - The priority order of the execution action member.
			* `status` - The current status of the execution action member. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, DURATION_EXCEEDED, RESCHEDULED and COMPLETED. enum:
				* SCHEDULED
				* IN_PROGRESS
				* FAILED
				* CANCELED
				* DURATION_EXCEEDED
				* RESCHEDULED
				* SUCCEEDED 
			* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.
		* `action_params` - Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
		* `action_type` - The action type of the execution action being performed
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `description` - Description of the execution action.
		* `display_name` - The user-friendly name for the execution action. The name does not need to be unique.
		* `estimated_time_in_mins` - The estimated time of the execution action in minutes.
		* `execution_action_order` - The priority order of the execution action.
		* `execution_window_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution window resource the execution action belongs to.
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution action.
		* `lifecycle_details` - Additional information about the current lifecycle state.
		* `lifecycle_substate` - The current sub-state of the execution action. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING. 
		* `state` - The current state of the execution action. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS. 
		* `time_created` - The date and time the execution action was created.
		* `time_updated` - The last date and time that the execution action was updated.
		* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.
	* `execution_window` - Details of an execution window. 
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `description` - Description of the execution window.
		* `display_name` - The user-friendly name for the execution window. The name does not need to be unique.
		* `estimated_time_in_mins` - The estimated time of the execution window in minutes.
		* `execution_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution resource the execution window belongs to.
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution window.
		* `is_enforced_duration` - Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
		* `lifecycle_details` - Additional information about the current lifecycle state.
		* `lifecycle_substate` - The current sub-state of the execution window. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING. 
		* `state` - The current state of the Schedule Policy. Valid states are CREATED, SCHEDULED, IN_PROGRESS, FAILED, CANCELED, UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS. 
		* `time_created` - The date and time the execution window was created.
		* `time_ended` - The date and time that the execution window ended.
		* `time_scheduled` - The scheduled start date and time of the execution window.
		* `time_started` - The date and time that the execution window was started.
		* `time_updated` - The last date and time that the execution window was updated.
		* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.
		* `window_duration_in_mins` - Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes. 
		* `window_type` - The execution window is of PLANNED or UNPLANNED type.
* `id` - The OCID of the maintenance run history.
* `maintenance_run_details` - Details of a maintenance run. 
	* `compartment_id` - The OCID of the compartment.
	* `current_custom_action_timeout_in_mins` - Extend current custom action timeout between the current database servers during waiting state, from 0 (zero) to 30 minutes.
	* `current_patching_component` - The name of the current infrastruture component that is getting patched.
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Specify a number of minutes, from 15 to 120. 
	* `database_software_image_id` - The Autonomous AI Database Software Image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
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
	* `is_maintenance_run_granular` - If `FALSE`, the maintenance run doesn't support granular maintenance.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `maintenance_subtype` - Maintenance sub-type.
	* `maintenance_type` - Maintenance type.
	* `patch_failure_count` - Contain the patch failure count.
	* `patch_id` - The unique identifier of the patch. The identifier string includes the patch type, the Oracle AI Database version, and the patch creation date (using the format YYMMDD). For example, the identifier `ru_patch_19.9.0.0_201030` is used for an RU patch for Oracle AI Database 19.9.0.0 that was released October 30, 2020.
	* `patching_end_time` - The time when the patching operation ended.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `patching_start_time` - The time when the patching operation started.
	* `patching_status` - The status of the patching operation.
	* `peer_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	* `peer_maintenance_run_ids` - The list of OCIDs for the maintenance runs associated with their Autonomous Data Guard peer container databases.
	* `state` - The current state of the maintenance run. For Autonomous AI Database Serverless instances, valid states are IN_PROGRESS, SUCCEEDED, and FAILED. 
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
	* `target_db_server_version` - The target software version for the database server patching operation.
	* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
	* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
	* `target_storage_server_version` - The target Cell version that is to be patched to.
	* `time_ended` - The date and time the maintenance run was completed.
	* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
	* `time_started` - The date and time the maintenance run starts.
	* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.

