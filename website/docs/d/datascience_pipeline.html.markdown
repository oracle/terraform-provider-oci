---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_pipeline"
sidebar_current: "docs-oci-datasource-datascience-pipeline"
description: |-
  Provides details about a specific Pipeline in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_pipeline
This data source provides details about a specific Pipeline resource in Oracle Cloud Infrastructure Data Science service.

Gets a Pipeline by identifier.

## Example Usage

```hcl
data "oci_datascience_pipeline" "test_pipeline" {
	#Required
	pipeline_id = oci_datascience_pipeline.test_pipeline.id
}
```

## Argument Reference

The following arguments are supported:

* `pipeline_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline.
* `configuration_details` - The configuration details of a pipeline.
	* `command_line_arguments` - The command line arguments to set for steps in the pipeline. 
	* `environment_variables` - Environment variables to set for steps in the pipeline.
	* `maximum_runtime_in_minutes` - A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
	* `type` - The type of pipeline.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the pipeline.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the pipeline.
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
* `infrastructure_configuration_details` - The infrastructure configuration details of a pipeline or a step.
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance. 
	* `shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
		* `memory_in_gbs` - A pipeline step run instance of type VM.Standard.E3.Flex allows memory to be specified. This specifies the size of the memory in GBs. 
		* `ocpus` - A pipeline step run instance of type VM.Standard.E3.Flex allows the ocpu count to be specified. 
	* `shape_name` - The shape used to launch the instance for all step runs in the pipeline.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `log_configuration_details` - The pipeline log configuration details.
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - If customer logging is enabled for pipeline.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline with.
* `state` - The current state of the pipeline.
* `delete_related_pipeline_runs` - If set to true will delete pipeline runs which are in a terminal state.
* `step_details` - Array of step details for each step.
	* `depends_on` - The list of step names this current step depends on for execution.
	* `description` - A short description of the step.
	* `is_artifact_uploaded` - A flag to indicate whether the artifact has been uploaded for this step or not.
	* `job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to be used as a step.
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
	* `step_infrastructure_configuration_details` - The infrastructure configuration details of a pipeline or a step.
		* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance. 
		* `shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `memory_in_gbs` - A pipeline step run instance of type VM.Standard.E3.Flex allows memory to be specified. This specifies the size of the memory in GBs. 
			* `ocpus` - A pipeline step run instance of type VM.Standard.E3.Flex allows the ocpu count to be specified. 
		* `shape_name` - The shape used to launch the instance for all step runs in the pipeline.
		* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
	* `step_name` - The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	* `step_type` - The type of step.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

