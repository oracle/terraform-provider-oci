---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system"
sidebar_current: "docs-oci-resource-database-db_system"
description: |-
  Provides the Db System resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_system
This resource provides the Db System resource in Oracle Cloud Infrastructure Database service.

Creates a new DB system in the specified compartment and availability domain. The Oracle
Database edition that you specify applies to all the databases on that DB system. The selected edition cannot be changed.

An initial database is created on the DB system based on the request parameters you provide and some default
options. For detailed information about default options, see [Bare metal and virtual machine DB system default options.](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/creatingDBsystem.htm#Default)

**Note:** Deprecated for Exadata Cloud Service systems. Use the [new resource model APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.

For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See [Switching an Exadata DB System to the New Resource Model and APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.

Use the [CreateCloudExadataInfrastructure](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudExadataInfrastructure/CreateCloudExadataInfrastructure/) and [CreateCloudVmCluster](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/CreateCloudVmCluster/) APIs to provision a new Exadata Cloud Service instance.

**Important:** When `auto_backup_enabled` is not present in the configuration or set to true, the `auto_backup_window` and `auto_full_backup_window` will be ignored

## Example Usage

```hcl
resource "oci_database_db_system" "test_db_system" {
	#Required
	availability_domain = var.db_system_availability_domain
	compartment_id = var.compartment_id
	db_home {
		#Required
		database {
			#Required
			admin_password = var.db_system_db_home_database_admin_password

			#Optional
			backup_id = oci_database_backup.test_backup.id
			backup_tde_password = var.db_system_db_home_database_backup_tde_password
			character_set = var.db_system_db_home_database_character_set
			database_id = oci_database_database.test_database.id
			database_software_image_id = oci_database_database_software_image.test_database_software_image.id
			db_backup_config {

				#Optional
				auto_backup_enabled = var.db_system_db_home_database_db_backup_config_auto_backup_enabled
				auto_backup_window = var.db_system_db_home_database_db_backup_config_auto_backup_window
				auto_full_backup_day = var.db_system_db_home_database_db_backup_config_auto_full_backup_day
				auto_full_backup_window = var.db_system_db_home_database_db_backup_config_auto_full_backup_window
				backup_deletion_policy = var.db_system_db_home_database_db_backup_config_backup_deletion_policy
				backup_destination_details {

					#Optional
					dbrs_policy_id = oci_identity_policy.test_policy.id
					id = var.db_system_db_home_database_db_backup_config_backup_destination_details_id
					type = var.db_system_db_home_database_db_backup_config_backup_destination_details_type
				}
				recovery_window_in_days = var.db_system_db_home_database_db_backup_config_recovery_window_in_days
				run_immediate_full_backup = var.db_system_db_home_database_db_backup_config_run_immediate_full_backup
			}
			db_domain = var.db_system_db_home_database_db_domain
			db_name = var.db_system_db_home_database_db_name
			db_workload = var.db_system_db_home_database_db_workload
			defined_tags = var.db_system_db_home_database_defined_tags
			freeform_tags = var.db_system_db_home_database_freeform_tags
			key_store_id = oci_database_key_store.test_key_store.id
			kms_key_id = oci_kms_key.test_key.id
			kms_key_version_id = oci_kms_key_version.test_key_version.id
			ncharacter_set = var.db_system_db_home_database_ncharacter_set
			pdb_name = var.db_system_db_home_database_pdb_name
			pluggable_databases = var.db_system_db_home_database_pluggable_databases
			sid_prefix = var.db_system_db_home_database_sid_prefix
			tde_wallet_password = var.db_system_db_home_database_tde_wallet_password
			time_stamp_for_point_in_time_recovery = var.db_system_db_home_database_time_stamp_for_point_in_time_recovery
			vault_id = oci_kms_vault.test_vault.id
		}

		#Optional
		database_software_image_id = oci_database_database_software_image.test_database_software_image.id
		db_unique_name = var.db_unique_name
		db_version = var.db_system_db_home_db_version
		defined_tags = var.db_system_db_home_defined_tags
		display_name = var.db_system_db_home_display_name
		freeform_tags = var.db_system_db_home_freeform_tags
	}
	hostname = var.db_system_hostname
	shape = var.db_system_shape
	ssh_public_keys = var.db_system_ssh_public_keys
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	backup_network_nsg_ids = var.db_system_backup_network_nsg_ids
	backup_subnet_id = oci_core_subnet.test_subnet.id
	cluster_name = var.db_system_cluster_name
	cpu_core_count = var.db_system_cpu_core_count
	data_collection_options {

		#Optional
		is_diagnostics_events_enabled = var.db_system_data_collection_options_is_diagnostics_events_enabled
		is_health_monitoring_enabled = var.db_system_data_collection_options_is_health_monitoring_enabled
		is_incident_logs_enabled = var.db_system_data_collection_options_is_incident_logs_enabled
	}
	data_storage_percentage = var.db_system_data_storage_percentage
	data_storage_size_in_gb = var.db_system_data_storage_size_in_gb
	database_edition = var.db_system_database_edition
	db_system_options {

		#Optional
		storage_management = var.db_system_db_system_options_storage_management
	}
	defined_tags = var.db_system_defined_tags
	disk_redundancy = var.db_system_disk_redundancy
	display_name = var.db_system_display_name
	domain = var.db_system_domain
	fault_domains = var.db_system_fault_domains
	freeform_tags = {"Department"= "Finance"}
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	license_model = var.db_system_license_model
	maintenance_window_details {

		#Optional
		custom_action_timeout_in_mins = var.db_system_maintenance_window_details_custom_action_timeout_in_mins
		days_of_week {

			#Optional
			name = var.db_system_maintenance_window_details_days_of_week_name
		}
		hours_of_day = var.db_system_maintenance_window_details_hours_of_day
		is_custom_action_timeout_enabled = var.db_system_maintenance_window_details_is_custom_action_timeout_enabled
		is_monthly_patching_enabled = var.db_system_maintenance_window_details_is_monthly_patching_enabled
		lead_time_in_weeks = var.db_system_maintenance_window_details_lead_time_in_weeks
		months {

			#Optional
			name = var.db_system_maintenance_window_details_months_name
		}
		patching_mode = var.db_system_maintenance_window_details_patching_mode
		preference = var.db_system_maintenance_window_details_preference
		weeks_of_month = var.db_system_maintenance_window_details_weeks_of_month
	}
	node_count = var.db_system_node_count
	nsg_ids = var.db_system_nsg_ids
	private_ip = var.db_system_private_ip
	source = var.db_system_source
	source_db_system_id = oci_database_db_system.test_db_system.id
	sparse_diskgroup = var.db_system_sparse_diskgroup
	storage_volume_performance_mode = var.db_system_storage_volume_performance_mode
	time_zone = var.db_system_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where the DB system is located.
* `backup_network_nsg_ids` - (Optional) (Updatable) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.

	**Subnet Restrictions:** See the subnet restrictions information for **subnetId**. 
* `cluster_name` - (Optional) The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DB system  belongs in.
* `cpu_core_count` - (Optional) (Updatable) The number of CPU cores to enable for a bare metal or Exadata DB system or AMD VMDB Systems. The valid values depend on the specified shape:
	* BM.DenseIO1.36 - Specify a multiple of 2, from 2 to 36.
	* BM.DenseIO2.52 - Specify a multiple of 2, from 2 to 52.
	* Exadata.Base.48 - Specify a multiple of 2, from 0 to 48.
	* Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	* Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	* Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	* Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	* Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	* Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.
	* VM.Standard.E4.Flex - Specify any thing from 1 to 64.

	This parameter is not used for INTEL virtual machine DB systems because virtual machine DB systems have a set number of cores for each shape. For information about the number of cores for a virtual machine DB system shape, see [Virtual Machine DB Systems](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/overview.htm#virtualmachine) 
* `data_collection_options` - (Optional) (Updatable) Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - (Optional) (Updatable) Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - (Optional) (Updatable) Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - (Optional) (Updatable) Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `data_storage_percentage` - (Optional) The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Specify 80 or 40. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems. Required for BMDBs.
* `data_storage_size_in_gb` - (Optional) (Updatable) Size (in GB) of the initial data volume that will be created and attached to a virtual machine DB system. You can scale up storage after provisioning, as needed. Note that the total storage size attached will be more than the amount you specify to allow for REDO/RECO space and software volume. Required for VMDBs.
* `database_edition` - (Required when source=DATABASE | DB_BACKUP | NONE) The Oracle Database Edition that applies to all the databases on the DB system. Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE. 
* `db_home` - (Required) (Updatable) Details for creating a Database Home if you are creating a database by restoring from a database backup.

	**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
	* `database` - (Required) (Updatable) Details for creating a database by restoring from a source database system.

		**Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API. 
		* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
		* `backup_id` - (Required when source=DB_BACKUP) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
		* `backup_tde_password` - (Applicable when source=DATABASE | DB_BACKUP) The password to open the TDE wallet.
		* `character_set` - (Applicable when source=NONE) The character set for the database.  The default is AL32UTF8. Allowed values are:

			AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
		* `database_id` - (Required when source=DATABASE) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
		* `database_software_image_id` - (Applicable when source=NONE) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
		* `db_backup_config` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
			* `auto_backup_enabled` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
			* `auto_backup_window` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
			* `auto_full_backup_day` - (Applicable when source=DB_SYSTEM | NONE) Day of the week the full backup should be applied on the database system. If no option is selected, the value is null and we will default to Sunday.
			* `auto_full_backup_window` - (Applicable when source=DB_SYSTEM | NONE) Time window selected for initiating full backup for the database system. There are twelve available two-hour time windows. If no option is selected, the value is null and a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
			* `backup_deletion_policy` - (Applicable when source=DB_SYSTEM | NONE) This defines when the backups will be deleted. - IMMEDIATE option keep the backup for predefined time i.e 72 hours and then delete permanently... - RETAIN will keep the backups as per the policy defined for database backups.
			* `backup_destination_details` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Backup destination details.
				* `dbrs_policy_id` - (Applicable when source=DB_SYSTEM | NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
				* `id` - (Applicable when source=DB_SYSTEM | NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
				* `type` - (Required when source=DB_SYSTEM | NONE) Type of the database backup destination.
			* `recovery_window_in_days` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
			* `run_immediate_full_backup` - (Applicable when source=DB_SYSTEM | NONE) If set to true, configures automatic full backups in the local region (the region of the DB system) for the first backup run immediately.
		* `db_domain` - (Applicable when source=DB_SYSTEM) The database domain. In a distributed database system, DB_DOMAIN specifies the logical location of the database within the network structure.
		* `db_name` - (Optional) The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
		* `db_workload` - (Applicable when source=NONE) **Deprecated.** The dbWorkload field has been deprecated for Exadata Database Service on Dedicated Infrastructure, Exadata Database Service on Cloud@Customer, and Base Database Service. Support for this attribute will end in November 2023. You may choose to update your custom scripts to exclude the dbWorkload attribute. After November 2023 if you pass a value to the dbWorkload attribute, it will be ignored.

			The database workload type. 
		* `defined_tags` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `freeform_tags` - (Applicable when source=DB_SYSTEM | NONE) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `key_store_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
		* `kms_key_id` - (Applicable when source=NONE) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
		* `kms_key_version_id` - (Applicable when source=NONE) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
		* `ncharacter_set` - (Applicable when source=NONE) The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
		* `pdb_name` - (Applicable when source=NONE) The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
		* `pluggable_databases` - (Applicable when source=DATABASE | DB_BACKUP) The list of pluggable databases that needs to be restored into new database.
		* `sid_prefix` - (Applicable when source=DB_BACKUP | NONE) Specifies a prefix for the `Oracle SID` of the database to be created. 
		* `tde_wallet_password` - (Applicable when source=NONE) The optional password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
		* `time_stamp_for_point_in_time_recovery` - (Applicable when source=DATABASE) The point in time of the original database from which the new database is created. If not specifed, the latest backup is used to create the database.
		* `vault_id` - (Applicable when source=NONE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
	* `database_software_image_id` - (Applicable when source=DB_BACKUP | NONE) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the image to be used to restore a database.
    * `db_unique_name` - (Optional) The `DB_UNIQUE_NAME` of the Oracle Database.
	* `db_version` - (Required when source=NONE) A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.

		This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier. 
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
	* `display_name` - (Optional) The user-provided name of the Database Home.
	* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `db_system_options` - (Optional) The DB system options.
	* `storage_management` - (Optional) The storage option used in DB system. ASM - Automatic storage management LVM - Logical Volume management 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `disk_redundancy` - (Applicable when source=DATABASE | DB_BACKUP | NONE) The type of redundancy configured for the DB system. Normal is 2-way redundancy, recommended for test and development systems. High is 3-way redundancy, recommended for production systems. 
* `display_name` - (Optional) The user-friendly name for the DB system. The name does not have to be unique.
* `domain` - (Optional) A domain name used for the DB system. If the Oracle-provided Internet and VCN Resolver is enabled for the specified subnet, the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. 
* `fault_domains` - (Optional) A Fault Domain is a grouping of hardware and infrastructure within an availability domain. Fault Domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or maintenance that affects one Fault Domain does not affect DB systems in other Fault Domains.

	If you do not specify the Fault Domain, the system selects one for you. To change the Fault Domain for a DB system, terminate it and launch a new DB system in the preferred Fault Domain.

	If the node count is greater than 1, you can specify which Fault Domains these nodes will be distributed into. The system assigns your nodes automatically to the Fault Domains you specify so that no Fault Domain contains more than one node.

	To get a list of Fault Domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - (Required) The hostname for the DB system. The hostname must begin with an alphabetic character, and can contain alphanumeric characters and hyphens (-). The maximum length of the hostname is 16 characters for bare metal and virtual machine DB systems, and 12 characters for Exadata DB systems.

	The maximum length of the combined hostname and domain is 63 characters.

	**Note:** The hostname must be unique within the subnet. If it is not unique, the DB system will fail to provision. 
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances. 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED. 
* `maintenance_window_details` - (Applicable when source=NONE) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - (Applicable when source=NONE) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - (Applicable when source=NONE) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required when source=NONE) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Applicable when source=NONE) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - (Applicable when source=NONE) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - (Applicable when source=NONE) (Updatable) If true, enables the monthly patching option.
	* `lead_time_in_weeks` - (Applicable when source=NONE) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Applicable when source=NONE) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required when source=NONE) (Updatable) Name of the month of the year.
	* `patching_mode` - (Applicable when source=NONE) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Applicable when source=NONE) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Applicable when source=NONE) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `node_count` - (Optional) The number of nodes to launch for a 2-node RAC virtual machine DB system. Specify either 1 or 2. 
