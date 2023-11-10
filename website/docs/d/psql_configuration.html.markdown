---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_configuration"
sidebar_current: "docs-oci-datasource-psql-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Psql service.

Gets a Configuration by identifier

## Example Usage

```hcl
data "oci_psql_configuration" "test_configuration" {
	#Required
	configuration_id = oci_psql_configuration.test_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `configuration_id` - (Required) unique Configuration identifier


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

