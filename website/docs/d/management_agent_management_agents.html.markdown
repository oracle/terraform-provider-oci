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

Returns a list of Management Agents.
If no explicit page size limit is specified, it will default to 1000 when compartmentIdInSubtree is true and 5000 otherwise.
The response is limited to maximum 1000 records when compartmentIdInSubtree is true.


## Example Usage

```hcl
data "oci_management_agent_management_agents" "test_management_agents" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.management_agent_access_level
	availability_status = var.management_agent_availability_status
	compartment_id_in_subtree = var.management_agent_compartment_id_in_subtree
	data_source_name = oci_management_agent_management_agent_data_source.test_management_agent_data_source.name
	data_source_type = var.management_agent_data_source_type
	display_name = var.management_agent_display_name
	gateway_id = oci_apigateway_gateway.test_gateway.id
	host_id = oci_management_agent_host.test_host.id
	wait_for_host_id = 10
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

* `access_level` - (Optional) When the value is "ACCESSIBLE", insufficient permissions for a compartment will filter out resources in that compartment without rejecting the request. 
* `availability_status` - (Optional) Filter to return only Management Agents in the particular availability status.
* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `compartment_id_in_subtree` - (Optional) if set to true then it fetches resources for all compartments where user has access to else only on the compartment specified.
* `data_source_name` - (Optional) Unique name of the dataSource.
* `data_source_type` - (Optional) The type of the dataSource.
* `display_name` - (Optional) Filter to return only Management Agents having the particular display name.
* `gateway_id` - (Optional) Filter to return only results having the particular gatewayId.
* `host_id` - (Optional) Filter to return only Management Agents having the particular agent host id.
* `wait_for_host_id` - (Optional) When host_id argument is set, the data source will wait for the given period of time (in minutes) for this host_id to become available. This can be used when compute instance with Management Agent has been recently created.
* `install_type` - (Optional) A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.
* `is_customer_deployed` - (Optional) true, if the agent image is manually downloaded and installed. false, if the agent is deployed as a plugin in Oracle Cloud Agent.
* `platform_type` - (Optional) Array of PlatformTypes to return only results having the particular platform types. Example: ["LINUX"]
* `plugin_name` - (Optional) Array of pluginName to return only Management Agents having the particular Plugins installed. A special pluginName of 'None' can be provided and this will return only Management Agents having no plugin installed. Example: ["PluginA"]
* `state` - (Optional) Filter to return only Management Agents in the particular lifecycle state.
* `version` - (Optional) Array of versions to return only Management Agents having the particular agent versions. Example: ["202020.0101","210201.0513"]


## Attributes Reference

The following attributes are exported:

* `management_agents` - The list of management_agents.

### ManagementAgent Reference

The following attributes are exported:

* `availability_status` - The current availability status of managementAgent
* `compartment_id` - Compartment Identifier
* `data_source_summary_list` - list of dataSources associated with the agent
    * `is_daemon_set` - If the Kubernetes cluster type is Daemon set then this will be set to true.
	* `key` - Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
	* `name` - Unique name of the DataSource.
	* `type` - The type of the DataSource.
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
* `management_agent_properties` - Additional properties for this Management Agent
	* `name` - Name of the property
	* `units` - Unit for the property
	* `values` - Values of the property
* `platform_name` - Platform Name
* `platform_type` - Platform Type
* `platform_version` - Platform Version
* `plugin_list` - list of managementAgentPlugins associated with the agent
	* `is_enabled` - flag indicating whether the plugin is in enabled mode or disabled mode.
	* `plugin_display_name` - Management Agent Plugin Identifier, can be renamed
	* `plugin_id` - Plugin Id
	* `plugin_name` - Management Agent Plugin Name
	* `plugin_status` - Plugin Status
	* `plugin_status_message` - Status message of the Plugin
	* `plugin_version` - Plugin Version
* `resource_artifact_version` - Version of the deployment artifact instantiated by this Management Agent. The format for Standalone resourceMode is YYMMDD.HHMM, and the format for other modes (whose artifacts are based upon Standalone but can advance independently) is YYMMDD.HHMM.VVVVVVVVVVVV. VVVVVVVVVVVV is always a numeric value between 000000000000 and 999999999999 
* `state` - The current state of managementAgent
* `time_created` - The time the Management Agent was created. An RFC3339 formatted datetime string
* `time_last_heartbeat` - The time the Management Agent has last recorded its health status in telemetry. This value will be null if the agent has not recorded its health status in last 7 days. An RFC3339 formatted datetime string
* `time_updated` - The time the Management Agent was last updated. An RFC3339 formatted datetime string
* `version` - Management Agent Version

