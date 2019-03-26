---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_budget"
sidebar_current: "docs-oci-datasource-budget-budget"
description: |-
  Provides details about a specific Budget in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_budget
This data source provides details about a specific Budget resource in Oracle Cloud Infrastructure Budget service.

Gets a Budget by identifier

## Example Usage

```hcl
data "oci_budget_budget" "test_budget" {
	#Required
	budget_id = "${oci_budget_budget.test_budget.id}"
}
```

## Argument Reference

The following arguments are supported:

* `budget_id` - (Required) The unique Budget OCID


## Attributes Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle
* `alert_rule_count` - Total number of alert rules in the budget
* `amount` - The amount of the budget expressed as a decimal number in the currency of the customer's rate card. 
* `compartment_id` - The OCID of the tenancy
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

