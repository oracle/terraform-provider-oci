---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_alert_rules"
sidebar_current: "docs-oci-datasource-budget-alert_rules"
description: |-
  Provides the list of Alert Rules in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_alert_rules
This data source provides the list of Alert Rules in Oracle Cloud Infrastructure Budget service.

Returns a list of Alert Rules for a specified Budget.


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

