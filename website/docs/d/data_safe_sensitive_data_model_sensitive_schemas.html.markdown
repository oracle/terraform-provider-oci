---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_sensitive_schemas"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_model_sensitive_schemas"
description: |-
  Provides the list of Sensitive Data Model Sensitive Schemas in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_model_sensitive_schemas
This data source provides the list of Sensitive Data Model Sensitive Schemas in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive schemas present in the specified sensitive data model based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_model_sensitive_schemas" "test_sensitive_data_model_sensitive_schemas" {
	#Required
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	schema_name = var.sensitive_data_model_sensitive_schema_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


## Attributes Reference

The following attributes are exported:

* `sensitive_schema_collection` - The list of sensitive_schema_collection.

### SensitiveDataModelSensitiveSchema Reference

The following attributes are exported:

* `items` - An array of sensitive schema summary objects.
	* `schema_name` - The database schema that contains the sensitive column.

