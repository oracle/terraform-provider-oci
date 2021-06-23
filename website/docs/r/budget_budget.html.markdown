---
subcategory: "Budget"
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
	amount = var.budget_amount
	compartment_id = var.tenancy_ocid
	reset_period = var.budget_reset_period

	#Optional
	budget_processing_period_start_offset = var.budget_budget_processing_period_start_offset
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.budget_description
	display_name = var.budget_display_name
	freeform_tags = {"Department"= "Finance"}
	target_compartment_id = oci_identity_compartment.test_compartment.id
	target_type = var.budget_target_type
	targets = var.budget_targets
}
```

## Argument Reference

The following arguments are supported:

* `amount` - (Required) (Updatable) The amount of the budget expressed as a whole number in the currency of the customer's rate card. 
* `budget_processing_period_start_offset` - (Optional) (Updatable) The number of days offset from the first day of the month, at which the budget processing period starts. In months that have fewer days than this value, processing will begin on the last day of that month. For example, for a value of 12, processing starts every month on the 12th at midnight.
* `compartment_id` - (Required) The OCID of the tenancy
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the budget.
* `display_name` - (Optional) (Updatable) The displayName of the budget.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `reset_period` - (Required) (Updatable) The reset period for the budget. Valid value is MONTHLY.
* `target_compartment_id` - (Optional) This is DEPRECTAED. Set the target compartment id in targets instead. 
* `target_type` - (Optional) The type of target on which the budget is applied. 
* `targets` - (Optional) The list of targets on which the budget is applied. If targetType is "COMPARTMENT", targets contains list of compartment OCIDs. If targetType is "TAG", targets contains list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}". Curerntly, the array should contain EXACT ONE item. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle
* `alert_rule_count` - Total number of alert rules in the budget
* `amount` - The amount of the budget expressed in the currency of the customer's rate card. 
* `budget_processing_period_start_offset` - The number of days offset from the first day of the month, at which the budget processing period starts. In months that have fewer days than this value, processing will begin on the last day of that month. For example, for a value of 12, processing starts every month on the 12th at midnight.
* `compartment_id` - The OCID of the tenancy
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Budget
	* `update` - (Defaults to 20 minutes), when updating the Budget
	* `delete` - (Defaults to 20 minutes), when destroying the Budget


## Import

Budgets can be imported using the `id`, e.g.

```
$ terraform import oci_budget_budget.test_budget "id"
```

