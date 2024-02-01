---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_asset"
sidebar_current: "docs-oci-resource-media_services-media_asset"
description: |-
  Provides the Media Asset resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_media_asset
This resource provides the Media Asset resource in Oracle Cloud Infrastructure Media Services service.

Creates a new MediaAsset.


## Example Usage

```hcl
resource "oci_media_services_media_asset" "test_media_asset" {
	#Required
	compartment_id = var.compartment_id
	type = var.media_asset_type

	#Optional
	bucket = var.media_asset_bucket
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.media_asset_display_name
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.media_asset_locks_type

		#Optional
		message = var.media_asset_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.media_asset_locks_time_created
	}
	master_media_asset_id = oci_media_services_media_asset.test_media_asset.id
	media_asset_tags {
		#Required
		value = var.media_asset_media_asset_tags_value

		#Optional
		type = var.media_asset_media_asset_tags_type
	}
	media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id
	metadata {
		#Required
		metadata = var.media_asset_metadata_metadata
	}
	namespace = var.media_asset_namespace
	object = var.media_asset_object
	object_etag = var.media_asset_object_etag
	parent_media_asset_id = oci_media_services_media_asset.test_media_asset.id
	segment_range_end_index = var.media_asset_segment_range_end_index
	segment_range_start_index = var.media_asset_segment_range_start_index
	source_media_workflow_id = oci_media_services_media_workflow.test_media_workflow.id
	source_media_workflow_version = var.media_asset_source_media_workflow_version
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Optional) The name of the object storage bucket where this asset is located.
* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Display name for the Media Asset. Does not have to be unique. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) (Updatable) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `master_media_asset_id` - (Optional) (Updatable) The ID of the senior most asset from which this asset is derived.
* `media_asset_tags` - (Optional) (Updatable) list of tags for the MediaAsset.
	* `type` - (Optional) (Updatable) Type of the tag.
	* `value` - (Required) (Updatable) Tag of the MediaAsset.
* `media_workflow_job_id` - (Optional) The ID of the MediaWorkflowJob used to produce this asset.
* `metadata` - (Optional) (Updatable) List of Metadata.
	* `metadata` - (Required) (Updatable) JSON string containing the technial metadata for the media asset.
* `namespace` - (Optional) The object storage namespace where this asset is located.
* `object` - (Optional) The object storage object name that identifies this asset.
* `object_etag` - (Optional) eTag of the underlying object storage object.
* `parent_media_asset_id` - (Optional) (Updatable) The ID of the parent asset from which this asset is derived.
* `segment_range_end_index` - (Optional) The end index for video segment files.
* `segment_range_start_index` - (Optional) The start index for video segment files.
* `source_media_workflow_id` - (Optional) The ID of the MediaWorkflow used to produce this asset.
* `source_media_workflow_version` - (Optional) The version of the MediaWorkflow used to produce this asset.
* `type` - (Required) (Updatable) The type of the media asset.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - The name of the object storage bucket where this represented asset is located.
* `compartment_id` - The ID of the compartment containing the MediaAsset.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `master_media_asset_id` - The ID of the senior most asset from which this asset is derived.
* `media_asset_tags` - List of tags for the MediaAsset.
	* `type` - Type of the tag.
	* `value` - Tag of the MediaAsset.
* `media_workflow_job_id` - The ID of the MediaWorkflowJob used to produce this asset.
* `metadata` - List of Metadata.
	* `metadata` - JSON string containing the technial metadata for the media asset.
* `namespace` - The object storage namespace where this asset is located.
* `object` - The object storage object name that identifies this asset.
* `object_etag` - eTag of the underlying object storage object.
* `parent_media_asset_id` - The ID of the parent asset from which this asset is derived.
* `segment_range_end_index` - The end index of video segment files.
* `segment_range_start_index` - The start index for video segment files.
* `source_media_workflow_id` - The ID of the MediaWorkflow used to produce this asset.
* `source_media_workflow_version` - The version of the MediaWorkflow used to produce this asset.
* `state` - The current state of the MediaAsset.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the MediaAsset was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the MediaAsset was updated. An RFC3339 formatted datetime string.
* `type` - The type of the media asset.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Media Asset
	* `update` - (Defaults to 20 minutes), when updating the Media Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Media Asset


## Import

MediaAssets can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_media_asset.test_media_asset "id"
```

