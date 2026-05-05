---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups"
sidebar_current: "docs-oci-datasource-database-autonomous_container_database_backup_list_autonomous_databases_in_backups"
description: |-
  Provides the list of Autonomous Container Database Backup List Autonomous Databases In Backups in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups
This data source provides the list of Autonomous Container Database Backup List Autonomous Databases In Backups in Oracle Cloud Infrastructure Database service.

Gets a list of Autonomous Databases associated with backups at the given timestamp for the specified Autonomous Container Database. If `compartmentId` is provided, filters to that compartment; otherwise, uses the container's compartment.


## Example Usage

```hcl
data "oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups" "test_autonomous_container_database_backup_list_autonomous_databases_in_backups" {
	#Required
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	time_stamp_requested = var.autonomous_container_database_backup_list_autonomous_databases_in_backup_time_stamp_requested

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If not provided, uses the Autonomous Container Database's compartment.
* `time_stamp_requested` - (Required) The timestamp at which the list of autonomous databases present to be returned (ISO-8601 format, e.g., "2025-12-02T06:39:15Z"). The requested value must be after the end timestamp of the oldest available Autonomous Container Database backup.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_in_backup_collection` - The list of autonomous_database_in_backup_collection.

### AutonomousContainerDatabaseBackupListAutonomousDatabasesInBackup Reference

The following attributes are exported:

* `items` - The list of Autonomous Databases that are part of the Autonomous Container Database Backup.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `display_name` - The user-friendly name for the Autonomous AI Database. The name does not have to be unique.
	* `state` - The current state of the Autonomous AI Database.

