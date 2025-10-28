---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database"
sidebar_current: "docs-oci-resource-database-autonomous_container_database"
description: |-
  Provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_container_database
This resource provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/AutonomousContainerDatabase

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

Creates an Autonomous Container Database in the specified Autonomous Exadata Infrastructure.


## Example Usage

```hcl
resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
	#Required
	display_name = var.autonomous_container_database_display_name
	patch_model = var.autonomous_container_database_patch_model

	#Optional
	autonomous_container_database_backup_id = oci_database_autonomous_container_database_backup.test_autonomous_container_database_backup.id
	autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			backup_retention_policy_on_terminate = var.autonomous_container_database_backup_config_backup_destination_details_backup_retention_policy_on_terminate
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			is_remote = var.autonomous_container_database_backup_config_backup_destination_details_is_remote
			is_retention_lock_enabled = var.autonomous_container_database_backup_config_backup_destination_details_is_retention_lock_enabled
			remote_region = var.autonomous_container_database_backup_config_backup_destination_details_remote_region
			vpc_password = var.autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
	}
	cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	compartment_id = var.compartment_id
	customer_contacts {

		#Optional
		email = var.autonomous_container_database_customer_contacts_email
	}
	database_software_image_id = oci_database_database_software_image.test_database_software_image.id
	db_name = var.autonomous_container_database_db_name
	db_split_threshold = var.autonomous_container_database_db_split_threshold
	db_unique_name = var.autonomous_container_database_db_unique_name
	db_version = var.autonomous_container_database_db_version
	defined_tags = {"Operations.CostCenter"= "42"}
	distribution_affinity = var.autonomous_container_database_distribution_affinity
	fast_start_fail_over_lag_limit_in_seconds = var.autonomous_container_database_fast_start_fail_over_lag_limit_in_seconds
	freeform_tags = {"Department"= "Finance"}
	is_automatic_failover_enabled = var.autonomous_container_database_is_automatic_failover_enabled
	is_dst_file_update_enabled = var.autonomous_container_database_is_dst_file_update_enabled
	key_store_id = oci_database_key_store.test_key_store.id
	kms_key_id = oci_kms_key.test_key.id
	kms_key_version_id = oci_kms_key_version.test_key_version.id
	maintenance_window_details {

		#Optional
		custom_action_timeout_in_mins = var.autonomous_container_database_maintenance_window_details_custom_action_timeout_in_mins
		days_of_week {
			#Required
			name = var.autonomous_container_database_maintenance_window_details_days_of_week_name
		}
		hours_of_day = var.autonomous_container_database_maintenance_window_details_hours_of_day
		is_custom_action_timeout_enabled = var.autonomous_container_database_maintenance_window_details_is_custom_action_timeout_enabled
		is_monthly_patching_enabled = var.autonomous_container_database_maintenance_window_details_is_monthly_patching_enabled
		lead_time_in_weeks = var.autonomous_container_database_maintenance_window_details_lead_time_in_weeks
		months {
			#Required
			name = var.autonomous_container_database_maintenance_window_details_months_name
		}
		patching_mode = var.autonomous_container_database_maintenance_window_details_patching_mode
		preference = var.autonomous_container_database_maintenance_window_details_preference
		skip_ru = var.autonomous_container_database_maintenance_window_details_skip_ru
		weeks_of_month = var.autonomous_container_database_maintenance_window_details_weeks_of_month
	}
	net_services_architecture = var.autonomous_container_database_net_services_architecture
	okv_end_point_group_name = oci_identity_group.test_group.name
	peer_autonomous_container_database_backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			backup_retention_policy_on_terminate = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_backup_retention_policy_on_terminate
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			is_remote = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_is_remote
			is_retention_lock_enabled = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_is_retention_lock_enabled
			remote_region = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_remote_region
			vpc_password = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_peer_autonomous_container_database_backup_config_recovery_window_in_days
	}
	peer_autonomous_container_database_compartment_id = oci_identity_compartment.test_compartment.id
	peer_autonomous_container_database_display_name = var.autonomous_container_database_peer_autonomous_container_database_display_name
	peer_autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
	peer_autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	peer_cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	peer_db_unique_name = var.autonomous_container_database_peer_db_unique_name
	protection_mode = var.autonomous_container_database_protection_mode
	service_level_agreement_type = var.autonomous_container_database_service_level_agreement_type
	source = var.autonomous_container_database_source
	standby_maintenance_buffer_in_days = var.autonomous_container_database_standby_maintenance_buffer_in_days
	vault_id = oci_kms_vault.test_vault.id
	version_preference = var.autonomous_container_database_version_preference
	vm_failover_reservation = var.autonomous_container_database_vm_failover_reservation
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_backup_id` - (Required when source=BACKUP_FROM_ID) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source ACD backup that you will clone to create a new ACD.
* `autonomous_exadata_infrastructure_id` - (Optional) **No longer used.** This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
* `autonomous_vm_cluster_id` - (Optional) The OCID of the Autonomous VM Cluster.
* `backup_config` - (Optional) (Updatable) Backup options for the Autonomous Container Database. 
	* `backup_destination_details` - (Optional) (Updatable) Backup destination details.
		* `backup_retention_policy_on_terminate` - (Optional) (Updatable) Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
		* `dbrs_policy_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) (Updatable) Proxy URL to connect to object store.
		* `is_remote` - (Optional) (Updatable) Indicates whether the backup destination is cross-region or local.
		* `is_retention_lock_enabled` - (Optional) (Updatable) Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
		* `remote_region` - (Optional) (Updatable) The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) (Updatable) Type of the database backup destination.
		* `vpc_password` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `cloud_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
