---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_snapshot_standby"
sidebar_current: "docs-oci-resource-database-database_snapshot_standby"
description: |-
  Provides the Database Snapshot Standby resource in Oracle Cloud Infrastructure Database service
---

# oci_database_database_snapshot_standby
This resource provides the Database Snapshot Standby resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/DatabaseSnapshotStandby

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

Performs transition from standby database into a snapshot standby and vice versa.
The transition performed based on the current role of the database, if the current role is standby then this operation will convert it to snapshot standby and if the current role is snapshot standby then this operation will convert it to standby.

This operation should be performed on respective standby/snapshot standby database.


## Example Usage

```hcl
resource "oci_database_database_snapshot_standby" "test_database_snapshot_standby" {
	#Required
	database_admin_password = var.database_snapshot_standby_database_admin_password
	database_id = oci_database_database.test_database.id
	standby_conversion_type = var.database_snapshot_standby_standby_conversion_type

	#Optional
	snapshot_duration_in_days = var.database_snapshot_standby_snapshot_duration_in_days
}
```

## Argument Reference

The following arguments are supported:

* `database_admin_password` - (Required) The administrator password of the primary database in this Data Guard association.

	**The password MUST be the same as the primary admin password.** 
* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `snapshot_duration_in_days` - (Optional) SnapshotDurationInDays is the duration in day(s) after which the Snapshot Standby Database will get converted back to Physical Standby. The minimum value of snapshotDurationInDays is 3 days and maximum value is 14 days. Default value will be 7 days if not provided in the Request.

	This field is only applicable if the requested database role is snapshot standby. 
* `standby_conversion_type` - (Required) Defines the conversion type of the standby database. Specify this to convert a physical standby to a snapshot standby and vice versa.

	Valid standbyConversionType:
	* SNAPSHOT 
	* PHYSICAL 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `character_set` - The character set for the database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The Connection strings used to connect to the Oracle Database.
	* `all_connection_strings` - All connection strings to use to connect to the Database.
	* `cdb_default` - Host name based CDB Connection String.
	* `cdb_ip_default` - IP based CDB Connection String.
* `data_guard_group` - Details of Data Guard setup that the given database is part of.  Also includes information about databases part of this Data Guard group and properties for their Data Guard configuration. 
	* `members` - List of Data Guard members, representing each database that is part of Data Guard.
		* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `1 second` 
		* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `102.96 MByte/s` 
		* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database.
		* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system, Cloud VM cluster or VM cluster.
		* `is_active_data_guard_enabled` - True if active Data Guard is enabled.
		* `role` - The role of the reporting database in this Data Guard association.
		* `transport_lag` - The rate at which redo logs are transported between the associated databases.  Example: `1 second` 
		* `transport_lag_refresh` - The date and time when last redo transport has been done.
		* `transport_type` - The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
			* MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
			* MAXIMUM_PERFORMANCE - ASYNC
			* MAXIMUM_PROTECTION - SYNC

			For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.

			**IMPORTANT** - The only transport type currently supported by the Database service is ASYNC. 
	* `protection_mode` - The protection mode of this Data Guard. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `database_management_config` - The configuration of the Database Management service.
	* `management_status` - The status of the Database Management service.
	* `management_type` - The Database Management type.
* `database_software_image_id` - The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `db_backup_config` - Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
	* `auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `auto_backup_window` - Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
	* `auto_full_backup_day` - Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
	* `auto_full_backup_window` - Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
	* `backup_deletion_policy` - This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
	* `backup_destination_details` - Backup destination details.
		* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
	* `run_immediate_full_backup` - If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `db_name` - The database name.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_unique_name` - A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed. 
* `db_workload` - **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service. Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.

	The database workload type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `encryption_key_location_details` - Types of providers supported for managing database encryption keys
	* `hsm_password` - Provide the HSM password as you would in RDBMS for External HSM.
	* `provider_type` - Use 'EXTERNAL' for creating a new database or migrate database key with External HSM.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `is_cdb` - True if the database is a container database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
* `last_backup_duration_in_seconds` - The duration when the latest database backup created.
* `last_backup_timestamp` - The date and time when the latest database backup was created.
* `last_failed_backup_timestamp` - The date and time when the latest database backup failed.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ncharacter_set` - The national character set for the database.
* `pdb_name` - The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `sid_prefix` - Specifies a prefix for the `Oracle SID` of the database to be created. 
* `source_database_point_in_time_recovery_timestamp` - Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Snapshot Standby
	* `update` - (Defaults to 20 minutes), when updating the Database Snapshot Standby
	* `delete` - (Defaults to 20 minutes), when destroying the Database Snapshot Standby


## Import

DatabaseSnapshotStandby can be imported using the `id`, e.g.

```
$ terraform import oci_database_database_snapshot_standby.test_database_snapshot_standby "id"
```

