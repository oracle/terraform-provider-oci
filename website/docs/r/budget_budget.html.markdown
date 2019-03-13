---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_budget"
sidebar_current: "docs-oci-resource-budget-budget"
description: |-
  Provides the Budget resource in Oracle Cloud Infrastructure Budget service
---

# oci_budget_budget
This resource provides the Budget resource in Oracle Cloud Infrastructure Budget service.

Creates a new Budget.


## Example Usage

```hcl
resource "oci_budget_budget" "test_budget" {
	#Required
	amount = "${var.budget_amount}"
	compartment_id = "${var.compartment_id}"
	reset_period = "${var.budget_reset_period}"
	target_compartment_id = "${oci_budget_target_compartment.test_target_compartment.id}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = "${var.budget_description}"
	display_name = "${var.budget_display_name}"
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `amount` - (Required) (Updatable) The amount of the budget expressed as a decimal number in the currency of the customer's rate card. 
* `compartment_id` - (Required) The OCID of the compartment
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The description of the budget.
* `display_name` - (Optional) (Updatable) The displayName of the budget.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `reset_period` - (Required) (Updatable) The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP. 
* `target_compartment_id` - (Required) The OCID of the compartment on which budget is applied


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

Budgets can be imported using the `id`, e.g.

```
$ terraform import oci_budget_budget.test_budget "id"
```

