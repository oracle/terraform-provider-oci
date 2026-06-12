---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_projects"
sidebar_current: "docs-oci-datasource-generative_ai-projects"
description: |-
  Provides the list of Projects in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_projects
This data source provides the list of Projects in Oracle Cloud Infrastructure Generative AI service.

Lists the generativeAiProjects of a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_projects" "test_projects" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.project_display_name
	id = var.project_id
	state = var.project_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the generativeAiProject.
* `state` - (Optional) A filter to return only resources whose lifecycle state matches the given value.


## Attributes Reference

The following attributes are exported:

* `generative_ai_project_collection` - The list of generative_ai_project_collection.

### Project Reference

The following attributes are exported:

* `compartment_id` - Owning compartment OCID for a GenerativeAiProject.
* `conversation_config` - Holds configuration related to conversation retention
	* `conversations_retention_in_hours` - Retention period (in hours) for conversations. The TTL starts from the time the conversation was last updated.
	* `responses_retention_in_hours` - Retention period (in hours) for responses. The TTL starts from the time the response was created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - An optional description of the GenerativeAiProject.
* `display_name` - A user-friendly name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - An OCID that uniquely identifies a GenerativeAiProject.
* `lifecycle_details` - A message describing the current state in more detail that can provide actionable information.
* `long_term_memory_config` - Configuration settings for long-term memory behavior.
	* `standard_long_term_memory_strategy` - Standard strategy settings for long-term memory.
		* `embedding_config` - Configuration for generating embeddings from extracted information.
			* `llm_selection` - LLM selection configuration.
				* `llm_selection_type` - The type of LLM selection.
				* `model_id` - The id of the GenAI model
		* `extraction_config` - Configuration for information extraction from conversation content.
			* `llm_selection` - LLM selection configuration.
				* `llm_selection_type` - The type of LLM selection.
				* `model_id` - The id of the GenAI model
		* `is_enabled` - Indicates whether long-term memory is enabled.
* `short_term_memory_optimization_config` - Configuration settings for short-term memory optimization.
	* `condenser_config` - Configuration for condensing conversation content.
		* `llm_selection` - LLM selection configuration.
			* `llm_selection_type` - The type of LLM selection.
			* `model_id` - The id of the GenAI model
	* `is_enabled` - Indicates whether short-term memory optimization is enabled.
* `state` - The lifecycle state of a GenerativeAiProject.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time that the generativeAiProject was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the generativeAiProject was updated in the format of an RFC3339 datetime string.
