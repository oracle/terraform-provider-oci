---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_pipeline"
sidebar_current: "docs-oci-resource-datascience-pipeline"
description: |-
  Provides the Pipeline resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_pipeline
This resource provides the Pipeline resource in Oracle Cloud Infrastructure Data Science service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-science/latest/Pipeline

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datascience

Creates a new Pipeline.


## Example Usage

```hcl
resource "oci_datascience_pipeline" "test_pipeline" {
	#Required
	compartment_id = var.compartment_id
	project_id = oci_datascience_project.test_project.id
	step_details {
		#Required
		step_name = var.pipeline_step_details_step_name
		step_type = var.pipeline_step_details_step_type

		#Optional

		application_id = oci_dataflow_application.test_application.id=
		depends_on = var.pipeline_step_details_depends_on
		description = var.pipeline_step_details_description
		is_artifact_uploaded = var.pipeline_step_details_is_artifact_uploaded
		job_id = oci_datascience_job.test_job.id
		step_configuration_details {

			#Optional
			command_line_arguments = var.pipeline_step_details_step_configuration_details_command_line_arguments
			environment_variables = var.pipeline_step_details_step_configuration_details_environment_variables
			maximum_runtime_in_minutes = var.pipeline_step_details_step_configuration_details_maximum_runtime_in_minutes
		}
		step_container_configuration_details {
			#Required
			container_type = var.pipeline_step_details_step_container_configuration_details_container_type
			image = var.pipeline_step_details_step_container_configuration_details_image

			#Optional
			cmd = var.pipeline_step_details_step_container_configuration_details_cmd
			entrypoint = var.pipeline_step_details_step_container_configuration_details_entrypoint
			image_digest = var.pipeline_step_details_step_container_configuration_details_image_digest
			image_signature_id = oci_datascience_image_signature.test_image_signature.id
		}
		step_dataflow_configuration_details {

			#Optional
			configuration = var.pipeline_step_details_step_dataflow_configuration_details_configuration
			driver_shape = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape
			driver_shape_config_details {

				#Optional
				cpu_baseline = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape_config_details_cpu_baseline
				memory_in_gbs = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape_config_details_memory_in_gbs
				memory_in_gbs_parameterized = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape_config_details_memory_in_gbs_parameterized
				ocpus = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape_config_details_ocpus
				ocpus_parameterized = var.pipeline_step_details_step_dataflow_configuration_details_driver_shape_config_details_ocpus_parameterized
			}
			executor_shape = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape
			executor_shape_config_details {

				#Optional
				cpu_baseline = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape_config_details_cpu_baseline
				memory_in_gbs = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape_config_details_memory_in_gbs
				memory_in_gbs_parameterized = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape_config_details_memory_in_gbs_parameterized
				ocpus = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape_config_details_ocpus
				ocpus_parameterized = var.pipeline_step_details_step_dataflow_configuration_details_executor_shape_config_details_ocpus_parameterized
			}
			logs_bucket_uri = var.pipeline_step_details_step_dataflow_configuration_details_logs_bucket_uri
			num_executors = var.pipeline_step_details_step_dataflow_configuration_details_num_executors
			warehouse_bucket_uri = var.pipeline_step_details_step_dataflow_configuration_details_warehouse_bucket_uri
		}
		step_infrastructure_configuration_details {

			#Optional
			block_storage_size_in_gbs = var.pipeline_step_details_step_infrastructure_configuration_details_block_storage_size_in_gbs
			block_storage_size_in_gbs_parameterized = var.pipeline_step_details_step_infrastructure_configuration_details_block_storage_size_in_gbs_parameterized
			shape_config_details {

				#Optional
				cpu_baseline = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_cpu_baseline
				memory_in_gbs = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_memory_in_gbs
				memory_in_gbs_parameterized = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_memory_in_gbs_parameterized
				ocpus = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_ocpus
				ocpus_parameterized = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_ocpus_parameterized
			}
			shape_name = oci_core_shape.test_shape.name
			subnet_id = oci_core_subnet.test_subnet.id
		}
		step_parameters {
			#Required
			output {
				#Required
				output_file = var.pipeline_step_details_step_parameters_output_output_file
				output_parameter_type = var.pipeline_step_details_step_parameters_output_output_parameter_type
				parameter_names = var.pipeline_step_details_step_parameters_output_parameter_names
			}
			parameter_type = var.pipeline_step_details_step_parameters_parameter_type
		}
		step_run_name = var.pipeline_step_details_step_run_name
		step_storage_mount_configuration_details_list {
			#Required
			destination_directory_name = var.pipeline_step_details_step_storage_mount_configuration_details_list_destination_directory_name
			storage_type = var.pipeline_step_details_step_storage_mount_configuration_details_list_storage_type

			#Optional
			bucket = var.pipeline_step_details_step_storage_mount_configuration_details_list_bucket
			destination_path = var.pipeline_step_details_step_storage_mount_configuration_details_list_destination_path
			export_id = oci_file_storage_export.test_export.id
			mount_target_id = oci_file_storage_mount_target.test_mount_target.id
			namespace = var.pipeline_step_details_step_storage_mount_configuration_details_list_namespace
			prefix = var.pipeline_step_details_step_storage_mount_configuration_details_list_prefix
		}
	}

	#Optional
	configuration_details {
		#Required
		type = var.pipeline_configuration_details_type

		#Optional
		command_line_arguments = var.pipeline_configuration_details_command_line_arguments
		environment_variables = var.pipeline_configuration_details_environment_variables
		maximum_runtime_in_minutes = var.pipeline_configuration_details_maximum_runtime_in_minutes
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.pipeline_description
	display_name = var.pipeline_display_name
	freeform_tags = {"Department"= "Finance"}
	infrastructure_configuration_details {
		#Required
		block_storage_size_in_gbs = var.pipeline_infrastructure_configuration_details_block_storage_size_in_gbs
		shape_name = oci_core_shape.test_shape.name

		#Optional
		block_storage_size_in_gbs_parameterized = var.pipeline_infrastructure_configuration_details_block_storage_size_in_gbs_parameterized
		shape_config_details {

			#Optional
			cpu_baseline = var.pipeline_infrastructure_configuration_details_shape_config_details_cpu_baseline
			memory_in_gbs = var.pipeline_infrastructure_configuration_details_shape_config_details_memory_in_gbs
			memory_in_gbs_parameterized = var.pipeline_infrastructure_configuration_details_shape_config_details_memory_in_gbs_parameterized
			ocpus = var.pipeline_infrastructure_configuration_details_shape_config_details_ocpus
			ocpus_parameterized = var.pipeline_infrastructure_configuration_details_shape_config_details_ocpus_parameterized
		}
		subnet_id = oci_core_subnet.test_subnet.id
	}
	log_configuration_details {

		#Optional
		enable_auto_log_creation = var.pipeline_log_configuration_details_enable_auto_log_creation
		enable_logging = var.pipeline_log_configuration_details_enable_logging
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
	parameters = var.pipeline_parameters
	storage_mount_configuration_details_list {
		#Required
		destination_directory_name = var.pipeline_storage_mount_configuration_details_list_destination_directory_name
		storage_type = var.pipeline_storage_mount_configuration_details_list_storage_type

		#Optional
		bucket = var.pipeline_storage_mount_configuration_details_list_bucket
		destination_path = var.pipeline_storage_mount_configuration_details_list_destination_path
		export_id = oci_file_storage_export.test_export.id
		mount_target_id = oci_file_storage_mount_target.test_mount_target.id
		namespace = var.pipeline_storage_mount_configuration_details_list_namespace
		prefix = var.pipeline_storage_mount_configuration_details_list_prefix
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the pipeline.
* `configuration_details` - (Optional) (Updatable) The configuration details of a pipeline.
	* `command_line_arguments` - (Optional) (Updatable) The command line arguments to set for steps in the pipeline. 
	* `environment_variables` - (Optional) (Updatable) Environment variables to set for steps in the pipeline.
	* `maximum_runtime_in_minutes` - (Optional) (Updatable) A time bound for the execution of the entire Pipeline. Timer starts when the Pipeline Run is in progress.
	* `type` - (Required) (Updatable) The type of pipeline.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the pipeline.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `infrastructure_configuration_details` - (Optional) (Updatable) The infrastructure configuration details of a pipeline or a step.
	* `block_storage_size_in_gbs` - (Required) (Updatable) The size of the block storage volume to attach to the instance. 
	* `block_storage_size_in_gbs_parameterized` - (Optional) (Updatable) The size of the block storage volume to attach to the pipeline step run instance specified as a parameter. This overrides the blockStorageSizeInGBs value. The request will fail if the parameters used are null or invalid. 
	* `shape_config_details` - (Optional) (Updatable) Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
		* `cpu_baseline` - (Optional) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the pipeline step run instance GBs. 
		* `memory_in_gbs_parameterized` - (Optional) (Updatable) The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
		* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to the pipeline step run instance. 
		* `ocpus_parameterized` - (Optional) (Updatable) The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
	* `shape_name` - (Required) (Updatable) The shape used to launch the instance for all step runs in the pipeline.
	* `subnet_id` - (Optional) (Updatable) The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
* `log_configuration_details` - (Optional) (Updatable) The pipeline log configuration details.
	* `enable_auto_log_creation` - (Optional) (Updatable) If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - (Optional) (Updatable) If customer logging is enabled for pipeline.
	* `log_group_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `parameters` - (Optional) (Updatable) Parameters used in the pipeline.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline with.
