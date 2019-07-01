---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream"
sidebar_current: "docs-oci-resource-streaming-stream"
description: |-
  Provides the Stream resource in Oracle Cloud Infrastructure Streaming service
---

# oci_streaming_stream
This resource provides the Stream resource in Oracle Cloud Infrastructure Streaming service.

Starts the provisioning of a new stream.
To track the progress of the provisioning, you can periodically call [GetStream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream/GetStream).
In the response, the `lifecycleState` parameter of the [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream/) object tells you its current state.


## Example Usage

```hcl
resource "oci_streaming_stream" "test_stream" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.stream_name}"
	partitions = "${var.stream_partitions}"

	#Optional
	defined_tags = "${var.stream_defined_tags}"
	freeform_tags = {"Department"= "Finance"}
	retention_in_hours = "${var.stream_retention_in_hours}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the stream.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the stream. Avoid entering confidential information.  Example: `TelemetryEvents` 
* `partitions` - (Required) The number of partitions in the stream.
* `retention_in_hours` - (Optional) The retention period of the stream, in hours. Accepted values are between 24 and 168 (7 days). If not specified, the stream will have a retention period of 24 hours. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the stream.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the stream.
* `lifecycle_state_details` - Any additional details about the current state of the stream.
* `messages_endpoint` - The endpoint to use when creating the StreamClient to consume or publish messages in the stream.
* `name` - The name of the stream. Avoid entering confidential information.  Example: `TelemetryEvents` 
* `partitions` - The number of partitions in the stream.
* `retention_in_hours` - The retention period of the stream, in hours. This property is read-only.
* `state` - The current state of the stream.
* `time_created` - The date and time the stream was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

## Import

Streams can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_stream.test_stream "id"
```

