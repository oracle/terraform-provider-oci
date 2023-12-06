---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_agents"
sidebar_current: "docs-oci-datasource-database_migration-agents"
description: |-
  Provides the list of Agents in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_agents
This data source provides the list of Agents in Oracle Cloud Infrastructure Database Migration service.

Display the name of all the existing ODMS Agents in the server.

## Example Usage

```hcl
data "oci_database_migration_agents" "test_agents" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.agent_display_name
	state = var.agent_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `state` - (Optional) The current state of the Database Migration Deployment. 


## Attributes Reference

The following attributes are exported:

* `agent_collection` - The list of agent_collection.

### Agent Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - ODMS Agent name 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `public_key` - ODMS Agent public key. 
* `state` - The current state of the ODMS on-premises Agent. 
* `stream_id` - The OCID of the Stream 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Agent was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time of the last Agent details update. An RFC3339 formatted datetime string. 
* `version` - ODMS Agent version 

