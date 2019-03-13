---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_alert_rules"
sidebar_current: "docs-oci-datasource-budget-alert_rules"
description: |-
  Provides the list of Alert Rules in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_alert_rules
This data source provides the list of Alert Rules in Oracle Cloud Infrastructure Budget service.

Returns a list of Alert Rules.


## Example Usage

```hcl
data "oci_budget_alert_rules" "test_alert_rules" {
	#Required
	budget_id = "${oci_budget_budget.test_budget.id}"

	#Optional
	display_name = "${var.alert_rule_display_name}"
	state = "${var.alert_rule_state}"
}
```

## Argument Reference

The following arguments are supported:

* `budget_id` - (Required) The unique Budget OCID
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) The current state of the resource to filter by.


## Attributes Reference

The following attributes are exported:

* `alert_rules` - The list of alert_rules.

### AlertRule Reference

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

