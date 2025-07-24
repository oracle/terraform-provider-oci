---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_inbound_replications"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_inbound_replications"
description: |-
  Provides the list of Managed My Sql Database Inbound Replications in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_inbound_replications
This data source provides the list of Managed My Sql Database Inbound Replications in Oracle Cloud Infrastructure Database Management service.

Retrieves information about the inbound replications of a specific MySQL server.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_inbound_replications" "test_managed_my_sql_database_inbound_replications" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `managed_my_sql_database_inbound_replication_collection` - The list of managed_my_sql_database_inbound_replication_collection.

### ManagedMySqlDatabaseInboundReplication Reference

The following attributes are exported:

* `inbound_replications_count` - The number of sources this server is replicating from.
* `items` - A list of ManagedMySqlDatabaseInboundReplicationSummary records.
	* `applier_filters` - A list of MySqlReplicationApplierFilter records.
		* `filter_name` - The type of replication filter that has been configured for the replication channel.
		* `filter_rule` - The rules configured for the replication filter type.
	* `apply_delay` - The time in seconds that the current transaction took between being committed on the source and being applied on the replica.
	* `apply_error` - Error from the apply operation of a MySQL server replication channel.
		* `last_error_message` - The error message of the most recent error that caused the SQL or coordinator thread to stop.
		* `last_error_number` - The error number of the most recent error that caused the SQL or coordinator thread to stop.
		* `time_last_error` - The timestamp when the most recent SQL or coordinator error occurred.
		* `worker_errors` - A list of MySqlApplyErrorWorker records.
			* `last_error_message` - The error message of the most recent error that caused the worker thread to stop.
			* `last_error_number` - The error number of the most recent error that caused the worker thread to stop.
			* `time_last_error` - The timestamp when the most recent worker error occurred.
	* `apply_status` - The current status of apply operations.
	* `busy_workers` - The number of workers currently busy applying transactions from the source server.
	* `channel_name` - The name of the replication channel.
	* `desired_delay_seconds` - The desired number of seconds that the replica must lag the source.
	* `fetch_error` - Error from the fetch operation of a MySQL server replication channel.
		* `last_error_message` - The error message of the most recent error that caused the I/O thread to stop.
		* `last_error_number` - The error number of the most recent error that caused the I/O thread to stop.
		* `time_last_error` - The timestamp when the most recent I/O error occurred.
	* `fetch_status` - The current status of fetch operations.
	* `gtid_assignment` - Indicates whether the channel assigns global transaction identifiers (GTIDs) to anonymous replicated transactions. OFF means no GTIDs are assigned. LOCAL means a GTID is assigned that includes this replica's own universally unique identifier (UUID). A UUID as value indicates that a GTID is assigned, which includes that manually set UUID value.
	* `relay_log_storage_space_used` - The total size in bytes of all the existing relay log files pertaining to this channel.
	* `remaining_delay_seconds` - If the replica is waiting for the desired delay seconds to pass since the source applied an event, this field contains the number of delay seconds remaining.
	* `retrieved_gtid_set` - The set of global transaction IDs corresponding to all transactions received by this replica from the source server. Empty if GTIDs are not in use.
	* `seconds_behind_source` - The number of seconds the replica is behind the source server.
	* `source_host` - The host name or IP address of the source this replica is connected to.
	* `source_port` - The port used to connect to the source.
	* `source_server_id` - The server ID value from the source server.
	* `source_uuid` - The Universally Unique Identifier (UUID) value from the source server.
	* `transactions_received` - The number of transactions received by this replica from the source server.
* `parallel_workers` - The number of applier threads for executing replication transactions in parallel.
* `preserve_commit_order` - For multi-threaded replicas, indicates if transactions are executed and committed on the replica in the same order as they appear in the relay log.
* `replica_server_id` - The server ID value of this replica.
* `replica_uuid` - The Universally Unique Identifier (UUID) value of this replica server.

