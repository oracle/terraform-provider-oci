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

Returns a responder recipe (ResponderRecipe resource) identified by responderRecipeId.

## Example Usage

```hcl
data "oci_cloud_guard_responder_recipe" "test_responder_recipe" {
	#Required
	responder_recipe_id = oci_cloud_guard_responder_recipe.test_responder_recipe.id
}
```

## Argument Reference

The following arguments are supported:

* `responder_recipe_id` - (Required) OCID of the responder recipe.


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

