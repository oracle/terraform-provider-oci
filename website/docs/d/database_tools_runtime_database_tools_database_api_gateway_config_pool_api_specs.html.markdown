---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_pool_api_specs"
description: |-
  Provides the list of Database Tools Database Api Gateway Config Pool Api Specs in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs
This data source provides the list of Database Tools Database Api Gateway Config Pool Api Specs in Oracle Cloud Infrastructure Database Tools Runtime service.

Returns a list of Database Tools database API gateway config API spec resources

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs" "test_database_tools_database_api_gateway_config_pool_api_specs" {
	#Required
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	pool_key = var.database_tools_database_api_gateway_config_pool_api_spec_pool_key

	#Optional
	display_name = var.database_tools_database_api_gateway_config_pool_api_spec_display_name
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `display_name` - (Optional) A filter to return only resources that match the entire specified display name.
* `pool_key` - (Required) The key of the pool config.


## Attributes Reference

The following attributes are exported:

* `database_tools_database_api_gateway_config_pool_api_spec_collection` - The list of database_tools_database_api_gateway_config_pool_api_spec_collection.

### DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec Reference

The following attributes are exported:

* `content` - The content of a string-escaped Open API spec in JSON format.
* `display_name` - A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `key` - A system generated string that uniquely identifies an API spec sub resource within a given pool.
* `time_created` - The time the resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the resource was updated. An RFC3339 formatted datetime string.
* `type` - The type of the Database Tools database API gateway config API spec sub resource.

