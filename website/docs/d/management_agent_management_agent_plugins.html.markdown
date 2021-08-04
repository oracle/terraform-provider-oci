---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_plugins"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_plugins"
description: |-
  Provides the list of Management Agent Plugins in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_plugins
This data source provides the list of Management Agent Plugins in Oracle Cloud Infrastructure Management Agent service.

Returns a list of managementAgentPlugins.


## Example Usage

```hcl
data "oci_management_agent_management_agent_plugins" "test_management_agent_plugins" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.management_agent_plugin_display_name
	platform_type = var.management_agent_plugin_platform_type
	state = var.management_agent_plugin_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `display_name` - (Optional) Filter to return only Management Agent Plugins having the particular display name.
* `platform_type` - (Optional) Filter to return only results having the particular platform type.
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.


## Attributes Reference

The following attributes are exported:

* `management_agent_plugins` - The list of management_agent_plugins.

### ManagementAgentPlugin Reference

The following attributes are exported:

* `description` - Management Agent Plugin description
* `display_name` - Management Agent Plugin Display Name
* `id` - Management Agent Plugin Id
* `is_console_deployable` - A flag to indicate whether a given plugin can be deployed from Agent Console UI or not.
* `name` - Management Agent Plugin Name
* `state` - The current state of Management Agent Plugin
* `supported_platform_types` - Supported Platform Types
* `version` - Management Agent Plugin Version

