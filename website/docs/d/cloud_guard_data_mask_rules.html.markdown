---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_mask_rules"
sidebar_current: "docs-oci-datasource-cloud_guard-data_mask_rules"
description: |-
  Provides the list of Data Mask Rules in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_data_mask_rules
This data source provides the list of Data Mask Rules in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all DataMaskRule resources in the specified compartmentId (OCID) and its subcompartments.


## Example Usage

```hcl
data "oci_cloud_guard_data_mask_rules" "test_data_mask_rules" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.data_mask_rule_access_level
	data_mask_rule_status = var.data_mask_rule_data_mask_rule_status
	display_name = var.data_mask_rule_display_name
	iam_group_id = oci_identity_group.test_group.id
	state = var.data_mask_rule_state
	target_id = oci_cloud_guard_target.test_target.id
	target_type = var.data_mask_rule_target_type
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `data_mask_rule_status` - (Optional) The status of the data mask rule
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `iam_group_id` - (Optional) OCID of the IAM group
* `state` - (Optional) The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
* `target_id` - (Optional) OCID of the target
* `target_type` - (Optional) Type of target


## Attributes Reference

The following attributes are exported:

* `data_mask_rule_collection` - The list of data_mask_rule_collection.

### DataMaskRule Reference

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

