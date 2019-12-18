---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream_pools"
sidebar_current: "docs-oci-datasource-streaming-stream_pools"
description: |-
  Provides the list of Stream Pools in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_stream_pools
This data source provides the list of Stream Pools in Oracle Cloud Infrastructure Streaming service.

List the stream pools for a given compartment ID.

## Example Usage

```hcl
data "oci_streaming_stream_pools" "test_stream_pools" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	id = "${var.stream_pool_id}"
	name = "${var.stream_pool_name}"
	state = "${var.stream_pool_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `id` - (Optional) A filter to return only resources that match the given ID exactly. 
* `name` - (Optional) A filter to return only resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `stream_pools` - The list of stream_pools.

### StreamPool Reference

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

