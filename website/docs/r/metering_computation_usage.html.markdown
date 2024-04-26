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

Returns usage for the given account.


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
	forecast {
		#Required
		time_forecast_ended = var.usage_forecast_time_forecast_ended

		#Optional
		forecast_type = var.usage_forecast_forecast_type
		time_forecast_started = var.usage_forecast_time_forecast_started
	}
	group_by = var.usage_group_by
	group_by_tag {

		#Optional
		key = var.usage_group_by_tag_key
		namespace = var.usage_group_by_tag_namespace
		value = var.usage_group_by_tag_value
	}
	is_aggregate_by_time = var.usage_is_aggregate_by_time
	query_type = var.usage_query_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_depth` - (Optional) The compartment depth level.
* `filter` - (Optional) The filter object for query usage.
* `forecast` - (Optional) Forecast configuration of usage/cost.
	* `forecast_type` - (Optional) BASIC uses the exponential smoothing (ETS) model to project future usage/costs based on history data. The basis for projections is a periodic set of equivalent historical days for which the projection is being made.
	* `time_forecast_ended` - (Required) The forecast end time.
	* `time_forecast_started` - (Optional) The forecast start time. Defaults to UTC-1 if not specified.
* `granularity` - (Required) The usage granularity. HOURLY - Hourly data aggregation. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. TOTAL - Not yet supported. 
* `group_by` - (Optional) Aggregate the result by. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
* `group_by_tag` - (Optional) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
	* `key` - (Optional) The tag key.
	* `namespace` - (Optional) The tag namespace.
	* `value` - (Optional) The tag value.
* `is_aggregate_by_time` - (Optional) Whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
* `query_type` - (Optional) The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data. Credit - Query the credit adjustments data. ExpiredCredit - Query the expired credits data. AllCredit - Query the credit adjustments and expired credit. 
* `tenant_id` - (Required) Tenant ID.
* `time_usage_ended` - (Required) The usage end time.
* `time_usage_started` - (Required) The usage start time.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `group_by` - Aggregate the result by.
* `items` - A list of usage items.
	* `ad` - The availability domain of the usage.
	* `attributed_cost` - The attributed cost with a max value of 9999999999.999999999999 and a minimum value of 0.
	* `attributed_usage` - The attributed usage with a max value of 9999999999.999999999999 and a minimum value of 0.
	* `compartment_id` - The compartment OCID.
	* `compartment_name` - The compartment name.
	* `compartment_path` - The compartment path, starting from root.
	* `computed_amount` - The computed cost.
	* `computed_quantity` - The usage number.
	* `currency` - The price currency.
	* `discount` - The discretionary discount applied to the SKU.
	* `is_forecast` - The forecasted data.
	* `list_rate` - The SKU list rate (not discount).
	* `overage` - The overage usage.
	* `overages_flag` - The SPM OverageFlag.
	* `platform` - Platform for the cost.
	* `region` - The region of the usage.
	* `resource_id` - The resource OCID that is incurring the cost.
	* `resource_name` - The resource name that is incurring the cost.
	* `service` - The service name that is incurring the cost.
	* `shape` - The resource shape.
	* `sku_name` - The SKU friendly name.
	* `sku_part_number` - The SKU part number.
	* `subscription_id` - The subscription ID.
	* `tags` - For grouping, a tag definition. For filtering, a definition and key.
		* `key` - The tag key.
		* `namespace` - The tag namespace.
		* `value` - The tag value.
	* `tenant_id` - The tenancy OCID.
	* `tenant_name` - The tenancy name.
	* `time_usage_ended` - The usage end time.
	* `time_usage_started` - The usage start time.
	* `unit` - The usage unit.
	* `unit_price` - The price per unit.
	* `weight` - The resource size being metered.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Usage
	* `update` - (Defaults to 20 minutes), when updating the Usage
	* `delete` - (Defaults to 20 minutes), when destroying the Usage


## Import

Import is not supported for this resource.

