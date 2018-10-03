# oci_database_db_home

## DbHome Singular DataSource

### DbHome Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_system_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - The Oracle Database version.
* `display_name` - The user-provided name for the database home. The name does not need to be unique.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database home.
* `last_patch_history_entry_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.



### Get Operation
Gets information about the specified database home.

The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


### Example Usage

```hcl
data "oci_database_db_home" "test_db_home" {
	#Required
	db_home_id = "${oci_database_db_home.test_db_home.id}"
}
```
# oci_database_db_homes

## DbHome DataSource

Gets a list of db_homes.

### List Operation
Gets a list of database homes in the specified DB system and compartment. A database home is a directory where Oracle database software is installed.

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.


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
