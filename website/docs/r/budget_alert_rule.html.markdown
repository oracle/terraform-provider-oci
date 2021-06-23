---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_alert_rule"
sidebar_current: "docs-oci-resource-budget-alert_rule"
description: |-
  Provides the Alert Rule resource in Oracle Cloud Infrastructure Budget service
---

# oci_budget_alert_rule
This resource provides the Alert Rule resource in Oracle Cloud Infrastructure Budget service.

Creates a new Alert Rule.


## Example Usage

```hcl
resource "oci_budget_alert_rule" "test_alert_rule" {
	#Required
	budget_id = oci_budget_budget.test_budget.id
	threshold = var.alert_rule_threshold
	threshold_type = var.alert_rule_threshold_type
	type = var.alert_rule_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.alert_rule_description
	display_name = var.alert_rule_display_name
	freeform_tags = {"Department"= "Finance"}
	message = var.alert_rule_message
	recipients = var.alert_rule_recipients
}
```

## Argument Reference

The following arguments are supported:

* `budget_id` - (Required) The unique Budget OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the alert rule.
* `display_name` - (Optional) (Updatable) The name of the alert rule.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `message` - (Optional) (Updatable) The message to be sent to the recipients when alert rule is triggered.
* `recipients` - (Optional) (Updatable) The audience that will receive the alert when it triggers. An empty string is interpreted as null.
* `threshold` - (Required) (Updatable) The threshold for triggering the alert expressed as a whole number or decimal value. If thresholdType is ABSOLUTE, threshold can have at most 12 digits before the decimal point and up to 2 digits after the decimal point. If thresholdType is PERCENTAGE, the maximum value is 10000 and can have up to 2 digits after the decimal point. 
* `threshold_type` - (Required) (Updatable) The type of threshold.
* `type` - (Required) (Updatable) Type of alert. Valid values are ACTUAL (the alert will trigger based on actual usage) or FORECAST (the alert will trigger based on predicted usage). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `budget_id` - The OCID of the budget
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the alert rule.
* `display_name` - The name of the alert rule.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the alert rule
* `message` - Custom message that will be sent when alert is triggered
* `recipients` - Delimited list of email addresses to receive the alert when it triggers. Delimiter character can be comma, space, TAB, or semicolon. 
* `state` - The current state of the alert rule.
* `threshold` - The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000. 
* `threshold_type` - The type of threshold.
* `time_created` - Time when budget was created
* `time_updated` - Time when budget was updated
* `type` - The type of alert. Valid values are ACTUAL (the alert will trigger based on actual usage) or FORECAST (the alert will trigger based on predicted usage). 
* `version` - Version of the alert rule. Starts from 1 and increments by 1.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Alert Rule
	* `update` - (Defaults to 20 minutes), when updating the Alert Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Alert Rule


## Import

AlertRules can be imported using the `id`, e.g.

```
$ terraform import oci_budget_alert_rule.test_alert_rule "budgets/{budgetId}/alertRules/{alertRuleId}" 
```

