---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_ratecards"
sidebar_current: "docs-oci-datasource-onesubscription-ratecards"
description: |-
  Provides the list of Ratecards in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_ratecards
This data source provides the list of Ratecards in Oracle Cloud Infrastructure Onesubscription service.

List API that returns all ratecards for given Subscription Id and Account ID (if provided) and
for a particular date range


## Example Usage

```hcl
data "oci_onesubscription_ratecards" "test_ratecards" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_onesubscription_subscription.test_subscription.id

	#Optional
	part_number = var.ratecard_part_number
	time_from = var.ratecard_time_from
	time_to = var.ratecard_time_to
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `part_number` - (Optional) This param is used to get the rate card(s) filterd by the partNumber
* `subscription_id` - (Required) Line level Subscription Id
* `time_from` - (Optional) This param is used to get the rate card(s) whose effective start date starts on or after a particular date
* `time_to` - (Optional) This param is used to get the rate card(s) whose effective end date ends on or before a particular date


## Attributes Reference

The following attributes are exported:

* `rate_cards` - The list of rate_cards.

### Ratecard Reference

The following attributes are exported:

* `currency` - Currency details 
	* `iso_code` - Currency Code 
	* `name` - Currency name 
	* `std_precision` - Standard Precision of the Currency 
* `discretionary_discount_percentage` - Rate card discretionary discount percentage 
* `is_tier` - Rate card price tier flag 
* `net_unit_price` - Rate card net unit price 
* `overage_price` - Rate card overage price 
* `product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part numner 
	* `product_category` - Product category 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of measure 
* `rate_card_tiers` - List of tiered rate card prices 
	* `net_unit_price` - Rate card tier net unit price 
	* `overage_price` - Rate card tier overage price 
	* `up_to_quantity` - Rate card tier quantity range 
* `subscribed_service_id` - SPM internal Subscribed Service ID 
* `time_end` - Rate card end date 
* `time_start` - Rate card start date 

