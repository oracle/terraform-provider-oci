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

* `db_servers_history_details` - List of database server history details.
	* `db_server_patching_details` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
		* `estimated_patch_duration` - Estimated time, in minutes, to patch one database server.
		* `patching_status` - The status of the patching operation.
		* `time_patching_ended` - The time when the patching operation ended.
		* `time_patching_started` - The time when the patching operation started.
	* `display_name` - The user-friendly name for the database server. The name does not need to be unique.
	* `id` - The OCID of the database server.
* `id` - The OCID of the maintenance run history.
* `maintenance_run_details` - Details of a maintenance run. 
	* `compartment_id` - The OCID of the compartment.
	* `current_custom_action_timeout_in_mins` - Extend current custom action timeout between the current database servers during waiting state, from 0 (zero) to 30 minutes.
	* `current_patching_component` - The name of the current infrastruture component that is getting patched.
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Specify a number of minutes, from 15 to 120. 
	* `database_software_image_id` - The Autonomous Database Software Image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
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

