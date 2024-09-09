---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_credential"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_credential"
description: |-
  Provides details about a specific Fleet Credential in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_credential
This data source provides details about a specific Fleet Credential resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a FleetCredential by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_credential" "test_fleet_credential" {
	#Required
	fleet_credential_id = oci_fleet_apps_management_fleet_credential.test_fleet_credential.id
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_credential_id` - (Required) unique FleetCredential identifier
* `fleet_id` - (Required) unique Fleet identifier


## Attributes Reference

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

