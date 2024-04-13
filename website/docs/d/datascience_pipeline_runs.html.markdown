---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_pipeline_runs"
sidebar_current: "docs-oci-datasource-datascience-pipeline_runs"
description: |-
  Provides the list of Pipeline Runs in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_pipeline_runs
This data source provides the list of Pipeline Runs in Oracle Cloud Infrastructure Data Science service.

Returns a list of PipelineRuns.

## Example Usage

```hcl
data "oci_datascience_pipeline_runs" "test_pipeline_runs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.pipeline_run_created_by
	display_name = var.pipeline_run_display_name
	id = var.pipeline_run_id
	pipeline_id = oci_datascience_pipeline.test_pipeline.id
	state = var.pipeline_run_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `pipeline_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
* `state` - (Optional) The current state of the PipelineRun.


## Attributes Reference

The following attributes are exported:

* `pipeline_runs` - The list of pipeline_runs.

### PipelineRun Reference

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

