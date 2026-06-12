---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_project"
sidebar_current: "docs-oci-datasource-generative_ai-project"
description: |-
  Provides details about a specific Project in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_project
This data source provides details about a specific Project resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about a generativeAiProject.

## Example Usage

```hcl
data "oci_generative_ai_project" "test_project" {
	#Required
	project_id = oci_generative_ai_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the generativeAiProject.


## Attributes Reference

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
