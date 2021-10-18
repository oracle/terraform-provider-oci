---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database"
sidebar_current: "docs-oci-resource-database-pluggable_database"
description: |-
  Provides the Pluggable Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_pluggable_database
This resource provides the Pluggable Database resource in Oracle Cloud Infrastructure Database service.

Creates and starts a pluggable database in the specified container database.
Use the [StartPluggableDatabase](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/PluggableDatabase/StartPluggableDatabase) and [StopPluggableDatabase](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/PluggableDatabase/StopPluggableDatabase) APIs to start and stop the pluggable database.


## Example Usage

```hcl
resource "oci_database_pluggable_database" "test_pluggable_database" {
	#Required
	container_database_id = oci_database_database.test_database.id
	pdb_name = var.pluggable_database_pdb_name

	#Optional
	defined_tags = var.pluggable_database_defined_tags
	freeform_tags = {"Department"= "Finance"}
	pdb_admin_password = var.pluggable_database_pdb_admin_password
	should_pdb_admin_account_be_locked = var.pluggable_database_should_pdb_admin_account_be_locked
	tde_wallet_password = var.pluggable_database_tde_wallet_password
}
```

## Argument Reference

The following arguments are supported:

* `container_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `pdb_admin_password` - (Optional) A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
* `pdb_name` - (Required) The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `should_pdb_admin_account_be_locked` - (Optional) The locked mode of the pluggable database admin account. If false, the user needs to provide the PDB Admin Password to connect to it. If true, the pluggable database will be locked and user cannot login to it. 
* `tde_wallet_password` - (Optional) The existing TDE wallet password of the CDB.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggable Database
	* `update` - (Defaults to 20 minutes), when updating the Pluggable Database
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggable Database


## Import

PluggableDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_pluggable_database.test_pluggable_database "id"
```

