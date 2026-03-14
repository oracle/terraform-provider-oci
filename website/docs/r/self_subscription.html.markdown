---
subcategory: "Self"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_self_subscription"
sidebar_current: "docs-oci-resource-self-subscription"
description: |-
  Provides the Subscription resource in Oracle Cloud Infrastructure Self service
---

# oci_self_subscription
This resource provides the Subscription resource in Oracle Cloud Infrastructure Self service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/self

Creates a Subscription.


## Example Usage

```hcl
resource "oci_self_subscription" "test_subscription" {
	#Required
	compartment_id = var.compartment_id
	product_id = oci_self_product.test_product.id
	seller_id = oci_self_seller.test_seller.id
	subscription_details {
		#Required
		billing_details {
			#Required
			meters {
				#Required
				name = var.subscription_subscription_details_billing_details_meters_name
				rate_allocation = var.subscription_subscription_details_billing_details_meters_rate_allocation

				#Optional
				extended_metadata {
					#Required
					key = var.subscription_subscription_details_billing_details_meters_extended_metadata_key
					value = var.subscription_subscription_details_billing_details_meters_extended_metadata_value
				}
			}
			metric_type = var.subscription_subscription_details_billing_details_metric_type
			rate_allocation = var.subscription_subscription_details_billing_details_rate_allocation
			sku = var.subscription_subscription_details_billing_details_sku

			#Optional
			has_gov_sku = var.subscription_subscription_details_billing_details_has_gov_sku
		}
		partner_registration_url = var.subscription_subscription_details_partner_registration_url
		pricing_plan {
			#Required
			billing_frequency = var.subscription_subscription_details_pricing_plan_billing_frequency
			plan_name = var.subscription_subscription_details_pricing_plan_plan_name
			plan_type = var.subscription_subscription_details_pricing_plan_plan_type
			rates {
				#Required
				currency = var.subscription_subscription_details_pricing_plan_rates_currency
				rate = var.subscription_subscription_details_pricing_plan_rates_rate
			}

			#Optional
			plan_description = var.subscription_subscription_details_pricing_plan_plan_description
			plan_duration = var.subscription_subscription_details_pricing_plan_plan_duration
		}

		#Optional
		amount = var.subscription_subscription_details_amount
		currency = var.subscription_subscription_details_currency
		is_auto_renew = var.subscription_subscription_details_is_auto_renew
	}
	tenant_id = oci_self_tenant.test_tenant.id

	#Optional
	additional_details {
		#Required
		key = var.subscription_additional_details_key
		value = var.subscription_additional_details_value
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.subscription_display_name
	freeform_tags = {"Department"= "Finance"}
	realm = var.subscription_realm
	region = var.subscription_region
	source_type = var.subscription_source_type
}
```

## Argument Reference

The following arguments are supported:

* `additional_details` - (Optional) Additional details that are specific for this subscription such as activation details.
	* `key` - (Required) Additional attribute for extendedMetadata.
	* `value` - (Required) It contains the value of above key.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the subscription in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The subscription name. Must be unique within the compartment. This value can be updated. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `product_id` - (Required) The unique identifier of the marketplace listing in Oracle Cloud Infrastructure.
* `realm` - (Optional) The realm from where customer is buying the subscription.
* `region` - (Optional) The region from where customer is buying the subscription.
* `seller_id` - (Required) The OCID for the seller in SELF Service.
* `source_type` - (Optional) The type of seller in SELF Service.
* `subscription_details` - (Required) The details of a subscription
	* `amount` - (Optional) Tha amount for the currency type.
	* `billing_details` - (Required) Sku details for billing subscription.
		* `has_gov_sku` - (Optional) Whether this sku is assign to gov product.
		* `meters` - (Required) The meters associated with sku.
			* `extended_metadata` - (Optional) Additional data give by sku.
				* `key` - (Required) Additional attribute for extendedMetadata.
				* `value` - (Required) It contains the value of above key.
			* `name` - (Required) Name of meter.
			* `rate_allocation` - (Required) Tha rate of this sku meter.
		* `metric_type` - (Required) The part's metric.
		* `rate_allocation` - (Required) Tha rate of this sku meter.
		* `sku` - (Required) Sku for service.
	* `currency` - (Optional) The currency supported, in the format specified by ISO-4217
	* `is_auto_renew` - (Optional) Whether subscription should be auto-renewed at the end of cycle.
	* `partner_registration_url` - (Required) The activation link given by the partner.
	* `pricing_plan` - (Required) A pricing plan details provided by the Publisher.
		* `billing_frequency` - (Required) Specifies the interval at which billing occurs for the subscription plan.
		* `plan_description` - (Optional) A detailed explanation of the subscription plan.
		* `plan_duration` - (Optional) Specifies the interval at which billing occurs for the subscription plan.
		* `plan_name` - (Required) The name of the subscription plan used to identify the plan.
		* `plan_type` - (Required) The type of the subscription plan.
		* `rates` - (Required) The pricing details of the subscription plan in various supported currencies.
			* `currency` - (Required) The currency supported, in the format specified by ISO-4217
			* `rate` - (Required) The amount charged for the plan in the specified currency.
* `tenant_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenant to create the subscription in. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscription
	* `update` - (Defaults to 20 minutes), when updating the Subscription
	* `delete` - (Defaults to 20 minutes), when destroying the Subscription


## Import

Subscriptions can be imported using the `id`, e.g.

```
$ terraform import oci_self_subscription.test_subscription "id"
```

