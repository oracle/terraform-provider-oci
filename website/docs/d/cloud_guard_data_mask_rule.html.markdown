---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_mask_rule"
sidebar_current: "docs-oci-datasource-cloud_guard-data_mask_rule"
description: |-
  Provides details about a specific Data Mask Rule in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_data_mask_rule
This data source provides details about a specific Data Mask Rule resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a DataMaskRule resource, identified by dataMaskRuleId.

## Example Usage

```hcl
data "oci_cloud_guard_data_mask_rule" "test_data_mask_rule" {
	#Required
	data_mask_rule_id = oci_cloud_guard_data_mask_rule.test_data_mask_rule.id
}
```

## Argument Reference

The following arguments are supported:

* `data_mask_rule_id` - (Required) OCID of the data mask rule


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID where the resource is created
* `data_mask_categories` - List of data mask rule categories
* `data_mask_rule_status` - The current status of the data mask rule
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The data mask rule description
* `display_name` - Data mask rule display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `iam_group_id` - IAM Group ID associated with the data mask rule
* `id` - Unique identifier that can't be changed after creation
* `lifecyle_details` - Additional details on the substate of the lifecycle state [DEPRECATE]
* `state` - The current lifecycle state of the data mask rule
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_selected` - Specification of how targets are to be selected (select ALL, or select by TargetResourceType or TargetId).
	* `kind` - Kind of target selection to be used
	* `values` - Types of targets
* `time_created` - The date and time the target was created. Format defined by RFC3339.
* `time_updated` - The date and time the target was updated. Format defined by RFC3339.

