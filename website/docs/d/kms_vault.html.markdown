---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault"
sidebar_current: "docs-oci-datasource-kms-vault"
description: |-
  Provides details about a specific Vault in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_vault
This data source provides details about a specific Vault resource in Oracle Cloud Infrastructure Kms service.

Gets the specified vault's configuration information.

As a provisioning operation, this call is subject to a Key Management limit that applies to
the total number of requests across all provisioning read operations. Key Management might
throttle this call to reject an otherwise valid request when the total rate of provisioning
read operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
data "oci_kms_vault" "test_vault" {
	#Required
	vault_id = oci_kms_vault.test_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `vault_id` - (Required) The OCID of the vault.


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
* `restore_from_file` - Details where vault was backed up.
    * `content_length` - content length of vault's backup binary file
    * `content_md5` - content md5 hashed value of vault's backup file
    * `restore_vault_from_file_details` - Vault backup file content
* `restore_from_object_store` - Details where vault was backed up
    * `bucket` - Name of the bucket where vault was backed up
    * `destination` - Type of backup to restore from. Values of "BUCKET", "PRE_AUTHENTICATED_REQUEST_URI" are supported
    * `namespace` - Namespace of the bucket where vault was backed up
    * `object` - Object containing the backup
    * `uri` - Pre-authenticated-request-uri of the backup
* `restore_trigger` - When flipped, triggers restore if restore options are provided. Values of 0 or 1 are supported
* `time_created` - The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property to indicate when to delete the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `vault_type` - The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing.

