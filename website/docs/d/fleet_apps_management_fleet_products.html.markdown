---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_products"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_products"
description: |-
  Provides the list of Fleet Products in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_products
This data source provides the list of Fleet Products in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of FleetProducts.


## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_products" "test_fleet_products" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.fleet_product_display_name
	resource_display_name = var.fleet_product_resource_display_name
	resource_id = oci_cloud_guard_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Required) unique Fleet identifier
* `resource_display_name` - (Optional) Resource Display Name
* `resource_id` - (Optional) Resource Identifier


## Attributes Reference

The following attributes are exported:

* `fleet_product_collection` - The list of fleet_product_collection.

### FleetProduct Reference

The following attributes are exported:

* `items` - List of fleetProducts.
	* `compartment_id` - Root Compartment Id.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `resource` - Resource Information for the Target
		* `resource_display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `resource_id` - The OCID of the resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_count` - Count of targets associated with the Product

