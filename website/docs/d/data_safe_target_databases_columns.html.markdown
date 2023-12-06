---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_databases_columns"
sidebar_current: "docs-oci-datasource-data_safe-target_databases_columns"
description: |-
  Provides the list of Target Databases Columns in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_databases_columns
This data source provides the list of Target Databases Columns in Oracle Cloud Infrastructure Data Safe service.

Returns a list of column metadata objects.


## Example Usage

```hcl
data "oci_data_safe_target_databases_columns" "test_target_databases_columns" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id

	#Optional
	column_name = var.target_databases_column_column_name
	column_name_contains = var.target_databases_column_column_name_contains
	datatype = var.target_databases_column_datatype
	schema_name = var.target_databases_column_schema_name
	schema_name_contains = var.target_databases_column_schema_name_contains
	table_name = oci_nosql_table.test_table.name
	table_name_contains = var.target_databases_column_table_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `column_name_contains` - (Optional) A filter to return only items if column name contains a specific string.
* `datatype` - (Optional) A filter to return only items related to specific datatype.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `schema_name_contains` - (Optional) A filter to return only items if schema name contains a specific string.
* `table_name` - (Optional) A filter to return only items related to specific table name.
* `table_name_contains` - (Optional) A filter to return only items if table name contains a specific string.
* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `columns` - The list of columns.

### TargetDatabasesColumn Reference

The following attributes are exported:

* `character_length` - Character length.
* `column_name` - Name of the column.
* `data_type` - Data type of the column.
* `length` - Length of the data represented by the column.
* `precision` - Precision of the column.
* `scale` - Scale of the column.
* `schema_name` - Name of the schema.
* `table_name` - Name of the table.

