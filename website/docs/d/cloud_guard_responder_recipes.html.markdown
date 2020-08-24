---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_responder_recipes"
sidebar_current: "docs-oci-datasource-cloud_guard-responder_recipes"
description: |-
  Provides the list of Responder Recipes in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_responder_recipes
This data source provides the list of Responder Recipes in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all ResponderRecipes in a compartment
The ListResponderRecipe operation returns only the targets in `compartmentId` passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListResponderRecipe on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_responder_recipes" "test_responder_recipes" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	access_level = "${var.responder_recipe_access_level}"
	compartment_id_in_subtree = "${var.responder_recipe_compartment_id_in_subtree}"
	display_name = "${var.responder_recipe_display_name}"
	resource_metadata_only = "${var.responder_recipe_resource_metadata_only}"
	state = "${var.responder_recipe_state}"
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `resource_metadata_only` - (Optional) Default is false. When set to true, the list of all Oracle Managed Resources Metadata supported by Cloud Guard is returned. 
* `state` - (Optional) The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `responder_recipe_collection` - The list of responder_recipe_collection.

### ResponderRecipe Reference

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

