---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_source"
sidebar_current: "docs-oci-datasource-cloud_guard-data_source"
description: |-
  Provides details about a specific Data Source in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_data_source
This data source provides details about a specific Data Source resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a data source (DataSource resource) identified by dataSourceId.

## Example Usage

```hcl
data "oci_cloud_guard_data_source" "test_data_source" {
	#Required
	data_source_id = oci_cloud_guard_data_source.test_data_source.id
}
```

## Argument Reference

The following arguments are supported:

* `data_source_id` - (Required) Data source OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID of data source
* `data_source_details` - Details specific to the data source type.
	* `additional_entities_count` - The additional entities count used for data source query
	* `data_source_feed_provider` - Type of data source feed provider (LoggingQuery)
	* `description` - Description text for the query
	* `interval_in_minutes` - Interval in minutes that query is run periodically.
	* `interval_in_seconds` - Interval in minutes which query is run periodically.
	* `logging_query_details` - Details for a logging query for a data source.
		* `key_entities_count` - The key entities count used for data source query
		* `logging_query_type` - Logging query type for data source
	* `logging_query_type` - Type of logging query for data source (Sighting/Insight)
	* `operator` - Operator used in data source
	* `query` - The continuous query expression that is run periodically.
	* `query_start_time` - Start policy for continuous query
		* `query_start_time` - Time when the query can start. If not specified it can start immediately
		* `start_policy_type` - Start policy delay timing
	* `regions` - List of logging query regions
	* `scheduled_query_scope_details` - Target information in which scheduled query will be run
		* `region` - region on which scheduled query needs to be run
		* `resource_ids` - List of OCIDs on scheduled query needs to run
		* `resource_type` - Type of resource
	* `threshold` - The integer value that must be exceeded, fall below or equal to (depending on the operator), for the query result to trigger an event
* `data_source_detector_mapping_info` - Information about the detector recipe and rule attached
	* `detector_recipe_id` - ID of the detector recipe attached to the data source
	* `detector_rule_id` - ID of the detector rule attached to the data source
* `data_source_feed_provider` - Possible type of dataSourceFeed Provider(LoggingQuery)
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Display name of the data source
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - OCID for the data source
* `region_status_detail` - Information about the region and status of query replication
	* `region` - Data source replication region
	* `status` - Data source replication region status
* `state` - The current lifecycle state of the resource.
* `status` - Enablement status of the data source
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Data source was created. Format defined by RFC3339.
* `time_updated` - The date and time the data source was updated. Format defined by RFC3339.

