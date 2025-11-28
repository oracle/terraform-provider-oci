---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_cost_alert_subscriptions"
sidebar_current: "docs-oci-datasource-budget-cost_alert_subscriptions"
description: |-
  Provides the list of Cost Alert Subscriptions in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_cost_alert_subscriptions
This data source provides the list of Cost Alert Subscriptions in Oracle Cloud Infrastructure Budget service.

Gets a list of Cost Alert Subscription in a compartment.


## Example Usage

```hcl
data "oci_budget_cost_alert_subscriptions" "test_cost_alert_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.cost_alert_subscription_name
	state = var.cost_alert_subscription_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `name` - (Optional) Unique, non-changeable resource name. 
* `state` - (Optional) The current state of the cost alert subscription.


## Attributes Reference

The following attributes are exported:

* `cost_alert_subscription_collection` - The list of cost_alert_subscription_collection.

### CostAlertSubscription Reference

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

