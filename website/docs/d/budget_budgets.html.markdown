---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_budgets"
sidebar_current: "docs-oci-datasource-budget-budgets"
description: |-
  Provides the list of Budgets in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_budgets
This data source provides the list of Budgets in Oracle Cloud Infrastructure Budget service.

Gets a list of Budgets in a compartment.

By default, ListBudgets returns budgets of 'COMPARTMENT' target type and the budget records with only ONE target compartment OCID.

To list ALL budgets, set the targetType query parameter to ALL.
Example:
  'targetType=ALL'

Additional targetTypes would be available in future releases. Clients should ignore new targetType 
or upgrade to latest version of client SDK to handle new targetType.


## Example Usage

```hcl
data "oci_budget_budgets" "test_budgets" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	display_name = var.budget_display_name
	state = var.budget_state
	target_type = var.budget_target_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) The current state of the resource to filter by.
* `target_type` - (Optional) The type of target to filter by.
	* ALL - List all budgets
	* COMPARTMENT - List all budgets with targetType == "COMPARTMENT"
	* TAG - List all budgets with targetType == "TAG" 


## Attributes Reference

The following attributes are exported:

* `budgets` - The list of budgets.

### Budget Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle
* `alert_rule_count` - Total number of alert rules in the budget
* `amount` - The amount of the budget expressed in the currency of the customer's rate card. 
* `compartment_id` - The OCID of the compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the budget.
* `display_name` - The display name of the budget.
* `forecasted_spend` - The forecasted spend in currency by the end of the current budget cycle
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the budget
* `reset_period` - The reset period for the budget. 
* `state` - The current state of the budget.
* `target_compartment_id` - This is DEPRECATED. For backwards compatability, the property will be populated when targetType is "COMPARTMENT" AND targets contains EXACT ONE target compartment ocid. For all other scenarios, this property will be left empty. 
* `target_type` - The type of target on which the budget is applied. 
* `targets` - The list of targets on which the budget is applied. If targetType is "COMPARTMENT", targets contains list of compartment OCIDs. If targetType is "TAG", targets contains list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}". 
* `time_created` - Time that budget was created
* `time_spend_computed` - The time that the budget spend was last computed
* `time_updated` - Time that budget was updated
* `version` - Version of the budget. Starts from 1 and increments by 1.

