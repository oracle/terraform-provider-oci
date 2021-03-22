---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_key"
sidebar_current: "docs-oci-resource-kms-key"
description: |-
  Provides the Key resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_key
This resource provides the Key resource in Oracle Cloud Infrastructure Kms service.

Creates a new master encryption key.

As a management operation, this call is subject to a Key Management limit that applies to the total
number of requests across all management write operations. Key Management might throttle this call
to reject an otherwise valid request when the total rate of management write operations exceeds 10
requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_key" "test_key" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.key_display_name
	key_shape {
		#Required
		algorithm = var.key_key_shape_algorithm
		length = var.key_key_shape_length

		#Optional
		curve_id = oci_kms_curve.test_curve.id
	}
	management_endpoint = var.key_management_endpoint

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	protection_mode = "${var.key_protection_mode}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the master encryption key.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `desired_state` - (Optional) (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
* `display_name` - (Required) (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `key_shape` - (Required) The cryptographic properties of a key.
	* `algorithm` - (Required) The algorithm used by a key's key versions to encrypt or decrypt.
	* `curve_id` - (Optional) Supported curve IDs for ECDSA keys.
	* `length` - (Required) The length of the key in bytes, expressed as an integer. Supported values include the following:
		* AES: 16, 24, or 32
		* RSA: 256, 384, or 512
		* ECDSA: 32, 48, or 66 
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
* `protection_mode` - (Optional) The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists  on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,  a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported. 
* `restore_from_file` - (Optional) (Updatable) Details where key was backed up.
    * `content_length` - (Required) (Updatable) content length of key's backup binary file
    * `content_md5` - (Optional) (Updatable) content md5 hashed value of key's backup file
    * `restore_key_from_file_details` - (Required) Key backup file content.
* `restore_from_object_store` - (Optional) (Updatable) Details where key was backed up
    * `bucket` - (Optional) (Updatable) Name of the bucket where key was backed up
    * `destination` - (Required) (Updatable) Type of backup to restore from. Values of "BUCKET", "PRE_AUTHENTICATED_REQUEST_URI" are supported
    * `namespace` - (Optional) (Updatable) Namespace of the bucket where key was backed up
    * `object` - (Optional) (Updatable) Object containing the backup
    * `uri` - (Optional) (Updatable) Pre-authenticated-request-uri of the backup
* `restore_trigger` - (Optional) (Updatable) An optional property when flipped triggers restore from restore option provided in config file. 
* `time_of_deletion` - (Optional) (Updatable) An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this master encryption key.
* `current_key_version` - The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The `currentKeyVersion` property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the key.
* `is_primary` - A boolean that will be true when key is primary, and will be false when key is a replica from a primary key.
* `key_shape` - The cryptographic properties of a key.
	* `algorithm` - The algorithm used by a key's key versions to encrypt or decrypt.
	* `curve_id` - Supported curve IDs for ECDSA keys.
	* `length` - The length of the key in bytes, expressed as an integer. Supported values include the following:
		* AES: 16, 24, or 32
		* RSA: 256, 384, or 512
		* ECDSA: 32, 48, or 66 
* `protection_mode` - The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists  on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,  a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported. 
* `replica_details` - Key replica details 
	* `replication_id` - ReplicationId associated with a key operation 
* `restored_from_key_id` - The OCID of the key from which this key was restored.
* `state` - The key's current lifecycle state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Key
	* `update` - (Defaults to 20 minutes), when updating the Key
	* `delete` - (Defaults to 20 minutes), when destroying the Key


## Import

Keys can be imported using the `id`, e.g.

```
$ terraform import oci_kms_key.test_key "managementEndpoint/{managementEndpoint}/keys/{keyId}" 
```

