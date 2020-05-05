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
	compartment_id = "${var.compartment_id}"
	display_name = "${var.key_display_name}"
	key_shape {
		#Required
		algorithm = "${var.key_key_shape_algorithm}"
		length = "${var.key_key_shape_length}"
	}
	management_endpoint = "${var.key_management_endpoint}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the master encryption key.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `desired_state` - (Optional) (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
* `display_name` - (Required) (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `key_shape` - (Required) 
	* `algorithm` - (Required) The algorithm used by a key's key versions to encrypt or decrypt.
	* `length` - (Required) The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
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
* `key_shape` - 
	* `algorithm` - The algorithm used by a key's key versions to encrypt or decrypt.
	* `length` - The length of the key, expressed as an integer. Values of 16, 24, or 32 are supported. 
* `restored_from_key_id` - The OCID of the key from which this key was restored.
* `state` - The key's current lifecycle state.  Example: `ENABLED` 
* `time_created` - The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key.

## Import

Keys can be imported using the `id`, e.g.

```
$ terraform import oci_kms_key.test_key "managementEndpoint/{managementEndpoint}/keys/{keyId}"
```

