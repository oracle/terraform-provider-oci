---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database"
sidebar_current: "docs-oci-datasource-database-pluggable_database"
description: |-
  Provides details about a specific Pluggable Database in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_pluggable_database
This data source provides details about a specific Pluggable Database resource in Oracle Cloud Infrastructure Database service.

Gets information about a specific pluggable database

## Example Usage

```hcl
data "oci_database_pluggable_database" "test_pluggable_database" {
	#Required
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
}
```

## Argument Reference

The following arguments are supported:

* `pluggable_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - Connection strings to connect to an Oracle Pluggable Database. 
	* `all_connection_strings` - All connection strings to use to connect to the Pluggable Database.
	* `pdb_default` - Host name based PDB Connection String.
	* `pdb_ip_default` - IP based PDB Connection String.
* `container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pluggable database.
* `is_restricted` - The restricted mode of pluggableDatabase. If a pluggableDatabase is opened in restricted mode, the user needs both Create a session and restricted session privileges to connect to it. 
* `lifecycle_details` - Detailed message for the lifecycle state.
* `open_mode` - The mode that pluggableDatabase is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend.
* `pdb_name` - The name for the pluggable database. The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `state` - The current state of the pluggable database.
* `time_created` - The date and time the pluggable database was created

