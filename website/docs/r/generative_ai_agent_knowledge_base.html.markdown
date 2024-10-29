---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_knowledge_base"
sidebar_current: "docs-oci-resource-generative_ai_agent-knowledge_base"
description: |-
  Provides the Knowledge Base resource in Oracle Cloud Infrastructure Generative Ai Agent service
---

# oci_generative_ai_agent_knowledge_base
This resource provides the Knowledge Base resource in Oracle Cloud Infrastructure Generative Ai Agent service.

**CreateKnowledgeBase**

Creates a knowledge base.


## Example Usage

```hcl
resource "oci_generative_ai_agent_knowledge_base" "test_knowledge_base" {
	#Required
	compartment_id = var.compartment_id
	index_config {
		#Required
		index_config_type = var.knowledge_base_index_config_index_config_type

		#Optional
		cluster_id = oci_containerengine_cluster.test_cluster.id
		database_connection {
			#Required
			connection_id = oci_database_migration_connection.test_connection.id
			connection_type = var.knowledge_base_index_config_database_connection_connection_type
		}
		database_functions {

			#Optional
			name = var.knowledge_base_index_config_database_functions_name
		}
		indexes {

			#Optional
			name = var.knowledge_base_index_config_indexes_name
			schema {

				#Optional
				body_key = var.knowledge_base_index_config_indexes_schema_body_key
				embedding_body_key = var.knowledge_base_index_config_indexes_schema_embedding_body_key
				title_key = var.knowledge_base_index_config_indexes_schema_title_key
				url_key = var.knowledge_base_index_config_indexes_schema_url_key
			}
		}
		secret_detail {
			#Required
			type = var.knowledge_base_index_config_secret_detail_type
			vault_secret_id = oci_vault_secret.test_secret.id

			#Optional
			client_id = oci_generative_ai_agent_client.test_client.id
			idcs_url = var.knowledge_base_index_config_secret_detail_idcs_url
			scope_url = var.knowledge_base_index_config_secret_detail_scope_url
		}
		should_enable_hybrid_search = var.knowledge_base_index_config_should_enable_hybrid_search
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.knowledge_base_description
	display_name = var.knowledge_base_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the knowledge base in. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user-friendly description of the knowledge base.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `index_config` - (Required) (Updatable) **IndexConfig**

	The index configuration of Knowledge bases. 
	* `cluster_id` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OpenSearch Cluster.
	* `database_connection` - (Required when index_config_type=OCI_DATABASE_CONFIG) (Updatable) **DatabaseConnection**

		The connection type for Databases. 
		* `connection_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Connection.
		* `connection_type` - (Required) (Updatable) The type of Database connection. The allowed values are:
			* `DATABASE_TOOL_CONNECTION`: This allows the service to connect to a vector store via a Database Tools Connection. 
	* `database_functions` - (Required when index_config_type=OCI_DATABASE_CONFIG) (Updatable) Array of Database functions to be used.
		* `name` - (Required when index_config_type=OCI_DATABASE_CONFIG) (Updatable) The name of the Database function. 
	* `index_config_type` - (Required) (Updatable) The type of index. The allowed values are:
		* `DEFAULT_INDEX_CONFIG`: DefaultIndexConfig allows the service to create and manage vector store on behalf of the customer.
		* `OCI_OPEN_SEARCH_INDEX_CONFIG`: OciOpenSearchIndexConfig allows customer to configure their OpenSearch cluster.
		* `OCI_DATABASE_CONFIG`: OciDatabaseConfig allows customer to configure their Database. 
	* `indexes` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) Index configuration for open search.
		* `name` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) The index name in opensearch.
		* `schema` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) **IndexSchema**

			The index schema details. 
			* `body_key` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) Body key name.
			* `embedding_body_key` - (Applicable when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) Field within customer managed Oracle Cloud Infrastructure OpenSearch document containing the vector embedding for queries.
			* `title_key` - (Applicable when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) Title key that stores the Title of a document, if available.
			* `url_key` - (Applicable when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) URL key that stores the URL of a document, if available.
	* `secret_detail` - (Required when index_config_type=OCI_OPEN_SEARCH_INDEX_CONFIG) (Updatable) **SecretDetail**

		The details of configured security configuration on OpenSearch. 
		* `client_id` - (Required when type=IDCS_SECRET) (Updatable) The IDCS Connect clientId.
		* `idcs_url` - (Required when type=IDCS_SECRET) (Updatable) The URL represent authentication url of the IDCS.
		* `scope_url` - (Required when type=IDCS_SECRET) (Updatable) Fully qualified scope url
		* `type` - (Required) (Updatable) The type of OpenID. The allowed values are:
			* `IDCS_SECRET`: The OpenID configuration used is OpenSearch is IDCS.
			* `BASIC_AUTH_SECRET`: Basic authentication use for OpenSearch 
		* `vault_secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret for basic authentication.
	* `should_enable_hybrid_search` - (Applicable when index_config_type=DEFAULT_INDEX_CONFIG) (Updatable) Whether to enable Hybrid search in service managed OpenSearch.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the knowledge base.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledge base.
* `index_config` - **IndexConfig**

	The index configuration of Knowledge bases. 
	* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OpenSearch Cluster.
	* `database_connection` - **DatabaseConnection**

		The connection type for Databases. 
		* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Connection.
		* `connection_type` - The type of Database connection. The allowed values are:
			* `DATABASE_TOOL_CONNECTION`: This allows the service to connect to a vector store via a Database Tools Connection. 
	* `database_functions` - Array of Database functions to be used.
		* `name` - The name of the Database function. 
	* `index_config_type` - The type of index. The allowed values are:
		* `DEFAULT_INDEX_CONFIG`: DefaultIndexConfig allows the service to create and manage vector store on behalf of the customer.
		* `OCI_OPEN_SEARCH_INDEX_CONFIG`: OciOpenSearchIndexConfig allows customer to configure their OpenSearch cluster.
		* `OCI_DATABASE_CONFIG`: OciDatabaseConfig allows customer to configure their Database. 
	* `indexes` - Index configuration for open search.
		* `name` - The index name in opensearch.
		* `schema` - **IndexSchema**

			The index schema details. 
			* `body_key` - Body key name.
			* `embedding_body_key` - Field within customer managed Oracle Cloud Infrastructure OpenSearch document containing the vector embedding for queries.
			* `title_key` - Title key that stores the Title of a document, if available.
			* `url_key` - URL key that stores the URL of a document, if available.
	* `secret_detail` - **SecretDetail**

		The details of configured security configuration on OpenSearch. 
		* `client_id` - The IDCS Connect clientId.
		* `idcs_url` - The URL represent authentication url of the IDCS.
		* `scope_url` - Fully qualified scope url
		* `type` - The type of OpenID. The allowed values are:
			* `IDCS_SECRET`: The OpenID configuration used is OpenSearch is IDCS.
			* `BASIC_AUTH_SECRET`: Basic authentication use for OpenSearch 
		* `vault_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret for basic authentication.
	* `should_enable_hybrid_search` - Whether to enable Hybrid search in service managed OpenSearch.
* `lifecycle_details` - A message that describes the current state of the knowledge base in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the knowledge base.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the knowledge base was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the knowledge base was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Knowledge Base
	* `update` - (Defaults to 20 minutes), when updating the Knowledge Base
	* `delete` - (Defaults to 20 minutes), when destroying the Knowledge Base


## Import

KnowledgeBases can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_agent_knowledge_base.test_knowledge_base "id"
```

