---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet"
description: |-
  Provides details about a specific Fleet in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet
This data source provides details about a specific Fleet resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Get the details of a fleet in Fleet Application Management.

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet" "test_fleet" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) Unique Fleet identifier.


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

