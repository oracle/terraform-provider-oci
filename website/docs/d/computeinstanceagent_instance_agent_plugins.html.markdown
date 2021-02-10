---
subcategory: "Computeinstanceagent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_agent_plugins"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_agent_plugins"
description: |-
  Provides the list of Instance Agent Plugins in Oracle Cloud Infrastructure Computeinstanceagent service
---

# Data Source: oci_computeinstanceagent_instance_agent_plugins
This data source provides the list of Instance Agent Plugins in Oracle Cloud Infrastructure Computeinstanceagent service.

The API to get one or more plugin information.

## Example Usage

```hcl
data "oci_computeinstanceagent_instance_agent_plugins" "test_instance_agent_plugins" {
	#Required
	instanceagent_id = oci_computeinstanceagent_instanceagent.test_instanceagent.id

	#Optional
	name = var.instance_agent_plugin_name
	status = var.instance_agent_plugin_status
}
```

## Argument Reference

The following arguments are supported:

* `instanceagent_id` - (Required) The OCID of the instance.
* `name` - (Optional) The plugin name
* `status` - (Optional) The plugin status


## Attributes Reference

The following attributes are exported:

* `instance_agent_plugins` - The list of instance_agent_plugins.

### InstanceAgentPlugin Reference

The following attributes are exported:

* `message` - The optional message from the agent plugin
* `name` - The plugin name
* `status` - The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
* `time_last_updated_utc` - The last update time of the plugin in UTC

