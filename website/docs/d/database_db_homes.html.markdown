---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_homes"
sidebar_current: "docs-oci-datasource-database-db_homes"
description: |-
  Provides the list of Db Homes in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_homes
This data source provides the list of Db Homes in Oracle Cloud Infrastructure Database service.

Lists the Database Homes in the specified DB system and compartment. A Database Home is a directory where Oracle Database software is installed.


## Example Usage

```hcl
data "oci_database_db_homes" "test_db_homes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	backup_id = oci_database_backup.test_backup.id
	db_system_id = oci_database_db_system.test_db_system.id
	db_version {
	}
	display_name = var.db_home_display_name
	state = var.db_home_state
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup. Specify a backupId to list only the DB systems or DB homes that support creating a database using this backup in this compartment.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Applicable when source=DATABASE | DB_BACKUP | NONE) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
* `db_version` - (Applicable when source=NONE | VM_CLUSTER_NEW) A filter to return only DB Homes that match the specified dbVersion.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `vm_cluster_id` - (Applicable when source=VM_CLUSTER_BACKUP | VM_CLUSTER_NEW) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.


## Attributes Reference

The following attributes are exported:

* `db_homes` - The list of db_homes.

### DbHome Reference

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

