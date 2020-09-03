---
subcategory: "Nosql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_indexes"
sidebar_current: "docs-oci-datasource-nosql-indexes"
description: |-
  Provides the list of Indexes in Oracle Cloud Infrastructure Nosql service
---

# Data Source: oci_nosql_indexes
This data source provides the list of Indexes in Oracle Cloud Infrastructure Nosql service.

Get a list of indexes on a table.

## Example Usage

```hcl
data "oci_nosql_indexes" "test_indexes" {
	#Required
	table_name_or_id = oci_nosql_table_name_or.test_table_name_or.id

	#Optional
	compartment_id = var.compartment_id
	name = var.index_name
	state = var.index_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of a table's compartment. When a table is identified by name, the compartmentId is often needed to provide context for interpreting the name. 
* `name` - (Optional) A shell-globbing-style (*?[]) filter for names.
* `state` - (Optional) Filter list by the lifecycle state of the item.
* `table_name_or_id` - (Required) A table name within the compartment, or a table OCID.


## Attributes Reference

The following attributes are exported:

* `index_collection` - The list of index_collection.

### Index Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `keys` - A set of keys for a secondary index.
	* `column_name` - The name of a column to be included as an index key.
	* `json_field_type` - If the specified column is of type JSON, jsonFieldType contains the type of the field indicated by jsonPath. 
	* `json_path` - If the specified column is of type JSON, jsonPath contains a dotted path indicating the field within the JSON object that will be the index key. 
* `lifecycle_details` - A message describing the current state in more detail. 
* `name` - Index name.
* `state` - The state of an index.
* `table_id` - the OCID of the table to which this index belongs.
* `table_name` - The name of the table to which this index belongs.

