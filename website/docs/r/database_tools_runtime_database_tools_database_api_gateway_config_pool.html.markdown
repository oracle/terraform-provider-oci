---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_database_api_gateway_config_pool"
description: |-
  Provides the Database Tools Database Api Gateway Config Pool resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_database_api_gateway_config_pool
This resource provides the Database Tools Database Api Gateway Config Pool resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Create a Database Tools database API gateway config pool resource

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool" "test_database_tools_database_api_gateway_config_pool" {
	#Required
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	display_name = var.database_tools_database_api_gateway_config_pool_display_name
	pool_route_value = var.database_tools_database_api_gateway_config_pool_pool_route_value
	type = var.database_tools_database_api_gateway_config_pool_type

	#Optional
	advanced_properties = var.database_tools_database_api_gateway_config_pool_advanced_properties
	database_actions_status = var.database_tools_database_api_gateway_config_pool_database_actions_status
	initial_pool_size = var.database_tools_database_api_gateway_config_pool_initial_pool_size
	jwt_profile_audience = var.database_tools_database_api_gateway_config_pool_jwt_profile_audience
	jwt_profile_issuer = var.database_tools_database_api_gateway_config_pool_jwt_profile_issuer
	jwt_profile_jwk_url = var.database_tools_database_api_gateway_config_pool_jwt_profile_jwk_url
	jwt_profile_role_claim_name = var.database_tools_database_api_gateway_config_pool_jwt_profile_role_claim_name
	max_pool_size = var.database_tools_database_api_gateway_config_pool_max_pool_size
	min_pool_size = var.database_tools_database_api_gateway_config_pool_min_pool_size
	rest_enabled_sql_status = var.database_tools_database_api_gateway_config_pool_rest_enabled_sql_status
}
```

## Argument Reference

The following arguments are supported:

* `advanced_properties` - (Optional) (Updatable) Advanced pool properties.
* `database_actions_status` - (Optional) (Updatable) Specifies to enable the Database Actions feature.
* `database_tools_connection_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection. Specifies the Oracle Cloud Infrastructure database tools connection ocid to build the connection pool from.
* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `initial_pool_size` - (Optional) (Updatable) Specifies the initial size for the number of database connections that will be created for the pool.
* `jwt_profile_audience` - (Optional) (Updatable) Specifies the expected audience for the JWT token. This value is used to validate the aud claim in the JWT token.
* `jwt_profile_issuer` - (Optional) (Updatable) Specifies the issuer of the JWT token. This value is used to validate the iss claim in the JWT token.
* `jwt_profile_jwk_url` - (Optional) (Updatable) Specifies the URL of the JSON Web Key (JWK) that is used to verify the signature of the JWT token.
* `jwt_profile_role_claim_name` - (Optional) (Updatable) Specifies the JSON pointer to the claim in the JWT token that contains the roles of the users.
* `max_pool_size` - (Optional) (Updatable) Specifies the maximum number of database connections allowed for the pool.
* `min_pool_size` - (Optional) (Updatable) Specifies the minimum number of database connections allowed for the pool.
* `pool_route_value` - (Required) (Updatable) The pool route value provided in requests to target this pool.
* `rest_enabled_sql_status` - (Optional) (Updatable) Specifies whether the REST-Enabled SQL service is active.
* `type` - (Required) (Updatable) The type of the Database Tools database API gateway config pool sub resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Database Api Gateway Config Pool
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Database Api Gateway Config Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Database Api Gateway Config Pool


## Import

DatabaseToolsDatabaseApiGatewayConfigPools can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool "databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}" 
```

