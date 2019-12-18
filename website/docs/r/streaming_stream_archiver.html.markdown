---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream_archiver"
sidebar_current: "docs-oci-resource-streaming-stream_archiver"
description: |-
  Provides the Stream Archiver resource in Oracle Cloud Infrastructure Streaming service
---

# oci_streaming_stream_archiver
This resource provides the Stream Archiver resource in Oracle Cloud Infrastructure Streaming service.

Starts the provisioning of a new stream archiver.
To track the progress of the provisioning, you can periodically call [GetArchiver](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream/GetArchiver).
In the response, the `lifecycleState` parameter of the [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Archiver) object tells you its current state.


## Example Usage

```hcl
resource "oci_streaming_stream_archiver" "test_stream_archiver" {
	#Required
	batch_rollover_size_in_mbs = "${var.stream_archiver_batch_rollover_size_in_mbs}"
	batch_rollover_time_in_seconds = "${var.stream_archiver_batch_rollover_time_in_seconds}"
	bucket = "${var.stream_archiver_bucket}"
	start_position = "${var.stream_archiver_start_position}"
	stream_id = "${oci_streaming_stream.test_stream.id}"
	use_existing_bucket = "${var.stream_archiver_use_existing_bucket}"

	#Optional
	state = "stopped"
}
```

## Argument Reference

The following arguments are supported:

* `batch_rollover_size_in_mbs` - (Required) (Updatable) The batch rollover size in megabytes.
* `batch_rollover_time_in_seconds` - (Required) (Updatable) The rollover time in seconds.
* `bucket` - (Required) (Updatable) The name of the bucket.
* `start_position` - (Required) (Updatable) The start message.
* `stream_id` - (Required) The OCID of the stream. 
* `use_existing_bucket` - (Required) (Updatable) The flag to create a new bucket or use existing one.
* `state` - (Optional) (Updatable) The target state for the instance pool. Could be set to RUNNING or STOPPED.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `batch_rollover_size_in_mbs` - The batch rollover size in megabytes.
* `batch_rollover_time_in_seconds` - The rollover time in seconds.
* `bucket` - The name of the bucket.
* `error` - 
	* `code` - A short error code that defines the error, meant for programmatic parsing.
	* `message` - A human-readable error string.
* `start_position` - The start message.
* `state` - The state of the stream archiver.
* `time_created` - Time when the resource was created.
* `use_existing_bucket` - The flag to create a new bucket or use existing one.

## Import

StreamArchiver can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_stream_archiver.test_stream_archiver "streams/{streamId}/archiver" 
```

