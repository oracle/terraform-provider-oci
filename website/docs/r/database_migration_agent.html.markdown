---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_agent"
sidebar_current: "docs-oci-resource-database_migration-agent"
description: |-
  Provides the Agent resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_agent
This resource provides the Agent resource in Oracle Cloud Infrastructure Database Migration service.

Modifies the ODMS Agent represented by the given ODMS Agent ID.

## Example Usage

```hcl
resource "oci_database_migration_agent" "test_agent" {
	#Required
	agent_id = oci_database_migration_agent.test_agent.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.agent_display_name
	freeform_tags = {"bar-key"= "value"}
	public_key = var.agent_public_key
	stream_id = oci_streaming_stream.test_stream.id
	version = var.agent_version
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The OCID of the agent 
* `compartment_id` - (Optional) (Updatable) OCID of the compartment 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) ODMS Agent name 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `public_key` - (Optional) (Updatable) ODMS Agent public key. 
* `stream_id` - (Optional) (Updatable) The OCID of the Stream 
* `version` - (Optional) (Updatable) ODMS Agent version 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

Agents can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_agent.test_agent "id"
```

