---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_default_configuration"
sidebar_current: "docs-oci-datasource-psql-default_configuration"
description: |-
  Provides details about a specific Default Configuration in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_default_configuration
This data source provides details about a specific Default Configuration resource in Oracle Cloud Infrastructure Psql service.

Gets a Default Configuration by identifier

## Example Usage

```hcl
data "oci_psql_default_configuration" "test_default_configuration" {
	#Required
	default_configuration_id = oci_psql_default_configuration.test_default_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `default_configuration_id` - (Required) unique Configuration identifier


## Attributes Reference

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

