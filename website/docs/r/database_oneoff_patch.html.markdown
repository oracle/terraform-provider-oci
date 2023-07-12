---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_oneoff_patch"
sidebar_current: "docs-oci-resource-database-oneoff_patch"
description: |-
  Provides the Oneoff Patch resource in Oracle Cloud Infrastructure Database service
---

# oci_database_oneoff_patch
This resource provides the Oneoff Patch resource in Oracle Cloud Infrastructure Database service.

Creates one-off patch for specified database version to download.


## Example Usage

```hcl
resource "oci_database_oneoff_patch" "test_oneoff_patch" {
	#Required
	compartment_id = var.compartment_id
	db_version = var.oneoff_patch_db_version
	display_name = var.oneoff_patch_display_name
	release_update = var.oneoff_patch_release_update

	#Optional
	defined_tags = var.oneoff_patch_defined_tags
	freeform_tags = {"Department"= "Finance"}
	one_off_patches = var.oneoff_patch_one_off_patches
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_version` - (Required) A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) One-off patch name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `one_off_patches` - (Optional) List of one-off patches for Database Homes.
* `release_update` - (Required) The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `download_oneoff_patch_trigger` - (Optional) (Updatable) An optional property when incremented triggers Download Oneoff Patch. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_version` - A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - One-off patch name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the one-off patch.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `one_off_patches` - List of one-off patches for Database Homes.
* `release_update` - The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `sha256sum` - SHA-256 checksum of the one-off patch.
* `size_in_kbs` - The size of one-off patch in kilobytes.
* `state` - The current state of the one-off patch.
* `time_created` - The date and time one-off patch was created.
* `time_of_expiration` - The date and time until which the one-off patch will be available for download.
* `time_updated` - The date and time one-off patch was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oneoff Patch
	* `update` - (Defaults to 20 minutes), when updating the Oneoff Patch
	* `delete` - (Defaults to 20 minutes), when destroying the Oneoff Patch


## Import

OneoffPatches can be imported using the `id`, e.g.

```
$ terraform import oci_database_oneoff_patch.test_oneoff_patch "id"
```

