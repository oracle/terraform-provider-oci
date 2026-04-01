---
subcategory: "Self"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_self_subscription_token"
sidebar_current: "docs-oci-datasource-self-subscription_token"
description: |-
  Provides details about a specific Subscription Token in Oracle Cloud Infrastructure Self service
---

# Data Source: oci_self_subscription_token
This data source provides details about a specific Subscription Token resource in Oracle Cloud Infrastructure Self service.

Gets a token of Subscriptions.


## Example Usage

```hcl
data "oci_self_subscription_token" "test_subscription_token" {
	#Required
	subscription_id = oci_self_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The unique identifier for the subscription.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `subscription_id` - The unique identifier for the subscription within a specific compartment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `token` - JWT token of subscriptions.

