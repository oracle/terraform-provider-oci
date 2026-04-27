---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_mcp_server"
sidebar_current: "docs-oci-resource-database_tools-database_tools_mcp_server"
description: |-
  Provides the Database Tools Mcp Server resource in Oracle Cloud Infrastructure Database Tools service
---

# oci_database_tools_database_tools_mcp_server
This resource provides the Database Tools Mcp Server resource in Oracle Cloud Infrastructure Database Tools service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-tools/latest/DatabaseToolsMcpServer

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databaseTools

Creates a new Database Tools MCP server.


## Example Usage

```hcl
resource "oci_database_tools_database_tools_mcp_server" "test_database_tools_mcp_server" {
	#Required
	compartment_id = var.compartment_id
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	display_name = var.database_tools_mcp_server_display_name
	domain_id = oci_identity_domain.test_domain.id
	storage {
		#Required
		type = var.database_tools_mcp_server_storage_type

		#Optional
		bucket {

			#Optional
			bucket = var.database_tools_mcp_server_storage_bucket_bucket
			namespace = var.database_tools_mcp_server_storage_bucket_namespace
		}
	}
	type = var.database_tools_mcp_server_type

	#Optional
	access_token_expiry_in_seconds = var.database_tools_mcp_server_access_token_expiry_in_seconds
	custom_roles {
		#Required
		description = var.database_tools_mcp_server_custom_roles_description
		display_name = var.database_tools_mcp_server_custom_roles_display_name
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.database_tools_mcp_server_description
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		type = var.database_tools_mcp_server_locks_type

		#Optional
		message = var.database_tools_mcp_server_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.database_tools_mcp_server_locks_time_created
	}
	refresh_token_expiry_in_seconds = var.database_tools_mcp_server_refresh_token_expiry_in_seconds
	runtime_identity = var.database_tools_mcp_server_runtime_identity
}
```

## Argument Reference

The following arguments are supported:

* `access_token_expiry_in_seconds` - (Optional) (Updatable) Access token expiry in seconds
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
* `custom_roles` - (Optional) (Updatable) Custom Roles associated with the MCP Server.
	* `description` - (Required) (Updatable) The description of the custom role.
	* `display_name` - (Required) (Updatable) The display name of the custom role.
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A human-readable description of the Database Tools MCP server.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
* `domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated identity domain.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `refresh_token_expiry_in_seconds` - (Optional) (Updatable) Refresh token expiry in seconds
* `runtime_identity` - (Optional) Specifies the identity used by the Database Tools MCP server to issue requests to other Oracle Cloud Infrastructure services (e.g., Secrets in Vault).
* `storage` - (Required) (Updatable) The storage option used when running a tool asynchronously.
	* `bucket` - (Required when type=OBJECT_STORAGE) (Updatable) A Cloud Storage bucket for an MCP Server.
		* `bucket` - (Required when type=OBJECT_STORAGE) (Updatable) The Object Storage bucket to use.
		* `namespace` - (Required when type=OBJECT_STORAGE) (Updatable) The Object Storage namespace to use.
	* `type` - (Required) (Updatable) The type of storage used for asynchronous tool calls.
* `type` - (Required) (Updatable) The Database Tools MCP server type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `access_token_expiry_in_seconds` - Access token expiry in seconds
* `built_in_roles` - Built-in roles associated with the MCP Server.
	* `description` - The description of the built-in role.
	* `display_name` - The display name of the built-in role.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools MCP server.
* `custom_roles` - Custom roles associated with the MCP Server.
	* `description` - The description of the custom role.
	* `display_name` - The display name of the custom role.
* `database_tools_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A human-readable description of the Database Tools MCP server.
* `display_name` - A user-friendly name. Does not have to be unique and can be updated. Avoid entering confidential information.
* `domain_app_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated domain application (Oracle Cloud Service).
* `domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated identity domain.
* `endpoints` - Invoke endpoints for the MCP server.
	* `endpoint` - The URI endpoint of the MCP server
	* `type` - The MCP server type
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools MCP server.
* `lifecycle_details` - A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `refresh_token_expiry_in_seconds` - Refresh token expiry in seconds
* `related_resource` - A related resource
	* `entity_type` - The resource entity type.
	* `identifier` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
* `runtime_identity` - Specifies the identity used by the Database Tools MCP server to issue requests to other Oracle Cloud Infrastructure services (e.g., Secrets in Vault).
* `state` - The current state of the Database Tools MCP server.
* `storage` - The storage option used when running a tool asynchronously.
	* `bucket` - A Cloud Storage bucket for an MCP Server.
		* `bucket` - The Object Storage bucket to use.
		* `namespace` - The Object Storage namespace to use.
	* `type` - The type of storage used for asynchronous tool calls.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools MCP server was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Database Tools MCP server was updated. An RFC3339 formatted datetime string.
* `type` - The Database Tools MCP server type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Mcp Server
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Mcp Server
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Mcp Server


## Import

DatabaseToolsMcpServers can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server "id"
```

