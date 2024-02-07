---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_distribution_channels"
sidebar_current: "docs-oci-datasource-media_services-stream_distribution_channels"
description: |-
  Provides the list of Stream Distribution Channels in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_stream_distribution_channels
This data source provides the list of Stream Distribution Channels in Oracle Cloud Infrastructure Media Services service.

Lists the Stream Distribution Channels.

## Example Usage

```hcl
data "oci_media_services_stream_distribution_channels" "test_stream_distribution_channels" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.stream_distribution_channel_display_name
	id = var.stream_distribution_channel_id
	state = var.stream_distribution_channel_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name given.
* `id` - (Optional) Unique Stream Distribution Channel identifier.
* `state` - (Optional) A filter to return only the resources with lifecycleState matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `stream_distribution_channel_collection` - The list of stream_distribution_channel_collection.

### StreamDistributionChannel Reference

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

