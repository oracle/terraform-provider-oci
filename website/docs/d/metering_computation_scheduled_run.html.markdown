---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_scheduled_run"
sidebar_current: "docs-oci-datasource-metering_computation-scheduled_run"
description: |-
  Provides details about a specific Scheduled Run in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_scheduled_run
This data source provides details about a specific Scheduled Run resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the saved schedule run.


## Example Usage

```hcl
data "oci_metering_computation_scheduled_run" "test_scheduled_run" {
	#Required
	scheduled_run_id = oci_metering_computation_scheduled_run.test_scheduled_run.id
}
```

## Argument Reference

The following arguments are supported:

* `scheduled_run_id` - (Required) The scheduledRun unique OCID


## Attributes Reference

The following attributes are exported:

* `id` - The ocid representing unique shedule run
* `lifecycle_details` - Additional details about scheduled run failure
* `schedule_id` - The ocid representing unique shedule
* `state` - Specifies if the schedule job was run successfully or not.
* `time_created` - The time when schedule started executing
* `time_finished` - The time when schedule finished executing

