---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model"
sidebar_current: "docs-oci-resource-datascience-model"
description: |-
  Provides the Model resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model
This resource provides the Model resource in Oracle Cloud Infrastructure Data Science service.

Creates a new model.

## Example Usage

```hcl
resource "oci_datascience_model" "test_model" {
	#Required
	compartment_id = var.compartment_id
	project_id = oci_datascience_project.test_project.id

	#Optional
	backup_setting {
		#Required
		backup_region = var.model_backup_setting_backup_region
		is_backup_enabled = var.model_backup_setting_is_backup_enabled

		#Optional
		customer_notification_type = var.model_backup_setting_customer_notification_type
	}
	custom_metadata_list {

		#Optional
		category = var.model_custom_metadata_list_category
		description = var.model_custom_metadata_list_description
		key = var.model_custom_metadata_list_key
		value = var.model_custom_metadata_list_value
	}
	defined_metadata_list {

		#Optional
		category = var.model_defined_metadata_list_category
		description = var.model_defined_metadata_list_description
		key = var.model_defined_metadata_list_key
		value = var.model_defined_metadata_list_value
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.model_description
	display_name = var.model_display_name
	freeform_tags = {"Department"= "Finance"}
	input_schema = var.model_input_schema
	output_schema = var.model_output_schema
	retention_setting {
		#Required
		archive_after_days = var.model_retention_setting_archive_after_days

		#Optional
		customer_notification_type = var.model_retention_setting_customer_notification_type
		delete_after_days = var.model_retention_setting_delete_after_days
	}
	version_label = var.model_version_label
}
```

## Argument Reference

The following arguments are supported:

* `backup_setting` - (Optional) (Updatable) Back up setting details of the model.
	* `backup_region` - (Required) (Updatable) Oracle Cloud Infrastructure backup region for the model.
	* `customer_notification_type` - (Optional) (Updatable) Customer notification on backup success/failure events.
	* `is_backup_enabled` - (Required) (Updatable) Boolean flag representing whether backup needs to be enabled/disabled for the model.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the model in.
* `custom_metadata_list` - (Optional) (Updatable) An array of custom metadata details for the model.
	* `category` - (Optional) (Updatable) Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - (Optional) (Updatable) Description of model metadata
	* `key` - (Optional) (Updatable) Key of the model Metadata. The key can either be user defined or Oracle Cloud Infrastructure defined. List of Oracle Cloud Infrastructure defined keys:
		* useCaseType
		* libraryName
		* libraryVersion
		* estimatorClass
		* hyperParameters
		* testartifactresults 
	* `value` - (Optional) (Updatable) Allowed values for useCaseType: binary_classification, regression, multinomial_classification, clustering, recommender, dimensionality_reduction/representation, time_series_forecasting, anomaly_detection, topic_modeling, ner, sentiment_analysis, image_classification, object_localization, other

		Allowed values for libraryName: scikit-learn, xgboost, tensorflow, pytorch, mxnet, keras, lightGBM, pymc3, pyOD, spacy, prophet, sktime, statsmodels, cuml, oracle_automl, h2o, transformers, nltk, emcee, pystan, bert, gensim, flair, word2vec, ensemble, other 
* `defined_metadata_list` - (Optional) (Updatable) An array of defined metadata details for the model.
	* `category` - (Optional) (Updatable) Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - (Optional) (Updatable) Description of model metadata
	* `key` - (Optional) (Updatable) Key of the model Metadata. The key can either be user defined or Oracle Cloud Infrastructure defined. List of Oracle Cloud Infrastructure defined keys:
		* useCaseType
		* libraryName
		* libraryVersion
		* estimatorClass
		* hyperParameters
		* testartifactresults 
	* `value` - (Optional) (Updatable) Allowed values for useCaseType: binary_classification, regression, multinomial_classification, clustering, recommender, dimensionality_reduction/representation, time_series_forecasting, anomaly_detection, topic_modeling, ner, sentiment_analysis, image_classification, object_localization, other

		Allowed values for libraryName: scikit-learn, xgboost, tensorflow, pytorch, mxnet, keras, lightGBM, pymc3, pyOD, spacy, prophet, sktime, statsmodels, cuml, oracle_automl, h2o, transformers, nltk, emcee, pystan, bert, gensim, flair, word2vec, ensemble, other 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the model.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My Model` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `input_schema` - (Optional) Input schema file content in String format
* `output_schema` - (Optional) Output schema file content in String format
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the model.
* `retention_setting` - (Optional) (Updatable) Retention setting details of the model.
	* `archive_after_days` - (Required) (Updatable) Number of days after which the model will be archived.
	* `customer_notification_type` - (Optional) (Updatable) Customer notification options on success/failure of archival, deletion events.
	* `delete_after_days` - (Optional) (Updatable) Number of days after which the archived model will be deleted.
* `version_label` - (Optional) (Updatable) The version label can add an additional description of the lifecycle state of the model or the application using/training the model.
* `model_artifact` - (Optional) The model artifact to upload. It is a ZIP archive of the files necessary to run the model. This can be done in a separate step or using cli/sdk. The Model will remain in "Creating" state until its artifact is uploaded.
* `artifact_content_disposition` - (Optional) This allows to specify a filename during upload. This file name is used to dispose of the file contents while downloading the file. Example: `attachment; filename=model-artifact.zip`
* `artifact_content_length` - (Optional, Required if `model_artifact` is set) The content length of the model_artifact.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model
	* `update` - (Defaults to 20 minutes), when updating the Model
	* `delete` - (Defaults to 20 minutes), when destroying the Model


## Import

Models can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model.test_model "id"
```

