---
subcategory: "Compute Instance Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_available_plugins"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_available_plugins"
description: |-
  Provides the list of Instance Available Plugins in Oracle Cloud Infrastructure Compute Instance Agent service
---

# Data Source: oci_computeinstanceagent_instance_available_plugins
This data source provides the list of Instance Available Plugins in Oracle Cloud Infrastructure Compute Instance Agent service.

Lists the Oracle Cloud Agent plugins that are available for compute instances in a specific compartment.

## Example Usage

```hcl
data "oci_computeinstanceagent_instance_available_plugins" "test_instance_available_plugins" {
	#Required
	compartment_id = var.compartment_id
	os_name = var.instance_available_plugin_os_name
	os_version = var.instance_available_plugin_os_version

	#Optional
	name = var.instance_available_plugin_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment for which the plugins are available
* `name` - (Optional) The plugin name.
* `os_name` - (Required) The image (OS) for the compute instance.

	If no match is found, all plugins are returned.

	Examples: `CentOS`, `Oracle Linux`, `Oracle Autonomous Linux`, `Canonical Ubuntu`, `Windows Server` 
* `os_version` - (Required) The OS version for the instance.

	If no match is found, all plugins are returned.

	Examples: `9.6`, `8` for CentOS and Oracle Linux. `22.04`, `22.04 Minimal` for Canonical Ubuntu. `2012 R2 Datacenter`, `2019 Standard` for Windows Server. 


## Attributes Reference

The following attributes are exported:

* `available_plugins` - The list of available_plugins.

### InstanceAvailablePlugin Reference

The following attributes are exported:

* `is_enabled_by_default` - Whether the plugin is enabled or disabled by default.
* `is_supported` - Whether the plugin is supported.
* `name` - The plugin name.
* `summary` - A brief description of the plugin's functionality.

