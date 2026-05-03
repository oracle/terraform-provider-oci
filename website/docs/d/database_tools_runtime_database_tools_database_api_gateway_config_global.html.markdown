---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_global"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_global"
description: |-
  Provides details about a specific Database Tools Database Api Gateway Config Global in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_global
This data source provides details about a specific Database Tools Database Api Gateway Config Global resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a Database Tools database API gateway config global resource

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_global" "test_database_tools_database_api_gateway_config_global" {
	#Required
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	global_key = var.database_tools_database_api_gateway_config_global_global_key
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `global_key` - (Required) The key of the global config.


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

