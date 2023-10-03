---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_alert_policy_association"
sidebar_current: "docs-oci-resource-data_safe-target_alert_policy_association"
description: |-
  Provides the Target Alert Policy Association resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_target_alert_policy_association
This resource provides the Target Alert Policy Association resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new target-alert policy association to track a alert policy applied on target.


## Example Usage

```hcl
resource "oci_data_safe_target_alert_policy_association" "test_target_alert_policy_association" {
	#Required
	compartment_id = var.compartment_id
	is_enabled = var.target_alert_policy_association_is_enabled
	policy_id = oci_identity_policy.test_policy.id
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.target_alert_policy_association_description
	display_name = var.target_alert_policy_association_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the target-alert policy association is created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Describes the target-alert policy association.
* `display_name` - (Optional) (Updatable) The display name of the target-alert policy association.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_enabled` - (Required) (Updatable) Indicates if the target-alert policy association is enabled or disabled by user.
* `policy_id` - (Required) The OCID of the alert policy.
* `target_id` - (Required) The OCID of the target.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Describes the target-alert policy association.
* `display_name` - The display name of the target-alert policy association.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the target-alert policy association.
* `is_enabled` - Indicates if the target-alert policy association is enabled or disabled by user.
* `policy_id` - The OCID of the alert policy.
* `state` - The current state of the target-alert policy association.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target on which alert policy is to be applied.
* `time_created` - Creation date and time of the alert policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert policy was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target Alert Policy Association
	* `update` - (Defaults to 20 minutes), when updating the Target Alert Policy Association
	* `delete` - (Defaults to 20 minutes), when destroying the Target Alert Policy Association


## Import

TargetAlertPolicyAssociations can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_target_alert_policy_association.test_target_alert_policy_association "id"
```

