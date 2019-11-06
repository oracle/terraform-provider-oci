---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream_archiver"
sidebar_current: "docs-oci-datasource-streaming-stream_archiver"
description: |-
  Provides details about a specific Stream Archiver in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_stream_archiver
This data source provides details about a specific Stream Archiver resource in Oracle Cloud Infrastructure Streaming service.

Returns the current state of the stream archiver.


## Example Usage

```hcl
data "oci_streaming_stream_archiver" "test_stream_archiver" {
	#Required
	stream_id = "${oci_streaming_stream.test_stream.id}"
}
```

## Argument Reference

The following arguments are supported:

* `stream_id` - (Required) The OCID of the stream. 


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

