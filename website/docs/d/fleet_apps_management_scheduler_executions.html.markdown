---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_executions"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_executions"
description: |-
  Provides the list of Scheduler Executions in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_executions
This data source provides the list of Scheduler Executions in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all executions that are scheduled.


## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_executions" "test_scheduler_executions" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.scheduler_execution_compartment_id_in_subtree
	display_name = var.scheduler_execution_display_name
	lifecycle_operation = var.scheduler_execution_lifecycle_operation
	resource_id = oci_cloud_guard_resource.test_resource.id
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
	runbook_version_name = oci_fleet_apps_management_runbook_version.test_runbook_version.name
	scheduler_defintion_id = oci_fleet_apps_management_scheduler_defintion.test_scheduler_defintion.id
	scheduler_job_id = oci_database_migration_job.test_job.id
	substate = var.scheduler_execution_substate
	time_scheduled_greater_than_or_equal_to = var.scheduler_execution_time_scheduled_greater_than_or_equal_to
	time_scheduled_less_than = var.scheduler_execution_time_scheduled_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `compartment_id_in_subtree` - (Optional) If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `lifecycle_operation` - (Optional) A filter to return only resources their lifecycleOperation matches the given lifecycleOperation.
* `resource_id` - (Optional) ResourceId filter (Example FleetId)
* `runbook_id` - (Optional) A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
* `runbook_version_name` - (Optional) RunbookVersion Name filter
* `scheduler_defintion_id` - (Optional) SchedulerDefinition identifier
* `scheduler_job_id` - (Optional) SchedulerJob identifier filter
* `substate` - (Optional) A filter to return only resources their subState matches the given subState.
* `time_scheduled_greater_than_or_equal_to` - (Optional) Scheduled Time
* `time_scheduled_less_than` - (Optional) Scheduled Time


## Attributes Reference

The following attributes are exported:

* `scheduler_execution_collection` - The list of scheduler_execution_collection.

### SchedulerExecution Reference

The following attributes are exported:

* `items` - List of schedulerExecutions.
	* `activity_id` - Action Group associated with the Schedule.
	* `compartment_id` - Compartment OCID
	* `compartment_name` - Name of the compartment in which resource exist.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the resource.
	* `latest_runbook_version_name` - Latest Runbook version available.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `resource_display_name` - Display Name of the Fleet associated with the Schedule.
	* `resource_id` - FleetId associated with the Schedule.
	* `runbook_display_name` - Display name of Runbook associated with the Schedule.
	* `runbook_id` - RunbookId associated with the Schedule.
	* `runbook_version_name` - Name of the Runbook version associated with the Schedule.
	* `scheduler_definition` - SchedulerDefinition  associated with the job.
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
		* `is_recurring` - Is this a recurring schedule?
	* `scheduler_job_id` - SchedulerJobId associated with the Schedule.
	* `state` - The current state of the Scheduler Execution.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_ended` - Actual end date and time for the Execution.
	* `time_scheduled` - The scheduled date and time for the Job.
	* `time_started` - Actual start date and time for the Execution.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

