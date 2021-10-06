---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_redeemable_user"
sidebar_current: "docs-oci-resource-usage_proxy-subscription_redeemable_user"
description: |-
  Provides the Subscription Redeemable User resource in Oracle Cloud Infrastructure Usage Proxy service
---

# oci_usage_proxy_subscription_redeemable_user
This resource provides the Subscription Redeemable User resource in Oracle Cloud Infrastructure Usage Proxy service.

Add list of redeemable user email ids for a subscription Id


## Example Usage

```hcl
resource "oci_usage_proxy_subscription_redeemable_user" "test_subscription_redeemable_user" {
	#Required
	subscription_id = oci_ons_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id

	#Optional
	items {
		#Required
		email_id = oci_usage_proxy_email.test_email.id
	}
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `items` - (Optional) The list of email ids to be added to the redeemable users list.
	* `email_id` - (Required) The Email Id for redeemable user.
* `subscription_id` - (Required) The subscriptionId for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.
* `user_id` - (Optional) The user Id of the person whose email we cc


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - The list of redeemable users email Ids.
	* `email_id` - The email Id of Redeemable User.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription Redeemable User
	* `update` - (Defaults to 20 minutes), when updating the Subscription Redeemable User
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription Redeemable User


## Import

SubscriptionRedeemableUsers can be imported using the `id`, e.g.

```
$ terraform import oci_usage_proxy_subscription_redeemable_user.test_subscription_redeemable_user "subscriptions/{subscriptionId}/redeemableUsers/tenancyId/{tenancyId}" 
```

