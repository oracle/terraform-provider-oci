---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_redeemable_users"
sidebar_current: "docs-oci-datasource-usage_proxy-subscription_redeemable_users"
description: |-
  Provides the list of Subscription Redeemable Users in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_subscription_redeemable_users
This data source provides the list of Subscription Redeemable Users in Oracle Cloud Infrastructure Usage Proxy service.

Provides emailids of redeemable users for the given subscriptionId


## Example Usage

```hcl
data "oci_usage_proxy_subscription_redeemable_users" "test_subscription_redeemable_users" {
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

* `redeemable_user_collection` - The list of redeemable_user_collection.

### SubscriptionRedeemableUser Reference

The following attributes are exported:

* `items` - The list of redeemable users email Ids.
	* `email_id` - The email Id of Redeemable User.

