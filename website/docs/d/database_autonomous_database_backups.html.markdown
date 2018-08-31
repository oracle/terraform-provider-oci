---
layout: "oci"
page_title: "OCI: oci_database_autonomous_database_backups"
sidebar_current: "docs-oci-datasource-database-autonomous_database_backups"
description: |-
  Provides a list of AutonomousDatabaseBackups
---

# Data Source: oci_database_autonomous_database_backups
The `oci_database_autonomous_database_backups` data source allows access to the list of OCI autonomous_database_backups

Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.


## Example Usage

```hcl
data "oci_database_autonomous_database_backups" "test_autonomous_database_backups" {

	#Optional
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.autonomous_database_backup_display_name}"
	state = "${var.autonomous_database_backup_state}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Optional) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_backups` - The list of autonomous_database_backups.

### AutonomousDatabaseBackup Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `compartment_id` - The OCID of the compartment.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

