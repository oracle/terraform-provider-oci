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
* `filter_column` - (Optional) The parameter to filter results by key criteria.
* `managed_my_sql_database_id` - (Required) The OCID of ManagedMySqlDatabase.
* `start_time` - (Required) The start time of the time range to retrieve the health metrics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 


## Attributes Reference

The following attributes are exported:

* `my_sql_data_collection` - The list of my_sql_data_collection.

### ManagedMySqlDatabaseSqlData Reference

The following attributes are exported:

* `items` - List of SQLDataSummary.
	* `avg_timer_wait` - The Average Execution Time.
	* `count_star` - The Number Of Times The Query Has Been Executed.
	* `digest` - The Digest Of The Normalized Query.
	* `digest_text` - The Normalized Query.
	* `first_seen` - When The Query Was First Seen. When The Table Is Truncated, The First Seen Value Is Also Reset.
	* `last_seen` - When The Query Was Seen The Last Time.
	* `max_timer_wait` - The Slowest The Query Has Been Executed.
	* `min_timer_wait` - The Fastest The Query Has Been Executed.
	* `quantile95` - The 95th Percentile Of The Query Latency. That Is, 95% Of The Queries Complete In The Time Given Or In Less Time.
	* `quantile99` - The 99th Percentile Of The Query Latency.
	* `quantile999` - The 99.9th Percentile Of The Query Latency.
	* `schema_name` - The Schema That Was The Default Schema When Executing The Query. If No Schema Was The Default, The Value Is NULL.
	* `sum_created_temp_disk_tables` - The Total Number Of On-Disk Internal Temporary Tables That Have Been Created By The Query.
	* `sum_created_temp_tables` - The Total Number Of Internal Temporary Tables – Whether Created In Memory Or On Disk – That Have Been Created By The Query.
	* `sum_errors` - The Total Number Of Errors That Have Been Encountered Executing The Query. 
	* `sum_lock_time` - The Total Amount Of Time That Has Been Spent Waiting For Table Locks.
	* `sum_no_good_index_used` - The Total Number Of Times No Good Index Was Used. This Means That The ExtraColumn In The EXPLAIN Output Includes “Range Checked For Each Record.”
	* `sum_no_index_used` - The Total Number Of Times No Index Was Used To Execute The Query.
	* `sum_rows_affected` - The Total Number Of Rows That Have Been Modified By The Query.
	* `sum_rows_examined` - The Total Number Of Rows That Have Been Examined By The Query.
	* `sum_rows_sent` - The Total Number Of Rows That Have Been Returned (Sent) To The Client.
	* `sum_select_full_join` - The Total Number Of Joins That Have Performed Full Table Scans As There Is No Index For The Join Condition Or There Is No Join Condition. This Is The Same That Increments The Select_full_join Status Variable.
	* `sum_select_full_range_join` - The Total Number Of Joins That Use A Full Range Search. This Is The Same That Increments The Select_full_range_join Status Variable.
	* `sum_select_range` - The Total Number Of Times The Query Has Used A Range Search. This Is The Same That Increments The Select_range Status Variable.
	* `sum_select_range_check` - The Total Number Of Joins By The Query Where The Join Does Not Have An Index That Checks For The Index Usage After Each Row. This Is The Same That Increments The Select_range_check Status Variable.
	* `sum_select_scan` - The Total Number Of Times The Query Has Performed A Full Table Scan On The First Table In The Join. This Is The Same That Increments The Select_scan Status Variable.
	* `sum_sort_merge_passes` - The Total Number Of Sort Merge Passes That Have Been Done To Sort The Result Of The Query. This Is The Same That Increments The Sort_merge_passes Status Variable.
	* `sum_sort_range` - The Total Number Of Times A Sort Was Done Using Ranges. This Is The Same That Increments The Sort_range Status Variable.
	* `sum_sort_rows` - The Total Number Of Rows Sorted. This Is The Same That Increments The Sort_rowsStatus Variable.
	* `sum_sort_scan` - The Total Number Of Times A Sort Was Done By Scanning The Table. This Is The Same That Increments The Sort_scan Status Variable.
	* `sum_timer_wait` - The Total Amount Of Time That Has Been Spent Executing The Query.
	* `sum_warnings` - The Total Number Of Warnings That Have Been Encountered Executing The Query. 

