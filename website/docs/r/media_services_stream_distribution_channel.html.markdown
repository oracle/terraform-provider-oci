---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_distribution_channel"
sidebar_current: "docs-oci-resource-media_services-stream_distribution_channel"
description: |-
  Provides the Stream Distribution Channel resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_stream_distribution_channel
This resource provides the Stream Distribution Channel resource in Oracle Cloud Infrastructure Media Services service.

Creates a new Stream Distribution Channel.


## Example Usage

```hcl
resource "oci_media_services_stream_distribution_channel" "test_stream_distribution_channel" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.stream_distribution_channel_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.stream_distribution_channel_locks_type

		#Optional
		message = var.stream_distribution_channel_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.stream_distribution_channel_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Stream Distribution Channel display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) (Updatable) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Stream Distribution Channel display name. Avoid entering confidential information.
* `domain_name` - Unique domain name of the Distribution Channel.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `state` - The current state of the Stream Distribution Channel.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the Stream Distribution Channel was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the Stream Distribution Channel was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Distribution Channel
	* `update` - (Defaults to 20 minutes), when updating the Stream Distribution Channel
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Distribution Channel


## Import

StreamDistributionChannels can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_stream_distribution_channel.test_stream_distribution_channel "id"
```

