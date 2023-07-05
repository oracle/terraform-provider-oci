---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_models"
sidebar_current: "docs-oci-datasource-ai_language-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Ai Language service
---

# Data Source: oci_ai_language_models
This data source provides the list of Models in Oracle Cloud Infrastructure Ai Language service.

Returns a list of models.


## Example Usage

```hcl
data "oci_ai_language_models" "test_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.model_display_name
	model_id = oci_ai_language_model.test_model.id
	project_id = oci_ai_language_project.test_project.id
	state = var.model_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `model_id` - (Optional) unique model OCID.
* `project_id` - (Optional) The ID of the project for which to list the objects.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `model_collection` - The list of model_collection.

### Model Reference

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

