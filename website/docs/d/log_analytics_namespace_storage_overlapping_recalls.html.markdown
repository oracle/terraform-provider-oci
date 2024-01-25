---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_storage_overlapping_recalls"
sidebar_current: "docs-oci-datasource-log_analytics-namespace_storage_overlapping_recalls"
description: |-
  Provides the list of Namespace Storage Overlapping Recalls in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_namespace_storage_overlapping_recalls
This data source provides the list of Namespace Storage Overlapping Recalls in Oracle Cloud Infrastructure Log Analytics service.

This API gets the list of overlapping recalls made in the given timeframe


## Example Usage

```hcl
data "oci_log_analytics_namespace_storage_overlapping_recalls" "test_namespace_storage_overlapping_recalls" {
	#Required
	namespace = var.namespace_storage_overlapping_recall_namespace

	#Optional
	time_data_ended = var.namespace_storage_overlapping_recall_time_data_ended
	time_data_started = var.namespace_storage_overlapping_recall_time_data_started
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `time_data_ended` - (Optional) This is the end of the time range for recalled data
* `time_data_started` - (Optional) This is the start of the time range for recalled data


## Attributes Reference

The following attributes are exported:

* `overlapping_recall_collection` - The list of overlapping_recall_collection.

### NamespaceStorageOverlappingRecall Reference

The following attributes are exported:

* `items` - This is the array of overlapping recall requests
	* `collection_id` - This is the id of the associated recalled data collection
	* `created_by` - This is the user who initiated the recall request
	* `log_sets` - This is the list of logsets associated with this recall
	* `purpose` - This is the purpose of the recall
	* `query_string` - This is the query associated with the recall
	* `recall_id` - This is the id for the recall request
	* `status` - This is the status of the recall
	* `time_data_ended` - This is the end of the time range of the archival data
	* `time_data_started` - This is the start of the time range of the archival data
	* `time_started` - This is the time when the recall operation was started for this recall request

