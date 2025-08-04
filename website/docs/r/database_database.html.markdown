---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database"
sidebar_current: "docs-oci-resource-database-database"
description: |-
  Provides the Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_database
This resource provides the Database resource in Oracle Cloud Infrastructure Database service.

Creates a new database in the specified Database Home. If the database version is provided, it must match the version of the Database Home. Applies only to Exadata systems.


## Example Usage

```hcl
resource "oci_database_database" "test_database" {
	#Required
	database {

		#Optional
		admin_password = var.database_database_admin_password
		backup_id = oci_database_backup.test_backup.id
		backup_tde_password = var.database_database_backup_tde_password
		character_set = var.database_database_character_set
		database_admin_password = var.database_database_database_admin_password
		database_software_image_id = oci_database_database_software_image.test_database_software_image.id
		db_backup_config {

			#Optional
			auto_backup_enabled = var.database_database_db_backup_config_auto_backup_enabled
			auto_backup_window = var.database_database_db_backup_config_auto_backup_window
			auto_full_backup_day = var.database_database_db_backup_config_auto_full_backup_day
			auto_full_backup_window = var.database_database_db_backup_config_auto_full_backup_window
			backup_deletion_policy = var.database_database_db_backup_config_backup_deletion_policy
			backup_destination_details {

				#Optional
				backup_retention_policy_on_terminate = var.database_database_db_backup_config_backup_destination_details_backup_retention_policy_on_terminate
				dbrs_policy_id = oci_identity_policy.test_policy.id
				id = var.database_database_db_backup_config_backup_destination_details_id
				is_remote = var.database_database_db_backup_config_backup_destination_details_is_remote
				is_retention_lock_enabled = var.database_database_db_backup_config_backup_destination_details_is_retention_lock_enabled
				remote_region = var.database_database_db_backup_config_backup_destination_details_remote_region
				type = var.database_database_db_backup_config_backup_destination_details_type
			}
			recovery_window_in_days = var.database_database_db_backup_config_recovery_window_in_days
			run_immediate_full_backup = var.database_database_db_backup_config_run_immediate_full_backup
		}
		db_name = var.database_database_db_name
		db_unique_name = var.database_database_db_unique_name
		db_workload = var.database_database_db_workload
		defined_tags = var.database_database_defined_tags
		encryption_key_location_details {
			#Required
			provider_type = var.database_database_encryption_key_location_details_provider_type

			#Optional
			azure_encryption_key_id = oci_kms_key.test_key.id
			hsm_password = var.database_database_encryption_key_location_details_hsm_password
		}
		freeform_tags = var.database_database_freeform_tags
		is_active_data_guard_enabled = var.database_database_is_active_data_guard_enabled
		key_store_id = oci_database_key_store.test_key_store.id
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		ncharacter_set = var.database_database_ncharacter_set
		pdb_name = var.database_database_pdb_name
		pluggable_databases = var.database_database_pluggable_databases
		protection_mode = var.database_database_protection_mode
		sid_prefix = var.database_database_sid_prefix
		source_database_id = oci_database_database.test_database.id
		source_encryption_key_location_details {
			#Required
			provider_type = var.database_database_source_encryption_key_location_details_provider_type

			#Optional
			azure_encryption_key_id = oci_kms_key.test_key.id
			hsm_password = var.database_database_source_encryption_key_location_details_hsm_password
		}
		
		storage_size_details {
			#Required
			data_storage_size_in_gb = var.database_database_storage_size_details_data_storage_size_in_gb
			reco_storage_size_in_gbs = var.database_database_storage_size_details_reco_storage_size_in_gbs
		}
		source_tde_wallet_password = var.database_database_source_tde_wallet_password
		tde_wallet_password = var.database_database_tde_wallet_password
		transport_type = var.database_database_transport_type
		vault_id = oci_kms_vault.test_vault.id
	}
	db_home_id = oci_database_db_home.test_db_home.id
	source = var.database_source

	#Optional
	db_version = var.database_db_version
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
}
```

## Argument Reference

The following arguments are supported:

