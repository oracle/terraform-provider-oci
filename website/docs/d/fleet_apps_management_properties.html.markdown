---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_properties"
sidebar_current: "docs-oci-datasource-fleet_apps_management-properties"
description: |-
  Provides the list of Properties in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_properties
This data source provides the list of Properties in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of Properties for Tenancy.


## Example Usage

```hcl
data "oci_fleet_apps_management_properties" "test_properties" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.property_display_name
	id = var.property_id
	scope = var.property_scope
	state = var.property_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique Property identifier
* `scope` - (Optional) A filter to return only resources their scope matches the given lifecycleState.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `property_collection` - The list of property_collection.

### Property Reference

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

