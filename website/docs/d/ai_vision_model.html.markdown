---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_model"
sidebar_current: "docs-oci-datasource-ai_vision-model"
description: |-
  Provides details about a specific Model in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_model
This data source provides details about a specific Model resource in Oracle Cloud Infrastructure Ai Vision service.

Get a model by identifier.

## Example Usage

```hcl
data "oci_ai_vision_model" "test_model" {
	#Required
	model_id = oci_ai_vision_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) A unique model identifier.


## Attributes Reference

The following attributes are exported:

<<<<<<< ours
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
=======
* `average_precision` - Average precision of the trained model
* `compartment_id` - Compartment Identifier
* `confidence_threshold` - Confidence ratio of the calculation
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the model.
* `display_name` - Model Identifier, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `is_quick_mode` - If It's true, Training is set for recommended epochs needed for quick training.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `max_training_duration_in_hours` - The maximum duration in hours for which the training will run.
* `metrics` - Complete Training Metrics for successful trained model
* `model_type` - Type of the Model.
* `model_version` - The version of the model
* `precision` - Precision of the trained model
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `recall` - Recall of the trained model
* `state` - The current state of the Model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `test_image_count` - Total number of testing Images
* `testing_dataset` - The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - Type of the Dataset.
	* `namespace_name` - The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `time_created` - The time the Model was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Model was updated. An RFC3339 formatted datetime string
* `total_image_count` - Total number of training Images
* `trained_duration_in_hours` - Total hours actually used for training
* `training_dataset` - The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - Type of the Dataset.
	* `namespace_name` - The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `validation_dataset` - The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - Type of the Dataset.
	* `namespace_name` - The namespace name of the ObjectStorage bucket that contains the input data file.
>>>>>>> theirs
	* `object` - The object name of the input data file.

