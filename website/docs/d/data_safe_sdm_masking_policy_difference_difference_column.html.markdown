---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sdm_masking_policy_difference_difference_column"
sidebar_current: "docs-oci-datasource-data_safe-sdm_masking_policy_difference_difference_column"
description: |-
  Provides details about a specific Sdm Masking Policy Difference Difference Column in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sdm_masking_policy_difference_difference_column
This data source provides details about a specific Sdm Masking Policy Difference Difference Column resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified SDM Masking policy difference column.

## Example Usage

```hcl
data "oci_data_safe_sdm_masking_policy_difference_difference_column" "test_sdm_masking_policy_difference_difference_column" {
	#Required
	difference_column_key = var.sdm_masking_policy_difference_difference_column_difference_column_key
	sdm_masking_policy_difference_id = oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id
}
```

## Argument Reference

The following arguments are supported:

* `difference_column_key` - (Required) The unique key that identifies the difference column.
* `sdm_masking_policy_difference_id` - (Required) The OCID of the SDM masking policy difference.


## Attributes Reference

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

