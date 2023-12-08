---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_carbon_emissions_query"
sidebar_current: "docs-oci-datasource-metering_computation-usage_carbon_emissions_query"
description: |-
  Provides details about a specific Usage Carbon Emissions Query in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_usage_carbon_emissions_query
This data source provides details about a specific Usage Carbon Emissions Query resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the usage carbon emissions saved query.


## Example Usage

```hcl
data "oci_metering_computation_usage_carbon_emissions_query" "test_usage_carbon_emissions_query" {
	#Required
	usage_carbon_emissions_query_id = oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query.id
}
```

## Argument Reference

The following arguments are supported:

* `usage_carbon_emissions_query_id` - (Required) The query unique OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID.
* `id` - The query OCID.
* `query_definition` - The common fields for queries.
	* `cost_analysis_ui` - The common fields for Cost Analysis UI rendering.
		* `graph` - The graph type.
		* `is_cumulative_graph` - A cumulative graph.
	* `display_name` - The query display name. Avoid entering confidential information.
	* `report_query` - The request of the generated usage carbon emissions report.
		* `compartment_depth` - The compartment depth level.
		* `date_range_name` - The UI date range, for example, LAST_THREE_MONTHS. It will override timeUsageStarted and timeUsageEnded properties.
		* `group_by` - Specifies what to aggregate the result by. For example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - The tag key.
			* `namespace` - The tag namespace.
			* `value` - The tag value.
		* `is_aggregate_by_time` - Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
		* `tenant_id` - Tenant ID.
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
		* `usage_carbon_emissions_query_filter` - The filter object for query usage.
	* `version` - The saved query version.

