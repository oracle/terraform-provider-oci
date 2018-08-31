# oci_database_autonomous_database_backup

## AutonomousDatabaseBackup Resource

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



### Create Operation
Creates a new Autonomous Database backup for the specified database based on the provided request parameters.


The following arguments are supported:

* `autonomous_database_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	display_name = "${var.autonomous_database_backup_display_name}"
}
```


## AutonomousDatabaseBackup Singular DataSource


### Get Operation
Gets information about the specified Autonomous Database backup.

The following arguments are supported:

* `autonomous_database_backup_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.


### Example Usage

```hcl
data "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_backup_id = "${var.autonomous_database_backup_autonomous_database_backup_id}"
}
```
# oci_database_autonomous_database_backups

## AutonomousDatabaseBackup DataSource

Gets a list of autonomous_database_backups.

### List Operation
Gets a list of Autonomous Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.

The following arguments are supported:

* `autonomous_database_id` - (Optional) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


The following attributes are exported:

* `autonomous_database_backups` - The list of autonomous_database_backups.

### Example Usage

```hcl
data "oci_database_autonomous_database_backups" "test_autonomous_database_backups" {

	#Optional
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.autonomous_database_backup_display_name}"
	state = "${var.autonomous_database_backup_state}"
}
```