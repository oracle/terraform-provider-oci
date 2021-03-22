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
	* `model_configuration_details` - (Required) (Updatable) The model configuration details.
		* `bandwidth_mbps` - (Optional) (Updatable) The network bandwidth for the model.
		* `instance_configuration` - (Required) (Updatable) The model deployment instance configuration
			* `instance_shape_name` - (Required) (Updatable) The shape used to launch the model deployment instances.
		* `model_id` - (Required) (Updatable) The OCID of the model you want to deploy.
		* `scaling_policy` - (Optional) (Updatable) The scaling policy to apply to each model of the deployment.
			* `instance_count` - (Required) (Updatable) The number of instances for the model deployment.
			* `policy_type` - (Required) (Updatable) The type of scaling policy.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model deployment.


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
	* `model_configuration_details` - The model configuration details.
		* `bandwidth_mbps` - The network bandwidth for the model.
		* `instance_configuration` - The model deployment instance configuration
			* `instance_shape_name` - The shape used to launch the model deployment instances.
		* `model_id` - The OCID of the model you want to deploy.
		* `scaling_policy` - The scaling policy to apply to each model of the deployment.
			* `instance_count` - The number of instances for the model deployment.
			* `policy_type` - The type of scaling policy.
* `model_deployment_url` - The URL to interact with the model deployment.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
* `state` - The state of the model deployment.
* `time_created` - The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Deployment
	* `update` - (Defaults to 20 minutes), when updating the Model Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Model Deployment


## Import

ModelDeployments can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_deployment.test_model_deployment "id"
```

