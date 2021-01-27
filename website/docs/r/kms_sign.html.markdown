---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_sign"
sidebar_current: "docs-oci-resource-kms-sign"
description: |-
  Provides the Sign resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_sign
This resource provides the Sign resource in Oracle Cloud Infrastructure Kms service.

Creates a digital signature for a message or message digest by using the private key in an asymmetric key. 
To verify the generated signature, you can use the Verify operation or use the public key in the same asymmetric key outside of KMS


## Example Usage

```hcl
resource "oci_kms_sign" "test_sign" {
	#Required
	crypto_endpoint = var.sign_message_crypto_endpoint
	key_id = oci_kms_key.test_key.id
	message = var.sign_message
	signing_algorithm = var.sign_signing_algorithm

	#Optional
	key_version_id = oci_kms_key_version.test_key_version.id
	message_type = var.sign_message_type
}
```

## Argument Reference

The following arguments are supported:

* `crypto_endpoint` - (Required) The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
* `key_id` - (Required) The OCID of the key used to sign the message
* `key_version_id` - (Optional) The OCID of the keyVersion used to sign the message.
* `message` - (Required) The Base64-encoded binary data object denoting the message or message digest to be signed. Message can be upto 4096 size in bytes. To sign a larger message, provide the message digest.
* `message_type` - (Optional) Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, RAW, indicates a message. To indicate a message digest, use DIGEST. 
* `signing_algorithm` - (Required) The algorithm to be used for signing the message or message digest For RSA keys, there are two supported Signature Schemes: PKCS1 and PSS along with  different Hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. In case of passing digest for signing, make sure the same hashing algorithm is  specified as used for created for digest.       


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `key_id` - The OCID of the key used to sign the message
* `key_version_id` - The OCID of the keyVersion used to sign the message
* `signature` - The Base64-encoded binary data object denoting the cryptographic signature that was generated for the message or message digest. 
* `signing_algorithm` - The algorithm to be used for signing the message or message digest For RSA keys, there are two supported Signature Schemes: PKCS1 and PSS along with  different Hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. In case of passing digest for signing, make sure the same hashing algorithm is  specified as used for created for digest.       

## Import

Sign can be imported using the `id`, e.g.

```
$ terraform import oci_kms_sign.test_sign "id"
```

