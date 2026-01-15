---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_cost_alert_subscription"
sidebar_current: "docs-oci-datasource-budget-cost_alert_subscription"
description: |-
  Provides details about a specific Cost Alert Subscription in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_cost_alert_subscription
This data source provides details about a specific Cost Alert Subscription resource in Oracle Cloud Infrastructure Budget service.

Gets a CostAlertSubscription by the identifier.

## Example Usage

```hcl
data "oci_budget_cost_alert_subscription" "test_cost_alert_subscription" {
	#Required
	cost_alert_subscription_id = oci_budget_cost_alert_subscription.test_cost_alert_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `cost_alert_subscription_id` - (Required) The unique costAlertSubscription OCID.


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

