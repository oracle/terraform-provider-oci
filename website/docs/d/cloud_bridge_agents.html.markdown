---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_agents"
sidebar_current: "docs-oci-datasource-cloud_bridge-agents"
description: |-
  Provides the list of Agents in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_agents
This data source provides the list of Agents in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of Agents.


## Example Usage

```hcl
data "oci_cloud_bridge_agents" "test_agents" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	agent_id = oci_cloud_bridge_agent.test_agent.id
	display_name = var.agent_display_name
	environment_id = oci_cloud_bridge_environment.test_environment.id
	state = var.agent_state
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Optional) A filter to return only resources that match the given Agent ID.
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `environment_id` - (Optional) A filter to return only resources that match the given environment ID.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `agent_collection` - The list of agent_collection.

### Agent Reference

The following attributes are exported:

* `agent_pub_key` - Resource principal public key.
* `agent_type` - Type of the Agent.
* `agent_version` - Agent identifier.
* `compartment_id` - Compartment identifier.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Agent identifier, can be renamed.
* `environment_id` - Environment identifier.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `heart_beat_status` - The current heartbeat status of the Agent based on its timeLastSyncReceived value.
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state of the Agent in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `os_version` - OS version.
* `plugin_list` - List of plugins associated with the agent.
	* `agent_id` - Agent identifier.
	* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
	* `name` - Plugin identifier, which can be renamed.
	* `plugin_version` - Plugin version.
	* `state` - The current state of the plugin.
	* `time_created` - The time when the plugin was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time when the plugin was updated. An RFC3339 formatted datetime string.
* `state` - The current state of the Agent.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the Agent was created. An RFC3339 formatted datetime string.
* `time_expire_agent_key_in_ms` - The time since epoch for when the public key will expire. An RFC3339 formatted datetime string.
* `time_last_sync_received` - The time when the last heartbeat of the Agent was noted. An RFC3339 formatted datetime string.
* `time_updated` - The time when the Agent was updated. An RFC3339 formatted datetime string.

