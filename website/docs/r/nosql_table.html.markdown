---
subcategory: "Nosql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_table"
sidebar_current: "docs-oci-resource-nosql-table"
description: |-
  Provides the Table resource in Oracle Cloud Infrastructure Nosql service
---

# oci_nosql_table
This resource provides the Table resource in Oracle Cloud Infrastructure Nosql service.

Create a new table.

## Example Usage

```hcl
resource "oci_nosql_table" "test_table" {
	#Required
	compartment_id = var.compartment_id
	ddl_statement = var.table_ddl_statement
	name = var.table_name
	table_limits {
		#Required
		max_read_units = var.table_table_limits_max_read_units
		max_storage_in_gbs = var.table_table_limits_max_storage_in_gbs
		max_write_units = var.table_table_limits_max_write_units
	}

	#Optional
	defined_tags = var.table_defined_tags
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `ddl_statement` - (Required) (Updatable) Complete CREATE TABLE DDL statement.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace": {"bar-key": "value"}}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) Table name.
* `table_limits` - (Required) (Updatable) 
	* `max_read_units` - (Required) (Updatable) Maximum sustained read throughput limit for the table.
	* `max_storage_in_gbs` - (Required) (Updatable) Maximum size of storage used by the table.
	* `max_write_units` - (Required) (Updatable) Maximum sustained write throughput limit for the table.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `ddl_statement` - A DDL statement representing the schema.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace": {"bar-key": "value"}}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable.
* `lifecycle_details` - A message describing the current state in more detail. 
* `name` - Human-friendly table name, immutable.
* `schema` - 
	* `columns` - The columns of a table.
		* `default_value` - The column default value.
		* `is_nullable` - The column nullable flag.
		* `name` - The column name.
		* `type` - The column type.
	* `primary_key` - A list of column names that make up a key.
	* `shard_key` - A list of column names that make up a key.
	* `ttl` - The default Time-to-Live for the table, in days.
* `state` - The state of a table.
* `table_limits` - 
	* `max_read_units` - Maximum sustained read throughput limit for the table.
	* `max_storage_in_gbs` - Maximum size of storage used by the table.
	* `max_write_units` - Maximum sustained write throughput limit for the table.
* `time_created` - The time the the table was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time the the table's metadata was last updated. An RFC3339 formatted datetime string. 

## Import

Tables can be imported using the `id`, e.g.

```
$ terraform import oci_nosql_table.test_table "id"
```

