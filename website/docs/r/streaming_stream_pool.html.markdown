---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_stream_pool"
sidebar_current: "docs-oci-resource-streaming-stream_pool"
description: |-
  Provides the Stream Pool resource in Oracle Cloud Infrastructure Streaming service
---

# oci_streaming_stream_pool
This resource provides the Stream Pool resource in Oracle Cloud Infrastructure Streaming service.

Starts the provisioning of a new stream pool.
To track the progress of the provisioning, you can periodically call GetStreamPool.
In the response, the `lifecycleState` parameter of the object tells you its current state.


## Example Usage

```hcl
resource "oci_streaming_stream_pool" "test_stream_pool" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${var.stream_pool_name}"

	#Optional
	defined_tags = "${var.stream_pool_defined_tags}"
	freeform_tags = {"Department"= "Finance"}
	kafka_settings {

		#Optional
		auto_create_topics_enable = "${var.stream_pool_kafka_settings_auto_create_topics_enable}"
		bootstrap_servers = "${var.stream_pool_kafka_settings_bootstrap_servers}"
		log_retention_hours = "${var.stream_pool_kafka_settings_log_retention_hours}"
		num_partitions = "${var.stream_pool_kafka_settings_num_partitions}"
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the stream.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `kafka_settings` - (Optional) (Updatable) 
	* `auto_create_topics_enable` - (Optional) (Updatable) Enable auto creation of topic on the server.
	* `bootstrap_servers` - (Optional) (Updatable) Bootstrap servers.
	* `log_retention_hours` - (Optional) (Updatable) The number of hours to keep a log file before deleting it (in hours).
	* `num_partitions` - (Optional) (Updatable) The default number of log partitions per topic.
* `name` - (Required) (Updatable) The name of the stream pool. Avoid entering confidential information.  Example: `MyStreamPool` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

StreamPools can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_stream_pool.test_stream_pool "id"
```

