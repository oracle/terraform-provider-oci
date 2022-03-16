---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_databases_schemas"
sidebar_current: "docs-oci-datasource-data_safe-target_databases_schemas"
description: |-
  Provides the list of Target Databases Schemas in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_databases_schemas
This data source provides the list of Target Databases Schemas in Oracle Cloud Infrastructure Data Safe service.

Returns list of schema.


## Example Usage

```hcl
data "oci_data_safe_target_databases_schemas" "test_target_databases_schemas" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id

	#Optional
	is_oracle_maintained = var.target_databases_schema_is_oracle_maintained
	schema_name = var.target_databases_schema_schema_name
	schema_name_contains = var.target_databases_schema_schema_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `is_oracle_maintained` - (Optional) A filter to return only items related to specific type of schema.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `schema_name_contains` - (Optional) A filter to return only items if schema name contains a specific string.
* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `schemas` - The list of schemas.

### TargetDatabasesSchema Reference

The following attributes are exported:

* `is_oracle_maintained` - Indicates if the schema is oracle supplied.
* `schema_name` - Name of the schema.

