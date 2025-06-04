---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_platform_configurations"
sidebar_current: "docs-oci-datasource-fleet_apps_management-platform_configurations"
description: |-
  Provides the list of Platform Configurations in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_platform_configurations
This data source provides the list of Platform Configurations in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the Platform Configurations in the specified compartment.
The query parameter `compartmentId` is required unless the query parameter `id` is specified.


## Example Usage

```hcl
data "oci_fleet_apps_management_platform_configurations" "test_platform_configurations" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.platform_configuration_compartment_id_in_subtree
	config_category = var.platform_configuration_config_category
	display_name = var.platform_configuration_display_name
	id = var.platform_configuration_id
	state = var.platform_configuration_state
	type = var.platform_configuration_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `compartment_id_in_subtree` - (Optional) If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. 
* `config_category` - (Optional) Config Category
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single Platform Configuration by id. Either compartmentId or id must be provided. 
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `type` - (Optional) A filter to return Platform Configurations whose type matches the given type.


## Attributes Reference

The following attributes are exported:

* `platform_configuration_collection` - The list of platform_configuration_collection.

### PlatformConfiguration Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID
* `config_category_details` - Config Category Details.
	* `compatible_products` - Products compatible with this Product. Provide products from the list of other products you have created that are compatible with the present one 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
	* `components` - Various components of the Product. For example:The administration server or node manager can be the components of the Oracle WebLogic Application server. Forms server or concurrent manager can be the components of the Oracle E-Business Suite. 
	* `config_category` - Category of configuration
	* `credentials` - OCID for the Credential name to be associated with the Product. These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
	* `instance_id` - The OCID of the resource.
	* `instance_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `patch_types` - Patch Types associated with this Product. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
	* `products` - Products that belong to the stack. For example, Oracle WebLogic and Java for the Oracle Fusion Middleware product stack. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
	* `sub_category_details` - ProductStack Config Category Details.
		* `components` - Various components of the Product. For example:The administration server or node manager can be the components of the Oracle WebLogic Application server. Forms server or concurrent manager can be the components of the Oracle E-Business Suite. 
		* `credentials` - OCID for the Credential name to be associated with the Product Stack. These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server. 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - The OCID of the resource.
		* `patch_types` - Patch Types associated with this Product Stack which will be considered as Product. 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - The OCID of the resource.
		* `sub_category` - SubCategory of Product Stack.
		* `versions` - Versions associated with the PRODUCT .  
	* `versions` - Versions associated with the PRODUCT .  
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `resource_region` - Associated region
* `state` - The current state of the PlatformConfiguration.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - The type of the configuration.

