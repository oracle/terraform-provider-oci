---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet"
sidebar_current: "docs-oci-resource-fleet_apps_management-fleet"
description: |-
  Provides the Fleet resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_fleet
This resource provides the Fleet resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new fleet instance that includes fleet resources and properties.
For more information, please see the documentation.


## Example Usage

```hcl
resource "oci_fleet_apps_management_fleet" "test_fleet" {
	#Required
	compartment_id = var.compartment_id
	fleet_type = var.fleet_fleet_type

	#Optional
	application_type = var.fleet_application_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.fleet_description
	display_name = var.fleet_display_name
	environment_type = var.fleet_environment_type
	freeform_tags = {"bar-key"= "value"}
	group_type = var.fleet_group_type
	is_target_auto_confirm = var.fleet_is_target_auto_confirm
	notification_preferences {
		#Required
		compartment_id = var.compartment_id
		topic_id = oci_ons_notification_topic.test_notification_topic.id

		#Optional
		preferences {

			#Optional
			on_job_failure = var.fleet_notification_preferences_preferences_on_job_failure
			on_topology_modification = var.fleet_notification_preferences_preferences_on_topology_modification
			on_upcoming_schedule = var.fleet_notification_preferences_preferences_on_upcoming_schedule
		}
	}
	products = var.fleet_products
	resource_selection_type = var.fleet_resource_selection_type
	resources {
		#Required
		compartment_id = var.compartment_id
		resource_id = oci_cloud_guard_resource.test_resource.id
		tenancy_id = oci_identity_tenancy.test_tenancy.id

		#Optional
		fleet_resource_type = var.fleet_resources_fleet_resource_type
	}
	rule_selection_criteria {
		#Required
		match_condition = var.fleet_rule_selection_criteria_match_condition
		rules {
			#Required
			compartment_id = var.compartment_id
			conditions {
				#Required
				attr_group = var.fleet_rule_selection_criteria_rules_conditions_attr_group
				attr_key = var.fleet_rule_selection_criteria_rules_conditions_attr_key
				attr_value = var.fleet_rule_selection_criteria_rules_conditions_attr_value
			}
			resource_compartment_id = oci_identity_compartment.test_compartment.id

			#Optional
			basis = var.fleet_rule_selection_criteria_rules_basis
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `application_type` - (Optional) Application Type associated with the Fleet.Applicable for Environment fleet types.
* `compartment_id` - (Required) Tenancy OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - (Optional) Environment Type associated with the Fleet.Applicable for Environment fleet types.
* `fleet_type` - (Required) Type of the Fleet
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group_type` - (Optional) Group Type associated with Group Fleet.Applicable for Group fleet types.
* `is_target_auto_confirm` - (Optional) (Updatable) A value which represents if auto confirming of the targets can be enabled
* `notification_preferences` - (Optional) (Updatable) Conditions when met to send notifications on the fleet activities
	* `compartment_id` - (Required) (Updatable) Copartment Id of the topic where the notifications will be directed
	* `preferences` - (Optional) (Updatable) Preferences to send notifications on the fleet activities
		* `on_job_failure` - (Optional) (Updatable) Enables or disables notification on Job Failures.'
		* `on_topology_modification` - (Optional) (Updatable) Enables or disables notification on Environment Fleet Topology Modification.
		* `on_upcoming_schedule` - (Optional) (Updatable) Enables notification on upcoming schedule.
	* `topic_id` - (Required) (Updatable) Topic Id where the notifications will be directed
* `products` - (Optional) Products associated with the Fleet
* `resource_selection_type` - (Optional) Type of resource selection in a fleet
* `resources` - (Optional) Resources to be added during fleet creation when Resource selection type is Manual.
	* `compartment_id` - (Required) Compartment Identifier.
	* `fleet_resource_type` - (Optional) Type of the FleetResource.
	* `resource_id` - (Required) OCID of the reosurce.
	* `tenancy_id` - (Required) Tenancy Identifier.
