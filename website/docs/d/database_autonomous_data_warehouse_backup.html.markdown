---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_data_warehouse_backup"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse_backup"
description: |-
  Provides details about a specific Autonomous Data Warehouse Backup in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_data_warehouse_backup
This data source provides details about a specific Autonomous Data Warehouse Backup resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous Data Warehouse backup.

**IMPORTANT:** This data source is being **deprecated**, use `oci_database_autonomous_database_backup` instead.

## Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_backup_id = "${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_backup_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.


## Attributes Reference

The following attributes are exported:

* `autonomous_data_warehouse_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

