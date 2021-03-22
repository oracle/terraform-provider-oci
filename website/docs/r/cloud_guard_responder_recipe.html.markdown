---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_responder_recipe"
sidebar_current: "docs-oci-resource-cloud_guard-responder_recipe"
description: |-
  Provides the Responder Recipe resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_responder_recipe
This resource provides the Responder Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Create a ResponderRecipe.


## Example Usage

```hcl
resource "oci_cloud_guard_responder_recipe" "test_responder_recipe" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.responder_recipe_display_name
	source_responder_recipe_id = oci_cloud_guard_responder_recipe.test_responder_recipe.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.responder_recipe_description
	freeform_tags = {"bar-key"= "value"}
	responder_rules {
		#Required
		details {
			#Required
			is_enabled = var.responder_recipe_responder_rules_details_is_enabled
		}
		responder_rule_id = oci_events_rule.test_rule.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) ResponderRecipe Description
* `display_name` - (Required) (Updatable) ResponderRecipe Display Name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `responder_rules` - (Optional) (Updatable) Responder Rules to override from source responder recipe
	* `compartment_id` - (Optional) (Updatable) Compartment Identifier
	* `details` - (Required) (Updatable) Details of UpdateResponderRuleDetails.
		* `is_enabled` - (Required) (Updatable) Identifies state for ResponderRule
	* `responder_rule_id` - (Required) (Updatable) ResponderRecipeRule Identifier
* `source_responder_recipe_id` - (Required) The id of the source responder recipe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - ResponderRecipe Description
* `display_name` - ResponderRecipe Display Name
* `effective_responder_rules` - List of responder rules associated with the recipe
	* `compartment_id` - Compartment Identifier
	* `description` - ResponderRule Description
	* `details` - Details of ResponderRule.
		* `condition` - 
		* `configurations` - ResponderRule configurations
			* `config_key` - Unique name of the configuration
			* `name` - configuration name
			* `value` - configuration value
		* `is_enabled` - Identifies state for ResponderRule
		* `mode` - Execution Mode for ResponderRule
	* `display_name` - ResponderRule Display Name
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `policies` - List of Policy
	* `responder_rule_id` - Identifier for ResponderRule.
	* `state` - The current state of the ResponderRule.
	* `supported_modes` - Supported Execution Modes
	* `time_created` - The date and time the responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the responder recipe rule was updated. Format defined by RFC3339.
	* `type` - Type of Responder
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Identifier for ResponderRecipe.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `owner` - Owner of ResponderRecipe
* `responder_rules` - List of responder rules associated with the recipe
	* `compartment_id` - Compartment Identifier
	* `description` - ResponderRule Description
	* `details` - Details of ResponderRule.
		* `condition` - 
		* `configurations` - ResponderRule configurations
			* `config_key` - Unique name of the configuration
			* `name` - configuration name
			* `value` - configuration value
		* `is_enabled` - Identifies state for ResponderRule
		* `mode` - Execution Mode for ResponderRule
	* `display_name` - ResponderRule Display Name
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `policies` - List of Policy
	* `responder_rule_id` - Identifier for ResponderRule.
	* `state` - The current state of the ResponderRule.
	* `supported_modes` - Supported Execution Modes
	* `time_created` - The date and time the responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the responder recipe rule was updated. Format defined by RFC3339.
	* `type` - Type of Responder
* `source_responder_recipe_id` - The id of the source responder recipe.
* `state` - The current state of the Example.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the responder recipe was created. Format defined by RFC3339.
* `time_updated` - The date and time the responder recipe was updated. Format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Responder Recipe
	* `update` - (Defaults to 20 minutes), when updating the Responder Recipe
	* `delete` - (Defaults to 20 minutes), when destroying the Responder Recipe


## Import

ResponderRecipes can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_responder_recipe.test_responder_recipe "id"
```

