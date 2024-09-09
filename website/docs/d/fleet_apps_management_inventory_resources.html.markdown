---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_inventory_resources"
sidebar_current: "docs-oci-datasource-fleet_apps_management-inventory_resources"
description: |-
  Provides the list of Inventory Resources in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_inventory_resources
This data source provides the list of Inventory Resources in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of InventoryResources.


## Example Usage

```hcl
data "oci_fleet_apps_management_inventory_resources" "test_inventory_resources" {
	#Required
	compartment_id = var.compartment_id
	resource_compartment_id = oci_identity_compartment.test_compartment.id

	#Optional
	defined_tag_equals = var.inventory_resource_defined_tag_equals
	display_name = var.inventory_resource_display_name
	freeform_tag_equals = var.inventory_resource_freeform_tag_equals
	inventory_properties = var.inventory_resource_inventory_properties
	matching_criteria = var.inventory_resource_matching_criteria
	resource_region = var.inventory_resource_resource_region
	state = var.inventory_resource_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `defined_tag_equals` - (Optional) A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned. Each item in the list has the format "{namespace}.{tagName}={value}".  All inputs are case-insensitive. Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR". Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND". 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `freeform_tag_equals` - (Optional) A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned. The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive. Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND". 
* `inventory_properties` - (Optional) A list of inventory properties filters to apply. The key for each inventory property and value for each resource type is "{resourceType}.{inventoryProperty}={value}". 
* `matching_criteria` - (Optional) Fetch resources matching matching ANY or ALL criteria passed as params in "tags" and "inventoryProperties"
* `resource_compartment_id` - (Required) Resource Compartment ID
* `resource_region` - (Optional) Resource Region
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `inventory_resource_collection` - The list of inventory_resource_collection.

### InventoryResource Reference

The following attributes are exported:

* `items` - List of InventoryResources.
	* `availability_domain` - Availability Domain of the resource
	* `compartment_id` - OCID of the compartment to which the resource belongs to.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the resource.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `resource_compartment_id` - Compartment Id of the resource
	* `resource_region` - Region the resource belongs to
	* `state` - The current state of the Resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `type` - Type of the Resource.

