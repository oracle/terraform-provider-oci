---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_job"
sidebar_current: "docs-oci-datasource-ai_vision-stream_job"
description: |-
  Provides details about a specific Stream Job in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_stream_job
This data source provides details about a specific Stream Job resource in Oracle Cloud Infrastructure Ai Vision service.

Get details of a stream analysis job.


## Example Usage

```hcl
data "oci_ai_vision_stream_job" "test_stream_job" {
	#Required
	stream_job_id = oci_ai_vision_stream_job.test_stream_job.id
}
```

## Argument Reference

The following arguments are supported:

* `stream_job_id` - (Required) Stream job id.


## Attributes Reference

The following attributes are exported:

* `agent_participant_id` - participant id of agent where results need to be sent
* `compartment_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - Stream job display name.
* `features` - a list of document analysis features.
	* `feature_type` - The feature of video analysis. Allowed values are:
		* OBJECT_TRACKING: Object tracking feature(OT).
		* FACE_DETECTION: Face detection feature(FD). 
	* `max_results` - The maximum number of results to return.
	* `should_return_landmarks` - Whether or not return face landmarks.
	* `tracking_types` - List of details of what to track.
		* `biometric_store_compartment_id` - compartment Id of biometric compartment.
		* `biometric_store_id` - Which biometric store user wants to do face recognition
		* `detection_model_id` - The detection model OCID.
		* `max_results` - The maximum number of results to return.
		* `objects` - List of the objects to be tracked.
		* `should_return_landmarks` - Whether or not return face landmarks.
		* `tracking_model_id` - The tracking model OCID.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamJob. 
* `lifecycle_details` - Additional details about current state of streamJob
* `state` - The current state of the Stream job.
* `stream_output_location` - Details about a where results will be Sent
	* `bucket` - The Object Storage bucket name.
	* `namespace` - The Object Storage namespace.
	* `obo_token` - Object storage output location
	* `output_location_type` - Type of device Allowed values are:
		* OBJECT_STORAGE
		* LIVEKIT_WEBRTC_AGENT 
	* `prefix` - The Object Storage folder name.
* `stream_source_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamSource 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `time_created` - When the streamJob was created, as an RFC3339 datetime string.
* `time_updated` - When the stream job was updated, as an RFC3339 datetime string.

