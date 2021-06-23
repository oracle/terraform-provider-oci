---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_software_image"
sidebar_current: "docs-oci-resource-database-database_software_image"
description: |-
  Provides the Database Software Image resource in Oracle Cloud Infrastructure Database service
---

# oci_database_database_software_image
This resource provides the Database Software Image resource in Oracle Cloud Infrastructure Database service.

create database software image in the specified compartment.


## Example Usage

```hcl
resource "oci_database_database_software_image" "test_database_software_image" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.database_software_image_display_name

	#Optional
	database_software_image_one_off_patches = var.database_software_image_database_software_image_one_off_patches
	database_version = var.database_software_image_database_version
	defined_tags = var.database_software_image_defined_tags
	freeform_tags = {"Department"= "Finance"}
	image_shape_family = var.database_software_image_image_shape_family
	image_type = var.database_software_image_image_type
	ls_inventory = var.database_software_image_ls_inventory
	patch_set = var.database_software_image_patch_set
	source_db_home_id = oci_database_db_home.test_db_home.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the database software image  belongs in.
* `database_software_image_one_off_patches` - (Optional) List of one-off patches for Database Homes.
* `database_version` - (Optional) The database version with which the database software image is to be built.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the database software image. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `image_shape_family` - (Optional) To what shape the image is meant for.
* `image_type` - (Optional) The type of software image. Can be grid or database.
* `ls_inventory` - (Optional) output from lsinventory which will get passed as a string
* `patch_set` - (Optional) The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `source_db_home_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `is_upgrade_supported` - True if this Database software image is supported for Upgrade.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `ls_inventory` - output from lsinventory which will get passed as a string
* `patch_set` - The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `state` - The current state of the database software image.
* `time_created` - The date and time the database software image was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 2 hours), when creating the Database Software Image
	* `update` - (Defaults to 30 minutes), when updating the Database Software Image
	* `delete` - (Defaults to 30 minutes), when destroying the Database Software Image


## Import

DatabaseSoftwareImages can be imported using the `id`, e.g.

```
$ terraform import oci_database_database_software_image.test_database_software_image "id"
```