* `customer_contacts` - (Optional) (Updatable) Customer Contacts. Setting this to an empty list removes all customer contacts. 
	* `email` - (Optional) (Updatable) The email address used by Oracle to send notifications regarding databases and infrastructure.
* `database_software_image_id` - (Optional) The Autonomous AI Database Software Image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_name` - (Optional) The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
* `db_split_threshold` - (Optional) The CPU value beyond which an Autonomous AI Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
* `db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - (Optional) The base version for the Autonomous Container Database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The display name for the Autonomous Container Database.
* `distribution_affinity` - (Optional) Determines whether an Autonomous AI Database must be opened across a minimum or maximum of nodes. By default, Minimum nodes is selected.
* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) (Updatable) The lag time for my preference based on data loss tolerance in seconds.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_automatic_failover_enabled` - (Optional) (Updatable) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
* `is_dst_file_update_enabled` - (Optional) (Updatable) Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
* `key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - (Optional) The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances. 
* `maintenance_window_details` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - (Optional) (Updatable) If true, enables the monthly patching option.
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.
	* `skip_ru` - (Optional) (Updatable) If true, skips the release update (RU) for the quarter. You cannot skip two consecutive quarters. An RU skip request will only be honoured if the current version of the Autonomous Container Database is supported for current quarter. 
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `net_services_architecture` - (Optional) Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
* `okv_end_point_group_name` - (Optional) (Updatable) The OKV End Point Group name for the Autonomous Container Database. 
* `patch_model` - (Required) (Updatable) Database Patch model preference.
* `peer_autonomous_container_database_backup_config` - (Optional) Backup options for the standby Autonomous Container Database. 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `backup_retention_policy_on_terminate` - (Optional) Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `is_remote` - (Optional) Indicates whether the backup destination is cross-region or local.
		* `is_retention_lock_enabled` - (Optional) Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
		* `remote_region` - (Optional) The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `peer_autonomous_container_database_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database will be created. 
* `peer_autonomous_container_database_display_name` - (Optional) The display name for the peer Autonomous Container Database.
* `peer_autonomous_exadata_infrastructure_id` - (Optional) *No longer used.* This parameter is no longer used for Autonomous AI Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail. 
* `peer_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard. 
* `peer_cloud_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
* `peer_db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `protection_mode` - (Optional) (Updatable) The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `service_level_agreement_type` - (Optional) The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
* `source` - (Optional) The source of the database. Use `NONE` to create a new Autonomous Container Database (ACD). Use `BACKUP_FROM_ID` to create a new ACD from a specified backup. 
* `standby_maintenance_buffer_in_days` - (Optional) (Updatable) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `version_preference` - (Optional) (Updatable) The next maintenance version preference. 
* `vm_failover_reservation` - (Optional) The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, and 50%, with 50% being the default option.
* `failover_trigger` - (Optional) (Updatable) An optional property when incremented triggers Failover. Could be set to any integer value.
* `reinstate_trigger` - (Optional) (Updatable) An optional property when incremented triggers Reinstate. Could be set to any integer value.
* `switchover_trigger` - (Optional) (Updatable) An optional property when incremented triggers Switchover. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `associated_backup_configuration_details` - A backup config object holds information about preferred backup destinations only. This object holds information about the associated backup destinations, such as secondary backup destinations created for local backups or remote replicated backups.
	* `backup_destination_attach_history` - The timestamps at which this backup destination is used as the preferred destination to host the Autonomous Container Database backups.
	* `backup_retention_policy_on_terminate` - Defines the automatic and manual backup retention policy for the Autonomous Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
	* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	* `internet_proxy` - Proxy URL to connect to object store.
	* `is_remote` - Indicates whether the backup destination is cross-region or local.
	* `is_retention_lock_enabled` - Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
	* `recovery_window_in_days` - Number of days between the current and earliest point of recoverability covered by automatic backups and manual backups, but not long term backups.
	* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
	* `space_utilized_in_gbs` - The total space utilized (in GBs) by this Autonomous Container Database on this backup destination, rounded to the nearest integer.
	* `time_at_which_storage_details_are_updated` - The latest timestamp when the backup destination details, such as 'spaceUtilized,' are updated.
	* `type` - Type of the database backup destination.
	* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
