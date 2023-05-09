---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_processor_job"
sidebar_current: "docs-oci-datasource-ai_document-processor_job"
description: |-
  Provides details about a specific Processor Job in Oracle Cloud Infrastructure Ai Document service
---

# Data Source: oci_ai_document_processor_job
This data source provides details about a specific Processor Job resource in Oracle Cloud Infrastructure Ai Document service.

Get the details of a processor job.


## Example Usage

```hcl
data "oci_ai_document_processor_job" "test_processor_job" {
	#Required
	processor_job_id = oci_ai_document_processor_job.test_processor_job.id
}
```

## Argument Reference

The following arguments are supported:

* `processor_job_id` - (Required) Processor job id.


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
* `output_location` - The Object Storage Location.
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
	* `is_zip_output_enabled` - Whether or not to generate a ZIP file containing the results.
	* `language` - The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	* `processor_type` - The type of the processor.
* `state` - The current state of the processor job.
* `time_accepted` - The job acceptance time.
* `time_finished` - The job finish time.
* `time_started` - The job start time.

