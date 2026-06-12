---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_project"
sidebar_current: "docs-oci-resource-generative_ai-project"
description: |-
  Provides the Project resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_project
This resource provides the Project resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/GenerativeAiProject

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a GenerativeAiProject.
The header contains an opc-work-request-id, which is the id for the WorkRequest that tracks the generativeAiProject creation progress.


## Example Usage

```hcl
resource "oci_generative_ai_project" "test_project" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	conversation_config {

		#Optional
		conversations_retention_in_hours = var.project_conversation_config_conversations_retention_in_hours
		responses_retention_in_hours = var.project_conversation_config_responses_retention_in_hours
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.project_description
	display_name = var.project_display_name
	freeform_tags = {"Department"= "Finance"}
	long_term_memory_config {

		#Optional
		standard_long_term_memory_strategy {
			#Required
			is_enabled = var.project_long_term_memory_config_standard_long_term_memory_strategy_is_enabled

			#Optional
			embedding_config {
				#Required
				llm_selection {
					#Required
					llm_selection_type = var.project_long_term_memory_config_standard_long_term_memory_strategy_embedding_config_llm_selection_llm_selection_type
					model_id = oci_generative_ai_model.test_model.id
				}
			}
			extraction_config {
				#Required
				llm_selection {
					#Required
					llm_selection_type = var.project_long_term_memory_config_standard_long_term_memory_strategy_extraction_config_llm_selection_llm_selection_type
					model_id = oci_generative_ai_model.test_model.id
				}
			}
		}
	}
	short_term_memory_optimization_config {
		#Required
		is_enabled = var.project_short_term_memory_optimization_config_is_enabled

		#Optional
		condenser_config {
			#Required
			llm_selection {
				#Required
				llm_selection_type = var.project_short_term_memory_optimization_config_condenser_config_llm_selection_llm_selection_type
				model_id = oci_generative_ai_model.test_model.id
			}
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Owning compartment OCID for a GenerativeAiProject.
* `conversation_config` - (Optional) (Updatable) Holds configuration related to conversation retention
	* `conversations_retention_in_hours` - (Optional) (Updatable) Retention period (in hours) for conversations. The TTL starts from the time the conversation was last updated.
	* `responses_retention_in_hours` - (Optional) (Updatable) Retention period (in hours) for responses. The TTL starts from the time the response was created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) An optional description of the GenerativeAiProject.
* `display_name` - (Optional) (Updatable) A user-friendly name.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `long_term_memory_config` - (Optional) (Updatable) Configuration settings for long-term memory behavior.
	* `standard_long_term_memory_strategy` - (Optional) (Updatable) Standard strategy settings for long-term memory.
		* `embedding_config` - (Optional) (Updatable) Configuration for generating embeddings from extracted information.
			* `llm_selection` - (Required) (Updatable) LLM selection configuration.
				* `llm_selection_type` - (Required) (Updatable) The type of LLM selection.
				* `model_id` - (Required) (Updatable) The id of the GenAI model
		* `extraction_config` - (Optional) (Updatable) Configuration for information extraction from conversation content.
			* `llm_selection` - (Required) (Updatable) LLM selection configuration.
				* `llm_selection_type` - (Required) (Updatable) The type of LLM selection.
				* `model_id` - (Required) (Updatable) The id of the GenAI model
		* `is_enabled` - (Required) (Updatable) Indicates whether long-term memory is enabled.
* `short_term_memory_optimization_config` - (Optional) (Updatable) Configuration settings for short-term memory optimization.
	* `condenser_config` - (Optional) (Updatable) Configuration for condensing conversation content.
		* `llm_selection` - (Required) (Updatable) LLM selection configuration.
			* `llm_selection_type` - (Required) (Updatable) The type of LLM selection.
			* `model_id` - (Required) (Updatable) The id of the GenAI model
	* `is_enabled` - (Required) (Updatable) Indicates whether short-term memory optimization is enabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Project
	* `update` - (Defaults to 20 minutes), when updating the Project
	* `delete` - (Defaults to 20 minutes), when destroying the Project


## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_project.test_project "id"
```
