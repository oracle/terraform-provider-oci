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
	autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
	backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			id = var.autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			vpc_password = var.autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
	}
	compartment_id = var.compartment_id
	db_unique_name = var.autonomous_container_database_db_unique_name
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_automatic_failover_enabled = var.autonomous_container_database_is_automatic_failover_enabled
	key_store_id = oci_database_key_store.test_key_store.id
	kms_key_id = oci_kms_key.test_key.id
	maintenance_window_details {
		#Required
		preference = var.autonomous_container_database_maintenance_window_details_preference

		#Optional
		days_of_week {
			#Required
			name = var.autonomous_container_database_maintenance_window_details_days_of_week_name
		}
		hours_of_day = var.autonomous_container_database_maintenance_window_details_hours_of_day
		lead_time_in_weeks = var.autonomous_container_database_maintenance_window_details_lead_time_in_weeks
		months {
			#Required
			name = var.autonomous_container_database_maintenance_window_details_months_name
		}
		weeks_of_month = var.autonomous_container_database_maintenance_window_details_weeks_of_month
	}
	peer_autonomous_container_database_display_name = var.autonomous_container_database_peer_autonomous_container_database_display_name
	peer_autonomous_exadata_infrastructure_id = oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id
	protection_mode = var.autonomous_container_database_protection_mode
	peer_autonomous_container_database_backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_peer_autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
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
	standby_maintenance_buffer_in_days = var.autonomous_container_database_standby_maintenance_buffer_in_days
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Optional) The OCID of the Autonomous Exadata Infrastructure.
* `autonomous_vm_cluster_id` - (Optional) The OCID of the Autonomous VM Cluster.
* `backup_config` - (Optional) (Updatable) Backup options for the Autonomous Container Database. 
	* `backup_destination_details` - (Optional) (Updatable) Backup destination details.
		* `id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) (Updatable) Proxy URL to connect to object store.
		* `type` - (Required) (Updatable) Type of the database backup destination.
		* `vpc_password` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) (Updatable) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
* `db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The display name for the Autonomous Container Database.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_automatic_failover_enabled` - (Optional) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
* `key_store_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `maintenance_window_details` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `preference` - (Required) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `patch_model` - (Required) (Updatable) Database Patch model preference.
* `peer_autonomous_container_database_display_name` - (Optional) The display name for the peer Autonomous Container Database.
* `peer_autonomous_exadata_infrastructure_id` - (Optional) The OCID of the peer Autonomous Exadata Infrastructure for autonomous dataguard.
* `protection_mode` - (Optional) The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `peer_autonomous_container_database_backup_config` - (Optional) 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `peer_autonomous_container_database_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database will be created. 
* `peer_autonomous_vm_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous VM cluster for Autonomous Data Guard. Required to enable Data Guard. 
* `peer_db_unique_name` - (Optional) **Deprecated.** The `DB_UNIQUE_NAME` of the peer Autonomous Container Database in a Data Guard association is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `protection_mode` - (Optional) The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `service_level_agreement_type` - (Optional) The service level agreement type of the Autonomous Container Database. The default is STANDARD. For an autonomous dataguard Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.
* `standby_maintenance_buffer_in_days` - (Optional) (Updatable) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database.  
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `rotate_key_trigger` - (Optional) (Updatable) An optional property when flipped triggers rotation of KMS key. It is only applicable on dedicated container databases i.e. where `autonomous_exadata_infrastructure_id` is set.
* `standby_maintenance_buffer_in_days` - (Optional) (Updatable) The scheduling detail for the quarterly maintenance window of standby Autonomous Container Database. This value represents the number of days before the primary database maintenance schedule. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_id` - The OCID of the Autonomous Exadata Infrastructure.
* `autonomous_vm_cluster_id` - The OCID of the Autonomous VM Cluster.
* `availability_domain` - The availability domain of the Autonomous Container Database.
* `backup_config` - Backup options for the Autonomous Container Database. 
	* `backup_destination_details` - Backup destination details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - The OCID of the compartment.
* `db_unique_name` - **Deprecated.** The `DB_UNIQUE_NAME` value is set by Oracle Cloud Infrastructure.  Do not specify a value for this parameter. Specifying a value for this field will cause Terraform operations to fail. 
* `db_version` - Oracle Database version of the Autonomous Container Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `key_history_entry` - Key History Entry.
	* `id` - The id of the Autonomous Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
	* `time_activated` - The date and time the kms key activated.
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch applied on the system.
* `patch_model` - Database patch model preference.
* `role` - The role of the dataguard enabled Autonomous Container Database.
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `standby_maintenance_buffer_in_days` - The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `state` - The current state of the Autonomous Container Database.
* `time_created` - The date and time the Autonomous Container Database was created.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Autonomous Container Database
	* `update` - (Defaults to 12 hours), when updating the Autonomous Container Database
	* `delete` - (Defaults to 12 hours), when destroying the Autonomous Container Database


## Import

AutonomousContainerDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_container_database.test_autonomous_container_database "id"
```