* `autonomous_exadata_infrastructure_id` - **No longer used.** For Autonomous AI Database on dedicated Exadata infrastructure, the container database is created within a specified `cloudAutonomousVmCluster`. 
* `autonomous_vm_cluster_id` - The OCID of the Autonomous VM Cluster.
* `availability_domain` - The availability domain of the Autonomous Container Database.
* `available_cpus` - Sum of CPUs available on the Autonomous VM Cluster + Sum of reclaimable CPUs available in the Autonomous Container Database. 
* `backup_config` - Backup options for the Autonomous Container Database. 
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
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. If the number of specified days is 0 then there will be no backups. 
* `backup_destination_properties_list` - This list describes the backup destination properties associated with the Autonomous Container Database (ACD) 's preferred backup destination. The object at a given index is associated with the destination present at the same index in the backup destination details list of the ACD Backup Configuration.
	* `backup_destination_attach_history` - The timestamps at which this backup destination is used as the preferred destination to host the Autonomous Container Database backups.
	* `space_utilized_in_gbs` - The total space utilized (in GBs) by this Autonomous Container Database on this backup destination, rounded to the nearest integer.
	* `time_at_which_storage_details_are_updated` - The latest timestamp when the backup destination details, such as 'spaceUtilized,' are updated.
* `cloud_autonomous_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous Exadata VM Cluster.
* `compartment_id` - The OCID of the compartment.
* `compute_model` - The compute model of the Autonomous Container Database. For Autonomous AI Database on Dedicated Exadata Infrastructure, the CPU type (ECPUs or OCPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model. ECPU compute model is the recommended model and OCPU compute model is legacy. See [Compute Models in Autonomous AI Database on Dedicated Exadata Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details. 
* `customer_contacts` - Customer Contacts. Setting this to an empty list removes all customer contacts. 
	* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `dataguard` - The properties that define Autonomous Container Databases Dataguard. 
	* `apply_lag` - The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database. Example: `9 seconds` 
	* `apply_rate` - The rate at which redo logs are synchronized between the associated Autonomous Container Databases. Example: `180 Mb per second` 
	* `automatic_failover_target` - Automatically selected by backend when observer is enabled. 
	* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. Used only by Autonomous AI Database on Dedicated Exadata Infrastructure. 
	* `availability_domain` - The domain of the Autonomous Container Database 
	* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
	* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
	* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
	* `protection_mode` - The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
	* `redo_transport_mode` - Automatically selected by backend based on the protection mode. 
	* `role` - The Data Guard role of the Autonomous Container Database or Autonomous AI Database, if Autonomous Data Guard is enabled. 
	* `state` - The current state of Autonomous Data Guard.
	* `time_created` - The date and time the Autonomous DataGuard association was created.
	* `time_lag_refreshed_on` - Timestamp when the lags were last calculated for a standby.
	* `time_last_role_changed` - The date and time when the last role change action happened.
	* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
	* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database. Example: `7 seconds` 
* `dataguard_group_members` - Array of Dg associations.
	* `apply_lag` - The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database. Example: `9 seconds` 
	* `apply_rate` - The rate at which redo logs are synchronized between the associated Autonomous Container Databases. Example: `180 Mb per second` 
	* `automatic_failover_target` - Automatically selected by backend when observer is enabled. 
	* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. Used only by Autonomous AI Database on Dedicated Exadata Infrastructure. 
	* `availability_domain` - The domain of the Autonomous Container Database 
	* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
	* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
	* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
	* `protection_mode` - The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
	* `redo_transport_mode` - Automatically selected by backend based on the protection mode. 
	* `role` - The Data Guard role of the Autonomous Container Database or Autonomous AI Database, if Autonomous Data Guard is enabled. 
	* `state` - The current state of Autonomous Data Guard.
	* `time_created` - The date and time the Autonomous DataGuard association was created.
	* `time_lag_refreshed_on` - Timestamp when the lags were last calculated for a standby.
	* `time_last_role_changed` - The date and time when the last role change action happened.
	* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
	* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database. Example: `7 seconds` 
* `db_name` - The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
* `db_split_threshold` - The CPU value beyond which an Autonomous AI Database will be opened across multiple nodes. The default value of this attribute is 16 for OCPUs and 64 for ECPUs.
* `db_unique_name` - **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - Oracle AI Database version of the Autonomous Container Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `distribution_affinity` - Determines whether an Autonomous AI Database must be opened across the maximum number of nodes or the least number of nodes. By default, Minimum nodes is selected.
* `dst_file_version` - DST Time-Zone File version of the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_data_guard_enabled` - **Deprecated.** Indicates whether the Autonomous AI Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to Autonomous AI Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dst_file_update_enabled` - Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
* `is_multiple_standby` - Indicates if it is multiple standby Autonomous Dataguard 
* `key_history_entry` - Key History Entry.
	* `id` - The id of the Autonomous AI Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
    * `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances.
    * `time_activated` - The date and time the kms key activated.
    * `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous AI Database Serverless does not use key versions, hence is not applicable for Autonomous AI Database Serverless instances. 
* `largest_provisionable_autonomous_database_in_cpus` - The largest Autonomous AI Database (CPU) that can be created in a new Autonomous Container Database.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `list_one_off_patches` - List of One-Off patches that has been successfully applied to Autonomous Container Database
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `skip_ru` - If true, skips the release update (RU) for the quarter. You cannot skip two consecutive quarters. An RU skip request will only be honoured if the current version of the Autonomous Container Database is supported for current quarter. 
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_per_compute_unit_in_gbs` - The amount of memory (in GBs) to be enabled per OCPU or ECPU. 
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs, rounded off to nearest integer value) enabled per ECPU or OCPU in the Autonomous VM Cluster. This is deprecated. Please refer to memoryPerComputeUnitInGBs for accurate value.
* `net_services_architecture` - Enabling SHARED server architecture enables a database server to allow many client processes to share very few server processes, thereby increasing the number of supported users.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `okv_end_point_group_name` - The OKV End Point Group name for the Autonomous Container Database. 
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
* `patch_model` - Database patch model preference.
* `provisionable_cpus` - An array of CPU values that can be used to successfully provision a single Autonomous AI Database. 
* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous Container Database.
* `reclaimable_cpus` - CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous AI Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database. 
* `recovery_appliance_details` - Information about the recovery appliance configuration associated with the Autonomous Container Database.
	* `allocated_storage_size_in_gbs` - The storage size of the backup destination allocated for an Autonomous Container Database to store backups on the recovery appliance, in GBs, rounded to the nearest integer.
	* `recovery_window_in_days` - Number of days between the current and earliest point of recoverability covered by automatic backups.
	* `time_recovery_appliance_details_updated` - The time when the recovery appliance details are updated.
* `reserved_cpus` - The number of CPUs reserved in an Autonomous Container Database.
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous AI Database, if Autonomous Data Guard is enabled.
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `state` - The current state of the Autonomous Container Database.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Autonomous Container Database was created.
* `time_of_last_backup` - The timestamp of last successful backup. Here NULL value represents either there are no successful backups or backups are not configured for this Autonomous Container Database.
* `time_snapshot_standby_revert` - The date and time the Autonomous Container Database will be reverted to Standby from Snapshot Standby.
* `total_cpus` - The number of CPUs allocated to the Autonomous VM cluster. 
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `version_preference` - The next maintenance version preference. 
* `vm_failover_reservation` - The percentage of CPUs reserved across nodes to support node failover. Allowed values are 0%, 25%, and 50%, with 50% being the default option.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Autonomous Container Database
	* `update` - (Defaults to 12 hours), when updating the Autonomous Container Database
	* `delete` - (Defaults to 12 hours), when destroying the Autonomous Container Database


## Import

AutonomousContainerDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_container_database.test_autonomous_container_database "id"
```

