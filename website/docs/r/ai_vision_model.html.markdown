---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_model"
sidebar_current: "docs-oci-resource-ai_vision-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Ai Vision service
---

# oci_ai_vision_model
This resource provides the Model resource in Oracle Cloud Infrastructure Ai Vision service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/vision/latest/Model

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/aiVision

Create a new model.


## Example Usage

```hcl
resource "oci_ai_vision_model" "test_model" {
	#Required
	compartment_id = var.compartment_id
	model_type = var.model_model_type
	project_id = oci_ai_vision_project.test_project.id
	training_dataset {
		#Required
		dataset_type = var.model_training_dataset_dataset_type

		#Optional
		bucket = var.model_training_dataset_bucket
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		namespace_name = var.model_training_dataset_namespace
		object = var.model_training_dataset_object
	}

	#Optional
	defined_tags = var.model_defined_tags
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = var.model_freeform_tags
	is_quick_mode = var.model_is_quick_mode
	max_training_duration_in_hours = var.model_max_training_duration_in_hours
	model_version = var.model_model_version
	testing_dataset {
		#Required
		dataset_type = var.model_testing_dataset_dataset_type

		#Optional
		bucket = var.model_testing_dataset_bucket
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		namespace_name = var.model_testing_dataset_namespace
		object = var.model_testing_dataset_object
	}
	validation_dataset {
		#Required
		dataset_type = var.model_validation_dataset_dataset_type

		#Optional
		bucket = var.model_validation_dataset_bucket
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		namespace_name = var.model_validation_dataset_namespace
		object = var.model_validation_dataset_object
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - (Optional) (Updatable) An optional description of the model.
* `display_name` - (Optional) (Updatable) A human-friendly name for the model, which can be changed.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `is_quick_mode` - (Optional) Set to true when experimenting with a new model type or dataset, so the model training is quick, with a predefined low number of passes through the training data.
* `max_training_duration_in_hours` - (Optional) The maximum model training duration in hours, expressed as a decimal fraction.
* `model_type` - (Required) Which type of Vision model this is.
* `model_version` - (Optional) The model version
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project that contains the model.
* `testing_dataset` - (Optional) The base entity which is the input for creating and training a model.
	* `bucket` - (Required when dataset_type=OBJECT_STORAGE) The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) OCID of the Data Labeling dataset.
	* `dataset_type` - (Required) The dataset type, based on where it is stored.
	* `namespace` - (Required when dataset_type=OBJECT_STORAGE) The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - (Required when dataset_type=OBJECT_STORAGE) The object name of the input data file.
* `training_dataset` - (Required) The base entity which is the input for creating and training a model.
	* `bucket` - (Required when dataset_type=OBJECT_STORAGE) The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) OCID of the Data Labeling dataset.
	* `dataset_type` - (Required) The dataset type, based on where it is stored.
	* `namespace` - (Required when dataset_type=OBJECT_STORAGE) The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - (Required when dataset_type=OBJECT_STORAGE) The object name of the input data file.
* `validation_dataset` - (Optional) The base entity which is the input for creating and training a model.
	* `bucket` - (Required when dataset_type=OBJECT_STORAGE) The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) OCID of the Data Labeling dataset.
	* `dataset_type` - (Required) The dataset type, based on where it is stored.
	* `namespace` - (Required when dataset_type=OBJECT_STORAGE) The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - (Required when dataset_type=OBJECT_STORAGE) The object name of the input data file.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_model.test_model "id"
```

