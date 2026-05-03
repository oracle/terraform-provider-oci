---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_database_api_gateway_config_pool_api_spec"
description: |-
  Provides the Database Tools Database Api Gateway Config Pool Api Spec resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec
This resource provides the Database Tools Database Api Gateway Config Pool Api Spec resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Create a Database Tools database API gateway config API spec resource

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec" "test_database_tools_database_api_gateway_config_pool_api_spec" {
	#Required
	content = var.database_tools_database_api_gateway_config_pool_api_spec_content
	database_tools_database_api_gateway_config_id = oci_apm_config_config.test_config.id
	display_name = var.database_tools_database_api_gateway_config_pool_api_spec_display_name
	pool_key = var.database_tools_database_api_gateway_config_pool_api_spec_pool_key
	type = var.database_tools_database_api_gateway_config_pool_api_spec_type
}
```

## Argument Reference

The following arguments are supported:

* `content` - (Required) (Updatable) The content of a string-escaped Open API spec in JSON format.
* `database_tools_database_api_gateway_config_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `pool_key` - (Required) The key of the pool config.
* `type` - (Required) (Updatable) The type of the Database Tools database API gateway config API spec sub resource.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `content` - The content of a string-escaped Open API spec in JSON format.
* `display_name` - A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
* `key` - A system generated string that uniquely identifies an API spec sub resource within a given pool.
* `time_created` - The time the resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the resource was updated. An RFC3339 formatted datetime string.
* `type` - The type of the Database Tools database API gateway config API spec sub resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Database Api Gateway Config Pool Api Spec
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Database Api Gateway Config Pool Api Spec
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Database Api Gateway Config Pool Api Spec


## Import

DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec.test_database_tools_database_api_gateway_config_pool_api_spec "databaseToolsDatabaseApiGatewayConfigs/{databaseToolsDatabaseApiGatewayConfigId}/pools/{poolKey}/apiSpecs/{apiSpecKey}" 
```

