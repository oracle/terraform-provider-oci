---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_library_managed_instances"
sidebar_current: "docs-oci-datasource-jms-fleet_library_managed_instances"
description: |-
  Provides the list of Fleet Library Managed Instances in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_library_managed_instances
This data source provides the list of Fleet Library Managed Instances in Oracle Cloud Infrastructure Jms service.

List managed instances where a library has been detected, filtered by query parameters.


## Example Usage

```hcl
data "oci_jms_fleet_library_managed_instances" "test_fleet_library_managed_instances" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	library_key = var.fleet_library_managed_instance_library_key

	#Optional
	application_id = oci_dataflow_application.test_application.id
	host_name = var.fleet_library_managed_instance_host_name
	hostname_contains = var.fleet_library_managed_instance_hostname_contains
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_end = var.fleet_library_managed_instance_time_end
	time_start = var.fleet_library_managed_instance_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `host_name` - (Optional) The host [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `hostname_contains` - (Optional) Filter the list with hostname contains the given value. 
* `library_key` - (Required) The unique identifier of a Java library.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `library_managed_instance_usage_collection` - The list of library_managed_instance_usage_collection.

### FleetLibraryManagedInstance Reference

The following attributes are exported:

* `items` - A list of LibraryManagedInstanceUsageSummaries. 
	* `application_count` - The count of applications where the specified library was detected. 
	* `first_seen_in_classpath` - The timestamp of the first time the specified library was detected in classpath. 
	* `hostname` - The hostname of the managed instance.
	* `last_detected_dynamically` - The date and time a library or Java package was _last_ detected in a dynamic library scan. 
	* `last_seen_in_classpath` - The timestamp of the last time the specified library was detected in classpath. 
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 