* `step_details` - (Required) (Updatable) Array of step details for each step.

	* `application_id` - (Required when step_type=DATAFLOW) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dataflow application to be used as a step.
	* `depends_on` - (Optional) The list of step names this current step depends on for execution.
	* `description` - (Optional) (Updatable) A short description of the step.
	* `is_artifact_uploaded` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) A flag to indicate whether the artifact has been uploaded for this step or not.
	* `job_id` - (Required when step_type=ML_JOB) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to be used as a step.
	* `step_configuration_details` - (Optional) (Updatable) The configuration details of a step.
		* `command_line_arguments` - (Optional) (Updatable) The command line arguments to set for step.
		* `environment_variables` - (Optional) (Updatable) Environment variables to set for step.
		* `maximum_runtime_in_minutes` - (Optional) (Updatable) A time bound for the execution of the step.
	* `step_container_configuration_details` - (Required when step_type=CONTAINER) Container Details for a step in pipeline.
		* `cmd` - (Optional) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
		* `container_type` - (Required) The type of container.
		* `entrypoint` - (Optional) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
		* `image` - (Required) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. 
		* `image_digest` - (Optional) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
		* `image_signature_id` - (Optional) OCID of the container image signature
	* `step_dataflow_configuration_details` - (Applicable when step_type=DATAFLOW) (Updatable) The configuration details of a Dataflow step.
		* `configuration` - (Applicable when step_type=DATAFLOW) (Updatable) The Spark configuration passed to the running process.
		* `driver_shape` - (Applicable when step_type=DATAFLOW) (Updatable) The VM shape for the driver.
		* `driver_shape_config_details` - (Applicable when step_type=DATAFLOW) (Updatable) Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - (Applicable when step_type=DATAFLOW) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - (Applicable when step_type=DATAFLOW) (Updatable) The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - (Applicable when step_type=DATAFLOW) (Updatable) The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - (Applicable when step_type=DATAFLOW) (Updatable) The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - (Applicable when step_type=DATAFLOW) (Updatable) The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
		* `executor_shape` - (Applicable when step_type=DATAFLOW) (Updatable) The VM shape for the executors.
		* `executor_shape_config_details` - (Applicable when step_type=DATAFLOW) (Updatable) Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - (Applicable when step_type=DATAFLOW) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - (Applicable when step_type=DATAFLOW) (Updatable) The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - (Applicable when step_type=DATAFLOW) (Updatable) The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - (Applicable when step_type=DATAFLOW) (Updatable) The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - (Applicable when step_type=DATAFLOW) (Updatable) The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
		* `logs_bucket_uri` - (Applicable when step_type=DATAFLOW) (Updatable) An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded.
		* `num_executors` - (Applicable when step_type=DATAFLOW) (Updatable) The number of executor VMs requested.
		* `warehouse_bucket_uri` - (Applicable when step_type=DATAFLOW) (Updatable) An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs.
	* `step_infrastructure_configuration_details` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The infrastructure configuration details of a pipeline or a step.
		* `block_storage_size_in_gbs` - (Required when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The size of the block storage volume to attach to the instance. 
		* `block_storage_size_in_gbs_parameterized` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The size of the block storage volume to attach to the pipeline step run instance specified as a parameter. This overrides the blockStorageSizeInGBs value. The request will fail if the parameters used are null or invalid. 
		* `shape_config_details` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) Details for the pipeline step run shape configuration. Specify only when a flex shape is selected.
			* `cpu_baseline` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The total amount of memory available to the pipeline step run instance GBs. 
			* `memory_in_gbs_parameterized` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The total amount of memory available to the pipeline step run instance in GBs specified as a parameter. This overrides the memoryInGBs value. The request will fail if the parameters used are null or invalid. 
			* `ocpus` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The total number of OCPUs available to the pipeline step run instance. 
			* `ocpus_parameterized` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The total number of OCPUs available to the pipeline step run instance specified as a parameter. This overrides the ocpus value. The request will fail if the parameters used are null or invalid. 
		* `shape_name` - (Required when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The shape used to launch the instance for all step runs in the pipeline.
		* `subnet_id` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The subnet to create a secondary vnic in to attach to the instance running the pipeline step. 
	* `step_name` - (Required) (Updatable) The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	* `step_parameters` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT | ML_JOB) (Updatable) Pipeline step parameter details
		* `output` - (Required) (Updatable) Pipeline output parameter details
			* `output_file` - (Required) (Updatable) Output file name
			* `output_parameter_type` - (Required) (Updatable) Type of output parameters
			* `parameter_names` - (Required) (Updatable) The list of parameter names that will be output by this step
		* `parameter_type` - (Required) (Updatable) Type of step parameter
	* `step_run_name` - (Applicable when step_type=ML_JOB) (Updatable) Name used when creating the steprun.
	* `step_storage_mount_configuration_details_list` - (Applicable when step_type=CONTAINER | CUSTOM_SCRIPT) (Updatable) The storage mount details to mount to the instance running the pipeline step.
		* `bucket` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage bucket
		* `destination_directory_name` - (Required) (Updatable) The local directory name to be mounted
		* `destination_path` - (Optional) (Updatable) The local path of the mounted directory, excluding directory name.
		* `export_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the export
		* `mount_target_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the mount target
		* `namespace` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage namespace
		* `prefix` - (Applicable when storage_type=OBJECT_STORAGE) (Updatable) Prefix in the bucket to mount
		* `storage_type` - (Required) (Updatable) The type of storage.
	* `step_type` - (Required) (Updatable) The type of step.
