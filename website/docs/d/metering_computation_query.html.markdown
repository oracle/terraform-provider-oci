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
		* `graph` - The graph type.
		* `is_cumulative_graph` - A cumulative graph.
	* `display_name` - The query display name. Avoid entering confidential information.
	* `report_query` - The request of the generated Cost Analysis report.
		* `compartment_depth` - The compartment depth level.
		* `date_range_name` - The UI date range, for example, LAST_THREE_MONTHS. Conflicts with timeUsageStarted and timeUsageEnded.
		* `filter` - The filter object for query usage.
		* `forecast` - Forecast configuration of usage/cost.
			* `forecast_type` - BASIC uses the exponential smoothing (ETS) model to project future usage/costs based on history data. The basis for projections is a periodic set of equivalent historical days for which the projection is being made
			* `time_forecast_ended` - The forecast end time.
			* `time_forecast_started` - The forecast start time. Defaults to UTC-1 if not specified. 
		* `granularity` - The usage granularity. HOURLY - Hourly data aggregation. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. TOTAL - Not yet supported. 
		* `group_by` - Specifies what to aggregate the result by. For example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - The tag key.
			* `namespace` - The tag namespace.
			* `value` - The tag value.
		* `is_aggregate_by_time` - Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
		* `query_type` - The query usage type. COST by default if it is missing.
			* Usage: Query the usage data.
			* Cost: Query the cost/billing data.
			* Credit: Query the credit adjustments data.
			* ExpiredCredit: Query the expired credits data.
			* AllCredit: Query the credit adjustments and expired credit. 
		* `query_filter` - The filter object for query usage.
		* `tenant_id` - Tenant ID.
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
	* `version` - The saved query version.

