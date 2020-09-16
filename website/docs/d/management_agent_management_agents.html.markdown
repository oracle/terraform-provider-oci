---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agents"
sidebar_current: "docs-oci-datasource-management_agent-management_agents"
description: |-
  Provides the list of Management Agents in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agents
This data source provides the list of Management Agents in Oracle Cloud Infrastructure Management Agent service.

Returns a list of Management Agent.


## Example Usage

```hcl
data "oci_management_agent_management_agents" "test_management_agents" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.management_agent_display_name
	platform_type = var.management_agent_platform_type
	plugin_name = var.management_agent_plugin_name
	state = var.management_agent_state
	version = var.management_agent_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment from which the Management Agents to be listed.
* `display_name` - (Optional) Filter to return only Management Agents having the particular display name.
* `platform_type` - (Optional) Filter to return only Management Agents having the particular platform type.
* `plugin_name` - (Optional) Filter to return only Management Agents having the particular Plugin installed.
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.
* `version` - (Optional) Filter to return only Management Agents having the particular agent version.


## Attributes Reference

The following attributes are exported:

* `management_agents` - The list of management_agents.

### ManagementAgent Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Management Agent Name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host` - Management Agent host machine name
* `id` - agent identifier
* `install_key_id` - agent install key identifier
* `install_path` - Path where Management Agent is installed
* `is_agent_auto_upgradable` - true if the agent can be upgraded automatically; false if it must be upgraded manually. true is currently unsupported.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `platform_name` - Platform Name
* `platform_type` - Platform Type
* `platform_version` - Platform Version
* `plugin_list` - list of managementAgentPlugins associated with the agent
	* `plugin_display_name` - Management Agent Plugin Identifier, can be renamed
	* `plugin_id` - Plugin Id
	* `plugin_name` - Management Agent Plugin Name
	* `plugin_version` - Plugin Version
* `state` - The current state of managementAgent
* `time_created` - The time the Management Agent was created. An RFC3339 formatted datetime string
* `time_last_heartbeat` - The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
* `time_updated` - The time the Management Agent was updated. An RFC3339 formatted datetime string
* `version` - Management Agent Version

