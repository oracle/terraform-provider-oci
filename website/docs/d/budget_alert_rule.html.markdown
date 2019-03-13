---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_alert_rule"
sidebar_current: "docs-oci-datasource-budget-alert_rule"
description: |-
  Provides details about a specific Alert Rule in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_alert_rule
This data source provides details about a specific Alert Rule resource in Oracle Cloud Infrastructure Budget service.

Gets an Alert Rule for a Budget by identifier

## Example Usage

```hcl
data "oci_budget_alert_rule" "test_alert_rule" {
	#Required
	alert_rule_id = "${oci_budget_alert_rule.test_alert_rule.id}"
	budget_id = "${oci_budget_budget.test_budget.id}"
}
```

## Argument Reference

The following arguments are supported:

* `alert_rule_id` - (Required) The unique Alert Rule OCID
* `budget_id` - (Required) The unique Budget OCID


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

