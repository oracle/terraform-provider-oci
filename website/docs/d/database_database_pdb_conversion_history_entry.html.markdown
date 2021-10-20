---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_pdb_conversion_history_entry"
sidebar_current: "docs-oci-datasource-database-database_pdb_conversion_history_entry"
description: |-
  Provides details about a specific Database Pdb Conversion History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_pdb_conversion_history_entry
This data source provides details about a specific Database Pdb Conversion History Entry resource in Oracle Cloud Infrastructure Database service.

Gets the details of operations performed to convert the specified database from non-container (non-CDB) to pluggable (PDB).


## Example Usage

```hcl
data "oci_database_database_pdb_conversion_history_entry" "test_database_pdb_conversion_history_entry" {
	#Required
	database_id = oci_database_database.test_database.id
	pdb_conversion_history_entry_id = oci_database_pdb_conversion_history_entry.test_pdb_conversion_history_entry.id
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `pdb_conversion_history_entry_id` - (Required) The database conversion history [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `action` - The operations used to convert a non-container database to a pluggable database.
	* Use `PRECHECK` to run a pre-check operation on non-container database prior to converting it into a pluggable database.
	* Use `CONVERT` to convert a non-container database into a pluggable database.
	* Use `SYNC` if the non-container database was manually converted into a pluggable database using the dbcli command-line utility. Databases may need to be converted manually if the CONVERT action fails when converting a non-container database using the API.
	* Use `SYNC_ROLLBACK` if the conversion of a non-container database into a pluggable database was manually rolled back using the dbcli command line utility. Conversions may need to be manually rolled back if the CONVERT action fails when converting a non-container database using the API. 
* `additional_cdb_params` - Additional container database parameter. 
* `cdb_name` - The database name. The name must begin with an alphabetic character and can contain a maximum of 8 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database conversion history.
* `lifecycle_details` - Additional information about the current lifecycle state for the conversion operation.
* `source_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `state` - Status of an operation performed during the conversion of a non-container database to a pluggable database.
* `target` - The target container database of the pluggable database created by the database conversion operation. Currently, the database conversion operation only supports creating the pluggable database in a new container database.
	* Use `NEW_DATABASE` to specify that the pluggable database be created within a new container database in the same database home. 
* `target_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `time_ended` - The date and time when the database conversion operation ended.
* `time_started` - The date and time when the database conversion operation started.

