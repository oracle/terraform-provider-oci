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

Creates a new DetectorRecipe resource.


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

* `compartment_id` - (Required) (Updatable) Compartment OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Detector recipe description.

	Avoid entering confidential information. 
* `detector` - (Optional) Detector for the rule
* `detector_rules` - (Optional) (Updatable) Detector rules to override from source detector recipe
	* `details` - (Required) (Updatable) Parameters to be updated for a detector rule within a detector recipe.
		* `condition` - (Optional) (Updatable) The base condition resource.
		* `configurations` - (Optional) (Updatable) List of detector rule configurations
			* `config_key` - (Required) (Updatable) Unique identifier of the configuration
			* `data_type` - (Optional) (Updatable) Configuration data type
			* `name` - (Required) (Updatable) Configuration name
			* `value` - (Optional) (Updatable) Configuration value
			* `values` - (Optional) (Updatable) List of configuration values
				* `list_type` - (Required) (Updatable) Configuration list item type (CUSTOM or MANAGED)
				* `managed_list_type` - (Required) (Updatable) Type of content in the managed list
				* `value` - (Required) (Updatable) Configuration value
		* `data_source_id` - (Optional) (Updatable) The unique identifier of the attached data source
		* `description` - (Optional) (Updatable) Description for the detector rule
		* `entities_mappings` - (Optional) (Updatable) Data source entities mapping for a detector rule
			* `display_name` - (Optional) (Updatable) Display name of the entity
			* `entity_type` - (Optional) (Updatable) Type of entity
			* `query_field` - (Required) (Updatable) The entity value mapped to a data source query
		* `is_enabled` - (Required) (Updatable) Enablement status of the detector rule
		* `labels` - (Optional) (Updatable) User-defined labels for a detector rule
		* `recommendation` - (Optional) (Updatable) Recommendation for the detector rule
		* `risk_level` - (Required) (Updatable) The risk level of the detector rule
	* `detector_rule_id` - (Required) (Updatable) Detector recipe rule ID
* `display_name` - (Required) (Updatable) Detector recipe display name.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `source_detector_recipe_id` - (Optional) The ID of the source detector recipe


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID of detector recipe
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Detector recipe description
* `detector` - Type of detector
* `detector_recipe_type` - Recipe type ( STANDARD, ENTERPRISE )
* `detector_rules` - List of detector rules for the detector type for recipe - user input
	* `candidate_responder_rules` - List of responder rules that can be used to remediate this detector rule
		* `display_name` - The display name of the responder rule
		* `id` - The unique identifier of the responder rule
		* `is_preferred` - Is this the preferred state?
	* `data_source_id` - The unique identifier of the attached data source
	* `description` - Description for DetectorRecipeDetectorRule resource
	* `details` - Detailed information for a detector.
		* `condition` - The base condition resource.
		* `configurations` - List of detector rule configurations
			* `config_key` - Unique identifier of the configuration
			* `data_type` - Configuration data type
			* `name` - Configuration name
			* `value` - Configuration value
			* `values` - List of configuration values
				* `list_type` - Configuration list item type (CUSTOM or MANAGED)
				* `managed_list_type` - Type of content in the managed list
				* `value` - Configuration value
		* `data_source_id` - The ID of the attached data source
		* `description` - Description for detector recipe detector rule
		* `entities_mappings` - Data source entities mapping for a detector rule
			* `display_name` - Display name of the entity
			* `entity_type` - Type of entity
			* `query_field` - The entity value mapped to a data source query
		* `is_configuration_allowed` - Can the rule be configured?
		* `is_enabled` - Enablement status for the rule
		* `labels` - User-defined labels for a detector rule
		* `recommendation` - Recommendation for detector recipe detector rule
		* `risk_level` - The risk level for the rule
	* `detector` - Detector recipe for the rule
	* `detector_rule_id` - The unique identifier of the detector rule.
	* `display_name` - Display name for DetectorRecipeDetectorRule resource
	* `entities_mappings` - Data source entities mapping for the detector rule
		* `display_name` - Display name of the entity
		* `entity_type` - Type of entity
		* `query_field` - The entity value mapped to a data source query
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `managed_list_types` - List of managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule resource
	* `resource_type` - Resource type of the configuration to which the rule is applied
	* `service_type` - Service type of the configuration to which the rule is applied
	* `state` - The current lifecycle state of the detector rule.
	* `time_created` - The date and time the detector recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the detector recipe rule was last updated. Format defined by RFC3339.
* `display_name` - Display name of detector recipe
* `effective_detector_rules` - List of effective detector rules for the detector type for recipe after applying defaults
	* `candidate_responder_rules` - List of responder rules that can be used to remediate this detector rule
		* `display_name` - The display name of the responder rule
		* `id` - The unique identifier of the responder rule
		* `is_preferred` - Is this the preferred state?
	* `data_source_id` - The unique identifier of the attached data source
	* `description` - Description for DetectorRecipeDetectorRule resource
	* `details` - Detailed information for a detector.
		* `condition` - The base condition resource.
		* `configurations` - List of detector rule configurations
			* `config_key` - Unique identifier of the configuration
			* `data_type` - Configuration data type
			* `name` - Configuration name
			* `value` - Configuration value
			* `values` - List of configuration values
				* `list_type` - Configuration list item type (CUSTOM or MANAGED)
				* `managed_list_type` - Type of content in the managed list
				* `value` - Configuration value
		* `data_source_id` - The ID of the attached data source
		* `description` - Description for detector recipe detector rule
		* `entities_mappings` - Data source entities mapping for a detector rule
			* `display_name` - Display name of the entity
			* `entity_type` - Type of entity
			* `query_field` - The entity value mapped to a data source query
		* `is_configuration_allowed` - Can the rule be configured?
		* `is_enabled` - Enablement status for the rule
		* `labels` - User-defined labels for a detector rule
		* `recommendation` - Recommendation for detector recipe detector rule
		* `risk_level` - The risk level for the rule
	* `detector` - Detector recipe for the rule
	* `detector_rule_id` - The unique identifier of the detector rule.
	* `display_name` - Display name for DetectorRecipeDetectorRule resource
	* `entities_mappings` - Data source entities mapping for the detector rule
		* `display_name` - Display name of the entity
		* `entity_type` - Type of entity
		* `query_field` - The entity value mapped to a data source query
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `managed_list_types` - List of managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule resource
	* `resource_type` - Resource type of the configuration to which the rule is applied
	* `service_type` - Service type of the configuration to which the rule is applied
	* `state` - The current lifecycle state of the detector rule.
	* `time_created` - The date and time the detector recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the detector recipe rule was last updated. Format defined by RFC3339.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - OCID for detector recipe
* `owner` - Owner of detector recipe
* `source_detector_recipe_id` - Recipe OCID of the source recipe to be cloned
* `state` - The current lifecycle state of the resource
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_ids` - List of target IDs to which the recipe is attached
* `time_created` - The date and time the detector recipe was created Format defined by RFC3339.
* `time_updated` - The date and time the detector recipe was last updated Format defined by RFC3339.

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

