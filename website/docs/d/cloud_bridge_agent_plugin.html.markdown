---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_agent_plugin"
sidebar_current: "docs-oci-datasource-cloud_bridge-agent_plugin"
description: |-
  Provides details about a specific Agent Plugin in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_agent_plugin
This data source provides details about a specific Agent Plugin resource in Oracle Cloud Infrastructure Cloud Bridge service.

Gets a plugin by identifier.

## Example Usage

```hcl
data "oci_cloud_bridge_agent_plugin" "test_agent_plugin" {
	#Required
	agent_id = oci_cloud_bridge_agent.test_agent.id
	plugin_name = var.agent_plugin_plugin_name
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) Unique Agent identifier path parameter.
* `plugin_name` - (Required) Unique plugin identifier path parameter.


## Attributes Reference

The following attributes are exported:

* `agent_id` - Agent identifier.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `desired_state` - State to which the customer wants the plugin to move to.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `name` - Plugin identifier, which can be renamed.
* `plugin_version` - Plugin version.
* `state` - The current state of the plugin.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the Agent was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the Agent was updated. An RFC3339 formatted datetime string.

