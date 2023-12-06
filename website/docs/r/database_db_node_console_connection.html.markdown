---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_connection"
sidebar_current: "docs-oci-resource-database-db_node_console_connection"
description: |-
  Provides the Db Node Console Connection resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_node_console_connection
This resource provides the Db Node Console Connection resource in Oracle Cloud Infrastructure Database service.

Creates a new console connection to the specified database node.
After the console connection has been created and is available,
you connect to the console using SSH.


## Example Usage

```hcl
resource "oci_database_db_node_console_connection" "test_db_node_console_connection" {
	#Required
	db_node_id = oci_database_db_node.test_db_node.id
	public_key = var.db_node_console_connection_public_key

	#Optional
	defined_tags = var.db_node_console_connection_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `public_key` - (Required) The SSH public key used to authenticate the console connection.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Node Console Connection
	* `update` - (Defaults to 20 minutes), when updating the Db Node Console Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Db Node Console Connection


## Import

DbNodeConsoleConnections can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_node_console_connection.test_db_node_console_connection "dbNodes/{dbNodeId}/consoleConnections/{consoleConnectionId}" 
```

