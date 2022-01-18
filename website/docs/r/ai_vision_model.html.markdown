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

Creates a new Model.


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
		namespace = var.model_training_dataset_namespace
		object = var.model_training_dataset_object
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = {"bar-key"= "value"}
	is_quick_mode = var.model_is_quick_mode
	max_training_duration_in_hours = var.model_max_training_duration_in_hours
	model_version = var.model_model_version
	testing_dataset {
		#Required
		dataset_type = var.model_testing_dataset_dataset_type

		#Optional
		bucket = var.model_testing_dataset_bucket
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		namespace = var.model_testing_dataset_namespace
		object = var.model_testing_dataset_object
	}
	validation_dataset {
		#Required
		dataset_type = var.model_validation_dataset_dataset_type

		#Optional
		bucket = var.model_validation_dataset_bucket
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		namespace = var.model_validation_dataset_namespace
		object = var.model_validation_dataset_object
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the Model.
* `display_name` - (Optional) (Updatable) Model Identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_quick_mode` - (Optional) If It's true, Training is set for recommended epochs needed for quick training.
* `max_training_duration_in_hours` - (Optional) The maximum duration in hours for which the training will run.
* `model_type` - (Required) The  type of the model.
* `model_version` - (Optional) Model version.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `testing_dataset` - (Optional) The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - (Applicable when dataset_type=OBJECT_STORAGE) The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - (Applicable when dataset_type=DATA_SCIENCE_LABELING) The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - (Required) Type of the Dataset.
	* `namespace` - (Applicable when dataset_type=OBJECT_STORAGE) The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - (Applicable when dataset_type=OBJECT_STORAGE) The object name of the input data file.
* `training_dataset` - (Required) The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - (Applicable when dataset_type=OBJECT_STORAGE) The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - (Applicable when dataset_type=DATA_SCIENCE_LABELING) The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - (Required) Type of the Dataset.
	* `namespace` - (Applicable when dataset_type=OBJECT_STORAGE) The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - (Applicable when dataset_type=OBJECT_STORAGE) The object name of the input data file.
* `validation_dataset` - (Optional) The base entity for a Dataset, which is the input for Model creation.
	* `bucket` - (Applicable when dataset_type=OBJECT_STORAGE) The name of the ObjectStorage bucket that contains the input data file.
	* `dataset_id` - (Applicable when dataset_type=DATA_SCIENCE_LABELING) The OCID of the Data Science Labeling Dataset.
	* `dataset_type` - (Required) Type of the Dataset.
	* `namespace` - (Applicable when dataset_type=OBJECT_STORAGE) The namespace name of the ObjectStorage bucket that contains the input data file.
	* `object` - (Applicable when dataset_type=OBJECT_STORAGE) The object name of the input data file.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_model.test_model "id"
```

