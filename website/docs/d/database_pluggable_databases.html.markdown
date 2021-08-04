---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_databases"
sidebar_current: "docs-oci-datasource-database-pluggable_databases"
description: |-
  Provides the list of Pluggable Databases in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_pluggable_databases
This data source provides the list of Pluggable Databases in Oracle Cloud Infrastructure Database service.

Gets a list of the pluggable databases in a database or compartment. You must provide either a `databaseId` or `compartmentId` value.


## Example Usage

```hcl
data "oci_database_pluggable_databases" "test_pluggable_databases" {

	#Optional
	compartment_id = var.compartment_id
	database_id = oci_database_database.test_database.id
	pdb_name = var.pluggable_database_pdb_name
	state = var.pluggable_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `pdb_name` - (Optional) A filter to return only pluggable databases that match the entire name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `pluggable_databases` - The list of pluggable_databases.

### PluggableDatabase Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - Connection strings to connect to an Oracle Pluggable Database. 
	* `all_connection_strings` - All connection strings to use to connect to the pluggable database.
	* `pdb_default` - A host name-based PDB connection string.
	* `pdb_ip_default` - An IP-based PDB connection string.
* `container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pluggable database.
* `is_restricted` - The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it. 
* `lifecycle_details` - Detailed message for the lifecycle state.
* `open_mode` - The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software). 
* `pdb_name` - The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `state` - The current state of the pluggable database.
* `time_created` - The date and time the pluggable database was created.

