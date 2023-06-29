---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_oneoff_patch"
sidebar_current: "docs-oci-datasource-database-oneoff_patch"
description: |-
  Provides details about a specific Oneoff Patch in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_oneoff_patch
This data source provides details about a specific Oneoff Patch resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified one-off patch.


## Example Usage

```hcl
data "oci_database_oneoff_patch" "test_oneoff_patch" {
	#Required
	oneoff_patch_id = oci_database_oneoff_patch.test_oneoff_patch.id
}
```

## Argument Reference

The following arguments are supported:

* `oneoff_patch_id` - (Required) The one-off patch [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

