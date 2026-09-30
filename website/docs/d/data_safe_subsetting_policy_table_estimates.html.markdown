---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_subsetting_policy_table_estimates"
sidebar_current: "docs-oci-datasource-data_safe-subsetting_policy_table_estimates"
description: |-
  Provides the list of Subsetting Policy Table Estimates in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_subsetting_policy_table_estimates
This data source provides the list of Subsetting Policy Table Estimates in Oracle Cloud Infrastructure Data Safe service.

Gets table size estimates for the specified subsetting policy.

## Example Usage

```hcl
data "oci_data_safe_subsetting_policy_table_estimates" "test_subsetting_policy_table_estimates" {
	#Required
	subsetting_policy_id = oci_data_safe_subsetting_policy.test_subsetting_policy.id

	#Optional
	object = var.subsetting_policy_table_estimate_object
	schema_name = var.subsetting_policy_table_estimate_schema_name
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `object` - (Optional) A filter to return only items related to a specific object name.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `subsetting_policy_id` - (Required) The OCID of the subsetting policy.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `table_estimate_collection` - The list of table_estimate_collection.

### SubsettingPolicyTableEstimate Reference

The following attributes are exported:

* `items` - An array of table estimate objects.
	* `estimated_row_count` - The estimated number of rows in the table after subsetting.
	* `estimated_size_in_kbs` - The estimated size of the table in KBs after subsetting.
	* `initial_row_count` - The initial number of rows in the table.
	* `initial_size_in_kbs` - The initial size of the table in KBs.
	* `object_type` - The type of the database object.
	* `schema_name` - The name of the schema that contains the table.
	* `table_name` - The name of the table.
	* `target_id` - The OCID of the target database associated with the table estimate object.

