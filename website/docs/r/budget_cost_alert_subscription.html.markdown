---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_cost_alert_subscription"
sidebar_current: "docs-oci-resource-budget-cost_alert_subscription"
description: |-
  Provides the Cost Alert Subscription resource in Oracle Cloud Infrastructure Budget service
---

# oci_budget_cost_alert_subscription
This resource provides the Cost Alert Subscription resource in Oracle Cloud Infrastructure Budget service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/budgets/latest/CostAlertSubscription

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/budget

Creates a new CostAlert Subscription.


## Example Usage

```hcl
resource "oci_budget_cost_alert_subscription" "test_cost_alert_subscription" {
	#Required
	channels = var.cost_alert_subscription_channels
	compartment_id = var.compartment_id
	name = var.cost_alert_subscription_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.cost_alert_subscription_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `channels` - (Required) (Updatable) The notification channels string.
* `compartment_id` - (Required) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the cost alert subscription.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the cost alert subscription. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `channels` - The notification channels string.
* `compartment_id` - The OCID of the compartment which hold the cost alert subscription resource.
* `cost_anomaly_monitors` - List of monitor identifiers
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the cost alert subscription.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Cost Alert Subscription.
* `name` - The name of the cost alert subscription. Avoid entering confidential information.
* `state` - The current state of the cost alert subscription.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the cost alert subscription was created.
* `time_updated` - The time that the cost alert subscription was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cost Alert Subscription
	* `update` - (Defaults to 20 minutes), when updating the Cost Alert Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Cost Alert Subscription


## Import

CostAlertSubscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_budget_cost_alert_subscription.test_cost_alert_subscription "id"
```

