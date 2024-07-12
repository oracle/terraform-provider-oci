---
subcategory: "Resource Scheduler"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_scheduler_schedule"
sidebar_current: "docs-oci-resource-resource_scheduler-schedule"
description: |-
  Provides the Schedule resource in Oracle Cloud Infrastructure Resource Scheduler service
---

# oci_resource_scheduler_schedule
This resource provides the Schedule resource in Oracle Cloud Infrastructure Resource Scheduler service.

Creates a Schedule


## Example Usage

```hcl
resource "oci_resource_scheduler_schedule" "test_schedule" {
  #Required
  action             = var.schedule_action
  compartment_id     = var.compartment_id
  recurrence_details = var.schedule_recurrence_details
  recurrence_type    = var.schedule_recurrence_type

  resource_filters {
    # Required
    attribute = "DEFINED_TAGS"
    value {
      namespace="SampleNamespace"
      tag_key="SampleTagKey"
      value="SampleValue"
    }
  }
  resource_filters {
    # Required
    attribute = "LIFECYCLE_STATE"
    value {
      value="SampleLifecycleState"
    }
  }
  resource_filters {
    # Required
    attribute = "COMPARTMENT_ID"
    value {
      value=var.compartment_id
    }
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.schedule_defined_tags_value)
  description   = var.schedule_description
  display_name  = var.schedule_display_name
  freeform_tags = var.schedule_freeform_tags
  time_ends   = var.schedule_time_ends
  time_starts = var.schedule_time_starts
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required) (Updatable) This is the action that will be executed by the schedule.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
* `recurrence_type` - (Required) (Updatable) Type of recurrence of a schedule. Could be set to `ICAL`, `CRON`
* `recurrence_details` - (Required) (Updatable) This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field. Example: `FREQ=WEEKLY;BYDAY=MO,TU,WE,TH;BYHOUR=10;INTERVAL=1`

* `resource_filters` - (Optional) (Updatable) This is a list of resources filters.  The schedule will be applied to resources matching all of them.
	* `attribute` - (Required) (Updatable) This is the resource attribute on which the threshold is defined. We support 5 different types of attributes: `DEFINED_TAGS`, `COMPARTMENT_ID`, `TIME_CREATED`, `LIFECYCLE_STATE` and `RESOURCE_TYPE`.
	* `value` - (Optional) (Updatable) This is a collection of resource filter values, different types of filter has different value format, see below:
	* When `attribute="DEFINED_TAGS"`:
		* `namespace` - (Updatable) This is the namespace of the defined tag.
		* `tag_key` - (Updatable) This is the key of the defined tag.
		* `value` - (Updatable) This is the value of the defined tag.
	* When `attribute="TIME_CREATED"`:
		* `condition` - (Applicable when attribute=TIME_CREATED) (Updatable) This is the condition for the filter in comparison to its creation time. Could be set to `EQUAL`, `BEFORE` and `AFTER`.
		* `value` - (Updatable) This is the date and time of resources used for filtering, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z`
	* When `attribute="COMPARTMENT_ID"`:
		* `should_include_child_compartments` - (Updatable) This boolean sets whether to include child compartments.
		* `value` - (Updatable) This is the compartment id used for filtering.
	* When `attribute="RESOURCE_TYPE"`:
		* `value` - (Updatable) This is the resource type used for filtering.
	* when `attribute="LIFECYCLE_STATE"`:
		* `value` - (Updatable) This is the lifecycle state value used for filtering.

* `resources` - (Optional) (Updatable) This is the list of resources to which the scheduled operation is applied.
	* `id` - (Required) (Updatable) This is the resource OCID.
	* `metadata` - (Optional) (Updatable) This is additional information that helps to identity the resource for the schedule.

	  { "id": "<OCID_of_bucket>" "metadata": { "namespaceName": "sampleNamespace", "bucketName": "sampleBucket" } }
* `defined_tags` - (Optional) (Updatable) These are defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) This is the description of the schedule.
* `display_name` - (Optional) (Updatable) This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `time_ends` - (Optional) (Updatable) This is the date and time the schedule ends, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_starts` - (Optional) (Updatable) This is the date and time the schedule starts, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `state` - (Optional) (Updatable) The target state for the Schedule. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
	* `value` - This is a collection of resource filter values.
		* `namespace` - This is the namespace of the defined tag.
		* `tag_key` - This is the key of the defined tag.
		* `value` - This is the value of the designated resource filter type value.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Schedule
	* `update` - (Defaults to 20 minutes), when updating the Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Schedule


## Import

Schedules can be imported using the `id`, e.g.

```
$ terraform import oci_resource_scheduler_schedule.test_schedule "id"
```

