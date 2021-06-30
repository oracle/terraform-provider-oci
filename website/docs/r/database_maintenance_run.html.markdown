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
	patching_mode = var.maintenance_run_patching_mode
	time_scheduled = var.maintenance_run_time_scheduled
}
```

## Argument Reference

The following arguments are supported:

* `is_enabled` - (Optional) (Updatable) If `FALSE`, skips the maintenance run.
* `is_patch_now_enabled` - (Optional) (Updatable) If set to `TRUE`, starts patching immediately.
* `maintenance_run_id` - (Required) The maintenance run OCID.
* `patch_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

	*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
* `time_scheduled` - (Optional) (Updatable) The scheduled date and time of the maintenance run to update.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - Description of the maintenance run.
* `display_name` - The user-friendly name for the maintenance run.
* `id` - The OCID of the maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `patch_failure_count` - Contain the patch failure count.
* `patch_id` - The unique identifier of the patch. The identifier string includes the patch type, the Oracle Database version, and the patch creation date (using the format YYMMDD). For example, the identifier `ru_patch_19.9.0.0_201030` is used for an RU patch for Oracle Database 19.9.0.0 that was released October 30, 2020.
* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

	*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
* `peer_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
* `state` - The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED. 
* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
* `time_ended` - The date and time the maintenance run was completed.
* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
* `time_started` - The date and time the maintenance run starts.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Run
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Run
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Run


## Import

MaintenanceRuns can be imported using the `id`, e.g.

```
$ terraform import oci_database_maintenance_run.test_maintenance_run "id"
```

