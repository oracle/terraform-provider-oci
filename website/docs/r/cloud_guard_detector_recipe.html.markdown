---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_detector_recipe"
sidebar_current: "docs-oci-resource-cloud_guard-detector_recipe"
description: |-
  Provides the Detector Recipe resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_detector_recipe
This resource provides the Detector Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a new DetectorRecipe object.


## Example Usage

```hcl
resource "oci_cloud_guard_detector_recipe" "test_detector_recipe" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.detector_recipe_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.detector_recipe_description
	detector = var.detector_recipe_detector
	detector_rules {
		#Required
		details {
			#Required
			is_enabled = var.detector_recipe_detector_rules_details_is_enabled
			risk_level = var.detector_recipe_detector_rules_details_risk_level

			#Optional
			condition = var.detector_recipe_detector_rules_details_condition
			configurations {
				#Required
				config_key = var.detector_recipe_detector_rules_details_configurations_config_key
				name = var.detector_recipe_detector_rules_details_configurations_name

				#Optional
				data_type = var.detector_recipe_detector_rules_details_configurations_data_type
				value = var.detector_recipe_detector_rules_details_configurations_value
				values {
					#Required
					list_type = var.detector_recipe_detector_rules_details_configurations_values_list_type
					managed_list_type = var.detector_recipe_detector_rules_details_configurations_values_managed_list_type
					value = var.detector_recipe_detector_rules_details_configurations_values_value
				}
			}
			data_source_id = oci_cloud_guard_data_source.test_data_source.id
			description = var.detector_recipe_detector_rules_details_description
			entities_mappings {
				#Required
				query_field = var.detector_recipe_detector_rules_details_entities_mappings_query_field

				#Optional
				display_name = var.detector_recipe_detector_rules_details_entities_mappings_display_name
				entity_type = var.detector_recipe_detector_rules_details_entities_mappings_entity_type
			}
			labels = var.detector_recipe_detector_rules_details_labels
			recommendation = var.detector_recipe_detector_rules_details_recommendation
		}
		detector_rule_id = oci_events_rule.test_rule.id
	}
	freeform_tags = {"bar-key"= "value"}
	source_detector_recipe_id = oci_cloud_guard_detector_recipe.test_detector_recipe.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Detector recipe description.

	Avoid entering confidential information. 
* `detector` - (Optional) detector for the rule
* `detector_rules` - (Optional) (Updatable) Detector Rules to override from source detector recipe
	* `details` - (Required) (Updatable) Details of a Detector Rule to be overriden in Detector Recipe
		* `condition` - (Optional) (Updatable) Base condition object
		* `configurations` - (Optional) (Updatable) Configuration details
			* `config_key` - (Required) (Updatable) Unique name of the configuration
			* `data_type` - (Optional) (Updatable) configuration data type
			* `name` - (Required) (Updatable) configuration name
			* `value` - (Optional) (Updatable) configuration value
			* `values` - (Optional) (Updatable) List of configuration values
				* `list_type` - (Required) (Updatable) configuration list item type, either CUSTOM or MANAGED
				* `managed_list_type` - (Required) (Updatable) type of the managed list
				* `value` - (Required) (Updatable) configuration value
		* `data_source_id` - (Optional) (Updatable) The id of the attached DataSource.
		* `description` - (Optional) (Updatable) Description for DetectorRecipeDetectorRule.
		* `entities_mappings` - (Optional) (Updatable) Data Source entities mapping for a Detector Rule
			* `display_name` - (Optional) (Updatable) The display name of entity
			* `entity_type` - (Optional) (Updatable) Possible type of entity
			* `query_field` - (Required) (Updatable) The entity value mapped to a data source query
		* `is_enabled` - (Required) (Updatable) Enables the control
		* `labels` - (Optional) (Updatable) user defined labels for a detector rule
		* `recommendation` - (Optional) (Updatable) Recommendation for DetectorRecipeDetectorRule
		* `risk_level` - (Required) (Updatable) The Risk Level
	* `detector_rule_id` - (Required) (Updatable) DetectorRecipeRule Identifier
* `display_name` - (Required) (Updatable) Detector recipe display name.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `source_detector_recipe_id` - (Optional) The id of the source detector recipe.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Detector Recipe
	* `update` - (Defaults to 20 minutes), when updating the Detector Recipe
	* `delete` - (Defaults to 20 minutes), when destroying the Detector Recipe


## Import

DetectorRecipes can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_detector_recipe.test_detector_recipe "id"
```

