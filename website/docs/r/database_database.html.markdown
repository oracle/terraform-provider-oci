---
layout: "oci"
page_title: "OCI: oci_database_database"
sidebar_current: "docs-oci-resource-database-database"
description: |-
  Creates and manages an OCI Database
---

# oci_database_database
The `oci_database_database` resource creates and manages an OCI Database



## Example Usage

```hcl
resource "oci_database_database" "test_database" {
}
```

## Argument Reference

The following arguments are supported:



** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `character_set` - The character set for the database.
* `compartment_id` - The OCID of the compartment.
* `db_backup_config` - 
	* `auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
* `db_home_id` - The OCID of the database home.
* `db_name` - The database name.
* `db_unique_name` - A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed. 
* `db_workload` - Database workload type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the database.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `ncharacter_set` - The national character set for the database.
* `pdb_name` - Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.

## Import

Databases can be imported using the `id`, e.g.

```
$ terraform import oci_database_database.test_database "id"
```
