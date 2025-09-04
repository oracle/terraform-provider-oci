---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_detector_recipe"
sidebar_current: "docs-oci-datasource-cloud_guard-detector_recipe"
description: |-
  Provides details about a specific Detector Recipe in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_detector_recipe
This data source provides details about a specific Detector Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a detector recipe (DetectorRecipe resource) identified by detectorRecipeId.

## Example Usage

```hcl
data "oci_cloud_guard_detector_recipe" "test_detector_recipe" {
	#Required
	detector_recipe_id = oci_cloud_guard_detector_recipe.test_detector_recipe.id
}
```

## Argument Reference

The following arguments are supported:

* `detector_recipe_id` - (Required) Detector recipe OCID


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
			* `additional_properties` - Map of additional property values for configuration
				* `key` - Name for Additional Property, for example, "interpreter", "router"
				* `property_type` - Property Type
				* `value` - Value for Property Name, for example, "generic", "cloudguard"
			* `allowed_values` - Map of possible values for configuration
				* `key` - key
				* `value` - value
			* `allowed_values_data_type` - Map property Value data type
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
	* `is_cloneable` - Is the rule cloneable?
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `managed_list_types` - List of managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule resource
	* `resource_type` - Resource type of the configuration to which the rule is applied
	* `rule_type` - Detector rule type
		* `key` - The unique identifier of the detector rule type
		* `value` - Detector rule type value
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
			* `additional_properties` - Map of additional property values for configuration
				* `key` - Name for Additional Property, for example, "interpreter", "router"
				* `property_type` - Property Type
				* `value` - Value for Property Name, for example, "generic", "cloudguard"
			* `allowed_values` - Map of possible values for configuration
				* `key` - key
				* `value` - value
			* `allowed_values_data_type` - Map property Value data type
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
	* `is_cloneable` - Is the rule cloneable?
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

	* `managed_list_types` - List of managed list types related to this rule
	* `recommendation` - Recommendation for DetectorRecipeDetectorRule resource
	* `resource_type` - Resource type of the configuration to which the rule is applied
	* `rule_type` - Detector rule type
		* `key` - The unique identifier of the detector rule type
		* `value` - Detector rule type value
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

