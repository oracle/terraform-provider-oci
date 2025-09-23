---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_uncorrelated_package_applications"
sidebar_current: "docs-oci-datasource-jms-fleet_uncorrelated_package_applications"
description: |-
  Provides the list of Fleet Uncorrelated Package Applications in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_uncorrelated_package_applications
This data source provides the list of Fleet Uncorrelated Package Applications in Oracle Cloud Infrastructure Jms service.

List applications where an uncorrelated package has been detected, filtered by query parameters.


## Example Usage

```hcl
data "oci_jms_fleet_uncorrelated_package_applications" "test_fleet_uncorrelated_package_applications" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	package_name = var.fleet_uncorrelated_package_application_package_name

	#Optional
	application_id = oci_dataflow_application.test_application.id
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_end = var.fleet_uncorrelated_package_application_time_end
	time_start = var.fleet_uncorrelated_package_application_time_start
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the managed instance.
* `package_name` - (Required) The unique identifier of a Java package.
* `time_end` - (Optional) The end of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_start` - (Optional) The start of the time period during which resources are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `uncorrelated_package_application_usage_collection` - The list of uncorrelated_package_application_usage_collection.

### FleetUncorrelatedPackageApplication Reference

The following attributes are exported:

* `items` - A list of UncorrelatedPackageApplicationUsageSummaries.
	* `application_key` - The internal identifier of a Java application. 
	* `application_name` - The displayed name of the Java application. 
	* `last_detected_dynamically` - The date and time a library or Java package was _last_ detected in a dynamic library scan. 
	* `managed_instance_count` - The count of managed instances wherein the specified library was detected. 

