---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_property"
sidebar_current: "docs-oci-resource-fleet_apps_management-property"
description: |-
  Provides the Property resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_property
This resource provides the Property resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new Property.


## Example Usage

```hcl
resource "oci_fleet_apps_management_property" "test_property" {
	#Required
	compartment_id = var.compartment_id
	selection = var.property_selection
	value_type = var.property_value_type

	#Optional
	display_name = var.property_display_name
	values = var.property_values
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `selection` - (Required) (Updatable) Text selection of the category
* `value_type` - (Required) (Updatable) Format of the value
* `values` - (Optional) (Updatable) Values of the property (must be a single value if selection = 'single choice')


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `resource_region` - Associated region
* `scope` - The scope of the property
* `selection` - Text selection of the category
* `state` - The current state of the Property.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - The type of the property.
* `value_type` - Format of the value
* `values` - Values of the property (must be a single value if selection = 'single choice')

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Property
	* `update` - (Defaults to 20 minutes), when updating the Property
	* `delete` - (Defaults to 20 minutes), when destroying the Property


## Import

Properties can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_property.test_property "id"
```

