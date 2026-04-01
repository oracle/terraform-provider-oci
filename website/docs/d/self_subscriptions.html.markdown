---
subcategory: "Self"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_self_subscriptions"
sidebar_current: "docs-oci-datasource-self-subscriptions"
description: |-
  Provides the list of Subscriptions in Oracle Cloud Infrastructure Self service
---

# Data Source: oci_self_subscriptions
This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Self service.

Lists the subscriptions which have been created in the specified compartment.
You can filter results by specifying query parameters.


## Example Usage

```hcl
data "oci_self_subscriptions" "test_subscriptions" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.subscription_display_name
	id = var.subscription_id
	lifecycle_details = var.subscription_lifecycle_details
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given name.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Subscription.
* `lifecycle_details` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `subscription_collection` - The list of subscription_collection.

### Subscription Reference

The following attributes are exported:

* `additional_details` - Additional details that are specific for this subscription such as activation details.
	* `key` - Additional attribute for extendedMetadata.
	* `value` - It contains the value of above key.
* `compartment_id` - The unique identifier for the compartment where the subscription was purchased.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The subscription name. Must be unique within the compartment. This value can be updated.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The unique identifier for the subscription within a specific compartment.
* `lifecycle_details` - A message that describes the current state of the Subscription in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `product_id` - The unique OCID of the product, effectively functioning as the listing ID.
* `realm` - The realm from where customer is buying the subscription.
* `region` - The region from where customer is buying the subscription.
* `seller_id` - The OCID that identifies the seller within the platform.
* `source_type` - The type of seller in SELF Service.
* `state` - The current lifecycle state of the Subscription.
* `subscription_details` - The details of a subscription
	* `amount` - Tha amount for the currency type.
	* `billing_details` - Sku details for billing subscription.
		* `has_gov_sku` - Whether this sku is assign to gov product.
		* `meters` - The meters associated with sku.
			* `extended_metadata` - Additional data give by sku.
				* `key` - Additional attribute for extendedMetadata.
				* `value` - It contains the value of above key.
			* `name` - Name of meter.
			* `rate_allocation` - Tha rate of this sku meter.
		* `metric_type` - The part's metric.
		* `rate_allocation` - Tha rate of this sku meter.
		* `sku` - Sku for service.
	* `currency` - The currency supported, in the format specified by ISO-4217
	* `is_auto_renew` - Whether subscription should be auto-renewed at the end of cycle.
	* `partner_registration_url` - The activation link given by the partner.
	* `pricing_plan` - A pricing plan details provided by the Publisher.
		* `billing_frequency` - Specifies the interval at which billing occurs for the subscription plan.
		* `plan_description` - A detailed explanation of the subscription plan.
		* `plan_duration` - Specifies the interval at which billing occurs for the subscription plan.
		* `plan_name` - The name of the subscription plan used to identify the plan.
		* `plan_type` - The type of the subscription plan.
		* `rates` - The pricing details of the subscription plan in various supported currencies.
			* `currency` - The currency supported, in the format specified by ISO-4217
			* `rate` - The amount charged for the plan in the specified currency.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - The unique identifier for the tenant where the subscription was purchased.
* `time_created` - The date and time the Subscription was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_ended` - The date and time the Subscription was ended, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_started` - The date and time the Subscription was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Subscription was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

