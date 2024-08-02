---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_default_configurations"
sidebar_current: "docs-oci-datasource-psql-default_configurations"
description: |-
  Provides the list of Default Configurations in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_default_configurations
This data source provides the list of Default Configurations in Oracle Cloud Infrastructure Psql service.

Returns a list of default configurations.


## Example Usage

```hcl
data "oci_psql_default_configurations" "test_default_configurations" {

	#Optional
	configuration_id = oci_psql_configuration.test_configuration.id
	db_version = var.default_configuration_db_version
	display_name = var.default_configuration_display_name
	shape = var.default_configuration_shape
	state = var.default_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `configuration_id` - (Optional) A unique identifier for the configuration.
* `db_version` - (Optional) Version of the PostgreSQL database, such as 14.9.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `shape` - (Optional) The name of the shape for the configuration. Example: `VM.Standard.E4.Flex` 
* `state` - (Optional) A filter to return only resources if their `lifecycleState` matches the given `lifecycleState`.


## Attributes Reference

The following attributes are exported:

* `default_configuration_collection` - The list of default_configuration_collection.

### DefaultConfiguration Reference

The following attributes are exported:

* `configuration_details` - List of default configuration values for databases.
	* `items` - List of ConfigParms object.
		* `allowed_values` - Range or list of allowed values.
		* `config_key` - The configuration variable name.
		* `data_type` - Data type of the variable.
		* `default_config_value` - Default value for the variable.
		* `description` - Details about the PostgreSQL variable.
		* `is_overridable` - Whether the value can be overridden or not.
		* `is_restart_required` - If true, modifying this configuration value will require a restart.
* `db_version` - Version of the PostgreSQL database.
* `description` - A description for the configuration.
* `display_name` - A user-friendly display name for the configuration.
* `id` - A unique identifier for the configuration.
* `instance_memory_size_in_gbs` - Memory size in gigabytes with 1GB increment.

	Its value is set to 0 if configuration is for a flexible shape. 
* `instance_ocpu_count` - CPU core count.

	Its value is set to 0 if configuration is for a flexible shape. 
* `is_flexible` - True if the configuration supports flexible shapes, false otherwise.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - The name of the shape for the configuration. Example: `VM.Standard.E4.Flex` 
* `state` - The current state of the configuration.
* `time_created` - The date and time that the configuration was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

