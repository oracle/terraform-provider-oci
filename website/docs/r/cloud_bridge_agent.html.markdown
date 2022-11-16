---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_agent"
sidebar_current: "docs-oci-resource-cloud_bridge-agent"
description: |-
  Provides the Agent resource in Oracle Cloud Infrastructure Cloud Bridge service
---

# oci_cloud_bridge_agent
This resource provides the Agent resource in Oracle Cloud Infrastructure Cloud Bridge service.

Creates an Agent.


## Example Usage

```hcl
resource "oci_cloud_bridge_agent" "test_agent" {
	#Required
	agent_type = var.agent_agent_type
	agent_version = var.agent_agent_version
	compartment_id = var.compartment_id
	display_name = var.agent_display_name
	environment_id = oci_cloud_bridge_environment.test_environment.id
	os_version = var.agent_os_version

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `agent_type` - (Required) Agent identifier.
* `agent_version` - (Required) Agent identifier.
* `compartment_id` - (Required) (Updatable) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Agent identifier.
* `environment_id` - (Required) Environment identifier.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `os_version` - (Required) OS version.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Agent
	* `update` - (Defaults to 20 minutes), when updating the Agent
	* `delete` - (Defaults to 20 minutes), when destroying the Agent


## Import

Agents can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_bridge_agent.test_agent "id"
```

