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
Pluggable Database can be created using different operations (e.g. LocalClone, RemoteClone, Relocate ) with this API.
Use the [StartPluggableDatabase](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/PluggableDatabase/StartPluggableDatabase) and [StopPluggableDatabase](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/PluggableDatabase/StopPluggableDatabase) APIs to start and stop the pluggable database.


## Example Usage

```hcl
resource "oci_database_pluggable_database" "test_pluggable_database" {
	#Required
	container_database_id = oci_database_database.test_database.id
	pdb_name = var.pluggable_database_pdb_name

	#Optional
	container_database_admin_password = var.pluggable_database_container_database_admin_password
	defined_tags = var.pluggable_database_defined_tags
	freeform_tags = {"Department"= "Finance"}
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	pdb_admin_password = var.pluggable_database_pdb_admin_password
	pdb_creation_type_details {
		#Required
		creation_type = var.pluggable_database_pdb_creation_type_details_creation_type
		source_pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id

		#Optional
		dblink_user_password = var.pluggable_database_pdb_creation_type_details_dblink_user_password
		dblink_username = var.pluggable_database_pdb_creation_type_details_dblink_username
		is_thin_clone = var.pluggable_database_pdb_creation_type_details_is_thin_clone
		refreshable_clone_details {

			#Optional
			is_refreshable_clone = var.pluggable_database_pdb_creation_type_details_refreshable_clone_details_is_refreshable_clone
		}
		source_container_database_admin_password = var.pluggable_database_pdb_creation_type_details_source_container_database_admin_password
	}
	should_create_pdb_backup = var.pluggable_database_should_create_pdb_backup
	should_pdb_admin_account_be_locked = var.pluggable_database_should_pdb_admin_account_be_locked
	tde_wallet_password = var.pluggable_database_tde_wallet_password
}
```

## Argument Reference

The following arguments are supported:

* `container_database_admin_password` - (Optional) The DB system administrator password of the Container Database.
* `container_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `pdb_admin_password` - (Optional) A strong password for PDB Admin. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
* `pdb_creation_type_details` - (Optional) The Pluggable Database creation type. Use `LOCAL_CLONE_PDB` for creating a new PDB using Local Clone on Source Pluggable Database. This will Clone and starts a pluggable database (PDB) in the same database (CDB) as the source PDB. The source PDB must be in the `READ_WRITE` openMode to perform the clone operation. isThinClone options are supported only for Exadata VM cluster on Exascale Infrastructure. Use `REMOTE_CLONE_PDB` for creating a new PDB using Remote Clone on Source Pluggable Database. This will Clone a pluggable database (PDB) to a different database from the source PDB. The cloned PDB will be started upon completion of the clone operation. The source PDB must be in the `READ_WRITE` openMode when performing the clone. For Exadata Cloud@Customer instances, the source pluggable database (PDB) must be on the same Exadata Infrastructure as the target container database (CDB) to create a remote clone. isThinClone options are supported only for Exadata VM cluster on Exascale Infrastructure.

	Use `RELOCATE_PDB` for relocating the Pluggable Database from Source CDB and creating it in target CDB. This will relocate a pluggable database (PDB) to a different database from the source PDB. The source PDB must be in the `READ_WRITE` openMode when performing the relocate. 
	* `creation_type` - (Required) The Pluggable Database creation type.
	* `dblink_user_password` - (Applicable when creation_type=RELOCATE_PDB | REMOTE_CLONE_PDB) The DB link user password.
	* `dblink_username` - (Applicable when creation_type=RELOCATE_PDB | REMOTE_CLONE_PDB) The name of the DB link user.
	* `is_thin_clone` - (Applicable when creation_type=LOCAL_CLONE_PDB | REMOTE_CLONE_PDB) True if Pluggable Database needs to be thin cloned and false if Pluggable Database needs to be thick cloned.
	* `refreshable_clone_details` - (Applicable when creation_type=REMOTE_CLONE_PDB) Parameters for creating Pluggable Database Refreshable Clone. **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
		* `is_refreshable_clone` - (Applicable when creation_type=REMOTE_CLONE_PDB) Indicates whether Pluggable Database is a refreshable clone.
	* `source_container_database_admin_password` - (Required when creation_type=RELOCATE_PDB | REMOTE_CLONE_PDB) The DB system administrator password of the source Container Database.
	* `source_pluggable_database_id` - (Required) The OCID of the Source Pluggable Database.
* `pdb_name` - (Required) The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `should_create_pdb_backup` - (Optional) Indicates whether to take Pluggable Database Backup after the operation.
* `should_pdb_admin_account_be_locked` - (Optional) The locked mode of the pluggable database admin account. If false, the user needs to provide the PDB Admin Password to connect to it. If true, the pluggable database will be locked and user cannot login to it. 
* `tde_wallet_password` - (Optional) The existing TDE wallet password of the CDB.
* `convert_to_regular_trigger` - (Optional) (Updatable) An optional property when incremented triggers Convert To Regular. Could be set to any integer value.
* `refresh_trigger` - (Optional) (Updatable) An optional property when incremented triggers Refresh. Could be set to any integer value.
* `rotate_key_trigger` - (Optional) (Updatable) An optional property when incremented triggers Rotate Key. Could be set to any integer value.


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
* `time_created` - The date and time the pluggable database was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggable Database
	* `update` - (Defaults to 20 minutes), when updating the Pluggable Database
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggable Database


## Import

PluggableDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_pluggable_database.test_pluggable_database "id"
```

