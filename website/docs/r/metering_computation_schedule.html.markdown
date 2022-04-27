---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_schedule"
sidebar_current: "docs-oci-resource-metering_computation-schedule"
description: |-
  Provides the Schedule resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_schedule
This resource provides the Schedule resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the created schedule.


## Example Usage

```hcl
resource "oci_metering_computation_schedule" "test_schedule" {
	#Required
	compartment_id = var.compartment_id
	name = var.schedule_name
	query_properties {
		#Required
		date_range {
			#Required
			date_range_type = var.schedule_query_properties_date_range_date_range_type

			#Optional
			dynamic_date_range_type = var.schedule_query_properties_date_range_dynamic_date_range_type
			time_usage_ended = var.schedule_query_properties_date_range_time_usage_ended
			time_usage_started = var.schedule_query_properties_date_range_time_usage_started
		}
		granularity = var.schedule_query_properties_granularity

		#Optional
		compartment_depth = var.schedule_query_properties_compartment_depth
		filter = var.schedule_query_properties_filter
		group_by = var.schedule_query_properties_group_by
		group_by_tag {

			#Optional
			key = var.schedule_query_properties_group_by_tag_key
			namespace = var.schedule_query_properties_group_by_tag_namespace
			value = var.schedule_query_properties_group_by_tag_value
		}
		is_aggregate_by_time = var.schedule_query_properties_is_aggregate_by_time
		query_type = var.schedule_query_properties_query_type
	}
	result_location {
		#Required
		bucket = var.schedule_result_location_bucket
		location_type = var.schedule_result_location_location_type
		namespace = var.schedule_result_location_namespace
		region = var.schedule_result_location_region
	}
	schedule_recurrences = var.schedule_schedule_recurrences
	time_scheduled = var.schedule_time_scheduled

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The tenancy of the customer
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}` 
* `name` - (Required) The unique name of the schedule created by the user
* `query_properties` - (Required) The query properties.
	* `compartment_depth` - (Optional) The depth level of the compartment.
	* `date_range` - (Required) Static or dynamic date range `dateRangeType`, which corresponds with type-specific characteristics. 
		* `date_range_type` - (Required) Defines whether the schedule date range is STATIC or DYNAMIC
		* `dynamic_date_range_type` - (Required when date_range_type=DYNAMIC) 
		* `time_usage_ended` - (Required when date_range_type=STATIC) The usage end time.
		* `time_usage_started` - (Required when date_range_type=STATIC) The usage start time.
	* `filter` - (Optional) 
	* `granularity` - (Required) The usage granularity. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation.   Allowed values are: DAILY MONTHLY 
	* `group_by` - (Optional) Aggregate the result by. For example: [ "tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName" ] 
	* `group_by_tag` - (Optional) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: [ { "namespace": "oracle", "key": "createdBy" ] 
		* `key` - (Optional) The tag key.
		* `namespace` - (Optional) The tag namespace.
		* `value` - (Optional) The tag value.
	* `is_aggregate_by_time` - (Optional) Specifies whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
	* `query_type` - (Optional) The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data.  Allowed values are: USAGE COST USAGE_AND_COST 
* `result_location` - (Required) The location where usage/cost CSVs will be uploaded defined by `locationType`, which corresponds with type-specific characteristics. 
	* `bucket` - (Required) The bucket name where usage/cost CSVs will be uploaded
	* `location_type` - (Required) Defines the type of location where the usage/cost CSVs will be stored 
	* `namespace` - (Required) The namespace needed to determine object storage bucket.
	* `region` - (Required) The destination Object Store Region specified by customer
* `schedule_recurrences` - (Required) In x-obmcs-recurring-time format shown here: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10 Describes the frequency of when the schedule will be run 
* `time_scheduled` - (Required) The date and time of the first time job execution


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The tenancy of the customer
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}` 
* `id` - The OCID representing unique shedule
* `name` - The unique name of the schedule created by the user
* `query_properties` - The query properties.
	* `compartment_depth` - The depth level of the compartment.
	* `date_range` - Static or dynamic date range `dateRangeType`, which corresponds with type-specific characteristics. 
		* `date_range_type` - Defines whether the schedule date range is STATIC or DYNAMIC
		* `dynamic_date_range_type` - 
		* `time_usage_ended` - The usage end time.
		* `time_usage_started` - The usage start time.
	* `filter` - 
	* `granularity` - The usage granularity. DAILY - Daily data aggregation. MONTHLY - Monthly data aggregation.   Allowed values are: DAILY MONTHLY 
	* `group_by` - Aggregate the result by. For example: [ "tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName" ] 
	* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: [ { "namespace": "oracle", "key": "createdBy" ] 
		* `key` - The tag key.
		* `namespace` - The tag namespace.
		* `value` - The tag value.
	* `is_aggregate_by_time` - Specifies whether aggregated by time. If isAggregateByTime is true, all usage/cost over the query time period will be added up.
	* `query_type` - The query usage type. COST by default if it is missing. Usage - Query the usage data. Cost - Query the cost/billing data.  Allowed values are: USAGE COST USAGE_AND_COST 
* `result_location` - The location where usage/cost CSVs will be uploaded defined by `locationType`, which corresponds with type-specific characteristics. 
	* `bucket` - The bucket name where usage/cost CSVs will be uploaded
	* `location_type` - Defines the type of location where the usage/cost CSVs will be stored 
	* `namespace` - The namespace needed to determine object storage bucket.
	* `region` - The destination Object Store Region specified by customer
* `schedule_recurrences` - In x-obmcs-recurring-time format shown here: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10 Describes the frequency of when the schedule will be run 
* `state` - The lifecycle state of the schedule
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time of when the schedule was created
* `time_scheduled` - The date and time of the first time job execution

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Schedule
	* `update` - (Defaults to 20 minutes), when updating the Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Schedule


## Import

Schedules can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_schedule.test_schedule "id"
```

