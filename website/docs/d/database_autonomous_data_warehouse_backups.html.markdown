---
layout: "oci"
page_title: "OCI: oci_database_autonomous_data_warehouse_backups"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse_backups"
description: |-
  Provides a list of AutonomousDataWarehouseBackups
---

# Data Source: oci_database_autonomous_data_warehouse_backups
The `oci_database_autonomous_data_warehouse_backups` data source allows access to the list of OCI autonomous_data_warehouse_backups

Gets a list of Autonomous Data Warehouse backups based on either the `autonomousDataWarehouseId` or `compartmentId` specified as a query parameter.


## Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_backups" "test_autonomous_data_warehouse_backups" {

	#Optional
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
	state = "${var.autonomous_data_warehouse_backup_state}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_id` - (Optional) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_data_warehouse_backups` - The list of autonomous_data_warehouse_backups.

### AutonomousDataWarehouseBackup Reference

The following attributes are exported:

* `autonomous_data_warehouse_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
* `compartment_id` - The OCID of the compartment.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

