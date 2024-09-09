---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_credential"
sidebar_current: "docs-oci-resource-fleet_apps_management-fleet_credential"
description: |-
  Provides the Fleet Credential resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_fleet_credential
This resource provides the Fleet Credential resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new FleetCredential.


## Example Usage

```hcl
resource "oci_fleet_apps_management_fleet_credential" "test_fleet_credential" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.fleet_credential_display_name
	entity_specifics {
		#Required
		credential_level = var.fleet_credential_entity_specifics_credential_level
		resource_id = oci_cloud_guard_resource.test_resource.id
		target = var.fleet_credential_entity_specifics_target
	}
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	password {
		#Required
		credential_type = var.fleet_credential_password_credential_type

		#Optional
		key_id = oci_kms_key.test_key.id
		key_version = var.fleet_credential_password_key_version
		secret_id = oci_vault_secret.test_secret.id
		secret_version = var.fleet_credential_password_secret_version
		value = var.fleet_credential_password_value
		vault_id = oci_kms_vault.test_vault.id
	}
	user {
		#Required
		credential_type = var.fleet_credential_user_credential_type

		#Optional
		key_id = oci_kms_key.test_key.id
		key_version = var.fleet_credential_user_key_version
		secret_id = oci_vault_secret.test_secret.id
		secret_version = var.fleet_credential_user_secret_version
		value = var.fleet_credential_user_value
		vault_id = oci_kms_vault.test_vault.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `entity_specifics` - (Required) (Updatable) Credential Details
	* `credential_level` - (Required) (Updatable) Credential Level.
	* `resource_id` - (Required) (Updatable) OCID of the resource associated with the target for which credential is created
	* `target` - (Required) (Updatable) Target associated with the Credential
* `fleet_id` - (Required) unique Fleet identifier
* `password` - (Required) (Updatable) Credential Details
	* `credential_type` - (Required) (Updatable) Credential Type
	* `key_id` - (Required when credential_type=KEY_ENCRYPTION) (Updatable) OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	* `key_version` - (Applicable when credential_type=KEY_ENCRYPTION) (Updatable) The Vault Key version.
	* `secret_id` - (Required when credential_type=VAULT_SECRET) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
	* `secret_version` - (Applicable when credential_type=VAULT_SECRET) (Updatable) The secret version.
	* `value` - (Required when credential_type=KEY_ENCRYPTION | PLAIN_TEXT) (Updatable) The value corresponding to the credential
	* `vault_id` - (Required when credential_type=KEY_ENCRYPTION) (Updatable) OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.
* `user` - (Required) (Updatable) Credential Details
	* `credential_type` - (Required) (Updatable) Credential Type
	* `key_id` - (Required when credential_type=KEY_ENCRYPTION) (Updatable) OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	* `key_version` - (Applicable when credential_type=KEY_ENCRYPTION) (Updatable) The Vault Key version.
	* `secret_id` - (Required when credential_type=VAULT_SECRET) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret.
	* `secret_version` - (Applicable when credential_type=VAULT_SECRET) (Updatable) The secret version.
	* `value` - (Required when credential_type=KEY_ENCRYPTION | PLAIN_TEXT) (Updatable) The value corresponding to the credential
	* `vault_id` - (Required when credential_type=KEY_ENCRYPTION) (Updatable) OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fleet Credential
	* `update` - (Defaults to 20 minutes), when updating the Fleet Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Fleet Credential


## Import

Import is not supported for this resource.

