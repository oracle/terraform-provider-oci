---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_tool"
sidebar_current: "docs-oci-resource-generative_ai_agent-tool"
description: |-
  Provides the Tool resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_tool
This resource provides the Tool resource in Oracle Cloud Infrastructure Generative Ai Agent service.

Creates a tool.


## Example Usage

```hcl
resource "oci_generative_ai_agent_tool" "test_tool" {
	#Required
	agent_id = oci_generative_ai_agent_agent.test_agent.id
	compartment_id = var.compartment_id
	description = var.tool_description
	tool_config {
		#Required
		tool_config_type = var.tool_tool_config_tool_config_type

		#Optional
		agent_endpoint_id = oci_generative_ai_agent_agent_endpoint.test_agent_endpoint.id
		api_schema {
			#Required
			api_schema_input_location_type = var.tool_tool_config_api_schema_api_schema_input_location_type

			#Optional
			bucket = var.tool_tool_config_api_schema_bucket
			content = var.tool_tool_config_api_schema_content
			namespace = var.tool_tool_config_api_schema_namespace
			object = var.tool_tool_config_api_schema_object
		}
		database_connection {
			#Required
			connection_id = oci_database_migration_connection.test_connection.id
			connection_type = var.tool_tool_config_database_connection_connection_type
		}
		database_schema {
			#Required
			input_location_type = var.tool_tool_config_database_schema_input_location_type

			#Optional
			bucket = var.tool_tool_config_database_schema_bucket
			content = var.tool_tool_config_database_schema_content
			namespace = var.tool_tool_config_database_schema_namespace
			prefix = var.tool_tool_config_database_schema_prefix
		}
		dialect = var.tool_tool_config_dialect
		function {

			#Optional
			description = var.tool_tool_config_function_description
			name = var.tool_tool_config_function_name
			parameters = var.tool_tool_config_function_parameters
		}
		generation_llm_customization {

			#Optional
			instruction = var.tool_tool_config_generation_llm_customization_instruction
		}
		http_endpoint_auth_config {

			#Optional
			http_endpoint_auth_sources {

				#Optional
				http_endpoint_auth_scope = var.tool_tool_config_http_endpoint_auth_config_http_endpoint_auth_sources_http_endpoint_auth_scope
				http_endpoint_auth_scope_config {
					#Required
					http_endpoint_auth_scope_config_type = var.tool_tool_config_http_endpoint_auth_config_http_endpoint_auth_sources_http_endpoint_auth_scope_config_http_endpoint_auth_scope_config_type

					#Optional
					client_id = oci_generative_ai_agent_client.test_client.id
					idcs_url = var.tool_tool_config_http_endpoint_auth_config_http_endpoint_auth_sources_http_endpoint_auth_scope_config_idcs_url
					key_location = var.tool_tool_config_http_endpoint_auth_config_http_endpoint_auth_sources_http_endpoint_auth_scope_config_key_location
					key_name = oci_kms_key.test_key.name
					scope_url = var.tool_tool_config_http_endpoint_auth_config_http_endpoint_auth_sources_http_endpoint_auth_scope_config_scope_url
					vault_secret_id = oci_vault_secret.test_secret.id
				}
			}
		}
		icl_examples {
			#Required
			input_location_type = var.tool_tool_config_icl_examples_input_location_type

			#Optional
			bucket = var.tool_tool_config_icl_examples_bucket
			content = var.tool_tool_config_icl_examples_content
			namespace = var.tool_tool_config_icl_examples_namespace
			prefix = var.tool_tool_config_icl_examples_prefix
		}
		knowledge_base_configs {

			#Optional
			knowledge_base_id = oci_generative_ai_agent_knowledge_base.test_knowledge_base.id
		}
		model_size = var.tool_tool_config_model_size
		should_enable_self_correction = var.tool_tool_config_should_enable_self_correction
		should_enable_sql_execution = var.tool_tool_config_should_enable_sql_execution
		subnet_id = oci_core_subnet.test_subnet.id
		table_and_column_description {
			#Required
			input_location_type = var.tool_tool_config_table_and_column_description_input_location_type

			#Optional
			bucket = var.tool_tool_config_table_and_column_description_bucket
			content = var.tool_tool_config_table_and_column_description_content
			namespace = var.tool_tool_config_table_and_column_description_namespace
			prefix = var.tool_tool_config_table_and_column_description_prefix
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.tool_display_name
	freeform_tags = {"Department"= "Finance"}
	metadata = var.tool_metadata
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The OCID of the agent that this Tool is attached to.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) Description about the Tool.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `metadata` - (Optional) (Updatable) Key-value pairs to allow additional configurations.
* `tool_config` - (Required) (Updatable) The configuration and type of Tool. 
	* `agent_endpoint_id` - (Required when tool_config_type=AGENT_TOOL_CONFIG) (Updatable) The AgentEndpoint OCID to be used as a tool in this agent.
	* `api_schema` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) The input location definition for Api schema.
		* `api_schema_input_location_type` - (Required) (Updatable) Type of Api Schema InputLocation. The allowed values are:
			* `INLINE`: The Api schema input location is inline.
			* `OBJECT_STORAGE_LOCATION`: The Api schema input location is object storage. 
		* `bucket` - (Required when api_schema_input_location_type=OBJECT_STORAGE_LOCATION) (Updatable) The bucket name of an object.
		* `content` - (Required when api_schema_input_location_type=INLINE) (Updatable) Inline content as input.
		* `namespace` - (Required when api_schema_input_location_type=OBJECT_STORAGE_LOCATION) (Updatable) The namespace name of an object.
		* `object` - (Required when api_schema_input_location_type=OBJECT_STORAGE_LOCATION) (Updatable) The location/name of object.
	* `database_connection` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) The connection type for Databases. 
		* `connection_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Connection.
		* `connection_type` - (Required) (Updatable) The type of Database connection. The allowed values are:
			* `DATABASE_TOOL_CONNECTION`: This allows the service to connect to a vector store via a Database Tools Connection. 
	* `database_schema` - (Required when tool_config_type=SQL_TOOL_CONFIG) (Updatable) The input location definition.
		* `bucket` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The bucket name of an object.
		* `content` - (Applicable when input_location_type=INLINE) (Updatable) Inline content as input.
		* `input_location_type` - (Required) (Updatable) Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The namespace name of an object.
		* `prefix` - (Applicable when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The prefix of file object(s) or folder prefix.
	* `dialect` - (Required when tool_config_type=SQL_TOOL_CONFIG) (Updatable) Dialect to be used for SQL generation.
	* `function` - (Required when tool_config_type=FUNCTION_CALLING_TOOL_CONFIG) (Updatable) Details of Function for Function calling tool.
		* `description` - (Applicable when tool_config_type=FUNCTION_CALLING_TOOL_CONFIG) (Updatable) A description of the function.
		* `name` - (Required when tool_config_type=FUNCTION_CALLING_TOOL_CONFIG) (Updatable) The name of the function to invoke.
		* `parameters` - (Applicable when tool_config_type=FUNCTION_CALLING_TOOL_CONFIG) (Updatable) The parameters the function accepts, defined using a JSON Schema object.  Refer to the guide for examples and the JSON Schema documentation for details on the format. 
	* `generation_llm_customization` - (Applicable when tool_config_type=RAG_TOOL_CONFIG | SQL_TOOL_CONFIG) (Updatable) Configuration to customize LLM. 
		* `instruction` - (Applicable when tool_config_type=RAG_TOOL_CONFIG | SQL_TOOL_CONFIG) (Updatable) If specified, the default instruction is replaced with provided instruction.
	* `http_endpoint_auth_config` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) Authentication configuration used for HTTP Endpoint tools. Defines the type of authentication and the source of credentials. 
		* `http_endpoint_auth_sources` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) A list of credential sources from which authentication credentials can be resolved. Only AGENT is supported for HTTP Endpoint Tool. 
			* `http_endpoint_auth_scope` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) Specifies the level from which credentials should be resolved.
			* `http_endpoint_auth_scope_config` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) Subset of AuthScopeConfig allowed for HTTP Endpoint Tool. 
				* `client_id` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG) (Updatable) IDCS client ID.
				* `http_endpoint_auth_scope_config_type` - (Required) (Updatable) The type of authentication to be applied for this HTTP Endpoint. 
				* `idcs_url` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG) (Updatable) IDCS OpenID discovery endpoint.
				* `key_location` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG) (Updatable) The location of the API key in the request.
				* `key_name` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG) (Updatable) The name of the key parameter in the location.
				* `scope_url` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG) (Updatable) OAuth2 scopes for token generation.
				* `vault_secret_id` - (Required when http_endpoint_auth_scope_config_type=HTTP_ENDPOINT_API_KEY_AUTH_SCOPE_CONFIG | HTTP_ENDPOINT_BASIC_AUTH_SCOPE_CONFIG | HTTP_ENDPOINT_BEARER_AUTH_SCOPE_CONFIG | HTTP_ENDPOINT_IDCS_AUTH_SCOPE_CONFIG) (Updatable) The OCID of the vault secret with username:password. Required when `authScope` is AGENT. 
	* `icl_examples` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) The input location definition.
		* `bucket` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The bucket name of an object.
		* `content` - (Applicable when input_location_type=INLINE) (Updatable) Inline content as input.
		* `input_location_type` - (Required) (Updatable) Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The namespace name of an object.
		* `prefix` - (Applicable when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The prefix of file object(s) or folder prefix.
	* `knowledge_base_configs` - (Required when tool_config_type=RAG_TOOL_CONFIG) (Updatable) The KnowledgeBase configurations that this RAG Tool uses
		* `knowledge_base_id` - (Required when tool_config_type=RAG_TOOL_CONFIG) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledgeBase this RAG Tool uses
	* `model_size` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) Size of the model.
	* `should_enable_self_correction` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) To enable/disable self correction.
	* `should_enable_sql_execution` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) To enable/disable SQL execution.
	* `subnet_id` - (Required when tool_config_type=HTTP_ENDPOINT_TOOL_CONFIG) (Updatable) The subnet ID from agent developer tenancy through which the egress is going to be routed.
	* `table_and_column_description` - (Applicable when tool_config_type=SQL_TOOL_CONFIG) (Updatable) The input location definition.
		* `bucket` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The bucket name of an object.
		* `content` - (Applicable when input_location_type=INLINE) (Updatable) Inline content as input.
		* `input_location_type` - (Required) (Updatable) Type of InputLocation. The allowed values are:
			* `INLINE`: The input location is inline.
			* `OBJECT_STORAGE_PREFIX`: The input location is object storage. 
		* `namespace` - (Required when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The namespace name of an object.
		* `prefix` - (Applicable when input_location_type=OBJECT_STORAGE_PREFIX) (Updatable) The prefix of file object(s) or folder prefix.
	* `tool_config_type` - (Required) (Updatable) The type of the Tool config. The allowed values are:
		* `SQL_TOOL_CONFIG`: The config for sql Tool.
		* `RAG_TOOL_CONFIG`: The config for rag Tool.
		* FUNCTION_CALLING_TOOL_CONFIG: The config for Function calling Tool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Tool
	* `update` - (Defaults to 20 minutes), when updating the Tool
	* `delete` - (Defaults to 20 minutes), when destroying the Tool


## Import

Tools can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_tool.test_tool "id"
```

