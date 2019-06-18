---
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
	budget_id = "${oci_budget_budget.test_budget.id}"
	recipients = "${var.alert_rule_recipients}"
	threshold = "${var.alert_rule_threshold}"
	threshold_type = "${var.alert_rule_threshold_type}"
	type = "${var.alert_rule_type}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = "${var.alert_rule_description}"
	display_name = "${var.alert_rule_display_name}"
	freeform_tags = {"bar-key"= "value"}
	message = "${var.alert_rule_message}"
}
```

## Argument Reference

The following arguments are supported:

* `budget_id` - (Required) The unique Budget OCID
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The description of the alert rule.
* `display_name` - (Optional) (Updatable) The name of the alert rule.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `message` - (Optional) (Updatable) The message to be sent to the recipients when alert rule is triggered.
* `recipients` - (Required) (Updatable) The audience that will received the alert when it triggers.
* `threshold` - (Required) (Updatable) The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000. 
* `threshold_type` - (Required) (Updatable) The type of threshold.
* `type` - (Required) (Updatable) ACTUAL means the alert will trigger based on actual usage. FORECAST means the alert will trigger based on predicted usage. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `budget_id` - The OCID of the budget
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the alert rule.
* `display_name` - The name of the alert rule.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the alert rule
* `message` - Custom message that will be sent when alert is triggered
* `recipients` - The audience that will received the alert when it triggers.
* `state` - The current state of the alert rule.
* `threshold` - The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000. 
* `threshold_type` - The type of threshold.
* `time_created` - Time when budget was created
* `time_updated` - Time when budget was updated
* `type` - ACTUAL means the alert will trigger based on actual usage. FORECAST means the alert will trigger based on predicted usage. 
* `version` - Version of the alert rule. Starts from 1 and increments by 1.

## Import

Import is not supported for this resource.

