---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_processor_job"
sidebar_current: "docs-oci-resource-ai_document-processor_job"
description: |-
  Provides the Processor Job resource in Oracle Cloud Infrastructure Ai Document service
---

# oci_ai_document_processor_job
This resource provides the Processor Job resource in Oracle Cloud Infrastructure Ai Document service.

Create a processor job for document analysis.


## Example Usage

```hcl
resource "oci_ai_document_processor_job" "test_processor_job" {
	#Required
	compartment_id = var.compartment_id
	input_location {
		#Required
		source_type = var.processor_job_input_location_source_type

		#Optional
		data = var.processor_job_input_location_data
		object_locations {

			#Optional
			bucket = var.processor_job_input_location_object_locations_bucket
			namespace = var.processor_job_input_location_object_locations_namespace
			object = var.processor_job_input_location_object_locations_object
		}
	}
	output_location {
		#Required
		bucket = var.processor_job_output_location_bucket
		namespace = var.processor_job_output_location_namespace
		prefix = var.processor_job_output_location_prefix
	}
	processor_config {
		#Required
		features {
			#Required
			feature_type = var.processor_job_processor_config_features_feature_type

			#Optional
			generate_searchable_pdf = var.processor_job_processor_config_features_generate_searchable_pdf
			max_results = var.processor_job_processor_config_features_max_results
			model_id = oci_ai_document_model.test_model.id
			tenancy_id = oci_identity_tenancy.test_tenancy.id
		}
		processor_type = var.processor_job_processor_config_processor_type

		#Optional
		document_type = var.processor_job_processor_config_document_type
		is_zip_output_enabled = var.processor_job_processor_config_is_zip_output_enabled
		language = var.processor_job_processor_config_language
	}

	#Optional
	display_name = var.processor_job_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment identifier.
* `display_name` - (Optional) The display name of the processor job.
* `input_location` - (Required) The location of the inputs.
	* `data` - (Required when source_type=INLINE_DOCUMENT_CONTENT) Raw document data with Base64 encoding.
	* `object_locations` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The list of ObjectLocations.
		* `bucket` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage bucket name.
		* `namespace` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage namespace name.
		* `object` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage object name.
	* `source_type` - (Required) The type of input location. The allowed values are:
		* `OBJECT_STORAGE_LOCATIONS`: A list of object locations in Object Storage.
		* `INLINE_DOCUMENT_CONTENT`: The content of an inline document. 
* `output_location` - (Required) The object storage location where to store analysis results.
	* `bucket` - (Required) The Object Storage bucket name.
	* `namespace` - (Required) The Object Storage namespace.
	* `prefix` - (Required) The Object Storage folder name.
* `processor_config` - (Required) The configuration of a processor.
	* `document_type` - (Optional) The document type.
	* `features` - (Required) The types of document analysis requested.
		* `feature_type` - (Required) The type of document analysis requested. The allowed values are:
			* `LANGUAGE_CLASSIFICATION`: Detect the language.
			* `TEXT_EXTRACTION`: Recognize text.
			* `TABLE_EXTRACTION`: Detect and extract data in tables.
			* `KEY_VALUE_EXTRACTION`: Extract form fields.
			* `DOCUMENT_CLASSIFICATION`: Identify the type of document. 
		* `generate_searchable_pdf` - (Applicable when feature_type=TEXT_EXTRACTION) Whether or not to generate a searchable PDF file.
		* `max_results` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | LANGUAGE_CLASSIFICATION) The maximum number of results to return.
		* `model_id` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | KEY_VALUE_EXTRACTION) The custom model ID.
		* `tenancy_id` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | KEY_VALUE_EXTRACTION) The custom model tenancy ID when modelId represents aliasName.
	* `is_zip_output_enabled` - (Optional) Whether or not to generate a ZIP file containing the results.
	* `language` - (Optional) The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	* `processor_type` - (Required) The type of the processor.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment identifier.
* `display_name` - The display name of the processor job.
* `id` - The id of the processor job.
* `input_location` - The location of the inputs.
	* `data` - Raw document data with Base64 encoding.
	* `object_locations` - The list of ObjectLocations.
		* `bucket` - The Object Storage bucket name.
		* `namespace` - The Object Storage namespace name.
		* `object` - The Object Storage object name.
	* `source_type` - The type of input location. The allowed values are:
		* `OBJECT_STORAGE_LOCATIONS`: A list of object locations in Object Storage.
		* `INLINE_DOCUMENT_CONTENT`: The content of an inline document. 
* `lifecycle_details` - The detailed status of FAILED state.
* `output_location` - The object storage location where to store analysis results.
	* `bucket` - The Object Storage bucket name.
	* `namespace` - The Object Storage namespace.
	* `prefix` - The Object Storage folder name.
* `percent_complete` - How much progress the operation has made, compared to the total amount of work to be performed.
* `processor_config` - The configuration of a processor.
	* `document_type` - The document type.
	* `features` - The types of document analysis requested.
		* `feature_type` - The type of document analysis requested. The allowed values are:
			* `LANGUAGE_CLASSIFICATION`: Detect the language.
			* `TEXT_EXTRACTION`: Recognize text.
			* `TABLE_EXTRACTION`: Detect and extract data in tables.
			* `KEY_VALUE_EXTRACTION`: Extract form fields.
			* `DOCUMENT_CLASSIFICATION`: Identify the type of document. 
		* `generate_searchable_pdf` - Whether or not to generate a searchable PDF file.
		* `max_results` - The maximum number of results to return.
		* `model_id` - The custom model ID.
		* `tenancy_id` - The custom model tenancy ID when modelId represents aliasName.
	* `is_zip_output_enabled` - Whether or not to generate a ZIP file containing the results.
	* `language` - The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	* `processor_type` - The type of the processor.
* `state` - The current state of the processor job.
* `time_accepted` - The job acceptance time.
* `time_finished` - The job finish time.
* `time_started` - The job start time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Processor Job
	* `update` - (Defaults to 20 minutes), when updating the Processor Job
	* `delete` - (Defaults to 20 minutes), when destroying the Processor Job


## Import

ProcessorJobs can be imported using the `id`, e.g.

```
$ terraform import oci_ai_document_processor_job.test_processor_job "id"
```

