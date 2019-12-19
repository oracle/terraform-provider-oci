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

Encrypts data using the given [EncryptDataDetails](https://docs.cloud.oracle.com/iaas/api/#/en/key/release/datatypes/EncryptDataDetails) resource.
Plaintext included in the example request is a base64-encoded value of a UTF-8 string.


## Example Usage

```hcl
resource "oci_kms_encrypted_data" "test_encrypted_data" {
	#Required
	crypto_endpoint = "${var.encrypted_data_crypto_endpoint}"
	key_id = "${oci_kms_key.test_key.id}"
	plaintext = "${var.encrypted_data_plaintext}"

	#Optional
	associated_data = "${var.encrypted_data_associated_data}"
	logging_context = "${var.encrypted_data_logging_context}"
}
```

## Argument Reference

The following arguments are supported:

* `associated_data` - (Optional) Information that can be used to provide an encryption context for the encrypted data. The length of the string representation of the associated data must be fewer than 4096 characters. 
* `crypto_endpoint` - (Required) The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. see Vault Crypto endpoint.
* `key_id` - (Required) The OCID of the key to encrypt with.
* `logging_context` - (Optional) Information that provides context for audit logging. You can provide this additional data as key-value pairs to include in the audit logs when audit logging is enabled. 
* `plaintext` - (Required) The plaintext data to encrypt.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ciphertext` - The encrypted data.

## Import

Import is not supported for this resource.

