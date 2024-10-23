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

	#Optional
	retention_period_in_days = var.backup_retention_period_in_days
	retention_period_in_years = var.backup_retention_period_in_years
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.
* `retention_period_in_days` - (Optional) (Updatable) The retention period of the long term backup in days.
* `retention_period_in_years` - (Optional) (Updatable) The retention period of the long term backup in years.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain where the database backup is stored.
* `backup_destination_type` - Type of the backup destination.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_edition` - The Oracle Database edition of the DB system from which the database backup was taken. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_size_in_gbs` - The size of the database in gigabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `is_using_oracle_managed_keys` - True if Oracle Managed Keys is required for restore of the backup.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `retention_period_in_days` - The retention period of the long term backup in days.
* `retention_period_in_years` - The retention period of the long term backup in years.
* `secondary_kms_key_ids` - List of OCIDs of the key containers used as the secondary encryption key in database transparent data encryption (TDE) operations.
* `shape` - Shape of the backup's source database.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_expiry_scheduled` - Expiration time of the long term database backup.
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

