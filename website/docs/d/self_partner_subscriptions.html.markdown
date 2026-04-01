---
subcategory: "Self"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_self_partner_subscriptions"
sidebar_current: "docs-oci-datasource-self-partner_subscriptions"
description: |-
  Provides the list of Partner Subscriptions in Oracle Cloud Infrastructure Self service
---

# Data Source: oci_self_partner_subscriptions
This data source provides the list of Partner Subscriptions in Oracle Cloud Infrastructure Self service.

Gets information about a Subscription.

## Example Usage

```hcl
data "oci_self_partner_subscriptions" "test_partner_subscriptions" {
	#Required
	listing_id = oci_marketplace_listing.test_listing.id

	#Optional
	display_name = var.partner_subscription_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given name.
* `listing_id` - (Required) The unique identifier for the listing.


## Attributes Reference

The following attributes are exported:

* `listing_subscriptions_collection` - The list of listing_subscriptions_collection.

### PartnerSubscription Reference

The following attributes are exported:

* `items` - List of subscriptions for particular listing.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The subscription name. Must be unique within the compartment. This value can be updated.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `lifecycle_details` - A message that describes the current state of the Subscription in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
	* `product_id` - The unique identifier of marketplace listing in Oracle Cloud Infrastructure.
	* `state` - The current state of the Subscription.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_ended` - The date and time the Subscription was ended, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
	* `time_started` - The date and time the Subscription was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

