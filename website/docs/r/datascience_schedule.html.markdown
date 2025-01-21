---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_schedule"
sidebar_current: "docs-oci-resource-datascience-schedule"
description: |-
  Provides the Schedule resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_schedule
This resource provides the Schedule resource in Oracle Cloud Infrastructure Data Science service.

Creates a new Schedule.


## Example Usage

```hcl
resource "oci_datascience_schedule" "test_schedule" {
	#Required
	action {
		#Required
		action_details {
			#Required
			http_action_type = var.schedule_action_action_details_http_action_type

			#Optional
			create_job_run_details {

				#Optional
				compartment_id = var.compartment_id
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.schedule_action_action_details_create_job_run_details_display_name
				freeform_tags = {"Department"= "Finance"}
				job_configuration_override_details {
					#Required
					job_type = var.schedule_action_action_details_create_job_run_details_job_configuration_override_details_job_type

					#Optional
					command_line_arguments = var.schedule_action_action_details_create_job_run_details_job_configuration_override_details_command_line_arguments
					environment_variables = var.schedule_action_action_details_create_job_run_details_job_configuration_override_details_environment_variables
					maximum_runtime_in_minutes = var.schedule_action_action_details_create_job_run_details_job_configuration_override_details_maximum_runtime_in_minutes
				}
				job_environment_configuration_override_details {
					#Required
					image = var.schedule_action_action_details_create_job_run_details_job_environment_configuration_override_details_image
					job_environment_type = var.schedule_action_action_details_create_job_run_details_job_environment_configuration_override_details_job_environment_type

					#Optional
					cmd = var.schedule_action_action_details_create_job_run_details_job_environment_configuration_override_details_cmd
					entrypoint = var.schedule_action_action_details_create_job_run_details_job_environment_configuration_override_details_entrypoint
					image_digest = var.schedule_action_action_details_create_job_run_details_job_environment_configuration_override_details_image_digest
					image_signature_id = oci_datascience_image_signature.test_image_signature.id
				}
				job_id = oci_datascience_job.test_job.id
				job_log_configuration_override_details {

					#Optional
					enable_auto_log_creation = var.schedule_action_action_details_create_job_run_details_job_log_configuration_override_details_enable_auto_log_creation
					enable_logging = var.schedule_action_action_details_create_job_run_details_job_log_configuration_override_details_enable_logging
					log_group_id = oci_logging_log_group.test_log_group.id
					log_id = oci_logging_log.test_log.id
				}
				project_id = oci_datascience_project.test_project.id
			}
			create_pipeline_run_details {

				#Optional
				compartment_id = var.compartment_id
				configuration_override_details {
					#Required
					type = var.schedule_action_action_details_create_pipeline_run_details_configuration_override_details_type

					#Optional
					command_line_arguments = var.schedule_action_action_details_create_pipeline_run_details_configuration_override_details_command_line_arguments
					environment_variables = var.schedule_action_action_details_create_pipeline_run_details_configuration_override_details_environment_variables
					maximum_runtime_in_minutes = var.schedule_action_action_details_create_pipeline_run_details_configuration_override_details_maximum_runtime_in_minutes
				}
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.schedule_action_action_details_create_pipeline_run_details_display_name
				freeform_tags = {"Department"= "Finance"}
				log_configuration_override_details {

					#Optional
					enable_auto_log_creation = var.schedule_action_action_details_create_pipeline_run_details_log_configuration_override_details_enable_auto_log_creation
					enable_logging = var.schedule_action_action_details_create_pipeline_run_details_log_configuration_override_details_enable_logging
					log_group_id = oci_logging_log_group.test_log_group.id
					log_id = oci_logging_log.test_log.id
				}
				pipeline_id = oci_datascience_pipeline.test_pipeline.id
				project_id = oci_datascience_project.test_project.id
				step_override_details {

					#Optional
					step_configuration_details {

						#Optional
						command_line_arguments = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_configuration_details_command_line_arguments
						environment_variables = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_configuration_details_environment_variables
						maximum_runtime_in_minutes = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_configuration_details_maximum_runtime_in_minutes
					}
					step_container_configuration_details {
						#Required
						container_type = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_container_configuration_details_container_type
						image = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_container_configuration_details_image

						#Optional
						cmd = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_container_configuration_details_cmd
						entrypoint = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_container_configuration_details_entrypoint
						image_digest = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_container_configuration_details_image_digest
						image_signature_id = oci_datascience_image_signature.test_image_signature.id
					}
					step_name = var.schedule_action_action_details_create_pipeline_run_details_step_override_details_step_name
				}
				system_tags = var.schedule_action_action_details_create_pipeline_run_details_system_tags
			}
			ml_application_instance_view_id = oci_dns_view.test_view.id
			trigger_ml_application_instance_view_flow_details {

				#Optional
				parameters {

					#Optional
					name = var.schedule_action_action_details_trigger_ml_application_instance_view_flow_details_parameters_name
					value = var.schedule_action_action_details_trigger_ml_application_instance_view_flow_details_parameters_value
				}
				trigger_name = oci_devops_trigger.test_trigger.name
			}
		}
		action_type = var.schedule_action_action_type
	}
	compartment_id = var.compartment_id
	display_name = var.schedule_display_name
	project_id = oci_datascience_project.test_project.id
	trigger {
		#Required
		trigger_type = var.schedule_trigger_trigger_type

		#Optional
		cron_expression = var.schedule_trigger_cron_expression
		frequency = var.schedule_trigger_frequency
		interval = var.schedule_trigger_interval
		is_random_start_time = var.schedule_trigger_is_random_start_time
		recurrence = var.schedule_trigger_recurrence
		time_end = var.schedule_trigger_time_end
		time_start = var.schedule_trigger_time_start
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.schedule_description
	freeform_tags = {"Department"= "Finance"}
	log_details {
		#Required
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required) (Updatable) The schedule action
	* `action_details` - (Required) (Updatable) Schedule Http action details
		* `create_job_run_details` - (Required when http_action_type=CREATE_JOB_RUN) (Updatable) Parameters needed to create a new job run. 
			* `compartment_id` - (Required when http_action_type=CREATE_JOB_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job run.
			* `defined_tags` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) A user-friendly display name for the resource.
			* `freeform_tags` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
			* `job_configuration_override_details` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) The job configuration details 
				* `command_line_arguments` - (Optional) (Updatable) The arguments to pass to the job. 
				* `environment_variables` - (Optional) (Updatable) Environment variables to set for the job. 
				* `job_type` - (Required) (Updatable) The type of job.
				* `maximum_runtime_in_minutes` - (Optional) (Updatable) A time bound for the execution of the job. Timer starts when the job becomes active. 
			* `job_environment_configuration_override_details` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) Environment configuration to capture job runtime dependencies.
				* `cmd` - (Optional) (Updatable) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
				* `entrypoint` - (Optional) (Updatable) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
				* `image` - (Required) (Updatable) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
				* `image_digest` - (Optional) (Updatable) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
				* `image_signature_id` - (Optional) (Updatable) OCID of the container image signature
				* `job_environment_type` - (Required) (Updatable) The environment configuration type used for job runtime.
			* `job_id` - (Required when http_action_type=CREATE_JOB_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to create a run for.
			* `job_log_configuration_override_details` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) Logging configuration for resource. 
				* `enable_auto_log_creation` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) If automatic on-behalf-of log object creation is enabled for job runs. 
				* `enable_logging` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) If customer logging is enabled for job runs.
				* `log_group_id` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) The log group id for where log objects are for job runs. 
				* `log_id` - (Applicable when http_action_type=CREATE_JOB_RUN) (Updatable) The log id the job run will push logs too. 
			* `project_id` - (Required when http_action_type=CREATE_JOB_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job run with.
		* `create_pipeline_run_details` - (Required when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The information about new PipelineRun.
			* `compartment_id` - (Required when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
			* `configuration_override_details` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The configuration details of a pipeline.
				* `command_line_arguments` - (Optional) (Updatable) The command line arguments to set for steps in the pipeline. 
				* `environment_variables` - (Optional) (Updatable) Environment variables to set for steps in the pipeline.
				* `maximum_runtime_in_minutes` - (Optional) (Updatable) A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
				* `type` - (Required) (Updatable) The type of pipeline.
			* `defined_tags` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) A user-friendly display name for the resource.
			* `freeform_tags` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
			* `log_configuration_override_details` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The pipeline log configuration details.
				* `enable_auto_log_creation` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) If automatic on-behalf-of log object creation is enabled for pipeline runs.
				* `enable_logging` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) If customer logging is enabled for pipeline.
				* `log_group_id` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
				* `log_id` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
			* `pipeline_id` - (Required when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
			* `project_id` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
			* `step_override_details` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Array of step override details. Only Step Configuration is allowed to be overridden.
				* `step_configuration_details` - (Required when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The configuration details of a step.
					* `command_line_arguments` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The command line arguments to set for step.
					* `environment_variables` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Environment variables to set for step.
					* `maximum_runtime_in_minutes` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) A time bound for the execution of the step.
				* `step_container_configuration_details` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Container Details for a step in pipeline.
					* `cmd` - (Optional) (Updatable) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
					* `container_type` - (Required) (Updatable) The type of container.
					* `entrypoint` - (Optional) (Updatable) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
					* `image` - (Required) (Updatable) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. 
					* `image_digest` - (Optional) (Updatable) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
					* `image_signature_id` - (Optional) (Updatable) OCID of the container image signature
				* `step_name` - (Required when http_action_type=CREATE_PIPELINE_RUN) (Updatable) The name of the step.
			* `system_tags` - (Applicable when http_action_type=CREATE_PIPELINE_RUN) (Updatable) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
		* `http_action_type` - (Required) (Updatable) The type of http action to trigger.
		* `ml_application_instance_view_id` - (Required when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
		* `trigger_ml_application_instance_view_flow_details` - (Required when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) Payload for trigger request endpoint
			* `parameters` - (Applicable when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) Parameters provided for given trigger invocation (they must match predefined schema)
				* `name` - (Required when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) Name of trigger parameter
				* `value` - (Required when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) Value of trigger parameter
			* `trigger_name` - (Required when http_action_type=INVOKE_ML_APPLICATION_PROVIDER_TRIGGER) (Updatable) Name of trigger
	* `action_type` - (Required) (Updatable) The Schedule Action type
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the schedule.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the schedule.
* `display_name` - (Required) (Updatable) A user-friendly name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `log_details` - (Optional) (Updatable) Custom logging details for schedule execution.
	* `log_group_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom log to be used for Schedule logging.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the schedule with.
* `trigger` - (Required) (Updatable) The trigger of the schedule can be UNIX cron or iCal expression or simple interval
	* `cron_expression` - (Required when trigger_type=CRON) (Updatable) Schedule cron expression
	* `frequency` - (Required when trigger_type=INTERVAL) (Updatable) The type of frequency
	* `interval` - (Required when trigger_type=INTERVAL) (Updatable) The interval of frequency.
	* `is_random_start_time` - (Applicable when trigger_type=INTERVAL) (Updatable) when true and timeStart is null, system generate a random start time between now and now + interval;  isRandomStartTime can be true if timeStart is null. 
	* `recurrence` - (Required when trigger_type=ICAL) (Updatable) This recurrence field conforms to RFC-5545 formatting
	* `time_end` - (Optional) (Updatable) The schedule end date time, if null, the schedule will never expire. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `time_start` - (Optional) (Updatable) The schedule starting date time, if null, System set the time when schedule is created. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `trigger_type` - (Required) (Updatable) The schedule trigger type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Schedule
	* `update` - (Defaults to 20 minutes), when updating the Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Schedule


## Import

Schedules can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_schedule.test_schedule "id"
```

