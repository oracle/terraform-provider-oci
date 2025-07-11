---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_high_availability_members"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_high_availability_members"
description: |-
  Provides the list of Managed My Sql Database High Availability Members in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_high_availability_members
This data source provides the list of Managed My Sql Database High Availability Members in Oracle Cloud Infrastructure Database Management service.

Retrieves information about the high availability members of a specific MySQL server's replication group.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_high_availability_members" "test_managed_my_sql_database_high_availability_members" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `managed_my_sql_database_high_availability_member_collection` - The list of managed_my_sql_database_high_availability_member_collection.

### ManagedMySqlDatabaseHighAvailabilityMember Reference

The following attributes are exported:

* `flow_control` - The mode used for flow control.
* `group_auto_increment` - The interval between successive values for auto-incremented columns for transactions that execute on this server instance.
* `group_name` - The name of the group to which this server instance belongs.
* `items` - A list of MySqlHighAvailabilityMember records.
	* `member_host` - The host name of the group member that clients use to connect to it.
	* `member_port` - The port number of the group member that clients use to connect to it.
	* `member_role` - The current role of the group member in the group.
	* `member_state` - The current state of the group member.
	* `member_uuid` - The Universally Unique Identifier (UUID) of the member server.
* `member_role` - The role of this server as a group replication member.
* `member_state` - The state of this server as a group replication member.
* `single_primary_mode` - Indicates if the replication group is running in single-primary mode.
* `status_summary` - High availability status summary of a MySQL server.
	* `channel_apply_errors` - A list of MySqlChannelApplyError records.
		* `apply_error` - Error from the apply operation of a MySQL server replication channel.
			* `last_error_message` - The error message of the most recent error that caused the SQL or coordinator thread to stop.
			* `last_error_number` - The error number of the most recent error that caused the SQL or coordinator thread to stop.
			* `time_last_error` - The timestamp when the most recent SQL or coordinator error occurred.
			* `worker_errors` - A list of MySqlApplyErrorWorker records.
				* `last_error_message` - The error message of the most recent error that caused the worker thread to stop.
				* `last_error_number` - The error number of the most recent error that caused the worker thread to stop.
				* `time_last_error` - The timestamp when the most recent worker error occurred.
		* `channel_name` - The name of the replication channel.
	* `channel_fetch_errors` - A list of MySqlChannelFetchError records.
		* `channel_name` - The name of the replication channel.
		* `fetch_error` - Error from the fetch operation of a MySQL server replication channel.
			* `last_error_message` - The error message of the most recent error that caused the I/O thread to stop.
			* `last_error_number` - The error number of the most recent error that caused the I/O thread to stop.
			* `time_last_error` - The timestamp when the most recent I/O error occurred.
* `transactions_in_gtid_executed` - The number of transactions that were replicated within the cluster.
* `view_id` - The current view identifier for this group.

