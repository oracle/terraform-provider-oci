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

Returns a list of data source events
(DataSourceEventCollection  resource) from the data source
(DataSource resource) identified by dataSourceId.


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

* `data_source_id` - (Required) Data source OCID.
* `region` - (Optional) A filter to return only resource where their region matches the given region.


## Attributes Reference

The following attributes are exported:

* `data_source_event_collection` - The list of data_source_event_collection.

### DataSourceEvent Reference

The following attributes are exported:

* `items` - List of events related to a data source
	* `comments` - Data source event comments
	* `data_source_id` - Unique identifier of data source.
	* `event_date` - Data source event date and time
	* `event_info` - This resource can have multiple subtypes, depending on the dataSourceFeedProvider value. For example, if dataSourceFeedProvider is LOGGINGQUERY, this resource will be of type LoggingEventInfo. 
		* `data_source_feed_provider` - Possible type of dataSourceFeed Provider (LoggingQuery)
		* `log_result` - Log result details of DataSource for a Problem
		* `observed_value` - Observed value of DataSource for a Problem
		* `operator` - Operator details of DataSource for a Problem
		* `trigger_value` - Triggered value of DataSource for a Problem
	* `region` - Data source event region
	* `status` - Current data source event info status
	* `time_created` - Data source event creation date and time
