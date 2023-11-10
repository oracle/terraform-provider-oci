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

Returns a list of Default Configurations.


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

* `configuration_id` - (Optional) unique Configuration identifier
* `db_version` - (Optional) Verison of the Postgresql DB like 14.5
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `shape` - (Optional) Shape name of the compute like VM.Standard.E4.Flex or VM.Standard3.Flex
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `default_configuration_collection` - The list of default_configuration_collection.

### DefaultConfiguration Reference

The following attributes are exported:

* `configuration_details` - List of DB default Configuration Values.
	* `items` - List of ConfigParms object.
		* `allowed_values` - Range or list of allowed values
		* `config_key` - Key is the configuration key.
		* `data_type` - Describes about the Datatype value.
		* `default_config_value` - Default value
		* `description` - Details about the Postgresql params.
		* `is_overridable` - This flags tells whether the value is overridable or not.
		* `is_restart_required` - If true, modfying this configuration value will requires restart.
* `db_version` - Version of the Postgresql DB
* `description` - Config description
* `display_name` - Config display name
* `id` - Unique identifier that is immutable on creation
* `instance_memory_size_in_gbs` - Memory Size in GB with 1GB increment. Min value matches the cpuCoreCount. Max value depends on the shape. 
* `instance_ocpu_count` - CPU cpuCoreCount. Min value is 1. Max value depends on the shape. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - Compute Shape Name like VM.Standard3.Flex.
* `state` - The current state of the Configuration.
* `time_created` - The time Configuration was created. An RFC3339 formatted datetime string

