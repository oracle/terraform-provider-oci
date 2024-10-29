---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_agent_endpoints"
sidebar_current: "docs-oci-datasource-generative_ai_agent-agent_endpoints"
description: |-
  Provides the list of Agent Endpoints in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_agent_endpoints
This data source provides the list of Agent Endpoints in Oracle Cloud Infrastructure Generative Ai Agent service.

**ListAgentEndpoints**

Gets a list of endpoints.


## Example Usage

```hcl
data "oci_generative_ai_agent_agent_endpoints" "test_agent_endpoints" {

	#Optional
	agent_id = oci_generative_ai_agent_agent.test_agent.id
	compartment_id = var.compartment_id
	display_name = var.agent_endpoint_display_name
	state = var.agent_endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `agent_endpoint_collection` - The list of agent_endpoint_collection.

### AgentEndpoint Reference

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

