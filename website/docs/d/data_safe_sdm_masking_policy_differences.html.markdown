---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sdm_masking_policy_differences"
sidebar_current: "docs-oci-datasource-data_safe-sdm_masking_policy_differences"
description: |-
  Provides the list of Sdm Masking Policy Differences in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sdm_masking_policy_differences
This data source provides the list of Sdm Masking Policy Differences in Oracle Cloud Infrastructure Data Safe service.

Gets a list of SDM and masking policy difference resources based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_sdm_masking_policy_differences" "test_sdm_masking_policy_differences" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.sdm_masking_policy_difference_compartment_id_in_subtree
	difference_access_level = var.sdm_masking_policy_difference_difference_access_level
	display_name = var.sdm_masking_policy_difference_display_name
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
	state = var.sdm_masking_policy_difference_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `difference_access_level` - (Optional) Valid value is ACCESSIBLE. Default is ACCESSIBLE. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `masking_policy_id` - (Optional) A filter to return only the resources that match the specified masking policy OCID.
* `sensitive_data_model_id` - (Optional) A filter to return only the resources that match the specified sensitive data model OCID.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle states.


## Attributes Reference

The following attributes are exported:

* `sdm_masking_policy_difference_collection` - The list of sdm_masking_policy_difference_collection.

### SdmMaskingPolicyDifference Reference

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

