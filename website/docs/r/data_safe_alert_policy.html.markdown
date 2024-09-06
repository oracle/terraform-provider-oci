---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policy"
sidebar_current: "docs-oci-resource-data_safe-alert_policy"
description: |-
  Provides the Alert Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_alert_policy
This resource provides the Alert Policy resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new user-defined alert policy.


## Example Usage

```hcl
resource "oci_data_safe_alert_policy" "test_alert_policy" {
	#Required
	alert_policy_type = var.alert_policy_alert_policy_type
	compartment_id = var.compartment_id
	severity = var.alert_policy_severity

	#Optional
	alert_policy_rule_details {
		#Required
		expression = var.alert_policy_alert_policy_rule_details_expression

		#Optional
		description = var.alert_policy_alert_policy_rule_details_description
		display_name = var.alert_policy_alert_policy_rule_details_display_name
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.alert_policy_description
	display_name = var.alert_policy_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_rule_details` - (Optional) The details of the alert policy rule.
	* `description` - (Optional) Describes the alert policy rule.
	* `display_name` - (Optional) The display name of the alert policy rule.
	* `expression` - (Required) The conditional expression of the alert policy rule which evaluates to boolean value.
* `alert_policy_type` - (Required) Indicates the Data Safe feature the alert policy belongs to
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the alert policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) The description of the alert policy.
* `display_name` - (Optional) (Updatable) The display name of the alert policy. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `severity` - (Required) (Updatable) Severity level of the alert raised by this policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alert_policy_type` - Indicates the Data Safe feature to which the alert policy belongs.
* `compartment_id` - The OCID of the compartment that contains the alert policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
* `description` - The description of the alert policy.
* `display_name` - The display name of the alert policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
* `id` - The OCID of the alert policy.
* `is_user_defined` - Indicates if the alert policy is user-defined (true) or pre-defined (false).
* `lifecycle_details` - Details about the current state of the alert policy.
* `severity` - Severity level of the alert raised by this policy.
* `state` - The current state of the alert.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Creation date and time of the alert policy, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - Last date and time the alert policy was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Alert Policy
* `update` - (Defaults to 20 minutes), when updating the Alert Policy
* `delete` - (Defaults to 20 minutes), when destroying the Alert Policy


## Import

AlertPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_alert_policy.test_alert_policy "id"
```
