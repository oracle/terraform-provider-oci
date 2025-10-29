---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_backups"
sidebar_current: "docs-oci-datasource-database-autonomous_database_backups"
description: |-
  Provides the list of Autonomous Database Backups in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_backups
This data source provides the list of Autonomous Database Backups in Oracle Cloud Infrastructure Database service.

Gets a list of Autonomous AI Database backups based on either the `autonomousDatabaseId` or `compartmentId` specified as a query parameter.


## Example Usage

```hcl
data "oci_database_autonomous_database_backups" "test_autonomous_database_backups" {

	#Optional
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	backup_destination_id = oci_database_backup_destination.test_backup_destination.id
	compartment_id = var.compartment_id
	display_name = var.autonomous_database_backup_display_name
	infrastructure_type = var.autonomous_database_backup_infrastructure_type
	key_store_id = oci_database_key_store.test_key_store.id
	state = var.autonomous_database_backup_state
	type = var.autonomous_database_backup_type
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Optional) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `backup_destination_id` - (Optional) A filter to return only resources that have the given backup destination id.
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `infrastructure_type` - (Optional) A filter to return only resources that match the given Infrastructure Type.
* `key_store_id` - (Optional) A filter to return only resources that have the given key store id.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `type` - (Optional) A filter to return only backups that matches with the given type of Backup.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_backups` - The list of autonomous_database_backups.

### AutonomousDatabaseBackup Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database.
* `backup_destination_details` - Backup destination details
	* `backup_retention_policy_on_terminate` - Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
	* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	* `internet_proxy` - Proxy URL to connect to object store.
	* `is_remote` - Indicates whether the backup destination is cross-region or local.
	* `is_retention_lock_enabled` - Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
	* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
	* `type` - Type of the database backup destination.
	* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_size_in_tbs` - The size of the database in terabytes at the time the backup was taken. 
* `db_version` - A valid Oracle AI Database version for Autonomous AI Database.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database backup.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `is_restorable` - Indicates whether the backup can be used to restore the associated Autonomous AI Database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `region` - Name of the region in which backup is taken in.
* `retention_period_in_days` - Retention period, in days, for long-term backups
* `size_in_tbs` - The backup size in terrabytes (TB).
* `source_database_details` - Source Autonomous Database details.
	* `autonomous_container_database_customer_contacts` - Customer Contacts for the Autonomous Container Database. Setting this to an empty list removes all customer contacts. 
		* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
	* `autonomous_container_database_display_name` - The user-provided name for the Autonomous Container Database.
	* `autonomous_container_database_dst_file_version` - DST Time-Zone File version of the Autonomous Container Database.
	* `autonomous_container_database_name` - Autonomous Container Database name.
	* `autonomous_database_customer_contacts` - Customer Contacts for the Autonomous database.
		* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
	* `autonomous_database_name` - Autonomous Database's name.
	* `autonomous_vm_cluster_display_name` - Autonomous VM cluster's user-friendly name.
	* `db_workload` - The Autonomous Database workload type. The following values are valid:
		* OLTP - indicates an Autonomous Transaction Processing database
		* DW - indicates an Autonomous Data Warehouse database
		* AJD - indicates an Autonomous JSON Database
		* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.

		This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, privateEndpointLabel, nsgIds, dbVersion, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
* `state` - The current state of the backup.
* `time_available_till` - Timestamp until when the backup will be available
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.

