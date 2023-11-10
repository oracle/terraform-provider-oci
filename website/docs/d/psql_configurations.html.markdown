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

Returns a list of Configurations.


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
* `configuration_id` - (Optional) unique Configuration identifier
* `db_version` - (Optional) Verison of the Postgresql DB like 14.5
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `shape` - (Optional) Shape name of the compute like VM.Standard.E4.Flex or VM.Standard3.Flex
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `configuration_collection` - The list of configuration_collection.

### Configuration Reference

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

