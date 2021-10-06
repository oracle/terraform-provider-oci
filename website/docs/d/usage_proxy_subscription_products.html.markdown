---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_subscription_products"
sidebar_current: "docs-oci-datasource-usage_proxy-subscription_products"
description: |-
  Provides the list of Subscription Products in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_subscription_products
This data source provides the list of Subscription Products in Oracle Cloud Infrastructure Usage Proxy service.

This API provides usage period specific product and its usage details.


## Example Usage

```hcl
data "oci_usage_proxy_subscription_products" "test_subscription_products" {
	#Required
	subscription_id = oci_ons_subscription.test_subscription.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id
	usage_period_key = var.subscription_product_usage_period_key

	#Optional
	producttype = var.subscription_product_producttype
}
```

## Argument Reference

The following arguments are supported:

* `producttype` - (Optional) The field to specify the type of product.
* `subscription_id` - (Required) The subscriptionId for which rewards information is requested for.
* `tenancy_id` - (Required) The OCID of the tenancy.
* `usage_period_key` - (Required) The SPM Identifier for the usage period.


## Attributes Reference

The following attributes are exported:

* `product_collection` - The list of product_collection.

### SubscriptionProduct Reference

The following attributes are exported:

* `items` - The list of product rewards summaries.
	* `earned_rewards` - The earned rewards for the product.
	* `is_eligible_to_earn_rewards` - The boolean flag to tell if the product is eligible for earning rewards.
	* `product_name` - The ratecard product Name.
	* `product_number` - The ratecard product number.
	* `usage_amount` - The ratecard product usage amount.

