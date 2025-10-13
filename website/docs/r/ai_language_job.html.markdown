---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_job"
sidebar_current: "docs-oci-resource-ai_language-job"
description: |-
  Provides the Job resource in Oracle Cloud Infrastructure Ai Language service
---

# oci_ai_language_job
This resource provides the Job resource in Oracle Cloud Infrastructure Ai Language service.

Creates a new language service async job.


## Example Usage

```hcl
resource "oci_ai_language_job" "test_job" {
	#Required
	compartment_id = var.compartment_id
	input_location {
		#Required
		bucket = var.job_input_location_bucket
		location_type = var.job_input_location_location_type
		namespace = var.job_input_location_namespace

		#Optional
		object_names = var.job_input_location_object_names
		prefix = var.job_input_location_prefix
	}
	model_metadata_details {

		#Optional
		configuration {

			#Optional
			configuration_map = var.job_model_metadata_details_configuration_configuration_map
		}
		endpoint_id = oci_ai_language_endpoint.test_endpoint.id
		language_code = var.job_model_metadata_details_language_code
		model_id = oci_ai_language_model.test_model.id
		model_type = var.job_model_metadata_details_model_type
	}
	output_location {
		#Required
		bucket = var.job_output_location_bucket
		namespace = var.job_output_location_namespace

		#Optional
		prefix = var.job_output_location_prefix
	}

	#Optional
	description = var.job_description
	display_name = var.job_display_name
	input_configuration {

		#Optional
		configuration {

			#Optional
			config = var.job_input_configuration_configuration_config
		}
		document_types = var.job_input_configuration_document_types
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
* `description` - (Optional) (Updatable) A short description of the job.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the job.
* `input_configuration` - (Optional) input documents configuration by default TXT files will be processed and this behaviour will not change in future after adding new types 
	* `configuration` - (Optional) meta data about documents For CSV valid JSON format is {"CSV" :{inputColumn: "reviewDetails", rowId: "reviewId", copyColumnsToOutput: ["reviewId" "userId"] , delimiter: ","} Note: In future if new file types added we will update here in documentation about input file meta data 
		* `config` - (Optional) meta data about documents For CSV valid JSON format is {"CSV" :{inputColumn: "reviewDetails", rowId: "reviewId", copyColumnsToOutput: ["reviewId" "userId"] , delimiter: ","} Note: In future if new file types added we will update here in documentation about input file meta data 
	* `document_types` - (Optional) Type of documents supported for this release only TXT,CSV  and one element is allowed here. for future scope this is marked as list 
* `input_location` - (Required) document location and other meta data about documents For TXT only ObjectStoragePrefixLocation supported For CSV only ObjectStorageFileNameLocation is supported For this release only one file is supported for ObjectStorageFileNameLocation i.e CSV file type 
	* `bucket` - (Required) Object Storage bucket name.
	* `location_type` - (Required) locationType 
	* `namespace` - (Required) Object Storage namespace name.
	* `object_names` - (Required when location_type=OBJECT_STORAGE_FILE_LIST) List of objects to be processed
	* `prefix` - (Applicable when location_type=OBJECT_STORAGE_PREFIX) The prefix (directory) in an Object Storage bucket.
* `model_metadata_details` - (Required) training model details For this release only one model is allowed to be input here. One of the three modelType, ModelId, endpointId should be given other wise error will be thrown from API 
	* `configuration` - (Optional) model configuration details For PII :  < ENTITY_TYPE , ConfigurationDetails> ex."ORACLE":{ "mode" : "MASK","maskingCharacter" : "&","leaveCharactersUnmasked": 3,"isUnmaskedFromEnd" : true  } For language translation : { "targetLanguageCodes" : ConfigurationDetails} 
		* `configuration_map` - (Optional) model configuration details For PII : ConfigurationDetails will be PiiEntityMasking can be anyone of the following ex.{ "mode" : "MASK","maskingCharacter" : "&","leaveCharactersUnmasked": 3,"isUnmaskedFromEnd" : true  } { "mode" : "MASK","replaceWith" : "&"  } { "mode" : "REPLACE" } For language translation :  { "languageCodes" : ["cs", "ar"]} Language code supported Automatically detect language - auto Arabic - ar Brazilian Portuguese -  pt-BR Czech - cs Danish - da Dutch - nl English - en Finnish - fi French - fr Canadian French - fr-CA German - de Italian - it Japanese - ja Korean - ko Norwegian - no Polish - pl Romanian - ro Simplified Chinese - zh-CN Spanish - es Swedish - sv Traditional Chinese - zh-TW Turkish - tr Greek - el Hebrew - he 
	* `endpoint_id` - (Optional) Unique identifier endpoint OCID that should be used for inference
	* `language_code` - (Optional) Language code supported
		* auto : Automatically detect language
		* ar : Arabic
		* pt-BR : Brazilian Portuguese
		* cs : Czech
		* da : Danish
		* nl : Dutch
		* en : English
		* fi : Finnish
		* fr : French
		* fr-CA : Canadian French
		* de : German
		* it : Italian
		* ja : Japanese
		* ko : Korean
		* no : Norwegian
		* pl : Polish
		* ro : Romanian
		* zh-CN : Simplified Chinese
		* es : Spanish
		* sv : Swedish
		* zh-TW : Traditional Chinese
		* tr : Turkish
		* el : Greek
		* he : Hebrew 
	* `model_id` - (Optional) Unique identifier model OCID that should be used for inference
	* `model_type` - (Optional) model type to used for inference allowed values are
		* LANGUAGE_SENTIMENT_ANALYSIS
		* LANGUAGE_DETECTION
		* TEXT_CLASSIFICATION
		* NAMED_ENTITY_RECOGNITION
		* KEY_PHRASE_EXTRACTION
		* LANGUAGE_PII_ENTITIES
		* LANGUAGE_TRANSLATION 
* `output_location` - (Required) Object storage output location to write inference results
	* `bucket` - (Required) Object Storage bucket name.
	* `namespace` - (Required) Object Storage namespace name.
	* `prefix` - (Optional) The prefix (directory) in an Object Storage bucket.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
* `completed_documents` - Number of documents processed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job.
* `description` - A short description of the job.
* `display_name` - A user-friendly display name for the job.
* `failed_documents` - Number of documents failed for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
* `input_configuration` - input documents configuration by default TXT files will be processed and this behaviour will not change in future after adding new types 
	* `configuration` - meta data about documents For CSV valid JSON format is {"CSV" :{inputColumn: "reviewDetails", rowId: "reviewId", copyColumnsToOutput: ["reviewId" "userId"] , delimiter: ","} Note: In future if new file types added we will update here in documentation about input file meta data 
		* `config` - meta data about documents For CSV valid JSON format is {"CSV" :{inputColumn: "reviewDetails", rowId: "reviewId", copyColumnsToOutput: ["reviewId" "userId"] , delimiter: ","} Note: In future if new file types added we will update here in documentation about input file meta data 
	* `document_types` - Type of documents supported for this release only TXT,CSV  and one element is allowed here. for future scope this is marked as list 
* `input_location` - document location and other meta data about documents For TXT only ObjectStoragePrefixLocation supported For CSV only ObjectStorageFileNameLocation is supported For this release only one file is supported for ObjectStorageFileNameLocation i.e CSV file type 
	* `bucket` - Object Storage bucket name.
	* `location_type` - locationType 
	* `namespace` - Object Storage namespace name.
	* `object_names` - List of objects to be processed
	* `prefix` - The prefix (directory) in an Object Storage bucket.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `model_metadata_details` - training model details For this release only one model is allowed to be input here. One of the three modelType, ModelId, endpointId should be given other wise error will be thrown from API 
	* `configuration` - model configuration details For PII :  < ENTITY_TYPE , ConfigurationDetails> ex."ORACLE":{ "mode" : "MASK","maskingCharacter" : "&","leaveCharactersUnmasked": 3,"isUnmaskedFromEnd" : true  } For language translation : { "targetLanguageCodes" : ConfigurationDetails} 
		* `configuration_map` - model configuration details For PII : ConfigurationDetails will be PiiEntityMasking can be anyone of the following ex.{ "mode" : "MASK","maskingCharacter" : "&","leaveCharactersUnmasked": 3,"isUnmaskedFromEnd" : true  } { "mode" : "MASK","replaceWith" : "&"  } { "mode" : "REPLACE" } For language translation :  { "languageCodes" : ["cs", "ar"]} Language code supported Automatically detect language - auto Arabic - ar Brazilian Portuguese -  pt-BR Czech - cs Danish - da Dutch - nl English - en Finnish - fi French - fr Canadian French - fr-CA German - de Italian - it Japanese - ja Korean - ko Norwegian - no Polish - pl Romanian - ro Simplified Chinese - zh-CN Spanish - es Swedish - sv Traditional Chinese - zh-TW Turkish - tr Greek - el Hebrew - he 
	* `endpoint_id` - Unique identifier endpoint OCID that should be used for inference
	* `language_code` - Language code supported
		* auto : Automatically detect language
		* ar : Arabic
		* pt-BR : Brazilian Portuguese
		* cs : Czech
		* da : Danish
		* nl : Dutch
		* en : English
		* fi : Finnish
		* fr : French
		* fr-CA : Canadian French
		* de : German
		* it : Italian
		* ja : Japanese
		* ko : Korean
		* no : Norwegian
		* pl : Polish
		* ro : Romanian
		* zh-CN : Simplified Chinese
		* es : Spanish
		* sv : Swedish
		* zh-TW : Traditional Chinese
		* tr : Turkish
		* el : Greek
		* he : Hebrew 
	* `model_id` - Unique identifier model OCID that should be used for inference
	* `model_type` - model type to used for inference allowed values are
		* LANGUAGE_SENTIMENT_ANALYSIS
		* LANGUAGE_DETECTION
		* TEXT_CLASSIFICATION
		* NAMED_ENTITY_RECOGNITION
		* KEY_PHRASE_EXTRACTION
		* LANGUAGE_PII_ENTITIES
		* LANGUAGE_TRANSLATION 
* `output_location` - Object storage output location to write inference results
	* `bucket` - Object Storage bucket name.
	* `namespace` - Object Storage namespace name.
	* `prefix` - The prefix (directory) in an Object Storage bucket.
* `pending_documents` - Number of documents still to process. For CSV this signifies number of rows and for TXT this signifies number of files.
* `percent_complete` - How much progress the operation has made, vs the total amount of work that must be performed.
* `state` - The current state of the Job.
* `time_accepted` - Job accepted time.
* `time_completed` - Job finished time.
* `time_started` - Job started time.
* `total_documents` - Total number of documents given as input for prediction. For CSV this signifies number of rows and for TXT this signifies number of files.
* `ttl_in_days` - Time to live duration in days for Job. Job will be available till max 90 days.
* `warnings_count` - warnings count

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Job
	* `update` - (Defaults to 20 minutes), when updating the Job
	* `delete` - (Defaults to 20 minutes), when destroying the Job


## Import

Jobs can be imported using the `id`, e.g.

```
$ terraform import oci_ai_language_job.test_job "id"
```

