---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_databases_remote_clone"
sidebar_current: "docs-oci-resource-database-pluggable_databases_remote_clone"
description: |-
  Provides the Pluggable Databases Remote Clone resource in Oracle Cloud Infrastructure Database service
---

# oci_database_pluggable_databases_remote_clone
This resource provides the Pluggable Databases Remote Clone resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/PluggableDatabasesRemoteClone

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

**Deprecated.** Use [CreatePluggableDatabase](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/PluggableDatabase/CreatePluggableDatabase) for Pluggable Database RemoteClone Operation.
Clones a pluggable database (PDB) to a different database from the source PDB. The cloned PDB will be started upon completion of the clone operation. The source PDB must be in the `READ_WRITE` openMode when performing the clone.
For Exadata Cloud@Customer instances, the source pluggable database (PDB) must be on the same Exadata Infrastructure as the target container database (CDB) to create a remote clone.


## Example Usage

```hcl
resource "oci_database_pluggable_databases_remote_clone" "test_pluggable_databases_remote_clone" {
	#Required
	cloned_pdb_name = var.pluggable_databases_remote_clone_cloned_pdb_name
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
	source_container_db_admin_password = var.pluggable_databases_remote_clone_source_container_db_admin_password
	target_container_database_id = oci_database_database.test_database.id

	#Optional
	pdb_admin_password = var.pluggable_databases_remote_clone_pdb_admin_password
	should_pdb_admin_account_be_locked = var.pluggable_databases_remote_clone_should_pdb_admin_account_be_locked
	target_tde_wallet_password = var.pluggable_databases_remote_clone_target_tde_wallet_password
}
```

## Argument Reference

The following arguments are supported:

* `cloned_pdb_name` - (Required) The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `pdb_admin_password` - (Optional) A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
* `pluggable_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `should_pdb_admin_account_be_locked` - (Optional) The locked mode of the pluggable database admin account. If false, the user needs to provide the PDB Admin Password to connect to it. If true, the pluggable database will be locked and user cannot login to it. 
* `source_container_db_admin_password` - (Required) The DB system administrator password of the source CDB.
* `target_container_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target CDB
* `target_tde_wallet_password` - (Optional) The existing TDE wallet password of the target CDB.


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
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
* `lifecycle_details` - Detailed message for the lifecycle state.
* `open_mode` - **Deprecated.** Use [PluggableDatabaseNodeLevelDetails](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/PluggableDatabaseNodeLevelDetails) for OpenMode details. The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software). 
* `pdb_name` - The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `pdb_node_level_details` - Pluggable Database Node Level Details. Example: [{"nodeName" : "node1", "openMode" : "READ_WRITE"}, {"nodeName" : "node2", "openMode" : "READ_ONLY"}] 
	* `node_name` - The Node name of the Database Instance.
	* `open_mode` - The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software). 
* `pluggable_database_management_config` - The configuration of the Pluggable Database Management service.
	* `management_status` - The status of the Pluggable Database Management service.
* `refreshable_clone_config` - Pluggable Database Refreshable Clone Configuration.
	* `is_refreshable_clone` - Indicates whether the Pluggable Database is a refreshable clone.
* `state` - The current state of the pluggable database.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the pluggable database was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggable Databases Remote Clone
	* `update` - (Defaults to 20 minutes), when updating the Pluggable Databases Remote Clone
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggable Databases Remote Clone


## Import

Import is not supported for this resource.

