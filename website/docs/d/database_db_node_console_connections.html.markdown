---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_connections"
sidebar_current: "docs-oci-datasource-database-db_node_console_connections"
description: |-
  Provides the list of Db Node Console Connections in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node_console_connections
This data source provides the list of Db Node Console Connections in Oracle Cloud Infrastructure Database service.

Lists the console connections for the specified Db node.


## Example Usage

```hcl
data "oci_database_db_node_console_connections" "test_db_node_console_connections" {
	#Required
	db_node_id = "${oci_database_db_node.test_db_node.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `console_connections` - The list of console_connections.

### DbNodeConsoleConnection Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment to contain the console connection.
* `connection_string` - The SSH connection string for the console connection.
* `db_node_id` - The OCID of the database node.
* `fingerprint` - The SSH public key fingerprint for the console connection.
* `id` - The OCID of the console connection.
* `state` - The current state of the console connection.

