---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_databases_tables"
sidebar_current: "docs-oci-datasource-data_safe-target_databases_tables"
description: |-
  Provides the list of Target Databases Tables in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_databases_tables
This data source provides the list of Target Databases Tables in Oracle Cloud Infrastructure Data Safe service.

Returns a list of table metadata objects.


## Example Usage

```hcl
data "oci_data_safe_target_databases_tables" "test_target_databases_tables" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id

	#Optional
	schema_name = var.target_databases_table_schema_name
	schema_name_contains = var.target_databases_table_schema_name_contains
	table_name = oci_nosql_table.test_table.name
	table_name_contains = var.target_databases_table_table_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `schema_name_contains` - (Optional) A filter to return only items if schema name contains a specific string.
* `table_name` - (Optional) A filter to return only items related to specific table name.
* `table_name_contains` - (Optional) A filter to return only items if table name contains a specific string.
* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `tables` - The list of tables.

### TargetDatabasesTable Reference

The following attributes are exported:

* `schema_name` - Name of the schema.
* `table_name` - Name of the table.

