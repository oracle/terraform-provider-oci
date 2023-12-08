---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_carbon_emissions_query"
sidebar_current: "docs-oci-resource-metering_computation-usage_carbon_emissions_query"
description: |-
  Provides the Usage Carbon Emissions Query resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_usage_carbon_emissions_query
This resource provides the Usage Carbon Emissions Query resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the created usage carbon emissions query.


## Example Usage

```hcl
resource "oci_metering_computation_usage_carbon_emissions_query" "test_usage_carbon_emissions_query" {
	#Required
	compartment_id = var.compartment_id
	query_definition {
		#Required
		cost_analysis_ui {

			#Optional
			graph = var.usage_carbon_emissions_query_query_definition_cost_analysis_ui_graph
			is_cumulative_graph = var.usage_carbon_emissions_query_query_definition_cost_analysis_ui_is_cumulative_graph
		}
		display_name = var.usage_carbon_emissions_query_query_definition_display_name
		report_query {
			#Required
			tenant_id = oci_metering_computation_tenant.test_tenant.id

			#Optional
			compartment_depth = var.usage_carbon_emissions_query_query_definition_report_query_compartment_depth
			date_range_name = var.usage_carbon_emissions_query_query_definition_report_query_date_range_name
			group_by = var.usage_carbon_emissions_query_query_definition_report_query_group_by
			group_by_tag {

				#Optional
				key = var.usage_carbon_emissions_query_query_definition_report_query_group_by_tag_key
				namespace = var.usage_carbon_emissions_query_query_definition_report_query_group_by_tag_namespace
				value = var.usage_carbon_emissions_query_query_definition_report_query_group_by_tag_value
			}
			is_aggregate_by_time = var.usage_carbon_emissions_query_query_definition_report_query_is_aggregate_by_time
			time_usage_ended = var.usage_carbon_emissions_query_query_definition_report_query_time_usage_ended
			time_usage_started = var.usage_carbon_emissions_query_query_definition_report_query_time_usage_started
			usage_carbon_emissions_query_filter = var.usage_carbon_emissions_query_query_definition_report_query_usage_carbon_emissions_query_filter
		}
		version = var.usage_carbon_emissions_query_query_definition_version
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
	* `report_query` - (Required) (Updatable) The request of the generated usage carbon emissions report.
		* `compartment_depth` - (Optional) (Updatable) The compartment depth level.
		* `date_range_name` - (Optional) (Updatable) The UI date range, for example, LAST_THREE_MONTHS. It will override timeUsageStarted and timeUsageEnded properties.
		* `group_by` - (Optional) (Updatable) Specifies what to aggregate the result by. For example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
		* `group_by_tag` - (Optional) (Updatable) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
			* `key` - (Optional) (Updatable) The tag key.
			* `namespace` - (Optional) (Updatable) The tag namespace.
			* `value` - (Optional) (Updatable) The tag value.
		* `is_aggregate_by_time` - (Optional) (Updatable) Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
		* `tenant_id` - (Required) (Updatable) Tenant ID.
		* `time_usage_ended` - (Optional) (Updatable) The usage end time.
		* `time_usage_started` - (Optional) (Updatable) The usage start time.
		* `usage_carbon_emissions_query_filter` - (Optional) (Updatable) The filter object for query usage.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Usage Carbon Emissions Query
	* `update` - (Defaults to 20 minutes), when updating the Usage Carbon Emissions Query
	* `delete` - (Defaults to 20 minutes), when destroying the Usage Carbon Emissions Query


## Import

UsageCarbonEmissionsQueries can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_usage_carbon_emissions_query.test_usage_carbon_emissions_query "id"
```

