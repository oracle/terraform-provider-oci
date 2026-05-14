---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_advanced_properties"
description: |-
  Provides the list of Database Tools Database Api Gateway Config Advanced Properties in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties
This data source provides the list of Database Tools Database Api Gateway Config Advanced Properties in Oracle Cloud Infrastructure Database Tools Runtime service.

Returns list of database API gateway config setting descriptions to be provided as advanced properties.

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties" "test_database_tools_database_api_gateway_config_advanced_properties" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `database_tools_database_api_gateway_config_advanced_property_summary_collection` - The list of database_tools_database_api_gateway_config_advanced_property_summary_collection.

### DatabaseToolsDatabaseApiGatewayConfigAdvancedProperty Reference

The following attributes are exported:

* `items` - List of database API gateway config setting descriptions to be provided as advanced properties.
	* `category_display_name` - A user-friendly name of a category.
	* `category_key` - The category of the Database Tools database API gateway config global setting.
	* `config_types` - The config types that support this advanced property. The supported types are GLOBAL and POOL. 
	* `data_type` - The data type of a database API gateway config setting.
	* `database_tools_connection_types` - The type of database (as determined by a type of Database Tools connection) to which this setting applies.  The advancedProperty applies to all types of Database Tools connection when null. This is only applicable when configTypes includes POOL. 
	* `default_value` - The default value (if applicable) of a database API gateway config setting.
	* `description` - A user-friendly description of a database API gateway config setting.
	* `display_name` - A user-friendly name.
	* `documentation_url` - Uniform resource locator (URL) of documentation related to this setting.
	* `hint_text` - Hint text for a database API gateway config setting.
	* `key` - A string that uniquely identifies a Database Tools database API gateway config global settings resource.
	* `list_of_values` - A list of string values (if applicable) supported by this database API gateway config setting.
	* `max_value` - A maximum numeric value (if applicable) of a database API gateway config setting.
	* `min_value` - A minimum numeric value (if applicable) of a database API gateway config setting.

