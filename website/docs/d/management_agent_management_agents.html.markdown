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
	availability_status = var.management_agent_availability_status
	display_name = var.management_agent_display_name
	host_id = oci_management_agent_host.test_host.id
	install_type = var.management_agent_install_type
	is_customer_deployed = var.management_agent_is_customer_deployed
	platform_type = var.management_agent_platform_type
	plugin_name = var.management_agent_plugin_name
	state = var.management_agent_state
	version = var.management_agent_version
}
```

## Argument Reference

The following arguments are supported:

* `availability_status` - (Optional) Filter to return only Management Agents in the particular availability status.
* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `display_name` - (Optional) Filter to return only Management Agents having the particular display name.
* `host_id` - (Optional) Filter to return only Management Agents having the particular agent host id.
* `install_type` - (Optional) A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
* `is_customer_deployed` - (Optional) true, if the agent image is manually downloaded and installed. false, if the agent is deployed as a plugin in Oracle Cloud Agent.
* `platform_type` - (Optional) Filter to return only results having the particular platform type.
* `plugin_name` - (Optional) Filter to return only Management Agents having the particular Plugin installed. A special pluginName of 'None' can be provided and this will return only Management Agents having no plugin installed.
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.
* `version` - (Optional) Filter to return only Management Agents having the particular agent version.


## Attributes Reference

The following attributes are exported:

* `management_agents` - The list of management_agents.

### ManagementAgent Reference

The following attributes are exported:

* `availability_status` - The current availability status of managementAgent
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Management Agent Name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host` - Management Agent host machine name
* `host_id` - Host resource ocid
* `id` - agent identifier
* `install_key_id` - agent install key identifier
* `install_path` - Path where Management Agent is installed
* `install_type` - The install type, either AGENT or GATEWAY
* `is_agent_auto_upgradable` - true if the agent can be upgraded automatically; false if it must be upgraded manually. This flag is derived from the tenancy level auto upgrade preference.
* `is_customer_deployed` - true, if the agent image is manually downloaded and installed. false, if the agent is deployed as a plugin in Oracle Cloud Agent.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `platform_name` - Platform Name
* `platform_type` - Platform Type
* `platform_version` - Platform Version
* `plugin_list` - list of managementAgentPlugins associated with the agent
	* `is_enabled` - flag indicating whether the plugin is in enabled mode or disabled mode.
	* `plugin_display_name` - Management Agent Plugin Identifier, can be renamed
	* `plugin_id` - Plugin Id
	* `plugin_name` - Management Agent Plugin Name
	* `plugin_version` - Plugin Version
* `resource_artifact_version` - Version of the deployment artifact instantiated by this Management Agent. The format for Standalone resourceMode is YYMMDD.HHMM, and the format for other modes (whose artifacts are based upon Standalone but can advance independently) is YYMMDD.HHMM.VVVVVVVVVVVV. VVVVVVVVVVVV is always a numeric value between 000000000000 and 999999999999 
* `state` - The current state of managementAgent
* `time_created` - The time the Management Agent was created. An RFC3339 formatted datetime string
* `time_last_heartbeat` - The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
* `time_updated` - The time the Management Agent was last updated. An RFC3339 formatted datetime string
* `version` - Management Agent Version

