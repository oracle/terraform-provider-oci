---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_pipeline_run"
sidebar_current: "docs-oci-resource-datascience-pipeline_run"
description: |-
  Provides the Pipeline Run resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_pipeline_run
This resource provides the Pipeline Run resource in Oracle Cloud Infrastructure Data Science service.

Creates a new PipelineRun.


## Example Usage

```hcl
resource "oci_datascience_pipeline_run" "test_pipeline_run" {
	#Required
	compartment_id = var.compartment_id
	pipeline_id = oci_datascience_pipeline.test_pipeline.id

	#Optional
	configuration_override_details {
		#Required
		type = var.pipeline_run_configuration_override_details_type

		#Optional
		command_line_arguments = var.pipeline_run_configuration_override_details_command_line_arguments
		environment_variables = var.pipeline_run_configuration_override_details_environment_variables
		maximum_runtime_in_minutes = var.pipeline_run_configuration_override_details_maximum_runtime_in_minutes
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.pipeline_run_display_name
	freeform_tags = {"Department"= "Finance"}
	log_configuration_override_details {

		#Optional
		enable_auto_log_creation = var.pipeline_run_log_configuration_override_details_enable_auto_log_creation
		enable_logging = var.pipeline_run_log_configuration_override_details_enable_logging
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
	opc_parent_rpt_url = var.pipeline_run_opc_parent_rpt_url
	project_id = oci_datascience_project.test_project.id
	step_override_details {
		#Required
		step_configuration_details {

			#Optional
			command_line_arguments = var.pipeline_run_step_override_details_step_configuration_details_command_line_arguments
			environment_variables = var.pipeline_run_step_override_details_step_configuration_details_environment_variables
			maximum_runtime_in_minutes = var.pipeline_run_step_override_details_step_configuration_details_maximum_runtime_in_minutes
		}
		step_name = var.pipeline_run_step_override_details_step_name

		#Optional
		step_container_configuration_details {
			#Required
			container_type = var.pipeline_run_step_override_details_step_container_configuration_details_container_type
			image = var.pipeline_run_step_override_details_step_container_configuration_details_image

			#Optional
			cmd = var.pipeline_run_step_override_details_step_container_configuration_details_cmd
			entrypoint = var.pipeline_run_step_override_details_step_container_configuration_details_entrypoint
			image_digest = var.pipeline_run_step_override_details_step_container_configuration_details_image_digest
			image_signature_id = oci_datascience_image_signature.test_image_signature.id
		}
	}
	system_tags = var.pipeline_run_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
* `configuration_override_details` - (Optional) The configuration details of a pipeline.
	* `command_line_arguments` - (Optional) The command line arguments to set for steps in the pipeline. 
	* `environment_variables` - (Optional) Environment variables to set for steps in the pipeline.
	* `maximum_runtime_in_minutes` - (Optional) A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
	* `type` - (Required) The type of pipeline.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `log_configuration_override_details` - (Optional) The pipeline log configuration details.
	* `enable_auto_log_creation` - (Optional) If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - (Optional) If customer logging is enabled for pipeline.
	* `log_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `opc_parent_rpt_url` - (Optional) URL to fetch the Resource Principal Token from the parent resource. 
* `pipeline_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
* `project_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
* `step_override_details` - (Optional) Array of step override details. Only Step Configuration is allowed to be overridden.
	* `step_configuration_details` - (Required) The configuration details of a step.
		* `command_line_arguments` - (Optional) The command line arguments to set for step.
		* `environment_variables` - (Optional) Environment variables to set for step.
		* `maximum_runtime_in_minutes` - (Optional) A time bound for the execution of the step.
	* `step_container_configuration_details` - (Optional) Container Details for a step in pipeline.
		* `cmd` - (Optional) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
		* `container_type` - (Required) The type of container.
		* `entrypoint` - (Optional) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
		* `image` - (Required) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. 
		* `image_digest` - (Optional) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
		* `image_signature_id` - (Optional) OCID of the container image signature
	* `step_name` - (Required) The name of the step.
* `system_tags` - (Optional) Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline run.
* `configuration_details` - The configuration details of a pipeline.
	* `command_line_arguments` - The command line arguments to set for steps in the pipeline. 
	* `environment_variables` - Environment variables to set for steps in the pipeline.
	* `maximum_runtime_in_minutes` - A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
	* `type` - The type of pipeline.
* `configuration_override_details` - The configuration details of a pipeline.
	* `command_line_arguments` - The command line arguments to set for steps in the pipeline. 
	* `environment_variables` - Environment variables to set for steps in the pipeline.
	* `maximum_runtime_in_minutes` - A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
	* `type` - The type of pipeline.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the pipeline run.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline run.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `log_configuration_override_details` - The pipeline log configuration details.
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - If customer logging is enabled for pipeline.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `log_details` - Customer logging details for pipeline run.
	* `log_group_id` - The log group id for where log objects will be for pipeline runs.
	* `log_id` - The log id of the log object the pipeline run logs will be shipped to.
* `pipeline_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
* `state` - The current state of the pipeline run.
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
* `step_runs` - Array of StepRun object for each step.
	* `job_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run triggered for this step run.
	* `lifecycle_details` - Details of the state of the step run.
	* `state` - The state of the step run.
	* `step_name` - The name of the step.
	* `step_type` - The type of step.
	* `time_finished` - The date and time the pipeline step run finshed executing in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	* `time_started` - The date and time the pipeline step run was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_accepted` - The date and time the pipeline run was accepted in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_finished` - The date and time the pipeline run request was finished in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_started` - The date and time the pipeline run request was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the pipeline run was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pipeline Run
	* `update` - (Defaults to 20 minutes), when updating the Pipeline Run
	* `delete` - (Defaults to 20 minutes), when destroying the Pipeline Run


## Import

PipelineRuns can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_pipeline_run.test_pipeline_run "id"
```

