---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_backup"
sidebar_current: "docs-oci-datasource-database-autonomous_database_backup"
description: |-
  Provides details about a specific Autonomous Database Backup in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_backup
This data source provides details about a specific Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous Database backup.

## Example Usage

```hcl
data "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_backup_id = oci_database_autonomous_database_backup.test_autonomous_database_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_backup_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `backup_destination_details` - Backup destination details
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	* `internet_proxy` - Proxy URL to connect to object store.
	* `type` - Type of the database backup destination.
	* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_size_in_tbs` - The size of the database in terabytes at the time the backup was taken. 
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `is_restorable` - Indicates whether the backup can be used to restore the associated Autonomous Database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `retention_period_in_days` - Retention period, in days, for long-term backups
* `size_in_tbs` - The backup size in terrabytes (TB).
* `state` - The current state of the backup.
* `time_available_till` - Timestamp until when the backup will be available
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.

