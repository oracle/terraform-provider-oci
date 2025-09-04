---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_target"
sidebar_current: "docs-oci-datasource-cloud_guard-target"
description: |-
  Provides details about a specific Target in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_target
This data source provides details about a specific Target resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a target (Target resource) identified by targetId.

## Example Usage

```hcl
data "oci_cloud_guard_target" "test_target" {
	#Required
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) OCID of the target


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID where the resource is created
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The target description
* `display_name` - Target display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that can't be changed after creation
* `inherited_by_compartments` - List of inherited compartments
* `lifecyle_details` - A message describing the current lifecycle state in more detail. For example, can be used to provide actionable information for a resource in Failed state. [DEPRECATE]
* `recipe_count` - Total number of recipes attached to target
* `state` - The current lifecycle state of the target
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_details` - Details specific to the target type.
	* `security_zone_display_name` - The name of the security zone to associate with this compartment.
	* `security_zone_id` - The OCID of the security zone to associate with this compartment
	* `target_resource_type` - Target type, determined by the type of resource for which the target was created
	* `target_security_zone_recipes` - The list of security zone recipes to associate with this compartment
		* `compartment_id` - The OCID of the compartment that contains the recipe
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
		* `description` - The recipe's description
		* `display_name` - The recipe's display name
		* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

			Avoid entering confidential information. 
		* `id` - Unique identifier that can’t be changed after creation
		* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a recipe in the `Failed` state.
		* `owner` - The owner of the recipe
		* `security_policies` - The list of security policy IDs that are included in the recipe
		* `state` - The current lifecycle state of the recipe
		* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
		* `time_created` - The time the recipe was created. An RFC3339 formatted datetime string.
		* `time_updated` - The time the recipe was last updated. An RFC3339 formatted datetime string.
* `target_detector_recipes` - List of detector recipes attached to target
	* `compartment_id` - Compartment OCID of the detector recipe
	* `description` - Detector recipe description.
	* `detector` - Type of detector
	* `detector_recipe_id` - Unique identifier for of original Oracle-managed detector recipe on which the TargetDetectorRecipe is based
	* `detector_recipe_type` - Recipe type ( STANDARD, ENTERPRISE )
	* `detector_rules` - List of detector rules for the detector recipe - user input
		* `data_source_id` - The ID of the attached data source
		* `description` - Description for TargetDetectorRecipeDetectorRule resource
		* `details` - Overriden settings of a detector rule in recipe attached to target.
			* `condition_groups` - Condition group corresponding to each compartment
				* `compartment_id` - Compartment OCID associated with condition
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
			* `is_configuration_allowed` - Configuration allowed or not
			* `is_enabled` - Enablement state of the detector rule
			* `labels` - User-defined labels for a detector rule
			* `risk_level` - The risk level of the detector rule
		* `detector` - Detector type for the rule
		* `detector_rule_id` - The unique identifier of the detector rule
		* `display_name` - Display name for TargetDetectorRecipeDetectorRule resource
		* `entities_mappings` - Data source entities mapping for a detector rule
			* `display_name` - Display name of the entity
			* `entity_type` - Type of entity
			* `query_field` - The entity value mapped to a data source query
		* `is_cloneable` - Is the rule cloneable?
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

		* `managed_list_types` - List of managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule resource
		* `resource_type` - The type of resource which is monitored by the detector rule. For example, Instance, Database, VCN, Policy. To find the resource type for a particular rule, see [Detector Recipe Reference] (/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).

			Or try [Detector Recipe Reference] (/cloud-guard/using/detect-recipes.htm#detect-recipes-reference). 
		* `rule_type` - Rule category type
			* `key` - The unique identifier of the detector rule type
			* `value` - Detector rule type value
		* `service_type` - Service type of the configuration to which the rule is applied
		* `state` - The current lifecycle state of the detector rule
		* `time_created` - The date and time the target detector recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target detector recipe rule was last updated. Format defined by RFC3339.
	* `display_name` - Display name of the detector recipe
	* `effective_detector_rules` - List of currently enabled detector rules for the detector type for recipe after applying defaults
		* `data_source_id` - The ID of the attached data source
		* `description` - Description for TargetDetectorRecipeDetectorRule resource
		* `details` - Overriden settings of a detector rule in recipe attached to target.
			* `condition_groups` - Condition group corresponding to each compartment
				* `compartment_id` - Compartment OCID associated with condition
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
			* `is_configuration_allowed` - Configuration allowed or not
			* `is_enabled` - Enablement state of the detector rule
			* `labels` - User-defined labels for a detector rule
			* `risk_level` - The risk level of the detector rule
		* `detector` - Detector type for the rule
		* `detector_rule_id` - The unique identifier of the detector rule
		* `display_name` - Display name for TargetDetectorRecipeDetectorRule resource
		* `entities_mappings` - Data source entities mapping for a detector rule
			* `display_name` - Display name of the entity
			* `entity_type` - Type of entity
			* `query_field` - The entity value mapped to a data source query
		* `is_cloneable` - Is the rule cloneable?
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

		* `managed_list_types` - List of managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule resource
		* `resource_type` - The type of resource which is monitored by the detector rule. For example, Instance, Database, VCN, Policy. To find the resource type for a particular rule, see [Detector Recipe Reference] (/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).

			Or try [Detector Recipe Reference] (/cloud-guard/using/detect-recipes.htm#detect-recipes-reference). 
		* `rule_type` - Rule category type
			* `key` - The unique identifier of the detector rule type
			* `value` - Detector rule type value
		* `service_type` - Service type of the configuration to which the rule is applied
		* `state` - The current lifecycle state of the detector rule
		* `time_created` - The date and time the target detector recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target detector recipe rule was last updated. Format defined by RFC3339.
	* `id` - OCID for the detector recipe

	* `owner` - Owner of the detector recipe
	* `state` - The current lifecycle state of the resource
	* `time_created` - The date and time the target detector recipe was created. Format defined by RFC3339.
	* `time_updated` - The date and time the target detector recipe was last updated. Format defined by RFC3339.
* `target_resource_id` - Resource ID which the target uses to monitor
* `target_resource_type` - Type of target
* `target_responder_recipes` - List of responder recipes attached to target
	* `compartment_id` - Compartment OCID
	* `description` - Target responder description
	* `display_name` - Target responder recipe display name
	* `effective_responder_rules` - List of currently enabled responder rules for the responder type for recipe after applying defaults
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
		* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target responder recipe rule was last updated. Format defined by RFC3339.
		* `type` - Type of responder
	* `id` - Unique identifier of target responder recipe that can't be changed after creation

	* `owner` - Owner of target responder recipe
	* `responder_recipe_id` - Unique identifier for the Oracle-managed responder recipe from which this recipe was cloned
	* `responder_rules` - List of responder rules associated with the recipe - user input
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
		* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
		* `time_updated` - The date and time the target responder recipe rule was last updated. Format defined by RFC3339.
		* `type` - Type of responder
	* `time_created` - The date and time the target responder recipe rule was created. Format defined by RFC3339.
	* `time_updated` - The date and time the target responder recipe rule was last updated. Format defined by RFC3339.
* `time_created` - The date and time the target was created. Format defined by RFC3339.
* `time_updated` - The date and time the target was last updated. Format defined by RFC3339.

