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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/Fleet

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Create a fleet in Fleet Application Management.


## Example Usage

```hcl
resource "oci_fleet_apps_management_fleet" "test_fleet" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.fleet_display_name
	resource_selection {
		#Required
		resource_selection_type = var.fleet_resource_selection_resource_selection_type

		#Optional
		rule_selection_criteria {

			#Optional
			match_condition = var.fleet_resource_selection_rule_selection_criteria_match_condition
			rules {

				#Optional
				basis = var.fleet_resource_selection_rule_selection_criteria_rules_basis
				compartment_id = var.compartment_id
				conditions {

					#Optional
					attr_group = var.fleet_resource_selection_rule_selection_criteria_rules_conditions_attr_group
					attr_key = var.fleet_resource_selection_rule_selection_criteria_rules_conditions_attr_key
					attr_value = var.fleet_resource_selection_rule_selection_criteria_rules_conditions_attr_value
				}
				resource_compartment_id = oci_identity_compartment.test_compartment.id
			}
		}
	}

	#Optional
	credentials {
		#Required
		compartment_id = var.compartment_id
		display_name = var.fleet_credentials_display_name
		entity_specifics {
			#Required
			credential_level = var.fleet_credentials_entity_specifics_credential_level

			#Optional
			resource_id = oci_cloud_guard_resource.test_resource.id
			target = var.fleet_credentials_entity_specifics_target
			variables {

				#Optional
				name = var.fleet_credentials_entity_specifics_variables_name
				value = var.fleet_credentials_entity_specifics_variables_value
			}
		}
		password {
			#Required
			credential_type = var.fleet_credentials_password_credential_type

			#Optional
			key_id = oci_kms_key.test_key.id
			key_version = var.fleet_credentials_password_key_version
			secret_id = oci_vault_secret.test_secret.id
			secret_version = var.fleet_credentials_password_secret_version
			value = var.fleet_credentials_password_value
			vault_id = oci_kms_vault.test_vault.id
		}
		user {
			#Required
			credential_type = var.fleet_credentials_user_credential_type

			#Optional
			key_id = oci_kms_key.test_key.id
			key_version = var.fleet_credentials_user_key_version
			secret_id = oci_vault_secret.test_secret.id
			secret_version = var.fleet_credentials_user_secret_version
			value = var.fleet_credentials_user_value
			vault_id = oci_kms_vault.test_vault.id
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.fleet_description
	details {

		#Optional
		fleet_type = var.fleet_details_fleet_type
	}
	environment_type = var.fleet_environment_type
	freeform_tags = {"bar-key"= "value"}
	is_target_auto_confirm = var.fleet_is_target_auto_confirm
	notification_preferences {
		#Required
		compartment_id = var.compartment_id
		topic_id = oci_ons_notification_topic.test_notification_topic.id

		#Optional
		preferences {

			#Optional
			on_job_failure = var.fleet_notification_preferences_preferences_on_job_failure
			on_resource_non_compliance = var.fleet_notification_preferences_preferences_on_resource_non_compliance
			on_runbook_newer_version = var.fleet_notification_preferences_preferences_on_runbook_newer_version
			on_task_failure = var.fleet_notification_preferences_preferences_on_task_failure
			on_task_pause = var.fleet_notification_preferences_preferences_on_task_pause
			on_task_success = var.fleet_notification_preferences_preferences_on_task_success
			on_topology_modification = var.fleet_notification_preferences_preferences_on_topology_modification
			upcoming_schedule {

				#Optional
				notify_before = var.fleet_notification_preferences_preferences_upcoming_schedule_notify_before
				on_upcoming_schedule = var.fleet_notification_preferences_preferences_upcoming_schedule_on_upcoming_schedule
			}
		}
	}
	parent_fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	products = var.fleet_products
	properties {
		#Required
		compartment_id = var.compartment_id
		fleet_property_type = var.fleet_properties_fleet_property_type

		#Optional
		display_name = var.fleet_properties_display_name
		is_required = var.fleet_properties_is_required
		value = var.fleet_properties_value
	}
	resources {
		#Required
		compartment_id = var.compartment_id
		resource_id = oci_cloud_guard_resource.test_resource.id
		tenancy_id = oci_identity_tenancy.test_tenancy.id

		#Optional
		fleet_resource_type = var.fleet_resources_fleet_resource_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) compartment OCID
* `credentials` - (Optional) Credentials associated with the Fleet.
	* `compartment_id` - (Required) (Updatable) Compartment OCID
	* `display_name` - (Required) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `entity_specifics` - (Required) Credential specific Details.
		* `credential_level` - (Required) At what level the credential is provided?
		* `resource_id` - (Required when credential_level=RESOURCE | TARGET) OCID of the resource associated with the target for which the credential is created.
		* `target` - (Required when credential_level=TARGET) Target name for which the credential is provided.
		* `variables` - (Applicable when credential_level=FLEET) List of fleet credential variables.
			* `name` - (Applicable when credential_level=FLEET) Name of the variable.
			* `value` - (Applicable when credential_level=FLEET) The value corresponding to the variable name.
	* `password` - (Required) Credential Details.
		* `credential_type` - (Required) Credential Type.
		* `key_id` - (Required when credential_type=KEY_ENCRYPTION) OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - (Applicable when credential_type=KEY_ENCRYPTION) The Vault Key version.
		* `secret_id` - (Required when credential_type=VAULT_SECRET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - (Applicable when credential_type=VAULT_SECRET) The secret version.
		* `value` - (Required when credential_type=KEY_ENCRYPTION | PLAIN_TEXT) The value corresponding to the credential.
		* `vault_id` - (Required when credential_type=KEY_ENCRYPTION) OCID for the Vault that will be used to fetch the key to encrypt/decrypt the value given.
	* `user` - (Required) Credential Details.
		* `credential_type` - (Required) Credential Type.
		* `key_id` - (Required when credential_type=KEY_ENCRYPTION) OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - (Applicable when credential_type=KEY_ENCRYPTION) The Vault Key version.
		* `secret_id` - (Required when credential_type=VAULT_SECRET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - (Applicable when credential_type=VAULT_SECRET) The secret version.
		* `value` - (Required when credential_type=KEY_ENCRYPTION | PLAIN_TEXT) The value corresponding to the credential.
		* `vault_id` - (Required when credential_type=KEY_ENCRYPTION) OCID for the Vault that will be used to fetch the key to encrypt/decrypt the value given.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - (Optional) Fleet Type
	* `fleet_type` - (Optional) Type of the Fleet. PRODUCT - A fleet of product-specific resources for a product type. ENVIRONMENT - A fleet of environment-specific resources for a product stack. GROUP - A fleet of a fleet of either environment or product fleets. GENERIC - A fleet of resources selected dynamically or manually for reporting purposes 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - (Optional) Environment Type associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_target_auto_confirm` - (Optional) (Updatable) A value that represents if auto-confirming of the targets can be enabled. This will allow targets to be auto-confirmed in the fleet without manual intervention. 
