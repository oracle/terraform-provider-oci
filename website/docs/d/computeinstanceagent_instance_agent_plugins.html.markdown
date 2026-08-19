---
subcategory: "Compute Instance Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_agent_plugins"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_agent_plugins"
description: |-
  Provides the list of Instance Agent Plugins in Oracle Cloud Infrastructure Compute Instance Agent service
---

# Data Source: oci_computeinstanceagent_instance_agent_plugins
This data source provides the list of Instance Agent Plugins in Oracle Cloud Infrastructure Compute Instance Agent service.

Gets information about the Oracle Cloud Agent plugins that are available on a specific compute instance.


## Example Usage

```hcl
data "oci_computeinstanceagent_instance_agent_plugins" "test_instance_agent_plugins" {
	#Required
	compartment_id = var.compartment_id
	instanceagent_id = var.instanceagent.id

	#Optional
	name = var.instance_agent_plugin_name
	status = var.instance_agent_plugin_status
}
```

## Argument Reference

The following arguments are supported:

* `instanceagent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `compartment_id` - (Required) The ID of the compartment in which the instance resides
* `name` - (Optional) The plugin name.
* `status` - (Optional) The plugin status.


## Attributes Reference

The following attributes are exported:

* `instance_agent_plugins` - The list of instance_agent_plugins.

### InstanceAgentPlugin Reference

The following attributes are exported:

* `message` - An optional message from the plugin.
* `name` - The plugin name.
* `status` - The plugin status.

	These are the available statuses:
	* `RUNNING` - The plugin is running.
	* `STOPPED` - The plugin is stopped.
	* `NOT_SUPPORTED` - The plugin is not supported on this platform.
	* `INVALID` - The plugin status is not recognizable by the service.

	To determine whether the plugin is enabled, use the [GetInstance](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/GetInstance) operation in the Core Services API. To enable or disable the plugin, use the [UpdateInstance](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/UpdateInstance) operation in the Core Services API. 
* `time_last_updated_utc` - The last updated time of the plugin, in UTC.

