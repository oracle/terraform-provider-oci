---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault"
sidebar_current: "docs-oci-resource-kms-vault"
description: |-
  Provides the Vault resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_vault
This resource provides the Vault resource in Oracle Cloud Infrastructure Kms service.

Creates a new vault. The type of vault you create determines key placement, pricing, and
available options. Options include storage isolation, a dedicated service endpoint instead
of a shared service endpoint for API calls, and either a dedicated hardware security module
(HSM) or a multitenant HSM.

As a provisioning operation, this call is subject to a Key Management limit that applies to
the total number of requests across all provisioning write operations. Key Management might
throttle this call to reject an otherwise valid request when the total rate of provisioning
write operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_vault" "test_vault" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vault_display_name
	vault_type = var.vault_vault_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	external_key_manager_metadata {
		#Required
		external_vault_endpoint_url = var.vault_external_key_manager_metadata_external_vault_endpoint_url
		oauth_metadata {
			#Required
			client_app_id = oci_kms_client_app.test_client_app.id
			client_app_secret = var.vault_external_key_manager_metadata_oauth_metadata_client_app_secret
			idcs_account_name_url = var.vault_external_key_manager_metadata_oauth_metadata_idcs_account_name_url
		}
		private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create this vault.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `external_key_manager_metadata` - (Optional) Metadata required for accessing External Key manager
	* `external_vault_endpoint_url` - (Required) URI of the vault on external key manager.
	* `oauth_metadata` - (Required) Authorization details required to get access token from IDP for accessing protected resources.
		* `client_app_id` - (Required) ID of the client app created in IDP.
		* `client_app_secret` - (Required) Secret of the client app created in IDP.
		* `idcs_account_name_url` - (Required) Base URL of the IDCS account where confidential client app is created.
	* `private_endpoint_id` - (Required) OCID of private endpoint created by customer.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `restore_from_file` - (Optional) (Updatable) Details where vault was backed up.
    * `content_length` - content length of vault's backup binary file
    * `content_md5` - (Optional) (Updatable) content md5 hashed value of vault's backup file
    * `restore_vault_from_file_details` - Vault backup file content
* `restore_from_object_store` - (Optional) (Updatable) Details where vault was backed up
    * `bucket` - (Optional) (Updatable) Name of the bucket where vault was backed up
    * `destination` - (Required) (Updatable) Type of backup to restore from. Values of "BUCKET", "PRE_AUTHENTICATED_REQUEST_URI" are supported
    * `namespace` - (Optional) (Updatable) Namespace of the bucket where vault was backed up
    * `object` - (Optional) (Updatable) Object containing the backup
    * `uri` - (Optional) (Updatable) Pre-authenticated-request-uri of the backup* `restore_trigger` - (Optional) (Updatable) An optional property when flipped triggers restore from restore option provided in config file. 
* `vault_type` - (Required) The type of vault to create. Each type of vault stores the key with different degrees of isolation and has different options and pricing. 
* `time_of_deletion` - (Optional) (Updatable) An optional property for the deletion time of the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains a particular vault.
* `crypto_endpoint` - The service endpoint to perform cryptographic operations against. Cryptographic operations include [Encrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/EncryptedData/Encrypt), [Decrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/DecryptedData/Decrypt), and [GenerateDataEncryptionKey](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/GeneratedKey/GenerateDataEncryptionKey) operations. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `external_key_manager_metadata_summary` - Summary about metadata of external key manager to be returned to the customer as a response.
	* `external_vault_endpoint_url` - URL of the vault on external key manager.
	* `oauth_metadata_summary` - Summary about authorization to be returned to the customer as a response.
		* `client_app_id` - ID of the client app created in IDP.
		* `idcs_account_name_url` - Base URL of the IDCS account where confidential client app is created.
	* `private_endpoint_id` - OCID of the private endpoint.
	* `vendor` - Vendor of the external key manager.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the vault.
* `is_primary` - A Boolean value that indicates whether the Vault is primary Vault or replica Vault.
* `is_vault_replicable` - A Boolean value that indicates whether the Vault has cross region replication capability. Always true for Virtual Private Vaults.
* `management_endpoint` - The service endpoint to perform management operations against. Management operations include "Create," "Update," "List," "Get," and "Delete" operations. 
* `replica_details` - Vault replica details 
	* `replication_id` - ReplicationId associated with a vault operation 
* `restored_from_vault_id` - The OCID of the vault from which this vault was restored, if it was restored from a backup file. If you restore a vault to the same region, the vault retains the same OCID that it had when you backed up the vault. 
* `state` - The vault's current lifecycle state.  Example: `DELETED` 
* `time_created` - The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property to indicate when to delete the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `vault_type` - The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vault
	* `update` - (Defaults to 20 minutes), when updating the Vault
	* `delete` - (Defaults to 20 minutes), when destroying the Vault


## Import

Vaults can be imported using the `id`, e.g.

```
$ terraform import oci_kms_vault.test_vault "id"
```

