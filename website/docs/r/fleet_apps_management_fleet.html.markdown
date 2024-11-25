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

Create a product, environment, group, or generic type of fleet in Fleet Application Management.


## Example Usage

```hcl
resource "oci_fleet_apps_management_fleet" "test_fleet" {
	#Required
	compartment_id = var.compartment_id
	fleet_type = var.fleet_fleet_type

	#Optional
	application_type = var.fleet_application_type
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

* `application_type` - (Optional) Product stack associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `compartment_id` - (Required) Tenancy OCID
* `credentials` - (Optional) Credentials associated with the Fleet.
	* `compartment_id` - (Required) Tenancy OCID
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
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - (Optional) Environment Type associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `fleet_type` - (Required) Type of the Fleet. PRODUCT - A fleet of product-specific resources for a product type. ENVIRONMENT - A fleet of environment-specific resources for a product stack. GROUP - A fleet of a fleet of either environment or product fleets. GENERIC - A fleet of resources selected dynamically or manually for reporting purposes 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group_type` - (Optional) Group Type associated with Group Fleet. 
* `is_target_auto_confirm` - (Optional) (Updatable) A value that represents if auto-confirming of the targets can be enabled. This will allow targets to be auto-confirmed in the fleet without manual intervention. 
* `notification_preferences` - (Optional) (Updatable) Notification information to get notified when the fleet status changes.
	* `compartment_id` - (Required) (Updatable) Compartment ID the topic belongs to.
	* `preferences` - (Optional) (Updatable) Preferences to send notifications on the fleet activities.
		* `on_job_failure` - (Optional) (Updatable) Enables or disables notification on Job Failures.
		* `on_topology_modification` - (Optional) (Updatable) Enables or disables notification on Environment Fleet Topology Modification.
		* `on_upcoming_schedule` - (Optional) (Updatable) Enables notification on upcoming schedule.
	* `topic_id` - (Required) (Updatable) Topic Id where the notifications will be directed. A topic is a communication channel for sending messages on chosen events to subscriptions. 
* `products` - (Optional) Products associated with the Fleet.
* `resource_selection_type` - (Optional) Type of resource selection in a Fleet. Select resources manually or select resources based on rules. 
* `resources` - (Optional) Resources associated with the Fleet if resourceSelectionType is MANUAL.
	* `compartment_id` - (Required) Compartment Identifier[OCID].
	* `fleet_resource_type` - (Optional) Type of the FleetResource.
	* `resource_id` - (Required) OCID of the resource.
	* `tenancy_id` - (Required) Tenancy Identifier[OCID].
* `rule_selection_criteria` - (Optional) (Updatable) Rule Selection Criteria for DYNAMIC resource selection for a GENERIC fleet. Rules define what resources are members of this fleet. All resources that meet the criteria are added automatically. 
	* `match_condition` - (Required) (Updatable) Match condition for the rule selection. Include resources that match all rules or any of the rules. 
	* `rules` - (Required) (Updatable) Rules.
		* `basis` - (Optional) (Updatable) Based on what the rule is created. It can be based on a resourceProperty or a tag.   If based on a tag, basis will be 'definedTagEquals' If based on a resource property, basis will be 'inventoryProperties' 
		* `compartment_id` - (Required) (Updatable) Tenancy Id (Root Compartment Id)for which the rule is created. 
		* `conditions` - (Required) (Updatable) Rule Conditions
			* `attr_group` - (Required) (Updatable) Attribute Group. Provide a Tag namespace if the rule is based on a tag. Provide resource type if the rule is based on a resource property. 
			* `attr_key` - (Required) (Updatable) Attribute Key.Provide Tag key if the rule is based on a tag. Provide resource property name if the rule is based on a resource property. 
			* `attr_value` - (Required) (Updatable) Attribute Value.Provide Tag value if the rule is based on a tag. Provide resource property value if the rule is based on a resource property. 
		* `resource_compartment_id` - (Required) (Updatable) The Compartment ID to dynamically search resources. Provide the compartment ID to which the rule is applicable. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `application_type` - Product stack associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `compartment_id` - Tenancy OCID
* `credentials` - Credentials associated with the Fleet.
	* `compartment_id` - Tenancy OCID
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
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - Environment Type associated with the Fleet. Applicable for ENVIRONMENT fleet types. 
* `fleet_type` - Type of the Fleet. PRODUCT - A fleet of product-specific resources for a product type. ENVIRONMENT - A fleet of environment-specific resources for a product stack. GROUP - A fleet of a fleet of either environment or product fleets. GENERIC - A fleet of resources selected dynamically or manually for reporting purposes 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `group_type` - Group Type associated with Group Fleet. Applicable for GROUP fleet types. 
* `id` - The OCID of the resource.
* `is_target_auto_confirm` - A value that represents if auto-confirming of the targets can be enabled. This will allow targets to be auto-confirmed in the fleet without manual intervention. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `notification_preferences` - Notification information to get notified when the fleet status changes.
	* `compartment_id` - Compartment ID the topic belongs to.
	* `preferences` - Preferences to send notifications on the fleet activities.
		* `on_job_failure` - Enables or disables notification on Job Failures.
		* `on_topology_modification` - Enables or disables notification on Environment Fleet Topology Modification.
		* `on_upcoming_schedule` - Enables notification on upcoming schedule.
	* `topic_id` - Topic Id where the notifications will be directed. A topic is a communication channel for sending messages on chosen events to subscriptions. 
* `products` - Products associated with the Fleet.
* `properties` - Properties associated with the Fleet.
	* `compartment_id` - Tenancy OCID
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_property_type` - Type of the FleetProperty.
	* `is_required` - Property is required or not.
	* `value` - Value of the Property.
* `resource_region` - Associated region
* `resource_selection_type` - Type of resource selection in a Fleet. Select resources manually or select resources based on rules. 
* `resources` - Resources associated with the Fleet if resourceSelectionType is MANUAL.
	* `compartment_id` - Compartment Identifier[OCID].
	* `fleet_resource_type` - Type of the FleetResource.
	* `resource_id` - OCID of the resource.
	* `tenancy_id` - Tenancy Identifier[OCID].
* `rule_selection_criteria` - Rule Selection Criteria for DYNAMIC resource selection for a GENERIC fleet. Rules define what resources are members of this fleet. All resources that meet the criteria are added automatically. 
	* `match_condition` - Match condition for the rule selection. Include resources that match all rules or any of the rules. 
	* `rules` - Rules.
		* `basis` - Based on what the rule is created. It can be based on a resourceProperty or a tag.   If based on a tag, basis will be 'definedTagEquals' If based on a resource property, basis will be 'inventoryProperties' 
		* `compartment_id` - Tenancy Id (Root Compartment Id)for which the rule is created. 
		* `conditions` - Rule Conditions
			* `attr_group` - Attribute Group. Provide a Tag namespace if the rule is based on a tag. Provide resource type if the rule is based on a resource property. 
			* `attr_key` - Attribute Key.Provide Tag key if the rule is based on a tag. Provide resource property name if the rule is based on a resource property. 
			* `attr_value` - Attribute Value.Provide Tag value if the rule is based on a tag. Provide resource property value if the rule is based on a resource property. 
		* `resource_compartment_id` - The Compartment ID to dynamically search resources. Provide the compartment ID to which the rule is applicable. 
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

