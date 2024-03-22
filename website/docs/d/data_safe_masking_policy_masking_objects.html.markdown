---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_masking_objects"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_masking_objects"
description: |-
  Provides the list of Masking Policy Masking Objects in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_masking_objects
This data source provides the list of Masking Policy Masking Objects in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking objects present in the specified masking policy and based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_masking_policy_masking_objects" "test_masking_policy_masking_objects" {
	#Required
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

	#Optional
	object = var.masking_policy_masking_object_object
	object_type = var.masking_policy_masking_object_object_type
	schema_name = var.masking_policy_masking_object_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `masking_policy_id` - (Required) The OCID of the masking policy.
* `object` - (Optional) A filter to return only items related to a specific object name.
* `object_type` - (Optional) A filter to return only items related to a specific object type.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.


## Attributes Reference

The following attributes are exported:

* `masking_object_collection` - The list of masking_object_collection.

### MaskingPolicyMaskingObject Reference

The following attributes are exported:

* `items` - An array of masking object summary objects.
	* `object` - The database object that contains the masking column.
	* `object_type` - The type of the database object that contains the masking column.
	* `schema_name` - The database schema that contains the masking column.

