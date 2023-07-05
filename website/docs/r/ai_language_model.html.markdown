---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_model"
sidebar_current: "docs-oci-resource-ai_language-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Ai Language service
---

# oci_ai_language_model
This resource provides the Model resource in Oracle Cloud Infrastructure Ai Language service.

Creates a new model for training and train the model with date provided.


## Example Usage

```hcl
resource "oci_ai_language_model" "test_model" {
	#Required
	compartment_id = var.compartment_id
	model_details {
		#Required
		model_type = var.model_model_details_model_type

		#Optional
		classification_mode {
			#Required
			classification_mode = var.model_model_details_classification_mode_classification_mode

			#Optional
			version = var.model_model_details_classification_mode_version
		}
		language_code = var.model_model_details_language_code
		version = var.model_model_details_version
	}
	project_id = oci_ai_language_project.test_project.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = {"bar-key"= "value"}
	test_strategy {
		#Required
		strategy_type = var.model_test_strategy_strategy_type
		testing_dataset {
			#Required
			dataset_type = var.model_test_strategy_testing_dataset_dataset_type

			#Optional
			dataset_id = oci_data_labeling_service_dataset.test_dataset.id
			location_details {
				#Required
				bucket = var.model_test_strategy_testing_dataset_location_details_bucket
				location_type = var.model_test_strategy_testing_dataset_location_details_location_type
				namespace = var.model_test_strategy_testing_dataset_location_details_namespace
				object_names = var.model_test_strategy_testing_dataset_location_details_object_names
			}
		}

		#Optional
		validation_dataset {
			#Required
			dataset_type = var.model_test_strategy_validation_dataset_dataset_type

			#Optional
			dataset_id = oci_data_labeling_service_dataset.test_dataset.id
			location_details {
				#Required
				bucket = var.model_test_strategy_validation_dataset_location_details_bucket
				location_type = var.model_test_strategy_validation_dataset_location_details_location_type
				namespace = var.model_test_strategy_validation_dataset_location_details_namespace
				object_names = var.model_test_strategy_validation_dataset_location_details_object_names
			}
		}
	}
	training_dataset {
		#Required
		dataset_type = var.model_training_dataset_dataset_type

		#Optional
		dataset_id = oci_data_labeling_service_dataset.test_dataset.id
		location_details {
			#Required
			bucket = var.model_training_dataset_location_details_bucket
			location_type = var.model_training_dataset_location_details_location_type
			namespace = var.model_training_dataset_location_details_namespace
			object_names = var.model_training_dataset_location_details_object_names
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)  for the models compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the a model.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `model_details` - (Required) Possible model types
	* `classification_mode` - (Applicable when model_type=TEXT_CLASSIFICATION) possible text classification modes
		* `classification_mode` - (Required) classification Modes
		* `version` - (Optional) Optional if nothing specified latest base model will be used for training. Supported versions can be found at /modelTypes/{modelType}
	* `language_code` - (Optional) supported language default value is en
	* `model_type` - (Required) Model type
	* `version` - (Applicable when model_type=NAMED_ENTITY_RECOGNITION | PRE_TRAINED_HEALTH_NLU | PRE_TRAINED_KEYPHRASE_EXTRACTION | PRE_TRAINED_LANGUAGE_DETECTION | PRE_TRAINED_NAMED_ENTITY_RECOGNITION | PRE_TRAINED_PHI | PRE_TRAINED_PII | PRE_TRAINED_SENTIMENT_ANALYSIS | PRE_TRAINED_SUMMARIZATION | PRE_TRAINED_TEXT_CLASSIFICATION | PRE_TRAINED_UNIVERSAL) Optional pre trained model version. if nothing specified latest pre trained model will be used.  Supported versions can be found at /modelTypes/{modelType} 
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `test_strategy` - (Optional) Possible strategy as testing and validation(optional) dataset. 
	* `strategy_type` - (Required) This information will define the test strategy different datasets for test and validation(optional) dataset. 
	* `testing_dataset` - (Required) Possible data set type
		* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) Data Science Labelling Service OCID
		* `dataset_type` - (Required) Possible data sets
		* `location_details` - (Required when dataset_type=OBJECT_STORAGE) Possible object storage location types
			* `bucket` - (Required) Object storage bucket name
			* `location_type` - (Required) Possible object storage location types
			* `namespace` - (Required) Object storage namespace
			* `object_names` - (Required) Array of files which need to be processed in the bucket
	* `validation_dataset` - (Optional) Possible data set type
		* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) Data Science Labelling Service OCID
		* `dataset_type` - (Required) Possible data sets
		* `location_details` - (Required when dataset_type=OBJECT_STORAGE) Possible object storage location types
			* `bucket` - (Required) Object storage bucket name
			* `location_type` - (Required) Possible object storage location types
			* `namespace` - (Required) Object storage namespace
			* `object_names` - (Required) Array of files which need to be processed in the bucket
* `training_dataset` - (Optional) Possible data set type
	* `dataset_id` - (Required when dataset_type=DATA_SCIENCE_LABELING) Data Science Labelling Service OCID
	* `dataset_type` - (Required) Possible data sets
	* `location_details` - (Required when dataset_type=OBJECT_STORAGE) Possible object storage location types
		* `bucket` - (Required) Object storage bucket name
		* `location_type` - (Required) Possible object storage location types
		* `namespace` - (Required) Object storage namespace
		* `object_names` - (Required) Array of files which need to be processed in the bucket


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)  for the model's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the Model.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `evaluation_results` - model training results of different models
	* `class_metrics` - List of text classification metrics
		* `f1` - F1-score, is a measure of a model’s accuracy on a dataset
		* `label` - Text classification label
		* `precision` - Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
		* `recall` - Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
		* `support` - number of samples in the test set
	* `confusion_matrix` - class level confusion matrix
		* `matrix` - confusion matrix data
	* `entity_metrics` - List of entity metrics
		* `f1` - F1-score, is a measure of a model’s accuracy on a dataset
		* `label` - Entity label
		* `precision` - Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
		* `recall` - Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
	* `labels` - labels
	* `metrics` - Model level named entity recognition metrics
		* `accuracy` - The fraction of the labels that were correctly recognised .
		* `macro_f1` - F1-score, is a measure of a model’s accuracy on a dataset
		* `macro_precision` - Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
		* `macro_recall` - Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
		* `micro_f1` - F1-score, is a measure of a model’s accuracy on a dataset
		* `micro_precision` - Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
		* `micro_recall` - Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
		* `weighted_f1` - F1-score, is a measure of a model’s accuracy on a dataset
		* `weighted_precision` - Precision refers to the number of true positives divided by the total number of positive predictions (i.e., the number of true positives plus the number of false positives)
		* `weighted_recall` - Measures the model's ability to predict actual positive classes. It is the ratio between the predicted true positives and what was actually tagged. The recall metric reveals how many of the predicted classes are correct.
	* `model_type` - Model type
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier model OCID of a model that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
* `model_details` - Possible model types
	* `classification_mode` - possible text classification modes
		* `classification_mode` - classification Modes
		* `version` - Optional if nothing specified latest base model will be used for training. Supported versions can be found at /modelTypes/{modelType}
	* `language_code` - supported language default value is en
	* `model_type` - Model type
	* `version` - Optional pre trained model version. if nothing specified latest pre trained model will be used.  Supported versions can be found at /modelTypes/{modelType} 
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `state` - The state of the model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `test_strategy` - Possible strategy as testing and validation(optional) dataset. 
	* `strategy_type` - This information will define the test strategy different datasets for test and validation(optional) dataset. 
	* `testing_dataset` - Possible data set type
		* `dataset_id` - Data Science Labelling Service OCID
		* `dataset_type` - Possible data sets
		* `location_details` - Possible object storage location types
			* `bucket` - Object storage bucket name
			* `location_type` - Possible object storage location types
			* `namespace` - Object storage namespace
			* `object_names` - Array of files which need to be processed in the bucket
	* `validation_dataset` - Possible data set type
		* `dataset_id` - Data Science Labelling Service OCID
		* `dataset_type` - Possible data sets
		* `location_details` - Possible object storage location types
			* `bucket` - Object storage bucket name
			* `location_type` - Possible object storage location types
			* `namespace` - Object storage namespace
			* `object_names` - Array of files which need to be processed in the bucket
* `time_created` - The time the the model was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the model was updated. An RFC3339 formatted datetime string.
* `training_dataset` - Possible data set type
	* `dataset_id` - Data Science Labelling Service OCID
	* `dataset_type` - Possible data sets
	* `location_details` - Possible object storage location types
		* `bucket` - Object storage bucket name
		* `location_type` - Possible object storage location types
		* `namespace` - Object storage namespace
		* `object_names` - Array of files which need to be processed in the bucket
* `version` - For pre trained models this will identify model type version used for model creation For custom identifying the model by model id is difficult. This param provides ease of use for end customer. <<service>>::<<service-name>>_<<model-type-version>>::<<custom model on which this training has to be done>> ex: ai-lang::NER_V1::CUSTOM-V0 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_ai_language_model.test_model "id"
```

