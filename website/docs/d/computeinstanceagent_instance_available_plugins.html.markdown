---
subcategory: "Computeinstanceagent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_computeinstanceagent_instance_available_plugins"
sidebar_current: "docs-oci-datasource-computeinstanceagent-instance_available_plugins"
description: |-
  Provides the list of Instance Available Plugins in Oracle Cloud Infrastructure Computeinstanceagent service
---

# Data Source: oci_computeinstanceagent_instance_available_plugins
This data source provides the list of Instance Available Plugins in Oracle Cloud Infrastructure Computeinstanceagent service.

The API to get the list of plugins that are available.

## Example Usage

```hcl
data "oci_computeinstanceagent_instance_available_plugins" "test_instance_available_plugins" {
	#Required
	os_name = var.instance_available_plugin_os_name
	os_version = var.instance_available_plugin_os_version

	#Optional
	name = var.instance_available_plugin_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The plugin name
* `os_name` - (Required) The OS for which the plugin is supported. Examples of OperatingSystemQueryParam:OperatingSystemVersionQueryParam are as follows: 'CentOS' '6.10' , 'CentOS Linux' '7', 'CentOS Linux' '8', 'Oracle Linux Server' '6.10', 'Oracle Linux Server' '8.0', 'Red Hat Enterprise Linux Server' '7.8', 'Windows' '10', 'Windows' '2008ServerR2', 'Windows' '2012ServerR2', 'Windows' '7', 'Windows' '8.1' 
* `os_version` - (Required) The OS version for which the plugin is supported.


## Attributes Reference

The following attributes are exported:

* `available_plugins` - The list of available_plugins.

### InstanceAvailablePlugin Reference

The following attributes are exported:

* `is_enabled_by_default` - Is the plugin enabled or disabled by default
* `is_supported` - Is the plugin supported or not
* `name` - The plugin name
* `summary` - A brief description of the plugin functionality

