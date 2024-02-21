---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database_sql_data"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database_sql_data"
description: |-
  Provides the list of Managed My Sql Database Sql Data in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database_sql_data
This data source provides the list of Managed My Sql Database Sql Data in Oracle Cloud Infrastructure Database Management service.

Retrieves SQL performance data for given MySQL Instance.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database_sql_data" "test_managed_my_sql_database_sql_data" {
	#Required
	end_time = var.managed_my_sql_database_sql_data_end_time
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
	start_time = var.managed_my_sql_database_sql_data_start_time

	#Optional
	filter_column = var.managed_my_sql_database_sql_data_filter_column
}
```

## Argument Reference

The following arguments are supported:

* `end_time` - (Required) The end time of the time range to retrieve the health metrics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
* `filter_column` - (Optional) The parameter to filter results by key criteria which include :
	* AVG_TIMER_WAIT
	* SUM_TIMER_WAIT
	* COUNT_STAR
	* SUM_ERRORS
	* SUM_ROWS_AFFECTED
	* SUM_ROWS_SENT
	* SUM_ROWS_EXAMINED
	* SUM_CREATED_TMP_TABLES
	* SUM_NO_INDEX_USED
	* SUM_NO_GOOD_INDEX_USED
	* FIRST_SEEN
	* LAST_SEEN
	* HEATWAVE_OFFLOADED
	* HEATWAVE_OUT_OF_MEMORY 
* `managed_my_sql_database_id` - (Required) The OCID of the Managed MySQL Database.
* `start_time` - (Required) The start time of the time range to retrieve the health metrics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 


## Attributes Reference

The following attributes are exported:

* `my_sql_data_collection` - The list of my_sql_data_collection.

### ManagedMySqlDatabaseSqlData Reference

The following attributes are exported:

* `items` - The list of SQLDataSummary records.
	* `avg_timer_wait` - The average execution time.
	* `count_star` - The number Of times the query has been executed.
	* `digest` - The digest information of the normalized query.
	* `digest_text` - The normalized query.
	* `first_seen` - The date and time the query was first seen. If the table is truncated, the first seen value is reset.
	* `heat_wave_offloaded` - The number of query executions offloaded to HeatWave.
	* `heat_wave_out_of_memory` - The number of query executions with HeatWave out-of-memory errors.
	* `last_seen` - The date and time the query was last seen.
	* `max_timer_wait` - The slowest the query has been executed.
	* `min_timer_wait` - The fastest the query has been executed.
	* `quantile95` - The 95th percentile of the query latency. That is, 95% of the queries complete in the time given or in less time.
	* `quantile99` - The 99th percentile of the query latency.
	* `quantile999` - The 99.9th percentile of the query latency.
	* `schema_name` - The name of the default schema when executing the query. If a schema is not set as the default, then the value is NULL.
	* `sum_created_temp_disk_tables` - The total number of On-Disk internal temporary tables that have been created by the query.
	* `sum_created_temp_tables` - The total number of internal temporary tables (in memory or on disk), which have been created by the query.
	* `sum_errors` - The total number of errors that have been encountered executing the query. 
	* `sum_lock_time` - The total amount of time that has been spent waiting for table locks.
	* `sum_no_good_index_used` - The total number of times no good index was used. This means that the extra column in The EXPLAIN output includes “Range Checked For Each Record.”
	* `sum_no_index_used` - The total number of times no index was used to execute the query.
	* `sum_rows_affected` - The total number of rows that have been modified by the query.
	* `sum_rows_examined` - The total number of rows that have been examined by the query.
	* `sum_rows_sent` - The total number of rows that have been returned (sent) to the client.
	* `sum_select_full_join` - The total number of joins that have performed full table scans as there was no join condition or no index for the join condition. This is the same as the select_full_join status variable.
	* `sum_select_full_range_join` - The total number of joins that use a full range search. This is the same as the select_full_range_join status variable.
	* `sum_select_range` - The total number of times the query has used a range search. This is the same as the select_range status variable.
	* `sum_select_range_check` - The total number of joins by the query where the join does not have an index that checks for the index usage after each row. This is the same as the select_range_check status variable.
	* `sum_select_scan` - The total number of times the query has performed a full table scan on the first table in the join. This is the same as the select_scan status variable.
	* `sum_sort_merge_passes` - The total number of sort merge passes that have been done to sort the result of the query. This is the same as the sort_merge_passes status variable.
	* `sum_sort_range` - The total number of times a sort was done using ranges. This is the same as the sort_range status variable.
	* `sum_sort_rows` - The total number of rows sorted. This is the same as the sort_rowsStatus variable.
	* `sum_sort_scan` - The total number of times a sort was done by scanning the table. This is the same as the sort_scan status variable.
	* `sum_timer_wait` - The total amount of time that has been spent executing the query.
	* `sum_warnings` - The total number of warnings that have been encountered executing the query.

