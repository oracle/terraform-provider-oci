# oci_database_backup

## Backup Resource

### Backup Reference

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



### Create Operation
Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.


The following arguments are supported:

* `database_id` - (Required) The OCID of the database.
* `display_name` - (Required) The user-friendly name for the backup. It does not have to be unique.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_database_backup" "test_backup" {
	#Required
	database_id = "${oci_database_database.test_database.id}"
	display_name = "${var.backup_display_name}"
}
```

# oci_database_backups

## Backup DataSource

Gets a list of backups.

### List Operation
Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.

The following arguments are supported:

* `compartment_id` - (Optional) The compartment OCID.
* `database_id` - (Optional) The OCID of the database.


The following attributes are exported:

* `backups` - The list of backups.

### Example Usage

```hcl
data "oci_database_backups" "test_backups" {

	#Optional
	database_id = "${oci_database_database.test_database.id}"
	// or
	// compartment_id = "${var.compartment_id}"
}
```