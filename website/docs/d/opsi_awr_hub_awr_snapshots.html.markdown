---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub_awr_snapshots"
sidebar_current: "docs-oci-datasource-opsi-awr_hub_awr_snapshots"
description: |-
  Provides the list of Awr Hub Awr Snapshots in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_awr_hub_awr_snapshots
This data source provides the list of Awr Hub Awr Snapshots in Oracle Cloud Infrastructure Opsi service.

Lists AWR snapshots for the specified source database in the AWR hub. The difference between the timeGreaterThanOrEqualTo and timeLessThanOrEqualTo should not exceed an elapsed range of 1 day.
The timeGreaterThanOrEqualTo & timeLessThanOrEqualTo params are optional. If these params are not provided, by default last 1 day snapshots will be returned.


## Example Usage

```hcl
data "oci_opsi_awr_hub_awr_snapshots" "test_awr_hub_awr_snapshots" {
	#Required
	awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
	awr_source_database_identifier = var.awr_hub_awr_snapshot_awr_source_database_identifier

	#Optional
	time_greater_than_or_equal_to = var.awr_hub_awr_snapshot_time_greater_than_or_equal_to
	time_less_than_or_equal_to = var.awr_hub_awr_snapshot_time_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `awr_hub_id` - (Required) Unique Awr Hub identifier
* `awr_source_database_identifier` - (Required) AWR source database identifier.
* `time_greater_than_or_equal_to` - (Optional) The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z 
* `time_less_than_or_equal_to` - (Optional) The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z 


## Attributes Reference

The following attributes are exported:

* `awr_snapshot_collection` - The list of awr_snapshot_collection.

### AwrHubAwrSnapshot Reference

The following attributes are exported:

* `items` - A list of AWR snapshot summary data.
	* `awr_source_database_id` - DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
	* `error_count` - The total number of errors.
	* `instance_number` - The database instance number.
	* `snapshot_identifier` - The identifier of the snapshot.
	* `time_db_startup` - The timestamp of the database startup.
	* `time_snapshot_begin` - The start time of the snapshot.
	* `time_snapshot_end` - The end time of the snapshot.

