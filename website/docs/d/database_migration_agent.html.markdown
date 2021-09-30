---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_agent"
sidebar_current: "docs-oci-datasource-database_migration-agent"
description: |-
  Provides details about a specific Agent in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_agent
This data source provides details about a specific Agent resource in Oracle Cloud Infrastructure Database Migration service.

Display the ODMS Agent configuration.


## Example Usage

```hcl
data "oci_database_migration_agent" "test_agent" {
	#Required
	agent_id = oci_database_migration_agent.test_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The OCID of the agent 


## Attributes Reference

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