* `nsg_ids` - (Optional) (Updatable) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `private_ip` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. Supported for VM BM shape.
* `shape` - (Required) (Updatable) The shape of the DB system. The shape determines resources allocated to the DB system.
	* For virtual machine shapes, the number of CPU cores and memory
	* For bare metal and Exadata shapes, the number of CPU cores, memory, and storage

	To get a list of shapes, use the [ListDbSystemShapes](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbSystemShapeSummary/ListDbSystemShapes) operation. 
* `source` - (Optional) The source of the database: Use `NONE` for creating a new database. Use `DB_BACKUP` for creating a new database by restoring from a backup. Use `DATABASE` for creating a new database from an existing database, including archive redo log data. The default is `NONE`. 
* `source_db_system_id` - (Required when source=DB_SYSTEM) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `sparse_diskgroup` - (Optional) If true, Sparse Diskgroup is configured for Exadata dbsystem. If False, Sparse diskgroup is not configured. Only applied for Exadata shape.
* `ssh_public_keys` - (Required) (Updatable) The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 40,000 characters.
* `storage_volume_performance_mode` - (Optional) The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See [Block Volume Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm) for more information. 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `time_zone` - (Optional) The time zone to use for the DB system. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the DB system is located in.
* `backup_network_nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.

	**Subnet Restriction:** See the subnet restrictions information for **subnetId**. 
* `cluster_name` - The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the DB system.
* `data_collection_options` - Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems. Required for BMDBs.
* `data_storage_size_in_gb` - The data storage size, in gigabytes, that is currently available to the DB system. Applies only for virtual machine DB systems. Required for VMDBs.
* `database_edition` - The Oracle Database edition that applies to all the databases on the DB system. 
* `db_system_options` - The DB system options.
	* `storage_management` - The storage option used in DB system. ASM - Automatic storage management LVM - Logical Volume management 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `disk_redundancy` - The type of redundancy configured for the DB system. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. 
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `domain` - The domain name for the DB system.
* `fault_domains` - List of the Fault Domains in which this DB system is provisioned.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for the DB system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `iorm_config_cache` - The IORM settings of the Exadata DB system. 
	* `db_plans` - An array of IORM settings for all the database in the Exadata DB system. 
		* `db_name` - The database name. For the default `DbPlan`, the `dbName` is `default`. 
		* `flash_cache_limit` - The flash cache limit for this database. This value is internally configured based on the share value assigned to the database. 
		* `share` - The relative priority of this database. 
	* `lifecycle_details` - Additional information about the current `lifecycleState`. 
	* `objective` - The current value for the IORM objective. The default is `AUTO`. 
	* `state` - The current state of IORM configuration for the Exadata DB system. 
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
* `license_model` - The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `listener_port` - The port number configured for the listener on the DB system.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_size_in_gbs` - Memory allocated to the DB system, in gigabytes.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `node_count` - The number of nodes in the DB system. For RAC DB systems, the value is greater than 1. 
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `os_version` - The most recent OS Patch Version applied on the DB system.
* `point_in_time_data_disk_clone_timestamp` - The point in time for a cloned database system when the data disks were cloned from the source database system, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `reco_storage_size_in_gb` - The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. Applies only for virtual machine DB systems. 
* `scan_dns_name` - The FQDN of the DNS record for the SCAN IP addresses that are associated with the DB system. 
* `scan_dns_record_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the DB system. 
* `scan_ip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the DB system. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `shape` - The shape of the DB system. The shape determines resources to allocate to the DB system.
	* For virtual machine shapes, the number of CPU cores and memory
	* For bare metal and Exadata shapes, the number of CPU cores, storage, and memory 
