---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_detector_recipes"
sidebar_current: "docs-oci-datasource-cloud_guard-detector_recipes"
description: |-
  Provides the list of Detector Recipes in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_detector_recipes
This data source provides the list of Detector Recipes in Oracle Cloud Infrastructure Cloud Guard service.

Returns a list of all detector recipes (DetectorRecipe objects) in a compartment, identified by compartmentId.

The ListDetectorRecipes operation returns only the detector recipes in `compartmentId` passed.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListDetectorRecipes on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_cloud_guard_detector_recipes" "test_detector_recipes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.detector_recipe_access_level
	compartment_id_in_subtree = var.detector_recipe_compartment_id_in_subtree
	display_name = var.detector_recipe_display_name
	resource_metadata_only = var.detector_recipe_resource_metadata_only
	state = var.detector_recipe_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to `RESTRICTED` permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `resource_metadata_only` - (Optional) Default is false. When set to true, the list of all Oracle Managed Resources Metadata supported by Cloud Guard are returned. 
* `state` - (Optional) The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.


## Attributes Reference

The following attributes are exported:

* `detector_recipe_collection` - The list of detector_recipe_collection.

### DetectorRecipe Reference

The following attributes are exported:

* `compartment_id` - compartmentId of detector recipe
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Detector recipe description.
* `detector` - Type of detector
* `detector_rules` - List of detector rules for the detector type for recipe - user input
	* `candidate_responder_rules` - List of CandidateResponderRule related to this rule
		* `display_name` - The display name of the Responder rule
		* `id` - The unique identifier of the Responder rule
		* `is_preferred` - Preferred state
	* `data_source_id` - The id of the attached DataSource.
	* `description` - Description for DetectorRecipeDetectorRule.
	* `details` - Details of a Detector Rule
		* `condition` - Base condition object
		* `configurations` - Configuration details
			* `config_key` - Unique name of the configuration
			* `data_type` - configuration data type
			* `name` - configuration name
			* `value` - configuration value
			* `values` - List of configuration values
				* `list_type` - configuration list item type, either CUSTOM or MANAGED
				* `managed_list_type` - type of the managed list
				* `value` - configuration value
		* `is_configuration_allowed` - configuration allowed or not
		* `is_enabled` - Enables the control
		* `labels` - user defined labels for a detector rule
		* `risk_level` - The Risk Level
	* `detector` - detector for the rule
	* `detector_rule_id` - The unique identifier of the detector rule.
	* `display_name` - Display name for DetectorRecipeDetectorRule.
	* `entities_mappings` - Data Source entities mapping for a Detector Rule
		* `display_name` - The display name of entity
		* `entity_type` - Possible type of entity
		* `query_field` - The entity value mapped to a data source query
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `managed_list_types` - List of cloudguard managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule
	* `resource_type` - resource type of the configuration to which the rule is applied
	* `service_type` - service type of the configuration to which the rule is applied
	* `state` - The current state of the DetectorRule.
	* `time_created` - The date and time the detector recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the detector recipe rule was updated. Format defined by RFC3339.
* `display_name` - DisplayName of detector recipe.
* `effective_detector_rules` - List of effective detector rules for the detector type for recipe after applying defaults
	* `candidate_responder_rules` - List of CandidateResponderRule related to this rule
		* `display_name` - The display name of the Responder rule
		* `id` - The unique identifier of the Responder rule
		* `is_preferred` - Preferred state
	* `data_source_id` - The id of the attached DataSource.
	* `description` - Description for DetectorRecipeDetectorRule.
	* `details` - Details of a Detector Rule
		* `condition` - Base condition object
		* `configurations` - Configuration details
			* `config_key` - Unique name of the configuration
			* `data_type` - configuration data type
			* `name` - configuration name
			* `value` - configuration value
			* `values` - List of configuration values
				* `list_type` - configuration list item type, either CUSTOM or MANAGED
				* `managed_list_type` - type of the managed list
				* `value` - configuration value
		* `is_configuration_allowed` - configuration allowed or not
		* `is_enabled` - Enables the control
		* `labels` - user defined labels for a detector rule
		* `risk_level` - The Risk Level
	* `detector` - detector for the rule
	* `detector_rule_id` - The unique identifier of the detector rule.
	* `display_name` - Display name for DetectorRecipeDetectorRule.
	* `entities_mappings` - Data Source entities mapping for a Detector Rule
		* `display_name` - The display name of entity
		* `entity_type` - Possible type of entity
		* `query_field` - The entity value mapped to a data source query
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `managed_list_types` - List of cloudguard managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule
	* `resource_type` - resource type of the configuration to which the rule is applied
	* `service_type` - service type of the configuration to which the rule is applied
	* `state` - The current state of the DetectorRule.
	* `time_created` - The date and time the detector recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the detector recipe rule was updated. Format defined by RFC3339.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Ocid for detector recipe
* `owner` - Owner of detector recipe
* `source_detector_recipe_id` - Recipe Ocid of the Source Recipe to be cloned
* `state` - The current state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_ids` - The recipe attached to targets
* `time_created` - The date and time the detector recipe was created. Format defined by RFC3339.
* `time_updated` - The date and time the detector recipe was updated. Format defined by RFC3339.

