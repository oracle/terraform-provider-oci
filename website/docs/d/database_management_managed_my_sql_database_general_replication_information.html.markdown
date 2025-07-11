---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_general_replication_information"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_general_replication_information"
description: |-
  Provides details about a specific Managed My Sql Database General Replication Information in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_general_replication_information
This data source provides details about a specific Managed My Sql Database General Replication Information resource in Oracle Cloud Infrastructure Database Management service.

Retrieves general information regarding the replication of a specific MySQL server.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_general_replication_information" "test_managed_my_sql_database_general_replication_information" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `apply_status_summary` - A summary of the current status of apply operations.
* `binary_log_format` - The binary logging format used by this server.
* `binary_logging` - The status of binary logging on this server.
* `executed_gtid_set` - The set of global transaction identifiers for transactions that have been executed on this source server.
* `fetch_status_summary` - A summary of the current status of fetch operations.
* `gtid_mode` - The Global Transaction Identifier (GTID) mode of this server.
* `high_availability_member_state` - The state of this server as a group replication member.
* `host_name` - This server's host name.
* `inbound_replications_count` - The number of sources this server is replicating from.
* `instance_type` - The type of the instance for example, Source, Replica, Primary Group Member, and Secondary Group Member. If the instance is replicating from one or more sources and has one or more replicas, which means, it belongs to a replication chain, the instance type can be Replica/Source.
* `is_high_availability_enabled` - Specifies if high availability is enabled on this server.
* `outbound_replications_count` - The number of replicas replicating from this server.
* `port` - The number of the port on which the server listens for TCP/IP connections.
* `read_only` - If the value is ON, the instance is configured as read_only. If the value is SUPER, the instance is configured as super_read_only. If the value is OFF, the instance is neither read_only nor super_read_only.
* `seconds_behind_source_max` - The number of seconds the replica is behind the source. When multiple sources are involved, this is the maximum value across all sources.
* `server_id` - This server's ID.
* `server_uuid` - This server's Universally Unique Identifier (UUID).

