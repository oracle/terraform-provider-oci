---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_job_job_activity_resources"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_job_job_activity_resources"
description: |-
  Provides the list of Scheduler Job Job Activity Resources in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_job_job_activity_resources
This data source provides the list of Scheduler Job Job Activity Resources in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of resources for an Activity Execution.

## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_job_job_activity_resources" "test_scheduler_job_job_activity_resources" {
	#Required
	job_activity_id = oci_fleet_apps_management_job_activity.test_job_activity.id
	scheduler_job_id = oci_database_migration_job.test_job.id

	#Optional
	resource_task_id = oci_fleet_apps_management_resource_task.test_resource_task.id
	sequence = var.scheduler_job_job_activity_resource_sequence
	step_name = var.scheduler_job_job_activity_resource_step_name
	target_name = oci_cloud_guard_target.test_target.name
}
```

## Argument Reference

The following arguments are supported:

* `job_activity_id` - (Required) unique jobActivity identifier
* `resource_task_id` - (Optional) Task Id
* `scheduler_job_id` - (Required) unique SchedulerJob identifier
* `sequence` - (Optional) Task Order Sequence
* `step_name` - (Optional) Unique step name
* `target_name` - (Optional) Unique target name


## Attributes Reference

The following attributes are exported:

* `resource_collection` - The list of resource_collection.

### SchedulerJobJobActivityResource Reference

The following attributes are exported:

* `items` - List of Execution Resources.
	* `description` - Description of the Resource Execution status. If there are any errors, this can also include a short error message. 
	* `resource_display_name` - Resource Display Name.
	* `resource_id` - Resource Identifier associated with the Work Request.
	* `sequence` - The sequence of the Resource.
	* `status` - Status of the Job at Resource Level.
	* `targets` - Targets associated with the resource.
		* `description` - Description of the Execution status. If there are any errors, this can also include a short error message.  
		* `status` - Status of the Job at target Level.
		* `target_name` - Target Name.
	* `time_ended` - The time the task ended for the resource. An RFC3339 formatted datetime string
	* `time_started` - The time the task started for the resource. An RFC3339 formatted datetime string

