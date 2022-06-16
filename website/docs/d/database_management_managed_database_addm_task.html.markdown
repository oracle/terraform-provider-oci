---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_addm_task"
sidebar_current: "docs-oci-datasource-database_management-managed_database_addm_task"
description: |-
  Provides details about a specific Managed Database Addm Task in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_addm_task
This data source provides details about a specific Managed Database Addm Task resource in Oracle Cloud Infrastructure Database Management service.

Lists the metadata for each ADDM task who's end snapshot time falls within the provided start and end time. Details include
the name of the ADDM task, description, user, status and creation date time.


## Example Usage

```hcl
data "oci_database_management_managed_database_addm_task" "test_managed_database_addm_task" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	time_end = var.managed_database_addm_task_time_end
	time_start = var.managed_database_addm_task_time_start
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `time_end` - (Required) The end of the time range to search for ADDM tasks as defined by date-time RFC3339 format.
* `time_start` - (Required) The beginning of the time range to search for ADDM tasks as defined by date-time RFC3339 format.


## Attributes Reference

The following attributes are exported:

* `items` - The list of ADDM task metadata.
	* `begin_snapshot_id` - The ID number of the beginning AWR snapshot. 
	* `db_user` - The database user who owns the ADDM task.
	* `description` - The description of the ADDM task.
	* `end_snapshot_id` - The ID number of the ending AWR snapshot. 
	* `end_snapshot_time` - The timestamp of the ending AWR snapshot used in the ADDM task as defined by date-time RFC3339 format. 
	* `findings` - The number of ADDM findings.
	* `how_created` - A description of how the task was created.
	* `start_snapshot_time` - The timestamp of the beginning AWR snapshot used in the ADDM task as defined by date-time RFC3339 format. 
	* `status` - The status of the ADDM task.
	* `task_id` - The ID number of the ADDM task.
	* `task_name` - The name of the ADDM task.
	* `time_created` - The creation date of the ADDM task.
* `managed_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.

