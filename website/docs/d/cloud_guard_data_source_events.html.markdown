---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_source_events"
sidebar_current: "docs-oci-datasource-cloud_guard-data_source_events"
description: |-
  Provides the list of Data Source Events in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_data_source_events
This data source provides the list of Data Source Events in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of events from CloudGuard DataSource


## Example Usage

```hcl
data "oci_cloud_guard_data_source_events" "test_data_source_events" {
	#Required
	data_source_id = oci_cloud_guard_data_source.test_data_source.id

	#Optional
	region = var.data_source_event_region
}
```

## Argument Reference

The following arguments are supported:

* `data_source_id` - (Required) DataSource OCID
* `region` - (Optional) A filter to return only resource their region matches the given region.


## Attributes Reference

The following attributes are exported:

* `data_source_event_collection` - The list of data_source_event_collection.

### DataSourceEvent Reference

The following attributes are exported:

* `items` - List of event related to a DataSource
	* `comments` - Data source event comments
	* `data_source_id` - Attached data Source
	* `event_date` - Data source event date time
	* `event_info` - Event info of a data source.
		* `data_source_feed_provider` - Possible type of dataSourceFeed Provider(LoggingQuery)
		* `log_result` - 
		* `observed_value` - 
		* `operator` - 
		* `trigger_value` - 
	* `region` - Data source event region
	* `status` - Current data source event info status
	* `time_created` - Data source event created time

