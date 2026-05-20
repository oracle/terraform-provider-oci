---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_mcp_toolset_versions"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_mcp_toolset_versions"
description: |-
  Provides the list of Database Tools Mcp Toolset Versions in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_mcp_toolset_versions
This data source provides the list of Database Tools Mcp Toolset Versions in Oracle Cloud Infrastructure Database Tools service.

Returns a list of Database Tools Toolset versions

## Example Usage

```hcl
data "oci_database_tools_database_tools_mcp_toolset_versions" "test_database_tools_mcp_toolset_versions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	database_tools_mcp_server_id = oci_database_tools_database_tools_mcp_server.test_database_tools_mcp_server.id
	display_name = var.database_tools_mcp_toolset_version_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `database_tools_mcp_server_id` - (Optional) A filter to return only resources matching the specified `databaseToolsMcpServerId`.
* `display_name` - (Optional) A filter to return only resources that match the entire specified display name.


## Attributes Reference

The following attributes are exported:

* `database_tools_mcp_toolset_version_collection` - The list of database_tools_mcp_toolset_version_collection.

### DatabaseToolsMcpToolsetVersion Reference

The following attributes are exported:

* `items` - Array of MCP toolset type version summary items.
	* `default_version` - The default version for this toolset type.
	* `type` - The MCP toolset type.
	* `versions` - The version configurations available for this toolset type.
		* `default_allowed_roles` - The roles granted access to this toolset version by default.
		* `default_report_allowed_roles` - The roles granted access to this toolset version by default.
		* `description` - A description of this version.
		* `features` - Optional feature flags or attributes for this version.
		* `tools` - The tools available in this version.
			* `default_allowed_roles` - The roles granted access to this tool by default.
			* `default_status` - The default status of the tool in this version.
			* `description` - The description of the tool.
			* `display_name` - The display name of the tool.
			* `name` - The unique name of the tool.
		* `version` - The version number.

