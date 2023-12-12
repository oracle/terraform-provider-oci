---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sdm_masking_policy_difference"
sidebar_current: "docs-oci-datasource-data_safe-sdm_masking_policy_difference"
description: |-
  Provides details about a specific Sdm Masking Policy Difference in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sdm_masking_policy_difference
This data source provides details about a specific Sdm Masking Policy Difference resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified SDM Masking policy difference.

## Example Usage

```hcl
data "oci_data_safe_sdm_masking_policy_difference" "test_sdm_masking_policy_difference" {
	#Required
	sdm_masking_policy_difference_id = oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference.id
}
```

## Argument Reference

The following arguments are supported:

* `sdm_masking_policy_difference_id` - (Required) The OCID of the SDM masking policy difference.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the Sensitive data model and masking policy difference resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `difference_type` - The type of the SDM masking policy difference. It defines the difference scope. NEW identifies new sensitive columns in the sensitive data model that are not in the masking policy. DELETED identifies columns that are present in the masking policy but have been deleted from the sensitive data model. MODIFIED identifies columns that are present in the sensitive data model as well as the masking policy but some of their attributes have been modified. ALL covers all the above three scenarios and reports new, deleted and modified columns. 
* `display_name` - The display name of the SDM masking policy difference.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Sensitive data model and masking policy difference resource.
* `masking_policy_id` - The OCID of the masking policy associated with the SDM masking policy difference.
* `sensitive_data_model_id` - The OCID of the sensitive data model associated with the SDM masking policy difference.
* `state` - The current state of the SDM masking policy difference.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the SDM masking policy difference was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_creation_started` - The date and time the SDM masking policy difference creation started, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