* `notification_preferences` - (Optional) (Updatable) Notification Preferences associated with the Fleet.
	* `compartment_id` - (Required) (Updatable) Compartment ID the topic belongs to.
	* `preferences` - (Optional) (Updatable) Preferences to send notifications on the fleet activities.
		* `on_job_failure` - (Optional) (Updatable) Enables or disables notification on Job Failures.
		* `on_resource_non_compliance` - (Optional) (Updatable) Enables or disables notification when fleet resource becomes non compliant.
		* `on_runbook_newer_version` - (Optional) (Updatable) Enables or disables notification when a newer version of runbook associated with a fleet is available
		* `on_task_failure` - (Optional) (Updatable) Enables or disables notification on task failure.
		* `on_task_pause` - (Optional) (Updatable) Enables or disables notification when a task is paused.
		* `on_task_success` - (Optional) (Updatable) Enables or disables notification on task success.
		* `on_topology_modification` - (Optional) (Updatable) Enables or disables notification on Environment Fleet Topology Modification.
		* `upcoming_schedule` - (Optional) (Updatable) Enables notification on upcoming schedule.
			* `notify_before` - (Optional) (Updatable) Specify when the notification should be sent. 
			* `on_upcoming_schedule` - (Optional) (Updatable) Enables notification on upcoming schedule.
	* `topic_id` - (Required) (Updatable) Topic Id where the notifications will be directed. A topic is a communication channel for sending messages on chosen events to subscriptions. 
