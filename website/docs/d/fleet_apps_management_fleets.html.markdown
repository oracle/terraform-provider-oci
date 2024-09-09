---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleets"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleets"
description: |-
  Provides the list of Fleets in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleets
This data source provides the list of Fleets in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of Fleets in the specified Tenancy.


## Example Usage

```hcl
data "oci_fleet_apps_management_fleets" "test_fleets" {

	#Optional
	application_type = var.fleet_application_type
	compartment_id = var.compartment_id
	display_name = var.fleet_display_name
	environment_type = var.fleet_environment_type
	fleet_type = var.fleet_fleet_type
	id = var.fleet_id
	product = var.fleet_product
	state = var.fleet_state
}
```

## Argument Reference

The following arguments are supported:

* `application_type` - (Optional) A filter to return only resources that match the Application Type given.
* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `environment_type` - (Optional) A filter to return only resources that match the Environment Type given.
* `fleet_type` - (Optional) A filter to return only resources their fleetType matches the given fleetType.
* `id` - (Optional) unique Fleet identifier
* `product` - (Optional) A filter to return only resources that match the Product Type given.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `fleet_collection` - The list of fleet_collection.

### Fleet Reference

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

