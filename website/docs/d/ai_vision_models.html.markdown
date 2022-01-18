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

Returns a list of Models.


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
* `id` - (Optional) unique Model identifier
* `project_id` - (Optional) The ID of the project for which to list the objects.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `model_collection` - The list of model_collection.

### Model Reference

The following attributes are exported:

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
	* `namespace` - The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `time_created` - The time the Model was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Model was updated. An RFC3339 formatted datetime string
* `total_image_count` - Total number of training Images
* `trained_duration_in_hours` - Total hours actually used for training
* `training_dataset` - The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - Type of the Dataset.
	* `namespace` - The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `validation_dataset` - The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - Type of the Dataset.
	* `namespace` - The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - The object name of the input data file.

