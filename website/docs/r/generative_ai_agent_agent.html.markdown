---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_agent"
sidebar_current: "docs-oci-resource-generative_ai_agent-agent"
description: |-
  Provides the Agent resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_agent
This resource provides the Agent resource in Oracle Cloud Infrastructure Generative Ai Agent service.

**CreateAgent**

Creates an agent.


## Example Usage

```hcl
resource "oci_generative_ai_agent_agent" "test_agent" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.agent_description
	display_name = var.agent_display_name
	freeform_tags = {"Department"= "Finance"}
	knowledge_base_ids = var.agent_knowledge_base_ids
	welcome_message = var.agent_welcome_message
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the agent in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description about the agent.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `knowledge_base_ids` - (Optional) (Updatable) List of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledgeBases associated with agent.
* `welcome_message` - (Optional) (Updatable) Details about purpose and responsibility of the agent


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description about the agent.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent.
* `knowledge_base_ids` - List of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledgeBases associated with agent.
* `lifecycle_details` - A message that describes the current state of the agent in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the agent.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the agent was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the agent was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `welcome_message` - Details about purpose and responsibility of the agent

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Agent
	* `update` - (Defaults to 20 minutes), when updating the Agent
	* `delete` - (Defaults to 20 minutes), when destroying the Agent


## Import

Agents can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_agent.test_agent "id"
```

