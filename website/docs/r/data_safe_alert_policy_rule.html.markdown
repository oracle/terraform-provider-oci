---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_policy_rule"
sidebar_current: "docs-oci-resource-data_safe-alert_policy_rule"
description: |-
	Provides the Alert Policy Rule resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_alert_policy_rule
This resource provides the Alert Policy Rule resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new rule for the alert policy.


## Example Usage

```hcl
resource "oci_data_safe_alert_policy_rule" "test_alert_policy_rule" {
	#Required
	alert_policy_id = oci_data_safe_alert_policy.test_alert_policy.id
	expression = var.alert_policy_rule_expression

	#Optional
	description = var.alert_policy_rule_description
	display_name = var.alert_policy_rule_display_name
}
```

## Argument Reference

The following arguments are supported:

* `alert_policy_id` - (Required) The OCID of the alert policy.
* `description` - (Optional) (Updatable) Describes the alert policy rule.
* `display_name` - (Optional) (Updatable) The display name of the alert policy rule.
* `expression` - (Required) (Updatable) The conditional expression of the alert policy rule which evaluates to boolean value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - Describes the alert policy rule.
* `display_name` - The display name of the alert policy rule.
* `expression` - The conditional expression of the alert policy rule which evaluates to boolean value.
* `key` - The unique key of the alert policy rule.
* `state` - The current state of the alert policy rule.
* `time_created` - Creation date and time of the alert policy rule, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Alert Policy Rule
* `update` - (Defaults to 20 minutes), when updating the Alert Policy Rule
* `delete` - (Defaults to 20 minutes), when destroying the Alert Policy Rule


## Import

AlertPolicyRules can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_alert_policy_rule.test_alert_policy_rule "alertPolicies/{alertPolicyId}/rules/{ruleKey}" 
```
