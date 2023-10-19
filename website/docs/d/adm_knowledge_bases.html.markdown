---
subcategory: "Adm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_adm_knowledge_bases"
sidebar_current: "docs-oci-datasource-adm-knowledge_bases"
description: |-
  Provides the list of Knowledge Bases in Oracle Cloud Infrastructure ADM service
---

# Data Source: oci_adm_knowledge_bases
This data source provides the list of Knowledge Bases in Oracle Cloud Infrastructure ADM service.

Returns a list of KnowledgeBases based on the specified query parameters.
At least id or compartmentId query parameter must be provided.


## Example Usage

```hcl
data "oci_adm_knowledge_bases" "test_knowledge_bases" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.knowledge_base_display_name
	id = var.knowledge_base_id
	state = var.knowledge_base_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) A filter to return only resources that belong to the specified compartment identifier. Required only if the id query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only resources that match the specified identifier. Required only if the compartmentId query parameter is not specified. 
* `state` - (Optional) A filter to return only Knowledge Bases that match the specified lifecycleState.


## Attributes Reference

The following attributes are exported:

* `knowledge_base_collection` - The list of knowledge_base_collection.

### KnowledgeBase Reference

The following attributes are exported:

* `compartment_id` - The compartment Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The name of the knowledge base.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the knowledge base.
* `state` - The current lifecycle state of the knowledge base.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The creation date and time of the knowledge base (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_updated` - The date and time the knowledge base was last updated (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).

