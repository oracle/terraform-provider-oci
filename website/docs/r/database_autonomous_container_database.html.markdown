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

Creates an Autonomous Container Database in the specified Autonomous Exadata Infrastructure.


## Example Usage

```hcl
resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
	#Required
	display_name = var.autonomous_container_database_display_name
	patch_model = var.autonomous_container_database_patch_model

	#Optional
	cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			vpc_password = var.autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
	}
	compartment_id = var.compartment_id
	db_name = var.autonomous_container_database_db_name
	db_unique_name = var.autonomous_container_database_db_unique_name
	db_version = var.autonomous_container_database_db_version
	defined_tags = {"Operations.CostCenter"= "42"}
	fast_start_fail_over_lag_limit_in_seconds = var.autonomous_container_database_fast_start_fail_over_lag_limit_in_seconds
	freeform_tags = {"Department"= "Finance"}
	is_automatic_failover_enabled = var.autonomous_container_database_is_automatic_failover_enabled
	is_dst_file_update_enabled = var.autonomous_container_database_is_dst_file_update_enabled
	key_store_id = oci_database_key_store.test_key_store.id
	kms_key_id = oci_kms_key.test_key.id
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
		weeks_of_month = var.autonomous_container_database_maintenance_window_details_weeks_of_month
	}
	peer_autonomous_container_database_display_name = var.autonomous_container_database_peer_autonomous_container_database_display_name
	peer_cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	protection_mode = var.autonomous_container_database_protection_mode
	peer_autonomous_container_database_backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			vpc_password = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_peer_autonomous_container_database_backup_config_recovery_window_in_days
	}
	peer_autonomous_container_database_compartment_id = oci_identity_compartment.test_compartment.id
	peer_autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	peer_db_unique_name = var.autonomous_container_database_peer_db_unique_name
	service_level_agreement_type = var.autonomous_container_database_service_level_agreement_type
	vault_id = oci_kms_vault.test_vault.id
	version_preference = var.autonomous_container_database_version_preference
	standby_maintenance_buffer_in_days = var.autonomous_container_database_standby_maintenance_buffer_in_days
}
```

## Argument Reference

The following arguments are supported:


* `autonomous_exadata_infrastructure_id` - (Optional) **No longer used.** This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `cloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail. 

* `autonomous_vm_cluster_id` - (Optional) The OCID of the Autonomous VM Cluster.
* `cloud_autonomous_vm_cluster_id` - (Optional) The OCID of the Cloud Autonomous VM Cluster.
* `backup_config` - (Optional) (Updatable) Backup options for the Autonomous Container Database. 
	* `backup_destination_details` - (Optional) (Updatable) Backup destination details.
		* `dbrs_policy_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) (Updatable) Proxy URL to connect to object store.
		* `type` - (Required) (Updatable) Type of the database backup destination.
		* `vpc_password` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
