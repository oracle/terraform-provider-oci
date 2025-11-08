---
subcategory: "Compute Instance Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_agent_plugin"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_agent_plugin"
description: |-
  Provides details about a specific Instance Agent Plugin in Oracle Cloud Infrastructure Compute Instance Agent service
---

# Data Source: oci_computeinstanceagent_instance_agent_plugin
This data source provides details about a specific Instance Agent Plugin resource in Oracle Cloud Infrastructure Compute Instance Agent service.

Gets information about a specific Oracle Cloud Agent plugin on a compute instance.

## Example Usage

```hcl
data "oci_computeinstanceagent_instance_agent_plugin" "test_instance_agent_plugin" {
	#Required
	compartment_id = var.compartment_id
	instanceagent_id = var.instanceagent.id
	plugin_name = var.instance_agent_plugin_plugin_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which the instance resides
* `instanceagent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `plugin_name` - (Required) The name of the plugin.


## Attributes Reference

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

