---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_platform_configuration"
sidebar_current: "docs-oci-resource-fleet_apps_management-platform_configuration"
description: |-
  Provides the Platform Configuration resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_platform_configuration
This resource provides the Platform Configuration resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new PlatformConfiguration.


## Example Usage

```hcl
resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
	#Required
	compartment_id = var.compartment_id
	config_category_details {
		#Required
		config_category = var.platform_configuration_config_category_details_config_category

		#Optional
		compatible_products {

			#Optional
			display_name = var.platform_configuration_config_category_details_compatible_products_display_name
			id = var.platform_configuration_config_category_details_compatible_products_id
		}
		components = var.platform_configuration_config_category_details_components
		credentials {

			#Optional
			display_name = var.platform_configuration_config_category_details_credentials_display_name
			id = var.platform_configuration_config_category_details_credentials_id
		}
		patch_types {

			#Optional
			display_name = var.platform_configuration_config_category_details_patch_types_display_name
			id = var.platform_configuration_config_category_details_patch_types_id
		}
		products {

			#Optional
			display_name = var.platform_configuration_config_category_details_products_display_name
			id = var.platform_configuration_config_category_details_products_id
		}
		sub_category_details {
			#Required
			sub_category = var.platform_configuration_config_category_details_sub_category_details_sub_category

			#Optional
			components = var.platform_configuration_config_category_details_sub_category_details_components
			credentials {

				#Optional
				display_name = var.platform_configuration_config_category_details_sub_category_details_credentials_display_name
				id = var.platform_configuration_config_category_details_sub_category_details_credentials_id
			}
			patch_types {

				#Optional
				display_name = var.platform_configuration_config_category_details_sub_category_details_patch_types_display_name
				id = var.platform_configuration_config_category_details_sub_category_details_patch_types_id
			}
			versions = var.platform_configuration_config_category_details_sub_category_details_versions
		}
		versions = var.platform_configuration_config_category_details_versions
	}

	#Optional
	description = var.platform_configuration_description
	display_name = var.platform_configuration_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `config_category_details` - (Required) (Updatable) Config Category Details.
	* `compatible_products` - (Applicable when config_category=PRODUCT) (Updatable) Products compatible with this Product. Provide products from the list of other products you have created that are compatible with the present one 
		* `display_name` - (Applicable when config_category=PRODUCT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - (Required when config_category=PRODUCT) (Updatable) The OCID of the resource.
	* `components` - (Applicable when config_category=PRODUCT) (Updatable) Various components of the Product. For example:The administration server or node manager can be the components of the Oracle WebLogic Application server. Forms server or concurrent manager can be the components of the Oracle E-Business Suite. 
	* `config_category` - (Required) (Updatable) Category of configuration
	* `credentials` - (Applicable when config_category=PRODUCT) (Updatable) OCID for the Credential name to be associated with the Product. These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server. 
		* `display_name` - (Applicable when config_category=PRODUCT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - (Required when config_category=PRODUCT) (Updatable) The OCID of the resource.
	* `patch_types` - (Applicable when config_category=PRODUCT) (Updatable) Patch Types associated with this Product. 
		* `display_name` - (Applicable when config_category=PRODUCT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - (Required when config_category=PRODUCT) (Updatable) The OCID of the resource.
	* `products` - (Required when config_category=PRODUCT_STACK) (Updatable) Products that belong to the stack. For example, Oracle WebLogic and Java for the Oracle Fusion Middleware product stack. 
		* `display_name` - (Applicable when config_category=PRODUCT_STACK) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
		* `id` - (Required when config_category=PRODUCT_STACK) (Updatable) The OCID of the resource.
	* `sub_category_details` - (Applicable when config_category=PRODUCT_STACK) (Updatable) ProductStack Config Category Details.
		* `components` - (Applicable when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) Various components of the Product. For example:The administration server or node manager can be the components of the Oracle WebLogic Application server. Forms server or concurrent manager can be the components of the Oracle E-Business Suite. 
		* `credentials` - (Applicable when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) OCID for the Credential name to be associated with the Product Stack. These are useful for target discovery or lifecycle management activities, for example, Oracle WebLogic admin credentials for Oracle WebLogic Application server. 
			* `display_name` - (Applicable when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - (Required when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) The OCID of the resource.
		* `patch_types` - (Applicable when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) Patch Types associated with this Product Stack which will be considered as Product. 
			* `display_name` - (Applicable when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - (Required when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) The OCID of the resource.
		* `sub_category` - (Required) (Updatable) SubCategory of Product Stack.
		* `versions` - (Required when sub_category=PRODUCT_STACK_AS_PRODUCT) (Updatable) Versions associated with the PRODUCT .  
	* `versions` - (Required when config_category=PRODUCT) (Updatable) Versions associated with the PRODUCT .  
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Platform Configuration
	* `update` - (Defaults to 20 minutes), when updating the Platform Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Platform Configuration


## Import

PlatformConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_platform_configuration.test_platform_configuration "id"
```

