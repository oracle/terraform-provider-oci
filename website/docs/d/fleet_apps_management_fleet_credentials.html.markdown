---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_credentials"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_credentials"
description: |-
  Provides the list of Fleet Credentials in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_credentials
This data source provides the list of Fleet Credentials in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of FleetCredentials.


## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_credentials" "test_fleet_credentials" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id

	#Optional
	compartment_id = var.compartment_id
	credential_level = var.fleet_credential_credential_level
	display_name = var.fleet_credential_display_name
	id = var.fleet_credential_id
	state = var.fleet_credential_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `credential_level` - (Optional) Credential Level.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Required) unique Fleet identifier
* `id` - (Optional) unique FleetCredential identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `fleet_credential_collection` - The list of fleet_credential_collection.

### FleetCredential Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `entity_specifics` - Credential Details
	* `credential_level` - Credential Level.
	* `resource_id` - OCID of the resource associated with the target for which credential is created
	* `target` - Target associated with the Credential
* `id` - The unique id of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `password` - Credential Details
	* `credential_type` - Credential Type
	* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	* `key_version` - The Vault Key version.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
	* `secret_version` - The secret version.
	* `value` - The value corresponding to the credential
	* `vault_id` - OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.
* `state` - The current state of the FleetCredential.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `user` - Credential Details
	* `credential_type` - Credential Type
	* `key_id` - OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	* `key_version` - The Vault Key version.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
	* `secret_version` - The secret version.
	* `value` - The value corresponding to the credential
	* `vault_id` - OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.

