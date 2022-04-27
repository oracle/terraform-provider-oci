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

Returns schedule history list.


## Example Usage

```hcl
data "oci_metering_computation_scheduled_runs" "test_scheduled_runs" {
	#Required
	schedule_id = oci_metering_computation_schedule.test_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `schedule_id` - (Required) The unique ID of a schedule.


## Attributes Reference

The following attributes are exported:

* `scheduled_run_collection` - The list of scheduled_run_collection.

### ScheduledRun Reference

The following attributes are exported:

* `id` - The ocid representing unique shedule run
* `lifecycle_details` - Additional details about scheduled run failure
* `schedule_id` - The ocid representing unique shedule
* `state` - Specifies if the schedule job was run successfully or not.
* `time_created` - The time when schedule started executing
* `time_finished` - The time when schedule finished executing

