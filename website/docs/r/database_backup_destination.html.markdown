---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backup_destination"
sidebar_current: "docs-oci-resource-database-backup_destination"
description: |-
  Provides the Backup Destination resource in Oracle Cloud Infrastructure Database service
---

# oci_database_backup_destination
This resource provides the Backup Destination resource in Oracle Cloud Infrastructure Database service.

Creates a backup destination.


## Example Usage

```hcl
resource "oci_database_backup_destination" "test_backup_destination" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.backup_destination_display_name}"
	type = "${var.backup_destination_type}"

	#Optional
	connection_string = "${var.backup_destination_connection_string}"
	defined_tags = "${var.backup_destination_defined_tags}"
	freeform_tags = {"Department"= "Finance"}
	local_mount_point_path = "${var.backup_destination_local_mount_point_path}"
	vpc_users = "${var.backup_destination_vpc_users}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_string` - (Required when type=RECOVERY_APPLIANCE) (Updatable) The connection string for connecting to the Recovery Appliance.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-provided name of the backup destination.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `local_mount_point_path` - (Required when type=NFS) (Updatable) The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes. 
* `type` - (Required) Type of the backup destination.
* `vpc_users` - (Required when type=RECOVERY_APPLIANCE) (Updatable) The Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `state` - The current lifecycle state of the backup destination.
* `time_created` - The date and time the backup destination was created.
* `type` - Type of the backup destination.
* `vpc_users` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.

## Import

BackupDestinations can be imported using the `id`, e.g.

```
$ terraform import oci_database_backup_destination.test_backup_destination "id"
```

