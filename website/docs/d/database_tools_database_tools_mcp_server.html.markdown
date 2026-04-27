---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_mcp_server"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_mcp_server"
description: |-
  Provides details about a specific Database Tools Mcp Server in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_mcp_server
This data source provides details about a specific Database Tools Mcp Server resource in Oracle Cloud Infrastructure Database Tools service.

Gets details of the specified Database Tools mcpserver.

## Example Usage

```hcl
data "oci_database_tools_database_tools_mcp_server" "test_database_tools_mcp_server" {
	#Required
	database_tools_mcp_server_id = oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_mcp_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools MCP server.


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

