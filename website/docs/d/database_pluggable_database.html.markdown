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

Gets information about the specified pluggable database.

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

