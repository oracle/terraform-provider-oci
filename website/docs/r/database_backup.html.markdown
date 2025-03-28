---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backup"
sidebar_current: "docs-oci-resource-database-backup"
description: |-
	Provides the Backup resource in Oracle Cloud Infrastructure Database service
---

# oci_database_backup
This resource provides the Backup resource in Oracle Cloud Infrastructure Database service.

Creates a new backup in the specified database based on the request parameters you provide. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.


## Example Usage

```hcl
resource "oci_database_backup" "test_backup" {
	#Required
	database_id = oci_database_database.test_database.id
	display_name = var.backup_display_name
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain where the database backup is stored.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_edition` - The Oracle Database Edition that applies to all the databases on the DB system. Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_size_in_gbs` - The size of the database in gigabytes at the time the backup was taken.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `encryption_key_location_details` - Types of providers supported for managing database encryption keys
	* `hsm_password` - Provide the HSM password as you would in RDBMS for External HSM.
	* `provider_type` - Use 'EXTERNAL' for creating a new database or migrate database key with External HSM.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `shape` - Shape of the backup's source database.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `version` - Version of the backup's source database

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 1 hours), when creating the Backup
* `update` - (Defaults to 1 hours), when updating the Backup
* `delete` - (Defaults to 1 hours), when destroying the Backup


## Import

Backups can be imported using the `id`, e.g.

```
$ terraform import oci_database_backup.test_backup "id"
```