---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_plugin_count"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_plugin_count"
description: |-
  Provides details about a specific Management Agent Plugin Count in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_plugin_count
This data source provides details about a specific Management Agent Plugin Count resource in Oracle Cloud Infrastructure Management Agent service.

Gets count of the inventory of management agent plugins for a given compartment id and group by parameter.
Supported groupBy parameter: pluginName


## Example Usage

```hcl
data "oci_management_agent_management_agent_plugin_count" "test_management_agent_plugin_count" {
	#Required
	compartment_id = var.compartment_id
	group_by = var.management_agent_plugin_count_group_by
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `group_by` - (Required) The field by which to group Management Agent Plugins


## Attributes Reference

The following attributes are exported:

* `items` - List in which each item describes an aggregation of Managment Agent Plugins
	* `count` - The number of Management Agent Plugins in this group
	* `dimensions` - The Aggregation of Management Agent Plugin Dimensions
		* `plugin_display_name` - Management Agent Plugin Display Name
		* `plugin_name` - Management Agent Plugin Name

