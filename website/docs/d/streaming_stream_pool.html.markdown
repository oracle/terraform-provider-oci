---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream_pool"
sidebar_current: "docs-oci-datasource-streaming-stream_pool"
description: |-
  Provides details about a specific Stream Pool in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_stream_pool
This data source provides details about a specific Stream Pool resource in Oracle Cloud Infrastructure Streaming service.

Gets detailed information about the stream pool, such as Kafka settings.

## Example Usage

```hcl
data "oci_streaming_stream_pool" "test_stream_pool" {
	#Required
	stream_pool_id = "${oci_streaming_stream_pool.test_stream_pool.id}"
}
```

## Argument Reference

The following arguments are supported:

* `stream_pool_id` - (Required) The OCID of the stream pool. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID that the pool belongs to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the stream pool.
* `kafka_settings` - 
	* `auto_create_topics_enable` - Enable auto creation of topic on the server.
	* `bootstrap_servers` - Bootstrap servers.
	* `log_retention_hours` - The number of hours to keep a log file before deleting it (in hours).
	* `num_partitions` - The default number of log partitions per topic.
* `lifecycle_state_details` - Any additional details about the current state of the stream.
* `name` - The name of the stream pool.
* `state` - The current state of the stream pool.
* `time_created` - The date and time the stream pool was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

