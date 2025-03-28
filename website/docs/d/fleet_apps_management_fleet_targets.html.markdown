---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_targets"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_targets"
description: |-
  Provides the list of Fleet Targets in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_targets
This data source provides the list of Fleet Targets in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns the list of all confirmed targets within a fleet.


## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_targets" "test_fleet_targets" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id

	#Optional
	display_name = var.fleet_target_display_name
	product = var.fleet_target_product
	resource_display_name = var.fleet_target_resource_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Required) Unique Fleet identifier.
* `product` - (Optional) Product Name.
* `resource_display_name` - (Optional) Resource Display Name.
* `resource_id` - (Optional) Resource Identifier


## Attributes Reference

The following attributes are exported:

* `fleet_target_collection` - The list of fleet_target_collection.

### FleetTarget Reference

The following attributes are exported:

* `items` - List of fleetTargets.
	* `compartment_id` - Tenancy OCID
	* `compliance_state` - The last known compliance state of the target.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `id` - The OCID of the resource.
	* `is_last_discovery_attempt_successful` - A boolean flag that represents whether the last discovery attempt was successful.
	* `product` - Product to which the target belongs to.
	* `resource` - Resource Information for the Target.
		* `resource_display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `resource_id` - The OCID of the resource.
	* `state` - The current state of the FleetTarget.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_of_last_discovery_attempt` - The time when last discovery was attempted.
	* `time_of_last_successful_discovery` - The time when the last successful discovery was made.
	* `version` - Current version of target.

