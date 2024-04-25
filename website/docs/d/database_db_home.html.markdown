---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_home"
sidebar_current: "docs-oci-datasource-database-db_home"
description: |-
  Provides details about a specific Db Home in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_home
This data source provides details about a specific Db Home resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Database Home.

## Example Usage

```hcl
data "oci_database_db_home" "test_db_home" {
	#Required
	db_home_id = var.db_home_id
}
```

## Argument Reference

The following arguments are supported:

* `db_home_id` - (Required) The Database Home [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_software_image_id` - The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `db_home_location` - The location of the Oracle Database Home.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - The Oracle Database version.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Database Home. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `is_unified_auditing_enabled` - Indicates whether unified autiding is enabled or not.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `one_off_patches` - List of one-off patches for Database Homes.
* `state` - The current state of the Database Home.
* `time_created` - The date and time the Database Home was created.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