* `source_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `sparse_diskgroup` - True, if Sparse Diskgroup is configured for Exadata dbsystem, False, if Sparse diskgroup was not configured. Only applied for Exadata shape.
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the DB system.
* `state` - The current state of the DB system.
* `storage_volume_performance_mode` - The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See [Block Volume Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm) for more information. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the DB system was created.
* `time_zone` - The time zone of the DB system. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `version` - The Oracle Database version of the DB system.
* `vip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the DB system. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB system to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `zone_id` - The OCID of the zone the DB system is associated with. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 2 hours), when creating the Db System
	* `update` - (Defaults to 2 hours), when updating the Db System
	* `delete` - (Defaults to 2 hours), when destroying the Db System


## Import

DbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_system.test_db_system "id"
```

Import is only supported for source=NONE

`db_home.0.database.0.admin_password` is not returned by the service for security reasons. To avoid a force new of the db_home on the next apply, add the following to the resource:

```
    lifecycle {
        ignore_changes = ["db_home.0.database.0.admin_password"]
    }
```
You may also need to add `hostname` to the ignore_changes list if you see a diff on a subsequent apply

If the oci_database_db_system being imported is missing a primary db_home, an empty placeholder for `db_home` will be set in the Terraform state.
To keep configurations consistent with the imported state, add an empty placeholder for `db_home` to your configuration like this:

```
  # Add this placeholder into your oci_database_db_system configuration to indicate that the primary db home is empty.
  db_home {
    database {
      admin_password = ""
    }
  }
```
