---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_target"
sidebar_current: "docs-oci-resource-cloud_guard-target"
description: |-
  Provides the Target resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_target
This resource provides the Target resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a target (Target resource), using parameters passed in a CreateTargetDetails resource.


## Example Usage

```hcl
resource "oci_cloud_guard_target" "test_target" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.target_display_name
	target_resource_id = oci_cloud_guard_resource.test_resource.id
	target_resource_type = var.target_target_resource_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.target_description
	freeform_tags = {"bar-key"= "value"}
	state = var.target_state
	target_detector_recipes {
		#Required
		detector_recipe_id = oci_cloud_guard_detector_recipe.test_detector_recipe.id

		#Optional
		detector_rules {
			#Required
			details {

				#Optional
				condition_groups {
					#Required
					compartment_id = var.compartment_id
					condition = var.target_target_detector_recipes_detector_rules_details_condition_groups_condition
				}
			}
			detector_rule_id = oci_events_rule.test_rule.id
		}
	}
	target_responder_recipes {
		#Required
		responder_recipe_id = oci_cloud_guard_responder_recipe.test_responder_recipe.id

		#Optional
		responder_rules {
			#Required
			details {

				#Optional
				condition = var.target_target_responder_recipes_responder_rules_details_condition
				configurations {
					#Required
					config_key = var.target_target_responder_recipes_responder_rules_details_configurations_config_key
					name = var.target_target_responder_recipes_responder_rules_details_configurations_name
					value = var.target_target_responder_recipes_responder_rules_details_configurations_value
				}
				mode = var.target_target_responder_recipes_responder_rules_details_mode
			}
			responder_rule_id = oci_events_rule.test_rule.id
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment OCID where the resource is created
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) The target description.

	Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) Display name for the target.

	Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `state` - (Optional) (Updatable) The enablement state of the detector rule
* `target_detector_recipes` - (Optional) (Updatable) List of detector recipes to attach to target
	* `detector_recipe_id` - (Required) Unique identifier for the target detector recipe
	* `detector_rules` - (Optional) (Updatable) List of overrides to be applied to detector rules associated with the target
		* `details` - (Required) (Updatable) Parameters to update detector rule configuration details in a detector recipe attached to a target.
			* `condition_groups` - (Optional) (Updatable) Condition group corresponding to each compartment
				* `compartment_id` - (Required) (Updatable) Compartment OCID associated with condition
				* `condition` - (Required) (Updatable) The base condition resource.
		* `detector_rule_id` - (Required) (Updatable) Unique identifier for the detector rule
* `target_resource_id` - (Required) Resource ID which the target uses to monitor
* `target_resource_type` - (Required) Type of resource that target support (COMPARTMENT/FACLOUD)
* `target_responder_recipes` - (Optional) (Updatable) List of responder recipes to attach to target
	* `responder_recipe_id` - (Required) Unique identifier for responder recipe
	* `responder_rules` - (Optional) (Updatable) List of overrides to be applied to responder rules associated with the target
		* `details` - (Required) (Updatable) Parameters to update details for a responder rule for a target responder recipe. TargetResponderRuleDetails contains all configurations associated with the ResponderRule, whereas UpdateTargetResponderRecipeResponderRuleDetails refers to the details that are to be updated for ResponderRule. 
			* `condition` - (Optional) (Updatable) The base condition resource.
			* `configurations` - (Optional) (Updatable) List of responder rule configurations
				* `config_key` - (Required) (Updatable) Unique identifier of the configuration
				* `name` - (Required) (Updatable) Configuration name
				* `value` - (Required) (Updatable) Configuration value
			* `mode` - (Optional) (Updatable) Execution mode for the responder rule
		* `responder_rule_id` - (Required) (Updatable) Unique identifier for target detector recipe


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

		* `managed_list_types` - List of managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule resource
		* `resource_type` - The type of resource which is monitored by the detector rule. For example, Instance, Database, VCN, Policy. To find the resource type for a particular rule, see [Detector Recipe Reference] (/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).

			Or try [Detector Recipe Reference] (/cloud-guard/using/detect-recipes.htm#detect-recipes-reference). 
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
		* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.

		* `managed_list_types` - List of managed list types related to this rule
		* `recommendation` - Recommendation for TargetDetectorRecipeDetectorRule resource
		* `resource_type` - The type of resource which is monitored by the detector rule. For example, Instance, Database, VCN, Policy. To find the resource type for a particular rule, see [Detector Recipe Reference] (/iaas/cloud-guard/using/detect-recipes.htm#detect-recipes-reference).

			Or try [Detector Recipe Reference] (/cloud-guard/using/detect-recipes.htm#detect-recipes-reference). 
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target
	* `update` - (Defaults to 20 minutes), when updating the Target
	* `delete` - (Defaults to 20 minutes), when destroying the Target


## Import

Targets can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_target.test_target "id"
```

