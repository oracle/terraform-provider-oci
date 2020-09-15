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

Lists the master encryption keys in the specified vault and compartment.

As a management operation, this call is subject to a Key Management limit that applies to the total number
of requests across all management read operations. Key Management might throttle this call to reject an
otherwise valid request when the total rate of management read operations exceeds 10 requests per second
for a given tenancy.


## Example Usage

```hcl
data "oci_kms_keys" "test_keys" {
	#Required
	compartment_id = var.compartment_id
	management_endpoint = var.key_management_endpoint
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

* `compartment_id` - The OCID of the compartment that contains this master encryption key.
* `current_key_version` - The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The `currentKeyVersion` property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the key.
* `key_shape` - 
	* `algorithm` - The algorithm used by a key's key versions to encrypt or decrypt.
	* `length` - The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `state` - The key's current lifecycle state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

