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

* `items` - List of Autonomous container database backups.
	* `acd_display_name` - The user-friendly name for the Autonomous Container Database when the Backup was initiated. This name need not be unique. This field captures the name at the time of backup creation, accounting for possible later updates to the display name.
	* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
	* `autonomous_databases` - List of Autonomous AI Databases that is part of this Autonomous Container Database Backup
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `display_name` - The user-friendly name for the Autonomous AI Database. The name does not have to be unique.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
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