* `rule_selection_criteria` - (Optional) (Updatable) Rule Selection Criteria
	* `match_condition` - (Required) (Updatable) Rule selection match condition.
	* `rules` - (Required) (Updatable) Rules.
		* `basis` - (Optional) (Updatable) Rule to be be applied on.
		* `compartment_id` - (Required) (Updatable) Please provide the root compartmentId (TenancyId).
		* `conditions` - (Required) (Updatable) Rule Conditions
			* `attr_group` - (Required) (Updatable) Attribute Group.
			* `attr_key` - (Required) (Updatable) Attribute Key.
			* `attr_value` - (Required) (Updatable) Attribute Value.
		* `resource_compartment_id` - (Required) (Updatable) Resource Compartment Id.Provide the compartmentId the resource belongs to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_type` - Application Type associated with the Fleet.Applicable for ENVIRONMENT fleet types.
* `compartment_id` - Tenancy OCID
* `credentials` - Credentials to be added during fleet creation.
	* `compartment_id` - Tenancy OCID
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `entity_specifics` - Credential Details
		* `credential_level` - Credential Level.
		* `resource_id` - OCID of the resource associated with the target for which credential is created
		* `target` - Target associated with the Credential
	* `password` - Credential Details
		* `credential_type` - Credential Type
		* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - The Vault Key version.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - The secret version.
		* `value` - The value corresponding to the credential
		* `vault_id` - OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.
	* `user` - Credential Details
		* `credential_type` - Credential Type
		* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - The Vault Key version.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - The secret version.
		* `value` - The value corresponding to the credential
		* `vault_id` - OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - Environment Type associated with the Fleet.Applicable for ENVIRONMENT fleet types.
* `fleet_type` - Type of the Fleet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group_type` - Group Type associated with Group Fleet.Applicable for GROUP fleet types.
* `id` - The OCID of the resource.
* `is_target_auto_confirm` - A value which represents if auto confirming of the targets can be enabled
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `notification_preferences` - Conditions when met to send notifications on the fleet activities
	* `compartment_id` - Copartment Id of the topic where the notifications will be directed
	* `preferences` - Preferences to send notifications on the fleet activities
		* `on_job_failure` - Enables or disables notification on Job Failures.'
		* `on_topology_modification` - Enables or disables notification on Environment Fleet Topology Modification.
		* `on_upcoming_schedule` - Enables notification on upcoming schedule.
	* `topic_id` - Topic Id where the notifications will be directed
* `products` - Products associated with the Fleet
* `properties` - Properties to be added during fleet creation.
	* `compartment_id` - Tenancy OCID
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_property_type` - Type of the FleetProperty.
	* `is_required` - Property is required or not
	* `value` - Value of the Property
* `resource_region` - Associated region
* `resource_selection_type` - Type of resource selection in a fleet.
* `resources` - Resources to be added during fleet creation when Resource selection type is Manual.
	* `compartment_id` - Compartment Identifier.
	* `fleet_resource_type` - Type of the FleetResource.
	* `resource_id` - OCID of the reosurce.
	* `tenancy_id` - Tenancy Identifier.
* `rule_selection_criteria` - Rule Selection Criteria
	* `match_condition` - Rule selection match condition.
	* `rules` - Rules.
		* `basis` - Rule to be be applied on.
		* `compartment_id` - Please provide the root compartmentId (TenancyId).
		* `conditions` - Rule Conditions
			* `attr_group` - Attribute Group.
			* `attr_key` - Attribute Key.
			* `attr_value` - Attribute Value.
		* `resource_compartment_id` - Resource Compartment Id.Provide the compartmentId the resource belongs to.
* `state` - The lifecycle state of the Fleet.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet
	* `update` - (Defaults to 20 minutes), when updating the Fleet
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet


## Import

Fleets can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_fleet.test_fleet "id"
```

