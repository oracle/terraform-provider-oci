# oci_database_database

## Database Data Source

An Oracle database on a DB System. For more information, see [Managing Oracle Databases](https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/overview.htm).

### Get Operation

Gets information about a specific database.


The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

### Database Reference

The following attributes are exported:

* `character_set` - The character set for the database.
* `compartment_id` - The OCID of the compartment.
* `db_backup_config` - 
	* `auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
* `db_home_id` - The OCID of the database home.
* `db_name` - The database name.
* `db_unique_name` - A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed. 
* `db_workload` - Database workload type.
* `id` - The OCID of the database.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `ncharacter_set` - The national character set for the database.
* `pdb_name` - Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.

### Example Usage

```hcl
data "oci_database_database" "test_database" {
	#Required
	database_id = "${var.database_id}"
}
```


# oci_database_databases

## Database DataSource

Gets a list of databases.

### List Operation
Gets a list of the databases in the specified database home.

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_home_id` - (Required) A database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `databases` - The list of databases.

### Example Usage

```hcl
data "oci_database_databases" "test_databases" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_home_id = "${oci_database_db_home.test_db_home.id}"
}
```