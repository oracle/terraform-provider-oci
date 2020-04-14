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
* `custom_encryption_key` - 
	* `key_state` - Life cycle State of the custom key
	* `kms_key_id` - Custom Encryption Key (Master Key) ocid.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `endpoint_fqdn` - The FQDN used to access the streams inside the stream pool (same FQDN as the messagesEndpoint attribute of a [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream) object). If the stream pool is private, the FQDN is customized and can only be accessed from inside the associated subnetId, otherwise the FQDN is publicly resolvable. Depending on which protocol you attempt to use, you need to either prepend https or append the Kafka port. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the stream pool.
* `is_private` - True if the stream pool is private, false otherwise. The associated endpoint and subnetId of a private stream pool can be retrieved through the [GetStreamPool](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/GetStreamPool) API. 
* `kafka_settings` - 
	* `auto_create_topics_enable` - Enable auto creation of topic on the server.
	* `bootstrap_servers` - Bootstrap servers.
	* `log_retention_hours` - The number of hours to keep a log file before deleting it (in hours).
	* `num_partitions` - The default number of log partitions per topic.
* `lifecycle_state_details` - Any additional details about the current state of the stream.
* `name` - The name of the stream pool.
* `private_endpoint_settings` - 
	* `nsg_ids` - The optional list of network security groups that are associated with the private endpoint of the stream pool.
	* `private_endpoint_ip` - The private IP associated with the stream pool in the associated subnetId. The stream pool's FQDN resolves to that IP and should be used - instead of the private IP - in order to not trigger any TLS issues. 
	* `subnet_id` - The subnet id from which the private stream pool can be accessed. Trying to access the streams from another network location will result in an error. 
* `state` - The current state of the stream pool.
* `time_created` - The date and time the stream pool was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

