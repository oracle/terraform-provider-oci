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

Creates a responder recipe (ResponderRecipe resource), from values passed in a
CreateResponderRecipeDetails resource.


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

* `compartment_id` - (Required) (Updatable) Compartment OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Responder recipe description.

	Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) Responder recipe display name.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `responder_rules` - (Optional) (Updatable) List of responder rules to override from source responder recipe
	* `compartment_id` - (Optional) (Updatable) Compartment OCID
	* `details` - (Required) (Updatable) Parameters to be updated for a responder rule within a responder recipe.
		* `is_enabled` - (Required) (Updatable) Enablement state for the responder rule
	* `responder_rule_id` - (Required) (Updatable) Unique identifier for the responder rule
* `source_responder_recipe_id` - (Required) The unique identifier of the source responder recipe


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Responder recipe description
* `display_name` - Responder recipe display name
* `effective_responder_rules` - List of currently enabled responder rules for the responder type, for recipe after applying defaults
	* `compartment_id` - Compartment OCID
	* `description` - Responder rule description
	* `details` - Detailed information for a responder rule
		* `condition` - The base condition resource.
		* `configurations` - List of responder rule configurations
			* `config_key` - Unique identifier of the configuration
			* `name` - Configuration name
			* `value` - Configuration value
		* `is_enabled` - Enabled state for the responder rule
		* `mode` - Execution mode for the responder rule
	* `display_name` - Responder rule display name
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `policies` - List of policies
	* `responder_rule_id` - Unique identifier for the responder rule
	* `state` - The current lifecycle state of the responder rule
	* `supported_modes` - Supported execution modes for the responder rule
	* `time_created` - The date and time the responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the responder recipe rule was last updated. Format defined by RFC3339.
	* `type` - Type of responder
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier for the responder recip
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `owner` - Owner of responder recipe
* `responder_rules` - List of responder rules associated with the recipe
	* `compartment_id` - Compartment OCID
	* `description` - Responder rule description
	* `details` - Detailed information for a responder rule
		* `condition` - The base condition resource.
		* `configurations` - List of responder rule configurations
			* `config_key` - Unique identifier of the configuration
			* `name` - Configuration name
			* `value` - Configuration value
		* `is_enabled` - Enabled state for the responder rule
		* `mode` - Execution mode for the responder rule
	* `display_name` - Responder rule display name
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `policies` - List of policies
	* `responder_rule_id` - Unique identifier for the responder rule
	* `state` - The current lifecycle state of the responder rule
	* `supported_modes` - Supported execution modes for the responder rule
	* `time_created` - The date and time the responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the responder recipe rule was last updated. Format defined by RFC3339.
	* `type` - Type of responder
* `source_responder_recipe_id` - The unique identifier of the source responder recipe
* `state` - The current lifecycle state of the example
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the responder recipe was created. Format defined by RFC3339.
* `time_updated` - The date and time the responder recipe was last updated. Format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Responder Recipe
	* `update` - (Defaults to 20 minutes), when updating the Responder Recipe
	* `delete` - (Defaults to 20 minutes), when destroying the Responder Recipe


## Import

ResponderRecipes can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_responder_recipe.test_responder_recipe "id"
```

