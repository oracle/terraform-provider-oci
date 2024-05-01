---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_masking_schemas"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_masking_schemas"
description: |-
  Provides the list of Masking Policy Masking Schemas in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_masking_schemas
This data source provides the list of Masking Policy Masking Schemas in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masking schemas present in the specified masking policy and based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_masking_policy_masking_schemas" "test_masking_policy_masking_schemas" {
	#Required
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

	#Optional
	schema_name = var.masking_policy_masking_schema_schema_name
}
```

## Argument Reference

The following arguments are supported:

* `masking_policy_id` - (Required) The OCID of the masking policy.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.


## Attributes Reference

The following attributes are exported:

* `masking_schema_collection` - The list of masking_schema_collection.

### MaskingPolicyMaskingSchema Reference

The following attributes are exported:

* `items` - An array of masking schema summary objects.
	* `schema_name` - The database schema that contains the masking column.

