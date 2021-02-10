---
subcategory: "Computeinstanceagent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_agent_plugin"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_agent_plugin"
description: |-
  Provides details about a specific Instance Agent Plugin in Oracle Cloud Infrastructure Computeinstanceagent service
---

# Data Source: oci_computeinstanceagent_instance_agent_plugin
This data source provides details about a specific Instance Agent Plugin resource in Oracle Cloud Infrastructure Computeinstanceagent service.

The API to get information for a plugin.

## Example Usage

```hcl
data "oci_computeinstanceagent_instance_agent_plugin" "test_instance_agent_plugin" {
	#Required
	instanceagent_id = oci_computeinstanceagent_instanceagent.test_instanceagent.id
	plugin_name = var.instance_agent_plugin_plugin_name
}
```

## Argument Reference

The following arguments are supported:

* `instanceagent_id` - (Required) The OCID of the instance.
* `plugin_name` - (Required) The name of the plugin.


## Attributes Reference

The following attributes are exported:

* `message` - The optional message from the agent plugin
* `name` - The plugin name
* `status` - The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
* `time_last_updated_utc` - The last update time of the plugin in UTC

