---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_histories"
sidebar_current: "docs-oci-datasource-database-db_node_console_histories"
description: |-
  Provides the list of Db Node Console Histories in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node_console_histories
This data source provides the list of Db Node Console Histories in Oracle Cloud Infrastructure Database service.

Lists the console histories for the specified database node.


## Example Usage

```hcl
data "oci_database_db_node_console_histories" "test_db_node_console_histories" {
	#Required
	db_node_id = oci_database_db_node.test_db_node.id

	#Optional
	display_name = var.db_node_console_history_display_name
	state = var.db_node_console_history_state
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `console_history_collection` - The list of console_history_collection.

### DbNodeConsoleHistory Reference

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

