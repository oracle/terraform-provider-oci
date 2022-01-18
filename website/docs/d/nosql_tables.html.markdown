---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_tables"
sidebar_current: "docs-oci-datasource-nosql-tables"
description: |-
  Provides the list of Tables in Oracle Cloud Infrastructure NoSQL Database service
---

# Data Source: oci_nosql_tables
This data source provides the list of Tables in Oracle Cloud Infrastructure NoSQL Database service.

Get a list of tables in a compartment.

## Example Usage

```hcl
data "oci_nosql_tables" "test_tables" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.table_name
	state = var.table_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of a table's compartment.
* `name` - (Optional) A shell-globbing-style (*?[]) filter for names.
* `state` - (Optional) Filter list by the lifecycle state of the item.


## Attributes Reference

The following attributes are exported:

* `table_collection` - The list of table_collection.

### Table Reference

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

