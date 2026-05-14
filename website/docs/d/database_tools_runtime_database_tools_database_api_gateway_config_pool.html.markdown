---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_pool"
description: |-
  Provides details about a specific Database Tools Database Api Gateway Config Pool in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool
This data source provides details about a specific Database Tools Database Api Gateway Config Pool resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a Database Tools database API gateway config pool resource

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool" "test_database_tools_database_api_gateway_config_pool" {
	#Required
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	pool_key = var.database_tools_database_api_gateway_config_pool_pool_key
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `pool_key` - (Required) The key of the pool config.


## Attributes Reference

The following attributes are exported:

* `advanced_properties` - Advanced pool properties.
* `database_actions_status` - Specifies to enable the Database Actions feature.
* `database_tools_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection. Specifies the Oracle Cloud Infrastructure database tools connection ocid to build the connection pool from.
* `display_name` - A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `initial_pool_size` - Specifies the initial size for the number of database connections that will be created for the pool.
* `jwt_profile_audience` - Specifies the expected audience for the JWT token. This value is used to validate the aud claim in the JWT token.
* `jwt_profile_issuer` - Specifies the issuer of the JWT token. This value is used to validate the iss claim in the JWT token.
* `jwt_profile_jwk_url` - Specifies the URL of the JSON Web Key (JWK) that is used to verify the signature of the JWT token.
* `jwt_profile_role_claim_name` - Specifies the JSON pointer to the claim in the JWT token that contains the roles of the users.
* `key` - A system generated string that uniquely identifies a pool sub resource.
* `max_pool_size` - Specifies the maximum number of database connections allowed for the pool.
* `min_pool_size` - Specifies the minimum number of database connections allowed for the pool.
* `pool_route_value` - The pool route value provided in requests to target this pool.
* `rest_enabled_sql_status` - Specifies whether the REST-Enabled SQL service is active.
* `time_created` - The time the resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the resource was updated. An RFC3339 formatted datetime string.
* `type` - The type of the Database Tools database API gateway config pool sub resource.

