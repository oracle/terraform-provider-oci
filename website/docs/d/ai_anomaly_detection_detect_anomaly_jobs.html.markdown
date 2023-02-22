---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_detect_anomaly_jobs"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-detect_anomaly_jobs"
description: |-
  Provides the list of Detect Anomaly Jobs in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_detect_anomaly_jobs
This data source provides the list of Detect Anomaly Jobs in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Returns a list of all the Anomaly Detection jobs in the specified compartment.


## Example Usage

```hcl
data "oci_ai_anomaly_detection_detect_anomaly_jobs" "test_detect_anomaly_jobs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	detect_anomaly_job_id = oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job.id
	display_name = var.detect_anomaly_job_display_name
	model_id = oci_ai_anomaly_detection_model.test_model.id
	project_id = oci_ai_anomaly_detection_project.test_project.id
	state = var.detect_anomaly_job_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `detect_anomaly_job_id` - (Optional) Unique Async Job identifier
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `model_id` - (Optional) The ID of the trained model for which to list the resources.
* `project_id` - (Optional) The ID of the project for which to list the objects.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `detect_anomaly_job_collection` - The list of detect_anomaly_job_collection.

### DetectAnomalyJob Reference

The following attributes are exported:

* `are_all_estimates_required` - Flag to enable the service to return estimates for all data points rather than just the anomalous data points
* `compartment_id` - The OCID of the compartment that starts the job.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Detect anomaly job description.
* `display_name` - Detect anomaly job display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Id of the job.
* `input_details` - Input details for detect anomaly job.
	* `input_type` - The type of input location Allowed values are:
		* `INLINE`: Inline input data.
		* `OBJECT_LIST`: Object store output location. 
	* `message` - Inline input details.
	* `object_locations` - List of ObjectLocations.
		* `bucket` - Object Storage bucket name.
		* `namespace` - Object Storage namespace name.
		* `object` - Object Storage object name.
* `lifecycle_state_details` - The current state details of the batch document job.
* `model_id` - The OCID of the trained model.
* `output_details` - Output details for detect anomaly job.
	* `bucket` - Object Storage bucket name.
	* `namespace` - Object Storage namespace.
	* `output_type` - The type of output location Allowed values are:
		* `OBJECT_STORAGE`: Object store output location. 
	* `prefix` - Object Storage folder name.
* `project_id` - The OCID of the project.
* `sensitivity` - The value that customer can adjust to control the sensitivity of anomaly detection
* `state` - The current state of the batch document job.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_accepted` - Job accepted time
* `time_finished` - Job finished time
* `time_started` - Job started time

