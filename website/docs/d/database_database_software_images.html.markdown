---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_software_images"
sidebar_current: "docs-oci-datasource-database-database_software_images"
description: |-
  Provides the list of Database Software Images in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_software_images
This data source provides the list of Database Software Images in Oracle Cloud Infrastructure Database service.

Gets a list of the database software images in the specified compartment.


## Example Usage

```hcl
data "oci_database_database_software_images" "test_database_software_images" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.database_software_image_display_name}"
	image_shape_family = "${var.database_software_image_image_shape_family}"
	image_type = "${var.database_software_image_image_type}"
	state = "${var.database_software_image_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `image_shape_family` - (Optional) A filter to return only resources that match the given image shape family exactly.
* `image_type` - (Optional) A filter to return only resources that match the given image type exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `database_software_images` - The list of database_software_images.

### DatabaseSoftwareImage Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_software_image_included_patches` - List of one-off patches for Database Homes.
* `database_software_image_one_off_patches` - List of one-off patches for Database Homes.
* `database_version` - The database version with which the database software image is to be built.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the database software image. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database software image.
* `image_shape_family` - To what shape the image is meant for.
* `image_type` - The type of software image. Can be grid or database.
* `included_patches_summary` - The patches included in the image and the version of the image
* `lifecycle_details` - Detailed message for the lifecycle state.
* `ls_inventory` - output from lsinventory which will get passed as a string
* `patch_set` - The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions) operation.
* `state` - The current state of the database software image.
* `time_created` - The date and time the database software image was created.

