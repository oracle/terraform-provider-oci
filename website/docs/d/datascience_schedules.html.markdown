---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_schedules"
sidebar_current: "docs-oci-datasource-datascience-schedules"
description: |-
  Provides the list of Schedules in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_schedules
This data source provides the list of Schedules in Oracle Cloud Infrastructure Data Science service.

Returns a list of Schedules.


## Example Usage

```hcl
data "oci_datascience_schedules" "test_schedules" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.schedule_display_name
	id = var.schedule_id
	project_id = oci_datascience_project.test_project.id
	state = var.schedule_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) unique Schedule identifier
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `schedules` - The list of schedules.

### Schedule Reference

The following attributes are exported:

* `action` - The schedule action
	* `action_details` - Schedule Http action details
		* `create_job_run_details` - Parameters needed to create a new job run. 
			* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job run.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly display name for the resource.
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
			* `job_configuration_override_details` - The job configuration details 
				* `command_line_arguments` - The arguments to pass to the job. 
				* `environment_variables` - Environment variables to set for the job. 
				* `job_type` - The type of job.
				* `maximum_runtime_in_minutes` - A time bound for the execution of the job. Timer starts when the job becomes active. 
			* `job_environment_configuration_override_details` - Environment configuration to capture job runtime dependencies.
				* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
				* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
				* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
				* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
				* `image_signature_id` - OCID of the container image signature
				* `job_environment_type` - The environment configuration type used for job runtime.
			* `job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to create a run for.
			* `job_log_configuration_override_details` - Logging configuration for resource. 
				* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
				* `enable_logging` - If customer logging is enabled for job runs.
				* `log_group_id` - The log group id for where log objects are for job runs. 
				* `log_id` - The log id the job run will push logs too. 
			* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job run with.
		* `create_pipeline_run_details` - The information about new PipelineRun.
			* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
			* `configuration_override_details` - The configuration details of a pipeline.
				* `command_line_arguments` - The command line arguments to set for steps in the pipeline. 
				* `environment_variables` - Environment variables to set for steps in the pipeline.
				* `maximum_runtime_in_minutes` - A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
				* `type` - The type of pipeline.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly display name for the resource.
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
			* `log_configuration_override_details` - The pipeline log configuration details.
				* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for pipeline runs.
				* `enable_logging` - If customer logging is enabled for pipeline.
				* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
				* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
			* `pipeline_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
			* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
			* `step_override_details` - Array of step override details. Only Step Configuration is allowed to be overridden.
				* `step_configuration_details` - The configuration details of a step.
					* `command_line_arguments` - The command line arguments to set for step.
					* `environment_variables` - Environment variables to set for step.
					* `maximum_runtime_in_minutes` - A time bound for the execution of the step.
				* `step_container_configuration_details` - Container Details for a step in pipeline.
					* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
					* `container_type` - The type of container.
					* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
					* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. 
					* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
					* `image_signature_id` - OCID of the container image signature
				* `step_name` - The name of the step.
			* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
		* `http_action_type` - The type of http action to trigger.
		* `ml_application_instance_view_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
		* `trigger_ml_application_instance_view_flow_details` - Payload for trigger request endpoint
			* `parameters` - Parameters provided for given trigger invocation (they must match predefined schema)
				* `name` - Name of trigger parameter
				* `value` - Value of trigger parameter
			* `trigger_name` - Name of trigger
	* `action_type` - The Schedule Action type
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the schedule.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the schedule.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the schedule.
* `display_name` - A user-friendly display name for the resource. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
* `last_schedule_run_details` - Details about the action performed by the last schedule execution. Example: `Invoked ML Application trigger.` 
* `lifecycle_details` - A message describing the current state in more detail.
* `log_details` - Custom logging details for schedule execution.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom log to be used for Schedule logging.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the schedule.
* `state` - The current state of the schedule.           Example: `ACTIVE` 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the schedule was created. Format is defined by RFC3339.           Example: `2022-08-05T01:02:29.600Z` 
* `time_last_schedule_run` - The last schedule execution time. Format is defined by RFC3339. Example: `2022-08-05T01:02:29.600Z` 
* `time_next_scheduled_run` - The next scheduled execution time for the schedule. Format is defined by RFC3339. Example: `2022-08-05T01:02:29.600Z` 
* `time_updated` - The date and time the schedule was updated. Format is defined by RFC3339.           Example: `2022-09-05T01:02:29.600Z` 
* `trigger` - The trigger of the schedule can be UNIX cron or iCal expression or simple interval
	* `cron_expression` - Schedule cron expression
	* `frequency` - The type of frequency
	* `interval` - The interval of frequency.
	* `is_random_start_time` - when true and timeStart is null, system generate a random start time between now and now + interval;  isRandomStartTime can be true if timeStart is null. 
	* `recurrence` - This recurrence field conforms to RFC-5545 formatting
	* `time_end` - The schedule end date time, if null, the schedule will never expire. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `time_start` - The schedule starting date time, if null, System set the time when schedule is created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `trigger_type` - The schedule trigger type

