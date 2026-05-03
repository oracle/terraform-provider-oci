---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_database_api_gateway_config_pool_auto_api_spec"
description: |-
  Provides the Database Tools Database Api Gateway Config Pool Auto Api Spec resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec
This resource provides the Database Tools Database Api Gateway Config Pool Auto Api Spec resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Create a Database Tools database API gateway config auto API spec resource

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec" "test_database_tools_database_api_gateway_config_pool_auto_api_spec" {
	#Required
	database_object_name = oci_objectstorage_object.test_object.name
	database_object_type = var.database_tools_database_api_gateway_config_pool_auto_api_spec_database_object_type
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	display_name = var.database_tools_database_api_gateway_config_pool_auto_api_spec_display_name
	pool_key = var.database_tools_database_api_gateway_config_pool_auto_api_spec_pool_key
	type = var.database_tools_database_api_gateway_config_pool_auto_api_spec_type

	#Optional
	alias = var.database_tools_database_api_gateway_config_pool_auto_api_spec_alias
	description = var.database_tools_database_api_gateway_config_pool_auto_api_spec_description
	operations = var.database_tools_database_api_gateway_config_pool_auto_api_spec_operations
	roles = var.database_tools_database_api_gateway_config_pool_auto_api_spec_roles
	scope = var.database_tools_database_api_gateway_config_pool_auto_api_spec_scope
	security_schemes = var.database_tools_database_api_gateway_config_pool_auto_api_spec_security_schemes
}
```

## Argument Reference

The following arguments are supported:

* `alias` - (Optional) (Updatable) Used as the URI path element for this object. When not specified the objectName lowercase is the default value.
* `database_object_name` - (Required) (Updatable) The name of the database object.
* `database_object_type` - (Required) (Updatable) The type of the database object.
* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `description` - (Optional) (Updatable) Description of the autoApiSpec.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `operations` - (Optional) (Updatable) The operations to limit access to this resource. If not specified then the default is ["READ","WRITE"].
* `pool_key` - (Required) The key of the pool config.
* `roles` - (Optional) (Updatable) The name of the database API gateway config roles protecting the resource. Only valid for RBAC JWT Profile pools and BEARER securitySchemes.
* `scope` - (Optional) (Updatable) The name of the database API gateway config privilege protecting the resource. Only valid for SCOPE JWT Profile pools and BEARER securitySchemes.
* `security_schemes` - (Optional) (Updatable) The security schemes that can access this resource. If not specified then the resource is public.
* `type` - (Required) (Updatable) The type of the Database Tools database API gateway config auto API spec sub resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Database Api Gateway Config Pool Auto Api Spec
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Database Api Gateway Config Pool Auto Api Spec
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Database Api Gateway Config Pool Auto Api Spec


## Import

DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec.test_database_tools_database_api_gateway_config_pool_auto_api_spec "databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/autoApiSpecs/{autoApiSpecKey}" 
```

