---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_sensitive_objects"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_model_sensitive_objects"
description: |-
  Provides the list of Sensitive Data Model Sensitive Objects in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_model_sensitive_objects
This data source provides the list of Sensitive Data Model Sensitive Objects in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive objects present in the specified sensitive data model based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_model_sensitive_objects" "test_sensitive_data_model_sensitive_objects" {
	#Required
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	object = var.sensitive_data_model_sensitive_object_object
	object_type = var.sensitive_data_model_sensitive_object_object_type
	schema_name = var.sensitive_data_model_sensitive_object_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `object_type` - (Optional) A filter to return only items related to a specific object type.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.


## Attributes Reference

The following attributes are exported:

* `sensitive_object_collection` - The list of sensitive_object_collection.

### SensitiveDataModelSensitiveObject Reference

The following attributes are exported:

* `items` - An array of sensitive object summary objects.
	* `object` - The database object that contains the sensitive column.
	* `object_type` - The type of the database object that contains the sensitive column.
	* `schema_name` - The database schema that contains the sensitive column.

