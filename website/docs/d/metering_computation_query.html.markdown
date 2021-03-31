---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_query"
sidebar_current: "docs-oci-datasource-metering_computation-query"
description: |-
  Provides details about a specific Query in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_query
This data source provides details about a specific Query resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the saved query.


## Example Usage

```hcl
data "oci_metering_computation_query" "test_query" {
	#Required
	query_id = oci_metering_computation_query.test_query.id
}
```

## Argument Reference

The following arguments are supported:

* `query_id` - (Required) The query unique OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID.
* `id` - The query OCID.
* `query_definition` - The common fields for queries.
	* `cost_analysis_ui` - The common fields for Cost Analysis UI rendering.
		* `graph` - the type of graph mode.
		* `is_cumulative_graph` - is cumulative graph.
	* `display_name` - The query display name. Avoid entering confidential information.
	* `report_query` - the request of generated cost analysis report.
		* `compartment_depth` - The compartment depth level.
		* `date_range_name` - the date range for ui, eg LAST_THREE_MONTHS. It is conflict with timeUsageStarted and timeUsageEnded
		* `filter` - 
		* `forecast` - Forecast configuration of usage/cost.
			* `forecast_type` - BASIC uses ETS to project future usage/cost based on history data. The basis for projections will be a rolling set of equivalent historical days for which projection is being made.
			* `time_forecast_ended` - forecast end time.
			* `time_forecast_started` - forecast start time. Will default to UTC-1 if not specified
		* `granularity` - The usage granularity. HOURLY - Hourly data aggregation. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. TOTAL - Not yet supported. 
		* `group_by` - Aggregate the result by. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - The tag key.
			* `namespace` - The tag namespace.
			* `value` - The tag value.
		* `is_aggregate_by_time` - is aggregated by time. true isAggregateByTime will add up all usage/cost over query time period
		* `query_type` - The query usage type. COST by default if it is missing Usage - Query the usage data. Cost - Query the cost/billing data. 
		* `tenant_id` - Tenant ID
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
	* `version` - the version of saved query.

