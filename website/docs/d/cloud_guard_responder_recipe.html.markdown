---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_responder_recipe"
sidebar_current: "docs-oci-datasource-cloud_guard-responder_recipe"
description: |-
  Provides details about a specific Responder Recipe in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_responder_recipe
This data source provides details about a specific Responder Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Get a ResponderRecipe by identifier

## Example Usage

```hcl
data "oci_cloud_guard_responder_recipe" "test_responder_recipe" {
	#Required
	responder_recipe_id = oci_cloud_guard_responder_recipe.test_responder_recipe.id
}
```

## Argument Reference

The following arguments are supported:

* `responder_recipe_id` - (Required) OCID of ResponderRecipe


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - ResponderRecipe Description
* `display_name` - ResponderRecipe Display Name
* `effective_responder_rules` - List of responder rules associated with the recipe
	* `compartment_id` - Compartment Identifier
	* `description` - ResponderRule Description
	* `details` - 
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
	* `details` - 
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

