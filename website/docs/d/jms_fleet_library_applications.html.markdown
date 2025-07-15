---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_library_applications"
sidebar_current: "docs-oci-datasource-jms-fleet_library_applications"
description: |-
  Provides the list of Fleet Library Applications in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_library_applications
This data source provides the list of Fleet Library Applications in Oracle Cloud Infrastructure Jms service.

List applications where a library has been detected filtered by query parameters.


## Example Usage

```hcl
data "oci_jms_fleet_library_applications" "test_fleet_library_applications" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	library_key = var.fleet_library_application_library_key

	#Optional
	application_id = oci_dataflow_application.test_application.id
	application_name = oci_dataflow_application.test_application.name
	application_name_contains = var.fleet_library_application_application_name_contains
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_end = var.fleet_library_application_time_end
	time_start = var.fleet_library_application_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the application.
* `application_name` - (Optional) The name of the application.
* `application_name_contains` - (Optional) Filter the list with application name contains the given value. 
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `library_key` - (Required) The unique identifier of a Java library.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the managed instance.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `library_application_usage_collection` - The list of library_application_usage_collection.

### FleetLibraryApplication Reference

The following attributes are exported:

* `items` - A list of LibraryApplicationUsageSummaries. 
	* `application_key` - The internal identifier of a Java application. 
	* `application_name` - The displayed name of the Java application. 
	* `first_seen_in_classpath` - The timestamp of the first time the specified library was detected in classpath. 
	* `last_detected_dynamically` - The date and time a library or Java package was _last_ detected in a dynamic library scan. 
	* `last_seen_in_classpath` - The timestamp of the last time the specified library was detected in classpath. 
	* `managed_instance_count` - The count of managed instances wherein the specified library was detected. 

