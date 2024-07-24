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

**Important:** When `auto_backup_enabled` is not present in the configuration or set to true, the `auto_backup_window` and `auto_full_backup_window` will be ignored

## Example Usage

```hcl
resource "oci_database_database" "test_database" {
	#Required
	database {
		#Required
		admin_password = var.database_database_admin_password
		db_name = var.database_database_db_name

		#Optional
		backup_id = oci_database_backup.test_backup.id
		backup_tde_password = var.database_database_backup_tde_password
		character_set = var.database_database_character_set
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
				dbrs_policy_id = oci_identity_policy.test_policy.id
				id = var.database_database_db_backup_config_backup_destination_details_id
				type = var.database_database_db_backup_config_backup_destination_details_type
			}
			recovery_window_in_days = var.database_database_db_backup_config_recovery_window_in_days
			run_immediate_full_backup = var.database_database_db_backup_config_run_immediate_full_backup
		}
		db_unique_name = var.database_database_db_unique_name
		db_workload = var.database_database_db_workload
		defined_tags = var.database_database_defined_tags
		freeform_tags = var.database_database_freeform_tags
		key_store_id = oci_database_key_store.test_key_store.id
		kms_key_id = oci_kms_key.test_key.id
		kms_key_version_id = oci_kms_key_version.test_key_version.id
		ncharacter_set = var.database_database_ncharacter_set
		pdb_name = var.database_database_pdb_name
		pluggable_databases = var.database_database_pluggable_databases
		sid_prefix = var.database_database_sid_prefix
		tde_wallet_password = var.database_database_tde_wallet_password
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
	* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	* `backup_id` - (Required when source=DB_BACKUP) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `backup_tde_password` - (Applicable when source=DB_BACKUP) The password to open the TDE wallet.
	* `character_set` - (Applicable when source=NONE) The character set for the database.  The default is AL32UTF8. Allowed values are:

		AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
	* `database_software_image_id` - (Applicable when source=NONE) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	* `db_backup_config` - (Applicable when source=NONE) (Updatable) Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
		* `auto_backup_enabled` - (Applicable when source=NONE) (Updatable) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
		* `auto_backup_window` - (Applicable when source=NONE) (Updatable) Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
		* `auto_full_backup_day` - (Applicable when source=NONE) Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
		* `auto_full_backup_window` - (Applicable when source=NONE) Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
		* `backup_deletion_policy` - (Applicable when source=NONE) This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
		* `backup_destination_details` - (Applicable when source=NONE) Backup destination details.
			* `dbrs_policy_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
			* `id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
			* `type` - (Required when source=NONE) Type of the database backup destination.
		* `recovery_window_in_days` - (Applicable when source=NONE) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
		* `run_immediate_full_backup` - (Applicable when source=NONE) If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
	* `db_name` - (Required) The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	* `db_unique_name` - (Optional) The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	* `db_workload` - (Applicable when source=NONE) **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service. Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.

		The database workload type. 
	* `defined_tags` - (Applicable when source=NONE) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
	* `freeform_tags` - (Applicable when source=NONE) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `key_store_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	* `kms_key_id` - (Applicable when source=NONE) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	* `kms_key_version_id` - (Applicable when source=NONE) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
	* `ncharacter_set` - (Applicable when source=NONE) The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
	* `pdb_name` - (Applicable when source=NONE) The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	* `pluggable_databases` - (Applicable when source=DB_BACKUP) The list of pluggable databases that needs to be restored into new database.
	* `sid_prefix` - (Optional) Specifies a prefix for the `Oracle SID` of the database to be created. 
	* `tde_wallet_password` - (Applicable when source=NONE) The optional password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
	* `vault_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `db_home_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `db_version` - (Optional) A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.

	This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_migration` - (Optional) The value to migrate to the kms version from none. Can only use once by setting value to true. You can not switch back to non-kms once you created or migrated.(https://www.oracle.com/security/cloud-security/key-management/faq/)
* `kms_key_rotation` - (Optional) The value to rotate the key version of current kms_key. Just change this value will trigger the rotation.
* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `source` - (Required) The source of the database: Use `NONE` for creating a new database. Use `DB_BACKUP` for creating a new database by restoring from a backup. The default is `NONE`. 


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
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `is_cdb` - True if the database is a container database.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
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
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
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

