# oci_database_db_home

### DbHome Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_system_id` - The OCID of the DB System.
* `db_version` - The Oracle database version.
* `display_name` - The user-provided name for the database home. It does not need to be unique.
* `id` - The OCID of the database home.
* `last_patch_history_entry_id` - The OCID of the last patch history. This is updated as soon as a patch operation is started.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.

## DbHome Data Source

The database home.

### Get Operation

Gets information about the specified database home.


The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


### Example Usage

```hcl
data "oci_database_db_home" "test_db_home" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

# oci_database_db_homes

## DbHome DataSource

Gets a list of db_homes.

### List Operation
Gets a list of database homes in the specified DB System and compartment. A database home is a directory where Oracle database software is installed.

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB System.


The following attributes are exported:

* `db_homes` - The list of db_homes.

### Example Usage

```hcl
data "oci_database_db_homes" "test_db_homes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```