---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_runs"
sidebar_current: "docs-oci-datasource-database-maintenance_runs"
description: |-
  Provides the list of Maintenance Runs in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_maintenance_runs
This data source provides the list of Maintenance Runs in Oracle Cloud Infrastructure Database service.

Gets a list of the Maintenance Runs in the specified compartment.


## Example Usage

```hcl
data "oci_database_maintenance_runs" "test_maintenance_runs" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.maintenance_run_availability_domain}"
	maintenance_type = "${var.maintenance_run_maintenance_type}"
	state = "${var.maintenance_run_state}"
	target_resource_id = "${oci_database_target_resource.test_target_resource.id}"
	target_resource_type = "${var.maintenance_run_target_resource_type}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `maintenance_type` - (Optional) The maintenance type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `target_resource_id` - (Optional) The target resource ID.
* `target_resource_type` - (Optional) The type of the target resource.


## Attributes Reference

The following attributes are exported:

* `maintenance_runs` - The list of maintenance_runs.

### MaintenanceRun Reference

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

