---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backup_destinations"
sidebar_current: "docs-oci-datasource-database-backup_destinations"
description: |-
  Provides the list of Backup Destinations in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_backup_destinations
This data source provides the list of Backup Destinations in Oracle Cloud Infrastructure Database service.

Gets a list of backup destinations in the specified compartment.


## Example Usage

```hcl
data "oci_database_backup_destinations" "test_backup_destinations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	type = var.backup_destination_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `type` - (Optional) A filter to return only resources that match the given type of the Backup Destination.


## Attributes Reference

The following attributes are exported:

* `backup_destinations` - The list of backup_destinations.

### BackupDestination Reference

The following attributes are exported:

* `associated_databases` - List of databases associated with the backup destination.
	* `db_name` - The display name of the database that is associated with the backup destination.
	* `id` - The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_string` - For a RECOVERY_APPLIANCE backup destination, the connection string for connecting to the Recovery Appliance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name of the backup destination.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text 
* `local_mount_point_path` - The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes. 
* `nfs_mount_type` - NFS Mount type for backup destination.
* `nfs_server` - Host names or IP addresses for NFS Auto mount.
* `nfs_server_export` - Specifies the directory on which to mount the file system
* `state` - The current lifecycle state of the backup destination.
* `time_at_which_storage_details_are_updated` - The time when the total storage size and the utilized storage size of the backup destination are updated.
* `time_created` - The date and time the backup destination was created.
* `total_storage_size_in_gbs` - The total storage size of the backup destination in GBs, rounded to the nearest integer.
* `type` - Type of the backup destination.
* `utilized_storage_size_in_gbs` - The total amount of space utilized on the backup destination (in GBs), rounded to the nearest integer.
* `vpc_users` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.

