---
subcategory: "Resource Scheduler"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_scheduler_schedules"
sidebar_current: "docs-oci-datasource-resource_scheduler-schedules"
description: |-
  Provides the list of Schedules in Oracle Cloud Infrastructure Resource Scheduler service
---

# Data Source: oci_resource_scheduler_schedules
This data source provides the list of Schedules in Oracle Cloud Infrastructure Resource Scheduler service.

This API gets a list of schedules


## Example Usage

```hcl
data "oci_resource_scheduler_schedules" "test_schedules" {

	#Optional
	compartment_id = var.compartment_id
	schedule_id = oci_resource_scheduler_schedule.test_schedule.id
	display_name = var.schedule_display_name
	state = var.schedule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) This is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. You need to at least provide either `compartment_id` or `schedule_id` or both.
* `schedule_id` - (Optional) This is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.  You need to at least provide either `compartment_id` or `schedule_id` or both.
* `display_name` - (Optional) This is a filter to return only resources that match the given display name exactly.
* `state` - (Optional) This is a filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `schedule_collection` - The list of schedule_collection.

### Schedule Reference

The following attributes are exported:

* `action` - This is the action that will be executed by the schedule.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
* `defined_tags` - These are defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - This is the description of the schedule.
* `display_name` - This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
* `freeform_tags` - These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule
* `recurrence_details` - This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field. 
* `recurrence_type` - Type of recurrence of a schedule
* `resource_filters` - This is a list of resources filters.  The schedule will be applied to resources matching all of them.
	* `attribute` - This is the resource attribute on which the threshold is defined.
	* `condition` - This is the condition for the filter in comparison to its creation time.
	* `should_include_child_compartments` - This sets whether to include child compartments.
	* `value` - This is a collection of resource lifecycle state values.
		* `namespace` - This is the namespace of the defined tag.
		* `tag_key` - This is the key of the defined tag.
		* `value` - This is the value of the defined tag.
* `resources` - This is the list of resources to which the scheduled operation is applied.
	* `id` - This is the resource OCID.
	* `metadata` - This is additional information that helps to identity the resource for the schedule.

		{ "id": "<OCID_of_bucket>" "metadata": { "namespaceName": "sampleNamespace", "bucketName": "sampleBucket" } } 
* `state` - This is the current state of a schedule.
* `system_tags` - These are system tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - This is the date and time the schedule was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_ends` - This is the date and time the schedule ends, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_last_run` - This is the date and time the schedule runs last time, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_next_run` - This is the date and time the schedule run the next time, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_starts` - This is the date and time the schedule starts, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - This is the date and time the schedule was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

