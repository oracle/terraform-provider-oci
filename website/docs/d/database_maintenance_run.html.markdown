---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run"
sidebar_current: "docs-oci-datasource-database-maintenance_run"
description: |-
  Provides details about a specific Maintenance Run in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_maintenance_run
This data source provides details about a specific Maintenance Run resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Maintenance Run.

## Example Usage

```hcl
data "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	maintenance_run_id = "${oci_database_maintenance_run.test_maintenance_run.id}"
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_run_id` - (Required) The Maintenance Run OCID.


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

