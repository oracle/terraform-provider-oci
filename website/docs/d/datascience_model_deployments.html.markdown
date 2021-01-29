---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployments"
sidebar_current: "docs-oci-datasource-datascience-model_deployments"
description: |-
  Provides the list of Model Deployments in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_deployments
This data source provides the list of Model Deployments in Oracle Cloud Infrastructure Datascience service.

Lists all model deployments in the specified compartment. Only one parameter other than compartmentId may also be included in a query. The query must include compartmentId. If the query does not include compartmentId, or includes compartmentId but two or more other parameters an error is returned.


## Example Usage

```hcl
data "oci_datascience_model_deployments" "test_model_deployments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.model_deployment_created_by
	display_name = var.model_deployment_display_name
	id = var.model_deployment_id
	project_id = oci_datascience_project.test_project.id
	state = var.model_deployment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `model_deployments` - The list of model_deployments.

### ModelDeployment Reference

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

