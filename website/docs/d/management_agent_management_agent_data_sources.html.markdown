---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_data_sources"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_data_sources"
description: |-
  Provides the list of Management Agent Data Sources in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_data_sources
This data source provides the list of Management Agent Data Sources in Oracle Cloud Infrastructure Management Agent service.

A list of Management Agent Data Sources for the given Management Agent Id.


## Example Usage

```hcl
data "oci_management_agent_management_agent_data_sources" "test_management_agent_data_sources" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id

	#Optional
	name = var.management_agent_data_source_name
}
```

## Argument Reference

The following arguments are supported:

* `management_agent_id` - (Required) Unique Management Agent identifier
* `name` - (Optional) Unique name of the dataSource.


## Attributes Reference

The following attributes are exported:

* `data_sources` - The list of data_sources.

### ManagementAgentDataSource Reference

The following attributes are exported:

* `data_source_key` - Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
* `name` - Unique name of the DataSource.
* `type` - The type of the DataSource.
