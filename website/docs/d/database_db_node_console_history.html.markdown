---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_history"
sidebar_current: "docs-oci-datasource-database-db_node_console_history"
description: |-
  Provides details about a specific Db Node Console History in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node_console_history
This data source provides details about a specific Db Node Console History resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified database node console history.


## Example Usage

```hcl
data "oci_database_db_node_console_history" "test_db_node_console_history" {
	#Required
	console_history_id = oci_core_console_history.test_console_history.id
	db_node_id = oci_database_db_node.test_db_node.id
}
```

## Argument Reference

The following arguments are supported:

* `console_history_id` - (Required) The OCID of the console history.
* `db_node_id` - (Required) The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

