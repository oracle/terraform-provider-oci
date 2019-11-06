---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_data_warehouse_backup"
sidebar_current: "docs-oci-resource-database-autonomous_data_warehouse_backup"
description: |-
  Provides the Autonomous Data Warehouse Backup resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_data_warehouse_backup
This resource provides the Autonomous Data Warehouse Backup resource in Oracle Cloud Infrastructure Database service.

Creates a new Autonomous Data Warehouse backup for the specified database based on the provided request parameters.

**IMPORTANT:** This resource is being **deprecated**, use `oci_database_autonomous_database_backup` instead.
Refer to the [Deprecation Guide](#deprecation-guide) below on how to rename and migrate existing resources.

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

* `autonomous_data_warehouse_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

AutonomousDataWarehouseBackups can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup "id"
```

## Deprecation Guide

To rename existing `oci_database_autonomous_data_warehouse_backup` resource in your Terraform configuration and state to the new type `oci_database_autonomous_database_backup`, follow the steps below.

1. Using Terraform, move the existing resource in the state with the following command:

    ```
    $ terraform state mv oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup oci_database_autonomous_database_backup.test_autonomous_data_warehouse_backup
    ```
    *Note:* Terraform will automatically backup your state file, alternatively you may use `-backup=PATH` to override where the backup is written.
2. Update the name of resource `oci_database_autonomous_data_warehouse_backup` to the new name `oci_database_autonomous_database_backup` in your Terraform configuration. 
Do not make any more changes to Terraform configuration at this point other than the resource rename.
3. Run a Terraform `plan` to ensure that there are no issues post the state migration.
4. Use Terraform `refresh` or `apply` command to update the local state before making any further configuration changes or updates to the resource.
This step ensures that any fields that are marked `Computed` and/or `ForceNew` do not cause Terraform errors because of missing entries in local state when updating Terraform configuration.
