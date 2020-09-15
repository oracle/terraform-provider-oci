---
subcategory: "Budget"
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
	budget_id = oci_budget_budget.test_budget.id
}
```

## Argument Reference

The following arguments are supported:

* `budget_id` - (Required) The unique Budget OCID


## Attributes Reference

The following attributes are exported:

* `actual_spend` - The actual spend in currency for the current budget cycle
* `alert_rule_count` - Total number of alert rules in the budget
* `amount` - The amount of the budget expressed in the currency of the customer's rate card. 
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

