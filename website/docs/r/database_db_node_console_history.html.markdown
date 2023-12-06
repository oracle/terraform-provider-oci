---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_history"
sidebar_current: "docs-oci-resource-database-db_node_console_history"
description: |-
  Provides the Db Node Console History resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_node_console_history
This resource provides the Db Node Console History resource in Oracle Cloud Infrastructure Database service.

Captures the most recent serial console data (up to a megabyte) for the specified database node.


## Example Usage

```hcl
resource "oci_database_db_node_console_history" "test_db_node_console_history" {
	#Required
	db_node_id = oci_database_db_node.test_db_node.id
	display_name = var.db_node_console_history_display_name

	#Optional
	defined_tags = var.db_node_console_history_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the console history. The name does not need to be unique. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the console history.
* `db_node_id` - The OCID of the database node.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the console history. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the console history.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the console history.
* `time_created` - The date and time the console history was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Node Console History
	* `update` - (Defaults to 20 minutes), when updating the Db Node Console History
	* `delete` - (Defaults to 20 minutes), when destroying the Db Node Console History


## Import

DbNodeConsoleHistories can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_node_console_history.test_db_node_console_history "dbNodes/{dbNodeId}/consoleHistories/{consoleHistoryId}" 
```

