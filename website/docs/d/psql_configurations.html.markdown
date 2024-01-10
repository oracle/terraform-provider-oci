---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_configurations"
sidebar_current: "docs-oci-datasource-psql-configurations"
description: |-
  Provides the list of Configurations in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_configurations
This data source provides the list of Configurations in Oracle Cloud Infrastructure Psql service.

Returns a list of configurations.


## Example Usage

```hcl
data "oci_psql_configurations" "test_configurations" {

	#Optional
	compartment_id = var.compartment_id
	configuration_id = oci_psql_configuration.test_configuration.id
	db_version = var.configuration_db_version
	display_name = var.configuration_display_name
	shape = var.configuration_shape
	state = var.configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `configuration_id` - (Optional) A unique identifier for the configuration.
* `db_version` - (Optional) Verison of the PostgreSQL database, such as 14.9.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `shape` - (Optional) The name of the shape for the configuration. Example: `VM.Standard.E4.Flex` 
* `state` - (Optional) A filter to return only resources if their `lifecycleState` matches the given `lifecycleState`.


## Attributes Reference

The following attributes are exported:

* `configuration_collection` - The list of configuration_collection.

### Configuration Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the configuration.
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
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description for the configuration.
* `display_name` - A user-friendly display name for the configuration. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier for the configuration. Immutable on creation.
* `instance_memory_size_in_gbs` - Memory size in gigabytes with 1GB increment. 
* `instance_ocpu_count` - CPU core count. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - The name of the shape for the configuration. Example: `VM.Standard.E4.Flex` 
* `state` - The current state of the configuration.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the configuration was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

