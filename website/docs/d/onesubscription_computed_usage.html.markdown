---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_computed_usage"
sidebar_current: "docs-oci-datasource-onesubscription-computed_usage"
description: |-
  Provides details about a specific Computed Usage in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_computed_usage
This data source provides details about a specific Computed Usage resource in Oracle Cloud Infrastructure Onesubscription service.

This is an API which returns Computed Usage corresponding to the id passed


## Example Usage

```hcl
data "oci_onesubscription_computed_usage" "test_computed_usage" {
	#Required
	compartment_id = var.compartment_id
	computed_usage_id = oci_onesubscription_computed_usage.test_computed_usage.id

	#Optional
	fields = var.computed_usage_fields
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `computed_usage_id` - (Required) The Computed Usage Id
* `fields` - (Optional) Partial response refers to an optimization technique offered by the RESTful web APIs to return only the information  (fields) required by the client. This parameter is used to control what fields to return. 


## Attributes Reference

The following attributes are exported:

* `commitment_service_id` - Subscribed service commitmentId. 
* `compute_source` - SPM Internal compute records source . 
* `cost` - Computed Line Amount not rounded 
* `cost_rounded` - Computed Line Amount rounded. 
* `currency_code` - Currency code 
* `data_center` - Data Center Attribute as sent by MQS to SPM. 
* `id` - SPM Internal computed usage Id , 32 character string 
* `is_invoiced` - Invoicing status for the aggregated compute usage 
* `mqs_message_id` - MQS Identfier send to SPM , SPM does not transform this attribute and is received as is. 
* `net_unit_price` - Net Unit Price for the product in consideration, price actual. 
* `original_usage_number` - SPM Internal Original usage Line number identifier in SPM coming from Metered Services entity. 
* `parent_product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part number 
	* `product_category` - Product category 
	* `provisioning_group` - Product provisioning group 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of Measure 
* `parent_subscribed_service_id` - Subscribed service line parent id 
* `plan_number` - Subscription plan number 
* `product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part number 
	* `product_category` - Product category 
	* `provisioning_group` - Product provisioning group 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of Measure 
* `quantity` - Total Quantity that was used for computation 
* `rate_card_id` - Ratecard Id at subscribed service level 
* `rate_card_tierd_id` - References the tier in the ratecard for that usage (OCI will be using the same reference to cross-reference for correctness on the usage csv report), comes from Entity OBSCNTR_IPT_PRODUCTTIER. 
* `time_created` - Computed Usage created time, expressed in RFC 3339 timestamp format. 
* `time_metered_on` - Metered Service date, expressed in RFC 3339 timestamp format. 
* `time_of_arrival` - Usae computation date, expressed in RFC 3339 timestamp format. 
* `time_updated` - Computed Usage updated time, expressed in RFC 3339 timestamp format. 
* `type` - Usage compute type in SPM. 
* `unit_of_measure` - Unit of Messure 
* `usage_number` - SPM Internal usage Line number identifier in SPM coming from Metered Services entity. 

