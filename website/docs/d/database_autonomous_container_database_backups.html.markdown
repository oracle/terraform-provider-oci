---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_backups"
sidebar_current: "docs-oci-datasource-database-autonomous_container_database_backups"
description: |-
  Provides the list of Autonomous Container Database Backups in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_database_backups
This data source provides the list of Autonomous Container Database Backups in Oracle Cloud Infrastructure Database service.

Gets a list of Autonomous Container Database backups by using either the 'autonomousDatabaseId' or 'compartmentId' as your query parameter.


## Example Usage

```hcl
data "oci_database_autonomous_container_database_backups" "test_autonomous_container_database_backups" {

	#Optional
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	compartment_id = var.compartment_id
	display_name = var.autonomous_container_database_backup_display_name
	infrastructure_type = var.autonomous_container_database_backup_infrastructure_type
	is_remote = var.autonomous_container_database_backup_is_remote
	state = var.autonomous_container_database_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `infrastructure_type` - (Optional) A filter to return only resources that match the given Infrastructure Type.
* `is_remote` - (Optional) call for all remote backups
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_container_database_backup_collection` - The list of autonomous_container_database_backup_collection.

### AutonomousContainerDatabaseBackup Reference

The following attributes are exported:

* `acd_display_name` - The user-friendly name for the Autonomous Container Database when the Backup was initiated. This name need not be unique. This field captures the name at the time of backup creation, accounting for possible later updates to the display name.
* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
* `autonomous_databases` - List of Autonomous AI Databases that is part of this Autonomous Container Database Backup
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `display_name` - The user-friendly name for the Autonomous AI Database. The name does not have to be unique.
	* `state` - The current state of the Autonomous AI Database.
* `backup_destination_details` - Backup destination details
	* `backup_retention_policy_on_terminate` - Defines the automatic and manual backup retention policy for the Autonomous AI Database termination.  The retention policy set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. Options are 'RETAIN_PER_RETENTION_WINDOW' or 'RETAIN_FOR_72_HOURS'.The default value is 'RETAIN_FOR_72_HOURS'. 
	* `dbrs_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
	* `internet_proxy` - Proxy URL to connect to object store.
	* `is_remote` - Indicates whether the backup destination is cross-region or local.
	* `is_retention_lock_enabled` - Indicates if backup retention is locked for all the database backups in the Autonomous Container Database (ACD). The retention window cannot be decreased if the backup retention lock is enabled. Once applied on the Autonomous Container Database, the retention lock cannot be removed, or the retention period cannot be decreased after a 14-day period. If the backup is a Long Term Backup and retention lock is enabled, the backup cannot be deleted and must expire. The retention lock set on the Autonomous Container Database is not applicable for cross region remote backups and backups hosted on recovery Appliance backup destination. 
	* `is_zero_data_loss_enabled` - Indicates whether Zero Data Loss functionality is enabled for a Recovery Appliance backup destination in an Autonomous Container Database. When enabled, the database automatically ships all redo logs in real-time to the Recovery Appliance for a Zero Data Loss recovery setup (sub-second RPO). Defaults to `TRUE` if no value is given.
	* `remote_region` - The name of the remote region where the remote automatic incremental backups will be stored.           For information about valid region names, see [Regions and Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm). 
	* `type` - Type of the database backup destination.
	* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
	* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_version` - A valid Oracle AI Database version for Autonomous AI Database. When you specify 23ai for dbversion, the system will provision a 23ai database, but the UI will display it as 26ai. When you specify 26ai for dbversion, the system will provision and display a 26ai database as expected. For new databases, it is recommended to use either 19c or 26ai.

	**Note** Starting December 2026, 23ai will not be supported as a valid value for this parameter. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - A user-friendly name for the backup. This name need not be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database backup.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `is_remote_backup` - Whether backup is for remote-region or local region
* `lifecycle_details` - Additional information about the current lifecycle state.
* `retention_period_in_days` - Retention period, in days, for long-term backups
* `state` - The current state of the backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

