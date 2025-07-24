---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_outbound_replications"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_outbound_replications"
description: |-
  Provides the list of Managed My Sql Database Outbound Replications in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_outbound_replications
This data source provides the list of Managed My Sql Database Outbound Replications in Oracle Cloud Infrastructure Database Management service.

Retrieves information pertaining to the outbound replications of a specific MySQL server.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_outbound_replications" "test_managed_my_sql_database_outbound_replications" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.


## Attributes Reference

The following attributes are exported:

* `managed_my_sql_database_outbound_replication_collection` - The list of managed_my_sql_database_outbound_replication_collection.

### ManagedMySqlDatabaseOutboundReplication Reference

The following attributes are exported:

* `items` - The list of ManagedMySqlDatabaseOutboundReplicationSummary records.
	* `replica_host` - The host name of the replica server, as specified on the replica with the --report-host option. This can differ from the machine name as configured in the operating system.
	* `replica_port` - The port on the replica server, as specified on the replica with the --report-port option. A zero in this column means that the replica port (--report-port) was not set.
	* `replica_server_id` - The server ID value of the replica.
	* `replica_uuid` - The Universally Unique Identifier (UUID) value of the replica server.
* `outbound_replications_count` - The number of outbound replications from the MySQL server.

