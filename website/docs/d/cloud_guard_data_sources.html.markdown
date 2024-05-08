---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_data_sources"
sidebar_current: "docs-oci-datasource-cloud_guard-data_sources"
description: |-
  Provides the list of Data Sources in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_data_sources
This data source provides the list of Data Sources in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all data sources (DataSource resources) for a compartment
identified by compartmentId. List is returned in a DataSourceCollection resource
with page of DataSourceSummary resources.

The ListAdhocQueries operation returns only the adhoc queries in 'compartmentId' passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAdhocQueries on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_data_sources" "test_data_sources" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.data_source_access_level
	compartment_id_in_subtree = var.data_source_compartment_id_in_subtree
	data_source_feed_provider = var.data_source_data_source_feed_provider
	display_name = var.data_source_display_name
	logging_query_type = var.data_source_logging_query_type
	state = var.data_source_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`. 
* `data_source_feed_provider` - (Optional) A filter to return only resources when their feed provider matches the given feed provider (`DataSourceFeedProvider` resource).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `logging_query_type` - (Optional) A filter to return only resources where their query type matches the given LoggingQueryType.
* `state` - (Optional) The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `data_source_collection` - The list of data_source_collection.

### DataSource Reference

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

