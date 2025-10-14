---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_job"
sidebar_current: "docs-oci-resource-ai_vision-stream_job"
description: |-
  Provides the Stream Job resource in Oracle Cloud Infrastructure Ai Vision service
---

# oci_ai_vision_stream_job
This resource provides the Stream Job resource in Oracle Cloud Infrastructure Ai Vision service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/vision/latest/StreamJob

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/aiVision

Create a stream analysis job with given inputs and features.


## Example Usage

```hcl
resource "oci_ai_vision_stream_job" "test_stream_job" {
	#Required
	compartment_id = var.compartment_id
	features {
		#Required
		feature_type = var.stream_job_features_feature_type

		#Optional
		max_results = var.stream_job_features_max_results
		should_return_landmarks = var.stream_job_features_should_return_landmarks
		tracking_types {

			#Optional
			biometric_store_compartment_id = oci_identity_compartment.test_compartment.id
			biometric_store_id = oci_ai_vision_biometric_store.test_biometric_store.id
			detection_model_id = oci_ai_vision_model.test_model.id
			max_results = var.stream_job_features_tracking_types_max_results
			objects = var.stream_job_features_tracking_types_objects
			should_return_landmarks = var.stream_job_features_tracking_types_should_return_landmarks
			tracking_model_id = oci_ai_vision_model.test_model.id
		}
	}
	stream_output_location {
		#Required
		bucket = var.stream_job_stream_output_location_bucket
		namespace = var.stream_job_stream_output_location_namespace
		output_location_type = var.stream_job_stream_output_location_output_location_type
		prefix = var.stream_job_stream_output_location_prefix

		#Optional
		obo_token = var.stream_job_stream_output_location_obo_token
	}
	stream_source_id = oci_ai_vision_stream_source.test_stream_source.id

	#Optional
	defined_tags = var.stream_job_defined_tags
	display_name = var.stream_job_display_name
	freeform_tags = var.stream_job_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - (Optional) (Updatable) Stream job display name.
* `features` - (Required) (Updatable) a list of stream analysis features.
	* `feature_type` - (Required) (Updatable) The feature of video analysis. Allowed values are:
		* OBJECT_TRACKING: Object tracking feature(OT).
		* FACE_DETECTION: Face detection feature(FD). 
	* `max_results` - (Applicable when feature_type=FACE_DETECTION) (Updatable) The maximum number of results to return.
	* `should_return_landmarks` - (Applicable when feature_type=FACE_DETECTION) (Updatable) Whether or not return face landmarks.
	* `tracking_types` - (Required when feature_type=OBJECT_TRACKING) (Updatable) List of details of what to track.
		* `biometric_store_compartment_id` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) compartment Id of biometric compartment.
		* `biometric_store_id` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) Which biometric store user wants to do face recognition
		* `detection_model_id` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) The detection model OCID.
		* `max_results` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) The maximum number of results to return.
		* `objects` - (Required when feature_type=OBJECT_TRACKING) (Updatable) List of the objects to be tracked.
		* `should_return_landmarks` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) Whether or not return face landmarks.
		* `tracking_model_id` - (Applicable when feature_type=OBJECT_TRACKING) (Updatable) The tracking model OCID.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `stream_output_location` - (Required) (Updatable) Details about a where results will be Sent
	* `bucket` - (Required) (Updatable) The Object Storage bucket name.
	* `namespace` - (Required) (Updatable) The Object Storage namespace.
	* `obo_token` - (Optional) (Updatable) Object storage output location
	* `output_location_type` - (Required) (Updatable) Type of device Allowed values are:
		* OBJECT_STORAGE
		* LIVEKIT_WEBRTC_AGENT 
	* `prefix` - (Required) (Updatable) The Object Storage folder name.
* `stream_source_id` - (Required) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of streamSource. 
* `state` - (Optional) (Updatable) The target state for the Stream Job. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Job
	* `update` - (Defaults to 20 minutes), when updating the Stream Job
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Job


## Import

StreamJobs can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_stream_job.test_stream_job "id"
```

