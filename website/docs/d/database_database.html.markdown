---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database"
sidebar_current: "docs-oci-datasource-database-database"
description: |-
  Provides details about a specific Database in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database
This data source provides details about a specific Database resource in Oracle Cloud Infrastructure Database service.

Gets information about a specific database.

## Example Usage

```hcl
data "oci_database_database" "test_database" {
	#Required
	database_id = "${var.database_id}"
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `character_set` - The character set for the database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The Connection strings used to connect to the Oracle Database.
	* `all_connection_strings` - All connection strings to use to connect to the Database.
	* `cdb_default` - Host name based CDB Connection String.
	* `cdb_ip_default` - IP based CDB Connection String.
* `db_backup_config` - 
	* `auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database home.
* `db_name` - The database name.
* `db_unique_name` - A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed. 
* `db_workload` - The database workload type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `ncharacter_set` - The national character set for the database.
* `pdb_name` - The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.