* `storage_mount_configuration_details_list` - (Optional) (Updatable) The storage mount details to mount to the instance running the pipeline step.
	* `bucket` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage bucket
	* `destination_directory_name` - (Required) (Updatable) The local directory name to be mounted
	* `destination_path` - (Optional) (Updatable) The local path of the mounted directory, excluding directory name.
	* `export_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the export
	* `mount_target_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the mount target
	* `namespace` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage namespace
	* `prefix` - (Applicable when storage_type=OBJECT_STORAGE) (Updatable) Prefix in the bucket to mount
	* `storage_type` - (Required) (Updatable) The type of storage.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `log_configuration_details` - The pipeline log configuration details.
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for pipeline runs.
	* `enable_logging` - If customer logging is enabled for pipeline.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `parameters` - Parameters used in the pipeline.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the pipeline with.
* `state` - The current state of the pipeline.
* `step_details` - Array of step details for each step.
	* `application_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dataflow application to be used as a step.
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
	* `step_name` - The name of the step. It must be unique within the pipeline. This is used to create the pipeline DAG.
	* `step_parameters` - Pipeline step parameter details
		* `output` - Pipeline output parameter details
			* `output_file` - Output file name
			* `output_parameter_type` - Type of output parameters
			* `parameter_names` - The list of parameter names that will be output by this step
		* `parameter_type` - Type of step parameter
	* `step_run_name` - Name used when creating the steprun.
	* `step_storage_mount_configuration_details_list` - The storage mount details to mount to the instance running the pipeline step.
		* `bucket` - The object storage bucket
		* `destination_directory_name` - The local directory name to be mounted
		* `destination_path` - The local path of the mounted directory, excluding directory name.
		* `export_id` - OCID of the export
		* `mount_target_id` - OCID of the mount target
		* `namespace` - The object storage namespace
		* `prefix` - Prefix in the bucket to mount
		* `storage_type` - The type of storage.
	* `step_type` - The type of step.
* `storage_mount_configuration_details_list` - The storage mount details to mount to the instance running the pipeline step.
	* `bucket` - The object storage bucket
	* `destination_directory_name` - The local directory name to be mounted
	* `destination_path` - The local path of the mounted directory, excluding directory name.
	* `export_id` - OCID of the export
	* `mount_target_id` - OCID of the mount target
	* `namespace` - The object storage namespace
	* `prefix` - Prefix in the bucket to mount
	* `storage_type` - The type of storage.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pipeline
	* `update` - (Defaults to 20 minutes), when updating the Pipeline
	* `delete` - (Defaults to 20 minutes), when destroying the Pipeline


## Import

Pipelines can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_pipeline.test_pipeline "id"
```

