---
subcategory: "Generative Ai Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_agent_knowledge_base"
sidebar_current: "docs-oci-datasource-generative_ai_agent-knowledge_base"
description: |-
  Provides details about a specific Knowledge Base in Oracle Cloud Infrastructure Generative Ai Agent service
---

# Data Source: oci_generative_ai_agent_knowledge_base
This data source provides details about a specific Knowledge Base resource in Oracle Cloud Infrastructure Generative Ai Agent service.

Gets information about a knowledge base.


## Example Usage

```hcl
data "oci_generative_ai_agent_knowledge_base" "test_knowledge_base" {
	#Required
	knowledge_base_id = oci_generative_ai_agent_knowledge_base.test_knowledge_base.id
}
```

## Argument Reference

The following arguments are supported:

* `knowledge_base_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledge base.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the knowledge base.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the knowledge base.
* `index_config` - The index configuration of Knowledge bases. 
	* `cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OpenSearch Cluster.
	* `database_connection` - The connection type for Databases. 
		* `connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Connection.
		* `connection_type` - The type of Database connection. 
	* `database_functions` - Array of Database functions to be used.
		* `name` - The name of the Database function. 
	* `index_config_type` - The type of index. 
	* `indexes` - Index configuration for open search.
		* `name` - The index name in opensearch.
		* `schema` - The index schema details. 
			* `body_key` - Body key name.
			* `embedding_body_key` - Field within customer managed Oracle Cloud Infrastructure OpenSearch document containing the vector embedding for queries.
			* `title_key` - Title key that stores the Title of a document, if available.
			* `url_key` - URL key that stores the URL of a document, if available.
	* `secret_detail` - The details of configured security configuration on OpenSearch. 
		* `client_id` - The IDCS Connect clientId.
		* `idcs_url` - The URL represent authentication url of the IDCS.
		* `scope_url` - Fully qualified scope url
		* `type` - The type of OpenID. 
		* `vault_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret for basic authentication.
	* `should_enable_hybrid_search` - Whether to enable Hybrid search in service managed OpenSearch.
* `knowledge_base_statistics` - Statistics for Default Knowledge Base.
	* `size_in_bytes` - Knowledge Base size in bytes.
	* `total_ingested_files` - Total number of ingested files in Knowledge Base.
* `lifecycle_details` - A message that describes the current state of the knowledge base in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of the knowledge base.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the knowledge base was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the knowledge base was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

