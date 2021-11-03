---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_redeemable_user"
sidebar_current: "docs-oci-datasource-usage_proxy-subscription_redeemable_user"
description: |-
  Provides details about a specific Subscription Redeemable User in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_subscription_redeemable_user
This data source provides details about a specific Subscription Redeemable User resource in Oracle Cloud Infrastructure Usage Proxy service.

Provides emailids of redeemable users for the given subscriptionId


## Example Usage

```hcl
data "oci_usage_proxy_subscription_redeemable_user" "test_subscription_redeemable_user" {
	#Required
	subscription_id = oci_ons_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The subscriptionId for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.


## Attributes Reference

The following attributes are exported:

* `items` - The list of redeemable users email Ids.
	* `email_id` - The email Id of Redeemable User.

