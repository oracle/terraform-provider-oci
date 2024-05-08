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

Creates a data source (DataSource resource), using parameters passed
through a CreateDataSourceDetails resource.


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
		description = var.data_source_data_source_details_description
		interval_in_minutes = var.data_source_data_source_details_interval_in_minutes
		interval_in_seconds = var.data_source_data_source_details_interval_in_seconds
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
		scheduled_query_scope_details {

			#Optional
			region = var.data_source_data_source_details_scheduled_query_scope_details_region
			resource_ids = var.data_source_data_source_details_scheduled_query_scope_details_resource_ids
			resource_type = var.data_source_data_source_details_scheduled_query_scope_details_resource_type
		}
		threshold = var.data_source_data_source_details_threshold
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = var.data_source_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment OCID of the data source
* `data_source_details` - (Optional) (Updatable) Details specific to the data source type.
	* `additional_entities_count` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) The additional entities count used for data source query
	* `data_source_feed_provider` - (Required) (Updatable) Type of data source feed provider (LoggingQuery)
	* `description` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) Description text for the query
	* `interval_in_minutes` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) Interval in minutes that query is run periodically.
	* `interval_in_seconds` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) Interval in minutes which query is run periodically.
	* `logging_query_details` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) Details for a logging query for a data source.
		* `key_entities_count` - (Optional) (Updatable) The key entities count used for data source query
		* `logging_query_type` - (Required) (Updatable) Logging query type for data source
	* `logging_query_type` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) Type of logging query for data source (Sighting/Insight)
	* `operator` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) Operator used in data source
	* `query` - (Optional) (Updatable) The continuous query expression that is run periodically.
	* `query_start_time` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) Start policy for continuous query
		* `query_start_time` - (Applicable when start_policy_type=ABSOLUTE_TIME_START_POLICY) (Updatable) Time when the query can start. If not specified it can start immediately
		* `start_policy_type` - (Required) (Updatable) Start policy delay timing
	* `regions` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) List of logging query regions
	* `scheduled_query_scope_details` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) Target information in which scheduled query will be run
		* `region` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) region on which scheduled query needs to be run
		* `resource_ids` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) List of OCIDs on scheduled query needs to run
		* `resource_type` - (Applicable when data_source_feed_provider=SCHEDULEDQUERY) (Updatable) Type of resource
	* `threshold` - (Applicable when data_source_feed_provider=LOGGINGQUERY) (Updatable) The integer value that must be exceeded, fall below or equal to (depending on the operator), for the query result to trigger an event
* `data_source_feed_provider` - (Required) Type of data source feed provider (LoggingQuery)
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Data source display name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `status` - (Optional) (Updatable) Enablement status of data source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

