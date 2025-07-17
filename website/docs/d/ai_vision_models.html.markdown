---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_models"
sidebar_current: "docs-oci-datasource-ai_vision-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_models
This data source provides the list of Models in Oracle Cloud Infrastructure Ai Vision service.

Returns a list of models in a compartment.


## Example Usage

```hcl
data "oci_ai_vision_models" "test_models" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.model_display_name
	id = var.model_id
	project_id = oci_ai_vision_project.test_project.id
	state = var.model_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The filter to find the model with the given identifier.
* `project_id` - (Optional) The ID of the project for which to list the objects.
* `state` - (Optional) The filter to match models with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `model_collection` - The list of model_collection.

### Model Reference

The following attributes are exported:

* `average_precision` - The mean average precision of the trained model.
* `compartment_id` - The compartment identifier.
* `confidence_threshold` - The intersection over the union threshold used for calculating precision and recall.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - An optional description of the model.
* `display_name` - A human-friendly name for the model, which can be changed.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - A unique identifier that is immutable after creation.
* `is_quick_mode` - Set to true when experimenting with a new model type or dataset, so model training is quick, with a predefined low number of passes through the training data.
* `lifecycle_details` - A message describing the current state in more detail, that can provide actionable information if training failed.
* `max_training_duration_in_hours` - The maximum model training duration in hours, expressed as a decimal fraction.
* `metrics` - The complete set of per-label metrics for successfully trained models.
* `model_type` - What type of Vision model this is.
* `model_version` - The version of the model.
* `precision` - The precision of the trained model.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project that contains the model.
* `recall` - Recall of the trained model.
* `state` - The current state of the model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `test_image_count` - The number of images set aside for evaluating model performance metrics after training.
* `testing_dataset` - The base entity which is the input for creating and training a model.
	* `bucket` - The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - OCID of the Data Labeling dataset.
	* `dataset_type` - The dataset type, based on where it is stored.
	* `namespace` - The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `time_created` - When the model was created, as an RFC3339 datetime string.
* `time_updated` - When the model was updated, as an RFC3339 datetime string.
* `total_image_count` - The number of images in the dataset used to train, validate, and test the model.
* `trained_duration_in_hours` - The total hours actually used for model training.
* `training_dataset` - The base entity which is the input for creating and training a model.
	* `bucket` - The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - OCID of the Data Labeling dataset.
	* `dataset_type` - The dataset type, based on where it is stored.
	* `namespace` - The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `validation_dataset` - The base entity which is the input for creating and training a model.
	* `bucket` - The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - OCID of the Data Labeling dataset.
	* `dataset_type` - The dataset type, based on where it is stored.
	* `namespace` - The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - The object name of the input data file.

