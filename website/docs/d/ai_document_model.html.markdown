---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_model"
sidebar_current: "docs-oci-datasource-ai_document-model"
description: |-
  Provides details about a specific Model in Oracle Cloud Infrastructure Ai Document service
---

# Data Source: oci_ai_document_model
This data source provides details about a specific Model resource in Oracle Cloud Infrastructure Ai Document service.

Get a model by identifier.

## Example Usage

```hcl
data "oci_ai_document_model" "test_model" {
	#Required
	model_id = oci_ai_document_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) A unique model identifier.


## Attributes Reference

The following attributes are exported:

* `alias_name` - the alias name of the model.
* `compartment_id` - The compartment identifier.
* `component_models` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) collection of active custom Key Value models that need to be composed.
	* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of active custom Key Value model that need to be composed.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - An optional description of the model.
* `display_name` - A human-friendly name for the model, which can be changed.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - A unique identifier that is immutable after creation.
* `is_composed_model` - Set to true when the model is created by using multiple key value extraction models.
* `is_quick_mode` - Set to true when experimenting with a new model type or dataset, so model training is quick, with a predefined low number of passes through the training data.
* `labels` - The collection of labels used to train the custom model.
* `lifecycle_details` - A message describing the current state in more detail, that can provide actionable information if training failed.
* `max_training_time_in_hours` - The maximum model training time in hours, expressed as a decimal fraction.
* `metrics` - Trained Model Metrics.
	* `dataset_summary` - Summary of count of samples used during model training.
		* `test_sample_count` - Number of samples used for testing the model.
		* `training_sample_count` - Number of samples used for training the model.
		* `validation_sample_count` - Number of samples used for validating the model.
	* `label_metrics_report` - List of metrics entries per label.
		* `confidence_entries` - List of document classification confidence report.
			* `accuracy` - accuracy under the threshold
			* `f1score` - f1Score under the threshold
			* `precision` - Precision under the threshold
			* `recall` - Recall under the threshold
			* `threshold` - Threshold used to calculate precision and recall.
		* `document_count` - Total test documents in the label.
		* `label` - Label name
		* `mean_average_precision` - Mean average precision under different thresholds
	* `model_type` - The type of custom model trained.
	* `overall_metrics_report` - Overall Metrics report for Document Classification Model.
		* `confidence_entries` - List of document classification confidence report.
			* `accuracy` - accuracy under the threshold
			* `f1score` - f1Score under the threshold
			* `precision` - Precision under the threshold
			* `recall` - Recall under the threshold
			* `threshold` - Threshold used to calculate precision and recall.
		* `document_count` - Total test documents in the label.
		* `mean_average_precision` - Mean average precision under different thresholds
* `model_type` - The type of the Document model.
* `model_version` - The version of the model.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project that contains the model.
* `state` - The current state of the model.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `tenancy_id` - The tenancy id of the model.
* `testing_dataset` - The base entity which is the input for creating and training a model.
	* `bucket` - The name of the Object Storage bucket that contains the input data file.
	* `dataset_id` - OCID of the Data Labeling dataset.
	* `dataset_type` - The dataset type, based on where it is stored.
	* `namespace` - The namespace name of the Object Storage bucket that contains the input data file.
	* `object` - The object name of the input data file.
* `time_created` - When the model was created, as an RFC3339 datetime string.
* `time_updated` - When the model was updated, as an RFC3339 datetime string.
* `trained_time_in_hours` - The total hours actually used for model training.
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

