---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_configuration"
sidebar_current: "docs-oci-resource-psql-configuration"
description: |-
  Provides the Configuration resource in Oracle Cloud Infrastructure Psql service
---

# oci_psql_configuration
This resource provides the Configuration resource in Oracle Cloud Infrastructure Psql service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/postgresql/latest/Configuration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/psql

Creates a new configuration.


## Example Usage

```hcl
resource "oci_psql_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_id
	db_configuration_overrides {
		#Required
		items {
			#Required
			config_key = var.configuration_db_configuration_overrides_items_config_key
			overriden_config_value = var.configuration_db_configuration_overrides_items_overriden_config_value
		}
	}
	db_version = var.configuration_db_version
	display_name = var.configuration_display_name

	#Optional
	compatible_shapes = var.configuration_compatible_shapes
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.configuration_description
	freeform_tags = {"bar-key"= "value"}
	instance_memory_size_in_gbs = var.configuration_instance_memory_size_in_gbs
	instance_ocpu_count = var.configuration_instance_ocpu_count
	is_flexible = var.configuration_is_flexible
	shape = var.configuration_shape
	system_tags = var.configuration_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the configuration.
* `compatible_shapes` - (Optional) (Updatable) Indicates the collection of compatible shapes for this configuration. 
* `db_configuration_overrides` - (Required) Configuration overrides for a PostgreSQL instance.
	* `items` - (Required) List of configuration overridden values.
		* `config_key` - (Required) Configuration variable name.
		* `overriden_config_value` - (Required) User-selected variable value.
* `db_version` - (Required) Version of the PostgreSQL database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Details about the configuration set.
* `display_name` - (Required) (Updatable) A user-friendly display name for the configuration. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `instance_memory_size_in_gbs` - (Optional) Memory size in gigabytes with 1GB increment.

	Skip or set it's value to 0 if configuration is for a flexible shape. 
* `instance_ocpu_count` - (Optional) CPU core count.

	Skip or set it's value to 0 if configuration is for a flexible shape. 
* `is_flexible` - (Optional) Whether the configuration supports flexible shapes.
* `shape` - (Optional) The name of the shape for the configuration. 

	For multi-shape enabled configurations, it is set to PostgreSQL.X86 or similar. Please use compatibleShapes property to set the list of supported shapes. 
* `system_tags` - (Optional) System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the configuration.
* `compatible_shapes` - Indicates the collection of compatible shapes for this configuration. 
* `config_type` - The type of configuration. Either user-created or a default configuration.
* `configuration_details` - List of configuration details.
	* `items` - List of ConfigParms object.
		* `allowed_values` - Range or list of allowed values.
		* `config_key` - The configuration variable name.
		* `data_type` - Data type of the variable.
		* `default_config_value` - Default value for the configuration variable.
		* `description` - Details about the PostgreSQL parameter.
		* `is_overridable` - Whether the value can be overridden or not.
		* `is_restart_required` - If true, modifying this configuration value will require a restart of the database.
		* `overriden_config_value` - User-selected configuration variable value.
* `db_version` - Version of the PostgreSQL database.
* `default_config_id` - The Default configuration used for this configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description for the configuration.
* `display_name` - A user-friendly display name for the configuration. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier for the configuration. Immutable on creation.
* `instance_memory_size_in_gbs` - Memory size in gigabytes with 1GB increment.

	It's value is set to 0 if configuration is for a flexible shape. 
* `instance_ocpu_count` - CPU core count.

	It's value is set to 0 if configuration is for a flexible shape. 
* `is_flexible` - Whether the configuration supports flexible shapes.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - The name of the shape for the configuration. 

	For multi-shape enabled configurations, it is set to PostgreSQL. Please use compatibleShapes property to get list of supported shapes for such configurations. 
* `state` - The current state of the configuration.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the configuration was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Configuration
	* `update` - (Defaults to 20 minutes), when updating the Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Configuration


## Import

Configurations can be imported using the `id`, e.g.

```
$ terraform import oci_psql_configuration.test_configuration "id"
```

