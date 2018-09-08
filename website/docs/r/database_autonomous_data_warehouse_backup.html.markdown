---
layout: "oci"
page_title: "OCI: oci_database_autonomous_data_warehouse_backup"
sidebar_current: "docs-oci-resource-database-autonomous_data_warehouse_backup"
description: |-
  Creates and manages an OCI AutonomousDataWarehouseBackup
---

# oci_database_autonomous_data_warehouse_backup
The `oci_database_autonomous_data_warehouse_backup` resource creates and manages an OCI AutonomousDataWarehouseBackup

Creates a new Autonomous Data Warehouse backup for the specified database based on the provided request parameters.


## Example Usage

```hcl
resource "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

AutonomousDataWarehouseBackups can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup "id"
```
