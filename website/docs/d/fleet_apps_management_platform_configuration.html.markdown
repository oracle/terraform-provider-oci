---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_platform_configuration"
sidebar_current: "docs-oci-datasource-fleet_apps_management-platform_configuration"
description: |-
  Provides details about a specific Platform Configuration in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_platform_configuration
This data source provides details about a specific Platform Configuration resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a PlatformConfiguration by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
	#Required
	platform_configuration_id = oci_fleet_apps_management_platform_configuration.test_platform_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `platform_configuration_id` - (Required) unique PlatformConfiguration identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `config_category_details` - Config Category Details.
	* `compatible_products` - Products compatible with this Product. Provide products from the list of other products you have created that are compatible with the present one 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
	* `components` - Various components of the Product. For example:The administration server or node manager can be the components of the Oracle WebLogic Application server. Forms server or concurrent manager can be the components of the Oracle E-Business Suite. 
	* `config_category` - Category of configuration
	* `credentials` - OCID for the Credential name to be associated with the Product. These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - The OCID of the resource.
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

