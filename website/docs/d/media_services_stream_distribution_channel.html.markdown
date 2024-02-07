---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_distribution_channel"
sidebar_current: "docs-oci-datasource-media_services-stream_distribution_channel"
description: |-
  Provides details about a specific Stream Distribution Channel in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_stream_distribution_channel
This data source provides details about a specific Stream Distribution Channel resource in Oracle Cloud Infrastructure Media Services service.

Gets a Stream Distribution Channel by identifier.

## Example Usage

```hcl
data "oci_media_services_stream_distribution_channel" "test_stream_distribution_channel" {
	#Required
	stream_distribution_channel_id = oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id
}
```

## Argument Reference

The following arguments are supported:

* `stream_distribution_channel_id` - (Required) Unique Stream Distribution Channel path identifier.


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

