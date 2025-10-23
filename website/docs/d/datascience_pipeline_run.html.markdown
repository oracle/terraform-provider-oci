---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_pipeline_run"
sidebar_current: "docs-oci-datasource-datascience-pipeline_run"
description: |-
  Provides details about a specific Pipeline Run in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_pipeline_run
This data source provides details about a specific Pipeline Run resource in Oracle Cloud Infrastructure Data Science service.

Gets a PipelineRun by identifier.

## Example Usage

```hcl
data "oci_datascience_pipeline_run" "test_pipeline_run" {
	#Required
	pipeline_run_id = oci_datascience_pipeline_run.test_pipeline_run.id
}
```

## Argument Reference

The following arguments are supported:

* `pipeline_run_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline run.


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
* `infrastructure_configuration_override_details` - The infrastructure configuration details of a pipeline or a step.
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance. 
	* `block_storage_size_in_gbs_parameterized` - The size of the block storage volume to attach to the pipeline step run instance specified as a parameter. This overrides the blockStorageSizeInGBs value. The request will fail if the parameters used are null or invalid. 
	* `shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
		* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory available to the pipeline step run instance GBs. 
		* `memory_in_gbs_parameterized` - The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
		* `ocpus` - The total number of OCPUs available to the pipeline step run instance. 
		* `ocpus_parameterized` - The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid.
	* `shape_name` - The shape used to launch the instance for all step runs in the pipeline.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `log_configuration_override_details` - The pipeline log configuration details.
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - If customer logging is enabled for pipeline.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `log_details` - Customer logging details for pipeline run.
	* `log_group_id` - The log group id for where log objects will be for pipeline runs.
	* `log_id` - The log id of the log object the pipeline run logs will be shipped to.
* `parameters_override` - Parameters override used in the pipeline run.
* `pipeline_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline for which pipeline run is created.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline run with.
* `state` - The current state of the pipeline run.
* `delete_related_job_runs` - If set to true will delete related job runs. 
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
	* `step_dataflow_configuration_details` - The configuration details of a Dataflow step.
		* `configuration` - The Spark configuration passed to the running process.
		* `driver_shape` - The VM shape for the driver.
		* `driver_shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
		* `executor_shape` - The VM shape for the executors.
		* `executor_shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
		* `logs_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
		* `num_executors` - The number of executor VMs requested.
		* `warehouse_bucket_uri` - An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs.
	* `step_infrastructure_configuration_details` - The infrastructure configuration details of a pipeline or a step.
		* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance. 
		* `block_storage_size_in_gbs_parameterized` - The size of the block storage volume to attach to the pipeline step run instance specified as a parameter. This overrides the blockStorageSizeInGBs value. The request will fail if the parameters used are null or invalid. 
		* `shape_config_details` - Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid.
		* `shape_name` - The shape used to launch the instance for all step runs in the pipeline.
		* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
	* `step_name` - The name of the step.
	* `step_storage_mount_configuration_details_list` - The storage mount details to mount to the instance running the pipeline step.
		* `bucket` - The object storage bucket
		* `destination_directory_name` - The local directory name to be mounted
		* `destination_path` - The local path of the mounted directory, excluding directory name.
		* `export_id` - OCID of the export
		* `mount_target_id` - OCID of the mount target
		* `namespace` - The object storage namespace
		* `prefix` - Prefix in the bucket to mount
		* `storage_type` - The type of storage.
* `step_runs` - Array of StepRun object for each step.
	* `dataflow_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dataflow run triggered for this step run.
	* `job_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run triggered for this step run.
	* `lifecycle_details` - Details of the state of the step run.
	* `state` - The state of the step run.
	* `step_name` - The name of the step.
	* `step_run_name` - Name used when creating the steprun.
	* `step_type` - The type of step.
	* `time_finished` - The date and time the pipeline step run finshed executing in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	* `time_started` - The date and time the pipeline step run was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `storage_mount_configuration_override_details_list` - The storage mount override details to mount to the instance running the pipeline step.
	* `bucket` - The object storage bucket
	* `destination_directory_name` - The local directory name to be mounted
	* `destination_path` - The local path of the mounted directory, excluding directory name.
	* `export_id` - OCID of the export
	* `mount_target_id` - OCID of the mount target
	* `namespace` - The object storage namespace
	* `prefix` - Prefix in the bucket to mount
	* `storage_type` - The type of storage.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_accepted` - The date and time the pipeline run was accepted in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_finished` - The date and time the pipeline run request was finished in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_started` - The date and time the pipeline run request was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the pipeline run was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

