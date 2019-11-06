---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream"
sidebar_current: "docs-oci-datasource-streaming-stream"
description: |-
  Provides details about a specific Stream in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_stream
This data source provides details about a specific Stream resource in Oracle Cloud Infrastructure Streaming service.

Gets detailed information about a stream, including the number of partitions.

## Example Usage

```hcl
data "oci_streaming_stream" "test_stream" {
	#Required
	stream_id = "${oci_streaming_stream.test_stream.id}"
}
```

## Argument Reference

The following arguments are supported:

* `stream_id` - (Required) The OCID of the stream to retrieve.


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

