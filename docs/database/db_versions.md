
# oci_database_db_versions

## DbVersion DataSource

Gets a list of db_versions.

### List Operation
Gets a list of supported Oracle database versions.
The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB system OCID. If provided, filters the results to the set of database versions which are supported for the DB system.
* `db_system_shape` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape.


The following attributes are exported:

* `db_versions` - The list of db_versions.

### Example Usage

```hcl
data "oci_database_db_versions" "test_db_versions" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	db_system_id = "${oci_database_db_system.test_db_system.id}"
	db_system_shape = "${var.db_version_db_system_shape}"
}
```
### DbVersion Reference

The following attributes are exported:

* `supports_pdb` - True if this version of the Oracle database software supports pluggable dbs.
* `version` - A valid Oracle database version.
