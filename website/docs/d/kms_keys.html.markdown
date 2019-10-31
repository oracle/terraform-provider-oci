---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_keys"
sidebar_current: "docs-oci-datasource-kms-keys"
description: |-
  Provides the list of Keys in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_keys
This data source provides the list of Keys in Oracle Cloud Infrastructure Kms service.

Lists the keys in the specified vault and compartment.


## Example Usage

```hcl
data "oci_kms_keys" "test_keys" {
	#Required
	compartment_id = "${var.compartment_id}"
	management_endpoint = "${var.key_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.


## Attributes Reference

The following attributes are exported:

* `keys` - The list of keys.

### Key Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key.
* `current_key_version` - The OCID of the KeyVersion resource used in cryptographic operations. During key rotation, service might be in a transitional state where this or a newer KeyVersion are used intermittently. The currentKeyVersion field is updated when the service is guaranteed to use the new KeyVersion for all subsequent encryption operations. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "foo-value"}` 
* `display_name` - A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the key.
* `key_shape` - 
	* `algorithm` - The algorithm used by a key's KeyVersions to encrypt or decrypt.
	* `length` - The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `state` - The key's current state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

