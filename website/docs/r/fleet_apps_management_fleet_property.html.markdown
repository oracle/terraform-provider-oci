---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_property"
sidebar_current: "docs-oci-resource-fleet_apps_management-fleet_property"
description: |-
  Provides the Fleet Property resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_fleet_property
This resource provides the Fleet Property resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new FleetProperty.


## Example Usage

```hcl
resource "oci_fleet_apps_management_fleet_property" "test_fleet_property" {
	#Required
	compartment_id = var.compartment_id
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	property_id = oci_fleet_apps_management_property.test_property.id
	value = var.fleet_property_value
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `fleet_id` - (Required) unique Fleet identifier
* `property_id` - (Required) Property Id.
* `value` - (Required) (Updatable) Value of the Property


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet Property
	* `update` - (Defaults to 20 minutes), when updating the Fleet Property
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet Property


## Import

Import is not supported for this resource.

