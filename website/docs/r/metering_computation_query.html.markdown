---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_query"
sidebar_current: "docs-oci-resource-metering_computation-query"
description: |-
  Provides the Query resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_query
This resource provides the Query resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the created query.


## Example Usage

```hcl
resource "oci_metering_computation_query" "test_query" {
	#Required
	compartment_id = var.compartment_id
	query_definition {
		#Required
		cost_analysis_ui {

			#Optional
			graph = var.query_query_definition_cost_analysis_ui_graph
			is_cumulative_graph = var.query_query_definition_cost_analysis_ui_is_cumulative_graph
		}
		display_name = var.query_query_definition_display_name
		report_query {
			#Required
			granularity = var.query_query_definition_report_query_granularity
			tenant_id = oci_metering_computation_tenant.test_tenant.id

			#Optional
			compartment_depth = var.query_query_definition_report_query_compartment_depth
			date_range_name = var.query_query_definition_report_query_date_range_name
			filter = var.query_query_definition_report_query_filter
			forecast {
				#Required
				time_forecast_ended = var.query_query_definition_report_query_forecast_time_forecast_ended

				#Optional
				forecast_type = var.query_query_definition_report_query_forecast_forecast_type
				time_forecast_started = var.query_query_definition_report_query_forecast_time_forecast_started
			}
			group_by = var.query_query_definition_report_query_group_by
			group_by_tag {

				#Optional
				key = var.query_query_definition_report_query_group_by_tag_key
				namespace = var.query_query_definition_report_query_group_by_tag_namespace
				value = var.query_query_definition_report_query_group_by_tag_value
			}
			is_aggregate_by_time = var.query_query_definition_report_query_is_aggregate_by_time
			query_type = var.query_query_definition_report_query_query_type
			time_usage_ended = var.query_query_definition_report_query_time_usage_ended
			time_usage_started = var.query_query_definition_report_query_time_usage_started
		}
		version = var.query_query_definition_version
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `query_definition` - (Required) (Updatable) The common fields for queries.
	* `cost_analysis_ui` - (Required) (Updatable) The common fields for Cost Analysis UI rendering.
		* `graph` - (Optional) (Updatable) The graph type.
		* `is_cumulative_graph` - (Optional) (Updatable) A cumulative graph.
	* `display_name` - (Required) (Updatable) The query display name. Avoid entering confidential information.
	* `report_query` - (Required) (Updatable) The request of the generated Cost Analysis report.
		* `compartment_depth` - (Optional) (Updatable) The compartment depth level.
		* `date_range_name` - (Optional) (Updatable) The UI date range, for example, LAST_THREE_MONTHS. Conflicts with timeUsageStarted and timeUsageEnded.
		* `filter` - (Optional) (Updatable) 
		* `forecast` - (Optional) (Updatable) Forecast configuration of usage/cost.
			* `forecast_type` - (Optional) (Updatable) BASIC uses the exponential smoothing (ETS) model to project future usage/costs based on history data. The basis for projections is a periodic set of equivalent historical days for which the projection is being made.
			* `time_forecast_ended` - (Required) (Updatable) The forecast end time.
			* `time_forecast_started` - (Optional) (Updatable) The forecast start time. Defaults to UTC-1 if not specified.
		* `granularity` - (Required) (Updatable) The usage granularity. HOURLY - Hourly data aggregation. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. TOTAL - Not yet supported. 
		* `group_by` - (Optional) (Updatable) Aggregate the result by. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - (Optional) (Updatable) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - (Optional) (Updatable) The tag key.
			* `namespace` - (Optional) (Updatable) The tag namespace.
			* `value` - (Optional) (Updatable) The tag value.
		* `is_aggregate_by_time` - (Optional) (Updatable) Whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
		* `query_type` - (Optional) (Updatable) The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data. Credit - Query the credit adjustments data. ExpiredCredit - Query the expired credits data AllCredit - Query the credit adjustments and expired credit 
		* `tenant_id` - (Required) (Updatable) Tenant ID.
		* `time_usage_ended` - (Optional) (Updatable) The usage end time.
		* `time_usage_started` - (Optional) (Updatable) The usage start time.
	* `version` - (Required) (Updatable) The saved query version.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
		* `filter` - 
		* `forecast` - Forecast configuration of usage/cost.
			* `forecast_type` - BASIC uses the exponential smoothing (ETS) model to project future usage/costs based on history data. The basis for projections is a periodic set of equivalent historical days for which the projection is being made.
			* `time_forecast_ended` - The forecast end time.
			* `time_forecast_started` - The forecast start time. Defaults to UTC-1 if not specified.
		* `granularity` - The usage granularity. HOURLY - Hourly data aggregation. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. TOTAL - Not yet supported. 
		* `group_by` - Aggregate the result by. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - The tag key.
			* `namespace` - The tag namespace.
			* `value` - The tag value.
		* `is_aggregate_by_time` - Whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
		* `query_type` - The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data. Credit - Query the credit adjustments data. ExpiredCredit - Query the expired credits data AllCredit - Query the credit adjustments and expired credit 
		* `tenant_id` - Tenant ID.
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
	* `version` - The saved query version.

## Import

Queries can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_query.test_query "id"
```

