---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_global"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_database_api_gateway_config_global"
description: |-
  Provides the Database Tools Database Api Gateway Config Global resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_database_api_gateway_config_global
This resource provides the Database Tools Database Api Gateway Config Global resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Update a Database Tools database API gateway config global resource

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_global" "test_database_tools_database_api_gateway_config_global" {
	#Required
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	global_key = var.database_tools_database_api_gateway_config_global_global_key
	type = var.database_tools_database_api_gateway_config_global_type

	#Optional
	advanced_properties = var.database_tools_database_api_gateway_config_global_advanced_properties
	certificate_bundle {
		#Required
		type = var.database_tools_database_api_gateway_config_global_certificate_bundle_type

		#Optional
		certificate_private_key {

			#Optional
			format = var.database_tools_database_api_gateway_config_global_certificate_bundle_certificate_private_key_format
			path = var.database_tools_database_api_gateway_config_global_certificate_bundle_certificate_private_key_path
		}
		certificate_public {

			#Optional
			format = var.database_tools_database_api_gateway_config_global_certificate_bundle_certificate_public_format
			path = var.database_tools_database_api_gateway_config_global_certificate_bundle_certificate_public_path
		}
	}
	database_api_status = var.database_tools_database_api_gateway_config_global_database_api_status
	document_root = var.database_tools_database_api_gateway_config_global_document_root
	http_port = var.database_tools_database_api_gateway_config_global_http_port
	https_port = var.database_tools_database_api_gateway_config_global_https_port
	pool_route = var.database_tools_database_api_gateway_config_global_pool_route
	pool_routing_header = var.database_tools_database_api_gateway_config_global_pool_routing_header
}
```

## Argument Reference

The following arguments are supported:

* `advanced_properties` - (Optional) (Updatable) Advanced global properties.
* `certificate_bundle` - (Optional) (Updatable) The certificate bundle that describes the SSL certicicate. Ignored if the httpsPort is 0.
	* `certificate_private_key` - (Applicable when type=FILENAME) (Updatable) Describes a certificate private key file to be used with SSL
		* `format` - (Applicable when type=FILENAME) (Updatable) The format of the file
		* `path` - (Applicable when type=FILENAME) (Updatable) The path to the file
	* `certificate_public` - (Applicable when type=FILENAME) (Updatable) Describes a certificate file to be used with SSL. Ignored if the httpsPort is 0.
		* `format` - (Applicable when type=FILENAME) (Updatable) The format of the file
		* `path` - (Applicable when type=FILENAME) (Updatable) The path to the file
	* `type` - (Required) (Updatable) The type of the certificate.
* `database_api_status` - (Optional) (Updatable) ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `document_root` - (Optional) (Updatable) The location of the static resources to be served under the / root server path.
* `global_key` - (Required) The key of the global config.
* `http_port` - (Optional) (Updatable) Specifies the HTTP listen port. 0 disables HTTP. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended.
* `https_port` - (Optional) (Updatable) Specifies the HTTPS listen port. 0 disables HTTPS. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended. ORDS will use a self-signed certificate if a certificate bundle is not provided.
* `pool_route` - (Optional) (Updatable) How the target pool route value is determined for a HTTP request.
* `pool_routing_header` - (Optional) (Updatable) The request header name providing the pool route value.
* `type` - (Required) (Updatable) The type of the Database Tools database API gateway config global settings resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advanced_properties` - Advanced global properties.
* `certificate_bundle` - The certificate bundle that describes the SSL certicicate. Ignored if the httpsPort is 0.
	* `certificate_private_key` - Describes a certificate private key file to be used with SSL
		* `format` - The format of the file
		* `path` - The path to the file
	* `certificate_public` - Describes a certificate file to be used with SSL. Ignored if the httpsPort is 0.
		* `format` - The format of the file
		* `path` - The path to the file
	* `type` - The type of the certificate.
* `database_api_status` - ORDS database API is a database management and monitoring REST API. Database Actions requires this feature.
* `document_root` - The location of the static resources to be served under the / root server path.
* `http_port` - Specifies the HTTP listen port. 0 disables HTTP. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended.
* `https_port` - Specifies the HTTPS listen port. 0 disables HTTPS. Use of ports below 1024 requires elevated (root) privileges and is generally discouraged; deployment on non-privileged ports (1024–65535) is recommended. ORDS will use a self-signed certificate if a certificate bundle is not provided.
* `key` - A string that uniquely identifies a Database Tools database API gateway config global settings resource.
* `metadata_source` - The RESTful service definition location.
* `pool_route` - How the target pool route value is determined for a HTTP request.
* `pool_routing_header` - The request header name providing the pool route value.
* `time_created` - The time the resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the resource was updated. An RFC3339 formatted datetime string.
* `type` - The type of the Database Tools database API gateway config global settings resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Database Api Gateway Config Global
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Database Api Gateway Config Global
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Database Api Gateway Config Global


## Import

DatabaseToolsDatabaseApiGatewayConfigGlobals can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_database_api_gateway_config_global.test_database_tools_database_api_gateway_config_global "databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/globals/{globalKey}" 
```

