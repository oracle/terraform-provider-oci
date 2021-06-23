---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream"
sidebar_current: "docs-oci-resource-streaming-stream"
description: |-
  Provides the Stream resource in Oracle Cloud Infrastructure Streaming service
---

# oci_streaming_stream
This resource provides the Stream resource in Oracle Cloud Infrastructure Streaming service.

Starts the provisioning of a new stream.
The stream will be created in the given compartment id or stream pool id, depending on which parameter is specified.
Compartment id and stream pool id cannot be specified at the same time.
To track the progress of the provisioning, you can periodically call [GetStream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream/GetStream).
In the response, the `lifecycleState` parameter of the [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream/) object tells you its current state.


## Example Usage

```hcl
resource "oci_streaming_stream" "test_stream" {
	#Required
	name = var.stream_name
	partitions = var.stream_partitions

	#Optional
	compartment_id = var.compartment_id
	defined_tags = var.stream_defined_tags
	freeform_tags = {"Department"= "Finance"}
	retention_in_hours = var.stream_retention_in_hours
	stream_pool_id = oci_streaming_stream_pool.test_stream_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the stream.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name of the stream. Avoid entering confidential information.  Example: `TelemetryEvents` 
* `partitions` - (Required) The number of partitions in the stream.
* `retention_in_hours` - (Optional) The retention period of the stream, in hours. Accepted values are between 24 and 168 (7 days). If not specified, the stream will have a retention period of 24 hours. 
* `stream_pool_id` - (Optional) (Updatable) The OCID of the stream pool that contains the stream.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the stream.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the stream.
* `lifecycle_state_details` - Any additional details about the current state of the stream.
* `messages_endpoint` - The endpoint to use when creating the StreamClient to consume or publish messages in the stream. If the associated stream pool is private, the endpoint is also private and can only be accessed from inside the stream pool's associated subnet. 
* `name` - The name of the stream. Avoid entering confidential information.  Example: `TelemetryEvents` 
* `partitions` - The number of partitions in the stream.
* `retention_in_hours` - The retention period of the stream, in hours. This property is read-only.
* `state` - The current state of the stream.
* `stream_pool_id` - The OCID of the stream pool that contains the stream.
* `time_created` - The date and time the stream was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream
	* `update` - (Defaults to 20 minutes), when updating the Stream
	* `delete` - (Defaults to 20 minutes), when destroying the Stream


## Import

Streams can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_stream.test_stream "id"
```

