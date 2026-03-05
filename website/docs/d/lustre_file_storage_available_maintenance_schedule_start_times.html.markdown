---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_available_maintenance_schedule_start_times"
sidebar_current: "docs-oci-datasource-lustre_file_storage-available_maintenance_schedule_start_times"
description: |-
  Provides the list of Available Maintenance Schedule Start Times in Oracle Cloud Infrastructure Lustre File Storage service
---

# Data Source: oci_lustre_file_storage_available_maintenance_schedule_start_times
This data source provides the list of Available Maintenance Schedule Start Times in Oracle Cloud Infrastructure Lustre File Storage service.

Gets the list of available maintenance schedule start times for both Create and Update operation

## Example Usage

```hcl
data "oci_lustre_file_storage_available_maintenance_schedule_start_times" "test_available_maintenance_schedule_start_times" {

	#Optional
	availability_domain = var.available_maintenance_schedule_start_time_availability_domain
	compartment_id = var.compartment_id
	day_of_week = var.available_maintenance_schedule_start_time_day_of_week
	id = var.available_maintenance_schedule_start_time_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `day_of_week` - (Optional) Day of the week filter
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.


## Attributes Reference

The following attributes are exported:

* `available_maintenance_schedule_start_time_collection` - The list of available_maintenance_schedule_start_time_collection.

### AvailableMaintenanceScheduleStartTime Reference

The following attributes are exported:

* `items` - List of available start times on every day of the week
	* `day_of_week` - Day of the week
	* `start_times` - List of available start times. Each array item is of the format `HH:mm`

