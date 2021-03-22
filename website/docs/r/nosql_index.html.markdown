---
subcategory: "NoSQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_nosql_index"
sidebar_current: "docs-oci-resource-nosql-index"
description: |-
  Provides the Index resource in Oracle Cloud Infrastructure NoSQL Database service
---

# oci_nosql_index
This resource provides the Index resource in Oracle Cloud Infrastructure NoSQL Database service.

Create a new index on the table identified by tableNameOrId.

## Example Usage

```hcl
resource "oci_nosql_index" "test_index" {
	#Required
	keys {
		#Required
		column_name = var.index_keys_column_name

		#Optional
		json_field_type = var.index_keys_json_field_type
		json_path = var.index_keys_json_path
	}
	name = var.index_name
	table_name_or_id = oci_nosql_table_name_or.test_table_name_or.id

	#Optional
	compartment_id = var.compartment_id
	is_if_not_exists = var.index_is_if_not_exists
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the table's compartment.  Required if the tableNameOrId path parameter is a table name. Optional if tableNameOrId is an OCID.  If tableNameOrId is an OCID, and compartmentId is supplied, the latter must match the identified table's compartmentId. 
* `is_if_not_exists` - (Optional) If true, the operation completes successfully even when the index exists.  Otherwise, an attempt to create an index that already exists will return an error. 
* `keys` - (Required) A set of keys for a secondary index.
	* `column_name` - (Required) The name of a column to be included as an index key.
	* `json_field_type` - (Optional) If the specified column is of type JSON, jsonFieldType contains the type of the field indicated by jsonPath. 
	* `json_path` - (Optional) If the specified column is of type JSON, jsonPath contains a dotted path indicating the field within the JSON object that will be the index key. 
* `name` - (Required) Index name.
* `table_name_or_id` - (Required) A table name within the compartment, or a table OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Index
	* `update` - (Defaults to 20 minutes), when updating the Index
	* `delete` - (Defaults to 20 minutes), when destroying the Index


## Import

Indexes can be imported using the `id`, e.g.

```
$ terraform import oci_nosql_index.test_index "tables/{tableNameOrId}/indexes/{indexName}" 
```

