---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_schedules"
sidebar_current: "docs-oci-datasource-metering_computation-schedules"
description: |-
  Provides the list of Schedules in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_schedules
This data source provides the list of Schedules in Oracle Cloud Infrastructure Metering Computation service.

Returns the saved schedule list.


## Example Usage

```hcl
data "oci_metering_computation_schedules" "test_schedules" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.schedule_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment ID in which to list resources.
* `name` - (Optional) Query parameter for filtering by name 


## Attributes Reference

The following attributes are exported:

* `schedule_collection` - The list of schedule_collection.

### Schedule Reference

The following attributes are exported:

* `compartment_id` - The customer tenancy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description of the schedule.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}` 
* `id` - The OCID representing a unique shedule.
* `name` - The unique name of the schedule created by the user.
* `output_file_format` - Specifies the supported output file format.
* `query_properties` - The query properties.
	* `compartment_depth` - The depth level of the compartment.
	* `date_range` - Static or dynamic date range `dateRangeType`, which corresponds with type-specific characteristics. 
		* `date_range_type` - Defines whether the schedule date range is STATIC or DYNAMIC.
		* `dynamic_date_range_type` - 
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
	* `filter` - The filter object for query usage.
	* `granularity` - The usage granularity. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation. Allowed values are: DAILY MONTHLY 
	* `group_by` - Aggregate the result by. For example: [ "tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName" ] 
	* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: [ { "namespace": "oracle", "key": "createdBy" ] 
		* `key` - The tag key.
		* `namespace` - The tag namespace.
		* `value` - The tag value.
	* `is_aggregate_by_time` - Specifies whether aggregated by time. If isAggregateByTime is true, all usage or cost over the query time period will be added up.
	* `query_type` - The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data. Allowed values are: USAGE COST USAGE_AND_COST 
* `result_location` - The location where usage or cost CSVs will be uploaded defined by `locationType`, which corresponds with type-specific characteristics. 
	* `bucket` - The bucket name where usage or cost CSVs will be uploaded.
	* `location_type` - Defines the type of location where the usage or cost CSVs will be stored. 
	* `namespace` - The namespace needed to determine the object storage bucket.
	* `region` - The destination Object Store Region specified by the customer.
* `saved_report_id` - The saved report ID which can also be used to generate a query.
* `schedule_recurrences` - Specifies the frequency according to when the schedule will be run,  in the x-obmcs-recurring-time format described in [RFC 5545 section 3.3.10](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). Supported values are : ONE_TIME, DAILY, WEEKLY and MONTHLY. 
* `state` - The schedule lifecycle state.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the schedule was created.
* `time_next_run` - The date and time of the next job execution.
* `time_scheduled` - The date and time of the first time job execution.