* `db_name` - (Optional) The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
* `db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - (Optional) The base version for the Autonomous Container Database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The display name for the Autonomous Container Database.
* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) The lag time for my preference based on data loss tolerance in seconds.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_dst_file_update_enabled` - (Optional) (Updatable) Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
* `is_automatic_failover_enabled` - (Optional) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association. Input DataType: boolean. Example : is_automatic_failover_enabled = true.
* `key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `maintenance_window_details` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - (Optional) (Updatable) Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - (Optional) (Updatable) If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - (Optional) (Updatable) If true, enables the monthly patching option.
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `patch_model` - (Required) (Updatable) Database Patch model preference.
* `peer_autonomous_container_database_display_name` - (Optional) The display name for the peer Autonomous Container Database.
* `peer_autonomous_exadata_infrastructure_id` - (End of Life) The OCID of the peer Autonomous Exadata Infrastructure for autonomous dataguard. Please use peer_cloud_autonomous_vm_cluster_id instead.
* `peer_cloud_autonomous_vm_cluster_id` - The OCID of the peer Autonomous Cloud VM Cluster for autonomous dataguard.  
* `protection_mode` - (Optional) The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `peer_autonomous_container_database_backup_config` - (Optional) 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `peer_autonomous_container_database_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database will be created.
* `peer_autonomous_container_database_display_name` - (Optional) The display name for the peer Autonomous Container Database.
* `peer_autonomous_exadata_infrastructure_id` - (Optional) *No longer used.* This parameter is no longer used for Autonomous Database on dedicated Exadata infrasture. Specify a `peerCloudAutonomousVmClusterId` instead. Using this parameter will cause the operation to fail.
* `peer_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard. 
* `peer_db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `protection_mode` - (Optional) The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `service_level_agreement_type` - (Optional) The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
* `standby_maintenance_buffer_in_days` - (Optional) (Updatable) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database.  
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `version_preference` - (Optional) (Updatable) The next maintenance version preference.
* `rotate_key_trigger` - (Optional) (Updatable) An optional property when flipped triggers rotation of KMS key. It is only applicable on dedicated container databases i.e. where `cloud_autonomous_vm_cluster_id` is set.
* `standby_maintenance_buffer_in_days` - (Optional) (Updatable) The scheduling detail for the quarterly maintenance window of standby Autonomous Container Database. This value represents the number of days before the primary database maintenance schedule.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_id` - **No longer used.** For Autonomous Database on dedicated Exadata infrastructure, the container database is created within a specified `cloudAutonomousVmCluster`. 
* `autonomous_vm_cluster_id` - The OCID of the Autonomous VM Cluster.
* `cloud_autonomous_vm_cluster_id` - The OCID of the Cloud Autonomous VM Cluster.
* `availability_domain` - The availability domain of the Autonomous Container Database
* `available_cpus` - Sum of CPUs available on the Autonomous VM Cluster + Sum of reclaimable CPUs available in the Autonomous Container Database.<br> For Autonomous Databases on Dedicated Exadata Infrastructure, the CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model. See [Compute Models in Autonomous Database on Dedicated Exadata Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
* `backup_config` - Backup options for the Autonomous Container Database. 
    * `backup_destination_details` - Backup destination details.
        * `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
        * `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
        * `internet_proxy` - Proxy URL to connect to object store.
        * `type` - Type of the database backup destination.
        * `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
        * `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
    * `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - The OCID of the compartment.
* `compute_model` - The compute model of the Autonomous VM Cluster.  
* `db_name` - The Database name for the Autonomous Container Database. The name must be unique within the Cloud Autonomous VM Cluster, starting with an alphabetic character, followed by 1 to 7 alphanumeric characters.
* `db_unique_name` - **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - Oracle Database version of the Autonomous Container Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `dst_file_version` - DST Time Zone File version of the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_dst_file_update_enabled` - Indicates if an automatic DST Time Zone file update is enabled for the Autonomous Container Database. If enabled along with Release Update, patching will be done in a Non-Rolling manner.
* `key_history_entry` - Key History Entry.
    * `id` - The id of the Autonomous Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
    * `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
    * `time_activated` - The date and time the kms key activated.
    * `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `largest_provisionable_autonomous_database_in_cpus` - The largest Autonomous Database (CPU) that can be created in a new Autonomous Container Database.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
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
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per OCPU or ECPU in the Autonomous VM Cluster. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
* `patch_model` - Database patch model preference.
* `provisionable_cpus` - An array of CPU values that can be used to successfully provision a single Autonomous Database.\ For Autonomous Database on Dedicated Exadata Infrastructure, the CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model. 
* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous Container Database.
* `reclaimable_cpus` - For Autonomous Databases on Dedicated Exadata Infrastructure:
    * These are the CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database.
    * The CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model.
* `reserved_cpus` - The number of CPUs reserved in an Autonomous Container Database.
    * These are the CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database. 
    * The CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model. See [Compute Models in Autonomous Database on Dedicated Exadata Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
* `resource_pool_leader_id` - (Optional) (Updatable) The unique identifier for leader autonomous database OCID [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `resource_pool_summary` - (Optional) (Updatable) The configuration details for resource pool
  * `is_disabled` - (Optional) (Updatable) Indicates if the resource pool should be deleted for the Autonomous Database.
  * `pool_size` - (Optional) (Updatable) Resource pool size.
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `state` - The current state of the Autonomous Container Database.
* `time_created` - The date and time the Autonomous Container Database was created.
* `time_snapshot_standby_revert` - The date and time the Autonomous Container Database will be reverted to Standby from Snapshot Standby.
* `total_cpus` - The number of CPUs allocated to the Autonomous VM cluster.<br> For Autonomous Databases on Dedicated Exadata Infrastructure, the CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model.  
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `version_preference` - The next maintenance version preference. 

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
