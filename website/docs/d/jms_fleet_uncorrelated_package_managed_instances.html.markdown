---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_uncorrelated_package_managed_instances"
sidebar_current: "docs-oci-datasource-jms-fleet_uncorrelated_package_managed_instances"
description: |-
  Provides the list of Fleet Uncorrelated Package Managed Instances in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_uncorrelated_package_managed_instances
This data source provides the list of Fleet Uncorrelated Package Managed Instances in Oracle Cloud Infrastructure Jms service.

List managed instances where an uncorrelated package has been detected, filtered by query parameters.


## Example Usage

```hcl
data "oci_jms_fleet_uncorrelated_package_managed_instances" "test_fleet_uncorrelated_package_managed_instances" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id
	package_name = var.fleet_uncorrelated_package_managed_instance_package_name

	#Optional
	application_id = oci_dataflow_application.test_application.id
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	time_end = var.fleet_uncorrelated_package_managed_instance_time_end
	time_start = var.fleet_uncorrelated_package_managed_instance_time_start
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

* `uncorrelated_package_managed_instance_usage_collection` - The list of uncorrelated_package_managed_instance_usage_collection.

### FleetUncorrelatedPackageManagedInstance Reference

The following attributes are exported:

* `items` - A list of UncorrelatedPackageManagedInstanceUsageSummaries.
	* `application_count` - The count of applications wherein the specified library was detected. 
	* `hostname` - The hostname of the managed instance.
	* `last_detected_dynamically` - The date and time a library or Java package was _last_ detected in a dynamic library scan. 
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 

