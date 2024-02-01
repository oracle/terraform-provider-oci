---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_assets"
sidebar_current: "docs-oci-datasource-media_services-media_assets"
description: |-
  Provides the list of Media Assets in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_assets
This data source provides the list of Media Assets in Oracle Cloud Infrastructure Media Services service.

Returns a list of MediaAssetSummary.


## Example Usage

```hcl
data "oci_media_services_media_assets" "test_media_assets" {

	#Optional
	bucket = var.media_asset_bucket
	compartment_id = var.compartment_id
	display_name = var.media_asset_display_name
	distribution_channel_id = oci_mysql_channel.test_channel.id
	master_media_asset_id = oci_media_services_media_asset.test_media_asset.id
	media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id
	object = var.media_asset_object
	parent_media_asset_id = oci_media_services_media_asset.test_media_asset.id
	source_media_workflow_id = oci_media_services_media_workflow.test_media_workflow.id
	source_media_workflow_version = var.media_asset_source_media_workflow_version
	state = var.media_asset_state
	type = var.media_asset_type
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Optional) Filter MediaAsset by the bucket where the object is stored.
* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name given.
* `distribution_channel_id` - (Optional) Unique DistributionChannel identifier.
* `master_media_asset_id` - (Optional) Unique MediaAsset identifier of the first asset upload.
* `media_workflow_job_id` - (Optional) The ID of the MediaWorkflowJob used to produce this asset, if this parameter is supplied then the workflow ID must also be supplied.
* `object` - (Optional) Filter MediaAsset by the name of the object in object storage.
* `parent_media_asset_id` - (Optional) Unique MediaAsset identifier of the asset from which this asset is derived.
* `source_media_workflow_id` - (Optional) The ID of the MediaWorkflow used to produce this asset.
* `source_media_workflow_version` - (Optional) The version of the MediaWorkflow used to produce this asset.
* `state` - (Optional) A filter to return only the resources with lifecycleState matching the given lifecycleState.
* `type` - (Optional) Filter MediaAsset by the asset type.


## Attributes Reference

The following attributes are exported:

* `media_asset_collection` - The list of media_asset_collection.

### MediaAsset Reference

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

