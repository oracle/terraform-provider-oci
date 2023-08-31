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

Lists the console connections for the specified database node.


## Example Usage

```hcl
data "oci_database_db_node_console_connections" "test_db_node_console_connections" {
	#Required
	db_node_id = oci_database_db_node.test_db_node.id
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
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `fingerprint` - The SSH public key fingerprint for the console connection.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the console connection.
* `lifecycle_details` - Information about the current lifecycle state.
* `service_host_key_fingerprint` - The SSH public key's fingerprint for the console connection service host.
* `state` - The current state of the console connection.

