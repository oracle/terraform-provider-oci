---
subcategory: "Osub Subscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osub_subscription_subscriptions"
sidebar_current: "docs-oci-datasource-osub_subscription-subscriptions"
description: |-
  Provides the list of Subscriptions in Oracle Cloud Infrastructure Osub Subscription service
---

# Data Source: oci_osub_subscription_subscriptions
This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Osub Subscription service.

This list API returns all subscriptions for a given plan number or subscription id or buyer email 
and provides additional parameters to include ratecard and commitment details.
This API expects exactly one of the above mentioned parameters as input. If more than one parameters are provided the API will throw
a 400 - invalid parameters exception and if no parameters are provided it will throw a 400 - missing parameter exception


## Example Usage

```hcl
data "oci_osub_subscription_subscriptions" "test_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	buyer_email = var.subscription_buyer_email
	is_commit_info_required = var.subscription_is_commit_info_required
	plan_number = var.subscription_plan_number
	subscription_id = oci_osub_subscription_subscription.test_subscription.id
	x_one_gateway_subscription_id = var.subscription_x_one_gateway_subscription_id
	x_one_origin_region = var.subscription_x_one_origin_region
}
```

## Argument Reference

The following arguments are supported:

* `buyer_email` - (Optional) Buyer Email Id
* `compartment_id` - (Required) The OCID of the compartment.
* `is_commit_info_required` - (Optional) Boolean value to decide whether commitment services will be shown
* `plan_number` - (Optional) The Plan Number
* `subscription_id` - (Optional) Line level Subscription Id
* `x_one_gateway_subscription_id` - (Optional) This header is meant to be used only for internal purposes and will be ignored on any public request. The purpose of this header is  to help on Gateway to API calls identification.  
* `x_one_origin_region` - (Optional) The Oracle Cloud Infrastructure home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc. 


## Attributes Reference

The following attributes are exported:

* `subscriptions` - The list of subscriptions.

### Subscription Reference

The following attributes are exported:

* `currency` - Currency details 
	* `iso_code` - Currency Code 
	* `name` - Currency name 
	* `std_precision` - Standard Precision of the Currency 
* `service_name` - Customer friendly service name provided by PRG 
* `status` - Status of the plan 
* `subscribed_services` - List of Subscribed Services of the plan  
	* `booking_opty_number` - Booking Opportunity Number of Subscribed Service 
	* `commitment_services` - List of Commitment services of a line  
		* `available_amount` - Commitment available amount 
		* `funded_allocation_value` - Funded Allocation line value 
		* `line_net_amount` - Commitment line net amount 
		* `quantity` - Commitment quantity 
		* `time_end` - Commitment end date 
		* `time_start` - Commitment start date 
	* `csi` - Subscribed service CSI number 
	* `data_center_region` - Subscribed service data center region 
	* `funded_allocation_value` - Funded Allocation line value example: 12000.00   
	* `id` - SPM internal Subscribed Service ID 
	* `is_intent_to_pay` - Subscribed service intent to pay flag 
	* `net_unit_price` - Subscribed service net unit price 
	* `operation_type` - Subscribed service operation type 
	* `order_number` - Sales Order Number associated to the subscribed service 
	* `partner_transaction_type` - This field contains the name of the partner to which the subscription belongs - depending on which the invoicing may differ 
	* `pricing_model` - Subscribed service pricing model 
	* `product` - Product description 
		* `name` - Product name 
		* `part_number` - Product part numner 
		* `provisioning_group` - Product provisioning group 
		* `unit_of_measure` - Unit of measure 
	* `program_type` - Subscribed service program type 
	* `promo_type` - Subscribed service promotion type 
	* `quantity` - Subscribed service quantity 
	* `status` - Subscribed service status 
	* `term_value` - Term value in Months 
	* `term_value_uom` - Term value UOM 
	* `time_end` - Subscribed service end date 
	* `time_start` - Subscribed service start date 
	* `total_value` - Subscribed service total value 
* `time_end` - Represents the date when the last service of the subscription ends 
* `time_start` - Represents the date when the first service of the subscription was activated 

