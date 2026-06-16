---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_insight_capabilities"
sidebar_current: "docs-oci-datasource-psql-insight_capabilities"
description: |-
  Provides the list of Insight Capabilities in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_insight_capabilities
This data source provides the list of Insight Capabilities in Oracle Cloud Infrastructure Psql service.

Returns the supported insight types and their capabilities.
This API allows clients to discover:
- Supported insight types
- Supported insight data types for each insight type
- Filters, sorting, pagination, limits, and data contracts
required to use the unified insights API.


## Example Usage

```hcl
data "oci_psql_insight_capabilities" "test_insight_capabilities" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `insight_capability_collection` - The list of insight_capability_collection.

### InsightCapability Reference

The following attributes are exported:

* `items` - PostgreSQL insight capabilities list.
	* `data_type_capabilities` - Supported insight data types for this insight type.
		* `data_contract` - Describes the response data format returned for an insight type.
			* `kind` - Indicates the structure of the insight data payload.
			* `unit` - Optional unit associated with numeric values.
		* `date_time_range_support` - Indicates whether a time range is required for the insight.
			* `is_required` - Indicates whether start and end time parameters are required.
		* `description` - Human-readable description of the insight data type.
		* `filters` - Supported filters for this insight data type.
			* `can_use_partial_match` - Indicates whether partial matching is supported.
			* `name` - Name of the filter parameter.
			* `type` - Data type of the filter parameter.
			* `values` - Allowed values for enum-based filters.
		* `granularity` - Describes time granularity behavior for time-series Insight.
			* `max_seconds` - Maximum supported granularity in seconds.
			* `min_seconds` - Minimum supported granularity in seconds.
			* `type` - Granularity selection strategy.
		* `insight_data_type` - Insight data type identifier (for example, AAS_TIME_SERIES).
		* `limits` - Defines limits applicable to an insight type.
			* `max_rows` - Maximum number of rows returned.
			* `max_time_range_days` - Maximum allowed time range in days.
		* `pagination` - Describes pagination support for an insight type.
			* `default_limit` - Default number of items per page.
			* `is_supported` - Indicates whether pagination is supported.
			* `max_limit` - Maximum number of items per page.
		* `sorting` - Describes sorting support for an insight type.
			* `default_sort` - Default sorting behavior for an insight type.
				* `field` - Default field used for sorting.
				* `order` - Default sort order.
			* `fields` - Fields that can be used for sorting.
			* `is_supported` - Indicates whether sorting is supported.
	* `description` - Human-readable description of the insight type.
	* `insight_type` - Echo of the requested insight type.