---
subcategory: "Osub Usage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osub_usage_computed_usages"
sidebar_current: "docs-oci-datasource-osub_usage-computed_usages"
description: |-
  Provides the list of Computed Usages in Oracle Cloud Infrastructure Osub Usage service
---

# Data Source: oci_osub_usage_computed_usages
This data source provides the list of Computed Usages in Oracle Cloud Infrastructure Osub Usage service.

This is a collection API which returns a list of Computed Usages for given filters.


## Example Usage

```hcl
data "oci_osub_usage_computed_usages" "test_computed_usages" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_ons_subscription.test_subscription.id
	time_from = var.computed_usage_time_from
	time_to = var.computed_usage_time_to

	#Optional
	computed_product = var.computed_usage_computed_product
	parent_product {
	}
	x_one_origin_region = var.computed_usage_x_one_origin_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `computed_product` - (Optional) Product part number for Computed Usage . 
* `parent_product` - (Optional) Product part number for subscribed service line, called parent product. 
* `subscription_id` - (Required) Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM.  
* `time_from` - (Required) Initial date to filter Computed Usage data in SPM. In the case of non aggregated data the time period between of fromDate and toDate , expressed in RFC 3339 timestamp format. 
* `time_to` - (Required) Final date to filter Computed Usage data in SPM, expressed in RFC 3339 timestamp format. 
* `x_one_origin_region` - (Optional) The Oracle Cloud Infrastructure home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc. 


## Attributes Reference

The following attributes are exported:

* `computed_usages` - The list of computed_usages.

### ComputedUsage Reference

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

