---
subcategory: "Kms"
layout: "oci"
page_title: "OCI: oci_kms_decrypted_data"
sidebar_current: "docs-oci-datasource-kms-decrypted_data"
description: |-
  Provides details about a specific DecryptedData
---

# Data Source: oci_kms_decrypted_data
The `oci_kms_decrypted_data` data source provides details about a specific DecryptedData

Decrypts data using the given DecryptDataDetails resource.


## Example Usage

```hcl
data "oci_kms_decrypted_data" "test_decrypted_data" {
	#Required
	ciphertext = var.decrypted_data_ciphertext
	crypto_endpoint = var.decrypted_data_crypto_endpoint
	key_id = oci_kms_key.test_key.id

	#Optional
	associated_data = var.decrypted_data_associated_data
}
```

## Argument Reference

The following arguments are supported:

* `associated_data` - (Optional) Information that can be used to provide an encryption context for the  encrypted data. The length of the string representation of the associatedData must be fewer than 4096 characters. 
* `ciphertext` - (Required) The encrypted data to decrypt.
* `crypto_endpoint` - (Required) The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. see Vault Crypto endpoint.
* `key_id` - (Required) The OCID of the key used to encrypt the ciphertext.


## Attributes Reference

The following attributes are exported:

* `plaintext` - The decrypted data, in the form of a base64-encoded value.
* `plaintext_checksum` - Checksum of the decrypted data.

