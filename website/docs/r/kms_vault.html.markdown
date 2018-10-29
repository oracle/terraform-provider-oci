---
layout: "oci"
page_title: "OCI: oci_kms_vault"
sidebar_current: "docs-oci-resource-kms-vault"
description: |-
  Creates and manages an OCI Vault
---

# oci_kms_vault
The `oci_kms_vault` resource creates and manages an OCI Vault

Creates a new vault. The type of vault you create determines key 
placement, pricing, and available options. Options include storage 
isolation, a dedicated service endpoint instead of a shared service
endpoint for API calls, and a dedicated HSM or a multitenant HSM.        


## Example Usage

```hcl
resource "oci_kms_vault" "test_vault" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.vault_display_name}"
	vault_type = "${var.vault_vault_type}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment where you want to create this vault.
* `display_name` - (Required) (Updatable) A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `vault_type` - (Required) The type of vault to create. Each type of vault stores the key with different degrees of isolation and has different options and pricing. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this vault.
* `crypto_endpoint` - The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. 
* `display_name` - A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `id` - The OCID of the vault.
* `management_endpoint` - The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. 
* `state` - The vault's current state.  Example: `DELETED` 
* `time_created` - The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property for the deletion time of the Vault expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `vault_type` - The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing.

## Import

Vaults can be imported using the `id`, e.g.

```
$ terraform import oci_kms_vault.test_vault "id"
```
