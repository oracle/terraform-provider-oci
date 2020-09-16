---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run"
sidebar_current: "docs-oci-resource-database-maintenance_run"
description: |-
  Provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service
---

# oci_database_maintenance_run
This resource provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service.

Updates the properties of a maintenance run, such as the state of a maintenance run.

## Example Usage

```hcl
resource "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	maintenance_run_id = oci_database_maintenance_run.test_maintenance_run.id

	#Optional
	is_enabled = var.maintenance_run_is_enabled
	is_patch_now_enabled = var.maintenance_run_is_patch_now_enabled
	patch_id = oci_database_patch.test_patch.id
	time_scheduled = var.maintenance_run_time_scheduled
}
```

## Argument Reference

The following arguments are supported:

* `is_enabled` - (Optional) (Updatable) If `FALSE`, skips the maintenance run.
* `is_patch_now_enabled` - (Optional) (Updatable) If set to `TRUE`, starts patching immediately.
* `maintenance_run_id` - (Required) The maintenance run OCID.
* `patch_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
* `time_scheduled` - (Optional) (Updatable) The scheduled date and time of the maintenance run to update.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - Description of the maintenance run.
* `display_name` - The user-friendly name for the maintenance run.
* `id` - The OCID of the maintenance run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
* `state` - The current state of the maintenance run.
* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
* `time_ended` - The date and time the maintenance run was completed.
* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
* `time_started` - The date and time the maintenance run starts.

## Import

MaintenanceRuns can be imported using the `id`, e.g.

```
$ terraform import oci_database_maintenance_run.test_maintenance_run "id"
```

