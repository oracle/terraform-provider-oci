---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_collection_log_insights"
sidebar_current: "docs-oci-datasource-data_safe-sql_collection_log_insights"
description: |-
  Provides the list of Sql Collection Log Insights in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_collection_log_insights
This data source provides the list of Sql Collection Log Insights in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of the SQL collection log analytics.


## Example Usage

```hcl
data "oci_data_safe_sql_collection_log_insights" "test_sql_collection_log_insights" {
	#Required
	sql_collection_id = oci_data_safe_sql_collection.test_sql_collection.id
	time_ended = var.sql_collection_log_insight_time_ended
	time_started = var.sql_collection_log_insight_time_started

	#Optional
	group_by = var.sql_collection_log_insight_group_by
}
```

## Argument Reference

The following arguments are supported:

* `group_by` - (Optional) The group by parameter to summarize SQL collection log insights aggregation.
* `sql_collection_id` - (Required) The OCID of the SQL collection resource.
* `time_ended` - (Required) An optional filter to return the stats of the SQL collection logs collected before the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_started` - (Required) An optional filter to return the stats of the SQL collection logs collected after the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 


## Attributes Reference

The following attributes are exported:

* `sql_collection_log_insights_collection` - The list of sql_collection_log_insights_collection.

### SqlCollectionLogInsight Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `dimensions` - The dimensions available for SQL collection analytics.
		* `client_ip` - The IP addresses for the SQL collection.
		* `client_os_user_name` - The operating system user names for the SQL collection.
		* `client_program` - The allowed client programs for the SQL collection.
	* `metric_name` - Name of the aggregation.
	* `sql_collection_log_insight_count` - Total count of aggregated value.
	* `time_ended` - The time at which the aggregation ended.
	* `time_started` - The time at which the aggregation started.

