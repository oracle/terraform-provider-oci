---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sdm_masking_policy_difference"
sidebar_current: "docs-oci-resource-data_safe-sdm_masking_policy_difference"
description: |-
  Provides the Sdm Masking Policy Difference resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sdm_masking_policy_difference
This resource provides the Sdm Masking Policy Difference resource in Oracle Cloud Infrastructure Data Safe service.

Creates SDM masking policy difference for the specified masking policy. It finds the difference between
masking columns of the masking policy and sensitive columns of the SDM. After performing this operation,
you can use ListDifferenceColumns to view the difference columns, PatchSdmMaskingPolicyDifferenceColumns
to specify the action you want perform on these columns, and then ApplySdmMaskingPolicyDifference to process the
difference columns and apply them to the masking policy.


## Example Usage

```hcl
resource "oci_data_safe_sdm_masking_policy_difference" "test_sdm_masking_policy_difference" {
	#Required
	compartment_id = var.compartment_id
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	difference_type = var.sdm_masking_policy_difference_difference_type
	display_name = var.sdm_masking_policy_difference_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the SDM masking policy difference resource should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `difference_type` - (Optional) The type of the SDM masking policy difference. It defines the difference scope. NEW identifies new sensitive columns in the sensitive data model that are not in the masking policy. DELETED identifies columns that are present in the masking policy but have been deleted from the sensitive data model. MODIFIED identifies columns that are present in the sensitive data model as well as the masking policy but some of their attributes have been modified. ALL covers all the above three scenarios and reports new, deleted and modified columns. 
* `display_name` - (Optional) (Updatable) A user-friendly name for the SDM masking policy difference. Does not have to be unique, and it is changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `masking_policy_id` - (Required) The OCID of the masking policy. Note that if the masking policy is not associated with an SDM, CreateSdmMaskingPolicyDifference operation won't be allowed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sdm Masking Policy Difference
	* `update` - (Defaults to 20 minutes), when updating the Sdm Masking Policy Difference
	* `delete` - (Defaults to 20 minutes), when destroying the Sdm Masking Policy Difference


## Import

SdmMaskingPolicyDifferences can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sdm_masking_policy_difference.test_sdm_masking_policy_difference "id"
```

