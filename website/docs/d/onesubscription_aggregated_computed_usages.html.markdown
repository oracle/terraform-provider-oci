---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_aggregated_computed_usages"
sidebar_current: "docs-oci-datasource-onesubscription-aggregated_computed_usages"
description: |-
  Provides the list of Aggregated Computed Usages in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_aggregated_computed_usages
This data source provides the list of Aggregated Computed Usages in Oracle Cloud Infrastructure Onesubscription service.

This is a collection API which returns a list of aggregated computed usage details (there can be multiple Parent Products under a given SubID each of which is represented under Subscription Service Line # in SPM).


## Example Usage

```hcl
data "oci_onesubscription_aggregated_computed_usages" "test_aggregated_computed_usages" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_onesubscription_subscription.test_subscription.id
	time_from = var.aggregated_computed_usage_time_from
	time_to = var.aggregated_computed_usage_time_to

	#Optional
	grouping = var.aggregated_computed_usage_grouping
	parent_product {
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `grouping` - (Optional) Grouping criteria to use for aggregate the computed Usage, either hourly (`HOURLY`), daily (`DAILY`), monthly(`MONTHLY`) or none (`NONE`) to not follow a grouping criteria by date. 
* `parent_product` - (Optional) Product part number for subscribed service line, called parent product. 
* `subscription_id` - (Required) Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM. 
* `time_from` - (Required) Initial date to filter Computed Usage data in SPM. In the case of non aggregated data the time period between of fromDate and toDate , expressed in RFC 3339 timestamp format. 
* `time_to` - (Required) Final date to filter Computed Usage data in SPM, expressed in RFC 3339 timestamp format. 


## Attributes Reference

The following attributes are exported:

* `aggregated_computed_usages` - The list of aggregated_computed_usages.

### AggregatedComputedUsage Reference

The following attributes are exported:

* `aggregated_computed_usages` - Aggregation of computed usages for the subscribed service. 
	* `cost` - Sum of Computed Line Amount rounded 
	* `cost_unrounded` - Sum of Computed Line Amount unrounded 
	* `data_center` - Data Center Attribute as sent by MQS to SPM. 
	* `net_unit_price` - Net Unit Price for the product in consideration. 
	* `product` - Product description 
		* `billing_category` - Metered service billing category 
		* `name` - Product name 
		* `part_number` - Product part number 
		* `product_category` - Product category 
		* `provisioning_group` - Product provisioning group 
		* `ucm_rate_card_part_type` - Rate card part type of Product 
		* `unit_of_measure` - Unit of Measure 
	* `quantity` - Total Quantity that was used for computation 
	* `time_metered_on` - Metered Service date , expressed in RFC 3339 timestamp format. 
	* `type` - Usage compute type in SPM. 
* `currency_code` - Currency code 
* `parent_product` - Product description 
	* `billing_category` - Metered service billing category 
	* `name` - Product name 
	* `part_number` - Product part number 
	* `product_category` - Product category 
	* `provisioning_group` - Product provisioning group 
	* `ucm_rate_card_part_type` - Rate card part type of Product 
	* `unit_of_measure` - Unit of Measure 
* `parent_subscribed_service_id` - Subscribed service line parent id 
* `plan_number` - Subscribed service asociated subscription plan number. 
* `pricing_model` - Subscribed services pricing model 
* `rate_card_id` - Inernal SPM Ratecard Id at line level 
* `subscription_id` - Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM 
* `time_end` - Subscribed services contract line end date, expressed in RFC 3339 timestamp format. 
* `time_start` - Subscribed services contract line start date, expressed in RFC 3339 timestamp format. 

