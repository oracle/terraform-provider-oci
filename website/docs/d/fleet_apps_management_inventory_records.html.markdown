---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_inventory_records"
sidebar_current: "docs-oci-datasource-fleet_apps_management-inventory_records"
description: |-
  Provides the list of Inventory Records in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_inventory_records
This data source provides the list of Inventory Records in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of inventoryDetails.


## Example Usage

```hcl
data "oci_fleet_apps_management_inventory_records" "test_inventory_records" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.inventory_record_compartment_id_in_subtree
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	is_details_required = var.inventory_record_is_details_required
	resource_id = oci_cloud_guard_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. 
* `fleet_id` - (Optional) unique Fleet identifier
* `is_details_required` - (Optional) If set to true, inventory details will be returned. 
* `resource_id` - (Optional) Resource Identifier


## Attributes Reference

The following attributes are exported:

* `inventory_record_collection` - The list of inventory_record_collection.

### InventoryRecord Reference

The following attributes are exported:

* `items` - List of inventory targets
	* `architecture` - Architecture of the resource associated with the target
	* `compartment_id` - OCID of the compartment to which the resource belongs to.
	* `components` - List of target components
		* `component_name` - Name of the target component
		* `component_path` - Path of the component
		* `component_version` - Version of the target component
		* `properties` - List of component properties
			* `name` - Name of the inventory target property.
			* `value` - Value of the inventory target property.
	* `installed_patches` - List of details on the patches currently installed on the target
		* `patch_description` - Description for the installed patch
		* `patch_id` - OCID of the installed patch
		* `patch_level` - Patch Level.
		* `patch_name` - Name of the installed patch
		* `patch_type` - Type of patch applied
		* `time_applied` - Date on which the patch was applied to the target
		* `time_released` - The date on which patch was released.
	* `os_type` - OS installed on the resource associated with the target
	* `properties` - List of target properties
		* `name` - Name of the inventory target property.
		* `value` - Value of the inventory target property.
	* `state` - The current state of the Inventory target.
	* `target_id` - The id of the Inventory target.
	* `target_name` - Name of the target
	* `target_product_id` - OCID of the product installed at the target path
	* `target_product_name` - Name of the product installed at the target path
	* `target_resource_id` - OCID of the resource associated with the target
	* `target_resource_name` - Name of the resource associated with the target
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
	* `version` - Version of the product on the target

