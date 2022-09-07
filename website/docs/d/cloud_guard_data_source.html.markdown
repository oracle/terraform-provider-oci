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

Returns a DataSource identified by dataSourceId

## Example Usage

```hcl
data "oci_cloud_guard_data_source" "test_data_source" {
	#Required
	data_source_id = oci_cloud_guard_data_source.test_data_source.id
}
```

## Argument Reference

The following arguments are supported:

* `data_source_id` - (Required) DataSource OCID


## Attributes Reference

The following attributes are exported:

* `compartment_id` - CompartmentId of Data source.
* `data_source_details` - Details specific to the data source type.
	* `additional_entities_count` - The additional entities count used for data source query.
	* `data_source_feed_provider` - Possible type of dataSourceFeed Provider(LoggingQuery)
	* `interval_in_minutes` - Interval in minutes that query is run periodically.
	* `logging_query_details` - Additional details specific to the data source type (Sighting/Insight).
		* `key_entities_count` - The key entities count used for data source query
		* `logging_query_type` - Logging query type for data source (Sighting/Insight)
	* `logging_query_type` - Logging query type for data source (Sighting/Insight)
	* `operator` - Operator used in Data Soruce
	* `query` - The continuous query expression that is run periodically.
	* `query_start_time` - Continuous query start policy object
		* `query_start_time` - Time when the query can start, if not specified it can start immediately.
		* `start_policy_type` - policy used for deciding the query start time
	* `regions` - Logging Query regions
	* `threshold` - The integer value that must be exceeded, fall below or equal to (depending on the operator), the query result to trigger an event.
* `data_source_detector_mapping_info` - Information about the detector recipe and rule attached
	* `detector_recipe_id` - Id of the attached detectorRecipeId to the Data Source.
	* `detector_rule_id` - Id of the attached detectorRuleId to the Data Source.
* `data_source_feed_provider` - Possible type of dataSourceFeed Provider(LoggingQuery)
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - DisplayName of Data source.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Ocid for Data source
* `region_status_detail` - Information about the region and status of query replication
	* `region` - Data Source replication region.
	* `status` - Data Source replication region status.
* `state` - The current state of the resource.
* `status` - Status of data Source
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Data source was created. Format defined by RFC3339.
* `time_updated` - The date and time the Data source was updated. Format defined by RFC3339.

