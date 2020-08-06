---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_maintenance_run"
sidebar_current: "docs-oci-datasource-database-maintenance_run"
description: |-
  Provides details about a specific Maintenance Run in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_maintenance_run
This data source provides details about a specific Maintenance Run resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified maintenance run.

## Example Usage

```hcl
data "oci_database_maintenance_run" "test_maintenance_run" {
	#Required
	maintenance_run_id = "${oci_database_maintenance_run.test_maintenance_run.id}"
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_run_id` - (Required) The maintenance run OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `description` - Description of the maintenance run.
* `display_name` - The user-friendly name for the maintenance run.
* `id` - The OCID of the maintenance run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance_subtype` - Maintenance sub-type.
* `maintenance_type` - Maintenance type.
* `state` - The current state of the maintenance run.
* `target_resource_id` - The ID of the target resource on which the maintenance run occurs.
* `target_resource_type` - The type of the target resource on which the maintenance run occurs.
* `time_ended` - The date and time the maintenance run was completed.
* `time_scheduled` - The date and time the maintenance run is scheduled to occur.
* `time_started` - The date and time the maintenance run starts.

