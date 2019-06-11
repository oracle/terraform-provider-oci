---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run"
sidebar_current: "docs-oci-resource-database-maintenance_run"
description: |-
  Provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service
---

# oci_database_maintenance_run
This resource provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service.

Updates the properties of a Maintenance Run, such as the state of a Maintenance Run.

## Example Usage

```hcl
resource "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	maintenance_run_id = "${oci_database_maintenance_run.test_maintenance_run.id}"

	#Optional
	is_enabled = "${var.maintenance_run_is_enabled}"
}
```

## Argument Reference

The following arguments are supported:

* `is_enabled` - (Optional) (Updatable) If set to false, skips the Maintenance Run.
* `maintenance_run_id` - (Required) The Maintenance Run OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - The text describing this Maintenance Run.
* `display_name` - The user-friendly name for the Maintenance Run.
* `id` - The OCID of the Maintenance Run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `state` - The current state of the Maintenance Run.
* `target_resource_id` - The ID of the target resource on which the Maintenance Run occurs.
* `target_resource_type` - The type of the target resource on which the Maintenance Run occurs.
* `time_ended` - The date and time the Maintenance Run was completed.
* `time_scheduled` - The date and time the Maintenance Run is scheduled for.
* `time_started` - The date and time the Maintenance Run starts.

## Import

MaintenanceRuns can be imported using the `id`, e.g.

```
$ terraform import oci_database_maintenance_run.test_maintenance_run "id"
```

