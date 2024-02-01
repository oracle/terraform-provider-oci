---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_stream_packaging_configs"
sidebar_current: "docs-oci-datasource-media_services-stream_packaging_configs"
description: |-
  Provides the list of Stream Packaging Configs in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_stream_packaging_configs
This data source provides the list of Stream Packaging Configs in Oracle Cloud Infrastructure Media Services service.

Lists the Stream Packaging Configurations.

## Example Usage

```hcl
data "oci_media_services_stream_packaging_configs" "test_stream_packaging_configs" {
	#Required
	distribution_channel_id = oci_mysql_channel.test_channel.id

	#Optional
	display_name = var.stream_packaging_config_display_name
	state = var.stream_packaging_config_state
	stream_packaging_config_id = oci_media_services_stream_packaging_config.test_stream_packaging_config.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only the resources that match the entire display name given.
* `distribution_channel_id` - (Required) Unique Stream Distribution Channel identifier.
* `state` - (Optional) A filter to return only the resources with lifecycleState matching the given lifecycleState.
* `stream_packaging_config_id` - (Optional) Unique Stream Packaging Configuration identifier.


## Attributes Reference

The following attributes are exported:

* `stream_packaging_config_collection` - The list of stream_packaging_config_collection.

### StreamPackagingConfig Reference

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

