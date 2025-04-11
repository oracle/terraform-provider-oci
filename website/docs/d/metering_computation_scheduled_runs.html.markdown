---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_scheduled_runs"
sidebar_current: "docs-oci-datasource-metering_computation-scheduled_runs"
description: |-
  Provides the list of Scheduled Runs in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_scheduled_runs
This data source provides the list of Scheduled Runs in Oracle Cloud Infrastructure Metering Computation service.

Returns the schedule history list.


## Example Usage

```hcl
data "oci_metering_computation_scheduled_runs" "test_scheduled_runs" {
	#Required
	schedule_id = oci_metering_computation_schedule.test_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `schedule_id` - (Required) The schedule unique ID.


## Attributes Reference

The following attributes are exported:

* `scheduled_run_collection` - The list of scheduled_run_collection.

### ScheduledRun Reference

The following attributes are exported:

* `id` - The OCID representing a unique shedule run.
* `lifecycle_details` - Additional details about the scheduled run.
* `schedule_id` - The OCID representing a unique shedule.
* `state` - Specifies whether or not the schedule job was successfully run.
* `time_created` - The time the schedule started executing.
* `time_finished` - The time the schedule finished executing.

