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
	compartment_id = var.compartment_id
	name = var.stream_pool_name

	#Optional
	custom_encryption_key {
		#Required
		kms_key_id = oci_kms_key.test_key.id
	}
	defined_tags = var.stream_pool_defined_tags
	freeform_tags = {"Department"= "Finance"}
	kafka_settings {

		#Optional
		auto_create_topics_enable = var.stream_pool_kafka_settings_auto_create_topics_enable
		bootstrap_servers = var.stream_pool_kafka_settings_bootstrap_servers
		log_retention_hours = var.stream_pool_kafka_settings_log_retention_hours
		num_partitions = var.stream_pool_kafka_settings_num_partitions
	}
	private_endpoint_settings {

		#Optional
		nsg_ids = var.stream_pool_private_endpoint_settings_nsg_ids
		private_endpoint_ip = var.stream_pool_private_endpoint_settings_private_endpoint_ip
		subnet_id = oci_core_subnet.test_subnet.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the stream.
* `custom_encryption_key` - (Optional) (Updatable) The OCID of the custom encryption key to be used or deleted if currently being used.
	* `kms_key_id` - (Required) (Updatable) Custom Encryption Key (Master Key) ocid.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `kafka_settings` - (Optional) (Updatable) Settings for the Kafka compatibility layer.
	* `auto_create_topics_enable` - (Optional) (Updatable) Enable auto creation of topic on the server.
	* `bootstrap_servers` - (Optional) (Updatable) Bootstrap servers.
	* `log_retention_hours` - (Optional) (Updatable) The number of hours to keep a log file before deleting it (in hours).
	* `num_partitions` - (Optional) (Updatable) The default number of log partitions per topic.
* `name` - (Required) (Updatable) The name of the stream pool. Avoid entering confidential information.  Example: `MyStreamPool` 
* `private_endpoint_settings` - (Optional) Optional parameters if a private stream pool is requested.
	* `nsg_ids` - (Optional) The optional list of network security groups to be used with the private endpoint of the stream pool. That value cannot be changed. 
	* `private_endpoint_ip` - (Optional) The optional private IP you want to be associated with your private stream pool. That parameter can only be specified when the subnetId parameter is set. It cannot be changed. The private IP needs to be part of the CIDR range of the specified subnetId or the creation will fail. If not specified a random IP inside the subnet will be chosen. After the stream pool is created, a custom FQDN, pointing to this private IP, is created. The FQDN is then used to access the service instead of the private IP. 
	* `subnet_id` - (Optional) If specified, the stream pool will be private and only accessible from inside that subnet. Producing-to and consuming-from a stream inside a private stream pool can also only be done from inside the subnet. That value cannot be changed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID that the pool belongs to.
* `custom_encryption_key` - Custom Encryption Key which will be used for encryption by all the streams in the pool.
	* `key_state` - Life cycle State of the custom key
	* `kms_key_id` - Custom Encryption Key (Master Key) ocid.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `endpoint_fqdn` - The FQDN used to access the streams inside the stream pool (same FQDN as the messagesEndpoint attribute of a [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream) object). If the stream pool is private, the FQDN is customized and can only be accessed from inside the associated subnetId, otherwise the FQDN is publicly resolvable. Depending on which protocol you attempt to use, you need to either prepend https or append the Kafka port. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the stream pool.
* `is_private` - True if the stream pool is private, false otherwise. The associated endpoint and subnetId of a private stream pool can be retrieved through the [GetStreamPool](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/GetStreamPool) API. 
* `kafka_settings` - Settings for the Kafka compatibility layer.
	* `auto_create_topics_enable` - Enable auto creation of topic on the server.
	* `bootstrap_servers` - Bootstrap servers.
	* `log_retention_hours` - The number of hours to keep a log file before deleting it (in hours).
	* `num_partitions` - The default number of log partitions per topic.
* `lifecycle_state_details` - Any additional details about the current state of the stream.
* `name` - The name of the stream pool.
* `private_endpoint_settings` - Optional settings if the stream pool is private.
	* `nsg_ids` - The optional list of network security groups that are associated with the private endpoint of the stream pool.
	* `private_endpoint_ip` - The private IP associated with the stream pool in the associated subnetId. The stream pool's FQDN resolves to that IP and should be used - instead of the private IP - in order to not trigger any TLS issues. 
	* `subnet_id` - The subnet id from which the private stream pool can be accessed. Trying to access the streams from another network location will result in an error. 
* `state` - The current state of the stream pool.
* `time_created` - The date and time the stream pool was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Pool
	* `update` - (Defaults to 20 minutes), when updating the Stream Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Pool


## Import

StreamPools can be imported using the `id`, e.g.

```
$ terraform import oci_streaming_stream_pool.test_stream_pool "id"
```

