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

Gets a default configuration by identifier.

## Example Usage

```hcl
data "oci_psql_default_configuration" "test_default_configuration" {
	#Required
	default_configuration_id = oci_psql_default_configuration.test_default_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `default_configuration_id` - (Required) A unique identifier for the configuration.


## Attributes Reference

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
* `instance_ocpu_count` - CPU core count. Minimum value is 1. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `shape` - The name of the shape for the configuration. Example: `VM.Standard.E4.Flex` 
* `state` - The current state of the configuration.
* `time_created` - The date and time that the configuration was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

