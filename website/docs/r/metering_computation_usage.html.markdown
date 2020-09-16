---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage"
sidebar_current: "docs-oci-resource-metering_computation-usage"
description: |-
  Provides the Usage resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_usage
This resource provides the Usage resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the usage for the given account


## Example Usage

```hcl
resource "oci_metering_computation_usage" "test_usage" {
	#Required
	granularity = var.usage_granularity
	tenant_id = oci_metering_computation_tenant.test_tenant.id
	time_usage_ended = var.usage_time_usage_ended
	time_usage_started = var.usage_time_usage_started

	#Optional
	compartment_depth = var.usage_compartment_depth
	filter = var.usage_filter
	group_by = var.usage_group_by
	query_type = var.usage_query_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_depth` - (Optional) The depth level of the compartment.
* `filter` - (Optional) 
* `granularity` - (Required) The granularity of the usage. HOURLY - Hourly aggregation of data DAILY - Daily aggregation of data MONTHLY - Monthly aggregation of data TOTAL - Not Supported Yet 
* `group_by` - (Optional) Aggregate the result by. example: `["service"]` 
* `query_type` - (Optional) The type of query of the usage. Usage - Query the usage data. Cost - Query the cost / billing data. 
* `tenant_id` - (Required) tenant id
* `time_usage_ended` - (Required) The end time of the usage.
* `time_usage_started` - (Required) The start time of the usage.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `group_by` - Aggregate the result by.
* `items` - A list of usage items.
	* `ad` - The availability domain of the usage.
	* `compartment_id` - The OCID of the compartment.
	* `compartment_name` - The name of the compartment.
	* `compartment_path` - The path of the compartment, starting from root.
	* `computed_amount` - The computed cost.
	* `computed_quantity` - The usage number.
	* `currency` - The currency for the price.
	* `discount` - The discretionary discount applied to the SKU.
	* `list_rate` - The list rate for the SKU (not discount).
	* `overage` - The overage usage.
	* `overages_flag` - The SPM OverageFlag.
	* `platform` - Platform for the cost.
	* `region` - The region of the usage.
	* `resource_id` - The Ocid of the resource that is incurring the cost.
	* `resource_name` - The name of the resource that is incurring the cost.
	* `service` - The name of the service that is incurring the cost.
	* `shape` - The shape of the resource.
	* `sku_name` - The friendly name for the SKU.
	* `sku_part_number` - The part number of the SKU.
	* `subscription_id` - The subscription Id.
	* `tags` - For grouping, a tag definition. For filtering, a definition and key
		* `key` - The key of the tag.
		* `namespace` - The tag namespace.
		* `value` - The value of the tag.
	* `time_usage_ended` - The end time of the usage.
	* `time_usage_started` - The start time of the usage.
	* `unit` - The unit of the usage.
	* `unit_price` - The price per unit.
	* `weight` - The size of resource being metered.

## Import

Import is not supported for this resource.

