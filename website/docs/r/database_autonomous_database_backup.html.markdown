---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_backup"
sidebar_current: "docs-oci-resource-database-autonomous_database_backup"
description: |-
  Provides the Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_backup
This resource provides the Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service.

Creates a new Autonomous Database backup for the specified database based on the provided request parameters.


## Example Usage

```hcl
resource "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	display_name = "${var.autonomous_database_backup_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_size_in_tbs` - The size of the database in terabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `is_restorable` - Indicates whether the backup can be used to restore the associated Autonomous Database.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

## Import

AutonomousDatabaseBackups can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database_backup.test_autonomous_database_backup "id"
```

