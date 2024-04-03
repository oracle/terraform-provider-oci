---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_source"
sidebar_current: "docs-oci-resource-cloud_guard-data_source"
description: |-
  Provides the Data Source resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_data_source
This resource provides the Data Source resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a DataSource


## Example Usage

```hcl
resource "oci_cloud_guard_data_source" "test_data_source" {
	#Required
	compartment_id = var.compartment_id
	data_source_feed_provider = var.data_source_data_source_feed_provider
	display_name = var.data_source_display_name

	#Optional
	data_source_details {
		#Required
		data_source_feed_provider = var.data_source_data_source_details_data_source_feed_provider

		#Optional
		additional_entities_count = var.data_source_data_source_details_additional_entities_count
		interval_in_minutes = var.data_source_data_source_details_interval_in_minutes
		logging_query_details {
			#Required
			logging_query_type = var.data_source_data_source_details_logging_query_details_logging_query_type

			#Optional
			key_entities_count = var.data_source_data_source_details_logging_query_details_key_entities_count
		}
		logging_query_type = var.data_source_data_source_details_logging_query_type
		operator = var.data_source_data_source_details_operator
		query = var.data_source_data_source_details_query
		query_start_time {
			#Required
			start_policy_type = var.data_source_data_source_details_query_start_time_start_policy_type

			#Optional
			query_start_time = var.data_source_data_source_details_query_start_time_query_start_time
		}
		regions = var.data_source_data_source_details_regions
		threshold = var.data_source_data_source_details_threshold
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = var.data_source_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) CompartmentId of Data Source.
* `data_source_details` - (Optional) (Updatable) Details specific to the data source type.
	* `additional_entities_count` - (Optional) (Updatable) The additional entities count used for data source query.
	* `data_source_feed_provider` - (Required) (Updatable) Possible type of dataSourceFeed Provider(LoggingQuery)
	* `interval_in_minutes` - (Optional) (Updatable) Interval in minutes that query is run periodically.
	* `logging_query_details` - (Optional) (Updatable) Additional details specific to the data source type (Sighting/Insight).
		* `key_entities_count` - (Optional) (Updatable) The key entities count used for data source query
		* `logging_query_type` - (Required) (Updatable) Logging query type for data source (Sighting/Insight)
	* `logging_query_type` - (Optional) (Updatable) Logging query type for data source (Sighting/Insight)
	* `operator` - (Optional) (Updatable) Operator used in Data Soruce
	* `query` - (Optional) (Updatable) The continuous query expression that is run periodically.
	* `query_start_time` - (Optional) (Updatable) Continuous query start policy object
		* `query_start_time` - (Applicable when start_policy_type=ABSOLUTE_TIME_START_POLICY) (Updatable) Time when the query can start, if not specified it can start immediately.
		* `start_policy_type` - (Required) (Updatable) policy used for deciding the query start time
	* `regions` - (Optional) (Updatable) Logging Query regions
	* `threshold` - (Optional) (Updatable) The integer value that must be exceeded, fall below or equal to (depending on the operator), the query result to trigger an event.
* `data_source_feed_provider` - (Required) Possible type of dataSourceFeed Provider(LoggingQuery)
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Data Source display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `status` - (Optional) (Updatable) Status of DataSource. Default value is DISABLED.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Data Source
	* `update` - (Defaults to 20 minutes), when updating the Data Source
	* `delete` - (Defaults to 20 minutes), when destroying the Data Source


## Import

DataSources can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_data_source.test_data_source "id"
```

