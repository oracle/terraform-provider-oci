---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_software_image"
sidebar_current: "docs-oci-datasource-database-database_software_image"
description: |-
  Provides details about a specific Database Software Image in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_software_image
This data source provides details about a specific Database Software Image resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified database software image.

## Example Usage

```hcl
data "oci_database_database_software_image" "test_database_software_image" {
	#Required
	database_software_image_id = oci_database_database_software_image.test_database_software_image.id
}
```

## Argument Reference

The following arguments are supported:

* `database_software_image_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

