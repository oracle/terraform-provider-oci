---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_appliance_image"
sidebar_current: "docs-oci-datasource-cloud_bridge-appliance_image"
description: |-
  Provides details about a specific Appliance Image in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_appliance_image
This data source provides details about a specific Appliance Image resource in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of Appliance Images.


## Example Usage

```hcl
data "oci_cloud_bridge_appliance_image" "test_appliance_image" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.appliance_image_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `items` - List of appliance images.
	* `checksum` - The checksum of the image file.
	* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The name of the image to be displayed.
	* `download_url` - The URL from which the appliance image can be downloaded.
	* `file_name` - The name of the appliance Image file.
	* `format` - The file format of the image file.
	* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - Unique identifier that is immutable on creation.
	* `platform` - The virtualization platform that the image file supports.
	* `size_in_mbs` - The size of the image file in megabytes.
	* `state` - The current state of the appliance image.
	* `time_created` - The time when the appliance image was created.An RFC3339 formatted datetime string.
	* `time_updated` - The time when the appliance image was last updated. An RFC3339 formatted datetime string.
	* `version` - The version of the image file.

