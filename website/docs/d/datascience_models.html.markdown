---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_models"
sidebar_current: "docs-oci-datasource-datascience-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_models
This data source provides the list of Models in Oracle Cloud Infrastructure Data Science service.

Lists models in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_models" "test_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.model_created_by
	display_name = var.model_display_name
	id = var.model_id
	model_version_set_name = oci_datascience_model_version_set.test_model_version_set.name
	project_id = oci_datascience_project.test_project.id
	state = var.model_state
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

* `models` - The list of models.

### Model Reference

The following attributes are exported:

* `backup_operation_details` - Backup operation details of the model.
	* `backup_state` - The backup status of the model.
	* `backup_state_details` - The backup execution status details of the model.
	* `time_last_backup` - The last backup execution time of the model.
* `backup_setting` - Back up setting details of the model.
	* `backup_region` - Oracle Cloud Infrastructure backup region for the model.
	* `customer_notification_type` - Customer notification on backup success/failure events.
	* `is_backup_enabled` - Boolean flag representing whether backup needs to be enabled/disabled for the model.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model.
* `custom_metadata_list` - An array of custom metadata details for the model.
	* `category` - Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - Description of model metadata
	* `key` - Key of the model Metadata. The key can either be user defined or Oracle Cloud Infrastructure defined. List of Oracle Cloud Infrastructure defined keys:
		* useCaseType
		* libraryName
		* libraryVersion
		* estimatorClass
		* hyperParameters
		* testartifactresults 
	* `value` - Allowed values for useCaseType: binary_classification, regression, multinomial_classification, clustering, recommender, dimensionality_reduction/representation, time_series_forecasting, anomaly_detection, topic_modeling, ner, sentiment_analysis, image_classification, object_localization, other

		Allowed values for libraryName: scikit-learn, xgboost, tensorflow, pytorch, mxnet, keras, lightGBM, pymc3, pyOD, spacy, prophet, sktime, statsmodels, cuml, oracle_automl, h2o, transformers, nltk, emcee, pystan, bert, gensim, flair, word2vec, ensemble, other 
* `defined_metadata_list` - An array of defined metadata details for the model.
	* `category` - Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - Description of model metadata
	* `key` - Key of the model Metadata. The key can either be user defined or Oracle Cloud Infrastructure defined. List of Oracle Cloud Infrastructure defined keys:
		* useCaseType
		* libraryName
		* libraryVersion
		* estimatorClass
		* hyperParameters
		* testartifactresults 
	* `value` - Allowed values for useCaseType: binary_classification, regression, multinomial_classification, clustering, recommender, dimensionality_reduction/representation, time_series_forecasting, anomaly_detection, topic_modeling, ner, sentiment_analysis, image_classification, object_localization, other

		Allowed values for libraryName: scikit-learn, xgboost, tensorflow, pytorch, mxnet, keras, lightGBM, pymc3, pyOD, spacy, prophet, sktime, statsmodels, cuml, oracle_automl, h2o, transformers, nltk, emcee, pystan, bert, gensim, flair, word2vec, ensemble, other 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the model.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `input_schema` - Input schema file content in String format
* `lifecycle_details` - Details about the lifecycle state of the model.
* `model_version_set_id` - The OCID of the model version set that the model is associated to.
* `model_version_set_name` - The name of the model version set that the model is associated to.
* `output_schema` - Output schema file content in String format
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
* `retention_operation_details` - Retention operation details for the model.
	* `archive_state` - The archival status of model.
	* `archive_state_details` - The archival state details of the model.
	* `delete_state` - The deletion status of the archived model.
	* `delete_state_details` - The deletion status details of the archived model.
	* `time_archival_scheduled` - The estimated archival time of the model based on the provided retention setting.
	* `time_deletion_scheduled` - The estimated deletion time of the model based on the provided retention setting.
* `retention_setting` - Retention setting details of the model.
	* `archive_after_days` - Number of days after which the model will be archived.
	* `customer_notification_type` - Customer notification options on success/failure of archival, deletion events.
	* `delete_after_days` - Number of days after which the archived model will be deleted.
* `state` - The state of the model.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

