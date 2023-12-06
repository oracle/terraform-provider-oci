---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_count"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_count"
description: |-
  Provides details about a specific Management Agent Count in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_count
This data source provides details about a specific Management Agent Count resource in Oracle Cloud Infrastructure Management Agent service.

Gets count of the inventory of agents for a given compartment id, group by, and isPluginDeployed parameters.
Supported groupBy parameters: availabilityStatus, platformType, version


## Example Usage

```hcl
data "oci_management_agent_management_agent_count" "test_management_agent_count" {
	#Required
	compartment_id = var.compartment_id
	group_by = var.management_agent_count_group_by

	#Optional
	has_plugins = var.management_agent_count_has_plugins
	install_type = var.management_agent_count_install_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to which a request will be scoped.
* `group_by` - (Required) The field by which to group Management Agents. Currently, only one groupBy dimension is supported at a time.
* `has_plugins` - (Optional) When set to true then agents that have at least one plugin deployed will be returned. When set to false only agents that have no plugins deployed will be returned.
* `install_type` - (Optional) A filter to return either agents or gateway types depending upon install type selected by user. By default both install type will be returned.


## Attributes Reference

The following attributes are exported:

* `items` - List in which each item describes an aggregation of Managment Agents
	* `count` - The number of Management Agents in this group
	* `dimensions` - The Aggregation of Management Agent Dimensions
		* `availability_status` - The availability status of managementAgent
		* `has_plugins` - Whether or not a managementAgent has at least one plugin
		* `install_type` - The install type, either AGENT or GATEWAY
		* `platform_type` - Platform Type
		* `version` - Agent image version

