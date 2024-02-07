---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_asset_distribution_channel_attachment"
sidebar_current: "docs-oci-datasource-media_services-media_asset_distribution_channel_attachment"
description: |-
  Provides details about a specific Media Asset Distribution Channel Attachment in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_asset_distribution_channel_attachment
This data source provides details about a specific Media Asset Distribution Channel Attachment resource in Oracle Cloud Infrastructure Media Services service.

Gets a MediaAssetDistributionChannelAttachment for a MediaAsset by identifiers.

## Example Usage

```hcl
data "oci_media_services_media_asset_distribution_channel_attachment" "test_media_asset_distribution_channel_attachment" {
	#Required
	distribution_channel_id = oci_mysql_channel.test_channel.id
	media_asset_id = oci_media_services_media_asset.test_media_asset.id

	#Optional
	version = var.media_asset_distribution_channel_attachment_version
}
```

## Argument Reference

The following arguments are supported:

* `distribution_channel_id` - (Required) Unique DistributionChannel identifier.
* `media_asset_id` - (Required) Unique MediaAsset identifier
* `version` - (Optional) Version of the attachment.


## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `distribution_channel_id` - OCID of associated Distribution Channel.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `media_workflow_job_id` - The ingest MediaWorkflowJob ID that created this attachment.
* `metadata_ref` - The identifier for the metadata.
* `state` - Lifecycle state of the attachment.
* `version` - Version of the attachment.

