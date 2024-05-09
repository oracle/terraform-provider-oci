---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_scheduled_jobs"
sidebar_current: "docs-oci-datasource-os_management_hub-scheduled_jobs"
description: |-
  Provides the list of Scheduled Jobs in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_scheduled_jobs
This data source provides the list of Scheduled Jobs in Oracle Cloud Infrastructure Os Management Hub service.

Lists scheduled jobs that match the specified compartment or scheduled job [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
data "oci_os_management_hub_scheduled_jobs" "test_scheduled_jobs" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.scheduled_job_compartment_id_in_subtree
	display_name = var.scheduled_job_display_name
	display_name_contains = var.scheduled_job_display_name_contains
	id = var.scheduled_job_id
	is_managed_by_autonomous_linux = var.scheduled_job_is_managed_by_autonomous_linux
	is_restricted = var.scheduled_job_is_restricted
	lifecycle_stage_id = oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id
	location = var.scheduled_job_location
	location_not_equal_to = var.scheduled_job_location_not_equal_to
	managed_compartment_id = oci_identity_compartment.test_compartment.id
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	operation_type = var.scheduled_job_operation_type
	schedule_type = var.scheduled_job_schedule_type
	state = var.scheduled_job_state
	time_end = var.scheduled_job_time_end
	time_start = var.scheduled_job_time_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `compartment_id_in_subtree` - (Optional) Indicates whether to include subcompartments in the returned results. Default is false.
* `display_name` - (Optional) A filter to return resources that match the given user-friendly name.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled job. A filter to return the specified job.
* `is_managed_by_autonomous_linux` - (Optional) Indicates whether to list only resources managed by the Autonomous Linux service. 
* `is_restricted` - (Optional) A filter to return only restricted scheduled jobs.
* `lifecycle_stage_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the lifecycle stage. This resource returns resources associated with this lifecycle stage.
* `location` - (Optional) A filter to return only resources whose location matches the given value.
* `location_not_equal_to` - (Optional) A filter to return only resources whose location does not match the given value.
* `managed_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed compartment. This filter returns resources associated with this compartment.
* `managed_instance_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group. This filter returns resources associated with this group.
* `managed_instance_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance. This filter returns resources associated with this managed instance.
* `operation_type` - (Optional) A filter to return only scheduled jobs with the given operation type.
* `schedule_type` - (Optional) A filter to return only scheduled jobs of the given scheduling type (one-time or recurring).
* `state` - (Optional) A filter to return only scheduled jobs currently in the given state.
* `time_end` - (Optional) A filter to return only resources with a date on or before the given value, in ISO 8601 format.  Example: 2017-07-14T02:40:00.000Z 
* `time_start` - (Optional) A filter to return only resources with a date on or after the given value, in ISO 8601 format.  Example: 2017-07-14T02:40:00.000Z 


## Attributes Reference

The following attributes are exported:

* `scheduled_job_collection` - The list of scheduled_job_collection.

### ScheduledJob Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the scheduled job.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the scheduled job.
* `display_name` - User-friendly name for the scheduled job.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled job.
* `is_managed_by_autonomous_linux` - Indicates whether this scheduled job is managed by the Autonomous Linux service.
* `is_restricted` - Indicates if the schedule job has restricted update and deletion capabilities. For restricted scheduled jobs,  you can update only the timeNextExecution, recurringRule, and tags. 
* `is_subcompartment_included` - Indicates whether to apply the scheduled job to all compartments in the tenancy when managedCompartmentIds specifies the tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) (root compartment). 
* `lifecycle_stage_ids` - The lifecycle stage [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.  A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with  managedInstanceIds, managedInstanceGroupIds, and managedCompartmentIds. 
* `locations` - The list of locations this scheduled job should operate on for a job targeting on compartments. (Empty list means apply to all locations). This can only be set when managedCompartmentIds is not empty.
* `managed_compartment_ids` - The compartment [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on. A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with managedInstanceIds, managedInstanceGroupIds, and lifecycleStageIds.
* `managed_instance_group_ids` - The managed instance group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on. A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with managedInstanceIds, managedCompartmentIds, and lifecycleStageIds.
* `managed_instance_ids` - The managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.  A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with  managedInstanceGroupIds, managedCompartmentIds, and lifecycleStageIds. 
* `operations` - The list of operations this scheduled job needs to perform. A scheduled job supports only one operation type, unless it is one of the following:
	* UPDATE_PACKAGES
	* UPDATE_ALL
	* UPDATE_SECURITY
	* UPDATE_BUGFIX
	* UPDATE_ENHANCEMENT
	* UPDATE_OTHER
	* UPDATE_KSPLICE_USERSPACE
	* UPDATE_KSPLICE_KERNEL 
	* `manage_module_streams_details` - The set of changes to make to the state of the modules, streams, and profiles on the managed target.
		* `disable` - The set of module streams to disable.
			* `module_name` - The name of a module.
			* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
			* `stream_name` - The name of a stream of the specified module.
		* `enable` - The set of module streams to enable.
			* `module_name` - The name of a module.
			* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
			* `stream_name` - The name of a stream of the specified module.
		* `install` - The set of module stream profiles to install.
			* `module_name` - The name of a module.
			* `profile_name` - The name of a profile of the specified module stream.
			* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
			* `stream_name` - The name of a stream of the specified module.
		* `remove` - The set of module stream profiles to remove.
			* `module_name` - The name of a module.
			* `profile_name` - The name of a profile of the specified module stream.
			* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
			* `stream_name` - The name of a stream of the specified module.
	* `operation_type` - The type of operation this scheduled job performs.
	* `package_names` - The names of the target packages. This parameter only applies when the scheduled job is for installing, updating, or removing packages.
	* `software_source_ids` - The software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).  This parameter only applies when the scheduled job is for attaching or detaching software sources. 
	* `switch_module_streams_details` - Provides the information used to update a module stream.
		* `module_name` - The name of a module.
		* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that contains the module stream. 
		* `stream_name` - The name of a stream of the specified module.
	* `windows_update_names` - Unique identifier for the Windows update. This parameter only applies if the scheduled job is for installing Windows updates. Note that this is not an OCID, but is a unique identifier assigned by Microsoft. For example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'. 
* `recurring_rule` - The frequency schedule for a recurring scheduled job.
* `retry_intervals` - The amount of time in minutes to wait until retrying the scheduled job. If set, the service will automatically retry  a failed scheduled job after the interval. For example, you could set the interval to [2,5,10]. If the initial  execution of the job fails, the service waits 2 minutes and then retries. If that fails, the service waits 5 minutes  and then retries. If that fails, the service waits 10 minutes and then retries. 
* `schedule_type` - The type of scheduling frequency for the job.
* `state` - The current state of the scheduled job.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this scheduled job was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_last_execution` - The time of the last execution of this scheduled job (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_next_execution` - The time of the next execution of this scheduled job (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `time_updated` - The time this scheduled job was updated (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).
* `work_request_ids` - The list of work request [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this scheduled job.

