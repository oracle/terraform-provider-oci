---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sdm_masking_policy_difference_difference_columns"
sidebar_current: "docs-oci-datasource-data_safe-sdm_masking_policy_difference_difference_columns"
description: |-
  Provides the list of Sdm Masking Policy Difference Difference Columns in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sdm_masking_policy_difference_difference_columns
This data source provides the list of Sdm Masking Policy Difference Difference Columns in Oracle Cloud Infrastructure Data Safe service.

Gets a list of columns of a SDM masking policy difference resource based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_sdm_masking_policy_difference_difference_columns" "test_sdm_masking_policy_difference_difference_columns" {
	#Required
	sdm_masking_policy_difference_id = oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id

	#Optional
	column_name = var.sdm_masking_policy_difference_difference_column_column_name
	difference_type = var.sdm_masking_policy_difference_difference_column_difference_type
	object = var.sdm_masking_policy_difference_difference_column_object
	planned_action = var.sdm_masking_policy_difference_difference_column_planned_action
	schema_name = var.sdm_masking_policy_difference_difference_column_schema_name
	sync_status = var.sdm_masking_policy_difference_difference_column_sync_status
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `difference_type` - (Optional) A filter to return only the SDM masking policy difference columns that match the specified difference type
* `object` - (Optional) A filter to return only items related to a specific object name.
* `planned_action` - (Optional) A filter to return only the SDM masking policy difference columns that match the specified planned action.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sdm_masking_policy_difference_id` - (Required) The OCID of the SDM masking policy difference.
* `sync_status` - (Optional) A filter to return the SDM masking policy difference columns based on the value of their syncStatus attribute.


## Attributes Reference

The following attributes are exported:

* `sdm_masking_policy_difference_column_collection` - The list of sdm_masking_policy_difference_column_collection.

### SdmMaskingPolicyDifferenceDifferenceColumn Reference

The following attributes are exported:

* `column_name` - The name of the difference column.
* `difference_type` - The type of the SDM masking policy difference column. It can be one of the following three types: NEW: A new sensitive column in the sensitive data model that is not in the masking policy. DELETED: A column that is present in the masking policy but has been deleted from the sensitive data model. MODIFIED: A column that is present in the masking policy as well as the sensitive data model but some of its attributes have been modified. 
* `key` - The unique key that identifies the SDM masking policy difference column.
* `masking_columnkey` - The unique key that identifies the masking column represented by the SDM masking policy difference column.
* `object` - The database object that contains the difference column.
* `planned_action` - Specifies how to process the difference column. It's set to SYNC by default. Use the PatchSdmMaskingPolicyDifferenceColumns operation to update this attribute. You can choose one of the following options: SYNC: To sync the difference column and update the masking policy to reflect the changes. NO_SYNC: To not sync the difference column so that it doesn't change the masking policy. After specifying the planned action, you can use the ApplySdmMaskingPolicyDifference operation to automatically process the difference columns. 
* `schema_name` - The database schema that contains the difference column.
* `sensitive_columnkey` - The unique key that identifies the sensitive column represented by the SDM masking policy difference column.
* `sensitive_type_id` - The OCID of the sensitive type associated with the difference column.
* `sync_status` - Indicates if the difference column has been processed. Use GetDifferenceColumn operation to  track whether the difference column has already been processed and applied to the masking policy. 
* `time_last_synced` - The date and time the SDM masking policy difference column was last synced, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

