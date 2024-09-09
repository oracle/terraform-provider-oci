---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_properties"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_properties"
description: |-
  Provides the list of Fleet Properties in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_properties
This data source provides the list of Fleet Properties in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of FleetProperties.


## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_properties" "test_fleet_properties" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.fleet_property_display_name
	id = var.fleet_property_id
	state = var.fleet_property_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Required) unique Fleet identifier
* `id` - (Optional) unique FleetProperty identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `fleet_property_collection` - The list of fleet_property_collection.

### FleetProperty Reference

The following attributes are exported:

* `allowed_values` - Values of the category (must be a single value if selection = 'single choice')
* `compartment_id` - Tenancy OCID
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `id` - The unique id of the resource.
* `property_id` - Property Id Ocid.
* `selection_type` - Text selection of the category
* `state` - The current state of the FleetProperty.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `value` - Value of the Property
* `value_type` - Format of the value

