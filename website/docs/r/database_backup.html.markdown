---
layout: "oci"
page_title: "OCI: oci_database_backup"
sidebar_current: "docs-oci-resource-database-backup"
description: |-
  Creates and manages an OCI Backup
---

# oci_database_backup
The `oci_database_backup` resource creates and manages an OCI Backup

Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.


## Example Usage

```hcl
resource "oci_database_backup" "test_backup" {
	#Required
	database_id = "${oci_database_database.test_database.id}"
	display_name = "${var.backup_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The OCID of the database.
* `display_name` - (Required) The user-friendly name for the backup. It does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain that the backup is located in.
* `compartment_id` - The OCID of the compartment.
* `database_edition` - The Oracle Database Edition of the DbSystem on which the backup was taken. 
* `database_id` - The OCID of the database.
* `db_data_size_in_mbs` - Size of the database in mega-bytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. It does not have to be unique.
* `id` - The OCID of the backup.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup starts.
* `type` - The type of backup.

## Import

Backups can be imported using the `id`, e.g.

```
$ terraform import oci_database_backup.test_backup "id"
```