* `parent_fleet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet that would be the parent for this fleet. 
* `products` - (Optional) (Updatable) Products associated with the Fleet.
* `properties` - (Optional) Properties associated with the Fleet.
	* `compartment_id` - (Required) (Updatable) compartment OCID
	* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_property_type` - (Required) Type of the FleetProperty.
	* `is_required` - (Optional) Property is required or not.
	* `value` - (Optional) Value of the Property.
* `resource_selection` - (Required) (Updatable) Resource Selection Type
	* `resource_selection_type` - (Required) (Updatable) Type of resource selection in a Fleet. Select resources manually or select resources based on rules. 
	* `rule_selection_criteria` - (Required when resource_selection_type=DYNAMIC) (Updatable) Rule Selection Criteria for DYNAMIC resource selection for a GENERIC fleet. Rules define what resources are members of this fleet. All resources that meet the criteria are added automatically. 
		* `match_condition` - (Required when resource_selection_type=DYNAMIC) (Updatable) Match condition for the rule selection. Include resources that match all rules or any of the rules. 
		* `rules` - (Required when resource_selection_type=DYNAMIC) (Updatable) Rules.
			* `basis` - (Applicable when resource_selection_type=DYNAMIC) (Updatable) Based on what the rule is created. It can be based on a resourceProperty or a tag.   If based on a tag, basis will be 'definedTagEquals' If based on a resource property, basis will be 'inventoryProperties' 
			* `compartment_id` - (Required when resource_selection_type=DYNAMIC) (Updatable) Compartment Id for which the rule is created. 
			* `conditions` - (Required when resource_selection_type=DYNAMIC) (Updatable) Rule Conditions
				* `attr_group` - (Required when resource_selection_type=DYNAMIC) (Updatable) Attribute Group. Provide a Tag namespace if the rule is based on a tag. Provide resource type if the rule is based on a resource property. 
				* `attr_key` - (Required when resource_selection_type=DYNAMIC) (Updatable) Attribute Key.Provide Tag key if the rule is based on a tag. Provide resource property name if the rule is based on a resource property. 
				* `attr_value` - (Required when resource_selection_type=DYNAMIC) (Updatable) Attribute Value.Provide Tag value if the rule is based on a tag. Provide resource property value if the rule is based on a resource property. 
			* `resource_compartment_id` - (Required when resource_selection_type=DYNAMIC) (Updatable) The Compartment ID to dynamically search resources. Provide the compartment ID to which the rule is applicable. 
