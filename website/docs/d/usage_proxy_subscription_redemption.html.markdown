---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_redemption"
sidebar_current: "docs-oci-datasource-usage_proxy-subscription_redemption"
description: |-
  Provides details about a specific Subscription Redemption in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_subscription_redemption
This data source provides details about a specific Subscription Redemption resource in Oracle Cloud Infrastructure Usage Proxy service.

Returns the list of redemption for the subscription ID.


## Example Usage

```hcl
data "oci_usage_proxy_subscription_redemption" "test_subscription_redemption" {
	#Required
	subscription_id = oci_onesubscription_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id

	#Optional
	time_redeemed_greater_than_or_equal_to = var.subscription_redemption_time_redeemed_greater_than_or_equal_to
	time_redeemed_less_than = var.subscription_redemption_time_redeemed_less_than
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) The subscription ID for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.
* `time_redeemed_greater_than_or_equal_to` - (Optional) The starting redeemed date filter for the redemption history.
* `time_redeemed_less_than` - (Optional) The ending redeemed date filter for the redemption history.


## Attributes Reference

The following attributes are exported:

* `items` - The list of redemption summary.
	* `base_rewards` - It provides the redeemed rewards in base/subscription currency.
	* `fx_rate` - It provides the fxRate between invoice currency and subscription currency.
	* `invoice_currency` - The currency associated with invoice.
	* `invoice_number` - It provides the invoice number against the redemption.
	* `invoice_total_amount` - It provides the invoice total amount of given redemption.
	* `redeemed_rewards` - It provides the redeemed rewards in invoice currency.
	* `redemption_code` - The redemption code used in the Billing Center during the reward redemption process.
	* `redemption_email` - It provides the redemption email id.
	* `time_invoiced` - It provides the invoice date.
	* `time_redeemed` - It provides redeem date.

