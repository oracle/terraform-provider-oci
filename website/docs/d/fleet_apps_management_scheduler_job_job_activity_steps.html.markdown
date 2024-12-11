---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_job_job_activity_steps"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_job_job_activity_steps"
description: |-
  Provides the list of Scheduler Job Job Activity Steps in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_job_job_activity_steps
This data source provides the list of Scheduler Job Job Activity Steps in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of Steps for an Activity Execution.

## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_job_job_activity_steps" "test_scheduler_job_job_activity_steps" {
	#Required
	job_activity_id = oci_fleet_apps_management_job_activity.test_job_activity.id
	scheduler_job_id = oci_database_migration_job.test_job.id

	#Optional
	resource_task_id = oci_fleet_apps_management_resource_task.test_resource_task.id
	sequence = var.scheduler_job_job_activity_step_sequence
	step_name = var.scheduler_job_job_activity_step_step_name
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

* `step_collection` - The list of step_collection.

### SchedulerJobJobActivityStep Reference

The following attributes are exported:

* `items` - List of Execution steps.
	* `description` - Description of the step Execution
	* `is_rollback_task` - Is this a rollback task?
	* `sequence` - The sequence of the step
	* `status` - Status of the Task
	* `step_name` - Name of the Step
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `task_record_id` - The OCID of taskRecord assocaited with the step
	* `time_ended` - The time the task ended. An RFC3339 formatted datetime string
	* `time_started` - The time the task started. An RFC3339 formatted datetime string

