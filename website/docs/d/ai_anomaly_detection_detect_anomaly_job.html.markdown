---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_detect_anomaly_job"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-detect_anomaly_job"
description: |-
  Provides details about a specific Detect Anomaly Job in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_detect_anomaly_job
This data source provides details about a specific Detect Anomaly Job resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Gets a detect anomaly asynchronous job by identifier.

## Example Usage

```hcl
data "oci_ai_anomaly_detection_detect_anomaly_job" "test_detect_anomaly_job" {
	#Required
	detect_anomaly_job_id = oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job.id
}
```

## Argument Reference

The following arguments are supported:

* `detect_anomaly_job_id` - (Required) Unique asynchronous job identifier.


## Attributes Reference

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

