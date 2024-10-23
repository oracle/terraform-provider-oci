---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backups"
sidebar_current: "docs-oci-datasource-database-backups"
description: |-
  Provides the list of Backups in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_backups
This data source provides the list of Backups in Oracle Cloud Infrastructure Database service.

Gets a list of backups based on the `databaseId` or `compartmentId` specified. Either one of these query parameters must be provided.


## Example Usage

```hcl
data "oci_database_backups" "test_backups" {

	#Optional
	backup_destination_type = var.backup_backup_destination_type
	compartment_id = var.compartment_id
	database_id = oci_database_database.test_database.id
	shape_family = var.backup_shape_family
	state = var.backup_state
	time_expiry_scheduled_greater_than_or_equal_to = var.backup_time_expiry_scheduled_greater_than_or_equal_to
	time_expiry_scheduled_less_than = var.backup_time_expiry_scheduled_less_than
	type = var.backup_type
	version = var.backup_version
}
```

## Argument Reference

The following arguments are supported:

* `backup_destination_type` - (Optional) A filter to return only resources that match the given backup destination type.
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `shape_family` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape family.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `time_expiry_scheduled_greater_than_or_equal_to` - (Optional) The start of date-time range of expiration for the long term backups to be fetched.
* `time_expiry_scheduled_less_than` - (Optional) The end of date-time range of expiration for the long term backups to be fetched.
* `type` - (Optional) A filter to return only backups that matches with the given type of Backup.
* `version` - (Optional) A filter to return only resources that match the given database version.


## Attributes Reference

The following attributes are exported:

* `backups` - The list of backups.

### Backup Reference

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

