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

Returns a list (ResponderRecipeCollection resource, with a page of ResponderRecipeSummary resources)
of all responder recipes (RespponderRecipe resources) in a compartment, identified by compartmentId.
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
	compartment_id = var.compartment_id

	#Optional
	access_level = var.responder_recipe_access_level
	compartment_id_in_subtree = var.responder_recipe_compartment_id_in_subtree
	display_name = var.responder_recipe_display_name
	resource_metadata_only = var.responder_recipe_resource_metadata_only
	state = var.responder_recipe_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the setting of `accessLevel`. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `resource_metadata_only` - (Optional) Default is false. When set to true, the list of all Oracle-managed resources metadata supported by Cloud Guard is returned. 
* `state` - (Optional) The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `responder_recipe_collection` - The list of responder_recipe_collection.

### ResponderRecipe Reference

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