* `database` - (Required) (Updatable) Details for creating a database.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `admin_password` - (Required when source=DB_BACKUP | NONE) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	* `backup_id` - (Required when source=DB_BACKUP) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `backup_tde_password` - (Applicable when source=DB_BACKUP) The password to open the TDE wallet.
	* `character_set` - (Applicable when source=NONE) The character set for the database.  The default is AL32UTF8. Allowed values are:

		AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
	* `database_admin_password` - (Required when source=DATAGUARD) The administrator password of the primary database in this Data Guard association.

		**The password MUST be the same as the primary admin password.** 
	* `database_software_image_id` - (Applicable when source=NONE) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	* `db_backup_config` - (Applicable when source=NONE) (Updatable) Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
		* `auto_backup_enabled` - (Applicable when source=NONE) (Updatable) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
		* `auto_backup_window` - (Applicable when source=NONE) (Updatable) Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
		* `auto_full_backup_day` - (Applicable when source=NONE) Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
		* `auto_full_backup_window` - (Applicable when source=NONE) Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
		* `backup_deletion_policy` - (Applicable when source=NONE) This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
		* `backup_destination_details` - (Applicable when source=NONE) Backup destination details.
			* `backup_retention_policy_on_terminate` - (Applicable when source=NONE) Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
			* `dbrs_policy_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
			* `id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
			* `is_remote` - (Applicable when source=NONE) Indicates whether the backup destination is cross-region or local.
			* `is_retention_lock_enabled` - (Applicable when source=NONE) Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
			* `remote_region` - (Applicable when source=NONE) The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
			* `type` - (Required when source=NONE) Type of the database backup destination.
		* `recovery_window_in_days` - (Applicable when source=NONE) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
		* `run_immediate_full_backup` - (Applicable when source=NONE) If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `db_name` - (Required when source=DB_BACKUP | NONE) The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	* `db_unique_name` - (Optional) Specifies the `DB_UNIQUE_NAME` of the peer database to be created. 
	* `db_workload` - (Applicable when source=NONE) **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service. Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.

		The database workload type. 
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
	* `encryption_key_location_details` - (Applicable when source=NONE) Types of providers supported for managing database encryption keys
		* `azure_encryption_key_id` - (Required when provider_type=AZURE) Provide the key OCID of a registered Azure key.
		* `hsm_password` - (Required when provider_type=EXTERNAL) Provide the HSM password as you would in RDBMS for External HSM.
		* `provider_type` - (Required) Use 'EXTERNAL' for creating a new database or migrating a database key to an External HSM. Use 'AZURE' for creating a new database or migrating a database key to Azure. 
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `is_active_data_guard_enabled` - (Applicable when source=DATAGUARD) True if active Data Guard is enabled.
	* `key_store_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	* `kms_key_id` - (Applicable when source=NONE) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Applicable when source=NONE) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
	* `ncharacter_set` - (Applicable when source=NONE) The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
	* `pdb_name` - (Applicable when source=NONE) The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	* `pluggable_databases` - (Applicable when source=DB_BACKUP) The list of pluggable databases that needs to be restored into new database.
	* `protection_mode` - (Required when source=DATAGUARD) The protection mode of this Data Guard. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
	* `sid_prefix` - (Optional) Specifies a prefix for the `Oracle SID` of the database to be created. 
	* `source_database_id` - (Required when source=DATAGUARD) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source database.
	* `source_encryption_key_location_details` - (Applicable when source=DATAGUARD | DB_BACKUP) Types of providers supported for managing database encryption keys
		* `azure_encryption_key_id` - (Required when provider_type=AZURE) Provide the key OCID of a registered Azure key.
		* `hsm_password` - (Required when provider_type=EXTERNAL) Provide the HSM password as you would in RDBMS for External HSM.
		* `provider_type` - (Required) Use 'EXTERNAL' for creating a new database or migrating a database key to an External HSM. Use 'AZURE' for creating a new database or migrating a database key to Azure. 
	* `source_tde_wallet_password` - (Required when source=DATAGUARD) The TDE wallet password of the source database specified by 'sourceDatabaseId'.
	* `storage_size_details` - (Optional) The database storage size details. This database option is supported for the Exadata VM cluster on Exascale Infrastructure. 
		* `data_storage_size_in_gb` - (Required) (Updatable) The DATA storage size, in gigabytes, that is applicable for the database. 
		* `reco_storage_size_in_gbs` - (Required) (Updatable) The RECO storage size, in gigabytes, that is applicable for the database. 
	* `tde_wallet_password` - (Applicable when source=NONE) The optional password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	* `transport_type` - (Required when source=DATAGUARD) The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
		* MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
		* MAXIMUM_PERFORMANCE - ASYNC
		* MAXIMUM_PROTECTION - SYNC

		For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.

		**IMPORTANT** - The only transport type currently supported by the Database service is ASYNC. 
	* `vault_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `db_home_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `db_version` - (Optional) A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
* `source` - (Required) The source of the database: Use `NONE` for creating a new database. Use `DB_BACKUP` for creating a new database by restoring from a backup. Use `DATAGUARD` for creating a new STANDBY database for a Data Guard setup.. The default is `NONE`. 
* `set_key_version_trigger` - (Optional) (Updatable) An optional property when incremented triggers Set Key Version. Could be set to any integer value.


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
		* `backup_retention_policy_on_terminate` - Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
		* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `is_remote` - Indicates whether the backup destination is cross-region or local.
		* `is_retention_lock_enabled` - Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
		* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - Type of the database backup destination.
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
* `storage_size_details` - The database storage size details. This database option is supported for the Exadata VM cluster on Exascale Infrastructure. 
	* `data_storage_size_in_gb` - The DATA storage size, in gigabytes, that is applicable for the database. 
	* `reco_storage_size_in_gbs` - The RECO storage size, in gigabytes, that is applicable for the database. 
	* `redo_log_storage_size_in_gbs` - The REDO Log storage size, in gigabytes, that is applicable for the database. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the database was created.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database
	* `update` - (Defaults to 20 minutes), when updating the Database
	* `delete` - (Defaults to 20 minutes), when destroying the Database


## Import

Databases can be imported using the `id`, e.g.

```
$ terraform import oci_database_database.test_database "id"
```

