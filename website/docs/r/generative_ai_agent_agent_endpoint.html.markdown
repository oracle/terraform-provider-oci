---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_agent_endpoint"
sidebar_current: "docs-oci-resource-generative_ai_agent-agent_endpoint"
description: |-
  Provides the Agent Endpoint resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_agent_endpoint
This resource provides the Agent Endpoint resource in Oracle Cloud Infrastructure Generative Ai Agent service.

**CreateAgentEndpoint**

Creates an endpoint.


## Example Usage

```hcl
resource "oci_generative_ai_agent_agent_endpoint" "test_agent_endpoint" {
	#Required
	agent_id = oci_generative_ai_agent_agent.test_agent.id
	compartment_id = var.compartment_id

	#Optional
	content_moderation_config {

		#Optional
		should_enable_on_input = var.agent_endpoint_content_moderation_config_should_enable_on_input
		should_enable_on_output = var.agent_endpoint_content_moderation_config_should_enable_on_output
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.agent_endpoint_description
	display_name = var.agent_endpoint_display_name
	freeform_tags = {"Department"= "Finance"}
	session_config {

		#Optional
		idle_timeout_in_seconds = var.agent_endpoint_session_config_idle_timeout_in_seconds
	}
	should_enable_citation = var.agent_endpoint_should_enable_citation
	should_enable_session = var.agent_endpoint_should_enable_session
	should_enable_trace = var.agent_endpoint_should_enable_trace
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The OCID of the agent that this endpoint is associated with.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the endpoint in. 
* `content_moderation_config` - (Optional) (Updatable) The configuration details about whether to apply the content moderation feature to input and output. Content moderation removes toxic and biased content from responses. It is recommended to use content moderation.
	* `should_enable_on_input` - (Optional) (Updatable) A flag to enable or disable content moderation on input.
	* `should_enable_on_output` - (Optional) (Updatable) A flag to enable or disable content moderation on output.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the endpoint.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `session_config` - (Optional) (Updatable) **SessionConfig**

	Session Configuration on AgentEndpoint. 
	* `idle_timeout_in_seconds` - (Optional) (Updatable) The session will become inactive after this timeout.
* `should_enable_citation` - (Optional) (Updatable) Whether to show citations in the chat result.
* `should_enable_session` - (Optional) Whether or not to enable Session-based chat.
* `should_enable_trace` - (Optional) (Updatable) Whether to show traces in the chat result.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agent_id` - The OCID of the agent that this endpoint is associated with.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `content_moderation_config` - The configuration details about whether to apply the content moderation feature to input and output. Content moderation removes toxic and biased content from responses. It is recommended to use content moderation.
	* `should_enable_on_input` - A flag to enable or disable content moderation on input.
	* `should_enable_on_output` - A flag to enable or disable content moderation on output.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.
* `lifecycle_details` - A message that describes the current state of the endpoint in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `session_config` - **SessionConfig**

	Session Configuration on AgentEndpoint. 
	* `idle_timeout_in_seconds` - The session will become inactive after this timeout.
* `should_enable_citation` - Whether to show citations in the chat result.
* `should_enable_session` - Whether or not to enable Session-based chat.
* `should_enable_trace` - Whether to show traces in the chat result.
* `state` - The current state of the endpoint.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the AgentEndpoint was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the endpoint was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Agent Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Agent Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Agent Endpoint


## Import

AgentEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_agent_endpoint.test_agent_endpoint "id"
```

