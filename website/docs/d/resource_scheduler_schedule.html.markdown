---
subcategory: "Resource Scheduler"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resource_scheduler_schedule"
sidebar_current: "docs-oci-datasource-resource_scheduler-schedule"
description: |-
  Provides details about a specific Schedule in Oracle Cloud Infrastructure Resource Scheduler service
---

# Data Source: oci_resource_scheduler_schedule
This data source provides details about a specific Schedule resource in Oracle Cloud Infrastructure Resource Scheduler service.

This API gets information about a schedule.

## Example Usage

```hcl
data "oci_resource_scheduler_schedule" "test_schedule" {
	#Required
	schedule_id = oci_resource_scheduler_schedule.test_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `schedule_id` - (Required) This is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.


## Attributes Reference

The following attributes are exported:

* `action` - This is the action that will be executed by the schedule.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
* `defined_tags` - These are defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - This is the description of the schedule.
* `display_name` - This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
* `freeform_tags` - These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule
* `last_run_status` - This is the status of the last work request.
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
	* `parameters` - This is the user input parameters to use when acting on the resource.

		{ "parameters": [ { "parameterType": "BODY", "value": { "ip": "192.168.44.44", "memory": "1024", "synced_folders": [ { "host_path": "data/", "guest_path": "/var/www", "type": "default" } ], "forwarded_ports": [] } }, { "parameterType": "PATH", "value": { "compartmentId": "ocid1.compartment.oc1..xxxxx", "instanceId": "ocid1.vcn.oc1..yyyy" } }, { "parameterType": "QUERY", "value": { "limit": "10", "tenantId": "ocid1.tenant.oc1..zzzz" } }, { "parameterType": "HEADER", "value": { "token": "xxxx" } } ] } 
		* `parameter_type` - This is the parameter type on which the input parameter is defined
		* `value` - This is the HTTP request header value.
* `state` - This is the current state of a schedule.
* `system_tags` - These are system tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - This is the date and time the schedule was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_ends` - This is the date and time the schedule ends, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_last_run` - This is the date and time the schedule runs last time, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_next_run` - This is the date and time the schedule run the next time, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_starts` - This is the date and time the schedule starts, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339)  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - This is the date and time the schedule was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

