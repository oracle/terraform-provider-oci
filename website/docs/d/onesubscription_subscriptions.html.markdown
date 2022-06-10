---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_subscriptions"
sidebar_current: "docs-oci-datasource-onesubscription-subscriptions"
description: |-
  Provides the list of Subscriptions in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_subscriptions
This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Onesubscription service.

This list API returns all subscriptions for a given plan number or subscription id or buyer email 
and provides additional parameters to include ratecard and commitment details.
This API expects exactly one of the above mentioned parameters as input. If more than one parameters are provided the API will throw
a 400 - invalid parameters exception and if no parameters are provided it will throw a 400 - missing parameter exception


## Example Usage

```hcl
data "oci_onesubscription_subscriptions" "test_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	buyer_email = var.subscription_buyer_email
	is_commit_info_required = var.subscription_is_commit_info_required
	plan_number = var.subscription_plan_number
	subscription_id = oci_onesubscription_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `buyer_email` - (Optional) Buyer Email Id
* `compartment_id` - (Required) The OCID of the root compartment.
* `is_commit_info_required` - (Optional) Boolean value to decide whether commitment services will be shown
* `plan_number` - (Optional) The Plan Number
* `subscription_id` - (Optional) Line level Subscription Id


## Attributes Reference

The following attributes are exported:

* `subscriptions` - The list of subscriptions.

### Subscription Reference

The following attributes are exported:

* `currency` - Currency details 
	* `iso_code` - Currency Code 
	* `name` - Currency name 
	* `std_precision` - Standard Precision of the Currency 
* `hold_reason` - Hold reason of the plan 
* `service_name` - Customer friendly service name provided by PRG 
* `status` - Status of the plan 
* `subscribed_services` - List of Subscribed Services of the plan  
	* `available_amount` - Subscribed sercice available or remaining amount 
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
	* `original_promo_amount` - Subscribed service Promotion Amount 
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
	* `used_amount` - Subscribed service used amount 
* `time_end` - Represents the date when the last service of the subscription ends 
* `time_hold_release_eta` - Represents the date of the hold release 
* `time_start` - Represents the date when the first service of the subscription was activated 

