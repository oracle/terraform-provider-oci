---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployment"
sidebar_current: "docs-oci-resource-datascience-model_deployment"
description: |-
  Provides the Model Deployment resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_deployment
This resource provides the Model Deployment resource in Oracle Cloud Infrastructure Datascience service.

Creates a new model deployment.

## Example Usage

```hcl
resource "oci_datascience_model_deployment" "test_model_deployment" {
	#Required
	compartment_id = var.compartment_id
	model_deployment_configuration_details {
		#Required
		deployment_type = var.model_deployment_model_deployment_configuration_details_deployment_type
		model_configuration_details {
			#Required
			instance_configuration {
				#Required
				instance_shape_name = oci_core_shape.test_shape.name

				#Optional
				model_deployment_instance_shape_config_details {

					#Optional
					memory_in_gbs = var.model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_memory_in_gbs
					ocpus = var.model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_ocpus
				}
			}
			model_id = oci_datascience_model.test_model.id

			#Optional
			bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps
			scaling_policy {
				#Required
				instance_count = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_instance_count
				policy_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_policy_type
			}
		}

		#Optional
		environment_configuration_details {
			#Required
			environment_configuration_type = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_configuration_type

			#Optional
			cmd = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_cmd
			entrypoint = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_entrypoint
			environment_variables = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_variables
			health_check_port = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_health_check_port
			image = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_image
			image_digest = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_image_digest
			server_port = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_server_port
		}
	}
	project_id = oci_datascience_project.test_project.id

	#Optional
	category_log_details {

		#Optional
		access {
			#Required
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
		predict {
			#Required
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.model_deployment_description
	display_name = var.model_deployment_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `category_log_details` - (Optional) (Updatable) The log details for each category.
	* `access` - (Optional) (Updatable) The log details.
		* `log_group_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
	* `predict` - (Optional) (Updatable) The log details.
		* `log_group_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the model deployment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) A short description of the model deployment.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My ModelDeployment`
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `model_deployment_configuration_details` - (Required) (Updatable) The model deployment configuration details.
	* `deployment_type` - (Required) (Updatable) The type of the model deployment.
	* `environment_configuration_details` - (Optional) (Updatable) The configuration to carry the environment details thats used in Model Deployment creation
		* `cmd` - (Applicable when environment_configuration_type=OCIR_CONTAINER) (Updatable) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
		* `entrypoint` - (Applicable when environment_configuration_type=OCIR_CONTAINER) (Updatable) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
		* `environment_configuration_type` - (Required) (Updatable) The environment configuration type
		* `environment_variables` - (Optional) (Updatable) Environment variables to set for the web server container. The size of envVars must be less than 2048 bytes. Key should be under 32 characters. Key should contain only letters, digits and underscore (_) Key should start with a letter. Key should have at least 2 characters. Key should not end with underscore eg. `TEST_` Key if added cannot be empty. Value can be empty. No specific size limits on individual Values. But overall environment variables is limited to 2048 bytes. Key can't be reserved Model Deployment environment variables. 
		* `health_check_port` - (Applicable when environment_configuration_type=OCIR_CONTAINER) (Updatable) The port on which the container [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) would listen. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
		* `image` - (Required when environment_configuration_type=OCIR_CONTAINER) (Updatable) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
		* `image_digest` - (Applicable when environment_configuration_type=OCIR_CONTAINER) (Updatable) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
		* `server_port` - (Applicable when environment_configuration_type=OCIR_CONTAINER) (Updatable) The port on which the web server serving the inference is running. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
	* `model_configuration_details` - (Required) (Updatable) The model configuration details.
		* `bandwidth_mbps` - (Optional) (Updatable) The network bandwidth for the model.
		* `instance_configuration` - (Required) (Updatable) The model deployment instance configuration
			* `instance_shape_name` - (Required) (Updatable) The shape used to launch the model deployment instances.
			* `model_deployment_instance_shape_config_details` - (Optional) (Updatable) Details for the model-deployment instance shape configuration.
				* `memory_in_gbs` - (Optional) (Updatable) A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the memory to be specified with in the range of 6 to 1024 GB. VM.Standard3.Flex memory range is between 6 to 512 GB and VM.Optimized3.Flex memory range is between 6 to 256 GB. 
				* `ocpus` - (Optional) (Updatable) A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the ocpu count to be specified with in the range of 1 to 64 ocpu. VM.Standard3.Flex OCPU range is between 1 to 32 ocpu and for VM.Optimized3.Flex OCPU range is 1 to 18 ocpu. 
		* `model_id` - (Required) (Updatable) The OCID of the model you want to deploy.
		* `scaling_policy` - (Optional) (Updatable) The scaling policy to apply to each model of the deployment.
			* `instance_count` - (Required) (Updatable) The number of instances for the model deployment.
			* `policy_type` - (Required) (Updatable) The type of scaling policy.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model deployment.
* `state` - (Optional) (Updatable) The target state for the Model Deployment. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `category_log_details` - The log details for each category.
	* `access` - The log details.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
	* `predict` - The log details.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log group to work with.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a log to work with.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model deployment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `description` - A short description of the model deployment.
* `display_name` - A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My ModelDeployment`
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
* `lifecycle_details` - Details about the state of the model deployment.
* `model_deployment_configuration_details` - The model deployment configuration details.
	* `deployment_type` - The type of the model deployment.
	* `environment_configuration_details` - The configuration to carry the environment details thats used in Model Deployment creation
		* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
		* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
		* `environment_configuration_type` - The environment configuration type
		* `environment_variables` - Environment variables to set for the web server container. The size of envVars must be less than 2048 bytes. Key should be under 32 characters. Key should contain only letters, digits and underscore (_) Key should start with a letter. Key should have at least 2 characters. Key should not end with underscore eg. `TEST_` Key if added cannot be empty. Value can be empty. No specific size limits on individual Values. But overall environment variables is limited to 2048 bytes. Key can't be reserved Model Deployment environment variables. 
		* `health_check_port` - The port on which the container [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) would listen. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
		* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
		* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
		* `server_port` - The port on which the web server serving the inference is running. The port can be anything between `1024` and `65535`. The following ports cannot be used `24224`, `8446`, `8447`. 
	* `model_configuration_details` - The model configuration details.
		* `bandwidth_mbps` - The network bandwidth for the model.
		* `instance_configuration` - The model deployment instance configuration
			* `instance_shape_name` - The shape used to launch the model deployment instances.
			* `model_deployment_instance_shape_config_details` - Details for the model-deployment instance shape configuration.
				* `memory_in_gbs` - A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the memory to be specified with in the range of 6 to 1024 GB. VM.Standard3.Flex memory range is between 6 to 512 GB and VM.Optimized3.Flex memory range is between 6 to 256 GB. 
				* `ocpus` - A model-deployment instance of type VM.Standard.E3.Flex or VM.Standard.E4.Flex allows the ocpu count to be specified with in the range of 1 to 64 ocpu. VM.Standard3.Flex OCPU range is between 1 to 32 ocpu and for VM.Optimized3.Flex OCPU range is 1 to 18 ocpu. 
		* `model_id` - The OCID of the model you want to deploy.
		* `scaling_policy` - The scaling policy to apply to each model of the deployment.
			* `instance_count` - The number of instances for the model deployment.
			* `policy_type` - The type of scaling policy.
* `model_deployment_url` - The URL to interact with the model deployment.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
* `state` - The state of the model deployment.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Deployment
	* `update` - (Defaults to 20 minutes), when updating the Model Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Model Deployment


## Import

ModelDeployments can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_deployment.test_model_deployment "id"
```
