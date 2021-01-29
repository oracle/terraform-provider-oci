---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployment"
sidebar_current: "docs-oci-datasource-datascience-model_deployment"
description: |-
  Provides details about a specific Model Deployment in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_deployment
This data source provides details about a specific Model Deployment resource in Oracle Cloud Infrastructure Datascience service.

Retrieves the model deployment for the specified `modelDeploymentId`.

## Example Usage

```hcl
data "oci_datascience_model_deployment" "test_model_deployment" {
	#Required
	model_deployment_id = oci_datascience_model_deployment.test_model_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `model_deployment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.


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