* `resources` - (Optional) Resources associated with the Fleet if resourceSelectionType is MANUAL.
	* `compartment_id` - (Required) (Updatable) Compartment Identifier[OCID].
	* `fleet_resource_type` - (Optional) Type of the FleetResource.
	* `resource_id` - (Required) OCID of the resource.
	* `tenancy_id` - (Required) Tenancy Identifier[OCID].


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID
* `credentials` - Credentials associated with the Fleet.
	* `compartment_id` - Compartment OCID
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `entity_specifics` - Credential specific Details.
		* `credential_level` - At what level the credential is provided?
		* `resource_id` - OCID of the resource associated with the target for which the credential is created.
		* `target` - Target name for which the credential is provided.
		* `variables` - List of fleet credential variables.
			* `name` - Name of the variable.
			* `value` - The value corresponding to the variable name.
	* `password` - Credential Details.
		* `credential_type` - Credential Type.
		* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - The Vault Key version.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - The secret version.
		* `value` - The value corresponding to the credential.
		* `vault_id` - OCID for the Vault that will be used to fetch the key to encrypt/decrypt the value given.
	* `user` - Credential Details.
		* `credential_type` - Credential Type.
		* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
		* `key_version` - The Vault Key version.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
		* `secret_version` - The secret version.
		* `value` - The value corresponding to the credential.
		* `vault_id` - OCID for the Vault that will be used to fetch the key to encrypt/decrypt the value given.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - Fleet Type
	* `fleet_type` - Type of the Fleet. PRODUCT - A fleet of product-specific resources for a product type. ENVIRONMENT - A fleet of environment-specific resources for a product stack. GROUP - A fleet of a fleet of either environment or product fleets. GENERIC - A fleet of resources selected dynamically or manually for reporting purposes 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - Environment Type associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `is_target_auto_confirm` - A value that represents if auto-confirming of the targets can be enabled. This will allow targets to be auto-confirmed in the fleet without manual intervention. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `notification_preferences` - Notification Preferences associated with the Fleet.
	* `compartment_id` - Compartment ID the topic belongs to.
	* `preferences` - Preferences to send notifications on the fleet activities.
		* `on_job_failure` - Enables or disables notification on Job Failures.
		* `on_resource_non_compliance` - Enables or disables notification when fleet resource becomes non compliant.
		* `on_runbook_newer_version` - Enables or disables notification when a newer version of runbook associated with a fleet is available
		* `on_task_failure` - Enables or disables notification on task failure.
		* `on_task_pause` - Enables or disables notification when a task is paused.
		* `on_task_success` - Enables or disables notification on task success.
		* `on_topology_modification` - Enables or disables notification on Environment Fleet Topology Modification.
		* `upcoming_schedule` - Enables notification on upcoming schedule.
			* `notify_before` - Specify when the notification should be sent. 
			* `on_upcoming_schedule` - Enables notification on upcoming schedule.
	* `topic_id` - Topic Id where the notifications will be directed. A topic is a communication channel for sending messages on chosen events to subscriptions. 
* `parent_fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet that would be the parent for this fleet. 
* `products` - Products associated with the Fleet.
* `properties` - Properties associated with the Fleet.
	* `compartment_id` - compartment OCID
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_property_type` - Type of the FleetProperty.
	* `is_required` - Property is required or not.
	* `value` - Value of the Property.
* `resource_region` - Associated region
* `resource_selection` - Resource Selection Type
	* `resource_selection_type` - Type of resource selection in a Fleet. Select resources manually or select resources based on rules. 
	* `rule_selection_criteria` - Rule Selection Criteria for DYNAMIC resource selection for a GENERIC fleet. Rules define what resources are members of this fleet. All resources that meet the criteria are added automatically. 
		* `match_condition` - Match condition for the rule selection. Include resources that match all rules or any of the rules. 
		* `rules` - Rules.
			* `basis` - Based on what the rule is created. It can be based on a resourceProperty or a tag.   If based on a tag, basis will be 'definedTagEquals' If based on a resource property, basis will be 'inventoryProperties' 
			* `compartment_id` - Compartment Id for which the rule is created. 
			* `conditions` - Rule Conditions
				* `attr_group` - Attribute Group. Provide a Tag namespace if the rule is based on a tag. Provide resource type if the rule is based on a resource property. 
				* `attr_key` - Attribute Key.Provide Tag key if the rule is based on a tag. Provide resource property name if the rule is based on a resource property. 
				* `attr_value` - Attribute Value.Provide Tag value if the rule is based on a tag. Provide resource property value if the rule is based on a resource property. 
			* `resource_compartment_id` - The Compartment ID to dynamically search resources. Provide the compartment ID to which the rule is applicable. 
* `resources` - Resources associated with the Fleet if resourceSelectionType is MANUAL.
	* `compartment_id` - Compartment Identifier[OCID].
	* `fleet_resource_type` - Type of the FleetResource.
	* `resource_id` - OCID of the resource.
	* `tenancy_id` - Tenancy Identifier[OCID].
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

