---
subcategory: "Announcements Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_announcements_service_services"
sidebar_current: "docs-oci-datasource-announcements_service-services"
description: |-
  Provides the list of Services in Oracle Cloud Infrastructure
---

# Data Source: oci_announcements_service_services
This data source provides the list of Services in Oracle Cloud Infrastructure.

List all OCI services


## Example Usage

```hcl
data "oci_announcements_service_services" "test_services" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	comms_manager_name = var.service_comms_manager_name
	platform_type = var.service_platform_type
}
```

## Argument Reference

The following arguments are supported:

* `comms_manager_name` - (Optional) Filter by comms manager name
* `compartment_id` - (Required) The OCID of the root compartment/tenancy. 
* `platform_type` - (Optional) A filter to return only services underlying a specific platform.


## Attributes Reference

The following attributes are exported:

* `services_collection` - The list of services_collection.

### Service Reference

The following attributes are exported:

	* `comms_manager_name` - Name of the comms manager team that manages Notifications to this service.
	* `excluded_realms` - The list of realms where this service is not available to be used.
	* `id` - ID of the service object.
	* `platform_type` - The platform type this service object is related to.
	* `previous_service_names` - The list of previously used names for this service object.
	* `service_name` - Name of the service represented by this object.
	* `short_name` - Short name of the team to whom this service object is related.
	* `state` - Current state of the service object.
	* `team_name` - Team name to which this service object is related.
	* `time_created` - The date and time when the service object was created.
	* `time_updated` - The date and time when the service object was updated.

