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

* `scheduled_run_id` - (Required) The scheduled run unique OCID.


## Attributes Reference

The following attributes are exported:

* `id` - The OCID representing a unique shedule run.
* `lifecycle_details` - Additional details about the scheduled run.
* `schedule_id` - The OCID representing a unique shedule.
* `state` - Specifies whether or not the schedule job was successfully run.
* `time_created` - The time the schedule started executing.
* `time_finished` - The time the schedule finished executing.

