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
		* `on_job_canceled` - Enables or disables notification on job canceled.
		* `on_job_failure` - Enables or disables notification on Job Failures.
		* `on_job_schedule_change` - Enables or disables notification on job schedule change.
		* `on_job_start` - Enables or disables notification on job start.
		* `on_job_success` - Enables or disables notification on job success.
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
* `products` - Products associated with the Fleet. PlatformConfiguration Ids corresponding to the Products. 
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
			* `compartment_id_in_subtree` - If set to true, resources will be returned for not only the provided compartment, but all compartments which descend from it. Which resources are returned and their field contents depends on the value of accessLevel. Default value for `compartmentIdInSubtree` is false 
			* `conditions` - Rule Conditions
				* `attr_group` - Attribute Group. Provide a Tag namespace if the rule is based on a tag. Provide resource type if the rule is based on a resource property. 
				* `attr_key` - Attribute Key.Provide Tag key if the rule is based on a tag. Provide resource property name if the rule is based on a resource property. 
				* `attr_value` - Attribute Value.Provide Tag value if the rule is based on a tag. Provide resource property value if the rule is based on a resource property. 
			* `match_condition` - Match condition for the rule selection. Include resources that match all rules or any of the rules. Default value for `matchCondition` is ANY 
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

