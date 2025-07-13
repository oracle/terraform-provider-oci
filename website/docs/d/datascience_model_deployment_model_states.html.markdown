---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_deployment_model_states"
sidebar_current: "docs-oci-datasource-datascience-model_deployment_model_states"
description: |-
  Provides the list of Model Deployment Model States in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model_deployment_model_states
This data source provides the list of Model Deployment Model States in Oracle Cloud Infrastructure Data Science service.

Lists the status of models in a model group deployment.

## Example Usage

```hcl
data "oci_datascience_model_deployment_model_states" "test_model_deployment_model_states" {
	#Required
	compartment_id = var.compartment_id
	model_deployment_id = oci_datascience_model_deployment.test_model_deployment.id

	#Optional
	display_name = var.model_deployment_model_state_display_name
	inference_key = var.model_deployment_model_state_inference_key
	model_id = oci_datascience_model.test_model.id
	project_id = oci_datascience_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `inference_key` - (Optional) <b>Filter</b> results by the inference key.
* `model_deployment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
* `model_id` - (Optional) <b>Filter</b> results by the model ocid.
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.


## Attributes Reference

The following attributes are exported:

* `model_deployment_model_states` - The list of model_deployment_model_states.

### ModelDeploymentModelState Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `inference_key` - SaaS friendly name for the model OCID.
* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployed model in model deployment.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
* `state` - The state of the deployed model in model deployment.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

