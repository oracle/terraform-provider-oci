---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_alert_rule"
sidebar_current: "docs-oci-datasource-budget-alert_rule"
description: |-
  Provides details about a specific Alert Rule in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_alert_rule
This data source provides details about a specific Alert Rule resource in Oracle Cloud Infrastructure Budget service.

Gets an Alert Rule for a specified budget.

## Example Usage

```hcl
data "oci_budget_alert_rule" "test_alert_rule" {
	#Required
	alert_rule_id = oci_budget_alert_rule.test_alert_rule.id
	budget_id = oci_budget_budget.test_budget.id
}
```

## Argument Reference

The following arguments are supported:

* `alert_rule_id` - (Required) The unique Alert Rule OCID.
* `budget_id` - (Required) The unique budget OCID.


## Attributes Reference

The following attributes are exported:

* `budget_id` - The OCID of the budget.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the alert rule.
* `display_name` - The name of the alert rule. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the alert rule.
* `message` - The custom message that will be sent when the alert is triggered.
* `recipients` - The delimited list of email addresses to receive the alert when it triggers. Delimiter characters can be a comma, space, TAB, or semicolon. 
* `state` - The current state of the alert rule.
* `threshold` - The threshold for triggering the alert. If the thresholdType is PERCENTAGE, the maximum value is 10000. 
* `threshold_type` - The type of threshold.
* `time_created` - The time when the budget was created.
* `time_updated` - The time when the budget was updated.
* `type` - The type of the alert. Valid values are ACTUAL (the alert triggers based on actual usage), or FORECAST (the alert triggers based on predicted usage). 
* `version` - The version of the alert rule. Starts from 1 and increments by 1.

