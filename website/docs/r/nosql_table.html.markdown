---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_table"
sidebar_current: "docs-oci-resource-nosql-table"
description: |-
  Provides the Table resource in Oracle Cloud Infrastructure NoSQL Database service
---

# oci_nosql_table
This resource provides the Table resource in Oracle Cloud Infrastructure NoSQL Database service.

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

		#Optional
		capacity_mode = var.table_table_limits_capacity_mode
	}

	#Optional
	defined_tags = var.table_defined_tags
	freeform_tags = {"bar-key"= "value"}
	is_auto_reclaimable = var.table_is_auto_reclaimable
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `ddl_statement` - (Required) (Updatable) Complete CREATE TABLE DDL statement. When update ddl_statement, it should be ALTER TABLE DDL statement.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace": {"bar-key": "value"}}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_auto_reclaimable` - (Optional) True if table can be reclaimed after an idle period.
* `name` - (Required) Table name.
* `table_limits` - (Required) (Updatable) Throughput and storage limits configuration of a table.
	* `capacity_mode` - (Optional) (Updatable) The capacity mode of the table.  If capacityMode = ON_DEMAND, maxReadUnits and maxWriteUnits are not used, and both will have the value of zero. 
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
* `is_auto_reclaimable` - True if this table can be reclaimed after an idle period.
* `lifecycle_details` - A message describing the current state in more detail. 
* `name` - Human-friendly table name, immutable.
* `schema` - The table schema information as a JSON object.
	* `columns` - The columns of a table.
		* `default_value` - The column default value.
		* `is_nullable` - The column nullable flag.
		* `name` - The column name.
		* `type` - The column type.
	* `primary_key` - A list of column names that make up a key.
	* `shard_key` - A list of column names that make up a key.
	* `ttl` - The default Time-to-Live for the table, in days.
* `state` - The state of a table.
* `system_tags` - Read-only system tag. These predefined keys are scoped to namespaces.  At present the only supported namespace is `"orcl-cloud"`; and the only key in that namespace is `"free-tier-retained"`. Example: `{"orcl-cloud"": {"free-tier-retained": "true"}}` 
* `table_limits` - Throughput and storage limits configuration of a table.
	* `capacity_mode` - The capacity mode of the table.  If capacityMode = ON_DEMAND, maxReadUnits and maxWriteUnits are not used, and both will have the value of zero. 
	* `max_read_units` - Maximum sustained read throughput limit for the table.
	* `max_storage_in_gbs` - Maximum size of storage used by the table.
	* `max_write_units` - Maximum sustained write throughput limit for the table.
* `time_created` - The time the the table was created. An RFC3339 formatted datetime string. 
* `time_of_expiration` - If lifecycleState is INACTIVE, indicates when this table will be automatically removed. An RFC3339 formatted datetime string. 
* `time_updated` - The time the the table's metadata was last updated. An RFC3339 formatted datetime string. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Table
	* `update` - (Defaults to 20 minutes), when updating the Table
	* `delete` - (Defaults to 20 minutes), when destroying the Table


## Import

Tables can be imported using the `id`, e.g.

```
$ terraform import oci_nosql_table.test_table "id"
```

