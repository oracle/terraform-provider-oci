---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_data_model_sensitive_types"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_data_model_sensitive_types"
description: |-
  Provides the list of Sensitive Data Model Sensitive Types in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_data_model_sensitive_types
This data source provides the list of Sensitive Data Model Sensitive Types in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive type Ids present in the specified sensitive data model.


## Example Usage

```hcl
data "oci_data_safe_sensitive_data_model_sensitive_types" "test_sensitive_data_model_sensitive_types" {
	#Required
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

	#Optional
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_data_model_id` - (Required) The OCID of the sensitive data model.
* `sensitive_type_id` - (Optional) A filter to return only items related to a specific sensitive type OCID.


## Attributes Reference

The following attributes are exported:

* `sensitive_data_model_sensitive_type_collection` - The list of sensitive_data_model_sensitive_type_collection.

### SensitiveDataModelSensitiveType Reference

The following attributes are exported:

* `items` - An array of sensitive types summary objects present in a sensitive data model.
	* `sensitive_data_model_sensitive_type_count` - The total number of sensitive columns linked to this specific sensitive type .
	* `sensitive_type_id` - The OCID of the sensitive type.

