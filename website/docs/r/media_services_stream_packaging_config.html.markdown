---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_packaging_config"
sidebar_current: "docs-oci-resource-media_services-stream_packaging_config"
description: |-
  Provides the Stream Packaging Config resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_stream_packaging_config
This resource provides the Stream Packaging Config resource in Oracle Cloud Infrastructure Media Services service.

Creates a new Packaging Configuration.


## Example Usage

```hcl
resource "oci_media_services_stream_packaging_config" "test_stream_packaging_config" {
	#Required
	display_name = var.stream_packaging_config_display_name
	distribution_channel_id = oci_mysql_channel.test_channel.id
	segment_time_in_seconds = var.stream_packaging_config_segment_time_in_seconds
	stream_packaging_format = var.stream_packaging_config_stream_packaging_format

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	encryption {
		#Required
		algorithm = var.stream_packaging_config_encryption_algorithm

		#Optional
		kms_key_id = oci_kms_key.test_key.id
	}
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.stream_packaging_config_locks_type

		#Optional
		message = var.stream_packaging_config_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.stream_packaging_config_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) The name of the stream Packaging Configuration. Avoid entering confidential information.
* `distribution_channel_id` - (Required) Unique identifier of the Distribution Channel that this stream packaging configuration belongs to.
* `encryption` - (Optional) The encryption used by the stream packaging configuration.
	* `algorithm` - (Required) The encryption algorithm for the stream packaging configuration.
	* `kms_key_id` - (Applicable when algorithm=AES128) The identifier of the customer managed Vault KMS symmetric encryption key (null if Oracle managed).
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `segment_time_in_seconds` - (Required) The duration in seconds for each fragment.
* `stream_packaging_format` - (Required) The output format for the package.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the stream packaging configuration. Avoid entering confidential information.
* `distribution_channel_id` - Unique identifier of the Distribution Channel that this stream packaging configuration belongs to.
* `encryption` - The encryption used by the stream packaging configuration.
	* `algorithm` - The encryption algorithm for the stream packaging configuration.
	* `kms_key_id` - The identifier of the customer managed Vault KMS symmetric encryption key (null if Oracle managed).
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `segment_time_in_seconds` - The duration in seconds for each fragment.
* `state` - The current state of the Packaging Configuration.
* `stream_packaging_format` - The output format for the package.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the Packaging Configuration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the Packaging Configuration was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Packaging Config
	* `update` - (Defaults to 20 minutes), when updating the Stream Packaging Config
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Packaging Config


## Import

StreamPackagingConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_stream_packaging_config.test_stream_packaging_config "id"
```

