---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_node_console_history_content"
sidebar_current: "docs-oci-datasource-database-db_node_console_history_content"
description: |-
  Provides details about a specific Db Node Console History Content in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_node_console_history_content
This data source provides details about a specific Db Node Console History Content resource in Oracle Cloud Infrastructure Database service.

Retrieves the specified database node console history contents upto a megabyte.


## Example Usage

```hcl
data "oci_database_db_node_console_history_content" "test_db_node_console_history_content" {
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


