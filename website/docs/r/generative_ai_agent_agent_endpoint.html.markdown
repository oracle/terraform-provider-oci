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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai-agents/latest/AgentEndpoint

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai_agent

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
	guardrail_config {

		#Optional
		content_moderation_config {

			#Optional
			input_guardrail_mode = var.agent_endpoint_guardrail_config_content_moderation_config_input_guardrail_mode
			output_guardrail_mode = var.agent_endpoint_guardrail_config_content_moderation_config_output_guardrail_mode
		}
		personally_identifiable_information_config {

			#Optional
			input_guardrail_mode = var.agent_endpoint_guardrail_config_personally_identifiable_information_config_input_guardrail_mode
			output_guardrail_mode = var.agent_endpoint_guardrail_config_personally_identifiable_information_config_output_guardrail_mode
		}
		prompt_injection_config {

			#Optional
			input_guardrail_mode = var.agent_endpoint_guardrail_config_prompt_injection_config_input_guardrail_mode
		}
	}
	human_input_config {
		#Required
		should_enable_human_input = var.agent_endpoint_human_input_config_should_enable_human_input
	}
	metadata = var.agent_endpoint_metadata
	output_config {
		#Required
		output_location {
			#Required
			bucket = var.agent_endpoint_output_config_output_location_bucket
			namespace = var.agent_endpoint_output_config_output_location_namespace
			output_location_type = var.agent_endpoint_output_config_output_location_output_location_type

			#Optional
			prefix = var.agent_endpoint_output_config_output_location_prefix
		}

		#Optional
		retention_period_in_minutes = var.agent_endpoint_output_config_retention_period_in_minutes
	}
	session_config {

		#Optional
		idle_timeout_in_seconds = var.agent_endpoint_session_config_idle_timeout_in_seconds
	}
	should_enable_citation = var.agent_endpoint_should_enable_citation
	should_enable_multi_language = var.agent_endpoint_should_enable_multi_language
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
* `guardrail_config` - (Optional) (Updatable) The configuration details about whether to apply the guardrail checks to input and output.
	* `content_moderation_config` - (Optional) (Updatable) The configuration details about whether to apply the content moderation feature to input and output. Content moderation removes toxic and biased content from responses. It is recommended to use content moderation.
		* `input_guardrail_mode` - (Optional) (Updatable) An input guardrail mode for content moderation.
		* `output_guardrail_mode` - (Optional) (Updatable) An output guardrail mode for content moderation.
	* `personally_identifiable_information_config` - (Optional) (Updatable) The configuration details for Personally Identifiable Information.
		* `input_guardrail_mode` - (Optional) (Updatable) An input guardrail mode for personally identifiable information.
		* `output_guardrail_mode` - (Optional) (Updatable) An output guardrail mode for personally identifiable information.
	* `prompt_injection_config` - (Optional) (Updatable) The configuration details for Prompt Injection.
		* `input_guardrail_mode` - (Optional) (Updatable) An input guardrail mode for prompt injection.
* `human_input_config` - (Optional) (Updatable) Human Input Configuration for an AgentEndpoint. 
	* `should_enable_human_input` - (Required) (Updatable) The Agent will request for human input for disambiguation or additional information gathering if this is enabled.
* `metadata` - (Optional) (Updatable) Key-value pairs to allow additional configurations.
* `output_config` - (Optional) (Updatable) Configuration to store results generated by agent. 
	* `output_location` - (Required) (Updatable) Location of the output. 
		* `bucket` - (Required) (Updatable) The name of the bucket.
		* `namespace` - (Required) (Updatable) The namespace of the object storage.
		* `output_location_type` - (Required) (Updatable) Type of OutputLocation.
		* `prefix` - (Optional) (Updatable) The prefix of the object storage.
	* `retention_period_in_minutes` - (Optional) (Updatable) Retention duration of the output data.
* `session_config` - (Optional) (Updatable) Session Configuration on AgentEndpoint. 
	* `idle_timeout_in_seconds` - (Optional) (Updatable) The session will become inactive after this timeout.
* `should_enable_citation` - (Optional) (Updatable) Whether to show citations in the chat result.
* `should_enable_multi_language` - (Optional) (Updatable) Whether to enable multi-language for chat.
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
* `guardrail_config` - The configuration details about whether to apply the guardrail checks to input and output.
	* `content_moderation_config` - The configuration details about whether to apply the content moderation feature to input and output. Content moderation removes toxic and biased content from responses. It is recommended to use content moderation.
		* `input_guardrail_mode` - An input guardrail mode for content moderation.
		* `output_guardrail_mode` - An output guardrail mode for content moderation.
	* `personally_identifiable_information_config` - The configuration details for Personally Identifiable Information.
		* `input_guardrail_mode` - An input guardrail mode for personally identifiable information.
		* `output_guardrail_mode` - An output guardrail mode for personally identifiable information.
	* `prompt_injection_config` - The configuration details for Prompt Injection.
		* `input_guardrail_mode` - An input guardrail mode for prompt injection.
* `human_input_config` - Human Input Configuration for an AgentEndpoint. 
	* `should_enable_human_input` - The Agent will request for human input for disambiguation or additional information gathering if this is enabled.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.
* `lifecycle_details` - A message that describes the current state of the endpoint in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `metadata` - Key-value pairs to allow additional configurations.
* `output_config` - Configuration to store results generated by agent. 
	* `output_location` - Location of the output. 
		* `bucket` - The name of the bucket.
		* `namespace` - The namespace of the object storage.
		* `output_location_type` - Type of OutputLocation.
		* `prefix` - The prefix of the object storage.
	* `retention_period_in_minutes` - Retention duration of the output data.
* `session_config` - Session Configuration on AgentEndpoint. 
	* `idle_timeout_in_seconds` - The session will become inactive after this timeout.
* `should_enable_citation` - Whether to show citations in the chat result.
* `should_enable_multi_language` - Whether to enable multi-language for chat.
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

