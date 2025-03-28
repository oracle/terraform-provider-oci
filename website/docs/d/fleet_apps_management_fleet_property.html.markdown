---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_property"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_property"
description: |-
  Provides details about a specific Fleet Property in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_property
This data source provides details about a specific Fleet Property resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a Fleet Property by identifier.

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_property" "test_fleet_property" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	fleet_property_id = oci_fleet_apps_management_property.test_property.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) Unique Fleet identifier.
* `fleet_property_id` - (Required) unique FleetProperty identifier.


## Attributes Reference

The following attributes are exported:

* `allowed_values` - Values of the property (must be a single value if selectionType = 'SINGLE_CHOICE').
* `compartment_id` - Tenancy OCID
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `id` - The unique id of the resource.
* `property_id` - OCID referring to global level metadata property.
* `selection_type` - Text selection of the property.
* `state` - The current state of the FleetProperty.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `value` - Value of the Property.
* `value_type` - Format of the value.

