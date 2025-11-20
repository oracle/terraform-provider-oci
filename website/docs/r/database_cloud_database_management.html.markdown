---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_database_management"
sidebar_current: "docs-oci-resource-database-cloud_database_management"
description: |-
  Provides the Database Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_database_management
This resource provides the Database Management resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/CloudDatabaseManagement
Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database
Enable / Update / Disable database management for the specified Oracle Database instance.

Database Management requires `USER_NAME`, `PASSWORD_SECRET_ID` and `PRIVATE_END_POINT_ID`.
`database.0.database_management_config` is updated to appropriate managementType and managementStatus for the specified Oracle Database instance.

## Example Usage

```hcl
resource "oci_database_cloud_database_management" "test" {
  database_id           = oci_database_database.test_database.id
  management_type       = var.database_cloud_database_management_details_management_type
  private_end_point_id  = var.database_cloud_database_management_details_private_end_point_id
  service_name          = var.database_cloud_database_management_details_service_name
  credentialdetails {
    user_name           = var.database_cloud_database_management_details_user_name
    password_secret_id  = var.database_cloud_database_management_details_password_secret_id
  }
  enable_management     = var.database_cloud_database_management_details_enable_management
  port = var.cloud_database_management_port
  protocol = var.cloud_database_management_protocol
  role = var.cloud_database_management_role
  ssl_secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `port` - (Optional) The port used to connect to the database.
* `protocol` - (Optional) Protocol used by the database connection.
* `role` - (Optional) The role of the user that will be connecting to the database.
* `ssl_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `private_end_point_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `service_name` - (Required) The name of the Oracle Database service that will be used to connect to the database.
* `management_type` - (Required) (Updatable) Specifies database management type
  enum:
    - `BASIC`
    - `ADVANCED`
* `credentaildetails` - (Required) (Updatable) Credential details to connect to the database
    * `user_name` - Database username
    * `password_secret_id` - Specific database username's password [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `enable_management` - (Required) (Updatable) Use this flag to enable/disable database management

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
		* `data_loss_exposure` - The Data loss exposure is the redo transport lag between the primary and standby databases.   Example: `2 seconds` 
		* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database.
		* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system, Cloud VM cluster or VM cluster.
		* `failover_readiness` - The failover readiness status of the Data Guard member. 
		* `failover_readiness_message` - The message explaining failover readiness status. Example: `This standby database is not failover ready.` 
		* `is_active_data_guard_enabled` - True if active Data Guard is enabled.
		* `role` - The role of the reporting database in this Data Guard association.
		* `switchover_readiness` - The switchover readiness status of the Data Guard member. 
		* `switchover_readiness_message` - The message explaining switchover readiness status. Example: `Address failed checks to avoid extended downtime.` 
		* `time_updated` - The date and time when the last successful Data Guard refresh occurred.
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
	* `backup_deletion_policy` - This defines when the backups will be deleted. - DELETE_IMMEDIATELY option keep the backup for predefined time i.e 72 hours and then delete permanently... - DELETE_AFTER_RETENTION_PERIOD will keep the backups as per the policy defined for database backups.
	* `backup_destination_details` - Backup destination details.
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
    * `azure_encryption_key_id` - Provide the key OCID of a registered Azure key.
    * `hsm_password` - Provide the HSM password as you would in RDBMS for External HSM.
    * `provider_type` - Use 'EXTERNAL' for creating a new database or migrating a database key to an External HSM. Use 'AZURE' for creating a new database or migrating a database key to Azure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `is_cdb` - True if the database is a container database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances. 
* `last_backup_duration_in_seconds` - The duration when the latest database backup created.
* `last_backup_timestamp` - The date and time when the latest database backup was created.
* `last_failed_backup_timestamp` - The date and time when the latest database backup failed.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ncharacter_set` - The national character set for the database.
* `pdb_name` - The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
* `sid_prefix` - Specifies a prefix for the `Oracle SID` of the database to be created.
* `source_database_point_in_time_recovery_timestamp` - Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
* `state` - The current state of the database.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `time_created` - The date and time the database was created.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Cloud Database Management
* `update` - (Defaults to 20 minutes), when updating the Cloud Database Management
* `delete` - (Defaults to 20 minutes), when destroying the Cloud Database Management

## Import

Import is not supported for this resource.