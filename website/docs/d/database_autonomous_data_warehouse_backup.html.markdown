---
layout: "oci"
page_title: "OCI: oci_database_autonomous_data_warehouse_backup"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse_backup"
description: |-
  Provides details about a specific AutonomousDataWarehouseBackup
---

# Data Source: oci_database_autonomous_data_warehouse_backup
The `oci_database_autonomous_data_warehouse_backup` data source provides details about a specific AutonomousDataWarehouseBackup

Gets information about the specified Autonomous Data Warehouse backup.

## Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_backup_id = "${var.autonomous_data_warehouse_backup_autonomous_data_warehouse_backup_id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_backup_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.


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

