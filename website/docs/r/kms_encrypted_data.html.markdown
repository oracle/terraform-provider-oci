---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_encrypted_data"
sidebar_current: "docs-oci-resource-kms-encrypted_data"
description: |-
  Provides the Encrypted Data resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_encrypted_data
This resource provides the Encrypted Data resource in Oracle Cloud Infrastructure Kms service.

Encrypts data using the given [EncryptDataDetails](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/datatypes/EncryptDataDetails) resource.
Plaintext included in the example request is a base64-encoded value of a UTF-8 string.


## Example Usage

```hcl
resource "oci_kms_encrypted_data" "test_encrypted_data" {
	#Required
	crypto_endpoint = var.encrypted_data_crypto_endpoint
	key_id = oci_kms_key.test_key.id
	plaintext = var.encrypted_data_plaintext

	#Optional
	associated_data = var.encrypted_data_associated_data
	encryption_algorithm = var.encrypted_data_encryption_algorithm
	key_version_id = oci_kms_key_version.test_key_version.id
	logging_context = var.encrypted_data_logging_context
}
```

## Argument Reference

The following arguments are supported:

* `associated_data` - (Optional) Information that can be used to provide an encryption context for the encrypted data. The length of the string representation of the associated data must be fewer than 4096 characters. 
* `crypto_endpoint` - (Required) The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. see Vault Crypto endpoint.
* `encryption_algorithm` - (Optional) The encryption algorithm to use to encrypt and decrypt data with a customer-managed key. `AES_256_GCM` indicates that the key is a symmetric key that uses the Advanced Encryption Standard (AES) algorithm and  that the mode of encryption is the Galois/Counter Mode (GCM). `RSA_OAEP_SHA_1` indicates that the  key is an asymmetric key that uses the RSA encryption algorithm and uses Optimal Asymmetric Encryption Padding (OAEP).  `RSA_OAEP_SHA_256` indicates that the key is an asymmetric key that uses the RSA encryption algorithm with a SHA-256 hash  and uses OAEP. 
* `key_id` - (Required) The OCID of the key to encrypt with.
* `key_version_id` - (Optional) The OCID of the key version used to encrypt the ciphertext.
* `logging_context` - (Optional) Information that provides context for audit logging. You can provide this additional data as key-value pairs to include in the audit logs when audit logging is enabled. 
* `plaintext` - (Required) The plaintext data to encrypt.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ciphertext` - The encrypted data.
* `encryption_algorithm` - The encryption algorithm to use to encrypt and decrypt data with a customer-managed key. `AES_256_GCM` indicates that the key is a symmetric key that uses the Advanced Encryption Standard (AES) algorithm and  that the mode of encryption is the Galois/Counter Mode (GCM). `RSA_OAEP_SHA_1` indicates that the  key is an asymmetric key that uses the RSA encryption algorithm and uses Optimal Asymmetric Encryption Padding (OAEP).  `RSA_OAEP_SHA_256` indicates that the key is an asymmetric key that uses the RSA encryption algorithm with a SHA-256 hash  and uses OAEP.    
* `key_id` - The OCID of the key used to encrypt the ciphertext.
* `key_version_id` - The OCID of the key version used to encrypt the ciphertext.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Encrypted Data
	* `update` - (Defaults to 20 minutes), when updating the Encrypted Data
	* `delete` - (Defaults to 20 minutes), when destroying the Encrypted Data


## Import

Import is not supported for this resource.

