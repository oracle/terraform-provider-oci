---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec"
sidebar_current: "docs-oci-datasource-database_tools_runtime-database_tools_database_api_gateway_config_pool_auto_api_spec"
description: |-
  Provides details about a specific Database Tools Database Api Gateway Config Pool Auto Api Spec in Oracle Cloud Infrastructure Database Tools Runtime service
---

# Data Source: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec
This data source provides details about a specific Database Tools Database Api Gateway Config Pool Auto Api Spec resource in Oracle Cloud Infrastructure Database Tools Runtime service.

Get a Database Tools database API gateway config auto API spec resource

## Example Usage

```hcl
data "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec" "test_database_tools_database_api_gateway_config_pool_auto_api_spec" {
	#Required
	auto_api_spec_key = var.database_tools_database_api_gateway_config_pool_auto_api_spec_auto_api_spec_key
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	pool_key = var.database_tools_database_api_gateway_config_pool_auto_api_spec_pool_key
}
```

## Argument Reference

The following arguments are supported:

* `auto_api_spec_key` - (Required) The key of the auto API spec config.
* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `pool_key` - (Required) The key of the pool config.


## Attributes Reference

The following attributes are exported:

* `alias` - Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
* `database_object_name` - The name of the database object.
* `database_object_type` - The type of the database object.
* `description` - Description of the autoApiSpec.
* `display_name` - A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `key` - A system generated string that uniquely identifies an auto API spec sub resource within a given pool.
* `operations` - The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
* `roles` - The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
* `scope` - The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
* `security_schemes` - The security schemes that can access this resource. If not specified then the resource is public.
* `time_created` - The time the resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the resource was updated. An RFC3339 formatted datetime string.
* `type` - The type of the Database Tools database API gateway config auto API spec sub resource.

