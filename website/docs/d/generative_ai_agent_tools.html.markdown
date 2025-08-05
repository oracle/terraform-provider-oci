---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_tools"
sidebar_current: "docs-oci-datasource-generative_ai_agent-tools"
description: |-
  Provides the list of Tools in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_tools
This data source provides the list of Tools in Oracle Cloud Infrastructure Generative Ai Agent service.

Gets a list of tools.


## Example Usage

```hcl
data "oci_generative_ai_agent_tools" "test_tools" {

	#Optional
	agent_id = oci_generative_ai_agent_agent.test_agent.id
	compartment_id = var.compartment_id
	display_name = var.tool_display_name
	state = var.tool_state
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

* `tool_collection` - The list of tool_collection.

### Tool Reference

The following attributes are exported:

* `agent_id` - The OCID of the agent that this Tool is attached to.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the Tool.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Tool.
* `metadata` - Key-value pairs to allow additional configurations.
* `state` - The current state of the Tool.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Tool was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Tool was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `tool_config` - The configuration and type of Tool. 
	* `agent_endpoint_id` - The AgentEndpoint OCID to be used as a tool in this agent.
	* `api_schema` - The input location definition for Api schema.
		* `api_schema_input_location_type` - Type of Api Schema InputLocation. The allowed values are:
			* `INLINE`: The Api schema input location is inline.
			* `OBJECT_STORAGE_LOCATION`: The Api schema input location is object storage. 
		* `bucket` - The bucket name of an object.
		* `content` - Inline content as input.
		* `namespace` - The namespace name of an object.
		* `object` - The location/name of object.
	* `database_connection` - The connection type for Databases. 
		* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Connection.
		* `connection_type` - The type of Database connection. The allowed values are:
			* `DATABASE_TOOL_CONNECTION`: This allows the service to connect to a vector store via a Database Tools Connection. 
	* `database_schema` - The input location definition.
		* `bucket` - The bucket name of an object.
		* `content` - Inline content as input.
		* `input_location_type` - Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - The namespace name of an object.
		* `prefix` - The prefix of file object(s) or folder prefix.
	* `dialect` - Dialect to be used for SQL generation.
	* `function` - Details of Function for Function calling tool.
		* `description` - A description of the function.
		* `name` - The name of the function to invoke.
		* `parameters` - The parameters the function accepts, defined using a JSON Schema object.  Refer to the guide for examples and the JSON Schema documentation for details on the format. 
	* `generation_llm_customization` - Configuration to customize LLM. 
		* `instruction` - If specified, the default instruction is replaced with provided instruction.
	* `http_endpoint_auth_config` - Authentication configuration used for HTTP Endpoint tools. Defines the type of authentication and the source of credentials. 
		* `http_endpoint_auth_sources` - A list of credential sources from which authentication credentials can be resolved. Only AGENT is supported for HTTP Endpoint Tool. 
			* `http_endpoint_auth_scope` - Specifies the level from which credentials should be resolved.
			* `http_endpoint_auth_scope_config` - Subset of AuthScopeConfig allowed for HTTP Endpoint Tool. 
				* `client_id` - IDCS client ID.
				* `http_endpoint_auth_scope_config_type` - The type of authentication to be applied for this HTTP Endpoint. 
				* `idcs_url` - IDCS OpenID discovery endpoint.
				* `key_location` - The location of the API key in the request.
				* `key_name` - The name of the key parameter in the location.
				* `scope_url` - OAuth2 scopes for token generation.
				* `vault_secret_id` - The OCID of the vault secret with username:password. Required when `authScope` is AGENT. 
	* `icl_examples` - The input location definition.
		* `bucket` - The bucket name of an object.
		* `content` - Inline content as input.
		* `input_location_type` - Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - The namespace name of an object.
		* `prefix` - The prefix of file object(s) or folder prefix.
	* `knowledge_base_configs` - The KnowledgeBase configurations that this RAG Tool uses
		* `knowledge_base_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledgeBase this RAG Tool uses
	* `model_size` - Size of the model.
	* `should_enable_self_correction` - To enable/disable self correction.
	* `should_enable_sql_execution` - To enable/disable SQL execution.
	* `subnet_id` - The subnet ID from agent developer tenancy through which the egress is going to be routed.
	* `table_and_column_description` - The input location definition.
		* `bucket` - The bucket name of an object.
		* `content` - Inline content as input.
		* `input_location_type` - Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - The namespace name of an object.
		* `prefix` - The prefix of file object(s) or folder prefix.
	* `tool_config_type` - The type of the Tool config. The allowed values are:
		* `SQL_TOOL_CONFIG`: The config for sql Tool.
		* `RAG_TOOL_CONFIG`: The config for rag Tool.
		* FUNCTION_CALLING_TOOL_CONFIG: The config for Function calling Tool.

