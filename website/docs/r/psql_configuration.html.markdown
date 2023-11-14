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

Creates a new Configuration Set.


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
	instance_memory_size_in_gbs = var.configuration_instance_memory_size_in_gbs
	instance_ocpu_count = var.configuration_instance_ocpu_count
	shape = var.configuration_shape

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.configuration_description
	freeform_tags = {"bar-key"= "value"}
	system_tags = var.configuration_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier
* `db_configuration_overrides` - (Required) Configuration Overrides for PGSQL instance.
	* `items` - (Required) List of configuration overriden values
		* `config_key` - (Required) Key is the configuration key.
		* `overriden_config_value` - (Required) User selected configuration value
* `db_version` - (Required) Version of the Postgresql DB
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Details about the Configuration Set.
* `display_name` - (Required) (Updatable) configuration display name
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `instance_memory_size_in_gbs` - (Required) Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape. 
* `instance_ocpu_count` - (Required) CPU cpuCoreCount. Min value is 1. Max value depends on the shape. 
* `shape` - (Required) Compute Shape Name like VM.Standard3.Flex.
* `system_tags` - (Optional) System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Config compartment identifier
* `configuration_details` - List of DB Configuration Values.
	* `items` - List of ConfigParms object.
		* `allowed_values` - Range or list of allowed values
		* `config_key` - Key is the configuration key.
		* `data_type` - Describes about the Datatype value.
		* `default_config_value` - Default value
		* `description` - Details about the Postgresql params.
		* `is_overridable` - This flags tells whether the value is overridable or not.
		* `is_restart_required` - If true, modfying this configuration value will requires restart.
		* `overriden_config_value` - User selected configuration value
* `db_version` - Version of the Postgresql DB
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Config description
* `display_name` - Config display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `instance_memory_size_in_gbs` - Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape. 
* `instance_ocpu_count` - CPU cpuCoreCount. Min value is 1. Max value depends on the shape. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - Compute Shape Name like VM.Standard3.Flex.
* `state` - The current state of the Configuration.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time Configuration was created. An RFC3339 formatted datetime string

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

