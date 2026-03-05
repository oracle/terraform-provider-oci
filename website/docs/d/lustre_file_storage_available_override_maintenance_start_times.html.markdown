---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_available_override_maintenance_start_times"
sidebar_current: "docs-oci-datasource-lustre_file_storage-available_override_maintenance_start_times"
description: |-
  Provides the list of Available Override Maintenance Start Times in Oracle Cloud Infrastructure Lustre File Storage service
---

# Data Source: oci_lustre_file_storage_available_override_maintenance_start_times
This data source provides the list of Available Override Maintenance Start Times in Oracle Cloud Infrastructure Lustre File Storage service.

Gets the list of available maintenance start times for Override operation

## Example Usage

```hcl
data "oci_lustre_file_storage_available_override_maintenance_start_times" "test_available_override_maintenance_start_times" {
	#Required
	id = var.available_override_maintenance_start_time_id

	#Optional
	date = var.available_override_maintenance_start_time_date
}
```

## Argument Reference

The following arguments are supported:

* `date` - (Optional) Date in format `YYYY-MM-DD`
* `id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.


## Attributes Reference

The following attributes are exported:

* `available_override_maintenance_start_time_collection` - The list of available_override_maintenance_start_time_collection.

### AvailableOverrideMaintenanceStartTime Reference

The following attributes are exported:

* `items` - List of available start times on every day of the week
	* `start_times` - List of available start times. Each array item is of the format `HH:mm`
	* `time_date_available` - The date corresponding to the list of start times available.  Example: `2024-04-25T00:00:00.000Z` 

