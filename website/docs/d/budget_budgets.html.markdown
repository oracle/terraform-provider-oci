---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_budgets"
sidebar_current: "docs-oci-datasource-budget-budgets"
description: |-
  Provides the list of Budgets in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_budgets
This data source provides the list of Budgets in Oracle Cloud Infrastructure Budget service.

Returns a list of Budgets.


## Example Usage

```hcl
data "oci_budget_budgets" "test_budgets" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.budget_display_name}"
	state = "${var.budget_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) The current state of the resource to filter by.


## Attributes Reference

The following attributes are exported:

* `budgets` - The list of budgets.

### Budget Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle
* `alert_rule_count` - Total number of alert rules in the budget
* `amount` - The amount of the budget expressed as a decimal number in the currency of the customer's rate card. 
* `compartment_id` - The OCID of the compartment
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the budget.
* `display_name` - The display name of the budget.
* `forecasted_spend` - The forecasted spend in currency by the end of the current budget cycle
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the budget
* `reset_period` - The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP. 
* `state` - The current state of the budget.
* `target_compartment_id` - The OCID of the compartment on which budget is applied
* `time_created` - Time when budget was created
* `time_spend_computed` - Time when the budget spend was last computed
* `time_updated` - Time when budget was updated
* `version` - Version of the budget. Starts from 1 and increments by 1.

