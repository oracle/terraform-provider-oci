---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent"
sidebar_current: "docs-oci-datasource-management_agent-management_agent"
description: |-
  Provides details about a specific Management Agent in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent
This data source provides details about a specific Management Agent resource in Oracle Cloud Infrastructure Management Agent service.

Gets complete details of the inventory of a given agent id

## Example Usage

```hcl
data "oci_management_agent_management_agent" "test_management_agent" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `management_agent_id` - (Required) Unique Management Agent identifier


## Attributes Reference

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

